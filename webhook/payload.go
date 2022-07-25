package webhook

type GithubEventData struct {
	RepositoryExtID string
	Ref             string
	Commits         []GithubCommit
}

type GithubCommit struct {
	Distinct bool   `json:"distinct"`
	Message  string `json:"message"`
	ExtID  string `json:"id"`
}
