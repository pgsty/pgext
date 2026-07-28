## Usage

Sources:

- [Official upstream README](https://github.com/zqqiang/node-cad/blob/739ff348b4d2c77b275c3a0fe87682c14ffd8181/README.md)
- [Official extension control file (cadfix.control)](https://github.com/zqqiang/node-cad/blob/739ff348b4d2c77b275c3a0fe87682c14ffd8181/db/cadfix/cadfix.control)
- [Official implementation source](https://github.com/zqqiang/node-cad/blob/739ff348b4d2c77b275c3a0fe87682c14ffd8181/db/cadfix/cadfix.c)

`cadfix` is the PostgreSQL-side C library in the historical `node-cad` prototype. It loads CAD data through Open CASCADE and exposes helpers used by the accompanying Node.js application; it is not a self-contained modern PostgreSQL extension.

### Core Workflow

The upstream README installs the `cadfix` shared library and registers its entry points manually:

```sql
CREATE FUNCTION cadinit(cstring)
RETURNS integer AS 'cadfix' LANGUAGE C;

CREATE FUNCTION full_edge(cstring, cstring)
RETURNS integer AS 'cadfix' LANGUAGE C;

SELECT cadinit('path/to/cad.step');
SELECT full_edge('evaluate', 'path/to/evaluate.csv');
SELECT full_edge('import', 'path/to/import.csv');
```

`cadinit` opens a CAD file. `full_edge` supports the upstream `evaluate` and `import` modes for edge data.

### Requirements and Caveats

- The reviewed control file identifies version `1.0`, names `$libdir/cadfix`, and is non-relocatable, but the repository does not provide versioned extension SQL.
- Creating C-language functions and loading an arbitrary server library requires elevated privileges.
- The documented build targets PostgreSQL 9.5.1, Visual Studio 2012, 32-bit Node.js 4.3.2, and a bundled Open CASCADE workflow. Treat it as historical source, not evidence of compatibility with current PostgreSQL.
- File paths are interpreted by the server process. Validate ownership, permissions, file formats, and failure handling before testing with untrusted CAD input.
