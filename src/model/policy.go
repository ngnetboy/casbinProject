package model

type PolicyRequest struct {
	Role   string `json:"role"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
