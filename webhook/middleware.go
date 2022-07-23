package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy-v2/database"
	"github.com/teezzan/commitspy-v2/response"
	"github.com/teezzan/commitspy-v2/validator"
)

func AuthenticateGithubWebhook(c *gin.Context) {

	var json validator.URIProjectUUID

	if err := c.ShouldBindUri(&json); err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sha256Signature := fetchSHA256Token(c)
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	isValidRequest := requestSHA256Validator(json.ProjectUUID, jsonData, sha256Signature)

	if !isValidRequest {
		response.WriteError(c, http.StatusForbidden, gin.H{"error": "sha256 keys do not match"})
		return
	}

	project, err := database.GetProjectByUUID(json.ProjectUUID)

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if project == nil {
		response.WriteSuccess(c, http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}

	c.Set("Project", &project)
	c.Next()
}

func requestSHA256Validator(token string, body []byte, requestSignature string) bool {
	key := []byte(token)

	sig := hmac.New(sha256.New, key)
	sig.Write(body)
	sumInHexString := fmt.Sprintf("%x", hex.EncodeToString(sig.Sum(nil)))
	return sumInHexString == requestSignature
}

func fetchSHA256Token(c *gin.Context) string {
	s := c.GetHeader("x-hub-signature-256")
	sha256Signature := strings.TrimSpace(strings.Replace(s, "sha256=", "", 1))
	return sha256Signature
}
