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

func WithFrontMatter(frontmatter doyoucompute.Frontmatter) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.frontmatter = frontmatter

		return nil, nil
	}
}

func WithExpectedBehavior(behavior doyoucompute.Section) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.expectedBehavior = behavior

		return nil, nil
	}
}

func WithActualBehavior(behavior doyoucompute.Section) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.actualBehavior = behavior

		return nil, nil
	}
}

func WithEnvironmentDetails(env doyoucompute.Section) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.environmentDetails = env

		return nil, nil
	}
}

func WithReproductionSteps(reproSteps doyoucompute.Section) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.reproductionSteps = reproSteps

		return nil, nil
	}
}

func WithCodeSamples(samples doyoucompute.Section) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.codeSamples = samples

		return nil, nil
	}
}

func WithErrorDetails(errDetails doyoucompute.Section) helpers.OptionsFunc[bugReportProps] {
	return func(p *bugReportProps) (helpers.PostEffect[bugReportProps], error) {
		p.environmentDetails = errDetails

		return nil, nil
	}
}

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

// Defaults
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
