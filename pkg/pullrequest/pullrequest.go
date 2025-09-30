// Package pullrequest provides a template for creating pull request documents.
//
// This package generates structured pull request templates with sections for
// describing changes, linking related issues, and explaining testing approaches.
//
// Basic usage:
//
//	doc, err := pullrequest.New()
//	if err != nil {
//		// handle error
//	}
//
// Customizing sections:
//
//	doc, err := pullrequest.New(
//		pullrequest.WithName("Feature PR"),
//		pullrequest.WithDescription(customSection),
//	)
package pullrequest

import (
	"fmt"

	"github.com/MoonMoon1919/doyoucompute"
)

type pullRequestProps struct {
	name         string
	description  doyoucompute.Section
	relatedIssue doyoucompute.Section
	testing      doyoucompute.Section
}

// WithName overrides the document name.
//
// Example:
//
//	pullrequest.WithName("Feature: Add authentication")
func WithName(name string) doyoucompute.OptionBuilder[pullRequestProps] {
	return func(p *pullRequestProps) (doyoucompute.Finalizer[pullRequestProps], error) {
		p.name = name

		return nil, nil
	}
}

// WithDescription overrides the description section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("Description")
//	section.WriteParagraph().Text("Added OAuth2 authentication")
//	pullrequest.WithDescription(section)
func WithDescription(description doyoucompute.Section) doyoucompute.OptionBuilder[pullRequestProps] {
	return func(p *pullRequestProps) (doyoucompute.Finalizer[pullRequestProps], error) {
		p.description = description

		return nil, nil
	}
}

// WithRelatedIssue overrides the related issue section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("Related issue")
//	section.WriteParagraph().Text("Fixes #123")
//	pullrequest.WithRelatedIssue(section)
func WithRelatedIssue(relatedIssue doyoucompute.Section) doyoucompute.OptionBuilder[pullRequestProps] {
	return func(p *pullRequestProps) (doyoucompute.Finalizer[pullRequestProps], error) {
		p.relatedIssue = relatedIssue

		return nil, nil
	}
}

// WithTesting overrides the testing section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("Testing")
//	section.WriteParagraph().Text("Added unit tests and integration tests")
//	pullrequest.WithTesting(section)
func WithTesting(testing doyoucompute.Section) doyoucompute.OptionBuilder[pullRequestProps] {
	return func(p *pullRequestProps) (doyoucompute.Finalizer[pullRequestProps], error) {
		p.testing = testing

		return nil, nil
	}
}

// DefaultName returns the default document name.
func DefaultName() string {
	return "Pull Request"
}

// DefaultDescription returns the default description section.
func DefaultDescription() doyoucompute.Section {
	section, _ := doyoucompute.SectionFactory("Description", func(s *doyoucompute.Section) error {
		s.WriteComment("What is this change and why are you making it?")
		return nil
	})

	return section
}

// DefaultRelatedIssue returns the default related issue section.
func DefaultRelatedIssue() doyoucompute.Section {
	section, _ := doyoucompute.SectionFactory("Related issue", func(s *doyoucompute.Section) error {
		s.WriteComment("Link to the relevant issue here.")
		return nil
	})

	return section
}

// DefaultTesting returns the default testing section.
func DefaultTesting() doyoucompute.Section {
	section, _ := doyoucompute.SectionFactory("How I tested", func(s *doyoucompute.Section) error {
		s.WriteComment("How did you test these changes?")
		return nil
	})

	return section
}

// New creates a new pull request document with default sections.
// Accepts zero or more option functions to customize the document.
//
// Example:
//
//	doc, err := pullrequest.New(
//		pullrequest.WithName("Bug Fix PR"),
//		pullrequest.WithTesting(customTestingSection),
//	)
func New(opts ...doyoucompute.OptionBuilder[pullRequestProps]) (doyoucompute.Document, error) {
	props := pullRequestProps{
		name:         DefaultName(),
		description:  DefaultDescription(),
		relatedIssue: DefaultRelatedIssue(),
		testing:      DefaultTesting(),
	}

	err := doyoucompute.ApplyOptions(&props, opts...)
	if err != nil {
		return doyoucompute.Document{}, err
	}

	// Validate
	if props.name == "" {
		return doyoucompute.Document{}, fmt.Errorf("pull request name cannot be empty")
	}

	return doyoucompute.DocumentFactory(props.name, func(d *doyoucompute.Document) error {
		d.AddSection(props.description)
		d.AddSection(props.relatedIssue)
		d.AddSection(props.testing)

		return nil
	})
}
