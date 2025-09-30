package docs

import (
	"github.com/MoonMoon1919/doyoucompute"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/contributing"
)

func writingDocs() doyoucompute.Section {
	docsSection := doyoucompute.NewSection("Writing documentation")

	// Review existing documentation
	docsSection.WriteParagraph().
		Text("Read the").
		Link("README", "./README.md").
		Text("to understand the project's structure and how it's used.")

	// Identify areas for improvement
	docsSection.WriteParagraph().
		Text("Look for documentation that is unclear, incomplete, or outdated.")

	// Make the changes
	docsSection.WriteParagraph().
		Text("Update the appropriate file in the").
		Link("docs folder", "./internal/docs").
		Text("since we're using doyoucompute to generate documents.")

	return docsSection
}

func license() doyoucompute.Section {
	licenseSection := doyoucompute.NewSection("License")
	licenseSection.WriteParagraph().
		Text("By contributing, you agree that your contributions will be licensed under the project's").
		Link("MIT License.", "./LICENSE")

	return licenseSection
}

func Contributing() (doyoucompute.Document, error) {
	return contributing.New(
		"https://github.com/MoonMoon1919/doyoucompute-templates",
		"https://github.com/MoonMoon1919/doyoucompute-templates/issues",
		contributing.WithWritingDocs(writingDocs()),
		contributing.WithLicense(license()),
	)
}
