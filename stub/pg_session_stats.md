## Usage

Sources:

- [Official extension control file](https://github.com/RyanMarcus/pg_session_stats/blob/8d114d5b653304a3039741c5cad8fee751bea1b3/pg_session_stats.control)
- [Official upstream documentation](https://github.com/RyanMarcus/pg_session_stats/blob/8d114d5b653304a3039741c5cad8fee751bea1b3/README.md)

`pg_session_stats` — Records per-session and parallel-worker CPU, process status, and I/O data into an external SQLite database.

The reviewed catalog snapshot records version `0.0.1`, kind `headless`, and implementation language `C`.
The curated compatibility set is `12`; confirm the exact build against the target server.

This component has no standalone extension DDL in the curated record; integrate it only through the upstream-documented load or provider workflow.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
