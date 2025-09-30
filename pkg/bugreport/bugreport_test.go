package bugreport

import (
	"reflect"
	"strings"
	"testing"

	"github.com/MoonMoon1919/doyoucompute"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/helpers"
)

func TestBugReport(t *testing.T) {
	customSection := doyoucompute.NewSection("Custom Section")
	customSection.WriteParagraph().Text("Custom content")

	customFrontmatter := doyoucompute.NewFrontmatter(map[string]interface{}{
		"name":      "Custom Bug",
		"about":     "Custom bug report",
		"title":     "Bug: ",
		"labels":    "bug,critical",
		"assignees": "maintainer",
	})

	tests := []struct {
		name               string
		opts               []helpers.OptionsFunc[bugReportProps]
		wantErr            bool
		wantName           string
		wantContentCount   int
		checkFrontmatter   bool
		wantFrontmatterKey string
		wantFrontmatterVal string
	}{
		{
			name:               "default bug report",
			opts:               nil,
			wantErr:            false,
			wantName:           "Bug Report",
			wantContentCount:   6,
			checkFrontmatter:   true,
			wantFrontmatterKey: "name",
			wantFrontmatterVal: "Bug Report",
		},
		{
			name: "with custom name",
			opts: []helpers.OptionsFunc[bugReportProps]{
				WithName("Critical Bug Report"),
			},
			wantErr:            false,
			wantName:           "Critical Bug Report",
			wantContentCount:   6,
			checkFrontmatter:   true,
			wantFrontmatterKey: "name",
			wantFrontmatterVal: "Critical Bug Report",
		},
		{
			name: "with custom frontmatter",
			opts: []helpers.OptionsFunc[bugReportProps]{
				WithFrontMatter(*customFrontmatter),
			},
			wantErr:            false,
			wantName:           "Bug Report",
			wantContentCount:   6,
			checkFrontmatter:   true,
			wantFrontmatterKey: "labels",
			wantFrontmatterVal: "bug,critical",
		},
		{
			name: "with custom expected behavior",
			opts: []helpers.OptionsFunc[bugReportProps]{
				WithExpectedBehavior(customSection),
			},
			wantErr:          false,
			wantName:         "Bug Report",
			wantContentCount: 6,
		},
		{
			name: "with custom actual behavior",
			opts: []helpers.OptionsFunc[bugReportProps]{
				WithActualBehavior(customSection),
			},
			wantErr:          false,
			wantName:         "Bug Report",
			wantContentCount: 6,
		},
		{
			name: "with custom environment details",
			opts: []helpers.OptionsFunc[bugReportProps]{
				WithEnvironmentDetails(customSection),
			},
			wantErr:          false,
			wantName:         "Bug Report",
			wantContentCount: 6,
		},
		{
			name: "with custom reproduction steps",
			opts: []helpers.OptionsFunc[bugReportProps]{
				WithReproductionSteps(customSection),
			},
			wantErr:          false,
			wantName:         "Bug Report",
			wantContentCount: 6,
		},
		{
			name: "with custom code samples",
			opts: []helpers.OptionsFunc[bugReportProps]{
				WithCodeSamples(customSection),
			},
			wantErr:          false,
			wantName:         "Bug Report",
			wantContentCount: 6,
		},
		{
			name: "with custom error details",
			opts: []helpers.OptionsFunc[bugReportProps]{
				WithErrorDetails(customSection),
			},
			wantErr:          false,
			wantName:         "Bug Report",
			wantContentCount: 6,
		},
		{
			name: "with multiple options",
			opts: []helpers.OptionsFunc[bugReportProps]{
				WithName("Production Bug"),
				WithExpectedBehavior(customSection),
				WithActualBehavior(customSection),
			},
			wantErr:          false,
			wantName:         "Production Bug",
			wantContentCount: 6,
		},
		{
			name: "with all options",
			opts: []helpers.OptionsFunc[bugReportProps]{
				WithName("Complete Bug Report"),
				WithFrontMatter(*customFrontmatter),
				WithExpectedBehavior(customSection),
				WithActualBehavior(customSection),
				WithEnvironmentDetails(customSection),
				WithReproductionSteps(customSection),
				WithCodeSamples(customSection),
				WithErrorDetails(customSection),
			},
			wantErr:          false,
			wantName:         "Complete Bug Report",
			wantContentCount: 6,
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

			emptyFrontmatter := doyoucompute.Frontmatter{}

			// Check frontmatter if specified
			if tt.checkFrontmatter {
				if reflect.DeepEqual(doc.Frontmatter, emptyFrontmatter) {
					t.Error("New() frontmatter is nil")
				} else if val, ok := doc.Frontmatter.Data[tt.wantFrontmatterKey]; !ok {
					t.Errorf("New() frontmatter missing key %v", tt.wantFrontmatterKey)
				} else if val != tt.wantFrontmatterVal {
					t.Errorf("New() frontmatter[%v] = %v, want %v", tt.wantFrontmatterKey, val, tt.wantFrontmatterVal)
				}
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

func TestBugReportContent(t *testing.T) {
	tests := []struct {
		name            string
		opts            []helpers.OptionsFunc[bugReportProps]
		wantContains    []string
		wantNotContains []string
	}{
		{
			name: "default sections contain expected content",
			opts: nil,
			wantContains: []string{
				"Expected behavior",
				"Actual behavior",
				"Environment details",
				"Steps to reproduce",
				"Code Samples",
				"Error Messages",
			},
		},
		{
			name: "custom section replaces default",
			opts: []helpers.OptionsFunc[bugReportProps]{
				WithExpectedBehavior(func() doyoucompute.Section {
					s := doyoucompute.NewSection("My Custom Expected")
					s.WriteParagraph().Text("Custom expected behavior")
					return s
				}()),
			},
			wantContains: []string{
				"My Custom Expected",
				"Custom expected behavior",
			},
			wantNotContains: []string{
				"Expected behavior",
				"What should happen?",
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

func TestBugReportValidation(t *testing.T) {
	tests := []struct {
		name    string
		opts    []helpers.OptionsFunc[bugReportProps]
		wantErr bool
		errMsg  string
	}{
		{
			name: "empty name should error",
			opts: []helpers.OptionsFunc[bugReportProps]{
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

func TestDefaultFunctions(t *testing.T) {
	tests := []struct {
		name     string
		testFunc func() interface{}
		wantNil  bool
	}{
		{
			name: "DefaultFrontMatter",
			testFunc: func() interface{} {
				return DefaultFrontMatter()
			},
			wantNil: false,
		},
		{
			name: "DefaultExpectedBehavior",
			testFunc: func() interface{} {
				return DefaultExpectedBehavior()
			},
			wantNil: false,
		},
		{
			name: "DefaultActualBehavior",
			testFunc: func() interface{} {
				return DefaultActualBehavior()
			},
			wantNil: false,
		},
		{
			name: "DefaultEnvironmentDetails",
			testFunc: func() interface{} {
				return DefaultEnvironmentDetails()
			},
			wantNil: false,
		},
		{
			name: "DefaultCodeSamples",
			testFunc: func() interface{} {
				return DefaultCodeSamples()
			},
			wantNil: false,
		},
		{
			name: "DefaultErrorMessages",
			testFunc: func() interface{} {
				return DefaultErrorMessages()
			},
			wantNil: false,
		},
		{
			name: "DefaultStepsToReproduce",
			testFunc: func() interface{} {
				return DefaultStepsToReproduce()
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
