# OperaTree Demo — Part 3: Power Features
# Run with: vhs demo-3-power.tape
# Requires: vhs, operatree (make install), yazi, git

# -------------------------------------------------------------------
# Output
# -------------------------------------------------------------------

Output ./operatree-demo-3-power.gif
# Output ./operatree-demo-3-power.mp4

# -------------------------------------------------------------------
# Terminal settings
# -------------------------------------------------------------------

Set Shell       "bash"
Set FontSize    14
Set Width       1200
Set Height      700
Set Padding     20
Set Framerate   30
Set WindowBar   Colorful

# -------------------------------------------------------------------
# Edit metadata
# -------------------------------------------------------------------

Type "# Edit subject metadata directly in your editor"
Sleep 2s
Enter
Sleep 500ms

Type "operatree metadata event cairo"
Enter
Sleep 2s

# Select the subject
Enter
Sleep 1s

# nvim opens — simulate quick look and quit
Sleep 2s
Type ":q"
Enter
Sleep 2s

# -------------------------------------------------------------------
# Sync
# -------------------------------------------------------------------

# Type "# Sync project metadata after manual edits"
# Sleep 1s
# Enter
# Sleep 500ms

# Type "operatree sync"
# Enter
# Sleep 2s

# -------------------------------------------------------------------
# Open in yazi
# -------------------------------------------------------------------

Type "# Open subject directory in file manager"
Sleep 1s
Enter
Sleep 500ms

Type "operatree open event cairo"
Enter
Sleep 3s

# Select subject
Enter
Sleep 2s

# yazi opens — navigate briefly then quit
Sleep 2s
Type "q"
Sleep 3s

Sleep 3s

# -------------------------------------------------------------------
# Pipe-friendly output
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Pipe-friendly — works with standard UNIX tools"
Sleep 1s
Enter
Sleep 500ms

Type "operatree describe --plain | grep name"
Enter
Sleep 4s

# -------------------------------------------------------------------
# Activity log
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Append-only audit trail at project root"
Sleep 3s
Enter
Sleep 500ms

Type "cat /tmp/operatree-demo/emissions-validation/activity.log"
Enter
Sleep 5s

Type "clear"
Enter
Sleep 500ms
Type "# Pipe the log with standard tools"
Sleep 3s
Enter
Sleep 500ms

Type "grep CREATE /tmp/operatree-demo/emissions-validation/activity.log | cut -f3 | sort | uniq -c"
Enter
Sleep 4s

# -------------------------------------------------------------------
# Track / untrack
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Untrack and re-track a project"
Sleep 1s
Enter
Sleep 500ms

Type "operatree untrack -d /tmp/operatree-demo/emissions-validation"
Enter
Sleep 4s

Type "operatree track -d /tmp/operatree-demo/emissions-validation"
Enter
Sleep 4s

# -------------------------------------------------------------------
# Version control
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# All data is plain text — perfect for git"
Sleep 1s
Enter
Sleep 500ms

Type "cd /tmp/operatree-demo/emissions-validation && git init && git add -A && git commit -m 'initial project snapshot'"
Enter
Sleep 3s

Type "git log --oneline"
Enter
Sleep 2s

Type "git show --stat HEAD"
Enter
Sleep 3s

# -------------------------------------------------------------------
# Explain
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Built-in directory philosophy guide"
Sleep 3s
Enter
Sleep 500ms

Type "operatree explain"
Enter
Sleep 5s

# -------------------------------------------------------------------
# End
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Part 3 complete — Power Features"
Sleep 1s
Enter
Type "# github.com/hanymamdouh82/operatree"
Sleep 3s
Enter
Sleep 2s
