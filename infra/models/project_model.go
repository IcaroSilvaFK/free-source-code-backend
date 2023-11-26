package models

import "time"

type UserProject struct {
	Email         string `json:"email"`
	AvatarUrl     string `json:"avatar_url"`
	LinkToProfile string `json:"link_to_profile"`
	Username      string `json:"username"`
}

type ProjectModel struct {
	ProjectName        string      `json:"project_name"`
	ProjectDescription string      `json:"project_description"`
	LinkToSocialMedia  string      `json:"link_to_social_media"`
	ID                 string      `json:"id"`
	CreatedAt          time.Time   `json:"created_at"`
	ProjectType        string      `json:"project_type"`
	Tecs               []string    `json:"tecs"`
	Stack              string      `json:"stack"`
	User               UserProject `json:"user"`
}
