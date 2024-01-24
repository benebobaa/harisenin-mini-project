package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v3/middleware/logger"

	"github.com/benebobaa/harisenin-mini-project/helper"
	"github.com/benebobaa/harisenin-mini-project/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/lib/pq"
	"mime/multipart"
	"path/filepath"
	"time"
)

var db *sql.DB

func initDB(dbDriver string, dbSource string) (*sql.DB, error) {

	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

type PostData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Image struct {
	ID       int
	PostID   int
	ImageURL string
	Filename string
	// Add other image-related fields as needed
}

func main() {
	fmt.Println("Hello World")
	configViper, err := utils.LoadConfig(".")
	helper.PanicIfError(err)

	db, err := initDB(configViper.DBDriver, configViper.DBSource)
	helper.PanicIfError(err)

	app := fiber.New(
		fiber.Config{
			IdleTimeout: time.Second * 5,
		})

	app.Use(logger.New())

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/ping", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/api/user", func(c fiber.Ctx) error {
		// Parse JSON body
		var postData PostData
		if err := json.Unmarshal(c.Body(), &postData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
		}

		// Insert data into the database
		SQL := `INSERT INTO "user" (username, password) VALUES ($1, $2)`
		_, err := db.Exec(SQL, postData.Username, postData.Password)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": pqErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Data inserted successfully"})
	})

	//upload test s3
	// Initialize AWS S3 client
	awsRegion := "us-east-1"

	appCreds := credentials.NewStaticCredentialsProvider(
		"AKIA3P6TUPNEM3ABS3OL",
		"rgUx0ZOak8vtpsSs2iCZj14ZSySMMQKb+D2rqsdO",
		"",
	)

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(appCreds),
	)
	helper.PanicIfError(err)

	s3Client := s3.NewFromConfig(cfg)

	app.Post("/api/upload", func(c fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}

		var image Image

		files := form.File["file"]
		for _, file := range files {
			imageURL, err := uploadToS3(file, s3Client, &image)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}

			image.ImageURL = imageURL
			fmt.Println("Save to database")
			err = saveToDatabase(image, db)
			fmt.Println("Save to database end")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
		}

		return c.SendString("Upload successful!")
	})

	app.Get("/images/:filename", func(c fiber.Ctx) error {
		filename := c.Params("filename")

		// Retrieve S3 URL from the database
		s3URL, err := getImageMetadata(db, filename)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Redirect to the S3 URL
		//err = c.Redirect().To(s3URL)
		//if err != nil {
		//	return err
		//}

		return c.SendFile(s3URL)
	})

	err = app.Listen(configViper.ServerAddress)
	helper.PanicIfError(err)
}

func uploadToS3(file *multipart.FileHeader, s3CLient *s3.Client, image *Image) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	fileName := generateFileName(file.Filename)
	image.Filename = fileName

	// Specify your S3 bucket name and folder (prefix) if needed
	s3BucketName := "tweets-harisenin-bucket"
	//uploader := s3manager.NewUploader(s3CLient)
	//_, err = uploader.Upload(&s3manager.UploadInput{
	//	Bucket: aws.String(s3BucketName),
	//	Key:    aws.String(fileName),
	//	Body:   src,
	//	ACL:    aws.String("public-read"), // Set appropriate ACL based on your requirements
	//})

	_, err = s3CLient.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &s3BucketName,
		Key:    &fileName,
		Body:   src,
		ACL:    "public-read", // Set appropriate ACL based on your requirements
	})

	if err != nil {
		return "", err
	}

	// Generate the S3 URL
	s3URL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s3BucketName, fileName)

	return s3URL, nil
}

func saveToDatabase(image Image, db *sql.DB) error {
	// Insert image data into PostgreSQL database
	fmt.Println("IMAGE", image)
	fmt.Println("IMAGE", image.ImageURL)
	SQL := `INSERT INTO "images"(image_url, filename) VALUES($1, $2)`
	_, err := db.Exec(SQL, image.ImageURL, image.Filename)
	fmt.Println("Error", err)
	return err
}

func generateFileName(originalName string) string {
	// Generate a unique filename based on timestamp and original filename
	extension := filepath.Ext(originalName)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)
	return fileName
}

func getImageMetadata(db *sql.DB, filename string) (string, error) {
	var s3URL string
	err := db.QueryRow(`SELECT "image_url" FROM images WHERE filename = $1`, filename).Scan(&s3URL)
	return s3URL, err
}
