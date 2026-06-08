# OperaTree User Guide

---

# Section 1 — Philosophy

## 1.1 The Problem with Existing Tools

Most project management tools share a common flaw: they own your data.

Your tasks, notes, events, and decisions live inside a database you cannot read, in a format you did not choose, on servers you do not control. When the tool shuts down, raises its prices, or simply stops fitting how you work — your history goes with it. Migration is painful. History is lost. You start over.

Even when tools do offer exports, what comes out is rarely useful. A CSV of task titles. A JSON blob that only their next product can import. The structure, the relationships, the context — gone.

This is not a minor inconvenience. It is a fundamental misalignment between the tool and the person using it. Your work should belong to you.

---

## 1.2 The Filesystem Is Already a Database

Here is something most project management tools want you to forget: your operating system already ships with a perfectly capable storage engine. It is called the filesystem.

Files and directories have existed for decades. They are readable by every text editor ever made. They are searchable with standard tools. They survive software changes, company acquisitions, and hardware migrations. They work offline, on any device, and with any version control system.

OperaTree does not replace your filesystem — it organises it.

Instead of inventing a proprietary storage format, OperaTree applies a consistent, opinionated directory structure to your project. Every subject — every event, task, topic, or objective — gets its own directory. Every directory contains a `META.yaml` file with structured metadata. The rest of the directory is yours to fill with whatever files your work produces.

The result is a project that any person can navigate in any file manager, search with any tool, and back up with any strategy — with or without OperaTree installed.

---

## 1.3 Why Plain Directories and YAML

The choice of directories and YAML is deliberate.

**Directories** map naturally to how people already think about projects. A folder for events. A folder for tasks. A folder for each deliverable. No translation layer, no abstraction — just the same mental model you use every day, made consistent and navigable.

**YAML** is the lightest structured format a human can read without tooling. Open a `META.yaml` in any text editor and you will understand it immediately. Fields are labelled. Values are plain text. There is no encoding to decode, no binary to parse. If OperaTree disappeared tomorrow, your metadata would still be readable.

Together, they mean your project is never locked to any software. The structure outlives the tool that created it.

---

## 1.4 Why OperaTree Never Touches Your File Contents

OperaTree creates directories and writes metadata. That is all it does to your filesystem.

It never opens your documents. It never modifies your reports, your notes, your spreadsheets, or any other file you create inside a subject directory. The contents of your files are entirely yours — OperaTree only manages the structure around them.

This boundary is intentional. The moment a tool starts reading or modifying your file contents, you have introduced a dependency. Your documents begin to depend on that tool's format, its assumptions, its continued existence. OperaTree refuses that dependency.

Your files are yours. OperaTree only provides the shelves.

---

## 1.5 The PMI/PRINCE2 Foundation

OperaTree's directory structure is not arbitrary. It is modelled on a hybrid of two widely used project management frameworks: **PMI** (Project Management Institute) and **PRINCE2** (Projects in Controlled Environments), adapted and generalised to fit a broader range of use cases.

From **PMI**, OperaTree inherits:

- The concept of a **work breakdown structure** — projects decomposed into discrete, trackable units of work
- Clear separation between **planning**, **execution**, and **delivery** layers
- Explicit **risk**, **budget**, and **communication** management areas

From **PRINCE2**, OperaTree inherits:

- A **governance and control layer** (`00_ADMIN/`) that sits above execution
- The idea of **stages** — work flows from inputs through working drafts to review and final output
- An **audit trail** — every action is logged, every change is traceable

These concepts are then generalised so they apply equally to a software development project, a research programme, a legal case, a company's day-to-day operations, or a personal knowledge base. The framework adapts to the work — not the other way around.

---

## 1.6 The Directory Layer Model

Every OperaTree project follows the same top-level structure:

```
your-project/
├── 00_ADMIN/               # governance, contacts, templates
├── 01_EVENTS/              # visits, workshops, meetings
├── 02_PROJECT_MANAGEMENT/  # tasks, reports, risks, budgets
├── 03_LEGAL/               # contracts, NDAs, compliance
├── 04_RESEARCH/            # topics, objectives, summaries
├── 05_ENGINEERING/         # architecture, specs, decisions
├── 06_DATA/                # raw → staging → processed pipeline
├── 07_MEDIA_LIBRARY/       # shared reusable assets
├── 08_DELIVERABLES/        # final external outputs
├── 99_ARCHIVE/             # historical storage
└── activity.log            # append-only audit trail
```

Each layer has a clear responsibility and a clear boundary. Events belong in `01_EVENTS/`. Signed contracts belong in `03_LEGAL/`. Final reports belong in `08_DELIVERABLES/`. When everyone on a project agrees on these boundaries, finding anything becomes trivial — you already know which layer to look in before you open a file manager.

**The numbered prefixes are not decorative.** They serve three practical purposes:

- **Natural sort order** — every terminal, every file manager, every operating system sorts numbered prefixes correctly without configuration. The governance layer is always at the top. The archive is always at the bottom.
- **Muscle memory** — once you have used OperaTree for a week, you know that `04` means research and `05` means engineering. Tab completion in the terminal becomes fast and precise.
- **Layer priority** — the numbers communicate importance and sequence. Administration comes before events. Events come before deliverables. The structure teaches the workflow.

---

## 1.7 Two Modes: Projects and Operations

OperaTree was designed for project management — a bounded piece of work with a start, a scope, and a set of deliverables. It works well for software products, research programmes, engineering projects, and consulting engagements.

But it turns out the same structure works equally well for something broader: **running an organisation**.

A company's day-to-day operations have the same ingredients as a project. There are meetings (events). There is ongoing work (tasks). There is knowledge to build and maintain (topics, objectives). There are goals to track and deliverables to produce. The only difference is that operations never end — there is no delivery date, only a continuous cycle.

OperaTree handles both modes without any special configuration. You simply treat the organisation itself as a project. One OperaTree project, continuously active, shared across everyone who needs it.

**Project mode** — one OperaTree project per client, product, or engagement:

```
FleetFix/           ← vehicle management platform
Anchor/             ← maritime logistics product
I-DNTITI/           ← identity verification service
OperaTree-Dev/      ← OperaTree itself
```

**Operations mode** — the organisation as a project:

```
atentec-management/ ← company operations, replicated across offices via Syncthing
```

In operations mode, a typical working day might look like this:

```bash
operatree summary                        # what do we have? what was added last?
operatree add event                      # record this morning's meeting
operatree edit                           # add notes after the meeting closes
operatree add task                       # the meeting generated a new piece of work
operatree add topic                      # studying something new? capture it
operatree add objective                  # set a goal and link it to relevant work
```

The tool does not change between modes. The structure does not change. What changes is only the scope of what you are managing.

---

## 1.8 How It Fits UNIX Pipelines

OperaTree is designed to work naturally with standard UNIX tools. Its output is plain text. Its metadata is plain YAML. Its activity log is tab-separated columns.

This means you can pipe, grep, cut, and sort OperaTree output without any special integration:

```bash
# count subjects created by type
grep CREATE activity.log | cut -f3 | sort | uniq -c

# find all events from last month
grep event activity.log | grep 2026-05

# list all tasks in raw YAML for processing
operatree describe --plain | grep -A5 "type: task"

# pipe project summary into a report
operatree describe --plain > project-snapshot.yaml
```

The `--plain` flag on `operatree describe` is specifically designed for this — it outputs raw YAML instead of the styled terminal view, making it suitable as input to other tools.

Non-interactive subject creation is also supported, enabling OperaTree to be driven from scripts:

```bash
# create a task from a script without interactive prompts
operatree add task --name "Deploy to staging" --date 2026-06-01
```

OperaTree does not try to replace your shell, your editor, or your other tools. It is one piece in a larger workflow — and it is designed to fit that role cleanly.

---

## 1.9 Your Data Outlives Any Software

This is the core promise of OperaTree, and it is worth stating plainly.

If OperaTree stopped being maintained tomorrow, your projects would be completely intact. Every event, task, topic, and objective would still be in its directory. Every `META.yaml` would still be readable in any text editor. Every file you created would be exactly where you left it.

You would lose the CLI interface. You would not lose your work.

This is the opposite of how most project management tools operate. It requires discipline in the design — no proprietary formats, no binary files, no database that only the tool can read. But it means that the value you put into OperaTree compounds over time, rather than being held hostage by a subscription.

Your project history is yours. Permanently.

---

_Next: Section 2 — Project Management Standard_
