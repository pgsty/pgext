## Usage

Sources:

- [Official upstream README](https://github.com/heterodb/toybox/blob/e4e9b0526432809bc2f6d50df35b23237ad6c74e/README.md)
- [Official extension control file (g_cube.control)](https://github.com/heterodb/toybox/blob/e4e9b0526432809bc2f6d50df35b23237ad6c74e/g_cube/g_cube.control)
- [Official implementation source](https://github.com/heterodb/toybox/blob/e4e9b0526432809bc2f6d50df35b23237ad6c74e/g_cube/g_cube.c)

`g_cube` is a PG-Strom plug-in that teaches PG-Strom how to represent and execute selected operations for PostgreSQL `cube` and `earth` values on a GPU. It augments existing types rather than installing a new user-facing type.

### Core Workflow

Build the library and its CUDA fat binaries against the same PostgreSQL and PG-Strom installation, then load it through the PG-Strom plug-in mechanism. Install `cube`—and `earthdistance` when `earth` values are needed—before testing GPU execution.

The reviewed source registers its descriptor from `_PG_init`; it does not provide versioned SQL or a standalone `CREATE EXTENSION g_cube` workflow.

### Accelerated Surface

- `cube_contains(cube, cube)`
- `cube_contained(cube, cube)`
- `cube_ll_coord(cube, integer)`
- casts between `cube` and the `earth` domain
- Arrow decoding for `cube` values

### Requirements and Caveats

- The reviewed control, registry, or catalog evidence identifies version `1.0`.
- The control file marks the extension as relocatable.
- The Makefile requires `nvcc`, PG-Strom server headers, and GPU architectures selected at build time.
- The C API calls PG-Strom's users-extra interface directly, so PostgreSQL, PG-Strom, CUDA, and this plug-in must be version-compatible.
- Confirm CPU fallback, numeric equivalence, device availability, and Arrow error handling on the target build.
