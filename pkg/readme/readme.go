package readme

import (
	"github.com/MoonMoon1919/doyoucompute"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/helpers"
)

type ReadmeProps struct {
	Name         string
	Intro        doyoucompute.Paragraph
	Features     doyoucompute.Section
	QuickStart   doyoucompute.Section
	contributing doyoucompute.Section
	license      doyoucompute.Section
}

func WithName(name string) helpers.OptionsFunc[ReadmeProps] {
	return func(p *ReadmeProps) (helpers.PostEffect[ReadmeProps], error) {
		p.Name = name

		return nil, nil
	}
}

func DefaultContributing() doyoucompute.Section {
	return helpers.SectionFactory("Contributing", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteIntro().
			Text("See").
			Link("CONTRIBUTING", "./CONTRIBUTING.md").
			Text("for details.")

		return s
	})
}

func DefaultLicense() doyoucompute.Section {
	return helpers.SectionFactory("License", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteIntro().
			Text("See").
			Link("LICENSE", "./LICENSE").
			Text("for details.")

		return s
	})
}

func New(props ReadmeProps, additionalSections []doyoucompute.Section, opts ...helpers.OptionsFunc[ReadmeProps]) (doyoucompute.Document, error) {
	sProps := ReadmeProps{
		Name:         props.Name,
		Intro:        props.Intro,
		Features:     props.Features,
		QuickStart:   props.QuickStart,
		contributing: DefaultContributing(),
		license:      DefaultLicense(),
	}

	err := helpers.ApplyOptions(sProps, opts...)
	if err != nil {
		return doyoucompute.Document{}, err
	}

	return helpers.DocumentBuilder(sProps.Name, func(d *doyoucompute.Document) error {
		d.AddIntro(&sProps.Intro)
		d.AddSection(sProps.Features)
		d.AddSection(sProps.QuickStart)
		d.AddSection(sProps.contributing)
		d.AddSection(sProps.license)

		for _, section := range additionalSections {
			d.AddSection(section)
		}

		return nil
	})
}
