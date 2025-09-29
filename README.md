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
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/helpers"
)

func Basics() {
	// With defaults
	br, err := bugreport.New()
	if err != nil {
		panic(err)
	}

	// With options
	bugreportOptions, err := bugreport.New(
		bugreport.WithName("Bug report - name override"),
		bugreport.WithExpectedBehavior(
			helpers.SectionFactory("Expected behavior",
				func(s doyoucompute.Section) doyoucompute.Section {
					s.WriteComment("A comment explaining how to use the section")

					return s
				},
			),
		),
		bugreport.WithCodeSamples(
			helpers.SectionFactory("Code samples",
				func(s doyoucompute.Section) doyoucompute.Section {
					s.WriteComment("A comment explaining how to use the section")
					s.WriteCodeBlock("go", []string{"# place code in here"}, doyoucompute.Static)

					return s
				},
			),
		),
	)
	if err != nil {
		panic(err)
	}

	fmt.Print(br)
	fmt.Print(bugreportOptions)
}

```

## Disclaimers

This work does not represent the interests or technologies of any employer, past or present. It is a personal project only.

## Contributing

See [CONTRIBUTING](./CONTRIBUTING.md) for details.

## License

See [LICENSE](./LICENSE) for details.
