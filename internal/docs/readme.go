package docs

import (
	"os"

	"github.com/MoonMoon1919/doyoucompute"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/helpers"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/readme"
)

func features() doyoucompute.Section {
	return helpers.SectionFactory("Features", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteIntro().
			Text("This package includes methods for creating the following documents:")

		featureList := s.CreateList(doyoucompute.BULLET)
		featureList.Append("README")
		featureList.Append("Contributing")
		featureList.Append("Pull request template")
		featureList.Append("Bug report")

		s.WriteParagraph().
			Text("These documents have a normalized structure to include sections that one would expect to see in the document").
			Text("For example, the Bug Report has exected and actual behavior, a section for example code, etc.").
			Text("Each document requires minimal inputs - in some cases, no input is required.")

		return s
	})
}

func basicUsage() (doyoucompute.Section, error) {
	sample, err := os.ReadFile("./internal/samples/basics.go")
	if err != nil {
		return doyoucompute.Section{}, err
	}

	return helpers.SectionFactory("Basic usage", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteIntro().
			Text("All documents support the functional options pattern to override defaults").
			Text("If an input is required it is included as an attribute on the").
			Code("New").
			Text("method for the associated document.")

		s.WriteCodeBlock("go", []string{string(sample)}, doyoucompute.Static)

		return s
	}), nil
}

func quickstart() (doyoucompute.Section, error) {
	basics, err := basicUsage()
	if err != nil {
		return doyoucompute.Section{}, err
	}

	return helpers.SectionFactory("Quickstart", func(s doyoucompute.Section) doyoucompute.Section {
		installation := s.CreateSection("Installation")
		installation.WriteCodeBlock("bash", []string{"go get github.com/MoonMoon1919/doyoucompute-templates"}, doyoucompute.Static)

		installation.AddSection(basics)

		return s
	}), nil
}

func disclaimers() doyoucompute.Section {
	return helpers.SectionFactory("Disclaimers", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteIntro().
			Text("This work does not represent the interests or technologies of any employer, past or present.").
			Text("It is a personal project only.")

		return s
	})
}

func ReadMe() (doyoucompute.Document, error) {
	quickstartSection, err := quickstart()
	if err != nil {
		return doyoucompute.Document{}, err
	}

	return readme.New(
		readme.ReadmeProps{
			Name: "DOYOUCOMPUTE-TEMPLATES",
			Intro: *doyoucompute.NewParagraph().
				Text("A collection of common documents created by").
				Link("doyoucompute.", "https://github.com/MoonMoon1919/doyoucompute"),
			Features:   features(),
			QuickStart: quickstartSection,
		},
		[]doyoucompute.Section{
			disclaimers(),
		},
	)
}
