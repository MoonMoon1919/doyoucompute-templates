package pullrequest

import (
	"strings"
	"testing"

	"github.com/MoonMoon1919/doyoucompute"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/helpers"
)

func TestPullRequest(t *testing.T) {
	customSection := doyoucompute.NewSection("Custom Section")
	customSection.WriteParagraph().Text("Custom content")

	tests := []struct {
		name             string
		opts             []helpers.OptionsFunc[pullRequestProps]
		wantErr          bool
		wantName         string
		wantContentCount int
	}{
		{
			name:             "default pull request",
			opts:             nil,
			wantErr:          false,
			wantName:         "Pull Request",
			wantContentCount: 3,
		},
		{
			name: "with custom name",
			opts: []helpers.OptionsFunc[pullRequestProps]{
				WithName("Feature: Authentication"),
			},
			wantErr:          false,
			wantName:         "Feature: Authentication",
			wantContentCount: 3,
		},
		{
			name: "with custom description",
			opts: []helpers.OptionsFunc[pullRequestProps]{
				WithDescription(customSection),
			},
			wantErr:          false,
			wantName:         "Pull Request",
			wantContentCount: 3,
		},
		{
			name: "with custom related issue",
			opts: []helpers.OptionsFunc[pullRequestProps]{
				WithRelatedIssue(customSection),
			},
			wantErr:          false,
			wantName:         "Pull Request",
			wantContentCount: 3,
		},
		{
			name: "with custom testing",
			opts: []helpers.OptionsFunc[pullRequestProps]{
				WithTesting(customSection),
			},
			wantErr:          false,
			wantName:         "Pull Request",
			wantContentCount: 3,
		},
		{
			name: "with multiple options",
			opts: []helpers.OptionsFunc[pullRequestProps]{
				WithName("Bug Fix: Memory Leak"),
				WithDescription(customSection),
				WithTesting(customSection),
			},
			wantErr:          false,
			wantName:         "Bug Fix: Memory Leak",
			wantContentCount: 3,
		},
		{
			name: "with all options",
			opts: []helpers.OptionsFunc[pullRequestProps]{
				WithName("Complete PR"),
				WithDescription(customSection),
				WithRelatedIssue(customSection),
				WithTesting(customSection),
			},
			wantErr:          false,
			wantName:         "Complete PR",
			wantContentCount: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := New(tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			// Check document name
			if doc.Name != tt.wantName {
				t.Errorf("New() name = %v, want %v", doc.Name, tt.wantName)
			}

			// Check content count
			if len(doc.Content) != tt.wantContentCount {
				t.Errorf("New() content count = %v, want %v", len(doc.Content), tt.wantContentCount)
			}

			// Verify document can be rendered
			renderer := doyoucompute.NewMarkdownRenderer()
			rendered, err := renderer.Render(&doc)
			if err != nil {
				t.Errorf("renderer.Render() error = %v", err)
			}
			if rendered == "" {
				t.Error("renderer.Render() returned empty string")
			}
		})
	}
}

func TestPullRequestContent(t *testing.T) {
	tests := []struct {
		name            string
		opts            []helpers.OptionsFunc[pullRequestProps]
		wantContains    []string
		wantNotContains []string
	}{
		{
			name: "default sections contain expected content",
			opts: nil,
			wantContains: []string{
				"Description",
				"Related issue",
				"How I tested",
			},
		},
		{
			name: "custom description replaces default",
			opts: []helpers.OptionsFunc[pullRequestProps]{
				WithDescription(func() doyoucompute.Section {
					s := doyoucompute.NewSection("My Description")
					s.WriteParagraph().Text("Added OAuth2 support")
					return s
				}()),
			},
			wantContains: []string{
				"My Description",
				"Added OAuth2 support",
			},
			wantNotContains: []string{
				"What is this change and why are you making it?",
			},
		},
		{
			name: "custom related issue replaces default",
			opts: []helpers.OptionsFunc[pullRequestProps]{
				WithRelatedIssue(func() doyoucompute.Section {
					s := doyoucompute.NewSection("Fixes")
					s.WriteParagraph().Text("Closes #42")
					return s
				}()),
			},
			wantContains: []string{
				"Fixes",
				"Closes #42",
			},
			wantNotContains: []string{
				"Link to the relevant issue here",
			},
		},
		{
			name: "custom testing replaces default",
			opts: []helpers.OptionsFunc[pullRequestProps]{
				WithTesting(func() doyoucompute.Section {
					s := doyoucompute.NewSection("Testing Details")
					s.WriteParagraph().Text("Ran unit tests and integration tests")
					return s
				}()),
			},
			wantContains: []string{
				"Testing Details",
				"Ran unit tests and integration tests",
			},
			wantNotContains: []string{
				"How did you test these changes?",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := New(tt.opts...)
			if err != nil {
				t.Fatalf("New() error = %v", err)
			}

			renderer := doyoucompute.NewMarkdownRenderer()
			rendered, err := renderer.Render(&doc)
			if err != nil {
				t.Fatalf("renderer.Render() error = %v", err)
			}

			for _, want := range tt.wantContains {
				if !strings.Contains(rendered, want) {
					t.Errorf("renderer.Render() missing expected content: %q", want)
				}
			}

			for _, notWant := range tt.wantNotContains {
				if strings.Contains(rendered, notWant) {
					t.Errorf("renderer.Render() contains unexpected content: %q", notWant)
				}
			}
		})
	}
}

func TestDefaultFunctions(t *testing.T) {
	tests := []struct {
		name     string
		testFunc func() interface{}
		wantNil  bool
	}{
		{
			name: "DefaultName",
			testFunc: func() interface{} {
				return DefaultName()
			},
			wantNil: false,
		},
		{
			name: "DefaultDescription",
			testFunc: func() interface{} {
				return DefaultDescription()
			},
			wantNil: false,
		},
		{
			name: "DefaultRelatedIssue",
			testFunc: func() interface{} {
				return DefaultRelatedIssue()
			},
			wantNil: false,
		},
		{
			name: "DefaultTesting",
			testFunc: func() interface{} {
				return DefaultTesting()
			},
			wantNil: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.testFunc()
			if tt.wantNil && result != nil {
				t.Errorf("%s() returned non-nil, want nil", tt.name)
			}
			if !tt.wantNil && result == nil {
				t.Errorf("%s() returned nil, want non-nil", tt.name)
			}
		})
	}
}

func TestPullRequestValidation(t *testing.T) {
	tests := []struct {
		name    string
		opts    []helpers.OptionsFunc[pullRequestProps]
		wantErr bool
		errMsg  string
	}{
		{
			name: "empty name should error",
			opts: []helpers.OptionsFunc[pullRequestProps]{
				WithName(""),
			},
			wantErr: true,
			errMsg:  "name cannot be empty",
		},
		{
			name:    "default name should pass",
			opts:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err != nil && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("New() error = %v, should contain %q", err, tt.errMsg)
			}
		})
	}
}
