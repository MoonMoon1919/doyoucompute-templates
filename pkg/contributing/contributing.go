// Package contributing provides a template for creating CONTRIBUTING.md documents.
//
// This package generates structured contribution guidelines for open source projects,
// including sections on setup, development workflow, submitting changes, reporting bugs,
// and writing documentation.
//
// Basic usage:
//
//	doc, err := contributing.New(
//		"https://github.com/username/project",
//		"https://github.com/username/project/issues",
//	)
//	if err != nil {
//		// handle error
//	}
//
// Customizing sections:
//
//	doc, err := contributing.New(
//		"https://github.com/username/project",
//		"https://github.com/username/project/issues",
//		contributing.WithName("Contributing Guidelines"),
//		contributing.WithSetup(customSetupSection),
//	)
package contributing

import (
	"fmt"
	"strings"

	"github.com/MoonMoon1919/doyoucompute"
)

type contributingProps struct {
	name            string
	projectUrl      string
	issueTrackerUrl string
	gettingStarted  doyoucompute.Section
	choseATask      doyoucompute.Section
	setup           doyoucompute.Section
	development     doyoucompute.Section
	submissions     doyoucompute.Section
	writingDocs     doyoucompute.Section
	reportingbugs   doyoucompute.Section
	license         doyoucompute.Section
}

// DefaultName returns the default document name.
func DefaultName() string {
	return "Contributing"
}

// WithName overrides the document name.
//
// Example:
//
//	contributing.WithName("Contributing Guidelines")
func WithName(name string) doyoucompute.OptionBuilder[contributingProps] {
	return func(p *contributingProps) (doyoucompute.Finalizer[contributingProps], error) {
		p.name = name

		return nil, nil
	}
}

// WithProjectUrl overrides the project URL.
//
// Example:
//
//	contributing.WithProjectUrl("https://github.com/username/project")
func WithProjectUrl(url string) doyoucompute.OptionBuilder[contributingProps] {
	return func(p *contributingProps) (doyoucompute.Finalizer[contributingProps], error) {
		p.projectUrl = url

		return nil, nil
	}
}

// WithIssueTrackerUrl overrides the issue tracker URL and updates dependent sections.
//
// Example:
//
//	contributing.WithIssueTrackerUrl("https://github.com/username/project/issues")
func WithIssueTrackerUrl(url string) doyoucompute.OptionBuilder[contributingProps] {
	return func(p *contributingProps) (doyoucompute.Finalizer[contributingProps], error) {
		p.issueTrackerUrl = url

		return func(p *contributingProps) error {
			p.choseATask = DefaultChoseATask(url)
			p.reportingbugs = DefaultReportingBugs(url)

			return nil
		}, nil
	}
}

// WithGettingStarted overrides the getting started section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("Getting started")
//	section.WriteParagraph().Text("Read our documentation first")
//	contributing.WithGettingStarted(section)
func WithGettingStarted(gettingStarted doyoucompute.Section) doyoucompute.OptionBuilder[contributingProps] {
	return func(p *contributingProps) (doyoucompute.Finalizer[contributingProps], error) {
		p.gettingStarted = gettingStarted

		return nil, nil
	}
}

// WithChoseATask overrides the task selection section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("Find a task")
//	section.WriteParagraph().Text("Check our project board")
//	contributing.WithChoseATask(section)
func WithChoseATask(task doyoucompute.Section) doyoucompute.OptionBuilder[contributingProps] {
	return func(p *contributingProps) (doyoucompute.Finalizer[contributingProps], error) {
		p.choseATask = task

		return nil, nil
	}
}

// WithSetup overrides the setup section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("Setup")
//	section.WriteParagraph().Text("Install Docker first")
//	contributing.WithSetup(section)
func WithSetup(setup doyoucompute.Section) doyoucompute.OptionBuilder[contributingProps] {
	return func(p *contributingProps) (doyoucompute.Finalizer[contributingProps], error) {
		p.setup = setup

		return nil, nil
	}
}

// WithDevelopment overrides the development workflow section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("Development")
//	section.WriteParagraph().Text("Use our pre-commit hooks")
//	contributing.WithDevelopment(section)
func WithDevelopment(development doyoucompute.Section) doyoucompute.OptionBuilder[contributingProps] {
	return func(p *contributingProps) (doyoucompute.Finalizer[contributingProps], error) {
		p.development = development

		return nil, nil
	}
}

// WithSubmissions overrides the submissions section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("Submitting PRs")
//	section.WriteParagraph().Text("Ensure CI passes before requesting review")
//	contributing.WithSubmissions(section)
func WithSubmissions(submissions doyoucompute.Section) doyoucompute.OptionBuilder[contributingProps] {
	return func(p *contributingProps) (doyoucompute.Finalizer[contributingProps], error) {
		p.submissions = submissions

		return nil, nil
	}
}

// WithWritingDocs overrides the documentation writing section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("Documentation")
//	section.WriteParagraph().Text("We use MkDocs for documentation")
//	contributing.WithWritingDocs(section)
func WithWritingDocs(docs doyoucompute.Section) doyoucompute.OptionBuilder[contributingProps] {
	return func(p *contributingProps) (doyoucompute.Finalizer[contributingProps], error) {
		p.writingDocs = docs

		return nil, nil
	}
}

// WithReportingbugs overrides the bug reporting section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("Bug Reports")
//	section.WriteParagraph().Text("Include system information")
//	contributing.WithReportingbugs(section)
func WithReportingbugs(bugs doyoucompute.Section) doyoucompute.OptionBuilder[contributingProps] {
	return func(p *contributingProps) (doyoucompute.Finalizer[contributingProps], error) {
		p.reportingbugs = bugs

		return nil, nil
	}
}

// WithLicense overrides the license section.
// This replaces the entire section, including the title.
//
// Example:
//
//	section := doyoucompute.NewSection("License")
//	section.WriteParagraph().Text("MIT License applies")
//	contributing.WithLicense(section)
func WithLicense(license doyoucompute.Section) doyoucompute.OptionBuilder[contributingProps] {
	return func(p *contributingProps) (doyoucompute.Finalizer[contributingProps], error) {
		p.license = license

		return nil, nil
	}
}

// DefaultGettingStarted returns the default getting started section.
func DefaultGettingStarted() doyoucompute.Section {
	section, _ := doyoucompute.SectionFactory("Getting started", func(s *doyoucompute.Section) error {
		s.WriteParagraph().
			Text("Read the").
			Link("README", "README.md").
			Text("to understand the project's scope and purpose.")

		return nil
	})

	return section
}

// DefaultChoseATask returns the default task selection section.
func DefaultChoseATask(issueTrackerUrl string) doyoucompute.Section {
	section, _ := doyoucompute.SectionFactory("Find a task", func(s *doyoucompute.Section) error {
		s.WriteParagraph().
			Text("Browse the").
			Link("issue tracker", issueTrackerUrl).
			Text(" to see what's being worked on and what needs attention.")

		s.WriteParagraph().
			Text("Don't see anything that interests you? Feel free to open a new issue to:")

		suggestionsList := s.CreateList(doyoucompute.BULLET)
		suggestionsList.Append("Suggest new features or improvements")
		suggestionsList.Append("Report documentation gaps or unclear examples")
		suggestionsList.Append("Propose improvements")
		suggestionsList.Append("Ask questions about implementation details")

		return nil
	})

	return section
}

// DefaultLicense returns the default license section.
func DefaultLicense() doyoucompute.Section {
	section, _ := doyoucompute.SectionFactory("License", func(s *doyoucompute.Section) error {
		s.WriteParagraph().
			Text("By contributing, you agree that your contributions will be licensed under the project's").
			Link("License.", "./LICENSE")

		return nil
	})

	return section
}

// DefaultWritingDocs returns the default documentation writing section.
func DefaultWritingDocs() doyoucompute.Section {
	section, _ := doyoucompute.SectionFactory("Writing documentation", func(s *doyoucompute.Section) error {
		s.WriteParagraph().
			Text("Read the").
			Link("README", "./README.md").
			Text("to understand the project's structure and how it's used.")

		// Identify areas for improvement
		s.WriteParagraph().
			Text("Look for documentation that is unclear, incomplete, or outdated and update it.")

		return nil
	})

	return section
}

// DefaultReportingBugs returns the default bug reporting section.
func DefaultReportingBugs(issueTrackerUrl string) doyoucompute.Section {
	section, _ := doyoucompute.SectionFactory("Reporting bugs", func(s *doyoucompute.Section) error {
		checkingSection := s.CreateSection("Checking for Existing Reports")
		checkingSection.WriteParagraph().
			Text("Before reporting a new bug, search the").
			Link("issue tracker", issueTrackerUrl).
			Text("to see if someone else has already reported the same issue.").
			Text("Check both open and closed issues - the bug might have been fixed in a recent version.")

		creatingSection := s.CreateSection("Reporting new bugs")

		creatingSection.WriteParagraph().
			Text("If you can't find an existing report, create a new issue and fill out the bug report form.")

		return nil
	})

	return section
}

// DefaultOpenSourceGoSetupGuidelines returns the default setup section for Go projects.
func DefaultOpenSourceGoSetupGuidelines(projectUrl string, projectName string) doyoucompute.Section {
	section, _ := doyoucompute.SectionFactory("Setting Up Your Development Environment", func(s *doyoucompute.Section) error {
		s.WriteParagraph().
			Text("First, fork the repository on GitHub at").
			Link(projectUrl, projectUrl).
			Text(" by clicking the \"Fork\" button.")

		s.WriteParagraph().
			Text("Then clone your forked repository to your local machine:")

		s.WriteCodeBlock("bash", []string{fmt.Sprintf("git clone <your_fork_url> %s", projectName)}, doyoucompute.Static)
		s.WriteCodeBlock("bash", []string{fmt.Sprintf("cd %s", projectName)}, doyoucompute.Static)

		s.WriteParagraph().
			Text("Install dependencies and verify you can run the tests:")

		s.WriteCodeBlock("bash", []string{"go mod tidy"}, doyoucompute.Static)
		s.WriteCodeBlock("bash", []string{"go test ./..."}, doyoucompute.Static)

		return nil
	})

	return section
}

// DefaultOpenSourceGoDevelopmentGuidelines returns the default development workflow section for Go projects.
func DefaultOpenSourceGoDevelopmentGuidelines() doyoucompute.Section {
	section, _ := doyoucompute.SectionFactory("Development Workflow", func(s *doyoucompute.Section) error {
		s.WriteParagraph().
			Text("Create a new branch for your feature or bug fix:")

		s.WriteCodeBlock("bash", []string{"git checkout -b feature/my-awesome-feature"}, doyoucompute.Static)

		s.WriteParagraph().
			Text("Make your changes and add tests for new functionality. Run tests to ensure changes work as expected:")

		s.WriteCodeBlock("bash", []string{"go test ./..."}, doyoucompute.Static)

		s.WriteParagraph().
			Text("If you're adding new features, consider adding example usage in the examples directory.")

		return nil
	})

	return section
}

// DefaultOpenSourceSubmittingGuidelines returns the default submission guidelines section.
func DefaultOpenSourceSubmittingGuidelines() doyoucompute.Section {
	section, _ := doyoucompute.SectionFactory("Submitting your changes", func(s *doyoucompute.Section) error {
		s.WriteParagraph().
			Text("Once you're satisfied with your changes, commit them with a descriptive message:")

		s.WriteCodeBlock("bash", []string{"git add ."}, doyoucompute.Static)
		s.WriteCodeBlock("bash", []string{"git commit -m \"Add feature: descriptive commit message\""}, doyoucompute.Static)

		s.WriteParagraph().
			Text("Push your changes to your forked repository:")

		s.WriteCodeBlock("bash", []string{"git push origin feature/my-awesome-feature"}, doyoucompute.Static)

		s.WriteParagraph().
			Text("Finally, create a pull request:")

		submissionSteps := s.CreateList(doyoucompute.BULLET)
		submissionSteps.Append("Go to the original repository on GitHub")
		submissionSteps.Append("Click \"Compare & pull request\"")
		submissionSteps.Append("Provide a clear description of your changes")
		submissionSteps.Append("Reference any relevant issues using #issue-number")
		submissionSteps.Append("Wait for review and address any feedback")

		return nil
	})

	return section
}

// New creates a new contributing guidelines document with default sections for Go projects.
// Accepts zero or more option functions to customize the document.
//
// The projectUrl is used to generate setup instructions and is parsed to extract the project name.
// The issueTrackerUrl is used in task selection and bug reporting sections.
//
// Example:
//
//	doc, err := contributing.New(
//		"https://github.com/username/project",
//		"https://github.com/username/project/issues",
//		contributing.WithName("How to Contribute"),
//	)
func New(projectUrl, issueTrackerUrl string, opts ...doyoucompute.OptionBuilder[contributingProps]) (doyoucompute.Document, error) {
	if projectUrl == "" {
		return doyoucompute.Document{}, fmt.Errorf("projectUrl cannot be empty")
	}
	if issueTrackerUrl == "" {
		return doyoucompute.Document{}, fmt.Errorf("issueTrackerUrl cannot be empty")
	}

	projectNameSplitter := strings.Split(projectUrl, "/")
	projectName := projectNameSplitter[len(projectNameSplitter)-1]

	// Validate project name was extracted
	if projectName == "" {
		return doyoucompute.Document{}, fmt.Errorf("could not extract project name from projectUrl: %s", projectUrl)
	}

	props := contributingProps{
		name:            DefaultName(),
		projectUrl:      projectUrl,
		issueTrackerUrl: issueTrackerUrl,
		gettingStarted:  DefaultGettingStarted(),
		choseATask:      DefaultChoseATask(issueTrackerUrl),
		setup:           DefaultOpenSourceGoSetupGuidelines(projectUrl, projectName),
		development:     DefaultOpenSourceGoDevelopmentGuidelines(),
		submissions:     DefaultOpenSourceSubmittingGuidelines(),
		writingDocs:     DefaultWritingDocs(),
		reportingbugs:   DefaultReportingBugs(issueTrackerUrl),
		license:         DefaultLicense(),
	}

	err := doyoucompute.ApplyOptions(&props, opts...)
	if err != nil {
		return doyoucompute.Document{}, err
	}

	// Validate props after options applied
	if props.name == "" {
		return doyoucompute.Document{}, fmt.Errorf("contributing guide name cannot be empty")
	}

	return doyoucompute.DocumentFactory(props.name, func(d *doyoucompute.Document) error {
		// Apply hierarchy here to improve flexibility of doc content using options
		props.gettingStarted.AddSection(props.choseATask)
		d.AddSection(props.gettingStarted)

		guidelines := d.CreateSection("Contribution guidelines")
		codeContributions := guidelines.CreateSection("Code contributions")
		codeContributions.AddSection(props.setup)
		codeContributions.AddSection(props.development)
		codeContributions.AddSection(props.submissions)

		guidelines.AddSection(props.reportingbugs)
		guidelines.AddSection(props.writingDocs)

		d.AddSection(props.license)

		return nil
	})
}
