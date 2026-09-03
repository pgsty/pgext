## Usage

Sources:

- [Official documentation](https://github.com/LadybugDB/pg_ladybug/blob/895cbdb5d1d0fb7db65e78ac7352795489b91f5d/README.md)
- [Extension control file](https://github.com/LadybugDB/pg_ladybug/blob/895cbdb5d1d0fb7db65e78ac7352795489b91f5d/pg_ladybug.control)
- [Official repository](https://github.com/LadybugDB/pg_ladybug)

`pg_ladybug` Cypher queries over PostgreSQL tables through an embedded Ladybug graph engine.

### Enablement

Install the files for the intended server, then create `pg_ladybug` in the target database:

```sql
CREATE EXTENSION pg_ladybug;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
┌─────────────────────────────────────────────┐
│          PostgreSQL backend                  │
│                                              │
│  Cypher query → ladybug.cypher()             │
│       │                                      │
│       ▼                                      │
│  Ladybug planner (liblbug, compile-time      │
│  linked via -llbug)                          │
│    - Parse Cypher                            │
│    - Discover tables via pg_client           │
│      (libpq → information_schema)            │
│    - Produce EXPLAIN plan                    │
│       │                                      │
│       ▼                                      │
│  extract_pushed_sql()                        │
│    - Parse plan text                         │
│    - Extract "Function:" section (SQL)       │
│    - OR construct SELECT from plan nodes     │
│       │                                      │
│       ▼                                      │
│  SPI_execute(pushed_sql)                     │
│    - Native Postgres execution               │
│    - Return rows via SRF                     │
│                                              │
└─────────────────────────────────────────────┘
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `ladybug.cypher` | FUNCTION | Callable function from the reviewed install surface. |
| `ladybug.explain` | FUNCTION | Callable function from the reviewed install surface. |
| `ladybug.pushed_sql` | FUNCTION | Callable function from the reviewed install surface. |
| `ladybug.register_node` | FUNCTION | Callable function from the reviewed install surface. |
| `ladybug.register_edge` | FUNCTION | Callable function from the reviewed install surface. |
| `ladybug.list_labels` | FUNCTION | Callable function from the reviewed install surface. |
| `ladybug.reset_graph` | FUNCTION | Callable function from the reviewed install surface. |
| `ladybug.sql_query` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 17, 18; do not infer unlisted majors.
- The extension fixes or creates schema objects under `ladybug`; include them in privilege and backup review.
