package readme

import (
	"strings"
	"testing"

	"github.com/MoonMoon1919/doyoucompute"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/helpers"
)

func TestReadme(t *testing.T) {
	introParagraph := doyoucompute.NewParagraph()
	introParagraph.Text("This is a test project")

	featuresSection := doyoucompute.NewSection("Features")
	featuresSection.WriteParagraph().Text("Feature 1")

	quickStartSection := doyoucompute.NewSection("Quick Start")
	quickStartSection.WriteParagraph().Text("Install and run")

	additionalSection := doyoucompute.NewSection("Usage")
	additionalSection.WriteParagraph().Text("How to use")

	tests := []struct {
		name               string
		props              ReadmeProps
		additionalSections []doyoucompute.Section
		opts               []helpers.OptionsFunc[ReadmeProps]
		wantErr            bool
		wantName           string
		wantMinContent     int
	}{
		{
			name: "default readme",
			props: ReadmeProps{
				Name:       "Test Project",
				Intro:      *introParagraph,
				Features:   featuresSection,
				QuickStart: quickStartSection,
			},
			additionalSections: nil,
			opts:               nil,
			wantErr:            false,
			wantName:           "Test Project",
			wantMinContent:     4, // Features, QuickStart, Contributing, License
		},
		{
			name: "with custom name",
			props: ReadmeProps{
				Name:       "Test Project",
				Intro:      *introParagraph,
				Features:   featuresSection,
				QuickStart: quickStartSection,
			},
			additionalSections: nil,
			opts: []helpers.OptionsFunc[ReadmeProps]{
				WithName("My Awesome Project"),
			},
			wantErr:        false,
			wantName:       "My Awesome Project",
			wantMinContent: 4,
		},
		{
			name: "with additional sections",
			props: ReadmeProps{
				Name:       "Test Project",
				Intro:      *introParagraph,
				Features:   featuresSection,
				QuickStart: quickStartSection,
			},
			additionalSections: []doyoucompute.Section{additionalSection},
			opts:               nil,
			wantErr:            false,
			wantName:           "Test Project",
			wantMinContent:     5, // Features, QuickStart, Usage, Contributing, License
		},
		{
			name: "with multiple additional sections",
			props: ReadmeProps{
				Name:       "Test Project",
				Intro:      *introParagraph,
				Features:   featuresSection,
				QuickStart: quickStartSection,
			},
			additionalSections: []doyoucompute.Section{
				additionalSection,
				doyoucompute.NewSection("API"),
				doyoucompute.NewSection("Examples"),
			},
			opts:           nil,
			wantErr:        false,
			wantName:       "Test Project",
			wantMinContent: 7, // Features, QuickStart, Usage, API, Examples, Contributing, License
		},
		{
			name: "with name and additional sections",
			props: ReadmeProps{
				Name:       "Test Project",
				Intro:      *introParagraph,
				Features:   featuresSection,
				QuickStart: quickStartSection,
			},
			additionalSections: []doyoucompute.Section{additionalSection},
			opts: []helpers.OptionsFunc[ReadmeProps]{
				WithName("Complete Documentation"),
			},
			wantErr:        false,
			wantName:       "Complete Documentation",
			wantMinContent: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := New(tt.props, tt.additionalSections, tt.opts...)
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

func TestReadmeContent(t *testing.T) {
	introParagraph := doyoucompute.NewParagraph()
	introParagraph.Text("This is a test project for testing")

	featuresSection := doyoucompute.NewSection("Features")
	featuresSection.WriteParagraph().Text("Feature 1: Amazing functionality")

	quickStartSection := doyoucompute.NewSection("Quick Start")
	quickStartSection.WriteParagraph().Text("Run npm install")

	tests := []struct {
		name               string
		props              ReadmeProps
		additionalSections []doyoucompute.Section
		opts               []helpers.OptionsFunc[ReadmeProps]
		wantContains       []string
		wantNotContains    []string
	}{
		{
			name: "default sections contain expected content",
			props: ReadmeProps{
				Name:       "Test Project",
				Intro:      *introParagraph,
				Features:   featuresSection,
				QuickStart: quickStartSection,
			},
			additionalSections: nil,
			opts:               nil,
			wantContains: []string{
				"This is a test project for testing",
				"Features",
				"Feature 1: Amazing functionality",
				"Quick Start",
				"Run npm install",
				"Contributing",
				"License",
				"CONTRIBUTING.md",
				"LICENSE",
			},
		},
		{
			name: "additional sections appear in content",
			props: ReadmeProps{
				Name:       "Test Project",
				Intro:      *introParagraph,
				Features:   featuresSection,
				QuickStart: quickStartSection,
			},
			additionalSections: []doyoucompute.Section{
				func() doyoucompute.Section {
					s := doyoucompute.NewSection("Usage")
					s.WriteParagraph().Text("How to use this")
					return s
				}(),
				func() doyoucompute.Section {
					s := doyoucompute.NewSection("API Reference")
					s.WriteParagraph().Text("API documentation")
					return s
				}(),
			},
			opts: nil,
			wantContains: []string{
				"Usage",
				"How to use this",
				"API Reference",
				"API documentation",
			},
		},
		{
			name: "contributing and license always present",
			props: ReadmeProps{
				Name:       "Test Project",
				Intro:      *introParagraph,
				Features:   featuresSection,
				QuickStart: quickStartSection,
			},
			additionalSections: []doyoucompute.Section{
				func() doyoucompute.Section {
					s := doyoucompute.NewSection("Usage")
					s.WriteParagraph().Text("Custom usage")
					return s
				}(),
			},
			opts: nil,
			wantContains: []string{
				"Contributing",
				"License",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := New(tt.props, tt.additionalSections, tt.opts...)
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

func TestReadmeSectionOrder(t *testing.T) {
	introParagraph := doyoucompute.NewParagraph()
	introParagraph.Text("Test intro")

	featuresSection := doyoucompute.NewSection("Features")
	featuresSection.WriteParagraph().Text("Feature content")

	quickStartSection := doyoucompute.NewSection("Quick Start")
	quickStartSection.WriteParagraph().Text("Quick start content")

	usage := doyoucompute.NewSection("Usage")
	usage.WriteParagraph().Text("Usage content")

	api := doyoucompute.NewSection("API Reference")
	api.WriteParagraph().Text("API content")

	props := ReadmeProps{
		Name:       "Test Project",
		Intro:      *introParagraph,
		Features:   featuresSection,
		QuickStart: quickStartSection,
	}

	doc, err := New(props, []doyoucompute.Section{usage, api})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	renderer := doyoucompute.NewMarkdownRenderer()
	rendered, err := renderer.Render(&doc)
	if err != nil {
		t.Fatalf("renderer.Render() error = %v", err)
	}

	// Find positions of each section
	featuresPos := strings.Index(rendered, "Features")
	quickStartPos := strings.Index(rendered, "Quick Start")
	usagePos := strings.Index(rendered, "Usage")
	apiPos := strings.Index(rendered, "API Reference")
	contributingPos := strings.Index(rendered, "Contributing")
	licensePos := strings.Index(rendered, "License")

	// Verify all sections are present
	if featuresPos == -1 {
		t.Error("Features section not found")
	}
	if quickStartPos == -1 {
		t.Error("Quick Start section not found")
	}
	if usagePos == -1 {
		t.Error("Usage section not found")
	}
	if apiPos == -1 {
		t.Error("API Reference section not found")
	}
	if contributingPos == -1 {
		t.Error("Contributing section not found")
	}
	if licensePos == -1 {
		t.Error("License section not found")
	}

	// Verify order: Features -> QuickStart -> Usage -> API -> Contributing -> License
	if featuresPos >= quickStartPos {
		t.Error("Features should come before Quick Start")
	}
	if quickStartPos >= usagePos {
		t.Error("Quick Start should come before Usage")
	}
	if usagePos >= apiPos {
		t.Error("Usage should come before API Reference")
	}
	if apiPos >= contributingPos {
		t.Error("API Reference should come before Contributing")
	}
	if contributingPos >= licensePos {
		t.Error("Contributing should come before License")
	}
}

func TestDefaultFunctions(t *testing.T) {
	tests := []struct {
		name     string
		testFunc func() interface{}
		wantNil  bool
	}{
		{
			name: "DefaultContributing",
			testFunc: func() interface{} {
				return DefaultContributing()
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
