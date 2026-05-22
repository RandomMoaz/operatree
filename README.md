# OperaTree

> Your project operating system — built on your filesystem.

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.21+-00ADD8.svg)](https://golang.org)
[![Status](https://img.shields.io/badge/status-alpha-orange.svg)]()

OperaTree is a CLI tool that brings structure, searchability, and intelligence to how you manage projects — using your filesystem and plain YAML as the only storage. No database required. No vendor lock-in. No proprietary formats.

It works the way your OS already works: files and directories. Everything is human-readable, Git-friendly, and pipes naturally into standard UNIX tools.

**Setup**

![demo](demo/operatree-demo-1-setup.gif)

**Working with Subjects**

![demo](demo/operatree-demo-2-creating.gif)

**Power Features**

![demo](demo/operatree-demo-3-power.gif)

---

## Philosophy

Most project management tools store your data in a database you don't control. When the tool dies, your data is trapped. When you switch tools, you lose history.

OperaTree takes the opposite approach:

- **Filesystem is the source of truth** — your project lives in directories you own
- **YAML is the metadata layer** — plain text, readable without any tool
- **The CLI is just an interface** — your data outlives any software

```
your-project/
├── 00_ADMIN/               # governance, contacts, templates
├── 01_EVENTS/              # visits, workshops, meetings
├── 02_PROJECT_MANAGEMENT/  # tasks, reports, risks
├── 03_LEGAL/               # contracts, NDAs, compliance
├── 04_RESEARCH/            # topics, objectives, summaries
├── 05_ENGINEERING/         # architecture, specs, decisions
├── 06_DATA/                # raw → staging → processed pipeline
├── 07_MEDIA_LIBRARY/       # shared reusable assets
├── 08_DELIVERABLES/        # final external outputs
├── 99_ARCHIVE/             # historical storage
└── activity.log            # append-only audit trail
```

Each subject (event, task, topic, objective) lives in its own directory with a `META.yaml` file that makes it searchable, filterable, and machine-readable.

---

## Features

- **Bootstrap projects** from opinionated templates in seconds
- **Create subjects** interactively — events, tasks, topics, objectives
- **Fuzzy search** across all metadata fields: names, tags, participants, notes, dates, locations
- **Edit subject metadata** directly in your preferred editor
- **Archive subject** by moving to `99_ARCHIVE` directory
- **Open subject directories** in your preferred file manager
- **Sync** project metadata with manually edited subject files
- **Track and untrack** existing projects without bootstrapping
- **Describe and summarize** projects with colored terminal output
- **Default project** — set once, never pass `-d` again
- **Activity log** — every create, edit, and archive appended to `activity.log` at project root
- **Version control ready** — filesystem layout is designed to work naturally with Git, rsync, Syncthing, or any file watcher
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

Installs to `/usr/local/bin` on Linux and macOS. May prompt for `sudo` if the directory is not writable.

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

Sets up `~/.config/operatree/operatree.yaml` with your standard projects directory, preferred editor, and file manager.

**2. Bootstrap a new project**

```bash
operatree bootstrap myproject
```

Creates the full directory structure and registers the project in config. Prompts to set it as the default project.

**3. Set a default project**

```bash
operatree default           # pick from tracked projects interactively
operatree default --show    # show current default
```

Once set, all commands use it automatically — no `-d` flag needed.

**4. Create a subject**

```bash
operatree new event
operatree new task
operatree new topic
operatree new objective
```

Launches an interactive form to capture metadata. Every creation is logged to `activity.log`.

**5. Find anything**

```bash
operatree find                    # browse all subjects interactively
operatree find hany               # fuzzy search across all metadata
operatree find event hany         # filter by type, then fuzzy search
```

**6. Edit and open**

```bash
operatree metadata                # find a subject, open META.yaml in editor
operatree metadata event cairo    # filter first, then open in editor
operatree open                    # find a subject, open its directory in file manager
operatree open task report        # filter first, then open directory
```

After editing metadata manually, sync the project:

```bash
operatree sync
```

**7. Describe and summarize**

```bash
operatree desc                    # styled project structure
operatree desc --plain            # raw YAML for piping
operatree summary                 # counts, types, status breakdown
operatree explain                 # directory philosophy guide
```

**8. Track existing projects**

```bash
operatree track                   # add current directory project to tracked list
operatree untrack                 # remove current directory project from tracked list
```

---

## Commands

| Command                            | Description                                                              |
| ---------------------------------- | ------------------------------------------------------------------------ |
| `operatree init`                   | Initialize OperaTree configuration                                       |
| `operatree bootstrap [name]`       | Bootstrap a new project                                                  |
| `operatree new [type]`             | Create a new subject interactively                                       |
| `operatree find [type] [term]`     | Fuzzy-find subjects across all metadata                                  |
| `operatree metadata [type] [term]` | Find a subject and open its metadata in editor                           |
| `operatree open [type] [term]`     | Find a subject and open its directory in file manager                    |
| `operatree archive [type] [term]`  | Find a subject and archives its directory in 99_ARCHIVE module directory |
| `operatree sync`                   | Sync project metadata with subject files on disk                         |
| `operatree track`                  | Add current project to tracked list                                      |
| `operatree untrack`                | Remove current project from tracked list                                 |
| `operatree default`                | Set default project interactively                                        |
| `operatree desc`                   | Describe project structure                                               |
| `operatree summary`                | Project summary with counts and status                                   |
| `operatree explain`                | Print directory philosophy guide                                         |
| `operatree version`                | Print version, commit, and build date                                    |

### Flags

| Flag             | Command     | Description                                     |
| ---------------- | ----------- | ----------------------------------------------- |
| `-d, --dest`     | all         | Override project directory                      |
| `-t, --template` | `bootstrap` | define project template                         |
| `--plain`        | `desc`      | Output raw YAML instead of styled view          |
| `--show`         | `default`   | Show current default project                    |
| `--name`         | `new`       | Sets name for subject and skips interactive CLI |
| `--date`         | `new`       | Sets date for subject and skips interactive CLI |
| `-v, --verbose`  | `bootstrap` | Print project structure after creation          |

### Project Directory Resolution

When `-d` is not provided, OperaTree resolves the project directory in this order:

1. Current directory contains a `META.yaml` → use current directory
2. A default project is set in config → use it
3. Neither → error with a helpful message

---

## Subject Types

A subject is any trackable unit of work or knowledge within a project.

| Type        | Purpose                                       | Key Fields                         |
| ----------- | --------------------------------------------- | ---------------------------------- |
| `event`     | A project activity — visit, workshop, meeting | date, location, participants, tags |
| `task`      | A unit of work with a lifecycle               | owner, status, related events      |
| `topic`     | A knowledge concept or domain area            | tags, notes, related objective     |
| `objective` | A goal driving research and decisions         | status, findings, strategy         |

All subjects share common fields: `name`, `date`, `tags`, `notes`. Type-specific fields use `omitempty` so they never appear in unrelated subjects.

---

## Search

OperaTree builds a rich search index from all subject metadata — not just names. A single fuzzy query matches across the full depth of the project tree including nested modules:

- Name, tags, participants
- Notes, date, location

```bash
operatree find cairo              # matches events in Cairo, notes mentioning Cairo, etc.
operatree find event emissions    # filter by type, then search
operatree find                    # open interactive browser across all subjects
```

The interactive finder shows a tabulated list with module path breadcrumbs and a live preview panel for the selected subject. The same finder is used by `metadata` and `open`.

---

## Editing Metadata

Subject metadata files (`META.yaml`) are plain YAML and can be edited directly in any text editor. OperaTree embraces this — it never locks your data.

The `metadata` command finds a subject and opens its `META.yaml` in your configured editor:

```bash
operatree metadata                # interactive finder, then open in editor
operatree metadata event cairo    # filter first, then open
```

After manual edits, run `sync` to update the project metadata index:

```bash
operatree sync
```

Sync walks the full project tree, reads each `META.yaml` from disk, and updates the project metadata file accordingly.

---

## Activity Log

Every subject creation, edit, and archive is recorded in `activity.log` at the project root:

```
2026-05-20T10:08:39Z    CREATE    event    "Site Visit Cairo"    hany@optiplex7040    v0.1.0
2026-05-20T11:22:14Z    CREATE    task     "Prepare Report"      hany@optiplex7040    v0.1.0
2026-05-20T14:05:03Z    EDIT      task     "Prepare Report"      hany@optiplex7040    v0.1.2
```

Tab-separated columns: `timestamp`, `action`, `type`, `name`, `user@host`, `version`.

The log is append-only and pipe-friendly:

```bash
grep CREATE activity.log | cut -f3 | sort | uniq -c    # count creations by type
grep event activity.log                                  # all event actions
grep hany activity.log | tail -20                        # last 20 actions by user
```

Add `activity.log` to `.gitignore` to exclude it, or commit it for a full audit trail — both are valid.

---

## Configuration

Config lives at `~/.config/operatree/operatree.yaml`:

```yaml
standardDir: /home/user/projects
editor: nvim # fallback to $EDITOR if not set
fileManager: nautilus
default:
  name: myproject
  absPath: /home/user/projects/myproject
  template: dev
projects:
  - name: myproject
    absPath: /home/user/projects/myproject
    template: dev
  - name: research-2026
    absPath: /home/user/projects/research-2026
    template: research
daemon:
  enabled: false
  host: localhost
  port: 7070
  dbDriver: sqlite
  dsn: ""
```

| Field         | Description                                              |
| ------------- | -------------------------------------------------------- |
| `standardDir` | Default base directory for new projects                  |
| `editor`      | Editor for `operatree metadata`. Falls back to `$EDITOR` |
| `fileManager` | File manager for `operatree open`                        |
| `default`     | Default project used when `-d` is not provided           |
| `projects`    | All tracked projects                                     |
| `daemon`      | Future daemon configuration (reserved)                   |

---

## Version Control & Backup

OperaTree's filesystem-first design makes it naturally compatible with any version control or backup strategy. Because everything is plain files and directories, no special integration is needed — your existing tools already work.

### The two-layer model

**Layer 1 — Change detection:** a file watcher monitors the project directory and triggers an action when files are added, modified, or deleted. Tools that fit this role include:

- [`watchexec`](https://github.com/hanymamdouh82/watchexec) — a general-purpose directoris watcher with flexible trigger rules, delays, exclusion lists, etc.
- [`inotifywait`](https://github.com/inotify-tools/inotify-tools) — Linux-native filesystem event monitoring
- [`fswatch`](https://github.com/emcrisostomo/fswatch) — cross-platform file change monitor

**Layer 2 — The action:** what happens when a change is detected. Common strategies:

| Strategy              | Tool         | Best for                                     |
| --------------------- | ------------ | -------------------------------------------- |
| Local version history | `git commit` | Full history, diffs, branching               |
| Local + remote backup | `git push`   | Team sharing, offsite backup                 |
| Directory sync        | `rsync`      | Mirror to another machine or drive           |
| Continuous sync       | `syncthing`  | Real-time multi-device sync without a server |
| Cloud backup          | `rclone`     | S3, Google Drive, Dropbox, etc.              |

OperaTree does not ship a built-in watcher or backup engine — this is an intentional design decision. The right strategy depends on your environment, team size, and infrastructure. The filesystem layout is the contract; how you protect it is yours to decide.

Native integration with file watchers and version control backends is on the roadmap as a community contribution area. See [CONTRIBUTING.md](CONTRIBUTING.md) for details.

---

## Roadmap

OperaTree is in active development. The foundation is filesystem-first and stable. Future layers will be built on top without breaking it.

| Phase                       | Description                                                                     | Status     |
| --------------------------- | ------------------------------------------------------------------------------- | ---------- |
| CLI                         | Filesystem engine, YAML metadata, fuzzy search, interactive forms, activity log | 🚧 Alpha   |
| Version control integration | Native watchexec/git hooks, configurable watcher and action backends            | 📋 Planned |
| Index sidecar               | SQLite mirror for fast queries, no filesystem writes                            | 📋 Planned |
| Daemon                      | API over the index, sync engine, configuration-driven engine selection          | 📋 Planned |
| Semantic search             | Embeddings and vector search over subject metadata, and contents (with Daemon)  | 📋 Planned |
| Web platform                | Multi-user, web UI, enterprise features (commercial)                            | 💡 Vision  |

---

## Contributing

OperaTree welcomes contributions. The most impactful areas for community contribution are:

- **New subject types** — extend the type system with domain-specific subjects
- **New project templates** — add templates for research, legal, creative, and other domains
- **Search enhancements** — improve matching algorithms, ranking, relevance
- **Version control backends** — watcher strategies, git hooks, rsync/syncthing integration
- **Output formatters** — new ways to display and export project data

See [CONTRIBUTING.md](CONTRIBUTING.md) for the full guide.

See [ARCHITECTURE.md](ARCHITECTURE.md) for the full architecture guide.

---

## License

Copyright 2026 Hany Mamdouh

Licensed under the [Apache License, Version 2.0](LICENSE).

> Commercial modules (daemon, semantic search, web platform) will be released
> under a separate commercial license when available. The CLI tool will always
> remain Apache 2.0.
