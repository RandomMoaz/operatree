# Section 4 — Getting Started

---

This section walks through setting up OperaTree for the first time and creating your first project. Rather than listing commands in isolation, we follow a single continuous story — a project manager named Alex who is setting up OperaTree to manage a new client engagement.

By the end of this section, you will have OperaTree configured, a project created, and your first subject added.

---

## 4.1 First Run — Initializing OperaTree

After installation, the first thing Alex does is open a terminal and run:

```bash
operatree init
```

OperaTree will refuse to run any other command until `init` has been completed. It checks for a configuration file at startup — if none exists, it directs you here.

The interactive prompt asks for three things:

**Standard directory** _(required)_
The base directory where new projects will be created by default. Alex sets this to `/home/alex/projects`. This becomes the default location for `operatree create` — you can always override it with `-d` when creating a project.

**Editor** _(optional)_
The text editor OperaTree opens when you run `operatree edit`. Alex sets this to `nvim`. If left empty, OperaTree falls back to your `$EDITOR` environment variable. If neither is set, OperaTree will not be able to open metadata files for editing until one is configured.

**File manager** _(optional)_
The file manager OperaTree opens when you run `operatree open` or `operatree goto`. Alex sets this to `nautilus`. If left empty, those commands will not function until a file manager is configured. There is no automatic fallback for this setting.

Once `init` completes, the configuration file is written to the appropriate location for your platform:

| Platform    | Config location                                          |
| ----------- | -------------------------------------------------------- |
| Linux       | `~/.config/operatree/operatree.yaml`                     |
| Linux (XDG) | `$XDG_CONFIG_HOME/operatree/operatree.yaml`              |
| macOS       | `~/Library/Application Support/operatree/operatree.yaml` |
| Windows     | `%APPDATA%\operatree\operatree.yaml`                     |

You only need to run `init` once per user account. If you reinstall OperaTree, restore your config file from backup, or copy it from another machine, `init` will detect the existing configuration and skip the setup.

---

## 4.2 Creating Your First Project

With OperaTree configured, Alex creates a new project for the client engagement:

```bash
operatree create fleetfix -t consulting
```

OperaTree scaffolds the full directory structure for the `consulting` template under the standard directory:

```
/home/alex/projects/fleetfix/
├── 00_ADMIN/
│   ├── contacts/
│   ├── governance/
│   ├── guidelines/
│   └── templates/
├── 01_EVENTS/
├── 02_PROJECT_MANAGEMENT/
│   ├── 07_TASKS/
│   ├── budgets/
│   ├── communications/
│   ├── planning/
│   ├── reports/
│   └── risks/
├── 03_LEGAL/
│   ├── contracts/
│   ├── ndas/
│   ├── compliance/
│   ├── approvals/
│   └── templates/
├── 04_RESEARCH/
│   ├── 08_INDEX/
│   ├── 09_TOPICS/
│   ├── 10_OBJECTIVES/
│   ├── 11_SUMMARIES/
│   ├── 12_REFERENCES/
│   ├── 13_AUDIO_NOTES/
│   └── 14_ATTACHMENTS/
├── 98_DELIVERABLES/
│   ├── client_documents/
│   ├── presentations/
│   ├── reports/
│   └── submissions/
├── 99_ARCHIVE/
│   ├── closed_tasks/
│   ├── deprecated/
│   └── old_versions/
└── activity.log
```

OperaTree then prompts Alex to set `fleetfix` as the default project. Alex confirms — from this point on, every command automatically targets `fleetfix` without needing the `-d` flag.

To verify the default is set:

```bash
operatree show default
```

To view a structured summary of the project:

```bash
operatree describe
```

---

## 4.3 Adding Your First Subject

The project is created. Now Alex needs to record the kickoff meeting that happened this morning.

```bash
operatree add event
```

OperaTree launches an interactive form. Alex fills in:

- **Name:** `FleetFix Kickoff Meeting`
- **Date:** `2026-06-07`
- **Location:** `Cairo HQ`
- **Participants:** `Alex`, `Sara Youssef`, `Omar Nabil`
- **Tags:** `kickoff`, `client`, `planning`
- **Notes:** `Discussed project scope, timeline, and initial deliverables.`

OperaTree creates the event directory inside `01_EVENTS/` and writes the `META.yaml`:

```
01_EVENTS/
└── fleetfix-kickoff-meeting/
    ├── 01_AGENDA/
    ├── 02_MEDIA/
    ├── 03_NOTES/
    ├── 04_DOCUMENTS/
    ├── 05_OUTCOMES/
    └── META.yaml
```

The creation is recorded in `activity.log`:

```
2026-06-07T09:14:00Z    CREATE    event    "FleetFix Kickoff Meeting"    alex@workstation    v0.1.2
```

Alex opens the event directory to add the meeting agenda document:

```bash
operatree open event kickoff
```

The file manager opens directly at the event directory. Alex drops the agenda PDF into `01_AGENDA/` and the meeting notes into `03_NOTES/`. OperaTree does not touch these files — they are entirely Alex's to manage.

---

## 4.4 Setting the Default Project

Later that day, Alex needs to switch context to check on another project. First, a quick look at what is currently set as default:

```bash
operatree show default
```

To switch the default to a different tracked project:

```bash
operatree use
```

An interactive picker lists all tracked projects. Alex selects a different one, and all subsequent commands now target it. To switch back to `fleetfix`:

```bash
operatree use
```

Pick `fleetfix` from the list. Done.

OperaTree always requires an explicit selection even if only one project is tracked — this is intentional. It keeps you aware of which project you are working on, which matters when project directories are shared across machines or team members.

---

## 4.5 Tracking an Existing Project

The following week, Alex's colleague Sara joins the engagement. Sara's machine already has the `fleetfix` project directory — it is synced via Syncthing from Alex's machine and appears at `/home/sara/shared/fleetfix`.

Sara installs OperaTree and runs `init` to set up her own configuration. Then, instead of creating a new project, she simply registers the existing one:

```bash
operatree track -d /home/sara/shared/fleetfix
```

OperaTree adds `fleetfix` to Sara's tracked projects list. She can now run every command against it — `find`, `add`, `edit`, `summary`, `describe` — exactly as Alex does. Her configuration is hers; the project data is shared.

This is one of OperaTree's core design principles: **configuration is per user, data is per project**. Two people can work on the same project directory with their own editors, file managers, and default settings, without any conflict.

If Sara's config file is ever restored from backup or copied from another machine, she does not need to run `track` again — the project is already registered in the restored config.

---

## 4.6 A Typical Morning

To close the section, here is what a typical morning looks like once OperaTree is part of your workflow:

```bash
# Check the pulse of the project
operatree summary

# A meeting is about to start — record it
operatree add event

# The meeting generated a task
operatree add task

# Something new to learn about for this task
operatree add topic

# Open the task directory to add files
operatree open task

# Find the event from last week
operatree find event kickoff

# Review and update the meeting notes
operatree edit event kickoff
```

Each command is small and focused. Together they form a complete workflow that keeps your project structured, searchable, and under your control.

---

_Next: Section 5 — Managing Projects_
