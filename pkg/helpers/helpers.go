package helpers

import "github.com/MoonMoon1919/doyoucompute"

// TODO: Move this module to main pkg
// Post Effects are useful for two attributes that are tightly coupled
// E.g., name and frontmatter in a github issue template
type PostEffect[T any] func(p *T) error
type OptionsFunc[T any] func(p *T) (PostEffect[T], error)

// TODO: Move contentFunc, sectionFactory, DocumentApplier, and documentBuilder to main pkg
type ContentFunc func(s doyoucompute.Section) doyoucompute.Section

func SectionFactory(name string, contentFuncs ...ContentFunc) doyoucompute.Section {
	s := doyoucompute.NewSection(name)

	for _, cFunc := range contentFuncs {
		s = cFunc(s)
	}

	return s
}

type DocumentApplier func(d *doyoucompute.Document) error

func DocumentBuilder(name string, appliers ...DocumentApplier) (doyoucompute.Document, error) {
	document, err := doyoucompute.NewDocument(name)
	if err != nil {
		return doyoucompute.Document{}, nil
	}

	for _, applier := range appliers {
		applier(&document)
	}

	return document, nil
}

func ApplyOptions[T any](props *T, opts ...OptionsFunc[T]) error {
	for _, opt := range opts {
		postEffect, err := opt(props)
		if err != nil {
			return err
		}

		if postEffect != nil {
			if err := postEffect(props); err != nil {
				return err
			}
		}
	}

	return nil
}
