package dto

type ListResponse struct {
    Offset        int                       `json:"offset"`
    Limit         int                       `json:"limit"`
    Total         int                       `json:"total"`
    Data          interface{}               `json:"data"`
}
