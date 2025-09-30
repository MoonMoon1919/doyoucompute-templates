package docs

import (
	"os"

	"github.com/MoonMoon1919/doyoucompute"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/readme"
)

func features() (doyoucompute.Section, error) {
	return doyoucompute.SectionFactory("Features", func(s *doyoucompute.Section) error {
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

		return nil
	})
}

func basicUsage() (doyoucompute.Section, error) {
	sample, err := os.ReadFile("./internal/samples/basics.go")
	if err != nil {
		return doyoucompute.Section{}, err
	}

	return doyoucompute.SectionFactory("Basic usage", func(s *doyoucompute.Section) error {
		s.WriteIntro().
			Text("All documents support the functional options pattern to override defaults").
			Text("If an input is required it is included as an attribute on the").
			Code("New").
			Text("method for the associated document.")

		s.WriteCodeBlock("go", []string{string(sample)}, doyoucompute.Static)

		return nil
	})
}

func quickstart() (doyoucompute.Section, error) {
	basics, err := basicUsage()
	if err != nil {
		return doyoucompute.Section{}, err
	}

	return doyoucompute.SectionFactory("Quickstart", func(s *doyoucompute.Section) error {
		installation := s.CreateSection("Installation")
		installation.WriteCodeBlock("bash", []string{"go get github.com/MoonMoon1919/doyoucompute-templates"}, doyoucompute.Static)

		installation.AddSection(basics)

		return nil
	})
}

func availableDocs() (doyoucompute.Section, error) {
	return doyoucompute.SectionFactory("Available documents", func(s *doyoucompute.Section) error {
		s.WriteIntro().
			Text("This package contains several different documents, each with configurable options")

		s.WriteParagraph().
			Text("For additional example usage").
			Text("See the docs in").
			Link("the docs and samples directory.", "./internal")

		readmeSection := s.CreateSection("README")
		readmeSection.WriteIntro().
			Text("README containing configurable introductory paragraph, features and quickstart sections").
			Text("an option to insert other content and default license and contributing sections with options for overrides.")

		readmeSection.WriteParagraph().Text("See").Link("the module", "./pkg/readme/readme.go").Text("for full details.")

		bugreportSection := s.CreateSection("Bug Report")
		bugreportSection.WriteIntro().
			Text("Bug Report template with Frontmatter for GitHub Issues.").
			Text("Contains defaults for expected/actual behavior, environment details,").
			Text("reproduction steps, code samples, and errors with options for overrides.")

		bugreportSection.WriteParagraph().Text("See").Link("the module", "./pkg/bugreport/bugreport.go").Text("for full details.")

		pullrequestSection := s.CreateSection("Pull Request")
		pullrequestSection.WriteIntro().
			Text("Pull Request template with default sections for description, issue link, and how it was tested with options for overrides.")

		pullrequestSection.WriteParagraph().Text("See").Link("the module", "./pkg/pullrequest/pullrequest.go").Text("for full details.")

		contributingSection := s.CreateSection("Contributing")
		contributingSection.WriteIntro().
			Text("Contributing document containing configurable sections for getting started, contribution guidelines, writing docs, and reporting bugs").
			Text("with configurable overrides.")

		contributingSection.WriteParagraph().Text("See").Link("the module", "./pkg/contributing/contributing.go").Text("for full details.")

		return nil
	})
}

func disclaimers() (doyoucompute.Section, error) {
	return doyoucompute.SectionFactory("Disclaimers", func(s *doyoucompute.Section) error {
		s.WriteIntro().
			Text("This work does not represent the interests or technologies of any employer, past or present.").
			Text("It is a personal project only.")

		return nil
	})
}

func ReadMe() (doyoucompute.Document, error) {
	quickstartSection, err := quickstart()
	if err != nil {
		return doyoucompute.Document{}, err
	}

	featuresSection, err := features()
	if err != nil {
		return doyoucompute.Document{}, err
	}

	disclaimerSection, err := disclaimers()
	if err != nil {
		return doyoucompute.Document{}, err
	}

	availableDocsSection, err := availableDocs()
	if err != nil {
		return doyoucompute.Document{}, err
	}

	return readme.New(
		readme.ReadmeProps{
			Name: "DOYOUCOMPUTE-TEMPLATES",
			Intro: *doyoucompute.NewParagraph().
				Text("A collection of common documents created by").
				Link("doyoucompute.", "https://github.com/MoonMoon1919/doyoucompute"),
			Features:   featuresSection,
			QuickStart: quickstartSection,
		},
		[]doyoucompute.Section{
			availableDocsSection,
			disclaimerSection,
		},
	)
}
