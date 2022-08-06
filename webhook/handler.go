package webhook

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy-v2/account"
	"github.com/teezzan/commitspy-v2/database"
	"github.com/teezzan/commitspy-v2/response"
)

type EventHandlers struct{}

func (ctrl EventHandlers) Github(c *gin.Context) {

	projectCtx, _ := ProjectFromCtx(c)
	evtDataCtx, _ := GithubEventDataFromCtx(c)

	evtType := c.GetHeader("x-github-event")
	if evtType == "ping" {
		projectCtx.ExternalID = evtDataCtx.RepositoryExtID
		err := database.UpdateProject(projectCtx)
		if err != nil {
			response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response.WriteSuccess(c, http.StatusAccepted, gin.H{})
		return
	} else if evtType == "push" {

		newCommit := &account.Commit{
			ProjectID: projectCtx.ID,
		}

		for _, commit := range evtDataCtx.Commits {
			if commit.Distinct {
				newCommit.Number += 1
			}
		}

		err := database.CreateCommit(newCommit)
		if err != nil {
			response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response.WriteSuccess(c, http.StatusOK, gin.H{})
	}

}

func (ctrl EventHandlers) Gitlab(c *gin.Context) {

	projectCtx, _ := ProjectFromCtx(c)
	evtDataCtx, _ := GitlabEventDataFromCtx(c)
	commitCount, err := database.CountCommitsByProjectUUID(fmt.Sprint(projectCtx.ID))
	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if *commitCount == 0 && projectCtx.ExternalID == "" {
		projectCtx.ExternalID = evtDataCtx.RepositoryExtID
		err := database.UpdateProject(projectCtx)
		if err != nil {
			response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response.WriteSuccess(c, http.StatusAccepted, gin.H{})
		return
	} else {

		newCommit := &account.Commit{
			ProjectID: projectCtx.ID,
		}

		for _, commit := range evtDataCtx.Commits {
			isMergeCommit :=
				strings.Contains(commit.Message, "Merge branch '") &&
					strings.Contains(commit.Message, "' into '") &&
					strings.Contains(commit.Message, "See merge request ")
			if !isMergeCommit {
				newCommit.Number += 1
			}
		}

		err := database.CreateCommit(newCommit)
		if err != nil {
			response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response.WriteSuccess(c, http.StatusOK, gin.H{})
	}

}
