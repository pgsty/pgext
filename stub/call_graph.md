## Usage

Sources:

- [Official README](https://github.com/johto/call_graph/blob/9bcba8927be7374263db9d0dd55ec6f8015664d9/README)
- [Official version 1.4 SQL](https://github.com/johto/call_graph/blob/9bcba8927be7374263db9d0dd55ec6f8015664d9/call_graph--1.4.sql)
- [Official control file](https://github.com/johto/call_graph/blob/9bcba8927be7374263db9d0dd55ec6f8015664d9/call_graph.control)

`call_graph` instruments PostgreSQL function execution and accumulates caller/callee edges that can be rendered as call graphs. It targets PostgreSQL 9.1-era servers and is primarily a diagnostic tool: tracking adds work to every instrumented function call, and optional table-usage tracking is explicitly expensive.

### Core Workflow

Preload the module and restart PostgreSQL before creating the extension:

```conf
shared_preload_libraries = 'call_graph'
```

Enable collection only for the session or transaction being measured. Disable it before consolidating buffer rows:

```sql
CREATE EXTENSION call_graph;

SET call_graph.enable = true;
SELECT application_function();
SET call_graph.enable = false;

SELECT *
FROM call_graph.processcallgraphbuffers(10000);
```

The processing function consumes up to the requested number of buffered records and updates the persistent graph/statistics tables. The repository's `utils/create_graphs.pl` tooling can turn those tables into graph output.

### Configuration and Data

- `call_graph.enable`: enables function-call collection; off by default.
- `call_graph.track_table_usage`: also records relations touched by a graph, with a significant performance penalty.
- `CallGraphBuffer` and `TableAccessBuffer`: unlogged append buffers used to reduce backend lock contention.
- `CallGraphs` and `Edges`: consolidated graph identity and caller/callee data.
- `TableUsage`, `DailyStats`, and `HourlyStats`: relation usage and aggregate statistics.
- `call_graph_version()`: returns the installed extension API version.

### Caveats

Version `1.4` is an old server-hook module; verify its C hooks against the exact PostgreSQL major version before loading it into an instance. Unlogged buffers are not crash durable, and consolidation should be scheduled with collection disabled in that worker to avoid self-generated records.

The installation grants `PUBLIC` insert access to both buffer tables. Upstream explicitly warns that malicious users can inject bogus records. Revoke those grants or isolate trusted roles before treating the graph as diagnostic evidence, and benchmark overhead before enabling table-usage tracking or instance-wide collection.
