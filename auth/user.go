package auth

type User struct {
	ID         int64
	Name       string
	ExternalID string
	Email      string
	Avatar     string
}
