# OperaTree Demo
# Run with: vhs demo.tape
# Requires: vhs (https://github.com/charmbracelet/vhs)
#           operatree installed (make install)
#           tree installed (sudo apt-get install tree or brew install tree)


# -------------------------------------------------------------------
# Output
# -------------------------------------------------------------------

Output ./operatree-demo.gif
Output ./operatree-demo.mp4

# -------------------------------------------------------------------
# Terminal settings
# -------------------------------------------------------------------

Set Shell       "bash"
Set FontSize    14
# Set FontFamily  "JetBrains Mono"
Set Width       1200
Set Height      700
# Set Theme       "Dracula"
Set Padding     20
Set Framerate   30
Set WindowBar Colorful

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
# 2. Initialize config
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Sleep 1s
Type "# Initialize OperaTree for the first time"
Sleep 1s
Enter
Sleep 500ms

Type "operatree init"
Enter
Sleep 1s

# Answer the init form: standard projects directory
Type "/tmp/operatree-demo"
Sleep 500ms
Enter
Sleep 3s

# -------------------------------------------------------------------
# 3. Bootstrap a project
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Sleep 1s
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
Sleep 4s

# -------------------------------------------------------------------
# 4. Set as default project (using interactive selection)
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Sleep 1s
Type "# Set emissions-validation as the default project"
Sleep 1s
Enter
Sleep 500ms

Type "operatree default"
Enter
Sleep 1s

# Select emissions-validation from the list
Sleep 500ms
Down
Sleep 300ms
Enter
Sleep 2s

# Verify default project is set
Type "clear"
Enter
Sleep 500ms
Type "# Verify the default project"
Sleep 1s
Enter
Sleep 500ms

Type "operatree default --show"
Enter
Sleep 3s

# -------------------------------------------------------------------
# 5. Describe the fresh project (using default)
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Sleep 1s
Type "# Describe the project structure"
Sleep 1s
Enter
Sleep 500ms

Type "operatree desc"
Enter
Sleep 4s

# -------------------------------------------------------------------
# 6. Create an event
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Sleep 1s
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

# Date - manual entry
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

# Show the event was created as a markdown file
Type "clear"
Enter
Sleep 500ms
Type "# Events are stored as markdown files - easily editable, version controllable"
Sleep 1s
Enter
Sleep 500ms

Type "tree /tmp/operatree-demo/emissions-validation/events"
Enter
Sleep 3s

Type "clear"
Enter
Sleep 500ms
Type "# View the event content"
Sleep 1s
Enter
Sleep 500ms

Type "head -20 /tmp/operatree-demo/emissions-validation/events/*.md"
Enter
Sleep 4s

# -------------------------------------------------------------------
# 7. Create a task
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Sleep 1s
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

# Date - manual entry
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

# Show tasks directory
Type "clear"
Enter
Sleep 500ms
Type "# Tasks are also markdown files with metadata"
Sleep 1s
Enter
Sleep 500ms

Type "tree /tmp/operatree-demo/emissions-validation/tasks"
Enter
Sleep 3s

# -------------------------------------------------------------------
# 8. Create a topic
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Sleep 1s
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

# Date - manual entry
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

# Show the complete filesystem structure
Type "clear"
Enter
Sleep 500ms
Type "# Everything is just files and directories - no database, no lock-in"
Sleep 1s
Enter
Sleep 500ms

Type "tree /tmp/operatree-demo/emissions-validation"
Enter
Sleep 4s

# -------------------------------------------------------------------
# 9. Project summary
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Sleep 1s
Type "# Project at a glance"
Sleep 1s
Enter
Sleep 500ms

Type "operatree summary"
Enter
Sleep 3s

# -------------------------------------------------------------------
# 10. Fuzzy find — open browser
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Sleep 1s
Type "# Find anything across all metadata"
Sleep 1s
Enter
Sleep 500ms

Type "operatree find cairo"
Enter
Sleep 2s

# Navigate the finder — just show it working
Sleep 1s
Down
Sleep 500ms
Up
Sleep 500ms
Enter
Sleep 2s

# -------------------------------------------------------------------
# 11. Find with type filter
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Sleep 1s
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
# 12. Explain
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Sleep 1s
Type "# Built-in directory philosophy guide"
Sleep 1s
Enter
Sleep 500ms

Type "operatree explain"
Enter
Sleep 3s

# -------------------------------------------------------------------
# 13. Pipe-friendly output
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Sleep 1s
Type "# Works with standard UNIX tools"
Sleep 1s
Enter
Sleep 500ms

Type "operatree desc --plain | grep name"
Enter
Sleep 2s

# -------------------------------------------------------------------
# 14. Bonus: Set a different default project
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Sleep 1s
Type "# Switch to a different default project"
Sleep 1s
Enter
Sleep 500ms

Type "# First, bootstrap another project"
Sleep 1s
Enter
Sleep 500ms

Type "operatree bootstrap analytics-dashboard"
Enter
Sleep 3s

Type "clear"
Enter
Sleep 500ms
Type "# Set analytics-dashboard as the new default"
Sleep 1s
Enter
Sleep 500ms

Type "operatree default"
Enter
Sleep 1s

# Select analytics-dashboard from the list
Down
Sleep 300ms
Enter
Sleep 2s

Type "clear"
Enter
Sleep 500ms
Type "# Verify the switch"
Sleep 1s
Enter
Sleep 500ms

Type "operatree default --show"
Enter
Sleep 2s

Type "clear"
Enter
Sleep 500ms
Type "# Show its summary"
Sleep 1s
Enter
Sleep 500ms

Type "operatree summary"
Enter
Sleep 3s

# -------------------------------------------------------------------
# 15. Switch back to original project
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Sleep 1s
Type "# Switch back to emissions-validation"
Sleep 1s
Enter
Sleep 500ms

Type "operatree default"
Enter
Sleep 1s

# Select emissions-validation from the list (first item)
Enter
Sleep 2s

Type "clear"
Enter
Sleep 500ms
Type "# Back to original project"
Sleep 1s
Enter
Sleep 500ms

Type "operatree default --show"
Enter
Sleep 2s

# -------------------------------------------------------------------
# 16. Demonstrate version control friendly
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Sleep 1s
Type "# All data is plain text - perfect for git"
Sleep 1s
Enter
Sleep 500ms

Type "cd /tmp/operatree-demo/emissions-validation && ls -la"
Enter
Sleep 2s

Type "clear"
Enter
Sleep 500ms
Type "# No hidden databases, no binary files - just your content"
Sleep 1s
Enter
Sleep 2s

# -------------------------------------------------------------------
# End
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Sleep 1s
Type "# github.com/hanymamdouh82/operatree"
Sleep 3s
Enter
Sleep 2s
