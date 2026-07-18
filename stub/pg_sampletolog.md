## Usage

Sources:

- [Official extension control file](https://github.com/anayrat/pg_sampletolog/blob/478aa18638c987efba6403e619f2ff41ad0e1d48/pg_sampletolog.control)
- [Official upstream documentation](https://github.com/anayrat/pg_sampletolog/blob/478aa18638c987efba6403e619f2ff41ad0e1d48/README.md)
- [Official PGXN distribution page](https://pgxn.org/dist/pg_sampletolog/)

`pg_sampletolog` — Samples statements or whole transactions into PostgreSQL logs through executor and utility hooks.

The reviewed catalog snapshot records version `2.0.0`, kind `headless`, and implementation language `C`.
The curated compatibility set is `10,11`; confirm the exact build against the target server.

This component has no standalone extension DDL in the curated record; integrate it only through the upstream-documented load or provider workflow.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
