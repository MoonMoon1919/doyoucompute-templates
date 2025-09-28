package contributing

import (
	"fmt"
	"strings"

	"github.com/MoonMoon1919/doyoucompute"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/helpers"
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

func WithName(name string) helpers.OptionsFunc[contributingProps] {
	return func(p *contributingProps) (helpers.PostEffect[contributingProps], error) {
		p.name = name

		return nil, nil
	}
}

func WithProjectUrl(url string) helpers.OptionsFunc[contributingProps] {
	return func(p *contributingProps) (helpers.PostEffect[contributingProps], error) {
		p.projectUrl = url

		return nil, nil
	}
}

func WithIssueTrackerUrl(url string) helpers.OptionsFunc[contributingProps] {
	return func(p *contributingProps) (helpers.PostEffect[contributingProps], error) {
		p.issueTrackerUrl = url

		return func(p *contributingProps) error {
			p.choseATask = DefaultChoseATask(url)
			p.reportingbugs = DefaultReportingBugs(url)

			return nil
		}, nil
	}
}

func WithGettingStarted(gettingStarted doyoucompute.Section) helpers.OptionsFunc[contributingProps] {
	return func(p *contributingProps) (helpers.PostEffect[contributingProps], error) {
		p.gettingStarted = gettingStarted

		return nil, nil
	}
}

func DefaultName() string {
	return "Contributing"
}

func DefaultGettingStarted() doyoucompute.Section {
	return helpers.SectionFactory("Getting started", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteParagraph().
			Text("Read the").
			Link("README", "README.md").
			Text("to understand the project's scope and purpose.")

		return s
	})
}

func DefaultChoseATask(issueTrackerUrl string) doyoucompute.Section {
	return helpers.SectionFactory("Find a task", func(s doyoucompute.Section) doyoucompute.Section {
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

		return s
	})
}

func DefaultLicense() doyoucompute.Section {
	return helpers.SectionFactory("License", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteParagraph().
			Text("By contributing, you agree that your contributions will be licensed under the project's").
			Link("License.", "./LICENSE")

		return s
	})
}

func DefaultWritingDocs() doyoucompute.Section {
	return helpers.SectionFactory("Writing documentation", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteParagraph().
			Text("Read the").
			Link("README", "./README.md").
			Text("to understand the project's structure and how it's used.")

		// Identify areas for improvement
		s.WriteParagraph().
			Text("Look for documentation that is unclear, incomplete, or outdated and update it.")

		return s
	})
}

func DefaultReportingBugs(issueTrackerUrl string) doyoucompute.Section {
	return helpers.SectionFactory("Reporting bugs", func(s doyoucompute.Section) doyoucompute.Section {
		checkingSection := s.CreateSection("Checking for Existing Reports")
		checkingSection.WriteParagraph().
			Text("Before reporting a new bug, search the").
			Link("issue tracker", issueTrackerUrl).
			Text("to see if someone else has already reported the same issue.").
			Text("Check both open and closed issues - the bug might have been fixed in a recent version.")

		creatingSection := s.CreateSection("Reporting new bugs")

		creatingSection.WriteParagraph().
			Text("If you can't find an existing report, create a new issue and fill out the bug report form.")

		return s
	})
}

func DefaultOpenSourceGoSetupGuidelines(projectUrl string, projectName string) doyoucompute.Section {
	return helpers.SectionFactory("Setting Up Your Development Environment", func(s doyoucompute.Section) doyoucompute.Section {
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

		return s
	})
}

func DefaultOpenSourceGoDevelopmentGuidelines() doyoucompute.Section {
	return helpers.SectionFactory("Development Workflow", func(s doyoucompute.Section) doyoucompute.Section {
		s.WriteParagraph().
			Text("Create a new branch for your feature or bug fix:")

		s.WriteCodeBlock("bash", []string{"git checkout -b feature/my-awesome-feature"}, doyoucompute.Static)

		s.WriteParagraph().
			Text("Make your changes and add tests for new functionality. Run tests to ensure changes work as expected:")

		s.WriteCodeBlock("bash", []string{"go test ./..."}, doyoucompute.Static)

		s.WriteParagraph().
			Text("If you're adding new features, consider adding example usage in the examples directory.")

		return s
	})
}

func DefaultOpenSourceSubmittingGuidelines() doyoucompute.Section {
	return helpers.SectionFactory("Submitting your changes", func(s doyoucompute.Section) doyoucompute.Section {
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

		return s
	})
}

func New(projectUrl, issueTrackerUrl string, opts ...helpers.OptionsFunc[contributingProps]) (doyoucompute.Document, error) {
	projectNameSplitter := strings.Split(projectUrl, "/")
	projectName := projectNameSplitter[len(projectNameSplitter)-1]

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

	err := helpers.ApplyOptions(props, opts...)
	if err != nil {
		return doyoucompute.Document{}, err
	}

	return helpers.DocumentBuilder(props.name, func(d *doyoucompute.Document) error {
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
