package contributing

import (
	"strings"
	"testing"

	"github.com/MoonMoon1919/doyoucompute"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/helpers"
)

func TestContributing(t *testing.T) {
	customSection := doyoucompute.NewSection("Custom Section")
	customSection.WriteParagraph().Text("Custom content")

	tests := []struct {
		name            string
		projectUrl      string
		issueTrackerUrl string
		opts            []helpers.OptionsFunc[contributingProps]
		wantErr         bool
		wantName        string
		wantMinContent  int
	}{
		{
			name:            "default contributing",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			opts:            nil,
			wantErr:         false,
			wantName:        "Contributing",
			wantMinContent:  2,
		},
		{
			name:            "with custom name",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			opts: []helpers.OptionsFunc[contributingProps]{
				WithName("Contributing Guidelines"),
			},
			wantErr:        false,
			wantName:       "Contributing Guidelines",
			wantMinContent: 2,
		},
		{
			name:            "with custom project url",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			opts: []helpers.OptionsFunc[contributingProps]{
				WithProjectUrl("https://github.com/other/repo"),
			},
			wantErr:        false,
			wantName:       "Contributing",
			wantMinContent: 2,
		},
		{
			name:            "with custom issue tracker url",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			opts: []helpers.OptionsFunc[contributingProps]{
				WithIssueTrackerUrl("https://github.com/other/repo/issues"),
			},
			wantErr:        false,
			wantName:       "Contributing",
			wantMinContent: 2,
		},
		{
			name:            "with custom getting started",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			opts: []helpers.OptionsFunc[contributingProps]{
				WithGettingStarted(customSection),
			},
			wantErr:        false,
			wantName:       "Contributing",
			wantMinContent: 2,
		},
		{
			name:            "with custom chose a task",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			opts: []helpers.OptionsFunc[contributingProps]{
				WithChoseATask(customSection),
			},
			wantErr:        false,
			wantName:       "Contributing",
			wantMinContent: 2,
		},
		{
			name:            "with custom setup",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			opts: []helpers.OptionsFunc[contributingProps]{
				WithSetup(customSection),
			},
			wantErr:        false,
			wantName:       "Contributing",
			wantMinContent: 2,
		},
		{
			name:            "with custom development",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			opts: []helpers.OptionsFunc[contributingProps]{
				WithDevelopment(customSection),
			},
			wantErr:        false,
			wantName:       "Contributing",
			wantMinContent: 2,
		},
		{
			name:            "with custom submissions",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			opts: []helpers.OptionsFunc[contributingProps]{
				WithSubmissions(customSection),
			},
			wantErr:        false,
			wantName:       "Contributing",
			wantMinContent: 2,
		},
		{
			name:            "with custom writing docs",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			opts: []helpers.OptionsFunc[contributingProps]{
				WithWritingDocs(customSection),
			},
			wantErr:        false,
			wantName:       "Contributing",
			wantMinContent: 2,
		},
		{
			name:            "with custom reporting bugs",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			opts: []helpers.OptionsFunc[contributingProps]{
				WithReportingbugs(customSection),
			},
			wantErr:        false,
			wantName:       "Contributing",
			wantMinContent: 2,
		},
		{
			name:            "with custom license",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			opts: []helpers.OptionsFunc[contributingProps]{
				WithLicense(customSection),
			},
			wantErr:        false,
			wantName:       "Contributing",
			wantMinContent: 2,
		},
		{
			name:            "with multiple options",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			opts: []helpers.OptionsFunc[contributingProps]{
				WithName("How to Contribute"),
				WithSetup(customSection),
				WithDevelopment(customSection),
			},
			wantErr:        false,
			wantName:       "How to Contribute",
			wantMinContent: 2,
		},
		{
			name:            "with all options",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			opts: []helpers.OptionsFunc[contributingProps]{
				WithName("Complete Contributing Guide"),
				WithProjectUrl("https://github.com/custom/project"),
				WithIssueTrackerUrl("https://github.com/custom/project/issues"),
				WithGettingStarted(customSection),
				WithChoseATask(customSection),
				WithSetup(customSection),
				WithDevelopment(customSection),
				WithSubmissions(customSection),
				WithWritingDocs(customSection),
				WithReportingbugs(customSection),
				WithLicense(customSection),
			},
			wantErr:        false,
			wantName:       "Complete Contributing Guide",
			wantMinContent: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := New(tt.projectUrl, tt.issueTrackerUrl, tt.opts...)
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

			// Check minimum content count
			if len(doc.Content) < tt.wantMinContent {
				t.Errorf("New() content count = %v, want at least %v", len(doc.Content), tt.wantMinContent)
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

func TestContributingValidation(t *testing.T) {
	tests := []struct {
		name            string
		projectUrl      string
		issueTrackerUrl string
		opts            []helpers.OptionsFunc[contributingProps]
		wantErr         bool
		errMsg          string
	}{
		{
			name:            "empty project URL should error",
			projectUrl:      "",
			issueTrackerUrl: "https://github.com/user/project/issues",
			wantErr:         true,
			errMsg:          "projectUrl cannot be empty",
		},
		{
			name:            "empty issue tracker URL should error",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "",
			wantErr:         true,
			errMsg:          "issueTrackerUrl cannot be empty",
		},
		{
			name:            "invalid project URL should error",
			projectUrl:      "https://github.com/",
			issueTrackerUrl: "https://github.com/user/project/issues",
			wantErr:         true,
			errMsg:          "could not extract project name",
		},
		{
			name:            "empty name after options should error",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			opts: []helpers.OptionsFunc[contributingProps]{
				WithName(""),
			},
			wantErr: true,
			errMsg:  "name cannot be empty",
		},
		{
			name:            "valid inputs should pass",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			wantErr:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.projectUrl, tt.issueTrackerUrl, tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err != nil && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("New() error = %v, should contain %q", err, tt.errMsg)
			}
		})
	}
}

func TestContributingContent(t *testing.T) {
	tests := []struct {
		name            string
		projectUrl      string
		issueTrackerUrl string
		opts            []helpers.OptionsFunc[contributingProps]
		wantContains    []string
		wantNotContains []string
	}{
		{
			name:            "default sections contain expected content",
			projectUrl:      "https://github.com/user/myproject",
			issueTrackerUrl: "https://github.com/user/myproject/issues",
			opts:            nil,
			wantContains: []string{
				"Getting started",
				"Find a task",
				"Setting Up Your Development Environment",
				"Development Workflow",
				"Submitting your changes",
				"Reporting bugs",
				"Writing documentation",
				"License",
				"myproject",
			},
		},
		{
			name:            "custom section replaces default",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/project/issues",
			opts: []helpers.OptionsFunc[contributingProps]{
				WithSetup(func() doyoucompute.Section {
					s := doyoucompute.NewSection("Custom Setup")
					s.WriteParagraph().Text("Use Docker Compose")
					return s
				}()),
			},
			wantContains: []string{
				"Custom Setup",
				"Use Docker Compose",
			},
			wantNotContains: []string{
				"Setting Up Your Development Environment",
			},
		},
		{
			name:            "issue tracker URL appears in content",
			projectUrl:      "https://github.com/user/project",
			issueTrackerUrl: "https://github.com/user/customissues",
			opts:            nil,
			wantContains: []string{
				"customissues",
			},
		},
		{
			name:            "project name extracted from URL",
			projectUrl:      "https://github.com/user/awesome-project",
			issueTrackerUrl: "https://github.com/user/awesome-project/issues",
			opts:            nil,
			wantContains: []string{
				"awesome-project",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := New(tt.projectUrl, tt.issueTrackerUrl, tt.opts...)
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

func TestContributingIssueTrackerUpdates(t *testing.T) {
	projectUrl := "https://github.com/user/project"
	initialIssueUrl := "https://github.com/user/project/issues"
	newIssueUrl := "https://github.com/user/newproject/issues"

	doc, err := New(
		projectUrl,
		initialIssueUrl,
		WithIssueTrackerUrl(newIssueUrl),
	)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	renderer := doyoucompute.NewMarkdownRenderer()
	rendered, err := renderer.Render(&doc)
	if err != nil {
		t.Fatalf("renderer.Render() error = %v", err)
	}

	// Should contain the new issue tracker URL
	if !strings.Contains(rendered, "newproject/issues") {
		t.Error("renderer.Render() should contain updated issue tracker URL")
	}

	// Should NOT contain the old issue tracker URL
	if strings.Contains(rendered, initialIssueUrl) {
		t.Error("renderer.Render() should not contain old issue tracker URL")
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
			name: "DefaultGettingStarted",
			testFunc: func() interface{} {
				return DefaultGettingStarted()
			},
			wantNil: false,
		},
		{
			name: "DefaultChoseATask",
			testFunc: func() interface{} {
				return DefaultChoseATask("https://github.com/user/project/issues")
			},
			wantNil: false,
		},
		{
			name: "DefaultLicense",
			testFunc: func() interface{} {
				return DefaultLicense()
			},
			wantNil: false,
		},
		{
			name: "DefaultWritingDocs",
			testFunc: func() interface{} {
				return DefaultWritingDocs()
			},
			wantNil: false,
		},
		{
			name: "DefaultReportingBugs",
			testFunc: func() interface{} {
				return DefaultReportingBugs("https://github.com/user/project/issues")
			},
			wantNil: false,
		},
		{
			name: "DefaultOpenSourceGoSetupGuidelines",
			testFunc: func() interface{} {
				return DefaultOpenSourceGoSetupGuidelines("https://github.com/user/project", "project")
			},
			wantNil: false,
		},
		{
			name: "DefaultOpenSourceGoDevelopmentGuidelines",
			testFunc: func() interface{} {
				return DefaultOpenSourceGoDevelopmentGuidelines()
			},
			wantNil: false,
		},
		{
			name: "DefaultOpenSourceSubmittingGuidelines",
			testFunc: func() interface{} {
				return DefaultOpenSourceSubmittingGuidelines()
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
