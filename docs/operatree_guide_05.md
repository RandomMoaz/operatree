# Section 5 — Managing Projects

---

This section covers the commands used to manage projects themselves — tracking, switching, inspecting, and keeping metadata in sync. We continue following Alex and Sara as the `fleetfix` engagement grows.

---

## 5.1 Viewing Tracked Projects

A few weeks into the engagement, Alex is now running several projects in parallel. To see everything currently tracked:

```bash
operatree show tracked
```

The output lists all registered projects with their names, paths, and templates:

```
  fleetfix          /home/alex/projects/fleetfix         consulting
  anchor            /home/alex/projects/anchor            dev
  atentec-mgmt      /home/alex/shared/atentec-mgmt       general
```

This is the starting point for any project management operation — knowing what you have before deciding what to do with it.

---

## 5.2 Switching Between Projects

Alex needs to shift focus from `fleetfix` to `anchor` for the afternoon. The current default is checked first:

```bash
operatree show default
```

```
Default project: fleetfix
Path: /home/alex/projects/fleetfix
```

To switch:

```bash
operatree use
```

An interactive picker lists all tracked projects. Alex selects `anchor`. From this point, every command targets `anchor` without needing `-d`.

OperaTree always requires an explicit selection — even if only one project is tracked. This is by design. When projects are shared across machines and team members, knowing exactly which project is active matters. There is no auto-select.

At the end of the day, Alex switches back:

```bash
operatree use
```

Selects `fleetfix`. Done.

---

## 5.3 Opening a Project in the File Manager

Sometimes Alex wants to browse the full project directory — not a specific subject, but the whole thing. The `goto` command opens the project root in the configured file manager:

```bash
operatree goto
```

An interactive picker lists all tracked projects. Alex selects `fleetfix` and the file manager opens at the project root.

Note that `goto` opens a file manager window — it does not change the working directory of your terminal session. This is a fundamental OS constraint: a child process cannot change the directory of its parent shell. If you need to navigate to a project directory in the terminal, use:

```bash
cd /home/alex/projects/fleetfix
```

Or set up a shell alias for projects you frequently navigate to.

---

## 5.4 Tracking a New Project

Two months in, a third team member — Omar — joins the engagement. Like Sara before him, Omar has the `fleetfix` directory already on his machine via Syncthing. After running `operatree init`, he registers the project:

```bash
operatree track -d /home/omar/shared/fleetfix
```

The `-d` flag is required for `track` — there is no interactive picker since OperaTree needs the exact path to register. Omar can now immediately run any command against `fleetfix`.

If Omar already has a config file restored from a previous machine, OperaTree detects it on startup and skips `init`. If `fleetfix` is already in that config, he does not need to run `track` either.

---

## 5.5 Untracking a Project

Six months later, `fleetfix` is delivered and closed. Alex removes it from the tracked list:

```bash
operatree untrack fleetfix
```

Or by path:

```bash
operatree untrack -d /home/alex/projects/fleetfix
```

The project disappears from `operatree show tracked` immediately. There is no confirmation step.

**The project directory and all its contents are completely untouched.** Untracking only removes the entry from OperaTree's configuration — it does not delete, move, or modify any files. The `fleetfix` directory remains exactly where it is, fully intact, readable by any tool.

If `fleetfix` was the default project, the default is cleared. OperaTree sets no new default automatically — the next `operatree show default` will show empty. This is safe: all commands continue to work normally when a project is specified with `-d`, which is the typical approach in scripted environments and when OperaTree is used as an underlying engine for higher-level tools.

To set a new default after untracking:

```bash
operatree use
```

---

## 5.6 Describing a Project

Alex wants to share a structured overview of the `fleetfix` project with a new stakeholder. The `describe` command prints a styled view of the project structure and metadata:

```bash
operatree describe
```

For piping into other tools or saving to a file, the `--plain` flag outputs raw YAML:

```bash
operatree describe --plain
operatree describe --plain > fleetfix-snapshot.yaml
operatree describe --plain | grep tags
```

To describe a project other than the current default:

```bash
operatree describe -d /home/alex/projects/anchor
```

---

## 5.7 Project Summary

For a quick pulse check — how many subjects are in the project, what types, what statuses — Alex uses:

```bash
operatree summary
```

This is typically the first command of the day. It gives a high-level count of subjects broken down by type and status, without opening a file manager or browsing individual entries. A useful orientation before deciding what to work on.

```bash
operatree summary -d /home/alex/projects/anchor    # summary of a specific project
```

---

## 5.8 Syncing Project Metadata

Sara has been working directly on the shared Syncthing directory. This morning she opened a `META.yaml` file in her text editor and updated the status of a task manually — a quick edit that did not go through `operatree edit`.

When subject files are edited outside of OperaTree — directly in a text editor, after a git pull, or after a Syncthing sync brings in changes from another machine — the project metadata index can fall out of sync with what is actually on disk. To reconcile:

```bash
operatree sync
```

OperaTree walks the full project tree, re-reads every `META.yaml` from disk, and updates the index. The command is safe to run at any time — it reads from disk and never overwrites your files.

Note that `operatree edit` runs sync automatically after the editor closes, so manual sync is only needed when edits happen outside of OperaTree. In shared environments like Sara and Alex's Syncthing setup, running `operatree sync` at the start of the day is a good habit — it ensures the index reflects any changes that arrived overnight from other team members.

```bash
operatree sync -d /home/alex/projects/fleetfix    # sync a specific project
```

---

## 5.9 The Configuration File

Everything OperaTree knows about your projects lives in the configuration file. It is plain YAML and can be read or edited directly in any text editor:

```yaml
standardDir: /home/alex/projects
editor: nvim
fileManager: nautilus
default:
  name: fleetfix
  absPath: /home/alex/projects/fleetfix
  template: consulting
projects:
  - name: fleetfix
    absPath: /home/alex/projects/fleetfix
    template: consulting
  - name: anchor
    absPath: /home/alex/projects/anchor
    template: dev
  - name: atentec-mgmt
    absPath: /home/alex/shared/atentec-mgmt
    template: general
```

To view the current configuration without opening the file:

```bash
operatree show config
```

If you edit the configuration file directly — for example to change your editor or file manager — no sync is needed. OperaTree reads the config file fresh on every command.

---

_Next: Section 6 — Working with Subjects_
