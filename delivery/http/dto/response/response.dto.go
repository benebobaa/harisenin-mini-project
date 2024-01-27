package response

type ResponseErrorDTO struct {
    StatusCode int `json:"status_code"`
    Error      any `json:"error"`
}
