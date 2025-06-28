package htemplate

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type HTemplate struct {
	rootDIR string
	tmpl    *template.Template
}

func NewHTemplate(rootDir string, pattern string) (*HTemplate, error) {
	tmpl, err := parseGlob(rootDir, pattern)
	if err != nil {
		return nil, err
	}
	return &HTemplate{
		rootDIR: rootDir,
		tmpl:    tmpl,
	}, nil
}

func parseGlob(rootDir, pattern string) (tmpl *template.Template, err error) {
	fullPattern := filepath.Join(rootDir, pattern)

	tmpl, err = template.ParseGlob(fullPattern)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template files with pattern %s: %w", fullPattern, err)
	}

	return
}

func (ht HTemplate) Execute(w http.ResponseWriter, filename string, data any) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := ht.tmpl.ExecuteTemplate(w, filename, data)
	if err != nil {
		http.Error(w, "Template rendering error", http.StatusInternalServerError)
		return err
	}

	return nil
}

type Template interface {
	ParseFiles(filename string) (*template.Template, error)
	Execute(w http.ResponseWriter, tmpl *template.Template, data any) error
}
