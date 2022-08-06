package validator

type URIProjectUUID struct {
	ProjectUUID string `uri:"uuid" binding:"required"`
}
