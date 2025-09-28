package bugreport

import (
	"github.com/MoonMoon1919/doyoucompute"
)

// START: Move these to main pkg
// Post Effects are useful for two attributes that are tightly coupled
// E.g., name and frontmatter in a github issue template
type PostEffect[T any] func(p *T) error
type OptionsFunc[T any] func(p *T) (PostEffect[T], error)

// TODO: Move contentFunc, sectionFactory, DocumentApplier, and documentBuilder to main pkg
type contentFunc func(s doyoucompute.Section) doyoucompute.Section

func sectionFactory(name string, contentFuncs ...contentFunc) doyoucompute.Section {
	s := doyoucompute.NewSection(name)

	for _, cFunc := range contentFuncs {
		s = cFunc(s)
	}

	return s
}

type DocumentApplier func(d *doyoucompute.Document) error

func documentBuilder(name string, appliers ...DocumentApplier) (doyoucompute.Document, error) {
	document, err := doyoucompute.NewDocument(name)
	if err != nil {
		return doyoucompute.Document{}, nil
	}

	for _, applier := range appliers {
		applier(&document)
	}

	return document, nil
}

// END: Move these to main pkg

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

func WithFrontMatter(frontmatter doyoucompute.Frontmatter) OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (PostEffect[bugReportProps], error) {
		p.frontmatter = frontmatter

		return nil, nil
	}
}

func WithExpectedBehavior(behavior doyoucompute.Section) OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (PostEffect[bugReportProps], error) {
		p.expectedBehavior = behavior

		return nil, nil
	}
}

func WithActualBehavior(behavior doyoucompute.Section) OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (PostEffect[bugReportProps], error) {
		p.actualBehavior = behavior

		return nil, nil
	}
}

func WithEnvironmentDetails(env doyoucompute.Section) OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (PostEffect[bugReportProps], error) {
		p.environmentDetails = env

		return nil, nil
	}
}

func WithReproductionSteps(reproSteps doyoucompute.Section) OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (PostEffect[bugReportProps], error) {
		p.reproductionSteps = reproSteps

		return nil, nil
	}
}

func WithCodeSamples(samples doyoucompute.Section) OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (PostEffect[bugReportProps], error) {
		p.codeSamples = samples

		return nil, nil
	}
}

func WithErrorDetails(errDetails doyoucompute.Section) OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (PostEffect[bugReportProps], error) {
		p.environmentDetails = errDetails

		return nil, nil
	}
}

func WithName(name string) OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (PostEffect[bugReportProps], error) {
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

// Defaults
func DefaultExpectedBehavior() doyoucompute.Section {
	return sectionFactory("Expected behavior", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("What should happen?")

		return s
	})
}

func DefaultActualBehavior() doyoucompute.Section {
	return sectionFactory("Actual behavior", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("What actually happens?")

		return s
	})
}

func DefaultEnvirionmentDetails() doyoucompute.Section {
	return sectionFactory("Environment details", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("Tell us what go version, os, package version, etc.")

		return s
	})
}

func DefaultCodeSamples() doyoucompute.Section {
	return sectionFactory("Code Samples", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("Share a snippet of code that demonstrates the bug.")
		s.WriteCodeBlock("sh", []string{"# place code in here"}, doyoucompute.Static)

		return s
	})
}

func DefaultErrorMessages() doyoucompute.Section {
	return sectionFactory("Error Messages", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("Add any relevant error messages/logs here.")

		return s
	})
}

func DefaultStepsToReproduce() doyoucompute.Section {
	return sectionFactory("Steps to reproduce", func(s doyoucompute.Section) doyoucompute.Section {
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

func New(opts ...OptionsFunc[bugReportProps]) (doyoucompute.Document, error) {
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

	for _, opt := range opts {
		postEffect, err := opt(&props)
		if err != nil {
			return doyoucompute.Document{}, err
		}

		if postEffect != nil {
			if err := postEffect(&props); err != nil {
				return doyoucompute.Document{}, err
			}
		}
	}

	return documentBuilder(props.name, func(d *doyoucompute.Document) error {
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
