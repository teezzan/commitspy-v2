package webhook

import (
	"net/http"

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
	}

	newCommit := &account.Commit{
		ProjectID: projectCtx.ID,
	}

	for _, commit := range evtDataCtx.Commits {
		if commit.Distinct {
			newCommit.Number += 1
			newCommit.ExternalIDs = append(newCommit.ExternalIDs, commit.ExtID)
		}
	}

	err := database.CreateCommit(newCommit)
	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.WriteSuccess(c, http.StatusOK, gin.H{})
}
