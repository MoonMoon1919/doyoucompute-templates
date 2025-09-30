# DOYOUCOMPUTE-TEMPLATES

A collection of common documents created by [doyoucompute.](https://github.com/MoonMoon1919/doyoucompute)

## Features

This package includes methods for creating the following documents:

- README
- Contributing
- Pull request template
- Bug report


These documents have a normalized structure to include sections that one would expect to see in the document For example, the Bug Report has exected and actual behavior, a section for example code, etc. Each document requires minimal inputs - in some cases, no input is required.

## Quickstart

### Installation

```bash
go get github.com/MoonMoon1919/doyoucompute-templates
```

#### Basic usage

All documents support the functional options pattern to override defaults If an input is required it is included as an attribute on the `New` method for the associated document.

```go
package samples

import (
	"fmt"

	"github.com/MoonMoon1919/doyoucompute"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/bugreport"
)

func Basics() {
	// With defaults
	br, err := bugreport.New()
	if err != nil {
		panic(err)
	}

	// With options
	expectedBehavior, err := doyoucompute.SectionFactory("Expected behavior",
		func(s *doyoucompute.Section) error {
			s.WriteComment("A comment explaining how to use the section")
			return nil
		},
	)
	if err != nil {
		panic(err)
	}

	codeSample, err := doyoucompute.SectionFactory("Code samples",
		func(s *doyoucompute.Section) error {
			s.WriteComment("A comment explaining how to use the section")
			s.WriteCodeBlock("go", []string{"# place code in here"}, doyoucompute.Static)

			return nil
		},
	)
	if err != nil {
		panic(err)
	}

	bugreportOptions, err := bugreport.New(
		bugreport.WithName("Bug report - name override"),
		bugreport.WithExpectedBehavior(expectedBehavior),
		bugreport.WithCodeSamples(codeSample),
	)
	if err != nil {
		panic(err)
	}

	fmt.Print(br)
	fmt.Print(bugreportOptions)
}

```

## Available documents

This package contains several different documents, each with configurable options

For additional example usage See the docs in [the docs and samples directory.](./internal)

### README

README containing configurable introductory paragraph, features and quickstart sections an option to insert other content and default license and contributing sections with options for overrides.

See [the module](./pkg/readme/readme.go) for full details.

### Bug Report

Bug Report template with Frontmatter for GitHub Issues. Contains defaults for expected/actual behavior, environment details, reproduction steps, code samples, and errors with options for overrides.

See [the module](./pkg/bugreport/bugreport.go) for full details.

### Pull Request

Pull Request template with default sections for description, issue link, and how it was tested with options for overrides.

See [the module](./pkg/pullrequest/pullrequest.go) for full details.

### Contributing

Contributing document containing configurable sections for getting started, contribution guidelines, writing docs, and reporting bugs with configurable overrides.

See [the module](./pkg/contributing/contributing.go) for full details.

## Disclaimers

This work does not represent the interests or technologies of any employer, past or present. It is a personal project only.

## Contributing

See [CONTRIBUTING](./CONTRIBUTING.md) for details.

## License

See [LICENSE](./LICENSE) for details.
