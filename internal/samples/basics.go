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
