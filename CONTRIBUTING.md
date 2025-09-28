# Contributing

## Getting started

Read the [README](README.md) to understand the project's scope and purpose.

### Find a task

Browse the [issue tracker](https://github.com/MoonMoon1919/doyoucompute-templates/issues)  to see what's being worked on and what needs attention.

Don't see anything that interests you? Feel free to open a new issue to:

- Suggest new features or improvements
- Report documentation gaps or unclear examples
- Propose improvements
- Ask questions about implementation details


## Contribution guidelines

### Code contributions

#### Setting Up Your Development Environment

First, fork the repository on GitHub at [https://github.com/MoonMoon1919/doyoucompute-templates](https://github.com/MoonMoon1919/doyoucompute-templates)  by clicking the "Fork" button.

Then clone your forked repository to your local machine:

```bash
git clone <your_fork_url> doyoucompute-templates
```

```bash
cd doyoucompute-templates
```

Install dependencies and verify you can run the tests:

```bash
go mod tidy
```

```bash
go test ./...
```

#### Development Workflow

Create a new branch for your feature or bug fix:

```bash
git checkout -b feature/my-awesome-feature
```

Make your changes and add tests for new functionality. Run tests to ensure changes work as expected:

```bash
go test ./...
```

If you're adding new features, consider adding example usage in the examples directory.

#### Submitting your changes

Once you're satisfied with your changes, commit them with a descriptive message:

```bash
git add .
```

```bash
git commit -m "Add feature: descriptive commit message"
```

Push your changes to your forked repository:

```bash
git push origin feature/my-awesome-feature
```

Finally, create a pull request:

- Go to the original repository on GitHub
- Click "Compare & pull request"
- Provide a clear description of your changes
- Reference any relevant issues using #issue-number
- Wait for review and address any feedback


### Reporting bugs

#### Checking for Existing Reports

Before reporting a new bug, search the [issue tracker](https://github.com/MoonMoon1919/doyoucompute-templates/issues) to see if someone else has already reported the same issue. Check both open and closed issues - the bug might have been fixed in a recent version.

#### Reporting new bugs

If you can't find an existing report, create a new issue and fill out the bug report form.

### Writing documentation

Read the [README](./README.md) to understand the project's structure and how it's used.

Look for documentation that is unclear, incomplete, or outdated and update it.

## License

By contributing, you agree that your contributions will be licensed under the project's [License.](./LICENSE)
