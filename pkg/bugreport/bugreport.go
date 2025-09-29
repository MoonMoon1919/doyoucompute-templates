// Package bugreport provides a template for creating bug report documents.
//
// This package is part of doyoucompute-templates and uses the doyoucompute
// library to generate structured bug report documents with customizable sections.
//
// Basic usage:
//
//	doc, err := bugreport.New()
//	if err != nil {
//		// handle error
//	}
//
// Customizing sections:
//
//	doc, err := bugreport.New(
//		bugreport.WithName("Critical Bug"),
//		bugreport.WithExpectedBehavior(customSection),
//		bugreport.WithActualBehavior(customSection),
//	)
package bugreport

import (
	"github.com/MoonMoon1919/doyoucompute"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/helpers"
)

const DEFAULT_NAME = "Bug Report"

type bugReportProps struct {
	name               string
	frontmatter        doyoucompute.Frontmatter
	expectedBehavior   doyoucompute.Section
	actualBehavior     doyoucompute.Section
	environmentDetails doyoucompute.Section
	reproductionSteps  doyoucompute.Section
	codeSamples        doyoucompute.Section
	errors             doyoucompute.Section
}

// TODO: Add GH frontmatter type to main pkg

// Overrides the default frontmatter for the document
// Example:
//
//	bugreport.WithFrontMatter(*doyoucompute.NewFrontmatter(map[string]interface{}{
//		"name":      "Some name",
//		"about":     "Report a bug",
//		"title":     "",
//		"labels":    "",
//		"assignees": "",
//	}))
func WithFrontMatter(frontmatter doyoucompute.Frontmatter) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.frontmatter = frontmatter

		return nil, nil
	}
}

// WithExpectedBehavior overrides the default expected behavior section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("Expected behavior")
//	// Add section content...
//	bugreport.WithExpectedBehavior(section)
func WithExpectedBehavior(behavior doyoucompute.Section) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.expectedBehavior = behavior

		return nil, nil
	}
}

// WithActualBehavior overrides the default actual behavior section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("Actual behavior")
//	// Add section content...
//	bugreport.WithActualBehavior(section)
func WithActualBehavior(behavior doyoucompute.Section) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.actualBehavior = behavior

		return nil, nil
	}
}

// WithEnvironmentDetails overrides the default environment section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("Environment")
//	// Add section content...
//	bugreport.WithEnvironmentDetails(section)
func WithEnvironmentDetails(env doyoucompute.Section) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.environmentDetails = env

		return nil, nil
	}
}

// WithReproductionSteps overrides the default reproduction steps section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("Repro steps")
//	list := section.CreateList(doyoucompute.NUMBERED)
//	list.Append("Run the program")
//	list.Append("Observe the error")
//	bugreport.WithReproductionSteps(section)
func WithReproductionSteps(reproSteps doyoucompute.Section) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.reproductionSteps = reproSteps

		return nil, nil
	}
}

// WithCodeSamples overrides the default code samples section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("Code samples")
//	section.WriteCodeBlock("go", []string{"fmt.Println(\"bug\")"}, doyoucompute.Static)
//	bugreport.WithCodeSamples(section)
func WithCodeSamples(samples doyoucompute.Section) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.codeSamples = samples

		return nil, nil
	}
}

// WithErrorDetails overrides the default errors section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("Errors")
//	// Add section content...
//	bugreport.WithErrorDetails(section)
func WithErrorDetails(errDetails doyoucompute.Section) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.errors = errDetails

		return nil, nil
	}
}

// WithName overrides the document name and updates the frontmatter accordingly.
//
// Example:
//
//	bugreport.WithName("foo")
func WithName(name string) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.name = name

		return func(p *bugReportProps) error {
			p.frontmatter = doyoucompute.Frontmatter{
				Data: map[string]interface{}{
					"name":      name,
					"about":     "Report a bug",
					"title":     "",
					"labels":    "",
					"assignees": "",
				},
			}

			return nil
		}, nil
	}
}

// DefaultExpectedBehavior returns the default expected behavior section.
func DefaultExpectedBehavior() doyoucompute.Section {
	return helpers.SectionFactory("Expected behavior", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("What should happen?")

		return s
	})
}

// DefaultActualBehavior returns the default actual behavior section.
func DefaultActualBehavior() doyoucompute.Section {
	return helpers.SectionFactory("Actual behavior", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("What actually happens?")

		return s
	})
}

// DefaultEnvirionmentDetails returns the default environment details section.
func DefaultEnvirionmentDetails() doyoucompute.Section {
	return helpers.SectionFactory("Environment details", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("Tell us what go version, os, package version, etc.")

		return s
	})
}

// DefaultCodeSamples returns the default code samples section.
func DefaultCodeSamples() doyoucompute.Section {
	return helpers.SectionFactory("Code Samples", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("Share a snippet of code that demonstrates the bug.")
		s.WriteCodeBlock("sh", []string{"# place code in here"}, doyoucompute.Static)

		return s
	})
}

// DefaultErrorMessages returns the default error messages section.
func DefaultErrorMessages() doyoucompute.Section {
	return helpers.SectionFactory("Error Messages", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("Add any relevant error messages/logs here.")

		return s
	})
}

// DefaultStepsToReproduce returns the default steps to reproduce section.
func DefaultStepsToReproduce() doyoucompute.Section {
	return helpers.SectionFactory("Steps to reproduce", func(s doyoucompute.Section) doyoucompute.Section {
		reproList := s.CreateList(doyoucompute.NUMBERED)
		reproList.Append("") // Intentionally empty
		reproList.Append("")
		reproList.Append("")

		return s
	})
}

// DefaultFrontMatter returns the default frontmatter for bug reports.
func DefaultFrontMatter() doyoucompute.Frontmatter {
	return *doyoucompute.NewFrontmatter(map[string]interface{}{
		"name":      DEFAULT_NAME,
		"about":     "Report a bug",
		"title":     "",
		"labels":    "",
		"assignees": "",
	})
}

// New creates a new bug report document with default sections.
// Accepts zero or more option functions to customize the document.
//
// Example:
//
//	doc, err := bugreport.New(
//		bugreport.WithName("API Bug"),
//		bugreport.WithExpectedBehavior(customSection),
//	)
func New(opts ...helpers.OptionsFunc[bugReportProps]) (doyoucompute.Document, error) {
	props := bugReportProps{
		name:               DEFAULT_NAME,
		frontmatter:        DefaultFrontMatter(),
		expectedBehavior:   DefaultExpectedBehavior(),
		actualBehavior:     DefaultActualBehavior(),
		environmentDetails: DefaultEnvirionmentDetails(),
		reproductionSteps:  DefaultStepsToReproduce(),
		codeSamples:        DefaultCodeSamples(),
		errors:             DefaultErrorMessages(),
	}

	err := helpers.ApplyOptions(props, opts...)
	if err != nil {
		return doyoucompute.Document{}, err
	}

	return helpers.DocumentBuilder(props.name, func(d *doyoucompute.Document) error {
		d.AddFrontmatter(props.frontmatter)
		d.AddSection(props.expectedBehavior)
		d.AddSection(props.actualBehavior)
		d.AddSection(props.environmentDetails)
		d.AddSection(props.reproductionSteps)
		d.AddSection(props.codeSamples)
		d.AddSection(props.errors)

		return nil
	})
}
