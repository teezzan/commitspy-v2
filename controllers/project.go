package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy-v2/account"
	"github.com/teezzan/commitspy-v2/auth"
	"github.com/teezzan/commitspy-v2/database"
	"github.com/teezzan/commitspy-v2/response"
	"github.com/teezzan/commitspy-v2/validator"
)

type Project struct{}

func (ctrl Project) Create(c *gin.Context) {
	userCtx, _ := auth.UserFromCtx(c)
	var json validator.CreateProject

	if err := c.ShouldBindJSON(&json); err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	project, err := database.GetUserProjectByNameOrURL(userCtx.ID, json.Name, json.URL)

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if project != nil {
		response.WriteSuccess(c, http.StatusConflict,
			gin.H{"error": fmt.Sprintf("project with name: %s or url: %s exists", json.Name, json.URL)})
		return
	}

	newProject := &account.Project{
		Name:   json.Name,
		URL:    json.URL,
		Type:   json.Type,
		UserID: userCtx.ID,
	}

	err = database.CreateProject(newProject)
	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.WriteSuccess(c, http.StatusCreated, gin.H{"project": newProject})

}
func (ctrl Project) FetchOne(c *gin.Context) {
	userCtx, _ := auth.UserFromCtx(c)

	var json validator.URIProjectID

	if err := c.ShouldBindUri(&json); err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	project, err := database.GetUserProjectById(userCtx.ID, json.ProjectID)

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if project == nil {
		response.WriteSuccess(c, http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}

	response.WriteSuccess(c, http.StatusOK, gin.H{"project": project})

}
func (ctrl Project) FetchAll(c *gin.Context) {
	userCtx, _ := auth.UserFromCtx(c)

	projects, err := database.GetUserProjects(userCtx.ID)

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var r gin.H
	if projects == nil {
		r = gin.H{"projects": make([]string, 0)}
	} else {
		r = gin.H{"projects": projects}
	}

	response.WriteSuccess(c, http.StatusOK, r)

}
func (ctrl Project) Update(c *gin.Context) {
	userCtx, _ := auth.UserFromCtx(c)

	var json validator.UpdateProject

	if err := c.ShouldBindUri(&json); err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	project, err := database.GetUserProjectById(userCtx.ID, json.ProjectID)

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if project == nil {
		response.WriteSuccess(c, http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}

	if json.Name != nil {
		project.Name = *json.Name
	}
	if json.CommitGoal != nil {
		project.CommitGoal = *json.CommitGoal
	}
	if json.CommitTimeWindow != nil {
		project.CommitTimeWindow = *json.CommitTimeWindow
		nextAlarm := time.Now().Local().Add(time.Hour * time.Duration(*json.CommitTimeWindow))
		project.CommitDeadline = &nextAlarm
	}
	err = database.UpdateProject(project)
	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.WriteSuccess(c, http.StatusAccepted, gin.H{"project": project})

}
func (ctrl Project) Delete(c *gin.Context) {
	userCtx, _ := auth.UserFromCtx(c)

	var json validator.URIProjectID

	if err := c.ShouldBindUri(&json); err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	project, err := database.GetUserProjectById(userCtx.ID, json.ProjectID)

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if project == nil {
		response.WriteSuccess(c, http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}

	err = database.DeleteProject(project)
	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.WriteSuccess(c, http.StatusOK, gin.H{})
}
