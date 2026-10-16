## Usage

Sources:

- [README.md](https://github.com/mejedi/planscape-postgresql-module/blob/54e7e6407c35a68741430bb50cb9872a863bd62e/README.md)
- [planscape.cpp](https://github.com/mejedi/planscape-postgresql-module/blob/54e7e6407c35a68741430bb50cb9872a863bd62e/planscape.cpp)
- [pg_hooks.cpp](https://github.com/mejedi/planscape-postgresql-module/blob/54e7e6407c35a68741430bb50cb9872a863bd62e/pg_hooks.cpp)
- [hook_engine.c](https://github.com/mejedi/planscape-postgresql-module/blob/54e7e6407c35a68741430bb50cb9872a863bd62e/hook_engine.c)
- [Makefile](https://github.com/mejedi/planscape-postgresql-module/blob/54e7e6407c35a68741430bb50cb9872a863bd62e/Makefile)
- [planscape.control](https://github.com/mejedi/planscape-postgresql-module/blob/54e7e6407c35a68741430bb50cb9872a863bd62e/planscape.control)

`planscape` is loaded into a PostgreSQL session to record planner alternatives while running EXPLAIN. Its control file identifies version 1.0, but the Makefile does not install extension SQL; use LOAD rather than CREATE EXTENSION.

### Core Workflow

```sql
LOAD 'planscape';
CREATE TEMP TABLE planscape_example (a integer);
EXPLAIN (PLANSCAPE) SELECT avg(a) FROM planscape_example;
```

### Operational Boundaries

The module intercepts planner and reporting functions through runtime code hooks and writes planning data as JSON. `EXPLAIN (PLANSCAPE)` is its documented entry point. Loading the library requires an appropriately privileged session. The low-level hook engine is architecture and server-build sensitive; the repository does not declare a current PostgreSQL support range or license. Use an isolated development environment and review its file-output behavior before testing sensitive queries.
