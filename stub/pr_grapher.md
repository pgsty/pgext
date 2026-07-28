## Usage

Sources:

- [Official upstream README](https://github.com/dalibo/opm/blob/6b757bfe413cb1a21d736cdc648971b4a2416213/README)
- [Official extension control file (pr_grapher.control)](https://github.com/dalibo/opm/blob/6b757bfe413cb1a21d736cdc648971b4a2416213/processes/pr_grapher/pr_grapher.control)
- [Official extension SQL (pr_grapher--1.0--1.1.sql)](https://github.com/dalibo/opm/blob/6b757bfe413cb1a21d736cdc648971b4a2416213/processes/pr_grapher/pr_grapher--1.0--1.1.sql)

`pr_grapher` — Grapher process for OPM. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pr_grapher;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pr_grapher.delete_graph(p_id bigint)` is an extension function and returns `boolean`.
- `pr_grapher.get_categories()` is an extension function and returns `TABLE`.
- `pr_grapher.js_time(timestamptz)` is an extension function and returns `bigint`.
- `pr_grapher.js_timetz(timestamptz)` is an extension function and returns `bigint`.
- `pr_grapher.list_graph()` is an extension function and returns `TABLE`.
- `pr_grapher.categories` is a table installed or managed by the extension.
- `pr_grapher.graph_categories` is a table installed or managed by the extension.
- `pr_grapher.graphs` is a table installed or managed by the extension.
- `pr_grapher.nested_categories` is a table installed or managed by the extension.
- `pr_grapher.series` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.1`.
- Install the confirmed extension dependencies first: `opm_core`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
