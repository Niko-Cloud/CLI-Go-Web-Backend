package domain

type Showcase struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	Summary string   `json:"summary"`
	Stack   []string `json:"stack"`
	RepoURL string   `json:"repoUrl"`
	LiveURL *string  `json:"liveUrl,omitempty"`
}
