package validator

type CreateProject struct {
	URL  string `json:"url" binding:"required,url"`
	Name string `json:"name" binding:"required"`
	Type int16  `json:"type" binding:"required"`
}

type UpdateProject struct {
	ProjectID        int64  `uri:"id" binding:"required"`
	Name             *string `json:"name" binding:"omitempty"`
	CommitGoal       *int64  `json:"commit_goal" binding:"omitempty,min=1"`
	CommitTimeWindow *int64  `json:"commit_time_window" binding:"omitempty,min=24"`
}
