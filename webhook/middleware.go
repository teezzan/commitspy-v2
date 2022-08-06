package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy-v2/account"
	"github.com/teezzan/commitspy-v2/database"
	"github.com/teezzan/commitspy-v2/response"
	"github.com/teezzan/commitspy-v2/validator"
	"github.com/tidwall/gjson"
)

func AuthenticateGithubWebhook(c *gin.Context) {

	var jsonPayload validator.URIProjectUUID

	if err := c.ShouldBindUri(&jsonPayload); err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sha256Signature := fetchSHA256Token(c)
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	jsonString := string(jsonData)
	isValidRequest := requestSHA256Validator(jsonPayload.ProjectUUID, jsonString, sha256Signature)

	if !isValidRequest {
		response.WriteError(c, http.StatusForbidden, gin.H{"error": "sha256 keys do not match"})
		return
	}

	evtData, err := parseGithubPayload(&jsonString, c.GetHeader("x-github-event"))

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Set("evtData", evtData)

	project, err := database.GetProjectByUUID(jsonPayload.ProjectUUID)

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if project == nil {
		response.WriteSuccess(c, http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}

	c.Set("Project", project)

	c.Next()
}

func requestSHA256Validator(token string, body string, requestSignature string) bool {
	h := hmac.New(sha256.New, []byte(token))
	h.Write([]byte(body))
	shaVal := hex.EncodeToString(h.Sum(nil))
	return shaVal == requestSignature
}

func fetchSHA256Token(c *gin.Context) string {
	s := c.GetHeader("x-hub-signature-256")
	sha256Signature := strings.TrimSpace(strings.Replace(s, "sha256=", "", 1))
	return sha256Signature
}

func ProjectFromCtx(c *gin.Context) (*account.Project, bool) {
	pCtx, ok := c.Get("Project")
	return pCtx.(*account.Project), ok
}

func GithubEventDataFromCtx(c *gin.Context) (*GithubEventData, bool) {
	evtCtx, ok := c.Get("evtData")
	return evtCtx.(*GithubEventData), ok
}

func parseGithubPayload(jsonBody *string, evtType string) (*GithubEventData, error) {

	evtData := GithubEventData{
		RepositoryExtID: gjson.Get(*jsonBody, "repository.id").String(),
	}

	if evtType == "push" {
		evtData.Ref = gjson.Get(*jsonBody, "ref").String()
		c := gjson.Get(*jsonBody, "commits").String()
		var commits []GithubCommit

		if err := json.Unmarshal([]byte(c), &commits); err != nil {
			return nil, err
		}

		evtData.Commits = commits
	}
	return &evtData, nil
}

func AuthenticateGitlabWebhook(c *gin.Context) {

	var jsonPayload validator.URIProjectUUID

	if err := c.ShouldBindUri(&jsonPayload); err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if jsonPayload.ProjectUUID != c.GetHeader("x-gitlab-token") {
		response.WriteError(c, http.StatusForbidden, gin.H{"error": "secret key do not match"})
		return
	}

	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	jsonString := string(jsonData)

	evtData, err := parseGitlabPayload(&jsonString, c.GetHeader("x-gitlab-event"))

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Set("evtData", evtData)

	project, err := database.GetProjectByUUID(jsonPayload.ProjectUUID)

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if project == nil {
		response.WriteSuccess(c, http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}

	c.Set("Project", project)

	c.Next()
}

func parseGitlabPayload(jsonBody *string, evtType string) (*GitlabEventData, error) {

	evtData := GitlabEventData{
		RepositoryExtID: gjson.Get(*jsonBody, "project_id").String(),
	}

	if evtType == "Push Hook" {
		evtData.Ref = gjson.Get(*jsonBody, "ref").String()
		c := gjson.Get(*jsonBody, "commits").String()
		var commits []GitlabCommit

		if err := json.Unmarshal([]byte(c), &commits); err != nil {
			return nil, err
		}

		evtData.Commits = commits
	}
	return &evtData, nil
}