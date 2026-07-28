## Usage

Sources:

- [Official upstream README](https://github.com/obartunov/pathprobe/blob/ae53ef0c7d9d266bd8c97c23b35a3aa6ff034143/README.md)
- [Official extension control file (pathprobe.control)](https://github.com/obartunov/pathprobe/blob/ae53ef0c7d9d266bd8c97c23b35a3aa6ff034143/pathprobe.control)
- [Official extension SQL (pathprobe--1.1.sql)](https://github.com/obartunov/pathprobe/blob/ae53ef0c7d9d266bd8c97c23b35a3aa6ff034143/pathprobe--1.1.sql)

`pathprobe` — pathprobe is a diagnostic extension for PostgreSQL planner work. It shows not only the final plan chosen by EXPLAIN, but also the path-level decisions made on the way: which paths were skipped, rejected, accepted, displaced, or survived. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pathprobe;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pathprobe(query text)` is an extension function and returns `text`.
- `pathprobe_json(query text)` is an extension function and returns `text`.
- `pathprobe_propose(query text, spec text)` is an extension function and returns `text`.
- `pathprobe_propose_json(query text, spec text)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
