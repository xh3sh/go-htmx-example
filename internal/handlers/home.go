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
		Projects   []model.Project
		Tags       []string
		ActiveTags []string
		Social     model.Social
	}{
		Projects:   h.config.Projects,
		Tags:       tags,
		ActiveTags: nil,
		Social:     h.config.Social,
	}
	return c.Render(http.StatusOK, "index", data)
}

// Helper functions
func getAllTags(projects []model.Project) []string {
	tagSet := make(map[string]struct{})
	for _, project := range projects {
		for _, tag := range project.Tags {
			tagSet[tag] = struct{}{}
		}
	}
	tags := make([]string, 0, len(tagSet))
	for tag := range tagSet {
		tags = append(tags, tag)
	}
	return tags
}
