package github

const (
	UrlBase      = "https://api.github.com"
	HeaderAccept = "application/vnd.github.v3+json"
)

type GithubAdapter struct{}

type Organization struct {
	Login       string `json:"login"`
	ID          int    `json:"id"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type Repository struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Private     bool   `json:"private"`
	Url         string `json:"url"`
	GitUrl      string `json:"git_url"`
	CloneUrl    string `json:"clone_url"`
	Description string `json:"description"`
}
