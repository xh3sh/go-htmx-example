package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xh3sh/go-htmx-example/internal/model"
)

type Handler struct {
	config *model.Config
}

func NewHandler(config *model.Config) *Handler {
	return &Handler{
		config: config,
	}
}

func (h *Handler) HandleHome(c echo.Context) error {
	tags := getAllTags(h.config.Projects)
	data := struct {
		Projects       []model.Project
		Tags           []string
		ActiveTags     []string
		Social         model.Social
		Languages      []model.SkillItem
		Infrastructure []model.SkillItem
		Technologies   []model.SkillItem
	}{
		Projects:       h.config.Projects,
		Tags:           tags,
		ActiveTags:     nil,
		Social:         h.config.Social,
		Languages:      h.config.Skills.Languages,
		Infrastructure: h.config.Skills.Infrastructure,
		Technologies:   h.config.Skills.Technologies,
	}
	return c.Render(http.StatusOK, "index", data)
}

func getAllTags(projects []model.Project) []string {
	tagSet := make(map[string]struct{})
	var tags []string
	for _, project := range projects {
		for _, tag := range project.Tags {
			if _, exists := tagSet[tag]; !exists {
				tagSet[tag] = struct{}{}
				tags = append(tags, tag)
			}
		}
	}

	return tags
}
