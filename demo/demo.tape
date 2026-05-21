# OperaTree Demo
# Run with: vhs demo.tape
# Requires: vhs (https://github.com/charmbracelet/vhs)
#           operatree installed (make install)
#           tree installed (sudo apt-get install tree or brew install tree)

# -------------------------------------------------------------------
# Output
# -------------------------------------------------------------------

Output ./operatree-demo.gif
# Output ./operatree-demo.mp4

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
# 1. Welcome
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "operatree"
Sleep 500ms
Enter
Sleep 2s

# -------------------------------------------------------------------
# 2. Version info
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "operatree version"
Enter
Sleep 2s

# -------------------------------------------------------------------
# 3. Initialize config
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Initialize OperaTree for the first time"
Sleep 1s
Enter
Sleep 500ms

Type "operatree init"
Enter
Sleep 1s

# Standard projects directory
Type "/tmp/operatree-demo"
Sleep 500ms
Enter
Sleep 500ms

# Editor
Type "nvim"
Sleep 300ms
Enter
Sleep 500ms

# File manager
Type "nautilus"
Sleep 300ms
Enter
Sleep 3s

# -------------------------------------------------------------------
# 4. Bootstrap a project
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Bootstrap a new project"
Sleep 1s
Enter
Sleep 500ms

Type "operatree bootstrap emissions-validation"
Enter
Sleep 3s

# Show the filesystem structure after bootstrap
Type "clear"
Enter
Sleep 500ms
Type "# OperaTree creates a standard directory structure on your filesystem"
Sleep 1s
Enter
Sleep 500ms

Type "tree /tmp/operatree-demo/emissions-validation | less"
Enter
Sleep 4s
Space
Type "q"
Enter
Sleep 2s

# -------------------------------------------------------------------
# 5. Set as default project
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Set emissions-validation as the default project"
Sleep 1s
Enter
Sleep 500ms

Type "operatree default"
Enter
Sleep 1s
Sleep 500ms
Down
Sleep 300ms
Enter
Sleep 2s

Type "operatree default --show"
Enter
Sleep 2s

# -------------------------------------------------------------------
# 6. Describe the fresh project
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Describe the project structure"
Sleep 1s
Enter
Sleep 500ms

Type "operatree desc"
Enter
Sleep 4s

# -------------------------------------------------------------------
# 7. Create an event
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Create a new event"
Sleep 1s
Enter
Sleep 500ms

Type "operatree new event"
Enter
Sleep 1s

# Name
Type "Site Visit - Cairo Plant"
Sleep 300ms
Enter
Sleep 500ms

# Date
Type "2026-05-15"
Sleep 300ms
Enter
Sleep 500ms

# Tags
Type "emissions, site-visit, cairo, measurement"
Sleep 300ms
Enter
Sleep 500ms

# Notes
Type "First on-site measurement session at the Cairo plant"
Sleep 300ms
Enter
Sleep 500ms

# Location
Type "Cairo Industrial Zone"
Sleep 300ms
Enter
Sleep 500ms

# Participants
Type "Hany Mamdouh, Foo Bar, Jane Doe"
Sleep 300ms
Enter
Sleep 3s

# Show the structure
Type "clear"
Enter
Sleep 500ms
Type "# Subjects are stored as directories with META.yaml — plain text, git-friendly"
Sleep 1s
Enter
Sleep 500ms

Type "tree /tmp/operatree-demo/emissions-validation/01_EVENTS"
Enter
Sleep 3s

Type "clear"
Enter
Sleep 500ms
Type "# View the subject metadata"
Sleep 1s
Enter
Sleep 500ms

Type "cat /tmp/operatree-demo/emissions-validation/01_EVENTS/*/META.yaml"
Enter
Sleep 4s

# -------------------------------------------------------------------
# 8. Create a task
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Create a new task"
Sleep 1s
Enter
Sleep 500ms

Type "operatree new task"
Enter
Sleep 1s

# Name
Type "Prepare Emissions Report"
Sleep 300ms
Enter
Sleep 500ms

# Date
Type "2026-05-20"
Sleep 300ms
Enter
Sleep 500ms

# Tags
Type "report, emissions, deliverable"
Sleep 300ms
Enter
Sleep 500ms

# Notes
Type "Compile all measurement data into the final client report"
Sleep 300ms
Enter
Sleep 500ms

# Owner
Type "Hany Mamdouh"
Sleep 300ms
Enter
Sleep 500ms

# Status — select in-progress
Down
Sleep 300ms
Enter
Sleep 500ms

# Related event — select first
Enter
Sleep 3s

# -------------------------------------------------------------------
# 9. Create a topic
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Create a knowledge topic"
Sleep 1s
Enter
Sleep 500ms

Type "operatree new topic"
Enter
Sleep 1s

# Name
Type "Emissions Measurement Standards"
Sleep 300ms
Enter
Sleep 500ms

# Date
Type "2026-05-18"
Sleep 300ms
Enter
Sleep 500ms

# Tags
Type "standards, ISO, emissions"
Sleep 300ms
Enter
Sleep 500ms

# Notes
Type "Overview of ISO 14064 and local regulatory standards"
Sleep 300ms
Enter
Sleep 3s

# Show full structure
Type "clear"
Enter
Sleep 500ms
Type "# Everything is just files and directories — no database, no lock-in"
Sleep 1s
Enter
Sleep 500ms

Type "tree /tmp/operatree-demo/emissions-validation"
Enter
Sleep 4s

# -------------------------------------------------------------------
# 10. Project summary
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Project at a glance"
Sleep 1s
Enter
Sleep 500ms

Type "operatree summary"
Enter
Sleep 3s

# -------------------------------------------------------------------
# 11. Fuzzy find
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Find anything across all metadata"
Sleep 1s
Enter
Sleep 500ms

Type "operatree find cairo"
Enter
Sleep 2s
Down
Sleep 500ms
Up
Sleep 500ms
Enter
Sleep 2s

# -------------------------------------------------------------------
# 12. Find with type filter
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Filter by type and search term"
Sleep 1s
Enter
Sleep 500ms

Type "operatree find task report"
Enter
Sleep 2s
Enter
Sleep 2s

# -------------------------------------------------------------------
# 13. Edit metadata
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Edit subject metadata directly in your editor"
Sleep 1s
Enter
Sleep 500ms

Type "operatree metadata event cairo"
Enter
Sleep 2s

# Select the subject
Enter
Sleep 1s

# Editor opens — wait, simulate a quick edit and quit
Sleep 2s
Type ":q"
Enter
Sleep 1s

# Sync after edit
Type "clear"
Enter
Sleep 500ms
Type "# Sync project metadata after manual edits"
Sleep 1s
Enter
Sleep 500ms

Type "operatree sync"
Enter
Sleep 2s

# -------------------------------------------------------------------
# 14. Open subject directory
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Open subject directory in file manager"
Sleep 1s
Enter
Sleep 500ms

Type "operatree open event cairo"
Enter
Sleep 2s
Enter
Sleep 2s

# -------------------------------------------------------------------
# 15. Pipe-friendly output
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Pipe-friendly — works with standard UNIX tools"
Sleep 1s
Enter
Sleep 500ms

Type "operatree desc --plain | grep name"
Enter
Sleep 2s

# -------------------------------------------------------------------
# 16. Activity log
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Append-only audit trail at project root"
Sleep 1s
Enter
Sleep 500ms

Type "cat /tmp/operatree-demo/emissions-validation/activity.log"
Enter
Sleep 3s

Type "clear"
Enter
Sleep 500ms
Type "# Pipe the log with standard tools"
Sleep 1s
Enter
Sleep 500ms

Type "grep CREATE /tmp/operatree-demo/emissions-validation/activity.log | cut -f3 | sort | uniq -c"
Enter
Sleep 2s

# -------------------------------------------------------------------
# 17. Track / untrack
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Track an existing project"
Sleep 1s
Enter
Sleep 500ms

Type "operatree -d /tmp/operatree-demo/emissions-validation track"
Enter
Sleep 2s

Type "operatree -d /tmp/operatree-demo/emissions-validation untrack"
Enter
Sleep 2s

Type "operatree -d /tmp/operatree-demo/emissions-validation track"
Enter
Sleep 2s

# -------------------------------------------------------------------
# 18. Second project — show multi-project workflow
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Bootstrap a second project"
Sleep 1s
Enter
Sleep 500ms

Type "operatree bootstrap analytics-dashboard"
Enter
Sleep 3s

Type "clear"
Enter
Sleep 500ms
Type "# Switch default project"
Sleep 1s
Enter
Sleep 500ms

Type "operatree default"
Enter
Sleep 1s
Down
Sleep 300ms
Enter
Sleep 2s

Type "operatree default --show"
Enter
Sleep 2s

Type "operatree summary"
Enter
Sleep 3s

# Switch back
Type "operatree default"
Enter
Sleep 1s
Enter
Sleep 2s

Type "operatree default --show"
Enter
Sleep 2s

# -------------------------------------------------------------------
# 19. Version control friendly
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# All data is plain text — initialize git and you have full version history"
Sleep 1s
Enter
Sleep 500ms

Type "cd /tmp/operatree-demo/emissions-validation && git init && git add -A && git commit -m 'initial project snapshot'"
Enter
Sleep 3s

Type "git log --oneline"
Enter
Sleep 2s

# -------------------------------------------------------------------
# 20. Explain
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Built-in directory philosophy guide"
Sleep 1s
Enter
Sleep 500ms

Type "operatree explain"
Enter
Sleep 4s

# -------------------------------------------------------------------
# End
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# github.com/hanymamdouh82/operatree"
Sleep 3s
Enter
Sleep 2s
