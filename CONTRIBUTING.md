# Contributing to OperaTree

First, thank you for considering a contribution to OperaTree. This project is built on the belief that a filesystem-first project operating system should be community-driven, extensible, and transparent.

This document explains how to contribute effectively — whether you're fixing a bug, adding a subject type, or proposing a new project template.

---

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Before You Start](#before-you-start)
- [Development Setup](#development-setup)
- [Project Structure](#project-structure)
- [Ways to Contribute](#ways-to-contribute)
  - [Bug Reports](#bug-reports)
  - [Feature Requests](#feature-requests)
  - [Adding a Subject Type](#adding-a-subject-type)
  - [Adding a Project Template](#adding-a-project-template)
  - [Improving Search](#improving-search)
- [Commit Convention](#commit-convention)
- [Pull Request Process](#pull-request-process)
- [License](#license)

---

## Code of Conduct

Be respectful. Be constructive. Disagreements about technical direction are welcome — personal attacks are not. Contributors are expected to maintain a professional and inclusive environment.

---

## Before You Start

- Check [open issues](https://github.com/hanymamdouh82/operatree/issues) to avoid duplicating effort
- For significant changes, open an issue first to discuss the approach before writing code
- Small fixes (typos, documentation, obvious bugs) can go straight to a PR

---

## Development Setup

**Requirements**

- Go 1.21 or higher
- Git
- Make

**Clone and build**

```bash
git clone https://github.com/hanymamdouh82/operatree.git
cd operatree
make build
```

**Run locally**

```bash
make run
```

**Run tests**

```bash
make test
```

**Install locally for manual testing**

```bash
make install
```

---

## Project Structure

```
operatree/
├── cmd/                    # Cobra CLI commands
│   ├── root.go             # root command and global flags
│   ├── bootstrap.go        # operatree bootstrap
│   ├── new.go              # operatree new
│   ├── find.go             # operatree find
│   ├── describe.go         # operatree desc
│   ├── summary.go          # operatree summary
│   ├── init.go             # operatree init
│   └── version.go          # operatree version
├── internal/
│   ├── config/             # config file management
│   ├── project/            # project struct, bootstrap, search, describe
│   ├── module/             # module struct and factories
│   ├── subject/            # subject struct, factory, interactive CLI
│   ├── metadata/           # tag and participant parsing utilities
│   ├── filesystem/         # filesystem helpers
│   └── help/               # embedded documentation
├── LICENSE
├── README.md
├── CONTRIBUTING.md
└── Makefile
```

---

## Ways to Contribute

### Bug Reports

Open an issue with:

- OperaTree version (`operatree version`)
- OS and terminal
- Steps to reproduce
- Expected vs actual behavior
- Relevant output or error messages

### Feature Requests

Open an issue describing:

- The problem you're trying to solve
- Your proposed solution
- Any alternatives you considered

---

### Adding a Subject Type

This is the most common and most valued contribution. Subject types are the extension point of OperaTree — they define what kinds of things a project can track.

**Step 1 — Add the constant**

```go
// internal/subject/subject.go
const (
    SubjectEvent     SubjectType = "event"
    SubjectTask      SubjectType = "task"
    SubjectTopic     SubjectType = "topic"
    SubjectObjective SubjectType = "objective"
    SubjectMeeting   SubjectType = "meeting"  // ← your new type
)
```

**Step 2 — Add type-specific fields to Subject**

All fields must use `omitempty` so they are invisible in YAML for subjects that don't use them. Never add required fields — every subject type shares the same struct.

```go
// internal/subject/subject.go
type Subject struct {
    // ... existing fields ...

    // Meeting-specific — omitempty keeps YAML clean for other types
    Agenda    string   `yaml:"agenda,omitempty"`
    MoMFile   string   `yaml:"momFile,omitempty"`
}
```

**Step 3 — Add the interactive form**

Add a branch in `interactiveCLI` for your type's specific fields:

```go
// internal/subject/cli.go
if st == SubjectMeeting {
    var agenda string

    err := huh.NewForm(
        huh.NewGroup(
            huh.NewText().
                Title("Agenda").
                Value(&agenda),
        ),
    ).Run()
    if err != nil {
        return err
    }

    s.Agenda = agenda
}
```

**Step 4 — Register in the CLI command**

```go
// cmd/new.go
var newCmd = &cobra.Command{
    ValidArgs: []cobra.Completion{"event", "task", "topic", "objective", "meeting"},
    // ...
}
```

Add a case in `newUnitEntity`:

```go
case "meeting":
    if err := newMeeting(&p); err != nil {
        log.Fatal(err)
    }
```

And implement `newMeeting` following the same pattern as `newEvent`.

**Step 5 — Update the find command type map**

```go
// cmd/find.go
var knownTypes = map[string]bool{
    "event":     true,
    "task":      true,
    "topic":     true,
    "objective": true,
    "meeting":   true,  // ← add here
}
```

**Step 6 — Update README**

Add your type to the Subject Types table in `README.md`.

That's it. No core changes, no breaking changes, no migration needed.

---

### Adding a Project Template

Templates define what a bootstrapped project looks like. The current template is `dev` — designed for software development companies.

**Step 1 — Create the template function**

```go
// internal/project/templates.go
func tmpltResearch(name string, bpth string) Project {
    ppth := path.Join(bpth, name)

    return Project{
        Name:    name,
        BaseDir: bpth,
        Modules: []module.Module{
            module.FactoryAdmin(ppth),
            module.FactoryResearch(ppth),
            module.FactoryData(ppth),
            module.FactoryDeliverables(ppth),
            module.FactoryArchive(ppth),
        },
    }
}
```

**Step 2 — Register in the template map**

```go
// internal/project/bootstrap.go
var templates = map[string]func(string, string) Project{
    "dev":      tmpltDev,
    "research": tmpltResearch,  // ← add here
}
```

**Step 3 — Update README**

Document your template in the README — what domain it targets and what modules it includes.

---

### Improving Search

The search pipeline is in `internal/project/search.go`. The current approach concatenates metadata fields into a `SearchStr` and runs fuzzy matching against it.

Potential improvements welcome:

- **Field weighting** — name matches should rank higher than note matches
- **Ranked results** — sort by relevance score not just match/no-match
- **Semantic search** — embedding-based similarity (this is on the roadmap as a commercial module, but algorithmic improvements to the fuzzy layer are always welcome)
- **Date-aware search** — `find last-week` or `find 2026-05`

If you're improving search, keep the `BuildSearchDB` / `SearchStr` pattern intact — it's the interface the rest of the system depends on.

---

## Commit Convention

OperaTree follows [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <short description>

[optional body]

[optional footer]
```

**Types:**

| Type       | When to use                                             |
| ---------- | ------------------------------------------------------- |
| `feat`     | New feature or subject type                             |
| `fix`      | Bug fix                                                 |
| `docs`     | Documentation only                                      |
| `refactor` | Code change that neither fixes a bug nor adds a feature |
| `test`     | Adding or updating tests                                |
| `chore`    | Build process, dependencies, tooling                    |

**Examples:**

```
feat(subject): add meeting subject type with agenda and MoM fields
fix(find): resolve ambiguity between type filter and search term
docs(contributing): add template contribution guide
refactor(search): extract walkModule into separate file
```

---

## Pull Request Process

1. Fork the repository
2. Create a branch from `main`: `git checkout -b feat/meeting-subject-type`
3. Make your changes following the patterns above
4. Run `make test` and ensure all tests pass
5. Run `make build` and test manually with `operatree init` and relevant commands
6. Commit using the convention above
7. Open a PR against `main` with a clear description of what you changed and why

**PR checklist:**

- [ ] Follows existing code patterns and naming conventions
- [ ] New subject types use `omitempty` on all type-specific fields
- [ ] New commands registered in both `cmd/` and `ValidArgs`
- [ ] README updated if user-facing behavior changed
- [ ] No hardcoded paths or test-specific defaults left in code

---

## License

By contributing to OperaTree, you agree that your contributions will be licensed under the [Apache License 2.0](LICENSE).

You retain copyright of your contributions. By submitting a pull request you grant the project maintainers the right to use your contribution under the project license.

> Note: OperaTree may in the future offer commercial modules built on top of the
> open source core. Community contributions to the open source CLI will always
> remain Apache 2.0 and will never be incorporated into commercial modules
> without explicit agreement from the contributor.
