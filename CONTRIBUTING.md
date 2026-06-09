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
  - [Version Control & Backup Backends](#version-control--backup-backends)
- [On the Use of AI Tools](#on-the-use-of-ai-tools)
- [Commit Convention](#commit-convention)
- [Pull Request Process](#pull-request-process)
- [License](#license)

---

## Code of Conduct

Be respectful. Be constructive. Disagreements about technical direction are welcome — personal attacks are not. Contributors are expected to maintain a professional and inclusive environment.

OperaTree is a deliberately minimal tool. The filesystem-first philosophy is a constraint, not a starting point — and that constraint applies to features too. Contributions are expected to reflect genuine understanding of the project's direction, not a desire to add surface area for its own sake.

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
│   ├── archive.go          # operatree archive
│   ├── bootstrap.go        # operatree bootstrap
│   ├── default.go          # operatree default
│   ├── describe.go         # operatree desc
│   ├── find.go             # operatree find
│   ├── help.go             # operatree help
│   ├── init.go             # operatree init
│   ├── jump.go             # operatree jump
│   ├── metadata.go         # operatree metadata
│   ├── new.go              # operatree new
│   ├── open.go             # operatree open
│   ├── root.go             # root command, global flags, project dir resolution
│   ├── summary.go          # operatree summary
│   ├── sync.go             # operatree sync
│   ├── track.go            # operatree track
│   ├── untrack.go          # operatree untrack
│   └── version.go          # operatree version
├── internal/
│   ├── activitylog/        # append-only activity log
│   ├── filesystem/         # filesystem helpers
│   ├── help/               # embedded documentation
│   ├── metadata/           # tag and participant parsing utilities
│   ├── runner/             # external binaries runner
│   ├── templates/          # templates yaml files, template loaders, listing, descriptions
│   └── ui/                 # ui utilities
├── pkg/
│   ├── config/             # config file management
│   ├── module/             # module struct and factories
│   ├── project/            # project struct, bootstrap, search, describe, sync
│   └── subject/            # subject struct, factory, interactive CLI
├── demo/                   # VHS tape and recorded demo
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
// pkg/subject/types.go
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
// pkg/subject/subject.go
type Subject struct {
    // ... existing fields ...

    // Meeting-specific — omitempty keeps YAML clean for other types
    Agenda  string `yaml:"agenda,omitempty"`
    MoMFile string `yaml:"momFile,omitempty"`
}
```

**Step 3 — Add the interactive form**

Add a branch in `interactiveCLI` for your type's specific fields:

```go
// pkg/subject/interactive.go
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

Add a map entry in `argsToSubject`:

```go
// cmd/new.go
var (
	argToSubject map[string]subject.SubjectType = map[string]subject.SubjectType{
		"event":     subject.SubjectEvent,
		"task":      subject.SubjectTask,
		"topic":     subject.SubjectTopic,
		"objective": subject.SubjectObjective,
		"meeting":   subject.SubjectMeeting,
	}
)
```

Add a map entry in `argsToSubject` for which module the subject belongs to:

```go
// pkg/project/types.go
// SubjectModuleMap maps each subject type to its corresponding storage modules
// It is used also to auto generate CMD accepted args
var SubjectModuleMap = map[subject.SubjectType]module.ModuleType{
	subject.SubjectEvent:     module.ModuleEvents,
	subject.SubjectTask:      module.ModuleTasks,
	subject.SubjectTopic:     module.ModuleTopics,
	subject.SubjectObjective: module.ModuleObjectives,
	subject.SubjectMeeting:   module.ModuleEvents,
}
```

**Step 4 - Update CMD flags**

Add new struct fields to `addCmd` command flags

```go
	addCmd.Flags().StringVar(&ns.Agenda, "agenda", "", "Meeting agenda description")
	addCmd.Flags().StringVar(&ns.MomFile, "mom-file", "", "minitues of meeting file link")
```

**Step 5 — Update README**

Add your type to the Subject Types table in `README.md`.

Update `doc/<section_.md>` for relevant section including your new subject.

That's it. No core changes, no breaking changes, no migration needed. Existing `META.yaml` files are unaffected since new fields use `omitempty`.

---

### Adding a Project Template

Templates define what a bootstrapped project looks like. The current templates are:

- `dev`: designed for software development companies
- `general`: designed to general purpose and minimal project structre
- `consulting`: for client engagement work
- `research`: for academic work, R&D, and analytical projects.

**IMPORTANT**

adding new template should follow approval process. Since the primary goal of `OperaTree` is to align with Porject Management standards,
adding a new template and its included modules should be discussed first for its feasibility and business adapatability. To propose a new template,

contributor should provide

- a clear explanation for template
- proposed full template yaml file with all descriptions clearly included
- reason for included modules and subdirs
- outcomes of the template
- target audience.

Any PR for a new template without linked discussion will be rejected.

After approval on template, follow the steps:

**Step 1 — Create the template file**

```yml
---
# /internal/templates/template_<name>.yml
name: dummy
description: dummy description
modules:
  - type: admin
    name: admin
    description: admin description
    modules: []
    subDirs:
      - sub1
  - type: project_management
    name: project_management
    description: project_management description
    modules:
      - type: tasks
        name: tasks
        description: tasks description
        subDirs:
          - sub2
```

**Step 2 — Register in the template map**

```go
// internal/templates/types.go
var Templates tmpltMap = tmpltMap{
	"general":           "template_general.yml",
	"dev":               "template_dev.yml",
	"consulting":        "template_consulting.yml",
	"research":          "template_research.yml",
	"<template_name>":   "template_<template_name>.yml",
}
```

**Step 3 — Update README**

Document your template — what domain it targets and what modules it includes.

---

### Improving Search

The search pipeline lives in `pkg/project/search_builder.go`. The current approach concatenates metadata fields into a `SearchStr` per subject and runs fuzzy matching against it. The full project tree is walked recursively via `walkModule`, building a flat `[]SearchDB` with `ModulePath` breadcrumbs.

The same search is used by `find`, `metadata`, and `open` — improvements benefit all three commands automatically.

Potential improvements welcome:

- **CLI only** — search based on CLI flags avoiding interactive CLI
- **Field weighting** — name matches should rank higher than note matches
- **Ranked results** — sort by relevance score not just match/no-match
- **Date-aware search** — `find last-week` or `find 2026-05`
- **Semantic search** — embedding-based similarity (on the roadmap as a commercial module, but algorithmic improvements to the fuzzy layer are always welcome)

If you're improving search, keep the `BuildSearchDB` / `SearchStr` / `walkModule` pattern intact — it's the interface the rest of the system depends on.

---

### Version Control & Backup Backends

OperaTree is designed to work naturally with external version control and backup tools. The filesystem layout is the contract — how changes are tracked and protected is a pluggable concern.

The model has two layers that work together:

**Layer 1 — Change detection:** a file watcher that monitors the project directory and triggers an action when files are added, modified, or deleted. OperaTree does not ship a watcher — it is designed to integrate with existing tools such as `watchexec`, `inotifywait`, or `fswatch`.

**Layer 2 — The action:** what runs when a change is detected. Possible backends include Git (local or remote), rsync, Syncthing, rclone, and others.

**Contribution ideas in this area:**

- A `operatree watch` command that wraps a configurable watcher + action pair
- Built-in Git integration — `operatree commit` to snapshot the current project state
- Config-driven backend selection (git, rsync, syncthing) under a new `backup` config section
- Git hook templates that users can install into their project with a single command
- Documentation and example scripts for common watcher + backend combinations

**Design constraints to respect:**

- The watcher and action must always be optional — OperaTree works fully without them
- No backup backend should be a required dependency
- The filesystem layout must never be modified to accommodate a specific backend
- All integration should be additive — existing projects must not need migration

If you're working on this area, open an issue first to discuss the approach. This is a high-impact contribution surface and coordination matters.

---

## On the Use of AI Tools

AI tools are welcome here — for brainstorming, exploring approaches, drafting, or checking your own work. Many contributors use them productively and that's fine.

What we ask is that **you** remain the author of your contribution. This means:

- You understand every line you submit and can explain the reasoning behind it
- The code follows the patterns already established in the codebase — not patterns an AI defaulted to
- You're not submitting a feature because it was easy to generate, but because it genuinely belongs in the project

OperaTree grows slowly and deliberately. A contribution that adds weight without a clear purpose will be declined regardless of how it was written. If a reviewer asks "why this approach?", the answer should come from you.

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
feat(cmd): add track and untrack commands
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
7. Open a PR against `main` with a clear description of what changed and why

**PR checklist:**

- [ ] Follows existing code patterns and naming conventions
- [ ] I can explain the reasoning behind every change in this PR
- [ ] New subject types use `omitempty` on all type-specific fields
- [ ] New subject types added to `knownTypes` map
- [ ] New commands registered in both `cmd/` and `ValidArgs` where applicable
- [ ] README updated if user-facing behavior changed
- [ ] No hardcoded paths or test-specific defaults left in code
- [ ] Activity log called for any action that creates, edits, or archives a subject

---

## License

By contributing to OperaTree, you agree that your contributions will be licensed under the [Apache License 2.0](LICENSE).

You retain copyright of your contributions. By submitting a pull request you grant the project maintainers the right to use your contribution under the project license.

> Note: OperaTree may in the future offer commercial modules built on top of the
> open source core. Community contributions to the open source CLI will always
> remain Apache 2.0 and will never be incorporated into commercial modules
> without explicit agreement from the contributor.
