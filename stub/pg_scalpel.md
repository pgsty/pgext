## Usage

Sources:

- [Upstream source 1](https://github.com/dbgreen/pg_scalpel/blob/3233758db18b2b664360037a43098a114799846f/README.md)
- [Upstream source 2](https://github.com/dbgreen/pg_scalpel/blob/3233758db18b2b664360037a43098a114799846f/pg_scalpel.control)

`pg_scalpel` 0.1 exposes a Janet interpreter inside a PostgreSQL 19 Linux x86-64 backend for debugging and instrumentation.

### Core Workflow

`janet_eval(text)` returns the last Janet value as text, with a VM lasting for the connection. The scalpel primitives resolve symbols, inspect structures, call C functions, install hooks and examine machine code. Static symbols and layouts require matching DWARF debug information.

```sql
CREATE EXTENSION pg_scalpel;
SELECT janet_eval('(+ 40 2)');
```

### Operational Boundaries

No preload or server restart is required. Installation and execution are enforced as superuser-only; revoking PUBLIC access is not a substitute for the runtime check. Janet and native primitives can execute OS commands, read/write arbitrary memory and patch code. Use disposable diagnostic environments and never treat this interpreter as a sandbox.
