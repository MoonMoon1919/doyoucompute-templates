package pullrequest

import (
	"github.com/MoonMoon1919/doyoucompute"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/helpers"
)

type pullRequestProps struct {
	name         string
	description  doyoucompute.Section
	relatedIssue doyoucompute.Section
	testing      doyoucompute.Section
}

func WithName(name string) helpers.OptionsFunc[pullRequestProps] {
	return func(p *pullRequestProps) (helpers.PostEffect[pullRequestProps], error) {
		p.name = name

		return nil, nil
	}
}

func WithDescription(description doyoucompute.Section) helpers.OptionsFunc[pullRequestProps] {
	return func(p *pullRequestProps) (helpers.PostEffect[pullRequestProps], error) {
		p.description = description

		return nil, nil
	}
}

func WithRelatedIssue(relatedIssue doyoucompute.Section) helpers.OptionsFunc[pullRequestProps] {
	return func(p *pullRequestProps) (helpers.PostEffect[pullRequestProps], error) {
		p.relatedIssue = relatedIssue

		return nil, nil
	}
}

func WithTesting(testing doyoucompute.Section) helpers.OptionsFunc[pullRequestProps] {
	return func(p *pullRequestProps) (helpers.PostEffect[pullRequestProps], error) {
		p.testing = testing

		return nil, nil
	}
}

func DefaultName() string {
	return "Pull Request"
}

func DefaultDescription() doyoucompute.Section {
	return helpers.SectionFactory("Description", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("What is this change and why are you making it?")
		return s
	})
}

func DefaultRelatedIssue() doyoucompute.Section {
	return helpers.SectionFactory("Related issue", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("Link to the relevant issue here.")
		return s
	})
}

func DefaultTesting() doyoucompute.Section {
	return helpers.SectionFactory("How I tested", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteComment("How did you test these changes?")
		return s
	})
}

func New(opts ...helpers.OptionsFunc[pullRequestProps]) (doyoucompute.Document, error) {
	props := pullRequestProps{
		name:         DefaultName(),
		description:  DefaultDescription(),
		relatedIssue: DefaultRelatedIssue(),
		testing:      DefaultTesting(),
	}

	err := helpers.ApplyOptions(props, opts...)
	if err != nil {
		return doyoucompute.Document{}, err
	}

	return helpers.DocumentBuilder(props.name, func(d *doyoucompute.Document) error {
		d.AddSection(props.description)
		d.AddSection(props.relatedIssue)
		d.AddSection(props.testing)

		return nil
	})
}
