package model

// Request structure
type Request struct {
	Input string `json:"input" binding:"required"`
}

// Response structure
type Response struct {
	Output string `json:"output"`
}
