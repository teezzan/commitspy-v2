package validator

type CreateProject struct {
	URL  string `json:"url" binding:"required|url"`
	Name string `json:"name" binding:"required"`
	Type int16  `json:"type" binding:"required"`
}
