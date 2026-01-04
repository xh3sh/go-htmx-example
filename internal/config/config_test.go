package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig_SuccessAndFields(t *testing.T) {
	td := t.TempDir()

	cfgDir := filepath.Join(td, "config")
	if err := os.MkdirAll(cfgDir, 0o755); err != nil {
		t.Fatalf("создание директории config: %v", err)
	}

	yml := `
server:
  port: ":8080"
projects: []
social:
  name: "me"
skills:
  languages:
    - name: "Go"
      tech: ["Tests"]
      level: 55
`
	if err := os.WriteFile(filepath.Join(cfgDir, "config.yaml"), []byte(yml), 0o644); err != nil {
		t.Fatalf("запись config.yaml: %v", err)
	}

	origWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("получение текущей директории: %v", err)
	}
	defer func() { _ = os.Chdir(origWd) }()

	if err := os.Chdir(td); err != nil {
		t.Fatalf("изменение директории: %v", err)
	}

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig вернул ошибку: %v", err)
	}
	if cfg == nil {
		t.Fatalf("ожидается не-nil конфиг")
	}
	if cfg.Server.Port != ":8080" {
		t.Fatalf("ожидается порт %q, получен %q", ":8080", cfg.Server.Port)
	}

	if cfg.Social.Name != "me" {
		t.Fatalf("ожидается social.name 'me', получено %q", cfg.Social.Name)
	}
	if len(cfg.Skills.Languages) != 1 || cfg.Skills.Languages[0].Name != "Go" {
		t.Fatalf("неожиданные skills.languages: %#v", cfg.Skills.Languages)
	}
}

func TestLoadConfig_FileMissing(t *testing.T) {
	td := t.TempDir()

	origWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	defer func() { _ = os.Chdir(origWd) }()

	if err := os.Chdir(td); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	_, err = LoadConfig()
	if err == nil {
		t.Fatalf("ожидается ошибка при отсутствии файла конфигурации")
	}
}

func TestLoadConfig_InvalidYAML(t *testing.T) {
	td := t.TempDir()
	cfgDir := filepath.Join(td, "config")
	if err := os.MkdirAll(cfgDir, 0o755); err != nil {
		t.Fatalf("создание директории config: %v", err)
	}

	// invalid YAML
	if err := os.WriteFile(filepath.Join(cfgDir, "config.yaml"), []byte("::: not yaml :::"), 0o644); err != nil {
		t.Fatalf("write config.yaml: %v", err)
	}

	origWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	defer func() { _ = os.Chdir(origWd) }()

	if err := os.Chdir(td); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	_, err = LoadConfig()
	if err == nil {
		t.Fatalf("ожидается ошибка при некорректном YAML")
	}
}
