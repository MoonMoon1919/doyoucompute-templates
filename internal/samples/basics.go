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
