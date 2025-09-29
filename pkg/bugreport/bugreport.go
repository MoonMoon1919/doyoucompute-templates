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

// Overrides the default expected behavior section of the document
// This overrides the entire section, including the section title
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

// Overrides the default actual behavior section of the document
// This overrides the entire section, including the section title
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

// Overrides the default environment section of the document
// This overrides the entire section, including the section title
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

// Overrides the default reproduction steps section of the document
// This overrides the entire section, including the section title
//
// Example:
//
//	section := doyoucompute.NewSection("Repro steps")
//	// Add section content...
//	bugreport.WithReproductionSteps(section)
func WithReproductionSteps(reproSteps doyoucompute.Section) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.reproductionSteps = reproSteps

		return nil, nil
	}
}

// Overrides the default code samples section of the document
// This overrides the entire section, including the section title
//
// Example:
//
//	section := doyoucompute.NewSection("Code samples")
//	// Add section content...
//	bugreport.WithCodeSamples(section)
func WithCodeSamples(samples doyoucompute.Section) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.codeSamples = samples

		return nil, nil
	}
}

// Overrides the default errors section of the document
// This overrides the entire section, including the section title
//
// Example:
//
//	section := doyoucompute.NewSection("Errors")
//	// Add section content...
//	bugreport.WithErrorDetails(section)
func WithErrorDetails(errDetails doyoucompute.Section) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.environmentDetails = errDetails

		return nil, nil
	}
}

// Overrides the name of the document
// This has a side effect of updating the name of the document in the frontmatter.
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

func DefaultExpectedBehavior() doyoucompute.Section {
	return helpers.SectionFactory("Expected behavior", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("What should happen?")

		return s
	})
}

func DefaultActualBehavior() doyoucompute.Section {
	return helpers.SectionFactory("Actual behavior", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("What actually happens?")

		return s
	})
}

func DefaultEnvirionmentDetails() doyoucompute.Section {
	return helpers.SectionFactory("Environment details", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("Tell us what go version, os, package version, etc.")

		return s
	})
}

func DefaultCodeSamples() doyoucompute.Section {
	return helpers.SectionFactory("Code Samples", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("Share a snippet of code that demonstrates the bug.")
		s.WriteCodeBlock("sh", []string{"# place code in here"}, doyoucompute.Static)

		return s
	})
}

func DefaultErrorMessages() doyoucompute.Section {
	return helpers.SectionFactory("Error Messages", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("Add any relevant error messages/logs here.")

		return s
	})
}

func DefaultStepsToReproduce() doyoucompute.Section {
	return helpers.SectionFactory("Steps to reproduce", func(s doyoucompute.Section) doyoucompute.Section {
		reproList := s.CreateList(doyoucompute.NUMBERED)
		reproList.Append("") // Intentionally empty
		reproList.Append("")
		reproList.Append("")

		return s
	})
}

func DefaultFrontMatter() doyoucompute.Frontmatter {
	return *doyoucompute.NewFrontmatter(map[string]interface{}{
		"name":      DEFAULT_NAME,
		"about":     "Report a bug",
		"title":     "",
		"labels":    "",
		"assignees": "",
	})
}

// New creates a new Bug Report document
// Uses defaults for all sections by defualt
// Takes in zero to many OptionsFunc to override defaults
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
