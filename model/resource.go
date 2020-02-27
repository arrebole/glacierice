package model

import "time"

// ResourceGithub ...
type ResourceGithub struct {
	ID            int64     `json:"id"`
	ResourceType  string    `json:"resource_type"`
	ResourceName  string    `json:"resource_name"`
	ContributorID string    `json:"contributor_id"`
	InitialTime   time.Time `json:"initial_time"`
	UpdateTime    time.Time `json:"update_time"`
	LinkGithub    string    `json:"link_github"`
	LinkReadme    string    `json:"link_readme"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
}

// TableName ...
func (ResourceGithub) TableName() string {
	return "resources_github"
}

// Resource3RD ...
type Resource3RD struct {
	ID              int64     `json:"id"`
	ResourceType    string    `json:"resource_type"`
	ResourceSubType string    `json:"resource_sub_type"`
	ResourceName    string    `json:"resource_name"`
	ContributorName string    `json:"contributor_name"`
	InitialTime     time.Time `json:"initial_time"`
	Link            string    `json:"link"`
	LinkType        string    `json:"link_type"`
	Description     string    `json:"description"`
	LanguageFlag    string    `json:"language_flag"`
}

// TableName ...
func (Resource3RD) TableName() string {
	return "resources_3rd"
}
