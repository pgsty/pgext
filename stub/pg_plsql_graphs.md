## Usage

Sources:

- [Official extension control file](https://github.com/KLIEBHAN/pg_plsql_graphs/blob/22b2eaef35a90234aa364aee8bf21f600c639a0e/pg_plsql_graphs.control)
- [Official upstream documentation](https://github.com/KLIEBHAN/pg_plsql_graphs/blob/22b2eaef35a90234aa364aee8bf21f600c639a0e/README.md)

`pg_plsql_graphs` 1.0 captures control-flow and data-dependence graphs when PL/pgSQL functions run, exposes the DOT text through views, and lets a client export it for Graphviz. It is an old experimental instrumentation module, not a static analyzer for every function in a database.

### Core Workflow

The documented build integrates iGraph into the PostgreSQL source tree. Preload the module, restart PostgreSQL, and create the extension.

```conf
shared_preload_libraries = 'pg_plsql_graphs'
```

```sql
CREATE EXTENSION pg_plsql_graphs;

SELECT my_plpgsql_function();
SELECT * FROM pg_plsql_graphs;

\COPY (SELECT * FROM pg_plsql_last_flowgraph_dot) TO 'flow.dot'
\COPY (SELECT * FROM pg_plsql_last_pdgs_dot) TO 'pdg.dot'
```

`pg_plsql_graphs` shows indented flow and dependence graphs for called functions; `pg_plsql_graphs_trimmed` is intended for export. The two `pg_plsql_last_*` views expose the most recently captured graphs.

### Caveats

The README targets PostgreSQL 9.3-era sources and iGraph 0.7.1, and requires invasive server build changes plus a preload hook. Captured state is in memory and tied to function execution. Treat it as isolated research tooling: test server compatibility and hook interactions, avoid sensitive function bodies, and do not rely on it as durable audit data.
