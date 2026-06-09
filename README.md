# OperaTree

> Your project operating system — built on your filesystem.

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.26+-00ADD8.svg)](https://golang.org)
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
├── 97_MEDIA_LIBRARY/       # shared reusable assets
├── 98_DELIVERABLES/        # final external outputs
├── 99_ARCHIVE/             # historical storage
└── activity.log            # append-only audit trail
```

Each subject (event, task, topic, objective, datasource) lives in its own directory with a `META.yaml` file that makes it searchable, filterable, and machine-readable.

> **Note:** All directory and file names are case-sensitive, including on Windows. `01_EVENTS` and `01_events` are different paths. Always use the exact casing shown above.

---

## Features

- **Create projects** from opinionated templates in seconds
- **Add subjects** interactively or fully non-interactively via flags — events, tasks, topics, objectives, datasources
- **Fuzzy search** across all metadata fields: names, tags, participants, notes, dates, locations
- **Non-interactive search** via `--term` and `--type` flags for scripting pipelines
- **Edit subject metadata** directly in your preferred editor — metadata index synced automatically on close
- **Archive subjects** interactively or directly by UUID — moves to `99_ARCHIVE/`
- **Rename subjects** interactively or directly by UUID — updates directory, metadata, and all cross-references
- **Open subject directories** in your preferred file manager
- **Sync** project metadata with manually edited subject files
- **Track and untrack** existing projects without creating them
- **Describe and summarize** projects with colored terminal output
- **Default project** — set once, never pass `-d` again
- **Activity log** — every create, edit, and archive appended to `activity.log` at project root
- **Version control ready** — filesystem layout is designed to work naturally with Git, rsync, Syncthing, or any file watcher
- **Pipe-friendly output** — chain with `grep`, `sed`, `cut`, and other UNIX tools
- **Self-documenting** — `operatree explain` renders the full directory philosophy in your terminal

---

## Installation

### Linux / macOS

```bash
curl -fsSL https://raw.githubusercontent.com/hanymamdouh82/operatree/main/install.sh | sh
```

No tools required beyond `curl`. The script auto-detects your OS and architecture, downloads the right binary, and installs it to `/usr/local/bin`.

### Windows

Open PowerShell and run:

```powershell
irm https://raw.githubusercontent.com/hanymamdouh82/operatree/main/install.ps1 | iex
```

No tools required. Installs to `%USERPROFILE%\bin` and adds it to your PATH automatically. Restart your terminal after installation.

### Manual download

Download the binary for your platform directly from the [Releases page](https://github.com/hanymamdouh82/operatree/releases):

| File                          | Platform                       |
| ----------------------------- | ------------------------------ |
| `operatree-linux-amd64`       | Linux x86-64                   |
| `operatree-linux-arm64`       | Linux ARM (Raspberry Pi, etc.) |
| `operatree-darwin-amd64`      | macOS Intel                    |
| `operatree-darwin-arm64`      | macOS Apple Silicon            |
| `operatree-windows-amd64.exe` | Windows x86-64                 |

Rename the downloaded file to `operatree` (or `operatree.exe` on Windows) and place it in a directory on your `PATH`.

### From source

Requires Go 1.26 or higher.

```bash
git clone https://github.com/hanymamdouh82/operatree.git
cd operatree
make install
```

Installs to `/usr/local/bin` on Linux and macOS. May prompt for `sudo` if the directory is not writable.

---

## Quick Start

**1. Initialize configuration**

```bash
operatree init
```

Sets up the config file with your standard projects directory, preferred editor, and file manager. Run once before anything else.

**2. Create a new project**

```bash
operatree create myproject -t dev
```

Creates the full directory structure and registers the project in config. Use `operatree show templates` to list available templates.

**3. Set a default project**

```bash
operatree use               # pick from tracked projects interactively
operatree show default      # show current default
```

Once set, all commands use it automatically — no `-d` flag needed.

**4. Add a subject**

```bash
operatree add event
operatree add task
operatree add topic
operatree add objective
operatree add datasource
```

Launches an interactive form to capture metadata. Every creation is logged to `activity.log`. Use flags to skip the form entirely:

```bash
operatree add event --name "Site Visit" --date 2026-06-01 --location Cairo --participants "Alex,Sara"
operatree add task --name "Prepare Report" --owner Alex --status active --related-events "Site Visit"
```

**5. Find anything**

```bash
operatree find                         # browse all subjects interactively
operatree find cairo                   # fuzzy search across all metadata
operatree find event cairo             # filter by type, then fuzzy search
operatree find --term cairo --plain    # non-interactive, raw YAML output
```

**6. Edit and open**

```bash
operatree edit                    # find a subject, open META.yaml in editor
operatree edit event cairo        # filter first, then open in editor
operatree open                    # find a subject, open its directory in file manager
operatree open task report        # filter first, then open directory
```

The metadata index is updated automatically when the editor closes — no manual sync needed.

**7. Describe and summarize**

```bash
operatree describe                # styled project structure
operatree describe --plain        # raw YAML for piping
operatree summary                 # counts, types, status breakdown
operatree explain                 # directory philosophy guide
```

**8. Track existing projects**

```bash
operatree track -d /path/to/project    # add an existing project to tracked list
operatree untrack myproject            # remove by name
operatree untrack -d .                 # remove by path
```

---

## Commands

| Command                           | Description                                           |
| --------------------------------- | ----------------------------------------------------- |
| `operatree init`                  | Initialize OperaTree configuration                    |
| `operatree create [name]`         | Create a new project from a template                  |
| `operatree use`                   | Set default project interactively                     |
| `operatree goto`                  | Open a tracked project in the file manager            |
| `operatree add [type]`            | Add a new subject interactively or via flags          |
| `operatree find [type] [term]`    | Fuzzy-find subjects across all metadata               |
| `operatree edit [type] [term]`    | Find a subject and open its metadata in editor        |
| `operatree open [type] [term]`    | Find a subject and open its directory in file manager |
| `operatree rename [type] [term]`  | Find a subject and rename it                          |
| `operatree archive [type] [term]` | Find a subject and move it to `99_ARCHIVE/`           |
| `operatree sync`                  | Sync project metadata with subject files on disk      |
| `operatree track`                 | Add a project to tracked list                         |
| `operatree untrack [name]`        | Remove a project from tracked list                    |
| `operatree describe`              | Describe project directory structure                  |
| `operatree summary`               | Project summary with counts and status                |
| `operatree explain`               | Print directory philosophy guide                      |
| `operatree show [verb]`           | Show OperaTree configuration and state                |
| `operatree version`               | Print version, commit, and build date                 |

### Flags

| Flag                  | Short | Command                     | Description                                  |
| --------------------- | ----- | --------------------------- | -------------------------------------------- |
| `--dest`              | `-d`  | all                         | Project directory to operate on              |
| `--template`          | `-t`  | `create`                    | Project template (required)                  |
| `--verbose`           | `-v`  | `create`                    | Print directory structure after creation     |
| `--plain`             | `-p`  | `describe`, `find`          | Output raw YAML instead of styled view       |
| `--name`              | —     | `add`                       | Subject name — triggers non-interactive mode |
| `--date`              | —     | `add`                       | Subject date                                 |
| `--notes`             | —     | `add`                       | Subject notes                                |
| `--tags`              | —     | `add`                       | Comma-delimited tags                         |
| `--location`          | —     | `add event`                 | Event location                               |
| `--participants`      | —     | `add event`                 | Comma-delimited participants                 |
| `--owner`             | —     | `add task`                  | Task owner                                   |
| `--status`            | —     | `add task`, `add objective` | Subject status                               |
| `--related-events`    | —     | `add task`                  | Comma-delimited related event names          |
| `--related-objective` | —     | `add topic`                 | Related objective name                       |
| `--outputs`           | —     | `add task`, `add objective` | Comma-delimited outputs                      |
| `--source`            | —     | `add datasource`            | Data origin                                  |
| `--source-link`       | —     | `add datasource`            | URL or path to source data                   |
| `--source-objective`  | —     | `add datasource`            | Related objective                            |
| `--source-datasize`   | —     | `add datasource`            | Dataset size                                 |
| `--term`              | `-t`  | `find`                      | Search term for non-interactive mode         |
| `--type`              | `-s`  | `find`                      | Subject type filter for non-interactive mode |
| `--uuid`              | `-u`  | `rename`, `archive`         | Subject UUID for non-interactive mode        |
| `--new-name`          | `-n`  | `rename`                    | New subject name (required with `--uuid`)    |

### Project Directory Resolution

When `-d` is not provided, OperaTree resolves the project directory in this order:

1. Current directory contains a `META.yaml` → use current directory
2. A default project is set in config → use it
3. Neither → error with a helpful message

---

## Subject Types

A subject is any trackable unit of work or knowledge within a project.

| Type         | Purpose                                       | Key Fields                                          |
| ------------ | --------------------------------------------- | --------------------------------------------------- |
| `event`      | A project activity — visit, workshop, meeting | date, location, participants, tags                  |
| `task`       | A unit of work with a lifecycle               | owner, status, related events, outputs              |
| `topic`      | A knowledge concept or domain area            | tags, notes, related objective                      |
| `objective`  | A goal driving research and decisions         | status, outputs, tags                               |
| `datasource` | An external dataset or data feed              | source, sourceLink, sourceObjective, sourceDataSize |

All subjects share common fields: `name`, `date`, `tags`, `notes`. Type-specific fields use `omitempty` so they never appear in unrelated subjects.

---

## Search

OperaTree builds a rich search index from all subject metadata — not just names. A single fuzzy query matches across the full depth of the project tree including nested modules:

- Name, tags, participants
- Notes, date, location

**Interactive mode** — launches the finder with a live preview panel:

```bash
operatree find cairo              # matches events in Cairo, notes mentioning Cairo, etc.
operatree find event emissions    # filter by type, then search
operatree find                    # open interactive browser across all subjects
```

**Non-interactive mode** — returns results directly for scripting:

```bash
operatree find --term cairo --type event          # filter events for "cairo"
operatree find --term done --type task --plain    # output as raw YAML

# Bulk archive all done tasks via pipeline
operatree find --term done --type task --plain \
  | grep uuid \
  | awk '{print $2}' \
  | xargs -I{} operatree archive --uuid {}
```

The interactive finder shows a tabulated list with module path breadcrumbs and a live preview panel for the selected subject. The same finder is used by `edit` and `open`.

---

## Editing Metadata

Subject metadata files (`META.yaml`) are plain YAML and can be edited directly in any text editor. OperaTree embraces this — it never locks your data.

The `edit` command finds a subject and opens its `META.yaml` in your configured editor:

```bash
operatree edit                    # interactive finder, then open in editor
operatree edit event cairo        # filter first, then open
```

The project metadata index is updated automatically when the editor closes — no manual sync needed.

For edits made outside of OperaTree (direct file edits, git pulls, Syncthing syncs), run `sync` manually to reconcile:

```bash
operatree sync
```

---

## Activity Log

Every subject creation, edit, and archive is recorded in `activity.log` at the project root:

```
2026-05-20T10:08:39Z    CREATE    event    "Site Visit Cairo"    hany@optiplex7040    v0.1.0
2026-05-20T11:22:14Z    CREATE    task     "Prepare Report"      hany@optiplex7040    v0.1.0
2026-05-20T14:05:03Z    EDIT      task     "Prepare Report"      hany@optiplex7040    v0.1.2
2026-05-20T16:45:00Z    ARCHIVE   task     "Old Vendor Analysis" hany@optiplex7040    v0.1.2
```

Tab-separated columns: `timestamp`, `action`, `type`, `name`, `user@host`, `version`.

The log is append-only and pipe-friendly:

```bash
grep CREATE activity.log | cut -f3 | sort | uniq -c    # count creations by type
grep event activity.log                                  # all event actions
grep hany activity.log | tail -20                        # last 20 actions by user
grep ARCHIVE activity.log | cut -f4                      # names of archived subjects
```

Add `activity.log` to `.gitignore` to exclude it, or commit it for a full audit trail — both are valid.

---

## Configuration

Config lives at `~/.config/operatree/operatree.yaml` (see platform-specific paths below):

```yaml
standardDir: /home/user/projects
editor: nvim # fallback to $EDITOR if not set
fileManager: nautilus # no fallback if not set
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

| Field         | Description                                                         |
| ------------- | ------------------------------------------------------------------- |
| `standardDir` | Default base directory for new projects                             |
| `editor`      | Editor for `operatree edit`. Falls back to `$EDITOR`                |
| `fileManager` | File manager for `operatree open` and `operatree goto`. No fallback |
| `default`     | Default project used when `-d` is not provided                      |
| `projects`    | All tracked projects                                                |
| `daemon`      | Reserved for future daemon configuration                            |

**Config file locations:**

| Platform    | Location                                                 |
| ----------- | -------------------------------------------------------- |
| Linux       | `~/.config/operatree/operatree.yaml`                     |
| Linux (XDG) | `$XDG_CONFIG_HOME/operatree/operatree.yaml`              |
| macOS       | `~/Library/Application Support/operatree/operatree.yaml` |
| Windows     | `%APPDATA%\operatree\operatree.yaml`                     |

Run `operatree init` again at any time to reconfigure — the prompt pre-fills your current values.

---

## Version Control & Backup

OperaTree's filesystem-first design makes it naturally compatible with any version control or backup strategy. Because everything is plain files and directories, no special integration is needed — your existing tools already work.

### The two-layer model

**Layer 1 — Change detection:** a file watcher monitors the project directory and triggers an action when files are added, modified, or deleted. Tools that fit this role include:

- [`watchexec`](https://github.com/hanymamdouh82/watchexec) — a config-driven directory watcher with flexible trigger rules, delays, and exclusion lists. Can run as a systemd service.
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

### Example: watchexec + operatree sync

```yaml
# conf.yml — auto-sync metadata index on any project change
dirs:
  /home/user/projects/myproject:
    bin: operatree
    args:
      - sync
      - -d
      - /home/user/projects/myproject
    stdout: true
    delay: 5
    exclude:
      - .git
      - activity.log
```

```bash
watchexec -c /home/user/.config/watchexec/conf.yml
```

OperaTree does not ship a built-in watcher or backup engine — this is an intentional design decision. The right strategy depends on your environment, team size, and infrastructure. The filesystem layout is the contract; how you protect it is yours to decide.

---

## Roadmap

OperaTree is in active development. The foundation is filesystem-first and stable. Future layers will be built on top without breaking it.

| Phase                       | Description                                                                                         | Status         |
| --------------------------- | --------------------------------------------------------------------------------------------------- | -------------- |
| CLI                         | Filesystem engine, YAML metadata, fuzzy search, interactive and non-interactive forms, activity log | 🚧 Alpha       |
| Version control integration | Native watchexec/git hooks, configurable watcher and action backends                                | 📋 Planned     |
| Index sidecar               | SQLite mirror for fast queries, no filesystem writes                                                | 📋 Planned     |
| Daemon                      | API over the index, sync engine, configuration-driven engine selection                              | 📋 Planned     |
| Semantic search             | Embeddings and vector search over subject metadata and file contents                                | 📋 In Progress |
| Web platform                | Multi-user, web UI, enterprise features (commercial)                                                | 💡 Vision      |

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
