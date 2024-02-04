package templates_test

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"go2api/pkg/templates"
)

func TestLoadTemplates(t *testing.T) {
	t.Run("Nonexistent Template Folder", func(t *testing.T) {
		testTemplate := &templates.Template{}
		pattern := "emptytemplate/*.html"

		expectedErrorSubstring := "matches no files"
		err := testTemplate.LoadTemplates(pattern)
		if assert.Error(t, err) {
			assert.ErrorContains(t, err, expectedErrorSubstring)
		}
	})

	t.Run("Pattern No Match", func(t *testing.T) {
		tempDir := "templatefolder1"
		if err := os.Mkdir(tempDir, os.ModePerm); err != nil {
			t.Error(err)
		}

		testTemplate := &templates.Template{}
		pattern := "templatefolder2/*.html"

		expectedErrorSubstring := "matches no files"
		err := testTemplate.LoadTemplates(pattern)
		if assert.Error(t, err) {
			assert.ErrorContains(t, err, expectedErrorSubstring)
		}

		if err := os.Remove(tempDir); err != nil {
			t.Error(err)
		}
	})

	t.Run("Folder Exists But No HTML", func(t *testing.T) {
		tempDir := "templatefolder888"
		if err := os.Mkdir(tempDir, os.ModePerm); err != nil {
			t.Error(err)
		}

		testTemplate := &templates.Template{}
		pattern := tempDir + "/*.html"

		expectedErrorSubstring := "matches no files"
		err := testTemplate.LoadTemplates(pattern)
		if assert.Error(t, err) {
			assert.ErrorContains(t, err, expectedErrorSubstring)
		}

		if err := os.Remove(tempDir); err != nil {
			t.Error(err)
		}
	})

	t.Run("Happy Path", func(t *testing.T) {
		tempDir := "templatefolder999"
		if err := os.Mkdir(tempDir, os.ModePerm); err != nil {
			t.Error(err)
		}
		if _, err := os.Create(tempDir + "/index.html"); err != nil {
			t.Error(err)
		}

		testTemplate := &templates.Template{}
		pattern := tempDir + "/*.html"

		err := testTemplate.LoadTemplates(pattern)
		if assert.NoError(t, err) && assert.NotNil(t, testTemplate.Tmpl) {
			if err := os.RemoveAll(tempDir); err != nil {
				t.Error(err)
			}
		}
	})
}

func TestExecuteTemplate(t *testing.T) {
	t.Run("Nonexistent Template", func(t *testing.T) {
		tempDir := "templatefolder11"
		if err := os.Mkdir(tempDir, os.ModePerm); err != nil {
			t.Error(err)
		}
		if _, err := os.Create(tempDir + "/index.html"); err != nil {
			t.Error(err)
		}

		testTemplate := &templates.Template{}
		pattern := tempDir + "/*.html"
		_ = testTemplate.LoadTemplates(pattern)

		w := httptest.NewRecorder()
		expectedErrorSubstring := "is undefined"
		err := testTemplate.ExecuteTemplate(w, "home.html", nil)

		if assert.Error(t, err) {
			assert.ErrorContains(t, err, expectedErrorSubstring)
			if err := os.RemoveAll(tempDir); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Happy Path", func(t *testing.T) {
		tempDir := "templatefolder111"
		if err := os.Mkdir(tempDir, os.ModePerm); err != nil {
			t.Error(err)
		}
		if _, err := os.Create(tempDir + "/index.html"); err != nil {
			t.Error(err)
		}

		testTemplate := &templates.Template{}
		pattern := tempDir + "/*.html"
		_ = testTemplate.LoadTemplates(pattern)

		w := httptest.NewRecorder()
		err := testTemplate.ExecuteTemplate(w, "index.html", nil)

		if assert.NoError(t, err) {
			if err := os.RemoveAll(tempDir); err != nil {
				t.Error(err)
			}
		}
	})
}
