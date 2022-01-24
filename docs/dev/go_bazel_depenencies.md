# Managing Go Dependencies

1. Install new module with go get i.e.
   `
   go get package
   `

2. Sync go.mod with GoLAnd IDE using context menu. Alternatively, update go.mod manually

3. Run a rebuild
   `
   make rebuild
   `

Rebuild does the following

* syncs go.mod with Bazel WORKSPACE,
* Generates new or updates existing build files as detected by Bazel & Gazelle
* Checks for breaking changes
* Runs a build

The exact build commands used by make rebuild are:

*Convert mod dependencies into bazel dependencies*

bazel run //:gazelle -- update-repos -from_file=go.mod

*Update all build files & dependencies*

bazel run //:gazelle

*Build all sources*

bazel build //:build_local

Each of these steps should run just fine after having added new dependencies. 