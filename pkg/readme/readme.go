// Package readme provides a template for creating README.md documents.
//
// This package generates structured README files with sections for introduction,
// features, quick start, and optional additional content. Contributing and license
// sections are always included at the end.
//
// Basic usage:
//
//	props := readme.ReadmeProps{
//		Name:       "My Project",
//		Intro:      introParagraph,
//		Features:   featuresSection,
//		QuickStart: quickStartSection,
//	}
//	doc, err := readme.New(props, nil)
//	if err != nil {
//		// handle error
//	}
//
// With additional sections:
//
//	doc, err := readme.New(
//		props,
//		[]doyoucompute.Section{usageSection, examplesSection},
//		readme.WithName("Project README"),
//	)
package readme

import (
	"github.com/MoonMoon1919/doyoucompute"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/helpers"
)

// ReadmeProps defines the required and optional properties for a README document.
// Name, Intro, Features, and QuickStart must be provided when creating a new document.
// The contributing and license fields are automatically set to defaults.
type ReadmeProps struct {
	// Document name
	Name string
	// Introduction paragraph
	Intro doyoucompute.Paragraph
	// Features section
	Features doyoucompute.Section
	// Quick start section
	QuickStart   doyoucompute.Section
	contributing doyoucompute.Section
	license      doyoucompute.Section
}

// WithName overrides the document name.
//
// Example:
//
//	readme.WithName("My Awesome Project")
func WithName(name string) helpers.OptionsFunc[ReadmeProps] {
	return func(p *ReadmeProps) (helpers.PostEffect[ReadmeProps], error) {
		p.Name = name

		return nil, nil
	}
}

// DefaultContributing returns the default contributing section.
func DefaultContributing() doyoucompute.Section {
	return helpers.SectionFactory("Contributing", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteIntro().
			Text("See").
			Link("CONTRIBUTING", "./CONTRIBUTING.md").
			Text("for details.")

		return s
	})
}

// DefaultLicense returns the default license section.
func DefaultLicense() doyoucompute.Section {
	return helpers.SectionFactory("License", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteIntro().
			Text("See").
			Link("LICENSE", "./LICENSE").
			Text("for details.")

		return s
	})
}

// New creates a new README document with the provided properties and optional additional sections.
// The contributing and license sections are always added at the end of the document.
// Accepts zero or more option functions to customize the document.
//
// Example:
//
//	props := readme.ReadmeProps{
//		Name:       "My Project",
//		Intro:      introParagraph,
//		Features:   featuresSection,
//		QuickStart: quickStartSection,
//	}
//	doc, err := readme.New(
//		props,
//		[]doyoucompute.Section{usageSection, apiSection},
//		readme.WithName("Project Documentation"),
//	)
func New(props ReadmeProps, additionalSections []doyoucompute.Section, opts ...helpers.OptionsFunc[ReadmeProps]) (doyoucompute.Document, error) {
	sProps := ReadmeProps{
		Name:         props.Name,
		Intro:        props.Intro,
		Features:     props.Features,
		QuickStart:   props.QuickStart,
		contributing: DefaultContributing(),
		license:      DefaultLicense(),
	}

	err := helpers.ApplyOptions(&sProps, opts...)
	if err != nil {
		return doyoucompute.Document{}, err
	}

	return helpers.DocumentBuilder(sProps.Name, func(d *doyoucompute.Document) error {
		d.AddIntro(&sProps.Intro)
		d.AddSection(sProps.Features)
		d.AddSection(sProps.QuickStart)

		for _, section := range additionalSections {
			d.AddSection(section)
		}

		// Always put contributing and license last in the document
		d.AddSection(sProps.contributing)
		d.AddSection(sProps.license)

		return nil
	})
}
