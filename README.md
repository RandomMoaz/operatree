# OperaTree

> Your project operating system — built on your filesystem.

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.21+-00ADD8.svg)](https://golang.org)
[![Status](https://img.shields.io/badge/status-alpha-orange.svg)]()

OperaTree is a CLI tool that brings structure, searchability, and intelligence to how you manage projects — using your filesystem and plain YAML as the only storage. No database required. No vendor lock-in. No proprietary formats.

It works the way your OS already works: files and directories. Everything is human-readable, Git-friendly, and pipes naturally into standard UNIX tools.

---

## Philosophy

Most project management tools store your data in a database you don't control. When the tool dies, your data is trapped. When you switch tools, you lose history.

OperaTree takes the opposite approach:

- **Filesystem is the source of truth** — your project lives in directories you own
- **YAML is the metadata layer** — plain text, readable without any tool
- **The CLI is just an interface** — your data outlives any software

```
your-project/
├── 00_ADMIN/          # governance, contacts, templates
├── 01_EVENTS/         # visits, workshops, meetings
├── 02_PROJECT_MANAGEMENT/  # tasks, reports, risks
├── 03_LEGAL/          # contracts, NDAs, compliance
├── 04_RESEARCH/       # topics, objectives, summaries
├── 05_ENGINEERING/    # architecture, specs, decisions
├── 06_DATA/           # raw → staging → processed pipeline
├── 07_MEDIA_LIBRARY/  # shared reusable assets
├── 08_DELIVERABLES/   # final external outputs
└── 99_ARCHIVE/        # historical storage
```

Each subject (event, task, topic, objective) lives in its own directory with a `META.yaml` file that makes it searchable, filterable, and machine-readable.

---

## Features

- **Bootstrap projects** from opinionated templates in seconds
- **Create subjects** interactively — events, tasks, topics, objectives
- **Fuzzy search** across all metadata fields: names, tags, participants, notes, dates, locations
- **Describe and summarize** projects with colored terminal output
- **Pipe-friendly output** — chain with `grep`, `sed`, `cut`, and other UNIX tools
- **Self-documenting** — `operatree explain` renders the full directory philosophy in your terminal

---

## Installation

### From source

```bash
git clone https://github.com/hanymamdouh82/operatree.git
cd operatree
make install
```

### Using go install

```bash
go install github.com/hanymamdouh82/operatree@latest
```

---

## Quick Start

**1. Initialize configuration**

```bash
operatree init
```

Sets up `~/.config/operatree/operatree.yaml` with your standard projects directory.

**2. Bootstrap a new project**

```bash
operatree bootstrap myproject
```

Creates the full directory structure under your configured standard directory.

**3. Create a subject**

```bash
operatree new event
operatree new task
operatree new topic
operatree new objective
```

Launches an interactive form to capture metadata — name, date, tags, participants, notes, and more.

**4. Find anything**

```bash
operatree find                    # browse all subjects interactively
operatree find hany               # fuzzy search across all metadata
operatree find event hany         # filter by type, then fuzzy search
```

**5. Describe and summarize**

```bash
operatree desc                    # full project structure
operatree summary                 # counts, types, status breakdown
operatree explain                 # directory philosophy guide
```

---

## Commands

| Command | Description |
|---|---|
| `operatree init` | Initialize OperaTree configuration |
| `operatree bootstrap [name]` | Bootstrap a new project |
| `operatree new [type]` | Create a new subject interactively |
| `operatree find [type] [term]` | Fuzzy-find subjects across all metadata |
| `operatree desc` | Describe project structure |
| `operatree summary` | Project summary with counts and status |
| `operatree explain` | Print directory philosophy guide |
| `operatree version` | Print version |

### Global Flags

| Flag | Default | Description |
|---|---|---|
| `-d, --dest` | from config | Project directory |

---

## Subject Types

A subject is any trackable unit of work or knowledge within a project.

| Type | Purpose | Key Fields |
|---|---|---|
| `event` | A project activity — visit, workshop, meeting | date, location, participants, tags |
| `task` | A unit of work with a lifecycle | owner, status, related events |
| `topic` | A knowledge concept or domain area | tags, notes, related objective |
| `objective` | A goal driving research and decisions | status, findings, strategy |

All subjects share common fields: `name`, `date`, `tags`, `notes`. Type-specific fields use `omitempty` so they never pollute unrelated subjects.

---

## Search

OperaTree builds a rich search index from all subject metadata — not just names. A single fuzzy query matches against:

- Name
- Tags
- Participants
- Notes
- Date
- Location

```bash
operatree find cairo              # finds events in Cairo, notes mentioning Cairo, etc.
operatree find event emissions    # finds events tagged or named around emissions
```

The interactive finder shows a live preview of each subject as you navigate.

---

## Configuration

Config lives at `~/.config/operatree/operatree.yaml`:

```yaml
standardDir: /home/user/projects
projects:
  - name: myproject
    absPath: /home/user/projects/myproject
    template: dev
daemon:
  enabled: false
  host: localhost
  port: 7070
  dbDriver: sqlite
  dsn: ""
```

Projects are automatically registered when bootstrapped.

---

## Roadmap

OperaTree is in active development. The foundation is filesystem-first and stable. Future layers will be built on top without breaking it.

| Phase | Description | Status |
|---|---|---|
| CLI | Filesystem engine, YAML metadata, fuzzy search, interactive forms | 🚧 Alpha |
| Index sidecar | SQLite mirror for fast queries, no filesystem writes | 📋 Planned |
| Daemon | API over the index, sync engine, configuration-driven engine selection | 📋 Planned |
| Semantic search | Embeddings and vector search over subject metadata | 📋 Planned |
| Web platform | Multi-user, web UI, enterprise features (commercial) | 💡 Vision |

---

## Contributing

OperaTree welcomes contributions. The most impactful areas for community contribution are:

- **New subject types** — extend the type system with domain-specific subjects
- **New project templates** — add templates for research, legal, creative, and other domains
- **Search enhancements** — improve matching algorithms, ranking, relevance
- **Output formatters** — new ways to display and export project data

See [CONTRIBUTING.md](CONTRIBUTING.md) for the full guide.

---

## License

Copyright 2026 Hany Mamdouh

Licensed under the [Apache License, Version 2.0](LICENSE).

> Commercial modules (daemon, semantic search, web platform) will be released
> under a separate commercial license when available. The CLI tool will always
> remain Apache 2.0.
