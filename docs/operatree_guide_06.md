# Section 6 — Working with Subjects

---

This section covers everything related to subjects — adding, finding, editing, opening, renaming, and archiving. We continue with Alex and Sara on the `fleetfix` engagement, now several weeks in with a growing collection of events, tasks, and topics.

---

## 6.1 Adding a Subject

Subjects are the atomic unit of work in OperaTree. Every meeting, task, research topic, and goal is a subject. Adding one is always the same command pattern:

```bash
operatree add [type]
```

Where `[type]` is one of: `event`, `task`, `topic`, `objective`, `datasource`.

Every `operatree add` command supports two modes:

- **Interactive** — launches a form that guides you through each field
- **Non-interactive** — provide `--name` and any other flags to create the subject directly without prompts, useful for scripting and automation

### Adding an Event

Alex just finished a client workshop. To record it interactively:

```bash
operatree add event
```

OperaTree launches an interactive form collecting:

- **Name** — the event name
- **Date** — when it occurred
- **Location** — where it took place
- **Participants** — who was present
- **Tags** — searchable labels
- **Notes** — free-form notes

OperaTree creates the event directory inside `01_EVENTS/` with its standard subdirectories and writes the `META.yaml`. The creation is appended to `activity.log`.

To create an event non-interactively — provide `--name` and any other fields as flags:

```bash
operatree add event \
  --name "Client Workshop" \
  --date 2026-06-10 \
  --location Cairo \
  --participants "Alex,Sara,Omar" \
  --tags "workshop,client,kickoff" \
  --notes "Discussed Q3 deliverables and timeline."
```

### Adding a Task

The workshop generated a deliverable. Alex creates a task interactively:

```bash
operatree add task
```

The form collects:

- **Name** — the task name
- **Date** — date created or started
- **Owner** — person responsible
- **Status** — current status (e.g. `active`, `blocked`, `done`)
- **Related events** — events that generated or relate to this task
- **Outputs** — expected or produced outputs
- **Tags** — searchable labels
- **Notes** — free-form notes

OperaTree creates the task directory inside `07_TASKS/` with four stage subdirectories automatically:

```
prepare-workshop-report/
├── 01_INPUTS/
├── 02_WORKING/
├── 03_REVIEW/
├── 04_FINAL/
└── META.yaml
```

These stage directories are created by OperaTree and reflect the task lifecycle — from initial inputs through working drafts, review, and final output. You are free to organise your own files inside any stage directory, but the stage directories themselves should not be renamed or removed.

To create a task non-interactively:

```bash
operatree add task \
  --name "Prepare Workshop Report" \
  --date 2026-06-10 \
  --owner Alex \
  --status active \
  --related-events "Client Workshop" \
  --outputs "Workshop Report v1.0" \
  --tags "report,workshop"
```

### Adding a Topic

To support the task, Alex needs to research a domain concept:

```bash
operatree add topic
```

The form collects:

- **Name** — the topic name
- **Date** — date created
- **Related objective** — the objective this topic supports
- **Tags** — searchable labels
- **Notes** — free-form notes

OperaTree creates the topic directory inside `09_TOPICS/` nested under `04_RESEARCH/`.

To create a topic non-interactively:

```bash
operatree add topic \
  --name "Predictive Maintenance" \
  --date 2026-06-10 \
  --related-objective "Reduce Equipment Downtime" \
  --tags "ml,iot,maintenance"
```

### Adding an Objective

Alex sets a project goal:

```bash
operatree add objective
```

The form collects:

- **Name** — the objective name
- **Date** — date created
- **Status** — current status (e.g. `active`, `achieved`, `deferred`)
- **Outputs** — decisions, strategies, or deliverables produced
- **Tags** — searchable labels
- **Notes** — free-form notes

OperaTree creates the objective directory inside `10_OBJECTIVES/` nested under `04_RESEARCH/`.

To create an objective non-interactively:

```bash
operatree add objective \
  --name "Reduce Equipment Downtime" \
  --date 2026-06-01 \
  --status active \
  --tags "maintenance,kpi,strategy"
```

### Adding a Data Source

For projects that involve data, Alex registers an external dataset:

```bash
operatree add datasource
```

The form collects:

- **Name** — dataset name
- **Date** — date acquired or registered
- **Source** — origin (e.g. Kaggle, internal team, API)
- **Source link** — URL or path to the original data
- **Source objective** — the objective this data supports
- **Source data size** — size or volume of the dataset
- **Tags** — searchable labels
- **Notes** — free-form notes

OperaTree creates the data source record inside `15_DATASOURCES/` nested under `06_DATA/`. The actual data files belong in `06_DATA/01_RAW/` — the data source subject is the metadata record that traces where the data came from and what it supports.

To create a data source non-interactively:

```bash
operatree add datasource \
  --name "Sensor Readings 2025" \
  --date 2026-06-01 \
  --source "IoT Team" \
  --source-link "/06_DATA/01_RAW/sensors_2025.csv" \
  --source-objective "Reduce Equipment Downtime" \
  --source-datasize "2.4GB" \
  --tags "sensors,iot,raw"
```

---

## 6.2 Finding Subjects

After several weeks of work, `fleetfix` has dozens of subjects. Finding the right one quickly is where OperaTree's search earns its place.

`operatree find` supports two modes:

### Interactive mode

```bash
operatree find
```

Opens an interactive finder showing all subjects across the full project tree. The finder displays a tabulated list with module path breadcrumbs and a live preview panel showing the key metadata fields of the selected subject. Navigate with arrow keys and press Enter to view the full formatted metadata.

To narrow the search before launching the finder using positional arguments:

```bash
operatree find event              # show only events
operatree find task report        # show tasks matching "report"
operatree find cairo              # search "cairo" across all subject types
```

### Non-interactive mode

Use `--term` and `--type` flags to search without launching the finder — results are returned directly, making it suitable for scripting:

```bash
operatree find --term cairo                       # search all types for "cairo"
operatree find --term report --type task          # search tasks for "report"
```

Add `--plain` to output results as raw YAML for piping into other tools:

```bash
operatree find --term cairo --plain
operatree find --term report --type task --plain | grep owner
```

The search is fuzzy and runs across all metadata fields — name, tags, participants, notes, date, and location. A search for `cairo` will match an event located in Cairo, a task with Cairo in its notes, and a topic tagged with `cairo`.

In interactive mode, selecting a subject displays the full `META.yaml` contents in a clean formatted view. `find` is a read-only command — it never modifies anything.

---

## 6.3 Editing Subject Metadata

Sara needs to update the status of a task and add some notes after a review session:

```bash
operatree edit
```

OperaTree opens the interactive finder. Sara selects the task — her configured editor opens with the `META.yaml` file. She updates the `status` field from `active` to `review` and adds notes. When she closes the editor, OperaTree automatically runs `sync` to update the project metadata index.

To filter before launching the finder:

```bash
operatree edit task               # filter to tasks, then pick one
operatree edit task report        # filter to tasks matching "report", then pick one
```

The editor used is whatever was configured during `operatree init`. It falls back to the `$EDITOR` environment variable if no editor was set in config.

**A note on shared environments:** In Sara and Alex's Syncthing setup, changes Sara makes via `operatree edit` are synced to Alex's machine automatically. When Alex starts his morning, running `operatree sync` ensures his local index reflects Sara's overnight edits — including any `META.yaml` files that Sara may have edited directly outside of OperaTree.

---

## 6.4 Opening a Subject Directory

Alex wants to add the final report PDF to the task directory:

```bash
operatree open
```

The interactive finder launches. Alex selects the task — the file manager opens directly at that subject's directory. He drops the PDF into `04_FINAL/` and closes the file manager.

To filter before launching the finder:

```bash
operatree open task               # filter to tasks, then pick one
operatree open task report        # filter to tasks matching "report"
```

OperaTree opens a file manager window and does not touch any files inside the subject directory. What you put there, how you organise it, and what you name your files is entirely yours.

---

## 6.5 Renaming a Subject

The kickoff event was initially named too generically. Alex wants to rename it:

```bash
operatree rename
```

The interactive finder launches. Alex selects the event — OperaTree prompts for the new name. On confirmation, it renames the subject directory and updates the `META.yaml` and the project metadata index in one operation. Any other subjects that reference this subject by name are updated automatically.

To filter before launching the finder:

```bash
operatree rename event            # filter to events, then pick one
operatree rename event kickoff    # filter to events matching "kickoff"
```

For scripting — target a subject directly by UUID without launching the finder:

```bash
operatree rename --uuid a1b2c3d4 --new-name "FleetFix Kickoff Meeting"
```

The UUID can be obtained from `operatree find --term [name] --plain` and extracted from the output. `--uuid` requires `--new-name` — providing one without the other is an error.

---

## 6.6 Archiving a Subject

Three months in, several tasks are complete and no longer actively referenced. Alex archives them to keep the active subject list clean:

```bash
operatree archive
```

The interactive finder launches. Alex selects the completed task — OperaTree moves the entire subject directory to `99_ARCHIVE/` at the project root. The `META.yaml` is preserved exactly as-is. The archive action is appended to `activity.log`.

```
2026-09-14T16:30:00Z    ARCHIVE    task    "Prepare Workshop Report"    alex@workstation    v0.1.2
```

To filter before launching the finder:

```bash
operatree archive task            # filter to tasks, then pick one
operatree archive task report     # filter to tasks matching "report"
```

For scripting — target a subject directly by UUID without launching the finder:

```bash
operatree archive --uuid a1b2c3d4
```

**Archiving is not deletion.** The subject and all its files remain in `99_ARCHIVE/` indefinitely. If you need to retrieve an archived subject, navigate to `99_ARCHIVE/` in your file manager and move it back manually. A formal restore command is planned for a future release.

---

## 6.7 Subject Workflow at a Glance

Here is how the subject commands fit together in a typical working sequence:

```bash
# Morning — check what is active
operatree summary
operatree find --term active --type task --plain   # list active tasks without the finder

# A meeting happens
operatree add event \
  --name "Vendor Review" \
  --date 2026-06-15 \
  --participants "Alex,Sara" \
  --tags "vendor,review"

# The meeting generates work
operatree add task \
  --name "Vendor Evaluation Report" \
  --owner Alex \
  --status active \
  --related-events "Vendor Review"

# Research needed for the task
operatree add topic \
  --name "Vendor Assessment Frameworks" \
  --related-objective "Select Primary Vendor"

# After the meeting — add notes and open directory for files
operatree edit event vendor       # update the event metadata
operatree open event vendor       # drop files into the event directory

# Task progresses — update its status
operatree edit task               # change status from active to review

# Task is done — archive it directly by UUID
operatree find --term "Vendor Evaluation" --type task --plain | grep uuid
operatree archive --uuid a1b2c3d4
```

Each command is small and focused. The finder, the editor, and the file manager work together as a natural workflow — OperaTree handles the structure, you handle the content.

---

_Next: Section 7 — Understanding Your Project_
