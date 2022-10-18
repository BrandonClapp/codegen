### Proposal

- `gen` will be installable via `go install`
- Run `gen <template>` inside of a directory with `gen.config.json`
- Generator uses variables found in `gen.config.json` to produce output files into an output directory found in config file (relative to config file)
- Generator will read template directory from `gen.config.json`
- Template will be a valid Go template
- The input directory is to be walked, transformed, and then output into the defined output directory found in config file

#### Considerations

- Verify that all variables are present that all files need before overwriting files in output directory
- Ensure that all templates process successfully before overwriting output directory
- Give feedback of missing template variables upon encountering missing variables

#### Questions

- How to change file names?
  - Should be able to do something like name the file `{{VARIABLE_NAME}}.go` and do variable replacement on filenames.
- How will templates support folder names?
  - Folder names should be supported by the same template syntax. `{{SOME_VARIABLE | Snake}}` -> `whatever_var_was/`
  - Alternatively, folders could just use a specific variable. It all depends on the template. `{{DEFINED_FOLDER_NAME}}`
