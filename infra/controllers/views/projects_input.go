package views

type UserProjectInput struct {
	Email         string `json:"email" validate:"required,email"`
	AvatarUrl     string `json:"avatar_url" validate:"required,url"`
	LinkToProfile string `json:"link_to_profile" validate:"required,url"`
	Username      string `json:"username" validate:"required"`
}

type ProjectInput struct {
	Title             string           `json:"title" validate:"required"`
	Description       string           `json:"description" validate:"required"`
	LinkToSocialMedia string           `json:"link_to_social_media" validate:"required"`
	ProjectType       string           `json:"project_type" validate:"required"`
	Tecs              []string         `json:"tecs" validate:"required"`
	User              UserProjectInput `json:"user" validate:"required"`
}
