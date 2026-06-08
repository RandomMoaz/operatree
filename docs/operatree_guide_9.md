# Section 9 — Configuration Reference

---

OperaTree's configuration is a single plain YAML file created by `operatree init`
and stored in the system config directory for your platform. It is the only place
OperaTree stores global state — everything else lives inside your project directories.

---

## 9.1 Config File Location

| Platform    | Location                                                 |
| ----------- | -------------------------------------------------------- |
| Linux       | `~/.config/operatree/operatree.yaml`                     |
| Linux (XDG) | `$XDG_CONFIG_HOME/operatree/operatree.yaml`              |
| macOS       | `~/Library/Application Support/operatree/operatree.yaml` |
| Windows     | `%APPDATA%\operatree\operatree.yaml`                     |

To view the current configuration without opening the file:

```bash
operatree show config
```

To reconfigure — change your editor, file manager, or standard directory:

```bash
operatree init
```

Running `init` on an existing installation opens the interactive prompt with your
current values pre-filled. Edit what you need and confirm. You do not need to
re-enter values you want to keep.

---

## 9.2 Full Configuration Reference

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
daemon:
  enabled: false
  host: localhost
  port: 7070
  dbDriver: sqlite
  dsn: ""
```

---

## 9.3 Field Reference

### `standardDir`

The base directory used when creating new projects without the `-d` flag.

```yaml
standardDir: /home/alex/projects
```

- **Required** — must be set during `operatree init`
- **Not validated at startup** — if the directory does not exist or is unavailable,
  OperaTree starts normally without error. The directory is only validated when
  `operatree create` is run without `-d`
- **Does not affect existing projects** — tracked projects store their own absolute
  paths independently. Changing `standardDir` only affects where new projects are
  created going forward
- **Supports removable drives and network shares** — if `standardDir` points to a
  location that is only available at certain times (a USB drive, a Syncthing share,
  a network mount), OperaTree will simply not be able to create new projects there
  when it is unavailable. All other commands that target tracked projects with
  absolute paths continue to work normally

---

### `editor`

The text editor opened by `operatree edit`.

```yaml
editor: nvim
```

- **Optional** — falls back to the `$EDITOR` environment variable if not set
- **Any editor works** — use any editor that can be launched from the terminal:
  `nvim`, `vim`, `nano`, `code`, `subl`, `hx`, and so on
- **GUI editors** — editors like VS Code (`code`) and Sublime Text (`subl`) work
  if they are available on your `$PATH` and can be launched from the terminal
- If neither `editor` nor `$EDITOR` is set, `operatree edit` will not function
  until one is configured. Run `operatree init` to set it

---

### `fileManager`

The file manager opened by `operatree open` and `operatree goto`.

```yaml
fileManager: nautilus
```

- **Optional** — no automatic fallback if not set
- **Any file manager works** — use any file manager that accepts a directory path
  as an argument: `nautilus`, `thunar`, `dolphin`, `nemo`, `finder` (macOS),
  `explorer` (Windows), and so on
- If not set, `operatree open` and `operatree goto` will not function until a
  file manager is configured. Run `operatree init` to set it

---

### `default`

The currently active default project. Set by `operatree use`.

```yaml
default:
  name: fleetfix
  absPath: /home/alex/projects/fleetfix
  template: consulting
```

- **Set interactively** via `operatree use` — do not edit this field manually
- **Cleared automatically** when the default project is untracked via
  `operatree untrack`
- **Empty default is safe** — all commands continue to work normally when `-d`
  is provided explicitly. An empty default only means commands without `-d` will
  require you to be inside a project directory that contains a `META.yaml`, or
  will return a helpful error

| Field      | Description                                |
| ---------- | ------------------------------------------ |
| `name`     | Project name as registered                 |
| `absPath`  | Absolute path on this machine              |
| `template` | Template used when the project was created |

---

### `projects`

The list of all tracked projects on this machine.

```yaml
projects:
  - name: fleetfix
    absPath: /home/alex/projects/fleetfix
    template: consulting
  - name: anchor
    absPath: /home/alex/projects/anchor
    template: dev
```

- **Managed by OperaTree** — entries are added by `operatree create` and
  `operatree track`, and removed by `operatree untrack`
- **Paths are absolute and per-machine** — two team members sharing the same
  project via Syncthing will have different `absPath` values in their respective
  config files, pointing to wherever the project lives on their own machine
- **Do not edit manually** — the `projects` list is the source of truth for
  `operatree show tracked`, `operatree use`, and `operatree goto`. Manual edits
  can cause inconsistencies

| Field      | Description                                |
| ---------- | ------------------------------------------ |
| `name`     | Project name                               |
| `absPath`  | Absolute path on this machine              |
| `template` | Template used when the project was created |

---

### `daemon`

Reserved for future use. The daemon is not yet available.

```yaml
daemon:
  enabled: false
  host: localhost
  port: 7070
  dbDriver: sqlite
  dsn: ""
```

These fields configure `operatree-daemon` — a background service that will
provide an API layer and fast query index over your projects. The daemon is
currently under active development. Leave this section as-is until the daemon
is released.

---

## 9.4 Config Validation

OperaTree validates the configuration file on every startup. If the file is
missing required fields or is malformed, OperaTree logs the specific issue and
exits with a helpful error message.

Common validation scenarios:

| Situation                              | Behaviour                                                                      |
| -------------------------------------- | ------------------------------------------------------------------------------ |
| Config file not found                  | OperaTree prompts you to run `operatree init`                                  |
| `standardDir` does not exist           | No error at startup — only validated during `operatree create` without `-d`    |
| `editor` not set and `$EDITOR` not set | Startup succeeds — error only when `operatree edit` is run                     |
| `fileManager` not set                  | Startup succeeds — error only when `operatree open` or `operatree goto` is run |
| Malformed YAML                         | OperaTree logs the parse error and exits                                       |
| Unknown fields                         | Logged as warnings — OperaTree continues normally                              |

The lazy validation of `standardDir`, `editor`, and `fileManager` is intentional.
A user who works entirely via `-d` flags and a terminal editor may never need
these fields to resolve. OperaTree only enforces a field when the command that
depends on it is actually run.

---

## 9.5 Config on Shared and Multi-Machine Setups

The config file is **per user and per machine**. It is never stored inside a
project directory and should never be committed to version control or synced
across machines via Syncthing.

Each user maintains their own config pointing to wherever their projects live
on their own machine:

```
Alex's machine                    Sara's machine
──────────────────────────────    ──────────────────────────────
~/.config/operatree/              ~/.config/operatree/
  operatree.yaml                    operatree.yaml
    standardDir: /home/alex/          standardDir: /home/sara/
    editor: nvim                      editor: code
    fileManager: nautilus             fileManager: dolphin
    projects:                         projects:
      - name: fleetfix                  - name: fleetfix
        absPath:                          absPath:
          /home/alex/projects/             /home/sara/shared/
          fleetfix                         fleetfix
```

Same project, different machines, different paths, different editors — each user
works with their own preferences while sharing the same project data.

If you set up a new machine and want to restore your OperaTree configuration,
copy your `operatree.yaml` to the correct platform location and run
`operatree init` to verify and update any paths that differ on the new machine.

---

_Next: Section 10 — Command Reference_
