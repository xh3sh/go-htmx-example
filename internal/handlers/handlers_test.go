package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/xh3sh/go-htmx-example/internal/model"
)

type testRenderer struct {
	name string
	data interface{}
}

func (t *testRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	t.name = name
	t.data = data
	return nil
}

func TestGetAllTags(t *testing.T) {
	projects := []model.Project{
		{Tags: []string{"go", "web"}},
		{Tags: []string{"web", "docker"}},
	}
	got := getAllTags(projects)
	want := []string{"go", "web", "docker"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("getAllTags = %v, ожидается %v", got, want)
	}
}

func TestHandleHome(t *testing.T) {
	cfg := &model.Config{
		Projects: []model.Project{
			{Tags: []string{"go", "web"}},
			{Tags: []string{"web", "docker"}},
		},
		Social: model.Social{Name: "me"},
		Skills: model.Skills{
			Languages:      []string{"Go"},
			Infrastructure: []string{"Docker"},
			Technologies:   []string{"HTMX"},
		},
	}

	e := echo.New()
	tr := &testRenderer{}
	e.Renderer = tr

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := NewHandler(cfg)
	if err := h.HandleHome(c); err != nil {
		t.Fatalf("HandleHome вернул ошибку: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("ожидается статус %d, получен %d", http.StatusOK, rec.Code)
	}

	if tr.name != "index" {
		t.Fatalf("ожидается имя рендера 'index', получено %q", tr.name)
	}

	if tr.data == nil {
		t.Fatalf("ожидаются не-nil данные рендера")
	}

	v := reflect.ValueOf(tr.data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	field := v.FieldByName("Projects")
	if !field.IsValid() {
		t.Fatalf("в данных рендера отсутствует поле Projects")
	}
	if field.Len() != 2 {
		t.Fatalf("ожидается 2 проекта, получено %d", field.Len())
	}
}
