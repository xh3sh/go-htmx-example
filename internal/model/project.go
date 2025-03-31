package model

type Project struct {
	ImageURL    string   `yaml:"image_url"`
	ProjectURL  string   `yaml:"project_url"` // Add ProjectURL field
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Link        string   `yaml:"link"`
	Tags        []string `yaml:"tags"` // Add Tags field
}
