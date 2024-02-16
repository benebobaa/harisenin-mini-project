package upload_image

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/benebobaa/harisenin-mini-project/domain/entity"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"net/http"
)

type AwsS3 struct {
	S3        *s3.S3
	AwsBucket string
}

func NewAwsS3(s3 *s3.S3, awsBucket string) AwsS3Action {
	return &AwsS3{S3: s3, AwsBucket: awsBucket}
}

type AwsS3Action interface {
	UploadFile(file *multipart.FileHeader) (entity.Image, error)
	//ValidateImageType(contentType string) bool
}

func (a *AwsS3) UploadFile(file *multipart.FileHeader) (entity.Image, error) {
	src, err := file.Open()
	if err != nil {
		return entity.Image{}, err
	}
	defer src.Close()

	// Read the file content into a byte slice
	fileBytes := make([]byte, file.Size)
	_, err = src.Read(fileBytes)
	if err != nil {
		return entity.Image{}, err
	}

	// Resize the image
	resizedBytes, err := a.resizeImageBytes(fileBytes, 800, 400) // Adjust dimensions as needed
	if err != nil {
		return entity.Image{}, err
	}

	// Specify the destination path in the S3 bucket (e.g., "uploads/")
	objectKey := "uploads/" + file.Filename

	// Get file information for content type and size
	fileType := http.DetectContentType(resizedBytes)
	size := int64(len(resizedBytes))

	if !a.validateImageType(fileType) {
		return entity.Image{}, fmt.Errorf("invalid image type")
	}

	// Create an io.Reader from the resizedBytes
	fileReader := bytes.NewReader(resizedBytes)

	// Upload the resized file to S3
	_, err = a.S3.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(a.AwsBucket),
		Key:           aws.String(objectKey),
		ACL:           aws.String("public-read"),
		Body:          fileReader,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	})

	if err != nil {
		return entity.Image{}, err
	}

	newImage := entity.Image{
		ImageType: fileType,
		ImageUrl:  fmt.Sprintf("https://%s.s3.amazonaws.com/%s", a.AwsBucket, objectKey),
	}

	return newImage, nil
}

func (a *AwsS3) validateImageType(contentType string) bool {
	// Add more allowed content types if needed
	allowedTypes := []string{"image/jpeg", "image/png", "image/gif"}
	for _, t := range allowedTypes {
		if t == contentType {
			return true
		}
	}
	return false
}

// ResizeImageBytes resizes the image represented as bytes to the specified width and height
func (a *AwsS3) resizeImageBytes(imageBytes []byte, width, height uint) ([]byte, error) {
	// Decode the original image
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}

	// Resize the image (adjust dimensions as needed)
	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	// Encode the resized image
	var resizedBuffer bytes.Buffer
	ext := a.detectImageFormat(imageBytes)
	switch ext {
	case "png":
		err = png.Encode(&resizedBuffer, resizedImg)
	case "jpeg":
		err = jpeg.Encode(&resizedBuffer, resizedImg, nil)
	default:
		return nil, fmt.Errorf("unsupported image format: %s", ext)
	}
	if err != nil {
		return nil, err
	}

	return resizedBuffer.Bytes(), nil
}

// detectImageFormat detects the image format based on the provided image bytes
func (a *AwsS3) detectImageFormat(imageBytes []byte) string {
	contentType := http.DetectContentType(imageBytes)
	switch contentType {
	case "image/jpeg":
		return "jpeg"
	case "image/png":
		return "png"
	// Add additional cases for other supported formats
	default:
		return ""
	}
}
