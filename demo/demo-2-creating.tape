# OperaTree Demo — Part 2: Creating & Finding
# Run with: vhs demo-2-creating.tape
# Requires: vhs, operatree (make install), tree
# Note: assumes Part 1 was recorded and default project is set

# -------------------------------------------------------------------
# Output
# -------------------------------------------------------------------

Output ./operatree-demo-2-creating.gif
Output ./operatree-demo-2-creating.mp4

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
# Create event
# -------------------------------------------------------------------

Type "# Create a new event"
Sleep 1s
Enter
Sleep 1500ms

Type "operatree new event"
Enter
Sleep 2s

Type "Site Visit - Cairo Plant"
Enter
Sleep 300ms

Type "2026-05-15"
Enter
Sleep 300ms

Type "emissions, site-visit, cairo, measurement"
Enter
Sleep 300ms

Type "First on-site measurement session at the Cairo plant"
Enter
Sleep 300ms

Type "Cairo Industrial Zone"
Enter
Sleep 300ms

Type "Hany Mamdouh, Foo Bar, Jane Doe"
Enter
Sleep 3s

# Show the subject metadata
Type "# Subjects are plain YAML — readable without any tool"
Sleep 2s
Enter
Sleep 500ms

Type "cat /tmp/operatree-demo/emissions-validation/01_EVENTS/*/METADATA.yml"
Enter
Sleep 4s

# -------------------------------------------------------------------
# Create task
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

Type "Prepare Emissions Report"
Enter
Sleep 300ms

Type "2026-05-20"
Enter
Sleep 300ms

Type "report, emissions, deliverable"
Enter
Sleep 300ms

Type "Compile all measurement data into the final client report"
Enter
Sleep 300ms

Type "Hany Mamdouh"
Enter
Sleep 300ms

# Status — select in-progress
Down
Sleep 300ms
Enter
Sleep 300ms

# Related event
Enter
Sleep 3s

# -------------------------------------------------------------------
# Create topic
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

Type "Emissions Measurement Standards"
Enter
Sleep 300ms

Type "2026-05-18"
Enter
Sleep 300ms

Type "standards, ISO, emissions"
Enter
Sleep 300ms

Type "Overview of ISO 14064 and local regulatory standards"
Enter
Sleep 3s

# Show full structure
Type "clear"
Enter
Sleep 500ms
Type "# Everything is files and directories — no database, no lock-in"
Sleep 1s
Enter
Sleep 500ms

Type "tree /tmp/operatree-demo/emissions-validation"
Enter
Sleep 4s

# -------------------------------------------------------------------
# Summary
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Project at a glance"
Sleep 2s
Enter
Sleep 500ms

Type "operatree summary"
Enter
Sleep 3s

# -------------------------------------------------------------------
# Fuzzy find
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Find anything across all metadata"
Sleep 3s
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
Sleep 4s

# -------------------------------------------------------------------
# Find with type filter
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Filter by type and search term"
Sleep 2s
Enter
Sleep 500ms

Type "operatree find task report"
Enter
Sleep 3s
Enter
Sleep 4s

# -------------------------------------------------------------------
# End
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Part 2 complete — Creating & Finding"
Sleep 1s
Enter
Type "# github.com/hanymamdouh82/operatree"
Sleep 3s
Enter
Sleep 2s
