## Usage

Sources:

- [Official extension control file (diag_planner.control)](https://github.com/masahikosawada/incubator/blob/ae553ea4cc728b2d7a742a7ecf721996b3ada8b5/diag_planner/diag_planner.control)
- [Official implementation source](https://github.com/masahikosawada/incubator/blob/ae553ea4cc728b2d7a742a7ecf721996b3ada8b5/diag_planner/diag_planner.c)

`diag_planner` is a lightweight planner-diagnostics module from an incubator repository. Loading the library installs relation- and join-path hooks that emit `NOTICE` messages describing candidate scan or join paths and their estimated costs.

### Core Workflow

Build and install the `diag_planner` library with the matching PostgreSQL server headers, then load it only in an isolated diagnostic session through the server's supported library-loading mechanism. Run a representative query and inspect the emitted scan and join path notices.

The reviewed source does not install SQL objects and does not document a standalone `CREATE EXTENSION` command.

### Diagnostic Output

- Scan notices distinguish sequential, sample, index, index-only, bitmap-index, and bitmap-heap paths.
- Join notices distinguish hash, merge, and nested-loop candidates.
- Each reported path includes the planner's startup and total cost estimates.

### Requirements and Caveats

- The reviewed control, registry, or catalog evidence identifies version `1.0`.
- The control file marks the extension as non-relocatable.
- The module hooks PostgreSQL planner internals and comes from an incubator tree; compatibility is tied to the exact server source it was built against.
- It emits output during planning and is intended for diagnostics, not routine production workloads.
