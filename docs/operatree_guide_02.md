# Section 2 — Project Management Standard

---

## 2.1 How OperaTree Thinks About Projects

OperaTree organises project knowledge according to its **purpose**, not its file type.

Most projects eventually become difficult to navigate because files are grouped by format — a folder for documents, a folder for photos, a folder for spreadsheets. This approach makes it nearly impossible to answer practical questions:

- Where is the research behind this decision?
- Which meeting produced this finding?
- Which files support this objective?
- Which datasets were used for this report?

OperaTree answers these questions by organising information around **business meaning and lifecycle**. Every file has a place determined by the question it answers:

| Layer              | Question answered          |
| ------------------ | -------------------------- |
| Events             | What happened?             |
| Project Management | What are we doing?         |
| Research           | What do we know?           |
| Engineering        | What are we building?      |
| Data               | What have we measured?     |
| Deliverables       | What are we communicating? |

When you need to find something, you start with the question — not a search box.

---

## 2.2 Projects, Modules, and Subjects

OperaTree introduces three levels of structure. Understanding these three levels is the key to understanding everything else.

### Project

A project is the top-level container. It is a directory on your filesystem that OperaTree has bootstrapped and registered. Everything lives inside it. A project has a name, a template, and a set of modules.

### Module

A module is a directory within a project that has a fixed, globally consistent prefix and a defined purpose. The prefix is always the same regardless of which project or template you are using — `99_ARCHIVE` is always the archive, `01_EVENTS` is always where events live.

Modules can contain:

- Other modules (nested structure)
- Free-form subdirectories (unmanaged, yours to organise)
- Subjects (managed by OperaTree)

The full set of available modules and their prefixes is:

| Prefix | Module             | Purpose                                                     |
| ------ | ------------------ | ----------------------------------------------------------- |
| `00`   | Admin              | Governance, contacts, templates, guidelines                 |
| `01`   | Events             | Project events — visits, meetings, workshops                |
| `02`   | Project Management | Execution layer — tasks, reports, risks, budgets            |
| `03`   | Legal              | Contracts, NDAs, compliance, approvals                      |
| `04`   | Research           | Knowledge and intelligence — topics, objectives, summaries  |
| `05`   | Engineering        | Architecture, specs, decisions, prototypes                  |
| `06`   | Data               | Full data lifecycle — raw through processed                 |
| `07`   | Tasks              | Active work units (nested under Project Management)         |
| `08`   | Index              | Navigation and knowledge maps                               |
| `09`   | Topics             | Concept-centric knowledge (nested under Research)           |
| `10`   | Objectives         | Goal-centric intelligence (nested under Research)           |
| `11`   | Summaries          | Distilled conclusions (nested under Research)               |
| `12`   | References         | External knowledge library (nested under Research)          |
| `13`   | Audio Notes        | Voice recordings and transcriptions (nested under Research) |
| `14`   | Attachments        | General supporting files (nested under Research)            |
| `15`   | Data Sources       | Source metadata (nested under Data)                         |
| `16`   | Publications       | Academic and formal publications                            |
| `97`   | Media Library      | Shared reusable assets                                      |
| `98`   | Deliverables       | Final external outputs                                      |
| `99`   | Archive            | Historical storage                                          |

**The numbered prefix is not decorative.** It gives every terminal and file manager a consistent, sorted view of the project — the governance layer at the top, the archive at the bottom, always. Once you have used OperaTree for a week, you will navigate by prefix instinctively.

### Subject

A subject is the atomic unit of trackable work or knowledge within a project. It lives inside a module, has its own directory, and contains a `META.yaml` file that makes it searchable and machine-readable.

Every subject has a type, a name, a date, tags, and notes. Beyond these common fields, each subject type has additional fields specific to its purpose. All type-specific fields are optional — they are only written to `META.yaml` when they apply, keeping metadata files clean for subjects that do not use them.

---

## 2.3 Templates

A template defines which modules are created when a project is bootstrapped. It does not change what the modules are — it only decides which ones are included.

Because module prefixes are globally fixed, a module that appears in two different templates will always have the same name, the same prefix, and the same internal structure. A `04_RESEARCH` directory created by the `dev` template is structurally identical to one created by the `research` template.

Four templates are currently available:

---

**`general`** — the minimal template for organised work. Suitable for personal projects, small teams, or any work that does not fit a more specialised category.

Includes: Admin, Events, Project Management (with Tasks), Media Library, Deliverables, Archive.

---

**`dev`** — for software development projects. Adds engineering, data pipeline, and legal modules to the standard base.

Includes: Admin, Events, Project Management (with Tasks), Legal, Research (full), Engineering, Data (with Sources), Media Library, Deliverables, Archive.

Project Management subdirectories: `budgets`, `communications`, `planning`, `reports`, `risks`.

Data subdirectories follow the full pipeline: `01_RAW`, `02_STAGING`, `03_PROCESSED`, `04_ANALYTICS`, `05_MODELS`, `06_EXPORTS`, `99_ARCHIVE`.

---

**`research`** — for academic work, R&D, and analytical projects. Replaces the engineering and data modules with a publications module.

Includes: Admin, Events, Project Management (with Tasks), Legal, Research (full), Publications, Deliverables, Archive.

Publications subdirectories: `drafts`, `review`, `published`.

Project Management subdirectories: `communications`, `planning`, `reports` (no risks or budgets).

---

**`consulting`** — for client engagement work. Combines a full project management layer with research and legal modules, without engineering or data pipeline infrastructure.

Includes: Admin, Events, Project Management (with Tasks), Legal, Research (full), Deliverables, Archive.

Project Management subdirectories: `budgets`, `communications`, `planning`, `reports`, `risks`.

---

### Choosing the right template

Template selection is a one-time decision made at project creation. Adding or removing modules after a project is bootstrapped is not currently supported — doing so risks breaking the structural conventions that OperaTree's project management standard depends on.

Choose carefully before running `operatree create`. As a guide:

| If your work involves...                             | Use template |
| ---------------------------------------------------- | ------------ |
| General organisation, personal projects, light teams | `general`    |
| Software products, APIs, technical platforms         | `dev`        |
| Academic research, R&D, analytical studies           | `research`   |
| Client engagements, consulting, advisory work        | `consulting` |

If you are unsure, err toward a more complete template. An unused module costs nothing — a missing one cannot be added later without risking the integrity of the project structure.

To see available templates before creating a project:

```bash
operatree show templates
```

---

## 2.4 The Five Subject Types

### Event

An event is anything that occurs at a specific point in time — a factory visit, a stakeholder meeting, a workshop, a field survey, a training session. Events are the chronological backbone of a project. They record what happened, when it happened, and who was involved.

Events live in the `01_EVENTS` module. Each event gets its own directory containing subdirectories for agenda, media, notes, documents, and outcomes.

**Fields:**

| Field          | Purpose                 |
| -------------- | ----------------------- |
| `name`         | Event name              |
| `date`         | Date the event occurred |
| `location`     | Where it took place     |
| `participants` | Who was present         |
| `tags`         | Searchable labels       |
| `notes`        | Free-form notes         |

**Example `META.yaml`:**

```yaml
uuid: a1b2c3d4
type: event
name: Cairo Factory Visit
date: "2026-05-14"
location: Cairo Industrial Zone
participants:
  - Hany Mamdouh
  - Ahmed Khalil
tags:
  - factory
  - inspection
  - cairo
notes: Initial site survey for production line assessment.
```

**Rule:** Store raw event information here. Store long-term conclusions and distilled findings in `04_RESEARCH/11_SUMMARIES`.

---

### Task

A task is a unit of work with a lifecycle. It has an owner, a status, and a working environment with defined stages. Tasks answer the question: what are we actively doing?

Tasks live in the `07_TASKS` module, which is nested inside `02_PROJECT_MANAGEMENT`. When you run `operatree add task`, OperaTree creates the task directory and its four stage subdirectories automatically:

```
website-development/
├── 01_INPUTS       # requirements, references, initial materials
├── 02_WORKING      # drafts, iterations, work in progress
├── 03_REVIEW       # reviewed documents, comments, approval stage
├── 04_FINAL        # approved outputs ready for delivery
└── META.yaml
```

You are free to create additional directories inside any of these stages to organise your own files, but the four stage directories themselves are created by OperaTree and should not be renamed or removed.

**Fields:**

| Field            | Purpose                                      |
| ---------------- | -------------------------------------------- |
| `name`           | Task name                                    |
| `date`           | Date created or started                      |
| `owner`          | Person responsible                           |
| `status`         | Current status (e.g. active, blocked, done)  |
| `related_events` | Events that generated or relate to this task |
| `outputs`        | Expected or produced outputs                 |
| `tags`           | Searchable labels                            |
| `notes`          | Free-form notes                              |

**Example `META.yaml`:**

```yaml
uuid: e5f6g7h8
type: task
name: Prepare Production Assessment Report
date: "2026-05-15"
owner: Hany Mamdouh
status: active
related_events:
  - Cairo Factory Visit
outputs:
  - Production Assessment Report v1.0
tags:
  - report
  - assessment
  - production
notes: Based on findings from the Cairo site visit. Draft due end of month.
```

**Rule:** If a file is actively supporting the execution of work, it belongs in a task. Drafts stay inside the task until approved — only finalized outputs move to `98_DELIVERABLES`.

---

### Topic

A topic is a knowledge concept or domain area. Topics answer the question: what is this? They are pure knowledge — not work items, not goals, just structured understanding of a concept that the project needs to develop or reference.

Topics live in the `09_TOPICS` module, nested inside `04_RESEARCH`. Each topic gets its own directory with an overview document, notes, diagrams, and attachments.

```
predictive-maintenance/
├── overview.md       # curated explanation of the concept
├── notes.md          # additional insights and details
├── diagrams/         # concept illustrations
├── attachments/      # supporting PDFs, whitepapers
└── META.yaml
```

**Fields:**

| Field               | Purpose                           |
| ------------------- | --------------------------------- |
| `name`              | Topic name                        |
| `date`              | Date created                      |
| `related_objective` | The objective this topic supports |
| `tags`              | Searchable labels                 |
| `notes`             | Free-form notes                   |

**Example `META.yaml`:**

```yaml
uuid: i9j0k1l2
type: topic
name: Predictive Maintenance
date: "2026-05-10"
related_objective: Reduce Equipment Downtime
tags:
  - maintenance
  - ml
  - iot
notes: Covers condition monitoring, failure prediction models, and sensor integration.
```

**Rule:** A topic contains your understanding of a concept. External sources that you read belong in `12_REFERENCES`. The topic is where you synthesise what you have learned.

---

### Objective

An objective is a goal driving research and decisions. Objectives answer the question: what are we trying to achieve? They are active thinking spaces — living documents that accumulate findings, evidence, and strategy as the project progresses.

Objectives live in the `10_OBJECTIVES` module, nested inside `04_RESEARCH`. Each objective gets its own directory with definition, findings, strategy, and supporting material.

```
reduce-equipment-downtime/
├── definition.md     # what the objective means and why it matters
├── findings.md       # observations and discoveries so far
├── strategy.md       # planned direction and decisions
├── notes/            # working notes and brainstorming
├── discussions/      # AI conversations, meeting extracts
├── attachments/      # supporting evidence
└── META.yaml
```

**Fields:**

| Field     | Purpose                                          |
| --------- | ------------------------------------------------ |
| `name`    | Objective name                                   |
| `date`    | Date created                                     |
| `status`  | Current status (e.g. active, achieved, deferred) |
| `outputs` | Decisions, strategies, or deliverables produced  |
| `tags`    | Searchable labels                                |
| `notes`   | Free-form notes                                  |

**Example `META.yaml`:**

```yaml
uuid: m3n4o5p6
type: objective
name: Reduce Equipment Downtime
date: "2026-04-01"
status: active
outputs:
  - Maintenance Strategy v1
  - Sensor Integration Proposal
tags:
  - downtime
  - maintenance
  - strategy
notes: Primary research objective for the FleetFix project. Linked to predictive maintenance topic.
```

**Rule:** Objectives are where strategy is built. They connect topics (what we know) to tasks (what we are doing) and ultimately to deliverables (what we produce).

---

### Data Source

A data source is a record of an external dataset or data feed used in the project. It does not store the data itself — it stores the metadata about where the data came from, how large it is, and which objective it supports. The actual data lives in `06_DATA/01_RAW`.

Data sources live in the `15_DATASOURCES` module, nested inside `06_DATA`.

**Fields:**

| Field             | Purpose                                  |
| ----------------- | ---------------------------------------- |
| `name`            | Dataset name                             |
| `date`            | Date acquired or registered              |
| `source`          | Origin (e.g. Kaggle, internal team, API) |
| `sourceLink`      | URL or path to the original data         |
| `sourceObjective` | The objective this data supports         |
| `sourceDataSize`  | Size or volume of the dataset            |
| `tags`            | Searchable labels                        |
| `notes`           | Free-form notes                          |

**Example `META.yaml`:**

```yaml
uuid: q7r8s9t0
type: datasource
name: Equipment Sensor Readings 2025
date: "2026-05-01"
source: Internal IoT team
sourceLink: /06_DATA/01_RAW/sensor_readings_2025.csv
sourceObjective: Reduce Equipment Downtime
sourceDataSize: 2.4 GB
tags:
  - sensors
  - iot
  - raw
notes: Hourly readings from 42 sensors across three production lines. Covers Jan–Dec 2025.
```

**Rule:** Register every external dataset as a data source. This creates a traceable lineage from raw data through processing to findings and decisions.

---

## 2.5 The Subject Lifecycle

Every subject follows the same basic lifecycle:

```
add → edit → archive
```

**Add** (`operatree add [type]`) — creates the subject directory, writes the initial `META.yaml`, and appends a `CREATE` entry to `activity.log`.

**Edit** (`operatree edit [type] [term]`) — opens the subject's `META.yaml` in your configured editor. The metadata index is updated automatically when the editor closes. Appends an `EDIT` entry to `activity.log`.

**Archive** (`operatree archive [type] [term]`) — moves the subject directory to `99_ARCHIVE`. Appends an `ARCHIVE` entry to `activity.log`.

Subjects can also be **renamed** (`operatree rename`) which is recorded in `activity.log`. A `delete` action is planned for a future release.

---

## 2.6 The Activity Log

Every action OperaTree takes on a subject is recorded in `activity.log` at the project root. The log is append-only, tab-separated, and pipe-friendly.

```
2026-05-20T10:08:39Z    CREATE    event    "Cairo Factory Visit"    hany@optiplex7040    v0.1.0
2026-05-20T11:22:14Z    CREATE    task     "Prepare Report"         hany@optiplex7040    v0.1.0
2026-05-20T14:05:03Z    EDIT      task     "Prepare Report"         hany@optiplex7040    v0.1.2
2026-05-21T09:14:22Z    ARCHIVE   task     "Old Vendor Analysis"    hany@optiplex7040    v0.1.2
```

Columns: `timestamp`, `action`, `type`, `name`, `user@host`, `version`.

Current actions recorded: `CREATE`, `EDIT`, `ARCHIVE`. A `DELETE` action is planned for a future release.

Because the log is tab-separated plain text, it pipes naturally into standard tools:

```bash
grep CREATE activity.log | cut -f3 | sort | uniq -c   # count creations by type
grep event activity.log                                 # all event actions
grep hany activity.log | tail -20                       # last 20 actions by user
```

You can commit `activity.log` to version control for a full audit trail, or add it to `.gitignore` to keep it local — both are valid choices depending on your team's needs.

---

## 2.7 Where Files Belong — A Decision Guide

When you create a file and need to decide where it goes, ask these questions in order:

**Was it created during a specific event?**
→ `01_EVENTS/[event-name]/`

**Is it helping execute current work?**
→ `02_PROJECT_MANAGEMENT/07_TASKS/[task-name]/`

**Is it a legal document?**
→ `03_LEGAL/`

**Is it knowledge or research?**
→ `04_RESEARCH/` — then ask:

- What is this concept? → `09_TOPICS/`
- What are we trying to achieve? → `10_OBJECTIVES/`
- What matters (conclusions)? → `11_SUMMARIES/`
- What have we read (external sources)? → `12_REFERENCES/`

**Is it technical design or system documentation?**
→ `05_ENGINEERING/`

**Is it data?**
→ `06_DATA/` — then ask which stage: raw, staging, processed, analytics, exports

**Is it a reusable asset used across multiple deliverables?**
→ `97_MEDIA_LIBRARY/`

**Is it a finalized output going to stakeholders?**
→ `98_DELIVERABLES/`

**Is it obsolete but must be retained?**
→ `99_ARCHIVE/`

The Research layer deserves special attention — it is the most important and most often misunderstood part of the structure:

| Module          | Question                       | What belongs here                                        |
| --------------- | ------------------------------ | -------------------------------------------------------- |
| `09_TOPICS`     | What is this?                  | Explanations, theory, frameworks, your synthesis         |
| `10_OBJECTIVES` | What are we trying to achieve? | Findings, strategy, evidence, brainstorming              |
| `11_SUMMARIES`  | What matters?                  | Conclusions distilled from events, research, discussions |
| `12_REFERENCES` | What have we read?             | Books, papers, standards, vendor docs — original sources |

The key distinction: `12_REFERENCES` holds what others wrote. `09_TOPICS` holds what you understood from it.

---

_Next: Section 3 — Installation_
