package dto

type ListRequest struct {
    Offset        int                       `json:"offset" default:"1"`
    Limit         int                       `json:"limit" default:"10"`
    Filter        string                    `json:"filter"`
    SortField     string                    `json:"sortField" default:"id"`
    SortOrder     string                    `json:"sortOrder" default:"desc"`
}
