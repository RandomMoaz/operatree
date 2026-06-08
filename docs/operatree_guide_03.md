# Section 3 — Installation

---

## 3.1 Current Status

OperaTree is currently in **alpha**. It is functional and actively used, but the API, command names, and configuration format may change between releases as the project matures.

Because OperaTree is in alpha, it is not yet available through system package managers such as `apt`, `brew`, `pacman`, or `aur`. Installation is done through one of three methods described below.

To follow releases and stay up to date:
[https://github.com/hanymamdouh82/operatree/releases](https://github.com/hanymamdouh82/operatree/releases)

---

## 3.2 Supported Platforms

| OS      | Architecture          | Supported |
| ------- | --------------------- | --------- |
| Linux   | x86_64 (amd64)        | ✓         |
| Linux   | arm64                 | ✓         |
| macOS   | x86_64 (amd64)        | ✓         |
| macOS   | arm64 (Apple Silicon) | ✓         |
| Windows | x86_64 (amd64)        | ✓         |

---

## 3.3 Method 1 — Install Script (Recommended)

The fastest way to install OperaTree. The script detects your OS and architecture automatically, downloads the latest release binary from GitHub, and installs it to the appropriate location for your platform.

### Linux and macOS

```bash
curl -fsSL https://raw.githubusercontent.com/hanymamdouh82/operatree/main/install.sh | sh
```

The script will:

1. Detect your OS and architecture
2. Fetch the latest release version from GitHub
3. Download the correct binary
4. Attempt to install to `/usr/local/bin` — prompts for `sudo` if needed

**Requirements:** `curl`, `sh`. Both are available by default on Linux and macOS.

### Windows

Open PowerShell and run:

```powershell
irm https://raw.githubusercontent.com/hanymamdouh82/operatree/main/install.ps1 | iex
```

The script will:

1. Fetch the latest release version from GitHub
2. Download the `windows/amd64` binary
3. Install to `%USERPROFILE%\bin`
4. Add `%USERPROFILE%\bin` to your user `PATH` automatically if not already present

**Note:** Restart your terminal after installation for the `PATH` change to take effect.

**Requirements:** PowerShell 5.1 or higher (included with Windows 10 and later).

---

Once complete, verify the installation:

```bash
operatree version
```

---

## 3.4 Method 2 — Using `go install`

If you have Go installed, you can install OperaTree directly from the module path:

```bash
go install github.com/hanymamdouh82/operatree@latest
```

The binary is placed in your `$GOPATH/bin` directory. Make sure this directory is in your `$PATH`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Add this line to your shell profile (`~/.bashrc`, `~/.zshrc`, etc.) to make it permanent.

**Requirements:** Go 1.21 or higher.

---

## 3.5 Method 3 — Build from Source

For contributors or users who want full control over the build:

```bash
git clone https://github.com/hanymamdouh82/operatree.git
cd operatree
make install
```

This compiles the binary and installs it to `/usr/local/bin`. You may be prompted for `sudo` if the directory is not writable by your user.

To build without installing:

```bash
make build
```

The binary is placed in the project root directory.

**Requirements:** Go 1.21 or higher, Git, Make.

---

## 3.6 Verifying the Installation

After installation, confirm everything is working:

```bash
operatree version
```

You should see output similar to:

```
OperaTree v0.1.2
  Commit:     a3f8c21
  Built:      2026-05-20T10:00:00Z
```

If the command is not found, check that the install directory is in your `$PATH`:

```bash
echo $PATH
```

---

## 3.7 Updating

To update to the latest release, run the install script again — it always fetches the latest version:

```bash
curl -fsSL https://raw.githubusercontent.com/hanymamdouh82/operatree/main/install.sh | sh
```

Or if you installed via `go install`:

```bash
go install github.com/hanymamdouh82/operatree@latest
```

---

## 3.8 Uninstalling

To remove OperaTree, delete the binary from your install directory.

**Linux and macOS:**

```bash
rm /usr/local/bin/operatree
```

**Windows:**

```powershell
Remove-Item "$env:USERPROFILE\bin\operatree.exe"
```

Your projects, configuration file, and activity logs are not affected. OperaTree never stores data outside of your project directories and the config file at `~/.config/operatree/operatree.yaml` (or the platform equivalent — see Section 3.9).

To remove the configuration file as well:

```bash
rm -rf ~/.config/operatree/                              # Linux
rm -rf ~/Library/Application\ Support/operatree/         # macOS
rm -rf $XDG_CONFIG_HOME/operatree/                       # if XDG_CONFIG_HOME is set
```

**Windows:**

```powershell
Remove-Item -Recurse "$env:APPDATA\operatree"
```

---

## 3.9 What Gets Installed Where

| Item         | Linux                                                   | macOS                                          | Windows                                        |
| ------------ | ------------------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| Binary       | `/usr/local/bin/operatree`                              | `/usr/local/bin/operatree`                     | `%USERPROFILE%\bin\operatree.exe`              |
| Config file  | `~/.config/operatree/` or `$XDG_CONFIG_HOME/operatree/` | `~/Library/Application Support/operatree/`     | `%APPDATA%\operatree\`                         |
| Project data | Wherever you bootstrap your projects                    | Wherever you bootstrap your projects           | Wherever you bootstrap your projects           |
| Activity log | Inside each project directory (`activity.log`)          | Inside each project directory (`activity.log`) | Inside each project directory (`activity.log`) |

Nothing is written to system directories beyond the binary install location. No background services are installed. No daemons are started.

---

_Next: Section 4 — Getting Started_
