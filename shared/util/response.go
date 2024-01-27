package util

import "github.com/benebobaa/harisenin-mini-project/delivery/http/dto/response"

func ConstructResponseError(statusCode int, errorMsg string) (response.ResponseErrorDTO, int) {
	resp := response.ResponseErrorDTO{
		StatusCode: statusCode,
		Error:      errorMsg,
	}

	return resp, statusCode
}
