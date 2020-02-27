package model

// Contributor 贡献者
type Contributor struct {
	ID       int64  `json:"id"`
	GithubID string `json:"github_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Link     string `json:"link"`
	Avatar   string `json:"avatar"`
}

// TableName ...
func (Contributor) TableName() string {
	return "contributor"
}
