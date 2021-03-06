# Manifest contains metadata about the driver. To learn about the different
# values in this manifest refer to:
#   https://github.com/bblfsh/sdk/blob/master/driver/manifest/manifest.go

# human-readable language name
name = "{{.Language}}"
# language identifier
language = "{{.Language}}"

# status describes the current development status of the driver, the valid
# values for status are: `planning`, `pre-alpha`, `alpha`, `beta`, `stable`,
# `mature` or `inactive`.
status = "planning"
features = ["ast"]

# documentation block is use to render the README.md file.
[documentation]
description  = """
"""
