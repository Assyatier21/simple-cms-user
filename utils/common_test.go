package utils

import (
	m "cms-user/models"
	"testing"
	"time"
)

func TestFormatTimeResArticle(t *testing.T) {
	// setup
	createdAt := time.Now().Format("2006-01-02T15:04:05Z")
	updatedAt := time.Now().Format("2006-01-02T15:04:05Z")
	article := &m.ResArticle{Title: "Test Article", CreatedAt: createdAt, UpdatedAt: updatedAt}

	// run test
	result := FormatTimeResArticle(article)

	// asserts
	expectedCreatedAt := FormattedTime(createdAt)
	expectedUpdatedAt := FormattedTime(updatedAt)

	if result.CreatedAt != expectedCreatedAt {
		t.Errorf("Expected created at time to be %s, but got %s", expectedCreatedAt, result.CreatedAt)
	}
	if result.UpdatedAt != expectedUpdatedAt {
		t.Errorf("Expected updated at time to be %s, but got %s", expectedUpdatedAt, result.UpdatedAt)
	}
}
func TestFormatTimeResCategory(t *testing.T) {
	// setup
	createdAt := time.Now().Format("2006-01-02T15:04:05Z")
	updatedAt := time.Now().Format("2006-01-02T15:04:05Z")
	category := &m.Category{Title: "Test Category", CreatedAt: createdAt, UpdatedAt: updatedAt}

	// run test
	result := FormatTimeResCategory(category)

	// asserts
	expectedCreatedAt := FormattedTime(createdAt)
	expectedUpdatedAt := FormattedTime(updatedAt)

	if result.CreatedAt != expectedCreatedAt {
		t.Errorf("Expected created at time to be %s, but got %s", expectedCreatedAt, result.CreatedAt)
	}
	if result.UpdatedAt != expectedUpdatedAt {
		t.Errorf("Expected updated at time to be %s, but got %s", expectedUpdatedAt, result.UpdatedAt)
	}
}
func TestFormattedTime(t *testing.T) {
	ts := time.Now().Format("2006-01-02T15:04:05Z")

	result := FormattedTime(ts)
	expected := time.Now().Format("2006-01-02 15:04:05")
	if result != expected {
		t.Errorf("Expected formatted time to be %s, but got %s", expected, result)
	}

	invalidTs := "invalid time format"
	result = FormattedTime(invalidTs)
	if result != "" {
		t.Errorf("Expected empty string, but got %s", result)
	}
}
