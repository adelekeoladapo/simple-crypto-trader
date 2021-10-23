package dto

type ServiceResponse struct {
    Status        string                    `json:"status"`
    Message       string                    `json:"message"`
    Data          interface{}               `json:"data"`
}
