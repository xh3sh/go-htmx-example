package handlers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/xh3sh/go-htmx-example/internal/model"
)

func (h *Handler) HandleFilter(c echo.Context) error {
	tagsParam := c.QueryParam("tags")

	seenTags := make(map[string]bool)
	var activeTags []string
	if tagsParam != "" {
		for _, t := range strings.Split(tagsParam, ",") {
			trimmed := strings.TrimSpace(t)
			if trimmed != "" && !seenTags[trimmed] {
				activeTags = append(activeTags, trimmed)
				seenTags[trimmed] = true
			}
		}
	}

	var filteredProjects []model.Project
	if len(activeTags) == 0 {
		filteredProjects = h.config.Projects
	} else {
		filteredProjects = filterProjectsByTags(h.config.Projects, activeTags)
	}

	data := struct {
		Projects   []model.Project
		ActiveTags []string
		Tags       []string
	}{
		Projects:   filteredProjects,
		ActiveTags: activeTags,
		Tags:       getAllTags(h.config.Projects),
	}

	return c.Render(http.StatusOK, "cards", data)
}

func filterProjectsByTags(projects []model.Project, tags []string) []model.Project {
	var filtered []model.Project
	for _, project := range projects {
		hasAllTags := true
		for _, targetTag := range tags {
			hasTag := false
			for _, projectTag := range project.Tags {
				if projectTag == targetTag {
					hasTag = true
					break
				}
			}
			if !hasTag {
				hasAllTags = false
				break
			}
		}
		if hasAllTags {
			filtered = append(filtered, project)
		}
	}
	return filtered
}
