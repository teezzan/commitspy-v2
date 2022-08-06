package webhook

type GithubEventData struct {
	RepositoryExtID string
	Ref             string
	Commits         []GithubCommit
}

type GithubCommit struct {
	Distinct bool   `json:"distinct"`
	Message  string `json:"message"`
	ExtID    string `json:"id"`
}

type GitlabEventData struct {
	RepositoryExtID string
	Ref             string
	Commits         []GitlabCommit
}

type GitlabCommit struct {
	Message string `json:"message"`
	ExtID   string `json:"id"`
}
