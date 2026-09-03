## Usage

Sources:

- [Official documentation](https://github.com/nminoru/pg_plan_tree_dot/blob/13eff984a59ccfeb332685e870195885db02a8e3/README.md)
- [Extension control file](https://github.com/nminoru/pg_plan_tree_dot/blob/13eff984a59ccfeb332685e870195885db02a8e3/pg_plan_tree_dot.control)
- [Official repository](https://github.com/nminoru/pg_plan_tree_dot)

`pg_plan_tree_dot` Graphviz DOT rendering of PostgreSQL execution-plan trees.

### Enablement

Install the files for the intended server, then create `pg_plan_tree_dot` in the target database:

```sql
CREATE EXTENSION pg_plan_tree_dot;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
CREATE EXTENSION pg_plan_tree_dot;
SELECT generate_plan_tree_dot('sql', 'output.dot');
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `public.generate_plan_tree_dot` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 9.0, 9.1, 9.2, 9.3, 9.4, 9.5, 9.6, 10, 11, 12; do not infer unlisted majors.
- Catalog lifecycle is abandoned; test upgrades, dump/restore, and server compatibility before production use.
