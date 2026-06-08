# Section 7 — Understanding Your Project

---

This section covers the three commands that help you orient yourself within a project —
`describe`, `summary`, and `explain`. They are all read-only: they never create,
modify, or move anything. Think of them as your project's dashboard.

---

## 7.1 Two Views of a Project

OperaTree gives you two complementary views of a project:

**Structure** — what the project looks like. Which modules exist, how they are
organised, what the directory tree is. This is `operatree describe`.

**Content** — what is in the project. How many subjects, what types, what statuses,
where they live. This is `operatree summary`.

Neither view is complete without the other. A new team member needs the structure
view first — to understand what the project is organised around. Then the content
view — to understand what work is currently happening inside it.

---

## 7.2 Describing the Project Structure

Omar has just joined the `fleetfix` engagement. Before he does anything else, he
wants to understand how the project is organised:

```bash
operatree describe
```

OperaTree prints a styled, colored view of the full directory tree for the project,
as defined by its template. This is the skeleton of the project — the modules and
subdirectories that OperaTree created during `operatree create`. Individual subjects
are not shown here; this is purely structural.

For a project created with the `consulting` template, the output looks like:

```
fleetfix/
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
│   ├── approvals/
│   ├── compliance/
│   ├── contracts/
│   ├── ndas/
│   └── templates/
├── 04_RESEARCH/
│   ├── 08_INDEX/
│   ├── 09_TOPICS/
│   ├── 10_OBJECTIVES/
│   ├── 11_SUMMARIES/
│   ├── 12_REFERENCES/
│   │   ├── articles/
│   │   ├── books/
│   │   ├── standards/
│   │   └── vendor_docs/
│   ├── 13_AUDIO_NOTES/
│   │   ├── indexed/
│   │   ├── raw/
│   │   └── transcriptions/
│   └── 14_ATTACHMENTS/
├── 98_DELIVERABLES/
│   ├── client_documents/
│   ├── presentations/
│   ├── reports/
│   └── submissions/
└── 99_ARCHIVE/
    ├── closed_tasks/
    ├── deprecated/
    └── old_versions/
```

This view answers the question: _what is this project organised around?_ It is
particularly useful when onboarding new team members, setting up a project on a
new machine, or simply reminding yourself which template was used and what modules
are available.

For piping into other tools or saving as a snapshot:

```bash
operatree describe --plain              # raw YAML output instead of styled view
operatree describe --plain | grep tags  # pipe into standard tools
operatree describe --plain > fleetfix-structure.yaml  # save to file
```

To describe a project other than the current default:

```bash
operatree describe -d /home/alex/projects/anchor
```

---

## 7.3 Summarising Project Content

Alex starts every morning with:

```bash
operatree summary
```

Where `describe` shows structure, `summary` shows content — what subjects exist,
how many, what types, what statuses, and which modules they live in. Here is
a sample output for `fleetfix` mid-engagement:

```
fleetfix  ·  summary
────────────────────────────────────────

Total Subjects   14
Latest Activity  2026-06-14

By Type
────────────────────────────────────────
  EVENT          ████████░░░░░░░░░░░░   6
  TASK           █████░░░░░░░░░░░░░░░   4
  TOPIC          ███░░░░░░░░░░░░░░░░░   3
  OBJECTIVE      █░░░░░░░░░░░░░░░░░░░   1

Modules
────────────────────────────────────────
  01_EVENTS                  6 subject(s)
  02_PROJECT_MANAGEMENT      4 subject(s)
    ↳ 07_TASKS               4
  04_RESEARCH                4 subject(s)
    ↳ 09_TOPICS              3
    ↳ 10_OBJECTIVES          1
```

At a glance this tells Alex: the project has 14 subjects, the last activity was
yesterday, most work is in events and tasks, and the research layer is starting
to grow. No file manager, no browsing — just orientation.

The `Latest Activity` field shows the date of the most recent entry in
`activity.log`. A useful complement would be showing the subject name of that
last action — this is planned as an enhancement in a future release.

To summarise a specific project:

```bash
operatree summary -d /home/alex/projects/anchor
```

---

## 7.4 Understanding the Directory Philosophy

New to OperaTree and unsure where a file belongs? The `explain` command renders
the full directory philosophy guide directly in your terminal:

```bash
operatree explain
```

The output covers every module — what it is for, what belongs in it, what does
not, and how it relates to adjacent modules. It is the same guide that informed
the design of the directory structure itself, rendered without leaving your terminal.

Since the output is long, piping to a pager makes it easier to read:

```bash
operatree explain | less
```

`explain` is particularly useful in two situations:

- **Onboarding** — a new team member who has never used OperaTree can run
  `explain` before touching anything and immediately understand the project's
  organisational logic
- **File placement decisions** — when you are unsure whether something belongs
  in `09_TOPICS` or `10_OBJECTIVES`, or in `07_TASKS` or `98_DELIVERABLES`,
  `explain` gives you the answer without leaving the terminal

---

## 7.5 Using the Three Commands Together

Here is how `describe`, `summary`, and `explain` work together as a complete
orientation toolkit:

```bash
# New to the project? Start here
operatree explain               # understand the directory philosophy
operatree describe              # see how this specific project is structured
operatree summary               # see what work is currently inside it

# Daily orientation
operatree summary               # pulse check — what do we have, what changed?

# Sharing project state
operatree describe --plain > structure.yaml   # export structure for reporting
operatree summary                             # review before a status meeting

# Investigating a specific project
operatree describe -d /path/to/project
operatree summary -d /path/to/project
```

Together these three commands give you a complete picture of any project in
seconds — without opening a file manager, without browsing directories, and
without querying a database.

---

_Next: Section 8 — Scripting and Automation_
