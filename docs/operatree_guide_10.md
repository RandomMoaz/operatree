# Section 10 — Command Reference

---

This section is a complete reference for all OperaTree commands and flags.
For conceptual explanations and workflow guidance, refer to the relevant
section of this guide. This section is designed for quick lookup.

---

## 10.1 Global Flags

These flags are available on all commands that operate on a project directory.

| Flag     | Short | Type   | Default                        | Description                                                                                                                                                                                                                    |
| -------- | ----- | ------ | ------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `--dest` | `-d`  | string | default project or current dir | Project directory to operate on. Accepts `.` for current directory, or a relative or absolute path. If unset, resolves in this order: (1) current directory contains a `META.yaml`, (2) default project from config, (3) error |

---

## 10.2 Command Index

| Command                           | Description                                  |
| --------------------------------- | -------------------------------------------- |
| [`operatree init`](#init)         | Initialize OperaTree configuration           |
| [`operatree create`](#create)     | Create a new project                         |
| [`operatree use`](#use)           | Set the default project                      |
| [`operatree goto`](#goto)         | Open a tracked project in the file manager   |
| [`operatree track`](#track)       | Register a project directory                 |
| [`operatree untrack`](#untrack)   | Remove a project from the tracked list       |
| [`operatree add`](#add)           | Add a new subject to the project             |
| [`operatree find`](#find)         | Fuzzy-find subjects across all metadata      |
| [`operatree edit`](#edit)         | Edit subject metadata in your editor         |
| [`operatree open`](#open)         | Open a subject directory in the file manager |
| [`operatree rename`](#rename)     | Rename a subject                             |
| [`operatree archive`](#archive)   | Archive a subject                            |
| [`operatree sync`](#sync)         | Sync the project metadata index from disk    |
| [`operatree describe`](#describe) | Print the project directory structure        |
| [`operatree summary`](#summary)   | Print a project content overview             |
| [`operatree explain`](#explain)   | Print the directory philosophy guide         |
| [`operatree show`](#show)         | Display OperaTree configuration and state    |
| [`operatree version`](#version)   | Print version information                    |

---

## 10.3 Command Details

---

### `init`

Create the OperaTree configuration file in the system config directory.
Interactively prompts for standard directory, editor, and file manager.
Run this once before using any other command. Running it again on an
existing installation opens the prompt with current values pre-filled
for reconfiguration.

```bash
operatree init
```

**Flags:** none

**Config file locations:**

| Platform    | Location                                                 |
| ----------- | -------------------------------------------------------- |
| Linux       | `~/.config/operatree/operatree.yaml`                     |
| Linux (XDG) | `$XDG_CONFIG_HOME/operatree/operatree.yaml`              |
| macOS       | `~/Library/Application Support/operatree/operatree.yaml` |
| Windows     | `%APPDATA%\operatree\operatree.yaml`                     |

---

### `create`

Scaffold a new project with the full OperaTree directory structure and
register it in config. A template is required and determines which modules
are included. Use `operatree show templates` to list available templates.

```bash
operatree create [project_name] -t [template]
```

**Arguments:**

| Argument       | Required | Description                             |
| -------------- | -------- | --------------------------------------- |
| `project_name` | Yes      | Name of the project directory to create |

**Flags:**

| Flag         | Short | Required | Default                   | Description                                      |
| ------------ | ----- | -------- | ------------------------- | ------------------------------------------------ |
| `--template` | `-t`  | Yes      | —                         | Project template to use                          |
| `--dest`     | `-d`  | No       | `standardDir` from config | Base directory for the new project               |
| `--verbose`  | `-v`  | No       | false                     | Print created directory structure after creation |

**Examples:**

```bash
operatree create fleetfix -t consulting
operatree create research-2026 -t research -v
operatree create myproject -t dev -d /home/alex/projects
```

---

### `use`

Interactively select a default project from your tracked projects.
Once set, all commands use it automatically without requiring `-d`.

```bash
operatree use
```

**Flags:** none

**Related:**

```bash
operatree show default     # view current default
```

---

### `goto`

Interactively select from tracked projects and open the chosen project
root directory in your configured file manager. Opens a file manager
window — does not change the terminal working directory.

```bash
operatree goto
```

**Flags:** none

---

### `track`

Register a project directory in your OperaTree configuration. The `-d`
flag is required — there is no interactive picker. Tracked projects are
available for `operatree use`, `operatree goto`, and all commands that
resolve the project directory automatically.

```bash
operatree track -d [path]
```

**Flags:**

| Flag     | Short | Required | Default | Description                   |
| -------- | ----- | -------- | ------- | ----------------------------- |
| `--dest` | `-d`  | Yes      | —       | Project directory to register |

**Examples:**

```bash
operatree track -d /home/alex/projects/fleetfix
operatree track -d .
```

---

### `untrack`

Remove a project from the tracked projects list. The project directory
and its contents are not affected. If the untracked project was the
default, the default is cleared.

```bash
operatree untrack [project_name]
operatree untrack -d [path]
```

**Arguments:**

| Argument       | Required | Description                                 |
| -------------- | -------- | ------------------------------------------- |
| `project_name` | No       | Name of the project as registered in config |

**Flags:**

| Flag     | Short | Required | Default         | Description                  |
| -------- | ----- | -------- | --------------- | ---------------------------- |
| `--dest` | `-d`  | No       | default project | Project directory to untrack |

**Resolution order:** project name argument → `-d` flag

**Examples:**

```bash
operatree untrack fleetfix
operatree untrack -d /home/alex/projects/fleetfix
operatree untrack -d .
```

---

### `add`

Launch an interactive form to add a new subject to the project, or create
one directly using flags for scripting. The subject type is required.
Providing `--name` skips the interactive form entirely and creates the
subject immediately with the supplied flag values.

```bash
operatree add [type]
```

**Valid types:** `event`, `task`, `topic`, `objective`, `datasource`

**Common flags (all subject types):**

| Flag      | Short | Required | Default         | Description                                         |
| --------- | ----- | -------- | --------------- | --------------------------------------------------- |
| `--name`  | —     | No       | —               | Subject name — triggers non-interactive mode        |
| `--date`  | —     | No       | —               | Subject date                                        |
| `--notes` | —     | No       | —               | Subject notes                                       |
| `--tags`  | —     | No       | —               | Comma-delimited tags (e.g. `"cairo,inspection,q2"`) |
| `--dest`  | `-d`  | No       | default project | Project directory to operate on                     |

**Type-specific flags:**

| Flag                  | Type            | Description                                        |
| --------------------- | --------------- | -------------------------------------------------- |
| `--location`          | event           | Location where the event took place                |
| `--participants`      | event           | Comma-delimited participant names                  |
| `--owner`             | task            | Person responsible for the task                    |
| `--status`            | task, objective | Current status (e.g. `active`, `done`, `deferred`) |
| `--related-events`    | task            | Comma-delimited related event names                |
| `--outputs`           | task, objective | Comma-delimited expected or produced outputs       |
| `--related-objective` | topic           | Name of the objective this topic supports          |
| `--source`            | datasource      | Data origin (e.g. Kaggle, internal team, API)      |
| `--source-link`       | datasource      | URL or path to the source data                     |
| `--source-objective`  | datasource      | Objective this data source supports                |
| `--source-datasize`   | datasource      | Size or volume of the dataset                      |

**Subject fields by type:**

| Field                 | event | task | topic | objective | datasource |
| --------------------- | ----- | ---- | ----- | --------- | ---------- |
| `--name`              | ✓     | ✓    | ✓     | ✓         | ✓          |
| `--date`              | ✓     | ✓    | ✓     | ✓         | ✓          |
| `--tags`              | ✓     | ✓    | ✓     | ✓         | ✓          |
| `--notes`             | ✓     | ✓    | ✓     | ✓         | ✓          |
| `--location`          | ✓     |      |       |           |            |
| `--participants`      | ✓     |      |       |           |            |
| `--owner`             |       | ✓    |       |           |            |
| `--status`            |       | ✓    |       | ✓         |            |
| `--related-events`    |       | ✓    |       |           |            |
| `--related-objective` |       |      | ✓     |           |            |
| `--outputs`           |       | ✓    |       | ✓         |            |
| `--source`            |       |      |       |           | ✓          |
| `--source-link`       |       |      |       |           | ✓          |
| `--source-objective`  |       |      |       |           | ✓          |
| `--source-datasize`   |       |      |       |           | ✓          |

**Subject module locations:**

| Type         | Module                            |
| ------------ | --------------------------------- |
| `event`      | `01_EVENTS/`                      |
| `task`       | `02_PROJECT_MANAGEMENT/07_TASKS/` |
| `topic`      | `04_RESEARCH/09_TOPICS/`          |
| `objective`  | `04_RESEARCH/10_OBJECTIVES/`      |
| `datasource` | `06_DATA/15_DATASOURCES/`         |

**Examples:**

```bash
operatree add event
operatree add event --name "Cairo Site Visit" --date 2026-06-01 --location Cairo --participants "Alex,Sara" --tags "site,inspection"
operatree add task --name "Prepare Report" --owner Alex --status active --related-events "Cairo Site Visit"
operatree add topic --name "Predictive Maintenance" --related-objective "Reduce Downtime" --tags "ml,iot"
operatree add objective --name "Reduce Downtime" --status active --tags "maintenance,kpi"
operatree add datasource --name "Sensor Readings 2025" --source "IoT Team" --source-link "/06_DATA/01_RAW/sensors.csv" --source-datasize "2.4GB"
operatree add task --name "Deploy" --date 2026-06-15 -d /path/to/project
```

---

### `find`

Fuzzy-find subjects across all metadata fields — name, tags, participants,
notes, date, and location. Supports interactive and non-interactive modes.
Read-only — never modifies anything.

**Interactive mode** — opens the finder with a live preview panel.
Selecting a subject displays its full formatted metadata:

```bash
operatree find [type] [term]
```

**Non-interactive mode** — returns results directly without launching the
finder. Use `--term` and `--type` flags:

```bash
operatree find --term [term] --type [type]
```

**Arguments (interactive mode):**

| Argument | Required | Description                                    |
| -------- | -------- | ---------------------------------------------- |
| `type`   | No       | Filter by subject type before launching finder |
| `term`   | No       | Fuzzy search term across all metadata fields   |

**Flags:**

| Flag      | Short | Required | Default         | Description                                          |
| --------- | ----- | -------- | --------------- | ---------------------------------------------------- |
| `--term`  | `-t`  | No       | —               | Search term for non-interactive mode                 |
| `--type`  | `-s`  | No       | —               | Subject type filter for non-interactive mode         |
| `--plain` | `-p`  | No       | false           | Output results as raw YAML instead of formatted view |
| `--dest`  | `-d`  | No       | default project | Project directory to operate on                      |

**Examples:**

```bash
operatree find
operatree find event
operatree find event cairo
operatree find cairo
operatree find --term cairo
operatree find --term report --type task
operatree find --term report --type task --plain
operatree find --term report --type task --plain | grep owner
operatree find --term done --type task --plain | grep uuid | awk '{print $2}' | xargs -I{} operatree archive --uuid {}
```

---

### `edit`

Fuzzy-find a subject and open its `META.yaml` in your configured editor.
The project metadata index is updated automatically when the editor closes.

```bash
operatree edit [type] [term]
```

**Arguments:**

| Argument | Required | Description                                    |
| -------- | -------- | ---------------------------------------------- |
| `type`   | No       | Filter by subject type before launching finder |
| `term`   | No       | Fuzzy search term across all metadata fields   |

**Flags:**

| Flag     | Short | Required | Default         | Description                     |
| -------- | ----- | -------- | --------------- | ------------------------------- |
| `--dest` | `-d`  | No       | default project | Project directory to operate on |

**Editor resolution order:** `editor` in config → `$EDITOR` environment variable → error

**Examples:**

```bash
operatree edit
operatree edit task
operatree edit task report
operatree edit event cairo -d /home/alex/projects/fleetfix
```

---

### `open`

Fuzzy-find a subject and open its directory in your configured file manager.

```bash
operatree open [type] [term]
```

**Arguments:**

| Argument | Required | Description                                    |
| -------- | -------- | ---------------------------------------------- |
| `type`   | No       | Filter by subject type before launching finder |
| `term`   | No       | Fuzzy search term across all metadata fields   |

**Flags:**

| Flag     | Short | Required | Default         | Description                     |
| -------- | ----- | -------- | --------------- | ------------------------------- |
| `--dest` | `-d`  | No       | default project | Project directory to operate on |

**Examples:**

```bash
operatree open
operatree open task
operatree open task report
operatree open event kickoff -d /home/alex/projects/fleetfix
```

---

### `rename`

Fuzzy-find a subject and rename it interactively, or target one directly
by UUID for scripting. Updates the subject directory name, `META.yaml`,
and all cross-references in the project metadata index in one operation.

**Interactive mode:**

```bash
operatree rename [type] [term]
```

**Non-interactive mode:**

```bash
operatree rename --uuid [uuid] --new-name [name]
```

**Arguments (interactive mode):**

| Argument | Required | Description                                    |
| -------- | -------- | ---------------------------------------------- |
| `type`   | No       | Filter by subject type before launching finder |
| `term`   | No       | Fuzzy search term across all metadata fields   |

**Flags:**

| Flag         | Short | Required | Default         | Description                                             |
| ------------ | ----- | -------- | --------------- | ------------------------------------------------------- |
| `--uuid`     | `-u`  | No       | —               | Subject UUID for non-interactive mode                   |
| `--new-name` | `-n`  | No       | —               | New name to assign — required when `--uuid` is provided |
| `--dest`     | `-d`  | No       | default project | Project directory to operate on                         |

**Note:** `--uuid` requires `--new-name`. Providing `--uuid` without `--new-name` is an error.

**Examples:**

```bash
operatree rename
operatree rename event
operatree rename event kickoff
operatree rename --uuid a1b2c3d4 --new-name "Cairo Factory Review"
operatree rename --uuid a1b2c3d4 --new-name "Cairo Factory Review" -d /path/to/project
```

---

### `archive`

Fuzzy-find a subject and move it to `99_ARCHIVE/` at the project root,
or target one directly by UUID for scripting. The subject directory and
`META.yaml` are preserved exactly as-is. The action is recorded in
`activity.log`.

**Interactive mode:**

```bash
operatree archive [type] [term]
```

**Non-interactive mode:**

```bash
operatree archive --uuid [uuid]
```

**Arguments (interactive mode):**

| Argument | Required | Description                                    |
| -------- | -------- | ---------------------------------------------- |
| `type`   | No       | Filter by subject type before launching finder |
| `term`   | No       | Fuzzy search term across all metadata fields   |

**Flags:**

| Flag     | Short | Required | Default         | Description                                                          |
| -------- | ----- | -------- | --------------- | -------------------------------------------------------------------- |
| `--uuid` | `-u`  | No       | —               | Subject UUID for non-interactive mode — bypasses the finder entirely |
| `--dest` | `-d`  | No       | default project | Project directory to operate on                                      |

**Examples:**

```bash
operatree archive
operatree archive task
operatree archive task report
operatree archive --uuid a1b2c3d4
operatree archive --uuid a1b2c3d4 -d /path/to/project

# Bulk archive all done tasks via pipeline
operatree find --term done --type task --plain \
  | grep uuid \
  | awk '{print $2}' \
  | xargs -I{} operatree archive --uuid {}
```

---

### `sync`

Walk the full project tree, re-read every `META.yaml` from disk, and
update the project metadata index. Run after editing subject files
manually outside of OperaTree. Safe to run at any time — reads from
disk, never overwrites files.

```bash
operatree sync
```

**Flags:**

| Flag     | Short | Required | Default         | Description                     |
| -------- | ----- | -------- | --------------- | ------------------------------- |
| `--dest` | `-d`  | No       | default project | Project directory to operate on |

**Examples:**

```bash
operatree sync
operatree sync -d /home/alex/projects/fleetfix
```

---

### `describe`

Print a structured view of the project directory tree as defined by its
template. Shows modules and subdirectories — not individual subjects.
Use `--plain` for raw YAML output suitable for piping.

```bash
operatree describe
```

**Flags:**

| Flag      | Short | Required | Default         | Description                            |
| --------- | ----- | -------- | --------------- | -------------------------------------- |
| `--plain` | `-p`  | No       | false           | Output raw YAML instead of styled view |
| `--dest`  | `-d`  | No       | default project | Project directory to operate on        |

**Examples:**

```bash
operatree describe
operatree describe --plain
operatree describe --plain | grep tags
operatree describe --plain > snapshot.yaml
operatree describe -d /home/alex/projects/anchor
```

---

### `summary`

Print a high-level overview of the project — subject counts by type,
status breakdown, and module distribution.

```bash
operatree summary
```

**Flags:**

| Flag     | Short | Required | Default         | Description                     |
| -------- | ----- | -------- | --------------- | ------------------------------- |
| `--dest` | `-d`  | No       | default project | Project directory to operate on |

**Examples:**

```bash
operatree summary
operatree summary -d /home/alex/projects/anchor
```

---

### `explain`

Print the full OperaTree directory philosophy guide — what each module
is for, what belongs in it, and how the layers relate to each other.

```bash
operatree explain
```

**Flags:** none

**Examples:**

```bash
operatree explain
operatree explain | less
```

---

### `show`

Display information about OperaTree configuration and state.

```bash
operatree show [verb]
```

**Valid verbs:**

| Verb        | Description                            |
| ----------- | -------------------------------------- |
| `tracked`   | List all tracked projects              |
| `config`    | Display current configuration          |
| `templates` | List available project templates       |
| `default`   | Show the currently set default project |

**Examples:**

```bash
operatree show tracked
operatree show config
operatree show templates
operatree show default
```

---

### `version`

Print the current OperaTree version, commit hash, and build date.

```bash
operatree version
```

**Flags:** none

**Example output:**

```
OperaTree v0.1.2
  Commit:     a3f8c21
  Built:      2026-05-20T10:00:00Z
```

---

## 10.4 Activity Log Reference

Every subject action is recorded in `activity.log` at the project root.
The log is append-only, tab-separated, and never modified by OperaTree.

**Format:**

```
timestamp    action    type    name    user@host    version
```

**Current actions:**

| Action    | Triggered by                  |
| --------- | ----------------------------- |
| `CREATE`  | `operatree add`               |
| `EDIT`    | `operatree edit`              |
| `ARCHIVE` | `operatree archive`           |
| `DELETE`  | Planned — not yet implemented |

**Common queries:**

```bash
cat activity.log                                        # full log
grep CREATE activity.log                                # all creations
grep CREATE activity.log | cut -f3 | sort | uniq -c    # count by type
grep task activity.log                                  # all task actions
grep "^2026-06" activity.log                           # all actions in June 2026
grep alex activity.log | tail -20                       # last 20 actions by user
grep ARCHIVE activity.log | cut -f4                     # names of archived subjects
```

---

## 10.5 Subject Directory Structure Reference

### Event

```
01_EVENTS/
└── [event-name]/
    ├── 01_AGENDA/
    ├── 02_MEDIA/
    ├── 03_NOTES/
    ├── 04_DOCUMENTS/
    ├── 05_OUTCOMES/
    └── META.yaml
```

### Task

```
02_PROJECT_MANAGEMENT/
└── 07_TASKS/
    └── [task-name]/
        ├── 01_INPUTS/
        ├── 02_WORKING/
        ├── 03_REVIEW/
        ├── 04_FINAL/
        └── META.yaml
```

### Topic

```
04_RESEARCH/
└── 09_TOPICS/
    └── [topic-name]/
        └── META.yaml
```

### Objective

```
04_RESEARCH/
└── 10_OBJECTIVES/
    └── [objective-name]/
        └── META.yaml
```

### Data Source

```
06_DATA/
└── 15_DATASOURCES/
    └── [datasource-name]/
        └── META.yaml
```

---

## 10.6 Template Reference

| Template     | Description                      | Includes                                                                                                    |
| ------------ | -------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `general`    | Minimal template for general use | Admin, Events, Project Management, Media Library, Deliverables, Archive                                     |
| `dev`        | Software development projects    | Admin, Events, Project Management, Legal, Research, Engineering, Data, Media Library, Deliverables, Archive |
| `research`   | Academic and R&D projects        | Admin, Events, Project Management, Legal, Research, Publications, Deliverables, Archive                     |
| `consulting` | Client engagement work           | Admin, Events, Project Management, Legal, Research, Deliverables, Archive                                   |

---

## 10.7 Module Prefix Reference

| Prefix | Module             | Notes                                           |
| ------ | ------------------ | ----------------------------------------------- |
| `00`   | Admin              | All templates                                   |
| `01`   | Events             | All templates                                   |
| `02`   | Project Management | All templates                                   |
| `03`   | Legal              | `dev`, `research`, `consulting`                 |
| `04`   | Research           | `dev`, `research`, `consulting`                 |
| `05`   | Engineering        | `dev` only                                      |
| `06`   | Data               | `dev` only                                      |
| `07`   | Tasks              | Nested under `02` — all templates               |
| `08`   | Index              | Nested under `04` — all templates with Research |
| `09`   | Topics             | Nested under `04` — all templates with Research |
| `10`   | Objectives         | Nested under `04` — all templates with Research |
| `11`   | Summaries          | Nested under `04` — all templates with Research |
| `12`   | References         | Nested under `04` — all templates with Research |
| `13`   | Audio Notes        | Nested under `04` — all templates with Research |
| `14`   | Attachments        | Nested under `04` — all templates with Research |
| `15`   | Data Sources       | Nested under `06` — `dev` only                  |
| `16`   | Publications       | `research` only                                 |
| `97`   | Media Library      | `dev`, `general`                                |
| `98`   | Deliverables       | All templates                                   |
| `99`   | Archive            | All templates                                   |
