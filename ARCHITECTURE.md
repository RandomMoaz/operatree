# OperaTree Architecture Guide

> **For Contributors:** This document explains OperaTree's design, codebase structure, and how all components work together. Use this to understand the project deeply and make informed contributions.

## Table of Contents

1. [Philosophy & Design Principles](#philosophy--design-principles)
2. [High-Level Architecture](#high-level-architecture)
3. [Package Organization](#package-organization)
4. [Data Flow & Request Lifecycle](#data-flow--request-lifecycle)
5. [Core Concepts](#core-concepts)
6. [Package Deep Dives](#package-deep-dives)
7. [Adding New Features](#adding-new-features)
8. [Common Patterns](#common-patterns)
9. [Testing Strategy](#testing-strategy)
10. [Troubleshooting Guide](#troubleshooting-guide)

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

- Each subject (Event, Task, Topic, Objective) has a `META.yaml` file
- Metadata is searchable, filterable, and machine-readable
- Content can live in the same directory alongside metadata
- Users can edit metadata with their preferred editor

---

## High-Level Architecture

### Layered Design

```
┌──────────────────────────────────────────────────┐
│          CLI Layer (cmd/)                        │
│  (Commands: new, find, metadata, archive, etc.)  │
└──────────────────┬───────────────────────────────┘
                   │
┌──────────────────▼───────────────────────────────┐
│       Business Logic Layer (internal/)           │
│  project, subject, module, metadata handling     │
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
User Command (e.g., "operatree new event")
        │
        ├─→ cmd/new.go (CLI command handler)
        │   └─→ Parses arguments, maps to subject type
        │
        ├─→ internal/project/new_subject.go
        │   └─→ NewSubject(project, name, date, subjectType)
        │       ├─→ Finds target module via findModule()
        │       ├─→ Creates subject via subject.SubjectFactory()
        │       ├─→ Writes to disk via subject.WriteToDisk()
        │       ├─→ Updates project metadata
        │       └─→ Logs to activity.log
        │
        └─→ internal/filesystem/ + internal/subject/
            └─→ Persist to disk
```

---

## Package Organization

### Top-Level Structure

```
operatree/
├── cmd/                    # CLI commands
│   ├── new.go             # Create new subject
│   ├── find.go            # Search subjects
│   ├── metadata.go        # Edit metadata
│   ├── archive.go         # Archive subject
│   ├── bootstrap.go       # Create new project
│   ├── sync.go            # Sync metadata
│   ├── root.go            # Cobra setup
│   └── ...
│
├── internal/              # Business logic (not exported)
│   ├── project/           # Project management
│   ├── subject/           # Subject types & operations
│   ├── module/            # Module (directory) structure
│   ├── config/            # Configuration management
│   ├── filesystem/        # File I/O operations
│   ├── metadata/          # Metadata parsing
│   ├── activitylog/       # Audit trail
│   ├── runner/            # External command execution
│   ├── ui/                # Terminal UI formatting
│   └── help/              # Help text
│
├── main.go                # Entry point
├── go.mod                 # Dependencies
├── go.sum                 # Dependency checksums
├── Makefile               # Build configuration
└── README.md              # User documentation
```

### Dependency Graph

```
cmd/ (depends on)
  └─→ internal/project/
  └─→ internal/subject/
  └─→ internal/config/
  └─→ internal/runner/
  └─→ internal/ui/

internal/project/ (depends on)
  └─→ internal/module/
  └─→ internal/subject/
  └─→ internal/filesystem/
  └─→ internal/activitylog/
  └─→ internal/metadata/

internal/subject/ (depends on)
  └─→ internal/metadata/
  └─→ internal/filesystem/

internal/module/ (depends on)
  └─→ internal/subject/
  └─→ internal/filesystem/

internal/filesystem/ (depends on)
  └─→ [Standard library only]
```

---

## Data Flow & Request Lifecycle

### Example: Creating a New Event

```
User Input: operatree new event --name "Cairo Visit" --date "2026-05-22"
│
├─→ cmd/new.go :: newSubject()
│   ├─ Parse "event" argument to subject.SubjectEvent
│   ├─ Load project from config (or current dir)
│   └─ Call project.NewSubject(p, "Cairo Visit", "2026-05-22", SubjectEvent)
│
├─→ internal/project/new_subject.go :: NewSubject()
│   ├─ Find module for SubjectEvent via findModule()
│   │  └─→ Recursively search p.Modules for ModuleEvents
│   │
│   ├─ Create subject instance via subject.SubjectFactory()
│   │  └─→ Validates input, assigns unique ID
│   │
│   ├─ Persist to disk via subject.WriteToDisk()
│   │  └─→ internal/filesystem/Create(subjectDir/META.yaml)
│   │
│   ├─ Update project metadata
│   │  └─→ p.WriteMetadata() → filesystem.StructToFile()
│   │
│   ├─ Log to activity.log
│   │  └─→ internal/activitylog.Log(CREATE, event, "Cairo Visit")
│   │
│   └─ Print confirmation to stdout
│
└─→ File System
    └─→ project/
        └─→ 01_EVENTS/
            └─→ cairo-visit/
                └─→ META.yaml (subject metadata)
```

### Data Structure Flow

```
Subject Type (CLI)    Subject Type (Internal)    Module Type (Storage)
─────────────────     ──────────────────────     ──────────────────
"event"        ──→    SubjectEvent         ──→   ModuleEvents (01_EVENTS/)
"task"         ──→    SubjectTask          ──→   ModuleTasks (02_PROJECT_MANAGEMENT/Tasks/)
"topic"        ──→    SubjectTopic         ──→   ModuleTopics (04_RESEARCH/Topics/)
"objective"    ──→    SubjectObjective     ──→   ModuleObjectives (04_RESEARCH/Objectives/)
```

**Mapping Logic:** `internal/project/types.go :: SubjectModuleMap`

---

## Core Concepts

### 1. **Projects**

- **What:** A collection of subjects organized into modules
- **Storage:** `~/projects/myproject/` directory
- **Metadata:** `METADATA.yml` in project root
- **Structure:** Nested modules (dirs) containing subjects (subdirs)

### 2. **Modules**

- **What:** Directories that organize subjects by category
- **Types:**
  - `00_ADMIN` — Governance, contacts, templates
  - `01_EVENTS` — Visits, workshops, meetings
  - `02_PROJECT_MANAGEMENT` — Tasks, reports, risks (nested: Tasks)
  - `03_LEGAL` — Contracts, NDAs, compliance
  - `04_RESEARCH` — Topics, objectives (nested: Topics, Objectives)
  - `05_ENGINEERING` — Architecture, specs, decisions
  - `06_DATA` — Raw → staging → processed pipeline
  - `07_MEDIA_LIBRARY` — Shared reusable assets
  - `08_DELIVERABLES` — Final external outputs
  - `99_ARCHIVE` — Historical storage (nested: closed_tasks)
- **Nesting:** Some modules contain submodules (e.g., Tasks under Project Management)

### 3. **Subjects**

- **What:** Trackable units of work or knowledge
- **Types:**
  - `Event` — Project activity (date, location, participants)
  - `Task` — Unit of work with lifecycle (owner, status)
  - `Topic` — Knowledge concept (tags, notes)
  - `Objective` — Goal driving decisions (status, findings)
- **Storage:** Each subject is a directory with `META.yaml`

### 4. **Metadata**

- **What:** YAML file containing subject properties
- **Location:** `subject-name/META.yaml`
- **Format:** YAML (human-readable, version-control friendly)
- **Editability:** Users can edit directly; sync updates project index

### 5. **Activity Log**

- **What:** Append-only audit trail
- **Location:** `project-root/activity.log`
- **Format:** Tab-separated, pipe-friendly
- **Entries:** Every CREATE, EDIT, DELETE action

---

## Package Deep Dives

### `cmd/` — CLI Layer

**Purpose:** Command-line interface, argument parsing, user interaction

**Key Files:**

- `root.go` — Cobra setup, global flags, project resolution
- `new.go` — Create new subject (unified command)
- `find.go` — Fuzzy search subjects
- `metadata.go` — Edit subject metadata
- `archive.go` — Archive (move to 99_ARCHIVE)
- `bootstrap.go` — Create new project
- `sync.go` — Sync project metadata from disk

**Patterns:**

```go
// Typical command handler pattern
func commandName(cmd *cobra.Command, args []string) {
    // 1. Load project
    p, err := project.Load(prjDir)
    if err != nil { log.Fatal(err) }

    // 2. Call business logic
    if err := project.SomeFunction(&p, arg1, arg2); err != nil {
        log.Fatal(err)
    }
}
```

**Dependencies:** None on other CLI files; all depend on `internal/project`

**How to Add a New Command:**

1. Create `cmd/newcommand.go`
2. Define `var newcommandCmd = &cobra.Command{...}`
3. Add initialization in `init()` function
4. Call `rootCmd.AddCommand(newcommandCmd)`

---

### `internal/project/` — Project Management

**Purpose:** High-level project operations, orchestration

**Key Files:**

- `project.go` — Project struct methods (ProjectDir, Describe, WriteMetadata, Archive)
- `types.go` — Project struct, SubjectModuleMap, project templates
- `new_subject.go` — Create new subject (unified function)
- `bootstrap.go` — Project initialization
- `archive.go` — Archive subjects
- `sync.go` — Sync metadata from disk
- `find_subjects.go` — Fuzzy search
- `summary.go` — Project statistics
- `describe.go` — Pretty-print project structure
- `search_builder.go` — Build searchable index
- `list.go` — List subjects by type
- `template_*.go` — Project templates (dev, general)

**Core Type:**

```go
type Project struct {
    Name     string          // e.g., "myproject"
    Template string          // e.g., "dev"
    BaseDir  string          // e.g., "/home/user/projects"
    Tags     []string        // Project-level tags
    Modules  []module.Module // Top-level modules
}
```

**Key Patterns:**

**Pattern 1: Unified Subject Creation**

```go
// Single function with subject type parameter
func NewSubject(p *Project, name, date string, st SubjectType) error {
    // Map subject type to module type
    tmt := SubjectModuleMap[st]

    // Find module recursively
    tm, err := findModule(p.Modules, tmt)

    // Create and persist subject
    s, err := subject.SubjectFactory(...)
    // ...
}
```

**Pattern 2: Search Index Building**

```go
// BuildSearchDB recursively walks all modules and subjects
// Returns []SearchDB (flattened list of all subjects with metadata)
db := BuildSearchDB(p)

// Used by: find, list, summary operations
for _, entry := range db {
    fmt.Println(entry.Subject.Name)
}
```

**Dependencies:**

- `internal/subject` — Subject operations
- `internal/module` — Module structure
- `internal/filesystem` — File I/O
- `internal/activitylog` — Logging
- `internal/metadata` — YAML handling

---

### `internal/subject/` — Subject Types & Operations

**Purpose:** Subject struct definitions, factory, persistence

**Key Concepts:**

**Subject Types:**

```go
type SubjectType string

const (
    SubjectEvent     SubjectType = "event"
    SubjectTask      SubjectType = "task"
    SubjectTopic     SubjectType = "topic"
    SubjectObjective SubjectType = "objective"
)
```

**Subject Struct:**

```go
type Subject struct {
    Type         SubjectType
    Name         string
    Date         string
    DirName      string  // Directory path
    Tags         []string
    Notes        string
    Status       string  // For Task, Objective
    Location     string  // For Event
    Participants []string // For Event
    Owner        string  // For Task
    // ... other type-specific fields
}
```

**Factory Pattern:**

```go
// SubjectFactory creates a new subject with validation
s, err := subject.SubjectFactory(initialSubject, modulePath, existingSubjects)
// - Validates input
// - Generates unique ID (directory name)
// - Prevents naming conflicts
```

**Key Operations:**

- `SubjectFactory()` — Create new subject
- `WriteToDisk()` — Persist to META.yaml
- `Load()` — Load from META.yaml

**Dependencies:**

- `internal/filesystem` — File I/O
- `internal/metadata` — YAML parsing

---

### `internal/module/` — Module Structure

**Purpose:** Directory structure organization, module types

**Module Types:**

```go
type ModuleType string

const (
    ModuleAdmin              ModuleType = "admin"
    ModuleEvents             ModuleType = "events"
    ModuleProjectManagement  ModuleType = "projectmanagement"
    ModuleTasks              ModuleType = "tasks"
    ModuleResearch           ModuleType = "research"
    ModuleTopics             ModuleType = "topics"
    // ... etc
)
```

**Module Struct:**

```go
type Module struct {
    Name     string
    Type     ModuleType
    AbsPath  string         // Absolute filesystem path
    Subjects []Subject      // Subjects at this level
    Modules  []Module       // Nested submodules
}
```

**Factory Functions:**

```go
// Templates create module hierarchies
module.FactoryAdmin(projectPath, prefix)      // Creates 00_ADMIN
module.FactoryEvents(projectPath, prefix)     // Creates 01_EVENTS
module.FactoryProjectManagement(projectPath)  // Creates 02_PROJECT_MANAGEMENT with Tasks submodule
// ... etc
```

**Key Operations:**

- `Bootstrap()` — Create module directories
- `Load()` — Load subjects from disk
- Nested structure support (for Tasks under ProjectManagement)

---

### `internal/config/` — Configuration Management

**Purpose:** User configuration, project tracking

**Config File Location:** `~/.config/operatree/operatree.yaml`

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

**Key Operations:**

- `Load()` — Load config from disk
- `Save()` — Persist config changes
- `AddProject()` — Register new project
- `SetDefault()` — Set default project

---

### `internal/filesystem/` — File I/O

**Purpose:** All filesystem operations encapsulated here

**Key Operations:**

- `CreateDir(path)` — Create directory
- `ReadFile(path)` — Read file contents
- `WriteFile(path, content)` — Write file
- `StructToFile(struct, path)` — Marshal struct to YAML file
- `Archive(src, dest)` — Move subject to archive
- `FileExists(path)` — Check if file exists

**Design:** Single responsibility — all filesystem I/O goes through this package. This makes it:

- Easy to mock for testing
- Centralized error handling
- Potential for future enhancements (permissions, backups, etc.)

---

### `internal/activitylog/` — Audit Trail

**Purpose:** Log all user actions for audit and undo

**Log Format:**

```
timestamp    action   type       name              user@host         version
2026-05-20T10:08:39Z   CREATE   event   "Cairo Visit"  hany@optiplex7040  v0.1.0
```

**Tab-separated columns:** timestamp, action, type, name, user@host, version

**Key Operations:**

- `Log(projectDir, action, type, name)` — Record action
- Actions: CREATE, EDIT, DELETE, ARCHIVE

**Design:** Append-only, pipe-friendly for Unix integration

---

### `internal/ui/` — Terminal Formatting

**Purpose:** Pretty-printing, colored output, terminal aesthetics

**Key Functions:**

- ANSI color codes (AnsiRed, AnsiGreen, AnsiBold, etc.)
- Progress bars, status indicators
- Formatted output for summary, describe

**Dependencies:** Charmbracelet libraries (lipgloss, glamour)

---

### `internal/metadata/` — YAML Serialization

**Purpose:** Marshal/unmarshal YAML, metadata validation

**Key Operations:**

- YAML parsing using `gopkg.in/yaml.v3`
- Struct ↔ YAML conversion
- Tag handling (omitempty for type-specific fields)

---

### `internal/runner/` — External Commands

**Purpose:** Execute external programs (editor, file manager, commands)

**Examples:**

- `runner.OpenInEditor(filePath, editorCmd)` — Open file in editor
- `runner.OpenInFileManager(dirPath, fmCmd)` — Open directory in file manager

**Design:** Encapsulates subprocess execution, error handling

---

## Adding New Features

### Scenario 1: Add a New Subject Type

**Steps:**

1. **Define Subject Type** (`internal/subject/types.go`):

```go
const SubjectMytype SubjectType = "mytype"
```

2. **Add to Module Mapping** (`internal/project/types.go`):

```go
var SubjectModuleMap = map[subject.SubjectType]module.ModuleType{
    // ... existing
    subject.SubjectMytype: module.ModuleMyModule,
}
```

3. **Add CLI Integration** (`cmd/new.go`):

```go
var argToSubject map[string]subject.SubjectType = map[string]subject.SubjectType{
    // ... existing
    "mytype": subject.SubjectMytype,
}
```

4. **Add Module Template** (`internal/module/factory.go`):

```go
func FactoryMyModule(projectPath, prefix string) Module {
    return Module{
        Name: fmt.Sprintf("%s_MYMODULE", prefix),
        Type: ModuleMyModule,
        AbsPath: path.Join(projectPath, name),
        Subjects: []Subject{},
        Modules: []Module{},
    }
}
```

5. **Update Project Templates** (`internal/project/template_dev.go`):

```go
p := Project{
    // ...
    Modules: []module.Module{
        // ... existing
        module.FactoryMyModule(ppth, "09"),
    },
}
```

6. **Test it:**

```bash
operatree new mytype --name "My New Type"
```

---

### Scenario 2: Add a New Command

**Steps:**

1. **Create Command File** (`cmd/mycommand.go`):

```go
package cmd

import "github.com/spf13/cobra"

var mycommandCmd = &cobra.Command{
    Use:   "mycommand [args]",
    Short: "Description of my command",
    Long:  "Longer description...",
    Run:   runMyCommand,
}

func runMyCommand(cmd *cobra.Command, args []string) {
    // Load project
    p, err := project.Load(prjDir)
    if err != nil { log.Fatal(err) }

    // Call business logic
    if err := project.MyFunction(&p); err != nil {
        log.Fatal(err)
    }
}

func init() {
    mycommandCmd.Flags().StringVar(&someFlag, "flag", "", "Description")
    rootCmd.AddCommand(mycommandCmd)
}
```

2. **Add Business Logic** (`internal/project/myfunction.go`):

```go
func MyFunction(p *Project) error {
    // Implementation
    return nil
}
```

3. **Wire Together** — The `init()` function in your cmd file adds the command to rootCmd

4. **Test:**

```bash
operatree mycommand
```

---

### Scenario 3: Add Search Enhancement

**Search Logic Location:** `internal/project/search_builder.go`

**How It Works:**

```go
// BuildSearchDB creates a flattened searchable index
// Each entry contains: subject data + concatenated metadata string
func BuildSearchDB(p *Project) []SearchDB {
    // Walks all modules recursively
    // For each subject, concatenates all searchable fields:
    // tags + participants + name + notes + date + location
    // Returns as one big searchable string
}

// Then fuzzy search happens in FindSubjects()
// Uses: github.com/lithammer/fuzzysearch
```

**To Enhance Search:**

1. Add new fields to concatenation in `BuildSearchDB()`
2. Or implement advanced ranking in `FindSubjects()`

---

## Common Patterns

### Pattern 1: Project Loading & Error Handling

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

// Always wrap errors with context
```

### Pattern 2: Recursive Module Traversal

```go
// For operations on nested modules (Tasks under ProjectManagement, etc.)
func processModule(m *module.Module) error {
    // Process this module's subjects
    for i, s := range m.Subjects {
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

### Pattern 3: Defensive Map Access

```go
// When accessing maps that might be missing keys
val, exists := SubjectModuleMap[subjectType]
if !exists {
    return fmt.Errorf("unsupported subject type: %s", string(subjectType))
}
```

### Pattern 4: Non-Fatal Error Handling

```go
// Some operations should not block others
if err := activitylog.Log(...); err != nil {
    fmt.Fprintf(os.Stderr, "warning: could not write activity log: %v\n", err)
    // Continue, don't return
}
```

### Pattern 5: Pointer Safety in Loops

```go
// When you need pointers to slice elements
for i := range modules {
    // Take pointer to actual element, not loop variable copy
    ptr := &modules[i]

    // Now mutations to ptr persist to the slice
    ptr.Subjects = append(ptr.Subjects, newSubject)
}
```

---

## Testing Strategy

### Test Organization

```
operatree/
├── cmd/
│   └── *_test.go
├── internal/
│   ├── project/
│   │   └── *_test.go
│   ├── subject/
│   │   └── *_test.go
│   └── ...
└── testdata/         # Test fixtures, example projects
```

### Test Patterns

**1. Unit Tests (Logic)**

```go
func TestNewSubject(t *testing.T) {
    // Arrange: Set up project
    p := createTestProject()

    // Act: Call function
    err := project.NewSubject(&p, "Test", "2026-05-22", subject.SubjectEvent)

    // Assert: Check results
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    // Verify side effects
    if len(p.Modules[0].Subjects) != 1 {
        t.Errorf("expected 1 subject, got %d", len(p.Modules[0].Subjects))
    }
}
```

**2. Integration Tests (Filesystem)**

```go
func TestProjectPersistence(t *testing.T) {
    // Create temp directory
    tmpDir := t.TempDir()

    // Create project
    p, err := project.Bootstrap("test", tmpDir, "dev")
    if err != nil {
        t.Fatalf("failed to bootstrap: %v", err)
    }

    // Verify files exist
    if _, err := os.Stat(path.Join(p.ProjectDir(), "METADATA.yml")); err != nil {
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

### Issue: "Project doesn't contain module X"

**Cause:** Module not found in project structure

**Debug Steps:**

1. Check project template: `operatree desc`
2. Verify module hierarchy is correct
3. Check `internal/project/template_*.go` for module definitions

**Fix:**

- Bootstrap project with correct template
- Or manually add module to project directory

---

### Issue: "Subject already exists"

**Cause:** Directory name collision

**Debug Steps:**

1. Check subject directory names: `ls -la module-dir/`
2. Look at generated directory name logic

**Fix:**

- Use different subject name to get different directory

---

### Issue: "Metadata sync fails"

**Cause:** Malformed YAML in subject directory

**Debug Steps:**

1. Check subject's META.yaml: `cat subject-dir/META.yaml`
2. Validate YAML syntax: `yamllint META.yaml`
3. Check `internal/project/sync.go` logging

**Fix:**

- Fix YAML syntax manually
- Delete and recreate subject
- Run `operatree sync` to repair index

---

### Issue: "Search not returning expected results"

**Cause:** Search index not built correctly

**Debug Steps:**

1. Check which fields are searchable in `internal/project/search_builder.go`
2. Search should match: name, tags, participants, notes, date, location
3. Verify metadata was synced: `operatree sync`

**Fix:**

- Run sync to rebuild index
- Check subject metadata is complete
- Try broader search terms

---

## Code Review Checklist for Contributors

Before submitting a PR:

- [ ] **Follows existing patterns** — Uses established patterns from codebase
- [ ] **Error handling** — All errors wrapped with context
- [ ] **Defensive checks** — Validates inputs, checks for nil pointers
- [ ] **Filesystem-first** — Data flows through filesystem, not memory
- [ ] **CLI separation** — Business logic separated from CLI layer
- [ ] **Comments** — Complex logic has clear comments
- [ ] **Testing** — Unit tests for new logic
- [ ] **YAML-friendly** — Config/metadata is valid YAML
- [ ] **No breaking changes** — Existing data structures remain valid
- [ ] **Activity log** — User actions logged appropriately

---

## Contribution Ideas

### High-Impact Areas

1. **New Subject Types**
   - Scientific experiment tracking
   - Meeting minutes
   - Budget tracking
   - Complexity: Low-Medium

2. **New Project Templates**
   - Legal/compliance template
   - Creative project template
   - Startup template
   - Complexity: Low

3. **Search Enhancements**
   - Regex support
   - Advanced filtering
   - Search ranking
   - Complexity: Medium

4. **Version Control Integration**
   - Git hooks
   - Automatic commits
   - Diff visualization
   - Complexity: Medium-High

5. **Output Formatters**
   - JSON export
   - CSV export
   - Markdown export
   - Complexity: Low-Medium

---

## Getting Help

- **Architecture Questions:** Open an issue with `[ARCHITECTURE]` label
- **Design Discussions:** Start a discussion in GitHub Discussions
- **Code Questions:** Tag maintainers in PRs for detailed review
- **Bug Reports:** Include code snippets, error logs, steps to reproduce

---

## Key Takeaways

1. **Filesystem is the source of truth** — Everything persists to disk
2. **Layered architecture** — CLI → Business Logic → Persistence
3. **Package responsibility** — Each package has one clear purpose
4. **Error handling first** — Defensive programming throughout
5. **YAML friendly** — Human-readable, version-control compatible
6. **No breaking changes** — Users own their data format
7. **Unix philosophy** — Compose small tools, output is pipe-friendly
8. **Testing matters** — Especially for filesystem operations

---

**Last Updated:** May 2026  
**OperaTree Version:** Alpha  
**Status:** Active Development
