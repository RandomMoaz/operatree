# OperaTree Demo — Part 1: Setup & Structure
# Run with: vhs demo-1-setup.tape
# Requires: vhs, operatree (make install), tree

# -------------------------------------------------------------------
# Output
# -------------------------------------------------------------------

Output ./operatree-demo-1-setup.gif
# Output ./operatree-demo-1-setup.mp4

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
# Welcome
# -------------------------------------------------------------------

Type "operatree"
Enter
Sleep 3s

# -------------------------------------------------------------------
# Version
# -------------------------------------------------------------------

Type "operatree version"
Enter
Enter
Enter
Sleep 3s

# -------------------------------------------------------------------
# Init
# -------------------------------------------------------------------

Type "# Initialize OperaTree for the first time"
Sleep 1s
Enter
Sleep 500ms

Type "operatree init"
Enter
Sleep 2s

Type "/tmp/operatree-demo"
Enter
Sleep 3s

# Type "nvim"
# Enter
# Sleep 500ms

Type "yazi"
Sleep 3s
Enter
Sleep 3s

# -------------------------------------------------------------------
# Bootstrap
# -------------------------------------------------------------------

Type "# Bootstrap a new project"
Sleep 2s
Enter
Sleep 1500ms

Type "operatree bootstrap emissions-validation"
Sleep 1s
Enter
Sleep 4s

# -------------------------------------------------------------------
# Show filesystem structure
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# OperaTree creates a standard directory structure on your filesystem"
Sleep 2s
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
# Set default
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Set emissions-validation as the default project"
Sleep 1s
Enter
Sleep 1500ms

Type "operatree default"
Enter
Sleep 1s
Down
Sleep 300ms
Enter
Sleep 3s

Type "operatree default --show"
Enter
Sleep 4s

# -------------------------------------------------------------------
# Describe
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Describe the project structure"
Sleep 1s
Enter
Sleep 500ms

Type "operatree describe"
Enter
Sleep 4s

# -------------------------------------------------------------------
# End
# -------------------------------------------------------------------

Type "clear"
Enter
Sleep 500ms
Type "# Part 1 complete — Setup & Structure"
Sleep 1s
Enter
Type "# github.com/hanymamdouh82/operatree"
Sleep 3s
Enter
Sleep 2s
