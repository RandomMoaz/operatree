# OperaTree Architecture Guide

> **For Contributors:** This document explains OperaTree's design, codebase structure, and how all components work together. Use this to understand the project deeply and make informed contributions.

## Table of Contents

1. [Philosophy & Design Principles](#philosophy--design-principles)
2. [High-Level Architecture](#high-level-architecture)
3. [Package Organization](#package-organization)
4. [Data Flow & Request Lifecycle](#data-flow--request-lifecycle)
5. [Path Resolution & Portability](#path-resolution--portability)
6. [Core Concepts](#core-concepts)
7. [Package Deep Dives](#package-deep-dives)
8. [Subject System Architecture](#subject-system-architecture)
9. [UUID-Based Subject Identification](#uuid-based-subject-identification)
10. [Non-Interactive Mode & Scripting](#non-interactive-mode--scripting)
11. [Template System](#template-system)
12. [Adding New Features](#adding-new-features)
13. [Common Patterns](#common-patterns)
14. [Testing Strategy](#testing-strategy)
15. [Troubleshooting Guide](#troubleshooting-guide)

---

## Philosophy & Design Principles

OperaTree is built on three foundational pillars:

### 1. **Filesystem-First**

- The filesystem is the **single source of truth**
- No database, no external dependencies for core functionality
- All data lives in YAML files under project directories
- Users own their data completely

### 2. **CLI is Just an Interface**

- The CLI is a convenience layer, not a requirement
- Your data remains valid and accessible even if the CLI breaks
- Users can manipulate files directly with standard Unix tools
- This ensures data longevity beyond the tool's lifetime

### 3. **Metadata Separation**

- Each subject (Event, Task, Topic, Objective, DataSource) has a `METADATA.yml` file
- Metadata is searchable, filterable, and machine-readable
- Content can live in the same directory alongside metadata
- Users can edit metadata with their preferred editor

---

## High-Level Architecture

### Layered Design

```
┌──────────────────────────────────────────────────┐
│          CLI Layer (cmd/)                        │
│  (Commands: add, find, create, sync, etc.)       │
└──────────────────┬───────────────────────────────┘
                   │
┌──────────────────▼───────────────────────────────┐
│     Business Logic Layer (pkg/ + internal/)      │
│  pkg/: project, subject, module, config          │
│  internal/: filesystem, metadata, activitylog    │
└──────────────────┬───────────────────────────────┘
                   │
┌──────────────────▼───────────────────────────────┐
│       Persistence Layer (internal/)              │
│  filesystem I/O, YAML serialization              │
└──────────────────┬───────────────────────────────┘
                   │
┌──────────────────▼───────────────────────────────┐
│            Operating System                      │
│         (Filesystem, File I/O)                   │
└──────────────────────────────────────────────────┘
```

### Component Relationships

```
User Command (e.g., "operatree add event --name 'Cairo Visit'")
        │
        ├─→ cmd/add.go (CLI command handler)
        │   ├─ Parse "event" argument (lowercase CLI input)
        │   ├─ Convert to SubjectType constant (UPPERCASE: "EVENT")
        │   ├─ Resolve project directory (-d flag)
        │   └─ Call pkg/project.NewSubject()
        │
        ├─→ pkg/project/new_subject.go
        │   ├─ Validate subject type against SubjectModuleMap
        │   ├─ Find target module recursively
        │   ├─ Create subject instance with all supplied fields
        │   ├─ Run through subject factory
        │   ├─ Write to disk (creates dirs, files, metadata)
        │   ├─ Update project metadata
        │   └─ Log to activity.log
        │
        └─→ internal/filesystem/ + pkg/subject/
            └─→ Persist to disk (directories, files, YAML)
```

---

## Package Organization

### Directory Structure & Layer Classification

```
operatree/
├── cmd/                    # CLI Layer (Cobra commands, 16+ files)
│   ├── root.go            # Cobra setup, global flags, project resolution
│   ├── add.go             # Create new subject (DYNAMIC subject type loading)
│   ├── find.go            # Search subjects (interactive & non-interactive)
│   ├── create.go          # Create new project
│   ├── edit.go            # Edit subject metadata
│   ├── rename.go          # Rename a subject
│   ├── archive.go         # Archive a subject
│   ├── sync.go            # Sync metadata from disk
│   ├── summary.go         # Project statistics
│   ├── describe.go        # Project structure
│   ├── explain.go         # Directory philosophy guide
│   ├── open.go            # Open subject in file manager
│   ├── goto.go            # Jump to tracked project
│   ├── show.go            # Display config/templates/tracked projects
│   ├── track.go           # Add project to tracked list
│   ├── untrack.go         # Remove project from tracked list
│   ├── use.go             # Set default project
│   ├── utilities.go       # Path resolution helpers
│   └── version.go         # Version info
│
├── pkg/                    # PUBLIC API - Business Logic Layer
│   ├── project/           # Project management & orchestration
│   │   ├── types.go            # Project struct, SubjectModuleMap
│   │   ├── new_subject.go      # Create new subject
│   │   ├── load.go             # Load project from disk
│   │   ├── hydrate.go          # Path hydration
│   │   ├── bootstrap.go        # Create new project structure
│   │   ├── factory.go          # Build project from template
│   │   ├── sync.go             # Sync metadata from disk
│   │   ├── list.go             # List all subjects
│   │   ├── find.go             # Search subjects
│   │   ├── project.go          # Project methods (Name, Dir, Archive, etc.)
│   │   ├── archive.go          # Archive subjects
│   │   ├── summary.go          # Project statistics
│   │   └── describe.go         # Pretty-print project
│   │
│   ├── subject/           # Subject types & operations
│   │   ├── types.go            # SubjectType constants, SubDirs, Files maps
│   │   ├── subject.go          # Subject methods (WriteToDisk, Edit, etc.)
│   │   ├── factory.go          # Subject creation factory
│   │   ├── interactive.go      # Interactive CLI prompts
│   │   └── name_factory.go     # Name generation logic
│   │
│   ├── module/            # Module (directory) structure
│   │   ├── types.go            # ModuleType constants, prefixes
│   │   └── module.go           # Module methods (MkDir, Bootstrap, etc.)
│   │
│   └── config/            # Configuration management
│       ├── types.go            # Config struct
│       ├── config.go           # Load/Save config
│       ├── find.go             # Fuzzy find projects
│       └── ...
│
├── internal/              # PRIVATE - NOT exported
│   ├── filesystem/        # File I/O operations (only internal used)
│   │   └── filesystem.go       # All filesystem operations
│   │
│   ├── activitylog/       # Audit trail (only internal used)
│   │   └── activitylog.go      # Action logging
│   │
│   ├── metadata/          # Metadata utilities
│   │   └── metadata.go         # Name formatting, etc.
│   │
│   ├── help/              # Embedded help files
│   │   ├── help.go             # Embedding setup
│   │   └── dir_struct.md       # Directory structure documentation
│   │
│   └── ui/                # Terminal UI formatting
│       └── ui.go               # ANSI colors and formatting
│
├── main.go                # Entry point
├── go.mod                 # Dependencies
├── go.sum                 # Dependency checksums
├── Makefile               # Build configuration
├── ARCHITECTURE.md        # This file
└── README.md              # User documentation
```

### Dependency Graph - CORRECTED

```
cmd/ (depends on)
  ├─→ pkg/project/
  ├─→ pkg/subject/
  ├─→ pkg/config/
  ├─→ internal/activitylog/
  └─→ internal/ui/

pkg/project/ (depends on)
  ├─→ pkg/module/
  ├─→ pkg/subject/
  ├─→ internal/filesystem/
  ├─→ internal/activitylog/
  └─→ internal/metadata/

pkg/subject/ (depends on)
  ├─→ internal/metadata/
  └─→ internal/filesystem/

pkg/module/ (depends on)
  └─→ internal/filesystem/

pkg/config/ (depends on)
  └─→ [Standard library + third-party fuzzyfinder]

internal/filesystem/ (depends on)
  └─→ [Standard library only]

internal/activitylog/ (depends on)
  └─→ [Standard library only]

internal/metadata/ (depends on)
  └─→ [Standard library only]

internal/ui/ (depends on)
  └─→ [Standard library only]
```

### Package Layer Classification

**`pkg/` — PUBLIC API**
- Meant to be imported by external code
- Contains core domain models and operations
- Exported types and functions
- **Includes:** project, subject, module, config

**`internal/` — PRIVATE IMPLEMENTATION**
- Not meant for external use
- Go compiler enforces this at build time
- Shared utilities and low-level operations
- **Includes:** filesystem, activitylog, metadata, help, ui

**`cmd/` — CLI COMMANDS**
- Command-line interface layer
- Depends on pkg/ and internal/
- Not importable (though technically not enforced)

---

## Data Flow & Request Lifecycle

### Example: Creating a New Event (Non-Interactive)

```
User Input: operatree add event --name "Cairo Visit" --date "2026-05-22" --location Cairo -d ~/myproject
│
├─→ cmd/root.go :: Execute()
│   └─→ Cobra parses flags and routes to addCmd
│
├─→ cmd/root.go :: resolveProjectDir() (PreRun hook)
│   ├─ Checks -d flag → actDir = ~/myproject
│   └─ Converts "." to absolute path if needed
│
├─→ cmd/add.go :: newSubject()
│   ├─ Get argument: "event" (lowercase from CLI)
│   ├─ Convert to uppercase: "event" → "EVENT"
│   ├─ Create SubjectType constant: subject.SubjectType("EVENT")
│   ├─ Build Subject struct with CLI flags:
│   │  Subject{
│   │    Name: "Cairo Visit",
│   │    Date: "2026-05-22",
│   │    Location: "Cairo",
│   │    Tags: ..., Notes: ..., etc.
│   │  }
│   ├─ Call pkg/project.Load(actDir)
│   └─ Call pkg/project.NewSubject(&p, ns, SubjectEvent)
│
├─→ pkg/project/new_subject.go :: NewSubject()
│   ├─ Get all existing subjects (for name collision detection)
│   ├─ Validate subject type exists in SubjectModuleMap
│   ├─ Map SubjectEvent to ModuleEvents
│   ├─ Recursively search project.Modules for ModuleEvents
│   │
│   ├─ Create initial subject struct with all fields populated
│   │
│   ├─ Call pkg/subject.SubjectFactory(initialSubject, modulePath, existSubjects)
│   │  └─→ Enters silent mode (name provided)
│   │
│   └─→ Call s.WriteToDisk()
│       └─→ pkg/subject/subject.go
│           ├─ Create subject directory: ~/myproject/01_EVENTS/2026-05-22-cairo-visit/
│           ├─ Create subdirs: 01_AGENDA, 02_MEDIA, 03_NOTES, 04_DOCUMENTS, 05_OUTCOMES
│           ├─ Create default files: (none for Events)
│           └─ Write METADATA.yml with subject data (including UUID)
│
├─→ Update project metadata
│   ├─ Append subject to module.Subjects[]
│   ├─ Write project METADATA.yml
│   └─ pkg/project/hydrate.go :: hydratePath() paths are recalculated
│
├─→ Log action
│   └─→ internal/activitylog/activitylog.go
│       ├─ Build entry: timestamp, action=CREATE, type=EVENT, name="2026-05-22-cairo-visit"
│       ├─ Get user/hostname info
│       └─ Append to activity.log in project root
│
└─→ Output confirmation with UUID
    └─→ "EVENT created: 2026-05-22-cairo-visit"
```

### Subject Type Conversion Flow

**Key Detail:** Subject types have **two representations** and use **dynamic loading**:

```
CLI Argument          Internal Constant          Storage Module Type
─────────────         ─────────────────          ───────────────────
"event"    (lower)  → SubjectEvent("EVENT")     → ModuleEvents
"task"     (lower)  → SubjectTask("TASK")       → ModuleTasks
"topic"    (lower)  → SubjectTopic("TOPIC")     → ModuleTopics
"objective"(lower)  → SubjectObjective("OBJ")   → ModuleObjectives
"datasource"(lower) → SubjectDataSource("DS")   → ModuleDataSources
```

**Dynamic Loading in `cmd/add.go`:**

```go
func init() {
    // Build completion slice DYNAMICALLY from SubjectModuleMap
    // This means adding a new subject type automatically updates the CLI!
    for k := range project.SubjectModuleMap {
        sn := strings.ToLower(string(k))
        validSubjects = append(validSubjects, sn)
    }

    addCmd = &cobra.Command{
        Use:       fmt.Sprintf("add [%s]", strings.Join(validSubjects, " | ")),
        ValidArgs: validSubjects,  // Dynamically populated
        Args:      cobra.MatchAll(cobra.OnlyValidArgs, cobra.ExactArgs(1)),
        Run:       newSubject,
    }
}

func newSubject(cmd *cobra.Command, args []string) {
    a := args[0]                              // "event" from CLI
    st := strings.ToUpper(a)                  // "EVENT"
    
    // Convert to SubjectType constant and pass complete subject with all fields
    if err := project.NewSubject(&p, ns, subject.SubjectType(st)); err != nil {
        log.Fatal(err)
    }
}
```

**Mapping Logic in `pkg/project/types.go`:**

```go
var SubjectModuleMap = map[subject.SubjectType]module.ModuleType{
    subject.SubjectEvent:      module.ModuleEvents,
    subject.SubjectTask:       module.ModuleTasks,
    subject.SubjectTopic:      module.ModuleTopics,
    subject.SubjectObjective:  module.ModuleObjectives,
    subject.SubjectDataSource: module.ModuleDataSources,
}
```

**Key Advantage:** When you add a new subject type, the CLI command automatically recognizes it without code changes to the add command!

---

## Path Resolution & Portability

### Universal `-d` Project Directory Flag

All OperaTree commands support the `-d` (or `--dest`) flag, which specifies the project directory to operate on:

- If `-d` is passed, the specified directory is used.
- If not, OperaTree's standard resolution applies:
  1. If the current directory is a project (contains `METADATA.yml`), it is used.
  2. If a default project is set in the config, it is used.
  3. If neither, a descriptive error is raised.

**Example:**

```bash
operatree add event -d ~/work/reports/sales-2026
```

### Path Hydration Mechanism

**Critical Concept:** Paths are **never persisted**; they're **hydrated at runtime**.

When a project is loaded, `pkg/project/hydrate.go` runs:

```go
func hydratePath(projectBaseDir string, p *Project) {
    p.absDir = projectBaseDir  // Set project's absolute path
    for i, m := range p.Modules {
        p.Modules[i].AbsPath = path.Join(projectBaseDir, m.Name)
        hydrateModule(&p.Modules[i])  // Recursively hydrate modules
    }
}

func hydrateModule(m *module.Module) {
    // Set each subject's absolute directory path
    for i, s := range m.Subjects {
        m.Subjects[i].DirName = path.Join(m.AbsPath, s.Name)
    }
    
    // Recurse into submodules (e.g., Tasks under ProjectManagement)
    for i, sm := range m.Modules {
        m.Modules[i].AbsPath = path.Join(m.AbsPath, sm.Name)
        hydrateModule(&m.Modules[i])
    }
}
```

**Workflow:**

1. CLI parses `-d` flag (or uses default project)
2. `project.Load(actDir)` reads METADATA.yml
3. `hydratePath(actDir, &project)` calculates absolute paths
4. All operations use hydrated paths—never persisted
5. When project moves, no config changes needed

### Why This Matters

- Projects can be moved, copied, or synced across filesystems without breaking
- Collaborators can use different base directories & everything works
- Config/backups are clean, lightweight, and future-proof
- Your data always belongs to you; location is context, not identity

---

## Core Concepts

### 1. **Projects**

- **What:** A collection of subjects organized into modules
- **Storage:** `~/projects/myproject/` directory
- **Metadata:** `METADATA.yml` in project root
- **Structure:** Nested modules (dirs) containing subjects (subdirs)

### 2. **Modules**

- **What:** Directories that organize subjects by category
- **Types:** See `pkg/module/types.go` for complete list:
  - `00_ADMIN` — Governance, contacts, templates
  - `01_EVENTS` — Visits, workshops, meetings
  - `02_PROJECT_MANAGEMENT` — Tasks, reports, risks
  - `03_LEGAL` — Contracts, NDAs, compliance
  - `04_RESEARCH` — Topics, objectives
  - `05_ENGINEERING` — Architecture, specs, decisions
  - `06_DATA` — Raw → staging → processed pipeline
  - `05_TASKS` (nested under PM) — Operational tasks
  - `07_TOPICS` (nested under Research) — Knowledge domains
  - `08_OBJECTIVES` (nested under Research) — Research goals
  - `97_MEDIA_LIBRARY` — Shared reusable assets
  - `98_DELIVERABLES` — Final external outputs
  - `99_ARCHIVE` — Historical storage
- **Nesting:** Some modules contain submodules (e.g., TASKS under PROJECT_MANAGEMENT)
- **Storage:** `Module.Subjects[]` contains direct subjects; `Module.Modules[]` contains nested modules

### 3. **Subjects**

- **What:** Trackable units of work or knowledge
- **Types:** EVENT, TASK, TOPIC, OBJECTIVE, DATASOURCE (see `pkg/subject/types.go`)
- **Storage:** Each subject is a directory with `METADATA.yml`
- **Structure:** Subjects auto-create subdirectories and default files based on their type
- **Identification:** Each subject has a unique UUID (v7, sortable by timestamp)

### 4. **Metadata**

- **What:** YAML file containing subject/project properties
- **Location:** `subject-dir/METADATA.yml` or `project-dir/METADATA.yml`
- **Format:** YAML (human-readable, version-control friendly)
- **Editability:** Users can edit directly; `sync` command updates project index
- **Auto-Sync:** `edit` command automatically syncs after editor closes

### 5. **Activity Log**

- **What:** Append-only audit trail of all changes
- **Location:** `project-root/activity.log`
- **Format:** Tab-separated values
- **Entries:** Timestamp, action (CREATE/EDIT/ARCHIVE/RENAME), type, name, user@host, version

**Example entry:**
```
2026-05-20T10:08:39Z	CREATE    	EVENT        	"2026-05-22-cairo-visit"	hany@optiplex7040	v0.1.0
2026-05-20T14:05:03Z	RENAME    	TASK         	"Prepare Report"        	hany@optiplex7040	v0.1.2
```

---

## Package Deep Dives

### `cmd/` — CLI Layer

**Purpose:** Command-line interface, argument parsing, user interaction

**Command Refactoring Summary:**

| Old Name | New Name | Purpose |
|----------|----------|---------|
| `new` | `add` | Create new subject |
| `bootstrap` | `create` | Create new project |
| `metadata` | `edit` | Edit subject metadata |
| `default` | `use` | Set default project |
| `jump` | `goto` | Open tracked project |

**Key Files:**

- **`root.go`** — Cobra setup, global flags, config loading, path resolution
- **`add.go`** — Create new subject (DYNAMIC subject type loading from `pkg/project.SubjectModuleMap`)
- **`find.go`** — Search subjects (interactive & non-interactive)
- **`create.go`** — Create new project (replaces `bootstrap`)
- **`edit.go`** — Edit subject metadata (auto-syncs after save)
- **`rename.go`** — Rename a subject with cross-reference updates
- **`archive.go`** — Archive a subject to 99_ARCHIVE
- **`sync.go`** — Sync project metadata from disk
- **`show.go`** — Display config/templates/tracked projects
- **`track.go`/`untrack.go`** — Manage tracked projects
- **`use.go`** — Set default project (replaces `default`)
- **`goto.go`** — Jump to tracked project (replaces `jump`)

**Import Pattern:**

```go
import (
    "github.com/hanymamdouh82/operatree/pkg/project"
    "github.com/hanymamdouh82/operatree/pkg/subject"
    "github.com/hanymamdouh82/operatree/pkg/config"
)
```

---

### `pkg/project/` — Project Management (PUBLIC API)

**Purpose:** High-level project operations, orchestration, template application

**Location:** `operatree/pkg/project/`

**Key Files:**

- **`types.go`** — Type definitions
  - `Project` struct with `absDir` (hydrated absolute path)
  - `SubjectModuleMap` — Maps each subject type to its storage module
  - Constants: `METADATA_FILE = "METADATA.yml"`, `ARCHIVED_DEST = "closed_tasks"`

- **`new_subject.go`** — Create new subject
  - `NewSubject()` — Main orchestration function
  - `findModule()` — Recursively search for target module by type

- **`load.go`** — Load project from disk
  - `Load(path)` — Reads METADATA.yml and calls `hydratePath()`

- **`hydrate.go`** — Path hydration at runtime
  - `hydratePath()` — Set absolute paths on project and all modules

- **`bootstrap.go`** — Create new project structure
  - `Bootstrap()` — Load template, create project, create directories

- **`factory.go`** — Build project from template
  - `Factory()` — Convert template to project structure

- **`sync.go`** — Sync metadata from disk
  - `Sync()` — Read all subject METADATA.yml files and update project

- **`list.go`** — List all subjects
  - `ListSubjects()` — Flatten project tree into list, optionally filtered by type

- **`find.go`** — Search subjects
  - `FindSubject()` — Fuzzy search (interactive, single result)
  - `FindSubjectsSilent()` — Fuzzy search (non-interactive, multiple results)

- **`project.go`** — Project methods
  - `ProjectName()` — Get project name
  - `ProjectDir()` — Get absolute project path
  - `WriteMetadata()` — Persist project METADATA.yml

- **`archive.go`** — Archive subjects
  - `Archive()` — Move subject to 99_ARCHIVE module

- **`summary.go`** — Project statistics
  - `Summary()` — Display subject counts by type

**Core Type:**

```go
type Project struct {
    Name     string          `yaml:"name"`
    Template string          `yaml:"template"`
    absDir   string          `yaml:"-"` // project absolute directory, hydrated during load
    Tags     []string        `yaml:"tags"`
    Modules  []module.Module `yaml:"modules"`
}
```

---

### `pkg/subject/` — Subject Types & Operations (PUBLIC API)

**Location:** `operatree/pkg/subject/`

**Key Files:**

- **`types.go`** — Type definitions and configuration maps
  - Subject type constants: `SubjectEvent`, `SubjectTask`, `SubjectTopic`, `SubjectObjective`, `SubjectDataSource`
  - Constant: `METADATA_FILE = "METADATA.yml"`
  - `SubDirs` map — Default subdirectories for each subject type
  - `Files` map — Default files for each subject type

- **`subject.go`** — Subject operations
  - `WriteToDisk()` — Orchestration: creates dir → subdirs → files → metadata

- **`factory.go`** — Subject creation factory
  - `SubjectFactory()` — Main factory function
  - `silent()` — Factory for non-interactive mode
  - `interactive()` — Factory for interactive mode

- **`interactive.go`** — Interactive CLI prompts
  - `interactiveCLI()` — Prompt user for subject properties

- **`name_factory.go`** — Name generation logic
  - `nameFactory()` — Generate directory name by type

**Subject Struct:**

```go
type Subject struct {
    UUID               string      `yaml:"uuid"`  // Unique identifier (v7 UUID)
    Type               SubjectType `yaml:"type"`
    Name               string      `yaml:"name"`
    DirName            string      `yaml:"-"`  // Not persisted, hydrated at load
    SubDirs            []string    `yaml:"subDirs"`
    Files              []string    `yaml:"-"` // Not persisted, used for creation only
    Date               string      `yaml:"date"`
    Tags               []string    `yaml:"tags"`
    Notes              string      `yaml:"notes"`
    // Type-specific fields
    Participants       []string `yaml:"participants,omitempty"`
    Location           string   `yaml:"location,omitempty"`
    Owner              string   `yaml:"owner,omitempty"`
    Status             string   `yaml:"status,omitempty"`
    RelatedObjective   string   `yaml:"related_objective,omitempty"`
    RelatedEvents      []string `yaml:"related_events,omitempty"`
    Outputs            []string `yaml:"outputs,omitempty"`
    Source             string   `yaml:"source,omitempty"`          // DataSource
    SourceLink         string   `yaml:"sourceLink,omitempty"`      // DataSource
    SourceObjective    string   `yaml:"sourceObjective,omitempty"` // DataSource
    SourceDataSize     string   `yaml:"sourceDataSize,omitempty"`  // DataSource
}
```

---

### `pkg/module/` — Module Structure (PUBLIC API)

**Location:** `operatree/pkg/module/`

**Purpose:** Directory structure organization, module types, filesystem bootstrap

**Key Files:**

- **`types.go`** — Type definitions and prefix mapping
  - `ModuleType` constants (uppercase): `ModuleAdmin = "ADMIN"`, etc.
  - `ModuleDirPrefixMap` — Maps module type to directory prefix (00-99)

- **`module.go`** — Module operations
  - `MkDir()` — Create module directory
  - `MkSubDirs()` — Create module's default subdirectories
  - `Bootstrap()` — Recursive: creates dir → subdirs → nested modules

**Module Types & Prefixes:**

```
ModuleAdmin              → "00"
ModuleEvents            → "01"
ModuleProjectManagement → "02"
ModuleLegal             → "03"
ModuleResearch          → "04"
ModuleEngineering       → "05"
ModuleData              → "06"
ModuleTasks             → "05" (nested under PROJECT_MANAGEMENT)
ModuleTopics            → "07" (nested under RESEARCH)
ModuleObjectives        → "08" (nested under RESEARCH)
ModuleDataSources       → "13" (nested under DATA)
ModuleMediaLibrary      → "97"
ModuleDeliverables      → "98"
ModuleArchive           → "99"
```

---

### `pkg/config/` — Configuration Management (PUBLIC API)

**Location:** `operatree/pkg/config/`

**Purpose:** User configuration, project tracking, persistence

**Key Files:**

- **`types.go`** — Type definitions
  - `Config` struct with projects list and default project
  - `Project` struct with name, path, template

- **`config.go`** — Configuration operations
  - `Load()` — Load config from YAML
  - `Save()` — Persist config to YAML
  - `AddProject()` — Register new project
  - `SetDefaultProject()` — Set default project

- **`find.go`** — Fuzzy find projects
  - `Find()` — Interactive project search

---

### `internal/filesystem/` — File I/O (PRIVATE)

**Location:** `operatree/internal/filesystem/`

**Purpose:** All filesystem operations encapsulated here

**Key Operations:**

- `CheckDirExists(path)` — Check if directory exists
- `CreateDir(path)` — Create directory (fails if exists)
- `ReadFile(path)` — Read file contents
- `StructToFile(struct, path)` — Marshal Go struct to YAML file
- `FileToStruct(struct, path)` — Unmarshal YAML file to Go struct
- `TextToMDFile(text, path)` — Write text to file
- `Archive(src, dest)` — Move file/directory to archive
- `RenameDir(src, dest)` — Rename directory

**Design Philosophy:** Single responsibility — all filesystem I/O goes through this package:
- Easy to mock for testing
- Centralized error handling
- Future enhancement opportunity (permissions, backups, etc.)

---

### `internal/activitylog/` — Audit Trail (PRIVATE)

**Location:** `operatree/internal/activitylog/`

**Purpose:** Log all user actions for audit and undo

**Key Types & Constants:**

```go
type Action string

const (
    ActionCreate  Action = "CREATE"
    ActionEdit    Action = "EDIT"
    ActionArchive Action = "ARCHIVE"
    ActionRename  Action = "RENAME"
    ActionDelete  Action = "DELETE"  // Planned
)
```

**Log Format (Tab-Separated):**

```
timestamp                 action    type        name                      user@host            version
2026-05-20T10:08:39Z     CREATE    EVENT       "2026-05-22-cairo-visit"  hany@optiplex7040    v0.1.0
2026-05-20T14:05:03Z     RENAME    TASK        "Prepare Report"          hany@optiplex7040    v0.1.2
```

**Key Operations:**

- `Log(projectRoot, action, subjectType, subjectName)` — Record action

**Design:** Append-only, pipe-friendly for Unix integration

---

### `internal/metadata/` — Metadata Utilities (PRIVATE)

**Location:** `operatree/internal/metadata/`

**Purpose:** Metadata parsing, name formatting

**Key Operations:**

- `FormatName(name)` — Sanitize and hyphenate names
- YAML marshaling/unmarshaling

---

### `internal/help/` — Embedded Help (PRIVATE)

**Location:** `operatree/internal/help/`

**Purpose:** Embedded documentation files

**Files:**

- `dir_struct.md` — Directory structure philosophy and guide
- `help.go` — Embedding setup with `go:embed`

---

### `internal/ui/` — Terminal UI Formatting (PRIVATE)

**Location:** `operatree/internal/ui/`

**Purpose:** Pretty-printing, colored output, terminal aesthetics

**ANSI Color Constants:**

```go
AnsiReset   = "\033[0m"
AnsiBold    = "\033[1m"
AnsiDim     = "\033[2m"
AnsiItalic  = "\033[3m"
AnsiPurple  = "\033[38;5;141m"
AnsiYellow  = "\033[38;5;221m"
AnsiGray    = "\033[38;5;244m"
AnsiGreen   = "\033[38;5;114m"
```

---

## Subject System Architecture

### Five Subject Types

The system now supports five types of subjects:

#### 1. **Event**
- **Module:** `01_EVENTS`
- **Purpose:** Record project activities (meetings, visits, workshops)
- **Key Fields:** location, participants, date
- **Subdirectories:** 01_AGENDA, 02_MEDIA, 03_NOTES, 04_DOCUMENTS, 05_OUTCOMES
- **Directory Naming:** `YYYY-MM-DD-hyphenated-name`

#### 2. **Task**
- **Module:** `02_PROJECT_MANAGEMENT/05_TASKS`
- **Purpose:** Unit of work with lifecycle
- **Key Fields:** owner, status, related_events, outputs
- **Subdirectories:** 01_INPUTS, 02_WORKING, 03_REVIEW, 04_FINAL
- **Directory Naming:** `YYYY-MM-DD-hyphenated-name`

#### 3. **Topic**
- **Module:** `04_RESEARCH/07_TOPICS`
- **Purpose:** Knowledge concept or domain area
- **Key Fields:** related_objective, tags, notes
- **Default Files:** overview.md, notes.md
- **Subdirectories:** 01_DIAGRAMS, 02_ATTACHMENTS
- **Directory Naming:** `hyphenated-name`

#### 4. **Objective**
- **Module:** `04_RESEARCH/08_OBJECTIVES`
- **Purpose:** Goal driving research and decisions
- **Key Fields:** status, outputs, tags
- **Default Files:** definitions.md, findings.md, strategy.md
- **Subdirectories:** 01_NOTES, 02_DISCUSSIONS, 03_ATTACHMENTS
- **Directory Naming:** `hyphenated-name`

#### 5. **DataSource** (New)
- **Module:** `06_DATA/13_DATASOURCES`
- **Purpose:** External dataset or data feed
- **Key Fields:** source, source_link, source_objective, source_datasize
- **Default Files:** source.md
- **Directory Naming:** `hyphenated-name`
- **Example:** "sensor-readings-2025", "kaggle-datasets-ml"

### Subject Lifecycle

```
add/interactive → interactive form
         ↓
add/flags → populated Subject struct
         ↓
SubjectFactory → validate, generate name, create UUID
         ↓
WriteToDisk → create dirs, files, metadata
         ↓
edit → open METADATA.yml in editor → auto-sync
         ↓
find → search by term/type/UUID
         ↓
rename → update dir, metadata, cross-references
         ↓
archive → move to 99_ARCHIVE
```

---

## UUID-Based Subject Identification

### Why UUIDs?

1. **Stable Identity:** Subjects can be renamed; UUID never changes
2. **Scriptable Operations:** Use `--uuid` flag to target subjects without interactive prompts
3. **Bulk Operations:** Combine UUID discovery with bash/awk for pipeline automation
4. **Portability:** UUID-based identity works across filesystem moves

### UUID System

- **Type:** v7 (sortable by timestamp)
- **Persistence:** Written to METADATA.yml `uuid` field
- **Generation:** Happens automatically during subject factory

**Example YAML:**

```yaml
uuid: a1b2c3d4-e5f6-7g8h-i9j0-k1l2m3n4o5p6
type: EVENT
name: 2026-05-22-cairo-visit
date: "2026-05-22"
location: Cairo
```

### Scripted Operations with UUIDs

**Example: Rename subject by UUID without interactive prompt**

```bash
operatree rename --uuid a1b2c3d4-e5f6-7g8h-i9j0-k1l2m3n4o5p6 --new-name "New Name"
```

**Example: Archive all done tasks via pipeline**

```bash
operatree find --term done --type task --plain \
  | grep uuid \
  | awk '{print $2}' \
  | xargs -I{} operatree archive --uuid {}
```

---

## Non-Interactive Mode & Scripting

### Two Modes of Operation

#### Interactive Mode
- User sees prompts
- Finds subjects via fuzzy search
- Ideal for daily interactive use
- Example: `operatree add event`

#### Non-Interactive Mode
- Controlled by flags
- No prompts or finders
- Ideal for scripts, cron jobs, pipelines
- Example: `operatree add event --name "Cairo Visit" --date 2026-05-22`

### Subject Creation with Flags

All flags are available for non-interactive subject creation:

**Common Flags (all types):**
```bash
--name          Subject name
--date          Date
--notes         Notes
--tags          Comma-delimited tags
```

**Event-Specific:**
```bash
--location      Location
--participants  Comma-delimited participant names
```

**Task-Specific:**
```bash
--owner              Person responsible
--status             Status (e.g., "active", "blocked", "done")
--related-events     Comma-delimited related event names
--outputs            Comma-delimited output names
```

**Topic-Specific:**
```bash
--related-objective  Related objective name
```

**DataSource-Specific:**
```bash
--source             Data origin (e.g., "Kaggle", "Internal API")
--source-link        URL or path to data
--source-objective   Related objective
--source-datasize    Dataset size/volume
```

**Examples:**

```bash
# Non-interactive event creation
operatree add event \
  --name "Cairo Factory Visit" \
  --date 2026-06-01 \
  --location Cairo \
  --participants "Alex,Sara,Omar" \
  --tags "factory,inspection" \
  --notes "Production line assessment"

# Non-interactive task creation
operatree add task \
  --name "Prepare Report" \
  --date 2026-06-01 \
  --owner Alex \
  --status active \
  --related-events "Cairo Factory Visit" \
  --outputs "Report v1.0" \
  --tags "report"
```

### Metadata Auto-Sync After Edit

When a user edits metadata via `operatree edit`:

1. File opens in editor
2. User makes changes and saves
3. Editor closes
4. OperaTree automatically runs `sync` (no user action needed)
5. Project metadata index is updated

**Before:** Manual `operatree sync` was required
**Now:** Automatic, reducing friction

---

## Adding New Features

### Adding a New Subject Type

To add a new subject type (e.g., "REPORT"):

**Step 1:** Add constant in `pkg/subject/types.go`

```go
const SubjectReport SubjectType = "REPORT"
```

**Step 2:** Add to `SubjectModuleMap` in `pkg/project/types.go`

```go
var SubjectModuleMap = map[subject.SubjectType]module.ModuleType{
    // ... existing entries ...
    subject.SubjectReport: module.ModuleReports,  // Module must exist
}
```

**Step 3:** Add module constant in `pkg/module/types.go`

```go
const ModuleReports ModuleType = "REPORTS"
```

**Step 4:** Add prefix mapping in `pkg/module/types.go`

```go
ModuleDirPrefixMap[ModuleReports] = "08"  // Or appropriate prefix
```

**Step 5:** (Optional) Configure default dirs/files in `pkg/subject/types.go`

```go
SubDirs[SubjectReport] = []string{...}
Files[SubjectReport] = []string{...}
```

**That's it!** The CLI command automatically recognizes the new type:

```bash
operatree add report --name "Q2 Results" --date 2026-06-30
# Works immediately without changes to cmd/add.go!
```

---

## Common Patterns

### Pattern 1: Resolve Project + Load + Operate

```go
// In cmd handler
func myCommand(cmd *cobra.Command, args []string) {
    // 1. Project already resolved by PreRun hook
    
    // 2. Load project (hydrates paths)
    p, err := project.Load(actDir)
    if err != nil { log.Fatal(err) }
    
    // 3. Operate on project
    if err := project.SomeOperation(&p, args); err != nil {
        log.Fatal(err)
    }
}
```

### Pattern 2: Find Subject → Operate → Update Metadata

```go
// In pkg/project or cmd
func OperateOnSubject(p *project.Project, searchTerm string) error {
    // 1. Find subject
    s, err := project.FindSubject(p, "", searchTerm)
    if err != nil { return err }
    
    // 2. Operate on subject
    s.DoSomething()
    
    // 3. Update project metadata
    return p.WriteMetadata()
}
```

### Pattern 3: Walk All Subjects Recursively

```go
func walkSubjects(modules []module.Module, callback func(*subject.Subject)) {
    for i, m := range modules {
        // Process subjects at this level
        for j := range m.Subjects {
            callback(&modules[i].Subjects[j])
        }
        
        // Recurse into nested modules
        walkSubjects(m.Modules, callback)
    }
}
```

---

## Testing Strategy

### Test Organization

```
operatree/
├── cmd/
│   ├── find_test.go
│   ├── add_test.go
│   └── ...
├── pkg/
│   ├── project/
│   │   ├── bootstrap_test.go
│   │   ├── find_test.go
│   │   └── ...
│   ├── subject/
│   │   ├── factory_test.go
│   │   └── ...
│   └── ...
├── internal/
│   ├── filesystem/
│   │   └── filesystem_test.go
│   └── ...
```

### Test Patterns

**Unit Tests:** Test individual functions with mocked dependencies

```go
func TestFindSubject_Fuzzy(t *testing.T) {
    p := &project.Project{...}
    s, err := project.FindSubject(p, "event", "cairo")
    if err != nil { t.Fatal(err) }
    if s.Name != "cairo-visit" {
        t.Errorf("expected cairo-visit, got %s", s.Name)
    }
}
```

---

## Troubleshooting Guide

### Common Issues

**Issue:** "Module not found in project"
- **Check:** Is the module template included in the project?
- **Check:** Has the project been bootstrapped?
- **Fix:** Re-run project bootstrap or check template

**Issue:** "Cannot find directory"
- **Check:** Are paths being hydrated correctly?
- **Check:** Is the project loaded before use?
- **Fix:** Ensure `project.Load()` is called

**Issue:** "Metadata out of sync"
- **Check:** Were METADATA.yml files edited outside of OperaTree?
- **Check:** Did files arrive via git pull or Syncthing?
- **Fix:** Run `operatree sync` to rebuild index

---

## Key Data Structures

### Project Metadata (METADATA.yml)

```yaml
name: fleetfix
template: consulting
modules:
  - type: EVENTS
    name: 01_EVENTS
    modules: []
    subjects:
      - uuid: a1b2c3d4-e5f6-7g8h-i9j0-k1l2m3n4o5p6
        type: EVENT
        name: 2026-05-22-cairo-visit
        date: "2026-05-22"
        location: Cairo
        tags: [factory, inspection]
        subDirs: [01_AGENDA, 02_MEDIA, 03_NOTES, 04_DOCUMENTS, 05_OUTCOMES]
```

---

**This document is the source of truth for OperaTree architecture. Keep it updated as the codebase evolves!**
