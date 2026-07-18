## Usage

Sources:

- [Official extension control file](https://github.com/fdr/pg_connlimit/blob/master/pg_connlimit.control)
- [Official upstream documentation](https://github.com/fdr/pg_connlimit/blob/master/pg_connlimit.c)

`pg_connlimit` — Limits per-role connection counts from filesystem configuration without catalog access.

The reviewed catalog snapshot records version `1.0`, kind `headless`, and implementation language `C`.

This component has no standalone extension DDL in the curated record; integrate it only through the upstream-documented load or provider workflow.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
