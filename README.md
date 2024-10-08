# cobra-cli (cobracli): better Cobra commands

[Use `pkg/cmd` not `cmd`.](https://rwxrob.github.io/zet/2709).

The original `cobra-cli` command creates unfortunately hobbled CLI command scaffolding by dumping all the commands---including subcommand of commands---into the same `cmd` directory. This clutters the code and makes understanding monolith commands with multiple layers of commands and subcommands virtually unusable. This approach also wastes the very useful ability to allow external project to import and single subcommand or even an entire subcommand tree. This monolith composition is now "table stakes" for new Go commands. By putting each command into its own exportable package not only can any subcommand be imported, but unit test code and embedded text files and other resources can be coupled with the subcommand in a clean way. Want to move that subcommand into another monolith repo? Just copy that command directory to the other repo. And all the main functionality of this command is also exported and available to anyone who wants to create their own code to do the same thing without a dependency on the commands themselves.

And who creates commands with dashes in them? I mean, really. This isn't the 90s.

## How is this different?

* Only exported packages at top level (never `main`)
* All packages shared between commands in `internal`
* All commands in `cmd` including main command 
* Exportable commands and subcommands using capitalization
* Initialize complex command sub from nested commands file

