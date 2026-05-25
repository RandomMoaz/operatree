# OperaTree Architecture Guide

> **For Contributors:** This document explains OperaTree's design, codebase structure, and how all components work together. Use this to understand the project deeply and make informed contribution decisions.

## Table of Contents

1. [Philosophy & Design Principles](#philosophy--design-principles)
2. [High-Level Architecture](#high-level-architecture)
3. [Package Organization](#package-organization)
4. [Data Flow & Request Lifecycle](#data-flow--request-lifecycle)
5. [Path Resolution & Portability](#path-resolution--portability)
6. [Core Concepts](#core-concepts)
7. [Package Deep Dives](#package-deep-dives)
8. [Subject Creation Deep Dive](#subject-creation-deep-dive)
9. [Template System](#template-system)
10. [Adding New Features](#adding-new-features)
11. [Common Patterns](#common-patterns)
12. [Testing Strategy](#testing-strategy)
13. [Troubleshooting Guide](#troubleshooting-guide)

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

- Each subject (Event, Task, Topic, Objective) has a `METADATA.yml` file
- Metadata is searchable, filterable, and machine-readable
- Content can live in the same directory alongside metadata
- Users can edit metadata with their preferred editor

---

## High-Level Architecture

### Layered Design

```
┌──────────────────────────────────────────────────┐
│          CLI Layer (cmd/)                        │
│  (Commands: new, find, bootstrap, sync, etc.)    │
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
User Command (e.g., "operatree new event --name 'Cairo Visit'")
        │
        ├─→ cmd/new.go (CLI command handler)
        │   ├─ Parse "event" argument (lowercase CLI input)
        │   ├─ Convert to SubjectType constant (UPPERCASE: "EVENT")
        │   ├─ Resolve project directory (-d flag)
        │   └─ Call project.NewSubject()
        │
        ├─→ internal/project/new_subject.go
        │   ├─ Validate subject type against SubjectModuleMap
        │   ├─ Find target module recursively
        │   ├─ Create subject instance
        │   ├─ Run through subject factory
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
├── cmd/                    # CLI commands (11 files)
│   ├── root.go            # Cobra setup, global flags, project resolution
│   ├── new.go             # Create new subject (DYNAMIC subject type loading)
│   ├── find.go            # Search subjects
│   ├── bootstrap.go       # Create new project
│   ├── sync.go            # Sync metadata from disk
│   ├── summary.go         # Project statistics
│   ├── open.go            # Open subject in file manager
│   ├── show.go            # Display config/templates/tracked projects
│   ├── track.go           # Add project to tracked list
│   ├── untrack.go         # Remove project from tracked list
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

### Example: Creating a New Event

```
User Input: operatree new event --name "Cairo Visit" --date "2026-05-22" -d ~/myproject
│
├─→ cmd/root.go :: Execute()
│   └─→ Cobra parses flags and routes to newCmd
│
├─→ cmd/root.go :: resolveProjectDir() (PreRun hook)
│   ├─ Checks -d flag → actDir = ~/myproject
│   └─ Converts "." to absolute path if needed
│
├─→ cmd/new.go :: newSubject()
│   ├─ Get argument: "event" (lowercase from CLI)
│   ├─ Convert to uppercase: "event" → "EVENT"
│   ├─ Create SubjectType constant: subject.SubjectType("EVENT")
│   ├─ Load project from actDir
│   └─ Call project.NewSubject(&p, "Cairo Visit", "2026-05-22", SubjectEvent)
│
├─→ internal/project/new_subject.go :: NewSubject()
│   ├─ Get all existing subjects (for name collision detection)
│   ├─ Validate subject type exists in SubjectModuleMap
│   ├─ Map SubjectEvent to ModuleEvents
│   ├─ Recursively search project.Modules for ModuleEvents
│   │
│   ├─ Create initial subject struct:
│   │  Subject{Type: EVENT, Name: "Cairo Visit", Date: "2026-05-22"}
│   │
│   ├─ Call subject.SubjectFactory(initialSubject, modulePath, existSubjects)
│   │  └─→ Enters silent mode (name provided)
│   │
│   └─→ Call s.WriteToDisk()
│       └─→ internal/subject/subject.go
│           ├─ Create subject directory: ~/myproject/01_EVENTS/2026-05-22-cairo-visit/
│           ├─ Create subdirs: 01_AGENDA, 02_MEDIA, 03_NOTES, 04_DOCUMENTS, 05_OUTCOMES
│           ├─ Create default files: (none for Events)
│           └─ Write METADATA.yml with subject data
│
├─→ Update project metadata
│   ├─ Append subject to module.Subjects[]
│   ├─ Write project METADATA.yml
│   └─ internal/project/hydrate.go :: hydratePath() paths are recalculated
│
├─→ Log action
│   └─→ internal/activitylog/activitylog.go
│       ├─ Build entry: timestamp, action=CREATE, type=EVENT, name="2026-05-22-cairo-visit"
│       ├─ Get user/hostname info
│       └─ Append to activity.log in project root
│
└─→ Output confirmation
    └─→ "EVENT created: 2026-05-22-cairo-visit"
```

### Subject Type Conversion Flow

**Key Detail:** Subject types have **two representations** and use **dynamic loading**:

```
CLI Argument          Internal Constant     Storage Module Type
─────────────         ─────────────────     ───────────────────
"event"    (lower)  → SubjectEvent("EVENT") → ModuleEvents
"task"     (lower)  → SubjectTask("TASK")   → ModuleTasks
"topic"    (lower)  → SubjectTopic("TOPIC") → ModuleTopics
"objective"(lower)  → SubjectObjective("OBJECTIVE") → ModuleObjectives
```

**Dynamic Loading in `cmd/new.go`:**

```go
func init() {
    // Build completion slice DYNAMICALLY from SubjectModuleMap
    // This means adding a new subject type automatically updates the CLI!
    for k := range project.SubjectModuleMap {
        sn := strings.ToLower(string(k))
        validSubjects = append(validSubjects, sn)
    }

    newCmd = &cobra.Command{
        Use:       fmt.Sprintf("new [%s]", strings.Join(validSubjects, " | ")),
        ValidArgs: validSubjects,  // Dynamically populated
        Args:      cobra.MatchAll(cobra.OnlyValidArgs, cobra.ExactArgs(1)),
        Run:       newSubject,
    }
}

func newSubject(cmd *cobra.Command, args []string) {
    a := args[0]                              // "event" from CLI
    st := strings.ToUpper(a)                  // "EVENT"
    
    // Convert to SubjectType constant
    if err := project.NewSubject(&p, subjectName, subjectDate, subject.SubjectType(st)); err != nil {
        log.Fatal(err)
    }
}
```

**Mapping Logic in `internal/project/types.go`:**

```go
var SubjectModuleMap = map[subject.SubjectType]module.ModuleType{
    subject.SubjectEvent:     module.ModuleEvents,
    subject.SubjectTask:      module.ModuleTasks,
    subject.SubjectTopic:     module.ModuleTopics,
    subject.SubjectObjective: module.ModuleObjectives,
}
```

**Key Advantage:** When you add a new subject type, the CLI command automatically recognizes it without code changes to the new command!

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
operatree new event -d ~/work/reports/sales-2026
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
- **Types:** See `internal/module/types.go` for complete list:
  - `00_ADMIN` — Governance, contacts, templates
  - `01_EVENTS` — Visits, workshops, meetings
  - `02_PROJECT_MANAGEMENT` — Tasks, reports, risks (contains nested `05_TASKS`)
  - `03_LEGAL` — Contracts, NDAs, compliance
  - `04_RESEARCH` — Topics, objectives (contains nested modules)
  - `05_ENGINEERING` — Architecture, specs, decisions
  - `06_DATA` — Raw → staging → processed pipeline
  - `97_MEDIA_LIBRARY` — Shared reusable assets
  - `98_DELIVERABLES` — Final external outputs
  - `99_ARCHIVE` — Historical storage
- **Nesting:** Some modules contain submodules (e.g., `05_TASKS` under `02_PROJECT_MANAGEMENT`)
- **Storage:** `Module.Subjects[]` contains direct subjects; `Module.Modules[]` contains nested modules

### 3. **Subjects**

- **What:** Trackable units of work or knowledge
- **Types:** EVENT, TASK, TOPIC, OBJECTIVE (see `internal/subject/types.go`)
- **Storage:** Each subject is a directory with `METADATA.yml`
- **Structure:** Subjects auto-create subdirectories and default files based on their type

### 4. **Metadata**

- **What:** YAML file containing subject/project properties
- **Location:** `subject-dir/METADATA.yml` or `project-dir/METADATA.yml`
- **Format:** YAML (human-readable, version-control friendly)
- **Editability:** Users can edit directly; `sync` command updates project index

### 5. **Activity Log**

- **What:** Append-only audit trail of all changes
- **Location:** `project-root/activity.log`
- **Format:** Tab-separated values
- **Entries:** Timestamp, action (CREATE/EDIT/DELETE/ARCHIVE), type, name, user@host, version

**Example entry:**
```
2026-05-20T10:08:39Z	CREATE  	EVENT        	"2026-05-22-cairo-visit"	hany@optiplex7040	v0.1.0
```

---

## Package Deep Dives

### `cmd/` — CLI Layer

**Purpose:** Command-line interface, argument parsing, user interaction

**Key Files:**

- **`root.go`** — Cobra setup, global flags, config loading, path resolution
  - Loads configuration at startup
  - Defines global variables: `destDir`, `actDir`, `cfg`, `verbose`
  - Sets up root command

- **`utilities.go`** — Path resolution helpers
  - `resolveProjectDir()` — Resolves `-d` flag for project commands
  - `resolveBaseDir()` — Resolves `-d` flag for base directory commands
  - `resolveProjectDirSkippingConfig()` — Ignores config, only uses explicit flags

- **`new.go`** — Create new subject
  - **DYNAMICALLY loads valid subject types** from `project.SubjectModuleMap`
  - Converts CLI argument (lowercase) to SubjectType constant (uppercase)
  - Calls `project.NewSubject()` with subject type
  - **No static map needed** — automatically picks up new subject types!

- **`bootstrap.go`** — Create new project
  - Takes project name and template name
  - Uses `-d` for base directory (where project is created)
  - Calls `project.Bootstrap()`

- **`find.go`** — Search subjects
  - Fuzzy search by type and term
  - Calls `project.FindSubjects()`

- **`sync.go`** — Sync project metadata
  - Walks all subjects on disk
  - Updates project in-memory from METADATA.yml files
  - Calls `project.Sync()`

- **`summary.go`** — Project statistics
  - Displays high-level project overview
  - Subject counts by type

- **`open.go`** — Open subject in file manager
  - Finds subject, opens its directory

- **`show.go`** — Display information
  - Shows tracked projects, config, templates, default project
  - No project `-d` flag needed

- **`track.go`** — Add project to tracked list
  - Adds project to config for future default project usage
  - Requires `-d` flag

- **`untrack.go`** — Remove project from tracked list
  - Removes project from config
  - Requires `-d` flag

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
  - `WriteMetadata()` — Persist project METADATA.yml
  - `ModuleExists()` — Check if module exists
  - `Archive()` — Archive a subject

- **`types.go`** — Type definitions
  - `Project` struct with `absDir` (hydrated absolute path)
  - `SubjectModuleMap` — Maps each subject type to its storage module
  - Metadata file constant: `METADATA_FILE = "METADATA.yml"`
  - Archive destination constant: `ARCHIVED_DEST = "closed_tasks"`

- **`new_subject.go`** — Create new subject
  - `NewSubject()` — Main orchestration function
  - `findModule()` — Recursively search for target module by type
  - Validates subject type, finds module, creates subject, persists to disk

- **`load.go`** — Load project from disk
  - `Load(path)` — Reads METADATA.yml and calls `hydratePath()`

- **`hydrate.go`** — Path hydration
  - `hydratePath()` — Set absolute paths on project and all modules
  - `hydrateModule()` — Recursive path hydration for nested modules
  - Sets `AbsPath` on all modules and `DirName` on all subjects

- **`bootstrap.go`** — Create new project structure
  - `Bootstrap()` — Load template, create project, create directories
  - Calls `internal/template/Load()` to get template
  - Calls `internal/project/Factory()` to build project structure
  - Calls `module.Bootstrap()` to create directories

- **`factory.go`** — Build project from template
  - `Factory()` — Convert template to project structure
  - `parseModule()` — Recursively parse template modules to module objects

- **`sync.go`** — Sync metadata from disk
  - `Sync()` — Read all subject METADATA.yml files and update project
  - `syncModule()` — Recursively sync all subjects in a module

- **`list_subjects.go`** — List all subjects
  - `ListSubjects()` — Flatten project tree into list of all subjects

- **`find_subjects.go`** — Search subjects
  - `FindSubjects()` — Fuzzy search by type and term

- **`describe.go`** — Pretty-print project
  - Formatted project output for terminal display

- **`archive.go`** — Archive subjects
  - `Archive()` — Move subject to 99_ARCHIVE module

- **`summary.go`** — Project statistics
  - `Summary()` — Display subject counts by type

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
func NewSubject(p *Project, subjectName, subjectDate string, st subject.SubjectType) error {
    // 1. Validate subject type exists in map
    tmt, exists := SubjectModuleMap[st]
    if !exists {
        return fmt.Errorf("unsupported subject type: %s", string(st))
    }

    // 2. Find target module recursively
    tm, err := findModule(p.Modules, tmt)
    if err != nil { return err }

    // 3. Create subject through factory
    s, err := subject.SubjectFactory(initialSubject, tm.AbsPath, allSubjects)
    if err != nil { return err }

    // 4. Persist to disk
    if err := s.WriteToDisk(); err != nil { return err }

    // 5. Update project metadata
    tm.Subjects = append(tm.Subjects, s)
    if err := p.WriteMetadata(); err != nil { return err }

    // 6. Log action
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
  - Metadata file constant: `METADATA_FILE = "METADATA.yml"`
  - **`SubDirs` map** — Defines default subdirectories created for each subject type
  - **`Files` map** — Defines default files created for each subject type

- **`subject.go`** — Subject operations
  - `MkDir()` — Create subject directory
  - `MkSubDirs()` — Create all subdirectories
  - `WriteFiles()` — Create default files with headers
  - `WriteMetadata()` — Write METADATA.yml
  - `WriteToDisk()` — Orchestration: creates dir → subdirs → files → metadata
  - `Describe()` — Pretty-print subject
  - `EditMetadata()` — Open metadata in editor

- **`factory.go`** — Subject creation factory
  - `SubjectFactory()` — Main factory function
  - `silent()` — Factory for non-interactive mode (used by CLI commands)
  - `interactive()` — Factory for interactive mode (prompts user)
  - `nameFactory()` — Generate directory name based on type

- **`interactive.go`** — Interactive CLI prompts
  - `interactiveCLI()` — Prompt user for subject properties
  - Type-specific prompts for EVENT, TASK, TOPIC, OBJECTIVE

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
    // Topics and Objectives have no default subdirs
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
    // Events and Tasks have no default files
}
```

**Subject Struct:**

```go
type Subject struct {
    Type              SubjectType `yaml:"type"`
    Name              string      `yaml:"name"`
    DirName           string      `yaml:"-"`  // Not persisted, hydrated at load
    SubDirs           []string    `yaml:"subDirs"`
    Files             []string    `yaml:"-"` // Not persisted, used for creation only
    Date              string      `yaml:"date"`
    Tags              []string    `yaml:"tags"`
    Notes             string      `yaml:"notes"`
    // Custom fields for specific types (use omitempty)
    Paricipants       []string `yaml:"paricipants,omitempty"`
    Location          string   `yaml:"location,omitempty"`
    Owner             string   `yaml:"owner,omitempty"`
    Status            string   `yaml:"status,omitempty"`
    RelatedObjective  string   `yaml:"related_objective,omitempty"`
    RelatedEvents     []string `yaml:"related_events,omitempty"`
    Outputs           []string `yaml:"outputs,omitempty"`
}
```

**Name Factory Logic:**

```go
// Names are auto-generated based on type and input
EVENT:     "2026-05-22-cairo-visit"      (date-hyphenated-name)
TASK:      "2026-05-22-fix-bug"          (date-hyphenated-name)
TOPIC:     "machine-learning"            (hyphenated-name only)
OBJECTIVE: "increase-reliability"        (hyphenated-name only)
```

---

### `internal/module/` — Module Structure

**Purpose:** Directory structure organization, module types, filesystem bootstrap

**Key Files:**

- **`types.go`** — Type definitions and prefix mapping
  - `ModuleType` constants (uppercase): `ModuleAdmin = "ADMIN"`, etc.
  - `ModuleDirPrefixMap` — Maps module type to directory prefix (00-99)
  - Complete list of all module types

- **`module.go`** — Module operations
  - `MkDir()` — Create module directory
  - `MkSubDirs()` — Create module's default subdirectories
  - `Bootstrap()` — Recursive: creates dir → subdirs → nested modules

**Module Types & Prefixes:**

```
ModuleAdmin              → "00_ADMIN"
ModuleEvents            → "01_EVENTS"
ModuleProjectManagement → "02_PROJECT_MANAGEMENT"
ModuleLegal             → "03_LEGAL"
ModuleResearch          → "04_RESEARCH"
ModuleEngineering       → "05_ENGINEERING"
ModuleData              → "06_DATA"
ModuleTasks             → "05_TASKS"              (nested under PM)
ModuleTopics            → "07_TOPICS"            (nested under Research)
ModuleObjectives        → "08_OBJECTIVES"        (nested under Research)
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
    SubDirs  []string          `yaml:"subDirs"`   // Flat subdirectories
}
```

---

### `internal/template/` — Project Templates

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

### `internal/config/` — Configuration Management

**Purpose:** User configuration, project tracking, persistence

**Key Files:**

- **`config.go`** — Configuration operations
  - Config file location: `~/.config/operatree/operatree.yaml`
  - Respects `XDG_CONFIG_HOME` environment variable for Linux users
  - `Load()` — Load config from YAML
  - `Save()` — Persist config to YAML
  - `AddProject()` — Register new project
  - `RemoveProject()` — Unregister project
  - `SetDefaultProject()` — Set default project

**Config File Location:**

```
Linux:   $XDG_CONFIG_HOME/operatree/operatree.yaml  (or ~/.config/operatree/operatree.yaml)
macOS:   ~/Library/Application Support/operatree/operatree.yaml
Windows: %APPDATA%\operatree\operatree.yaml
```

**Config Structure:**

```yaml
standardDir: /home/user/projects
editor: nvim
fileManager: nautilus
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
```

---

### `internal/filesystem/` — File I/O

**Purpose:** All filesystem operations encapsulated here

**Key Operations:**

- `CheckDirExists(path)` — Check if directory exists
- `CreateDir(path)` — Create directory (fails if exists)
- `ReadFile(path)` — Read file contents
- `StructToFile(struct, path)` — Marshal Go struct to YAML file
- `TextToMDFile(text, path)` — Write text to file
- `Archive(src, dest)` — Move file/directory to archive

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
    ActionDelete  Action = "DELETE"
    ActionArchive Action = "ARCHIVE"
)
```

**Log Format (Tab-Separated):**

```
timestamp                 action    type        name                      user@host            version
2026-05-20T10:08:39Z     CREATE    EVENT       "2026-05-22-cairo-visit"  hany@optiplex7040    v0.1.0
```

**Key Operations:**

- `Log(projectRoot, action, subjectType, subjectName)` �� Record action
- `AppVersion` — Set from main.go build flags

**Design:** Append-only, pipe-friendly for Unix integration (can be piped to `grep`, `cut`, `awk`, etc.)

---

### `internal/metadata/` — Metadata Utilities

**Purpose:** Metadata parsing, name formatting

**Key Operations:**

- `FormatName(name)` — Sanitize and hyphenate names
- YAML marshaling/unmarshaling using `gopkg.in/yaml.v3`

---

### `internal/runner/` — External Command Execution

**Purpose:** Execute external programs (editor, file manager)

**Key Operations:**

- `EditFile(path)` — Open file in configured editor
- `OpenFileManager(path)` — Open directory in file manager

---

### `internal/ui/` — Terminal Formatting

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

## Subject Creation Deep Dive

Understanding subject creation is crucial for contributors. Here's the complete flow:

### Step 1: CLI Parsing with Dynamic Loading

```bash
operatree new event --name "Cairo Visit" --date "2026-05-22"
```

**In `cmd/new.go` init():**

```go
func init() {
    // 1. DYNAMICALLY build valid subjects from SubjectModuleMap
    for k := range project.SubjectModuleMap {
        sn := strings.ToLower(string(k))
        validSubjects = append(validSubjects, sn)
    }
    
    // 2. Create command with dynamic ValidArgs
    newCmd = &cobra.Command{
        Use:       fmt.Sprintf("new [%s]", strings.Join(validSubjects, " | ")),
        ValidArgs: validSubjects,  // ["event", "task", "topic", "objective"]
        Args:      cobra.MatchAll(cobra.OnlyValidArgs, cobra.ExactArgs(1)),
        Run:       newSubject,
    }
    
    // 3. Add flags and set PreRun hook
    newCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
    newCmd.Flags().StringVar(&subjectName, "name", "", "subject name")
    newCmd.Flags().StringVar(&subjectDate, "date", "", "subject date")
    newCmd.PreRun = resolveProjectDir
    rootCmd.AddCommand(newCmd)
}
```

**In `cmd/new.go` newSubject():**

```go
func newSubject(cmd *cobra.Command, args []string) {
    a := args[0]  // "event" (lowercase CLI input)
    
    // Convert to uppercase and create SubjectType constant
    st := strings.ToUpper(a)  // "EVENT"
    
    // Load project with path hydration
    p, err := project.Load(actDir)
    if err != nil { log.Fatal(err) }
    
    // Pass to business logic - no mapping needed, direct type creation
    if err := project.NewSubject(&p, subjectName, subjectDate, subject.SubjectType(st)); err != nil {
        log.Fatal(err)
    }
}
```

### Step 2: Subject Factory

**In `internal/project/new_subject.go`:**

```go
func NewSubject(p *Project, subjectName, subjectDate string, st subject.SubjectType) error {
    // 1. Create initial subject struct
    is := subject.Subject{
        Type: st,                // SubjectEvent ("EVENT")
        Name: subjectName,       // "Cairo Visit"
        Date: subjectDate,       // "2026-05-22"
    }
    
    // 2. Get target module
    tmt := SubjectModuleMap[st]  // ModuleEvents
    tm, err := findModule(p.Modules, tmt)
    if err != nil { return err }
    
    // 3. Call factory - this does all the setup
    s, err := subject.SubjectFactory(is, tm.AbsPath, listOfExistingSubjects)
    if err != nil { return err }
    
    // 4. Persist to disk
    if err := s.WriteToDisk(); err != nil { return err }
    
    // 5. Update project and log
    tm.Subjects = append(tm.Subjects, s)
    if err := p.WriteMetadata(); err != nil { return err }
    
    activitylog.Log(p.ProjectDir(), activitylog.ActionCreate, string(st), s.Name)
    
    return nil
}
```

### Step 3: Subject Factory - Silent Mode

**In `internal/subject/factory.go`:**

```go
func SubjectFactory(s Subject, ppth string, pss []Subject) (Subject, error) {
    // Since Name is provided, use silent mode (not interactive)
    if s.Name == "" {
        return interactive(s.Type, ppth, pss)  // Interactive mode
    }
    
    return silent(s, ppth)  // Silent mode (CLI commands)
}

func silent(s Subject, ppth string) (Subject, error) {
    // 1. Set default subdirectories from configuration map
    s.SubDirs = SubDirs[s.Type]  // For EVENT: [01_AGENDA, 02_MEDIA, ...]
    s.Files = Files[s.Type]      // For EVENT: [] (no files)
    
    // 2. Generate directory name using name factory
    s.Name = nameFactory(s)      // "2026-05-22-cairo-visit"
    
    // 3. Compute absolute path
    s.DirName = path.Join(ppth, s.Name)  // ~/project/01_EVENTS/2026-05-22-cairo-visit
    
    return s, nil
}

func nameFactory(s Subject) string {
    switch s.Type {
    case SubjectEvent:
        sn := metadata.FormatName(s.Name)      // "cairo-visit"
        if s.Date != "" {
            sn = fmt.Sprintf("%s-%s", s.Date, sn)  // "2026-05-22-cairo-visit"
        }
        return sn
    case SubjectTask:
        sn := metadata.FormatName(s.Name)
        return fmt.Sprintf("%s-%s", s.Date, sn)
    case SubjectTopic:
        return metadata.FormatName(s.Name)
    default:
        return s.Name
    }
}
```

### Step 4: Write to Disk

**In `internal/subject/subject.go`:**

```go
func (s *Subject) WriteToDisk() error {
    // 1. Create subject directory
    if err := s.MkDir(); err != nil { return err }
    // Creates: ~/project/01_EVENTS/2026-05-22-cairo-visit/
    
    // 2. Create all subdirectories
    if err := s.MkSubDirs(); err != nil { return err }
    // Creates: 01_AGENDA/, 02_MEDIA/, 03_NOTES/, 04_DOCUMENTS/, 05_OUTCOMES/
    
    // 3. Create default files
    if err := s.WriteFiles(); err != nil { return err }
    // For EVENT: creates nothing
    // For TOPIC: creates overview.md, notes.md
    
    // 4. Write metadata file
    if err := s.WriteMetadata(); err != nil { return err }
    // Creates: METADATA.yml with subject data (marshaled to YAML)
    
    return nil
}
```

### Step 5: Result on Disk

After successful creation:

```
~/project/01_EVENTS/
└── 2026-05-22-cairo-visit/
    ├── METADATA.yml              # Subject metadata (YAML)
    ├── 01_AGENDA/                # Empty directory
    ├── 02_MEDIA/                 # Empty directory
    ├── 03_NOTES/                 # Empty directory
    ├── 04_DOCUMENTS/             # Empty directory
    └── 05_OUTCOMES/              # Empty directory
```

**METADATA.yml Content:**

```yaml
type: EVENT
name: Cairo Visit
date: 2026-05-22
tags: []
notes: ""
subDirs:
  - 01_AGENDA
  - 02_MEDIA
  - 03_NOTES
  - 04_DOCUMENTS
  - 05_OUTCOMES
paricipants: []
location: ""
```

### Configuration Maps Control Everything

These maps in `internal/subject/types.go` define the directory and file structure:

```go
// What subdirectories are created for each subject type
var SubDirs SubjectDirMap = SubjectDirMap{
    SubjectEvent: {"01_AGENDA", "02_MEDIA", "03_NOTES", "04_DOCUMENTS", "05_OUTCOMES"},
    SubjectTask: {"01_INPUTS", "02_WORKING", "03_REVIEW", "04_FINAL"},
    // SubjectTopic and SubjectObjective have NO subdirectories
}

// What files are created for each subject type
var Files SubjectFilesMap = SubjectFilesMap{
    SubjectTopic: {"overview.md", "notes.md"},
    SubjectObjective: {"definitions.md", "findings.md", "strategy.md"},
    // SubjectEvent and SubjectTask have NO files
}
```

---

## Template System

### How Templates Work

Templates define the structure of a bootstrapped project. The workflow:

```
1. User runs: operatree bootstrap myproject --template dev -d ~/projects
2. CLI loads template: internal/template/Load("dev")
3. Factory converts template to project structure: internal/project/Factory()
4. Modules are recursively created: internal/module/Bootstrap()
5. Project METADATA.yml is written
```

### Template Structure

Templates are YAML files that define:

- Module hierarchy
- Nesting relationships (Tasks under ProjectManagement)
- Default subdirectories for each module

**Example template entry:**

```yaml
name: dev
description: Software development project template
modules:
  - type: ADMIN
    name: ADMIN
    subDirs: []
  - type: PROJECT_MANAGEMENT
    name: PROJECT_MANAGEMENT
    subDirs: []
    modules:  # ← Nested module
      - type: TASKS
        name: TASKS
        subDirs: []
```

### Available Templates

- **general** — Minimal structure, good starting point
- **dev** — Software development projects
- **consulting** — Client engagement work
- **research** — Academic and R&D projects

---

## Adding New Features

### Scenario 1: Add a New Subject Type

**Complete step-by-step guide:**

**Step 1: Define Subject Type** (`internal/subject/types.go`):

```go
const SubjectMeeting SubjectType = "MEETING"
```

**Step 2: Add Configuration Maps** (`internal/subject/types.go`):

```go
var SubDirs SubjectDirMap = SubjectDirMap{
    // ... existing ...
    SubjectMeeting: {
        "01_AGENDA",
        "02_MOM",           // Minutes of Meeting
        "03_ATTENDEES",
    },
}

var Files SubjectFilesMap = SubjectFilesMap{
    // ... existing ...
    SubjectMeeting: {
        "notes.md",
    },
}
```

**Step 3: Add Module Mapping** (`internal/project/types.go`):

```go
var SubjectModuleMap = map[subject.SubjectType]module.ModuleType{
    subject.SubjectEvent:     module.ModuleEvents,
    subject.SubjectTask:      module.ModuleTasks,
    subject.SubjectTopic:     module.ModuleTopics,
    subject.SubjectObjective: module.ModuleObjectives,
    subject.SubjectMeeting:   module.ModuleEvents,  // Store in Events module
}
```

**That's it!** The CLI will automatically recognize the new subject type because `cmd/new.go` dynamically loads from `project.SubjectModuleMap`. No changes needed to the `new` command!

**Step 4: Optional - Add Interactive Mode** (`internal/subject/interactive.go`):

If you want interactive prompts for the new type:

```go
if st == SubjectMeeting {
    var agenda string
    var attendees string
    
    err := huh.NewForm(
        huh.NewGroup(
            huh.NewText().Title("Agenda").Value(&agenda),
            huh.NewText().Title("Attendees").Value(&attendees),
        ),
    ).Run()
    if err != nil { return err }
    
    s.Agenda = agenda
    s.Attendees = attendees
}
```

**Step 5: Update Subject Struct** (`internal/subject/types.go`):

If you added type-specific fields in interactive mode, add them to `Subject` struct:

```go
type Subject struct {
    // ... existing fields ...
    Agenda    string   `yaml:"agenda,omitempty"`
    Attendees []string `yaml:"attendees,omitempty"`
}
```

**Step 6: Update Project Templates** (if needed):

Add the new subject's module to templates if appropriate.

**Test it:**

```bash
operatree new meeting --name "Q2 Planning" --date "2026-06-01"
```

---

### Scenario 2: Add a New Command

**Steps:**

1. **Create Command File** (`cmd/mycommand.go`):

```go
package cmd

import (
    "log"
    "github.com/hanymamdouh82/operatree/internal/project"
    "github.com/spf13/cobra"
)

func init() {
    myCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
    myCmd.PreRun = resolveProjectDir
    rootCmd.AddCommand(myCmd)
}

var myCmd = &cobra.Command{
    Use:   "mycommand [args]",
    Short: "Short description",
    Long:  "Longer description of what this command does",
    Args:  cobra.ExactArgs(0),
    Run:   runMyCommand,
}

func runMyCommand(cmd *cobra.Command, args []string) {
    // 1. Load project
    p, err := project.Load(actDir)
    if err != nil { log.Fatal(err) }

    // 2. Call business logic
    if err := project.MyFunction(&p); err != nil {
        log.Fatal(err)
    }
}
```

2. **Add Business Logic** (`internal/project/myfunction.go`):

```go
package project

func MyFunction(p *Project) error {
    // Implementation
    return nil
}
```

3. **Test it:**

```bash
operatree mycommand
```

---

### Scenario 3: Enhance Search

**Location:** `internal/project/find_subjects.go`

Current search uses fuzzy matching via `github.com/lithammer/fuzzysearch`.

To enhance:

1. Add searchable fields to the concatenated string
2. Implement ranking/scoring logic
3. Support advanced query syntax

---

## Common Patterns

### Pattern 1: Loading & Error Handling

```go
// Always load project first
p, err := project.Load(projectDir)
if err != nil {
    return fmt.Errorf("failed to load project: %w", err)
}

// Work with project
if err := project.NewSubject(&p, name, date, subjectType); err != nil {
    return fmt.Errorf("failed to create subject: %w", err)
}
```

### Pattern 2: Recursive Module Traversal

```go
// For operations on nested modules
func processModule(m *module.Module) error {
    // Process subjects at this level
    for _, s := range m.Subjects {
        if err := processSubject(s); err != nil {
            return err
        }
    }

    // Recurse into submodules
    for i := range m.Modules {
        if err := processModule(&m.Modules[i]); err != nil {
            return err
        }
    }

    return nil
}
```

### Pattern 3: Pointer Safety in Loops

```go
// When you need pointers to slice elements
for i := range modules {
    ptr := &modules[i]  // Pointer to actual slice element
    ptr.Subjects = append(ptr.Subjects, newSubject)  // Persists to slice
}
```

### Pattern 4: Defensive Map Access

```go
// Always check if key exists
val, exists := SubjectModuleMap[subjectType]
if !exists {
    return fmt.Errorf("unsupported subject type: %s", string(subjectType))
}
```

### Pattern 5: Non-Fatal Errors

```go
// Some operations should not block others
if err := activitylog.Log(...); err != nil {
    fmt.Fprintf(os.Stderr, "warning: could not write activity log: %v\n", err)
    // Continue, don't return
}
```

---

## Testing Strategy

### Test Organization

```
operatree/
├── cmd/
│   └── *_test.go          # Command handler tests
├── internal/
│   ├── project/
│   │   └── *_test.go      # Business logic tests
│   ├── subject/
│   │   └── *_test.go      # Subject factory tests
│   └── ...
└── testdata/              # Test fixtures, example projects
```

### Test Patterns

**1. Unit Tests (Logic)**

```go
func TestNewSubject(t *testing.T) {
    // Arrange
    p := createTestProject()

    // Act
    err := project.NewSubject(&p, "Test", "2026-05-22", subject.SubjectEvent)

    // Assert
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if len(p.Modules[0].Subjects) != 1 {
        t.Errorf("expected 1 subject, got %d", len(p.Modules[0].Subjects))
    }
}
```

**2. Integration Tests (Filesystem)**

```go
func TestProjectBootstrap(t *testing.T) {
    // Create temp directory
    tmpDir := t.TempDir()

    // Bootstrap project
    p, err := project.Bootstrap("test", tmpDir, "dev")
    if err != nil {
        t.Fatalf("failed to bootstrap: %v", err)
    }

    // Verify files exist
    metadata := path.Join(p.ProjectDir(), project.METADATA_FILE)
    if _, err := os.Stat(metadata); err != nil {
        t.Errorf("metadata file not found: %v", err)
    }
}
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run specific package
go test ./internal/project

# Run with coverage
go test -cover ./...

# Verbose output
go test -v ./...
```

---

## Troubleshooting Guide

### Issue: "Unsupported subject type"

**Cause:** Subject type not defined or not registered

**Debug:**

1. Check CLI accepts it: `operatree new --help`
2. Verify constant in `internal/subject/types.go`
3. Verify mapping in `internal/project/types.go`

**Fix:**

- Add subject type definition (Step 1 from Scenario 1)
- Add to SubjectModuleMap (Step 3 from Scenario 1)
- CLI will automatically pick it up!

---

### Issue: "Module type X not found"

**Cause:** Module missing from project structure or not mapped

**Debug:**

1. Check project template: `operatree show templates`
2. Verify module in template YAML
3. Verify `SubjectModuleMap` has entry

**Fix:**

- Update template file
- Or manually add module directory to project

---

### Issue: "Subject already exists"

**Cause:** Directory name collision

**Debug:**

1. List subjects: `ls -la module-dir/`
2. Check name factory logic

**Fix:**

- Use different subject name to get different directory name

---

### Issue: "Metadata sync fails"

**Cause:** Malformed YAML in subject METADATA.yml

**Debug:**

1. Check file: `cat subject-dir/METADATA.yml`
2. Validate YAML: `yamllint subject-dir/METADATA.yml`

**Fix:**

- Fix YAML syntax in metadata file
- Run `operatree sync` to rebuild index

---

### Issue: "Config file not found"

**Cause:** First run, no config yet

**Debug:**

1. Check config location: `echo $XDG_CONFIG_HOME` (Linux)
2. Run: `operatree show config`

**Fix:**

- Run `operatree bootstrap` to create first project
- Config is automatically created on first project creation

---

### Issue: Paths break after moving project

**Expected:** Paths should NOT break

If they do, it's a bug. Projects should be fully portable.

**Verify:**

- Use `-d` with new location: `operatree show -d /new/location`
- No config changes should be needed
- If this fails, file an issue
