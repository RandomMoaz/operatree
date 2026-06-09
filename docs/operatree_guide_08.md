# Section 8 — Scripting, Automation, and the OperaTree Ecosystem

---

OperaTree is designed to be one piece in a larger workflow. Its plain text output,
filesystem-first storage, and pipe-friendly design make it a natural fit for
scripting, automation, and integration with other tools. This section covers how
to use OperaTree beyond the interactive terminal — and introduces the broader
ecosystem of tools being built around it.

---

## 8.1 Non-Interactive Subject Creation

By default, `operatree add` launches an interactive form. Providing `--name`
bypasses the interactive form entirely and creates the subject immediately.
All subject fields are available as flags, enabling fully scripted subject
creation without any user interaction.

Common fields available for all subject types:

```bash
operatree add event \
  --name "Weekly Sync" \
  --date 2026-06-15 \
  --location "Cairo HQ" \
  --participants "Alex,Sara" \
  --tags "sync,weekly" \
  --notes "Weekly team sync."

operatree add task \
  --name "Prepare Q3 Report" \
  --date 2026-06-15 \
  --owner Alex \
  --status active \
  --related-events "Weekly Sync" \
  --outputs "Q3 Report v1.0" \
  --tags "report,q3"

operatree add topic \
  --name "Decarbonization Frameworks" \
  --date 2026-06-15 \
  --related-objective "Reduce Carbon Footprint" \
  --tags "sustainability,frameworks"

operatree add objective \
  --name "Reduce Carbon Footprint" \
  --date 2026-06-15 \
  --status active \
  --tags "sustainability,kpi"

operatree add datasource \
  --name "Emissions Dataset 2025" \
  --date 2026-06-15 \
  --source "Environmental Team" \
  --source-link "/06_DATA/01_RAW/emissions_2025.csv" \
  --source-objective "Reduce Carbon Footprint" \
  --source-datasize "450MB" \
  --tags "emissions,raw"
```

When `--name` is provided, fields not supplied by flags are left at their zero
values in `META.yaml` and can be filled in later with `operatree edit`.

This makes it straightforward to create subjects from scripts, cron jobs, or
other automated pipelines:

```bash
# Create a daily standup event from a cron job
0 9 * * 1-5 operatree add event \
  --name "Daily Standup $(date +%Y-%m-%d)" \
  --date "$(date +%Y-%m-%d)" \
  --participants "Alex,Sara,Omar" \
  --tags "standup,daily" \
  -d /home/alex/projects/fleetfix
```

---

## 8.2 Pipe-Friendly Output

OperaTree's output is designed to work naturally with standard UNIX tools.

### Project metadata as YAML

```bash
operatree describe --plain                          # raw YAML project structure
operatree describe --plain | grep -A5 "type: task" # extract task sections
operatree describe --plain > snapshot.yaml          # save project snapshot
```

### Activity log analysis

The `activity.log` file is tab-separated plain text — it pipes directly into
`grep`, `cut`, `sort`, `awk`, and any other standard tool:

```bash
# Count subject creations by type
grep CREATE activity.log | cut -f3 | sort | uniq -c

# All actions on events
grep event activity.log

# Last 20 actions by a specific user
grep alex activity.log | tail -20

# Everything that happened this week
grep "^2026-06-1" activity.log

# Tasks created this month
grep CREATE activity.log | grep task | grep "^2026-06"

# Pipe into a report
grep CREATE activity.log | awk -F'\t' '{print $3, $4}' > creation-report.txt
```

### Non-interactive find

`operatree find` supports non-interactive mode via `--term` and `--type` flags,
returning results directly without launching the finder:

```bash
operatree find --term cairo                        # search all types for "cairo"
operatree find --term report --type task           # search tasks for "report"
operatree find --term active --type task --plain   # output as raw YAML
```

The `--plain` flag outputs results as raw YAML, making `find` a query engine
for scripting pipelines. A powerful pattern combining `find`, `grep`, and
`archive` to bulk-archive completed tasks:

```bash
# Find all done tasks and archive them by UUID
operatree find --term done --type task --plain \
  | grep uuid \
  | awk '{print $2}' \
  | xargs -I{} operatree archive --uuid {}
```

Similarly, rename a subject by UUID without interactive selection:

```bash
# Get the UUID of a subject
operatree find --term "kickoff" --type event --plain | grep uuid

# Rename it directly
operatree rename --uuid a1b2c3d4 --new-name "FleetFix Kickoff Meeting"
```

---

## 8.3 Version Control Integration

Because OperaTree projects are plain files and directories, they work naturally
with any version control or sync strategy. No special integration is needed —
your existing tools already understand the structure.

### The two-layer model

Automated version control works best as two layers:

**Layer 1 — Change detection:** a file watcher monitors the project directory
and triggers an action when files are added, modified, or deleted.

**Layer 2 — The action:** what runs when a change is detected — a git commit,
an rsync, a push to a remote.

### watchexec

[`watchexec`](https://github.com/hanymamdouh82/watchexec) is a lightweight,
config-driven file watcher built in Go. It monitors specified directories and
executes predefined commands when changes are detected. Each directory can have
its own command, check interval, and exclusion list. It can run as a systemd
service for continuous background monitoring.

All behaviour is defined in a `conf.yml` file:

```yaml
dirs:
  /home/alex/projects/fleetfix:
    bin: operatree
    args:
      - sync
      - -d
      - /home/alex/projects/fleetfix
    stdout: true
    delay: 5
    exclude:
      - .git
      - activity.log
```

This configuration watches the `fleetfix` project directory and runs
`operatree sync` whenever a change is detected, with a 5-second interval
between checks. The `.git` directory and `activity.log` are excluded to
avoid unnecessary trigger cycles.

**Auto-commit on every change:**

```yaml
dirs:
  /home/alex/projects/fleetfix:
    bin: bash
    args:
      - -c
      - "git -C /home/alex/projects/fleetfix add -A && git -C /home/alex/projects/fleetfix commit -m 'auto: $(date -u +%Y-%m-%dT%H:%M:%SZ)'"
    stdout: true
    delay: 10
    exclude:
      - .git
```

**Mirror to a backup location on every change:**

```yaml
dirs:
  /home/alex/projects/fleetfix:
    bin: rsync
    args:
      - -av
      - /home/alex/projects/fleetfix/
      - /backup/fleetfix/
    stdout: false
    delay: 10
    exclude:
      - .git
```

**Watching multiple projects in a single config:**

```yaml
dirs:
  /home/alex/projects/fleetfix:
    bin: operatree
    args:
      - sync
      - -d
      - /home/alex/projects/fleetfix
    stdout: true
    delay: 5
    exclude:
      - .git
      - activity.log

  /home/alex/projects/anchor:
    bin: operatree
    args:
      - sync
      - -d
      - /home/alex/projects/anchor
    stdout: true
    delay: 5
    exclude:
      - .git
      - activity.log
```

Run watchexec with a config file:

```bash
watchexec -c /home/alex/.config/watchexec/conf.yml
```

Or install it as a systemd service for continuous background operation —
see the [watchexec repository](https://github.com/hanymamdouh82/watchexec)
for the service file and setup instructions.

The Syncthing pattern from earlier works naturally with watchexec as a
service: when Syncthing brings in updated `META.yaml` files from another
machine, watchexec detects the changes and automatically runs
`operatree sync` to update the local index without any manual intervention.

### Syncthing

[`Syncthing`](https://syncthing.net) handles Layer 2 for multi-device and
multi-user setups. It continuously replicates the project directory across
all connected machines in real time, with no server required.

The recommended pattern for shared OperaTree projects:

```
Machine A (Alex)          Machine B (Sara)
fleetfix/ ←──────────────── fleetfix/
    ↑    Syncthing sync          ↑
watchexec                   watchexec
    ↓                           ↓
operatree sync              operatree sync
```

Each machine runs its own `watchexec` instance watching for `META.yaml` changes.
When Sara edits a subject on her machine, Syncthing replicates the change to
Alex's machine, `watchexec` detects the incoming `META.yaml` update, and
`operatree sync` runs automatically. Alex's metadata index stays current without
any manual action.

---

## 8.4 The OperaTree Ecosystem

OperaTree CLI is the foundation. A growing set of tools is being built on top of
it, each adding a layer of capability without changing the underlying filesystem
structure. All ecosystem tools treat the OperaTree project directory as the source
of truth — they read from it, they may write metadata, but they never reorganise
or lock your files.

---

### operatree-rag

`operatree-rag` adds semantic search and AI-assisted knowledge retrieval to any
OperaTree project.

It works by embedding and vectorising two sources of knowledge from your project:

- **Subject metadata** — every `META.yaml` across all modules
- **Module file contents** — documents, notes, and text files inside your project directories

Once indexed, you can query your project in natural language:

```bash
operatree answer "what was my latest meeting about?"
operatree answer "what do I know about decarbonization?"
operatree answer "which tasks are related to the vendor evaluation?"
```

The result is a ranked list of subjects and files that match the query using
vector similarity — not keyword matching. A topic you tagged `emissions-reduction`
will surface for a query about `carbon footprint` because the concepts are
semantically close, even if the words do not match.

`operatree-rag` can also pipe results directly to a local LLM to generate a
synthesised answer from your project knowledge:

```bash
operatree answer "summarise what we know about predictive maintenance" --llm ollama
```

This turns your OperaTree project into a private, local knowledge base that you
can query conversationally — without sending your project data to any external
service.

`operatree-rag` is currently under active development.

---

### operatree-daemon

`operatree-daemon` is a background service that provides an API layer over your
OperaTree projects. It maintains a SQLite index mirror of project metadata for
fast queries, runs a sync engine to keep the index current, and exposes a
configuration-driven API that higher-level tools can use.

The daemon is designed so that the CLI continues to work independently of it —
it is an optional performance and integration layer, not a dependency.

`operatree-daemon` is currently under active development.

---

### operatree-gui

`operatree-gui` is a graphical interface for OperaTree built on top of
`operatree-daemon`. It provides a visual project browser, subject management,
and search — for users who prefer a GUI over the terminal without giving up
the filesystem-first data model.

`operatree-gui` is currently under active development.

---

### Design principle across the ecosystem

Every tool in the OperaTree ecosystem follows the same rule: **the filesystem
is the source of truth**. No ecosystem tool stores data in a format that the
others cannot read. No tool requires another to be running. If `operatree-daemon`
stops, the CLI still works. If `operatree-rag` is not installed, your project
is unchanged. Each tool adds capability — none of them adds dependency.

---

## 8.5 Aliases and Shell Integration

OperaTree's binary name is intentionally descriptive. For day-to-day use, a
short alias reduces friction significantly. Add your preferred alias to your
shell profile (`~/.bashrc`, `~/.zshrc`, or equivalent):

```bash
alias ot='operatree'
```

With this alias, the full command set becomes:

```bash
ot summary
ot add event
ot find task
ot edit task report
ot open event kickoff
```

On Windows, add the alias to your PowerShell profile (`$PROFILE`):

```powershell
Set-Alias ot operatree
```

The alias is personal — it lives in your shell config, not in any OperaTree
file. Team members can each choose their own alias or none at all.

---

_Next: Section 9 — Configuration Reference_
