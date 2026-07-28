## Usage

Sources:

- [Official upstream README](https://github.com/pgrouting/pgorpy/blob/9ec1b516e1358d6fbf05dfaf0da93f10dd1d766f/README.md)
- [Official extension control file (pgorpy.control)](https://github.com/pgrouting/pgorpy/blob/9ec1b516e1358d6fbf05dfaf0da93f10dd1d766f/sql/pg_controls/pgorpy.control)

`pgorpy` — OR-tools python accessible from the database. Use it when SQL needs these specialized functions or aggregates. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pgorpy;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- Install the confirmed extension dependencies first: `plpython3u`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
