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

- Each subject (Event, Task, Topic, Objective, DataSource) has a `META.yaml` file
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
│       Business Logic Layer (internal/)           │
│  project, subject, module, template, metadata    │
└──────────────────┬───────────────────────────────┘
                   │
┌──────────────────▼───────────────────────────────┐
│       Persistence Layer (internal/)              │
│  filesystem I/O, YAML serialization, config mgmt │
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
        │   └─ Call project.NewSubject()
        │
        ├─→ internal/project/new_subject.go
        │   ├─ Validate subject type against SubjectModuleMap
        │   ├─ Find target module recursively
        │   ├─ Create subject instance with all supplied fields
        │   ├─ Run through subject factory
        │   ├─ Assign UUID to subject (v7 UUID)
        │   ├─ Write to disk (creates dirs, files, metadata)
        │   ├─ Update project metadata
        │   └─ Log to activity.log
        │
        └─→ internal/filesystem/ + internal/subject/
            └─→ Persist to disk (directories, files, YAML)
```

---

## Package Organization

### Top-Level Structure

```
operatree/
├── cmd/                    # CLI commands (16 files)
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
├── internal/              # Business logic (not exported)
│   ├── project/           # Project management & orchestration
│   ├── subject/           # Subject types & operations
│   ├── module/            # Module (directory) structure
│   ├── template/          # Project templates
│   ├── config/            # Configuration management
│   ├── filesystem/        # File I/O operations
│   ├── metadata/          # Metadata utilities
│   ├── activitylog/       # Audit trail
│   ├── runner/            # External command execution
│   ├── help/              # Embedded help files
│   └── ui/                # Terminal UI formatting
│
├── main.go                # Entry point
├── go.mod                 # Dependencies
├── go.sum                 # Dependency checksums
├── Makefile               # Build configuration
├── ARCHITECTURE.md        # This file
└── README.md              # User documentation
```

### Dependency Graph

```
cmd/ (depends on)
  ├─→ internal/project/
  ├─→ internal/subject/
  ├─→ internal/config/
  ├─→ internal/template/
  ├─→ internal/runner/
  └─→ internal/ui/

internal/project/ (depends on)
  ├─→ internal/module/
  ├─→ internal/subject/
  ├─→ internal/template/
  ├─→ internal/filesystem/
  ├─→ internal/activitylog/
  └─→ internal/metadata/

internal/subject/ (depends on)
  ├─→ internal/metadata/
  ├─→ internal/filesystem/
  └─→ internal/runner/

internal/module/ (depends on)
  ├─→ internal/subject/
  └─→ internal/filesystem/

internal/template/ (depends on)
  ├─→ internal/module/
  └─→ internal/project/

internal/filesystem/ (depends on)
  └─→ [Standard library only]
```

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
│   ├─ Load project from actDir
│   └─ Call project.NewSubject(&p, ns, SubjectEvent)
│
├─→ internal/project/new_subject.go :: NewSubject()
│   ├─ Get all existing subjects (for name collision detection)
│   ├─ Validate subject type exists in SubjectModuleMap
│   ├─ Map SubjectEvent to ModuleEvents
│   ├─ Recursively search project.Modules for ModuleEvents
│   │
│   ├─ Create initial subject struct with all fields populated
│   │
│   ├─ Call subject.SubjectFactory(initialSubject, modulePath, existSubjects)
│   │  └─→ Enters silent mode (name provided)
│   │
│   ├─ Call s.SetID() (generates v7 UUID)
│   │
│   └─→ Call s.WriteToDisk()
│       └─→ internal/subject/subject.go
│           ├─ Create subject directory: ~/myproject/01_EVENTS/2026-05-22-cairo-visit/
│           ├─ Create subdirs: 01_AGENDA, 02_MEDIA, 03_NOTES, 04_DOCUMENTS, 05_OUTCOMES
│           ├─ Create default files: (none for Events)
│           └─ Write META.yaml with subject data (including UUID)
│
├─→ Update project metadata
│   ├─ Append subject to module.Subjects[]
│   ├─ Write project META.yaml
│   └─ internal/project/hydrate.go :: hydratePath() paths are recalculated
│
├─→ Log action
│   └─→ internal/activitylog/activitylog.go
│       ├─ Build entry: timestamp, action=CREATE, type=EVENT, name="2026-05-22-cairo-visit"
│       ├─ Get user/hostname info
│       └─ Append to activity.log in project root
│
└─→ Output confirmation with UUID
    └─→ "EVENT created: 2026-05-22-cairo-visit (uuid: a1b2c3d4...)"
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

**Mapping Logic in `internal/project/types.go`:**

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
  1. If the current directory is a project (contains `META.yaml`), it is used.
  2. If a default project is set in the config, it is used.
  3. If neither, a descriptive error is raised.

**Example:**

```bash
operatree add event -d ~/work/reports/sales-2026
```

### Path Hydration Mechanism

**Critical Concept:** Paths are **never persisted**; they're **hydrated at runtime**.

When a project is loaded, `internal/project/hydrate.go` runs:

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
2. `project.Load(actDir)` reads META.yaml
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
- **Metadata:** `META.yaml` in project root
- **Structure:** Nested modules (dirs) containing subjects (subdirs)

### 2. **Modules**

- **What:** Directories that organize subjects by category
- **Types:** See `internal/module/types.go` for complete list:
  - `00_ADMIN` — Governance, contacts, templates
  - `01_EVENTS` — Visits, workshops, meetings
  - `02_PROJECT_MANAGEMENT` — Tasks, reports, risks (contains nested `07_TASKS`)
  - `03_LEGAL` — Contracts, NDAs, compliance
  - `04_RESEARCH` — Topics, objectives (contains nested modules)
  - `05_ENGINEERING` — Architecture, specs, decisions
  - `06_DATA` — Raw → staging → processed pipeline
  - `97_MEDIA_LIBRARY` — Shared reusable assets
  - `98_DELIVERABLES` — Final external outputs
  - `99_ARCHIVE` — Historical storage
- **Nesting:** Some modules contain submodules (e.g., `07_TASKS` under `02_PROJECT_MANAGEMENT`)
- **Storage:** `Module.Subjects[]` contains direct subjects; `Module.Modules[]` contains nested modules

### 3. **Subjects**

- **What:** Trackable units of work or knowledge
- **Types:** EVENT, TASK, TOPIC, OBJECTIVE, DATASOURCE (see `internal/subject/types.go`)
- **Storage:** Each subject is a directory with `META.yaml`
- **Structure:** Subjects auto-create subdirectories and default files based on their type
- **Identification:** Each subject has a unique UUID (v7, sortable by timestamp)

### 4. **Metadata**

- **What:** YAML file containing subject/project properties
- **Location:** `subject-dir/META.yaml` or `project-dir/META.yaml`
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

The following commands were renamed for clarity and consistency:

| Old Name | New Name | Purpose |
|----------|----------|---------|
| `new` | `add` | Create new subject |
| `bootstrap` | `create` | Create new project |
| `metadata` | `edit` | Edit subject metadata |
| `default` | `use` | Set default project |
| `jump` | `goto` | Open tracked project |

**Key Files:**

- **`root.go`** — Cobra setup, global flags, config loading, path resolution
  - Loads configuration at startup
  - Defines global variables: `destDir`, `actDir`, `cfg`, `verbose`
  - Sets up root command

- **`utilities.go`** — Path resolution helpers
  - `resolveProjectDir()` — Resolves `-d` flag for project commands
  - `resolveBaseDir()` — Resolves `-d` flag for base directory commands
  - `resolveProjectDirSkippingConfig()` — Ignores config, only uses explicit flags

- **`add.go`** — Create new subject (replaces old `new.go`)
  - **DYNAMICALLY loads valid subject types** from `project.SubjectModuleMap`
  - Converts CLI argument (lowercase) to SubjectType constant (uppercase)
  - Accepts flags for all subject fields: `--name`, `--date`, `--location`, `--owner`, `--status`, etc.
  - Calls `project.NewSubject()` with populated Subject struct
  - **No static map needed** — automatically picks up new subject types!

- **`create.go`** — Create new project (replaces old `bootstrap.go`)
  - Takes project name and template name
  - Uses `-d` for base directory (where project is created)
  - Calls `project.Bootstrap()`

- **`find.go`** — Search subjects (enhanced)
  - Fuzzy search by type and term (interactive)
  - Non-interactive mode via `--term`, `--type` flags
  - `--plain` flag for raw YAML output
  - Calls `project.FindSubject()` or `project.FindSubjectsSilent()`

- **`edit.go`** — Edit subject metadata (replaces old `metadata.go`)
  - Finds subject, opens `META.yaml` in editor
  - Automatically runs `sync` after editor closes
  - Calls `subject.EditMetadata()`

- **`rename.go`** — Rename a subject (new)
  - Interactive mode: finds subject, prompts for new name
  - Non-interactive mode: `--uuid` + `--new-name` flags
  - Updates subject directory, metadata, and all cross-references
  - Calls `project.RenameSubject()`

- **`archive.go`** — Archive a subject (enhanced)
  - Interactive mode: finds subject via fuzzy search
  - Non-interactive mode: `--uuid` flag
  - Moves subject to `99_ARCHIVE`
  - Calls `project.Archive()`

- **`sync.go`** — Sync project metadata
  - Walks all subjects on disk
  - Updates project in-memory from META.yaml files
  - Calls `project.Sync()`

- **`summary.go`** — Project statistics
  - Displays high-level project overview
  - Subject counts by type

- **`describe.go`** — Project structure
  - Pretty-prints project directory tree
  - `--plain` flag for raw YAML

- **`explain.go`** — Directory philosophy guide (new)
  - Prints full OperaTree directory structure philosophy

- **`open.go`** — Open subject in file manager
  - Finds subject, opens its directory

- **`goto.go`** — Jump to tracked project (replaces old `jump.go`)
  - Opens tracked project root in file manager

- **`use.go`** — Set default project (replaces old `default.go`)
  - Interactively select default project

- **`show.go`** — Display information
  - Shows tracked projects, config, templates, default project
  - No project `-d` flag needed

- **`track.go`** — Add project to tracked list
  - Adds project to config for future default project usage
  - Requires `-d` flag

- **`untrack.go`** — Remove project from tracked list
  - Removes project from config
  - Requires `-d` flag or name argument

**Command Pattern:**

```go
// Typical command handler pattern
func init() {
    cmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
    cmd.PreRun = resolveProjectDir  // Resolve -d before running
    rootCmd.AddCommand(cmd)
}

func commandHandler(cmd *cobra.Command, args []string) {
    // 1. Load project (uses already-resolved actDir)
    p, err := project.Load(actDir)
    if err != nil { log.Fatal(err) }

    // 2. Call business logic
    if err := project.SomeFunction(&p, args); err != nil {
        log.Fatal(err)
    }
}
```

---

### `internal/project/` — Project Management

**Purpose:** High-level project operations, orchestration, template application

**Key Files:**

- **`project.go`** — Project type methods
  - `ProjectName()` — Get project name
  - `ProjectDir()` — Get absolute project path
  - `ProjectBaseDir()` — Get parent directory
  - `Describe()` — Pretty-print project
  - `WriteMetadata()` — Persist project META.yaml
  - `ModuleExists()` — Check if module exists
  - `Archive()` — Archive a subject
  - `RenameSubject()` — Rename a subject

- **`types.go`** — Type definitions
  - `Project` struct with `absDir` (hydrated absolute path)
  - `SubjectModuleMap` — Maps each subject type to its storage module
  - Metadata file constant: `METADATA_FILE = "META.yaml"`

- **`new_subject.go`** — Create new subject
  - `NewSubject()` — Main orchestration function
  - Accepts fully populated Subject struct (from CLI flags)
  - `findModule()` — Recursively search for target module by type
  - Validates subject type, finds module, creates subject, persists to disk
  - Generates UUID and logs action

- **`load.go`** — Load project from disk
  - `Load(path)` — Reads META.yaml and calls `hydratePath()`
  - Calls `backfillUUIDs()` to ensure all subjects have UUIDs
  - Performs automatic migrations on load

- **`hydrate.go`** — Path hydration
  - `hydratePath()` — Set absolute paths on project and all modules
  - `hydrateModule()` — Recursive path hydration for nested modules
  - Sets `AbsPath` on all modules and `DirName` on all subjects

- **`backfill_uuid.go`** — UUID backfill (new)
  - `backfillUUIDs()` — Ensures all subjects have UUIDs
  - Generates missing UUIDs and writes to disk
  - Removes orphaned subjects from metadata

- **`bootstrap.go`** — Create new project structure
  - `Bootstrap()` — Load template, create project, create directories
  - Calls `internal/template/Load()` to get template
  - Calls `internal/project/Factory()` to build project structure
  - Calls `module.Bootstrap()` to create directories

- **`factory.go`** — Build project from template
  - `Factory()` — Convert template to project structure
  - `parseModule()` — Recursively parse template modules to module objects

- **`sync.go`** — Sync metadata from disk
  - `Sync()` — Read all subject META.yaml files and update project
  - `syncModule()` — Recursively sync all subjects in a module

- **`list_subjects.go`** — List all subjects
  - `ListSubjects()` — Flatten project tree into list of all subjects

- **`find_subjects.go`** — Search subjects (enhanced)
  - `FindSubject()` — Fuzzy search, interactive mode (single result)
  - `FindSubjectsSilent()` — Fuzzy search, non-interactive (multiple results)
  - `FindSubjectByID()` — Find subject by UUID (direct, non-interactive)
  - Supports type filtering

- **`find_module.go`** — Find modules (new)
  - `findModule()` — Recursively search for module by type

- **`rename_subject.go`** — Rename subjects (new)
  - `RenameSubject()` — Main orchestration
  - Updates subject directory, metadata, and cross-references
  - `updateReferences()` — Updates related_events and related_objective fields
  - `identifyUUIDsToBeUpdated()` — Find all subjects referencing old name

- **`describe.go`** — Pretty-print project
  - Formatted project output for terminal display

- **`archive.go`** — Archive subjects
  - `Archive()` — Move subject to 99_ARCHIVE module

- **`summary.go`** — Project statistics
  - `Summary()` — Display subject counts by type
  - Shows latest activity date and subject name

**Core Type:**

```go
type Project struct {
    Name     string          // e.g., "myproject"
    Template string          // e.g., "dev"
    absDir   string          // project absolute directory, hydrated at load time
    Tags     []string        // Project-level tags
    Modules  []module.Module // Top-level modules (can contain nested modules)
}
```

**Key Pattern:**

```go
// Create new subject - orchestration function
func NewSubject(p *Project, cliSubject subject.Subject, st subject.SubjectType) error {
    // 1. Validate subject type exists in map
    tmt, exists := SubjectModuleMap[st]
    if !exists {
        return fmt.Errorf("unsupported subject type: %s", string(st))
    }

    // 2. Find target module recursively
    tm, err := findModule(p.Modules, tmt)
    if err != nil { return err }

    // 3. Create subject through factory
    s, err := subject.SubjectFactory(cliSubject, tm.AbsPath, allSubjects)
    if err != nil { return err }

    // 4. Assign UUID
    if err := s.SetID(); err != nil { return err }

    // 5. Persist to disk
    if err := s.WriteToDisk(); err != nil { return err }

    // 6. Update project metadata
    tm.Subjects = append(tm.Subjects, s)
    if err := p.WriteMetadata(); err != nil { return err }

    // 7. Log action
    activitylog.Log(p.ProjectDir(), activitylog.ActionCreate, string(st), s.Name)
    
    return nil
}
```

---

### `internal/subject/` — Subject Types & Operations

**Purpose:** Subject struct definitions, factory pattern, persistence, configuration

**Key Files:**

- **`types.go`** — Type definitions and configuration maps
  - Subject type constants (uppercase): `SubjectEvent = "EVENT"`, etc.
  - `SubjectDataSource = "DATASOURCE"` (new)
  - Metadata file constant: `METADATA_FILE = "META.yaml"`
  - **`SubDirs` map** — Defines default subdirectories created for each subject type
  - **`Files` map** — Defines default files created for each subject type

- **`subject.go`** — Subject operations
  - `SetID()` — Generate and assign UUID v7 (new)
  - `MkDir()` — Create subject directory
  - `MkSubDirs()` — Create all subdirectories
  - `WriteFiles()` — Create default files with headers
  - `WriteMetadata()` — Write META.yaml
  - `ReadMetadata()` — Read META.yaml from disk (new)
  - `WriteToDisk()` — Orchestration: creates dir → subdirs → files → metadata
  - `Describe()` — Pretty-print subject
  - `EditMetadata()` — Open metadata in editor
  - `Rename()` — Rename subject (new)

- **`factory.go`** — Subject creation factory
  - `SubjectFactory()` — Main factory function
  - `silent()` — Factory for non-interactive mode (used by CLI commands)
  - `interactive()` — Factory for interactive mode (prompts user)
  - `nameFactory()` — Generate directory name based on type

- **`name_factory.go`** — Name generation logic (new)
  - `nameFactory()` — Generate directory name by type
  - EVENT/TASK: `YYYY-MM-DD-hyphenated-name`
  - TOPIC/OBJECTIVE/DATASOURCE: `hyphenated-name` only

- **`interactive.go`** — Interactive CLI prompts
  - `interactiveCLI()` — Prompt user for subject properties
  - Type-specific prompts for EVENT, TASK, TOPIC, OBJECTIVE, DATASOURCE

- **`rename.go`** — Subject rename logic (new)
  - `renameSubject()` — Rename directory and update metadata
  - `interactiveRename()` — Interactive prompt for new name

**Configuration Maps:**

```go
// Default subdirectories created during subject creation
var SubDirs SubjectDirMap = SubjectDirMap{
    SubjectEvent: {
        "01_AGENDA",
        "02_MEDIA",
        "03_NOTES",
        "04_DOCUMENTS",
        "05_OUTCOMES",
    },
    SubjectTask: {
        "01_INPUTS",
        "02_WORKING",
        "03_REVIEW",
        "04_FINAL",
    },
    // Topics, Objectives, DataSources have no default subdirs
}

// Default files created during subject creation
var Files SubjectFilesMap = SubjectFilesMap{
    SubjectTopic: {
        "overview.md",
        "notes.md",
    },
    SubjectObjective: {
        "definitions.md",
        "findings.md",
        "strategy.md",
    },
    // Events, Tasks, DataSources have no default files
}
```

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
    
    // Event-specific
    Location           string   `yaml:"location,omitempty"`
    Participants       []string `yaml:"participants,omitempty"`
    
    // Task-specific
    Owner              string   `yaml:"owner,omitempty"`
    Status             string   `yaml:"status,omitempty"`
    RelatedEvents      []string `yaml:"related_events,omitempty"`
    RelatedObjective   string   `yaml:"related_objective,omitempty"`
    Outputs            []string `yaml:"outputs,omitempty"`
    
    // DataSource-specific
    Source             string   `yaml:"source,omitempty"`
    SourceLink         string   `yaml:"source_link,omitempty"`
    SourceObjective    string   `yaml:"source_objective,omitempty"`
    SourceDataSize     string   `yaml:"source_datasize,omitempty"`
}
```

**Name Factory Logic:**

```go
// Names are auto-generated based on type and input
EVENT:      "2026-05-22-cairo-visit"      (date-hyphenated-name)
TASK:       "2026-05-22-fix-bug"          (date-hyphenated-name)
TOPIC:      "machine-learning"            (hyphenated-name only)
OBJECTIVE:  "increase-reliability"        (hyphenated-name only)
DATASOURCE: "sensor-readings-2025"        (hyphenated-name only)
```

---

### `internal/module/` — Module Structure

**Purpose:** Directory structure organization, module types, filesystem bootstrap

**Key Files:**

- **`types.go`** — Type definitions and prefix mapping
  - `ModuleType` constants (uppercase): `ModuleAdmin = "ADMIN"`, etc.
  - `ModuleDirPrefixMap` — Maps module type to directory prefix (00-99)
  - **Updated prefixes:** Tasks (07), Topics (09), Objectives (10), etc.
  - Complete list of all module types including new `ModuleDataSources`

- **`module.go`** — Module operations
  - `MkDir()` — Create module directory
  - `MkSubDirs()` — Create module's default subdirectories
  - `Bootstrap()` — Recursive: creates dir → subdirs → nested modules

**Module Types & Prefixes (Updated):**

```
ModuleAdmin              → "00_ADMIN"
ModuleEvents            → "01_EVENTS"
ModuleProjectManagement → "02_PROJECT_MANAGEMENT"
ModuleLegal             → "03_LEGAL"
ModuleResearch          → "04_RESEARCH"
ModuleEngineering       → "05_ENGINEERING"
ModuleData              → "06_DATA"
ModuleTasks             → "07_TASKS"              (nested under PM, was 05)
ModuleIndex             → "08_INDEX"             (nested under Research, was 06)
ModuleTopics            → "09_TOPICS"            (nested under Research, was 07)
ModuleObjectives        → "10_OBJECTIVES"        (nested under Research, was 08)
ModuleSummaries         → "11_SUMMARIES"         (nested under Research, was 09)
ModuleReferences        → "12_REFERENCES"        (nested under Research, was 10)
ModuleAudioNotes        → "13_AUDIO_NOTES"       (nested under Research, was 11)
ModuleAttachments       → "14_ATTACHMENTS"       (nested under Research, was 12)
ModuleDataSources       → "15_DATASOURCES"       (nested under Data, was 13)
ModulePublications      → "16_PUBLICATIONS"      (was 14)
ModuleMediaLibrary      → "97_MEDIA_LIBRARY"
ModuleDeliverables      → "98_DELIVERABLES"
ModuleArchive           → "99_ARCHIVE"
```

**Module Struct:**

```go
type Module struct {
    Type     ModuleType        `yaml:"type"`
    Name     string            `yaml:"name"`      // e.g., "01_EVENTS"
    AbsPath  string            `yaml:"-"`         // Absolute path, hydrated at load
    Modules  []Module          `yaml:"modules"`   // Nested modules
    Subjects []subject.Subject `yaml:"subjects"`  // Direct subjects
    SubDirs  []string          `yaml:"subDirs"`   // Flat subdirectory
}
```

---

### `internal/filesystem/` — File I/O

**Purpose:** All filesystem operations encapsulated here

**Key Operations:**

- `CheckDirExists(path)` — Check if directory exists
- `CreateDir(path)` — Create directory (fails if exists)
- `ReadFile(path)` — Read file contents
- `StructToFile(struct, path)` — Marshal Go struct to YAML file
- `FileToStruct(struct, path)` — Unmarshal YAML file to Go struct (new)
- `TextToMDFile(text, path)` — Write text to file
- `Archive(src, dest)` — Move file/directory to archive
- `RenameDir(src, dest)` — Rename directory (new)

**Design Philosophy:** Single responsibility — all filesystem I/O goes through this package. Makes it:

- Easy to mock for testing
- Centralized error handling
- Future enhancement opportunity (permissions, backups, etc.)

---

### `internal/activitylog/` — Audit Trail

**Purpose:** Log all user actions for audit and undo

**Key Types & Constants:**

```go
type Action string

const (
    ActionCreate  Action = "CREATE"
    ActionEdit    Action = "EDIT"
    ActionArchive Action = "ARCHIVE"
    ActionRename  Action = "RENAME"  // New
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
- `AppVersion` — Set from main.go build flags

**Design:** Append-only, pipe-friendly for Unix integration (can be piped to `grep`, `cut`, `awk`, etc.)

---

## Subject System Architecture

### Five Subject Types

The system now supports five types of subjects:

#### 1. **Event**
- **Module:** `01_EVENTS`
- **Purpose:** Record project activities (meetings, visits, workshops)
- **Key Fields:** location, participants, date
- **Module Prefix:** 01
- **Subdirectories:** 01_AGENDA, 02_MEDIA, 03_NOTES, 04_DOCUMENTS, 05_OUTCOMES
- **Directory Naming:** `YYYY-MM-DD-hyphenated-name`

#### 2. **Task**
- **Module:** `02_PROJECT_MANAGEMENT/07_TASKS`
- **Purpose:** Unit of work with lifecycle
- **Key Fields:** owner, status, related_events, outputs
- **Module Prefix:** 07 (nested under 02)
- **Subdirectories:** 01_INPUTS, 02_WORKING, 03_REVIEW, 04_FINAL
- **Directory Naming:** `YYYY-MM-DD-hyphenated-name`

#### 3. **Topic**
- **Module:** `04_RESEARCH/09_TOPICS`
- **Purpose:** Knowledge concept or domain area
- **Key Fields:** related_objective, tags, notes
- **Module Prefix:** 09 (nested under 04)
- **Default Files:** overview.md, notes.md
- **Directory Naming:** `hyphenated-name`

#### 4. **Objective**
- **Module:** `04_RESEARCH/10_OBJECTIVES`
- **Purpose:** Goal driving research and decisions
- **Key Fields:** status, outputs, tags
- **Module Prefix:** 10 (nested under 04)
- **Default Files:** definitions.md, findings.md, strategy.md
- **Directory Naming:** `hyphenated-name`

#### 5. **DataSource** (New)
- **Module:** `06_DATA/15_DATASOURCES`
- **Purpose:** External dataset or data feed
- **Key Fields:** source, source_link, source_objective, source_datasize
- **Module Prefix:** 15 (nested under 06)
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
edit → open META.yaml in editor → auto-sync
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

### UUID Generation

- **Type:** v7 (sortable by timestamp)
- **Dependency:** `github.com/google/uuid`
- **Generation:** Happens automatically during `SubjectFactory` → `SetID()`
- **Persistence:** Written to META.yaml `uuid` field

```go
func (s *Subject) SetID() error {
    id7, err := uuid.NewV7()
    if err != nil {
        return err
    }
    s.UUID = id7.String()
    return nil
}
```

### UUID Backfill on Project Load

When a project is loaded, `backfillUUIDs()` ensures all subjects have UUIDs:

```go
func (p *Project) backfillUUIDs() (bool, error) {
    // Walk all subjects
    // If UUID missing → generate and write to disk
    // If subject dir doesn't exist on disk → remove from metadata
    // Return true if any changes made → triggers project.WriteMetadata()
}
```

**This means:**
- Old projects automatically gain UUIDs on first load
- Orphaned subjects are cleaned up
- No migration script needed

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

# Non-interactive datasource creation
operatree add datasource \
  --name "Sensor Readings 2025" \
  --source "IoT Team" \
  --source-link "/06_DATA/01_RAW/sensors.csv" \
  --source-objective "Reduce Downtime" \
  --source-datasize "2.4GB"
```

### Finding Subjects Non-Interactively

```bash
# Get all subjects matching "cairo" (any type)
operatree find --term cairo --plain

# Get all tasks matching "report"
operatree find --term report --type task --plain

# Bulk operation: archive all done tasks
operatree find --term done --type task --plain \
  | grep uuid \
  | awk '{print $2}' \
  | xargs -I{} operatree archive --uuid {}
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

## Template System

**Purpose:** Define project structures, module hierarchies, default layouts

**Key Files:**

- **`types.go`** — Type definitions
  - `OTTemplate` — Template structure
  - `ModuleTemplate` — Nested template structure
  - `Templates` map — Available templates: "general", "dev", "consulting", "research"

- **`list.go`** — List templates
  - `ListTemplates()` — Display all available templates

- **`load.go`** — Load template from embedded YAML
  - `Load(name)` — Load template by name

- **`template_*.yml`** — Template YAML files (embedded)
  - `template_general.yml` — Minimal structure
  - `template_dev.yml` — Software development
  - `template_consulting.yml` — Client engagement
  - `template_research.yml` — Academic/R&D

**Template Structure Example:**

```yaml
# template_dev.yml
name: dev
description: Software development project template
modules:
  - type: ADMIN
    name: ADMIN
    subDirs: []
  - type: EVENTS
    name: EVENTS
    subDirs: []
  - type: PROJECT_MANAGEMENT
    name: PROJECT_MANAGEMENT
    subDirs: []
    modules:
      - type: TASKS
        name: TASKS
        subDirs: []
  # ... more modules
```

---

## Adding New Features

### Adding a New Subject Type

To add a new subject type (e.g., "REPORT"):

**Step 1:** Add constant in `internal/subject/types.go`

```go
const SubjectReport SubjectType = "REPORT"
```

**Step 2:** Add to `SubjectModuleMap` in `internal/project/types.go`

```go
var SubjectModuleMap = map[subject.SubjectType]module.ModuleType{
    // ... existing entries ...
    subject.SubjectReport: module.ModuleReports,  // Module must exist
}
```

**Step 3:** Add module constant in `internal/module/types.go`

```go
const ModuleReports ModuleType = "REPORTS"
```

**Step 4:** Add prefix mapping in `internal/module/types.go`

```go
ModuleDirPrefixMap[ModuleReports] = "08"  // Or appropriate prefix
```

**Step 5:** (Optional) Configure default dirs/files in `internal/subject/types.go`

```go
SubDirs[SubjectReport] = []string{...}
Files[SubjectReport] = []string{...}
```

**Step 6:** (Optional) Add interactive prompts in `internal/subject/interactive.go`

```go
if st == SubjectReport {
    // Add type-specific fields to huh form
}
```

**That's it!** The CLI command automatically recognizes the new type:

```bash
operatree add report --name "Q2 Results" --date 2026-06-30
# Works immediately without changes to cmd/add.go!
```

### Adding a New Command

To add a new command (e.g., `operatree bulk-rename`):

**Step 1:** Create `cmd/bulkrename.go`

```go
package cmd

import (
    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(bulkRenameCmd)
}

var bulkRenameCmd = &cobra.Command{
    Use:   "bulkrename",
    Short: "Rename multiple subjects at once",
    Run:   bulkRename,
}

func bulkRename(cmd *cobra.Command, args []string) {
    // Implementation
}
```

**Step 2:** The command is automatically registered with the root command

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
// In internal/project or cmd
func OperateOnSubject(p *Project, searchTerm string) error {
    // 1. Find subject
    s, err := FindSubject(p, "", searchTerm)
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
├── internal/
│   ├── project/
│   │   ├── bootstrap_test.go
│   │   ├── find_subjects_test.go
│   │   └── ...
│   ├── subject/
│   │   ├── factory_test.go
│   │   └── ...
│   └── ...
```

### Test Patterns

**Unit Tests:** Test individual functions with mocked dependencies

```go
func TestFindSubject_Fuzzy(t *testing.T) {
    // Create test project in memory
    p := &Project{...}
    
    // Test fuzzy search
    s, err := FindSubject(p, "event", "cairo")
    if err != nil { t.Fatal(err) }
    
    if s.Name != "cairo-visit" {
        t.Errorf("expected cairo-visit, got %s", s.Name)
    }
}
```

**Integration Tests:** Test CLI + project operations end-to-end

```go
func TestAddEventCommand(t *testing.T) {
    // Create temp directory
    tmpdir := t.TempDir()
    
    // Bootstrap test project
    p, err := project.Bootstrap("test", tmpdir, "general")
    if err != nil { t.Fatal(err) }
    
    // Run add command
    // Verify result on disk
}
```

---

## Troubleshooting Guide

### Common Issues

**Issue:** "Project not found"
- **Check:** Does `-d` flag point to valid project directory?
- **Check:** Does project contain META.yaml?
- **Fix:** Run `operatree show tracked` to see registered projects

**Issue:** UUID not found
- **Check:** Is UUID correct? (copy-paste from `operatree find --plain`)
- **Check:** Has subject been archived?
- **Fix:** Try `operatree find --term [name]` to get current UUID

**Issue:** Editor not opening
- **Check:** Is `editor` set in config? (`operatree show config`)
- **Check:** Is $EDITOR environment variable set?
- **Fix:** Run `operatree init` to reconfigure editor

**Issue:** Metadata out of sync
- **Check:** Were META.yaml files edited outside of OperaTree?
- **Check:** Did files arrive via git pull or Syncthing?
- **Fix:** Run `operatree sync` to rebuild index

**Issue:** Subject directory structure missing
- **Check:** Does subject directory exist on disk?
- **Check:** Are expected subdirectories present?
- **Fix:** The project likely crashed during WriteToDisk; manually create missing dirs

### Debug Mode

Enable verbose logging (when implemented):

```bash
operatree --verbose add event --name "Test"
```

Check activity log:

```bash
# View all actions
cat activity.log

# View recent actions
tail -20 activity.log

# View all actions by type
grep TASK activity.log
```

---

## Key Data Structures

### Project Metadata (META.yaml)

```yaml
name: fleetfix
template: consulting
modules:
  - type: ADMIN
    name: 00_ADMIN
    modules: []
    subjects: []
    subDirs: []
  - type: EVENTS
    name: 01_EVENTS
    modules: []
    subjects:
      - uuid: a1b2c3d4-e5f6-7g8h-i9j0-k1l2m3n4o5p6
        type: EVENT
        name: 2026-05-22-cairo-visit
        date: "2026-05-22"
        location: Cairo
        participants:
          - Alex
          - Sara
        tags:
          - factory
          - inspection
        notes: "Initial site survey"
        subDirs:
          - 01_AGENDA
          - 02_MEDIA
          - 03_NOTES
          - 04_DOCUMENTS
          - 05_OUTCOMES
    subDirs: []
  # ... more modules ...
```

### Subject Metadata (META.yaml)

```yaml
uuid: a1b2c3d4-e5f6-7g8h-i9j0-k1l2m3n4o5p6
type: EVENT
name: 2026-05-22-cairo-visit
date: "2026-05-22"
location: Cairo
participants:
  - Alex
  - Sara
tags:
  - factory
  - inspection
notes: "Initial site survey"
subDirs:
  - 01_AGENDA
  - 02_MEDIA
  - 03_NOTES
  - 04_DOCUMENTS
  - 05_OUTCOMES
```

---

## Performance Considerations

### Path Hydration Cost

- Hydration runs on every project load
- Cost: O(n) where n = total subjects + modules
- Optimization: Paths only calculated for accessed subjects (lazy loading is not implemented yet)

### Search Performance

- Fuzzy search: O(n*m) where n = subjects, m = search term length
- Index built dynamically on every search (no persistent cache)
- Future: SQLite index sidecar planned for faster queries

### Metadata Sync Performance

- Sync walks entire filesystem tree: O(n)
- Reads all META.yaml files
- Rewrites project META.yaml
- Future: Incremental sync (only changed files) planned

---

This completes the updated ARCHITECTURE.md with all major changes from the `rename-uuid` branch!
