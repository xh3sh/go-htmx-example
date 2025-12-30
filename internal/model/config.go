package model

type Config struct {
	Server   Server    `yaml:"server"`
	Projects []Project `yaml:"projects"`
	Social   Social    `yaml:"social"`
	Skills   Skills    `yaml:"skills"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Social struct {
	Name        string `yaml:"name"`
	IconURL     string `yaml:"icon_url"`
	Description string `yaml:"description"`
	GitHubURL   string `yaml:"github_url"`
	ContactURL  string `yaml:"contact_url"`
}

type Skills struct {
	Languages      []string `yaml:"languages"`
	Infrastructure []string `yaml:"infrastructure"`
	Technologies   []string `yaml:"technologies"`
}
