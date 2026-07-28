## Usage

Sources:

- [Official upstream README](https://github.com/pgrouting/vrprouting/blob/b34a8c323b8a0e3a81f7b2e5f4261cbd8634c929/README.md)
- [Official extension control file (vrprouting.control)](https://github.com/pgrouting/vrprouting/blob/b34a8c323b8a0e3a81f7b2e5f4261cbd8634c929/sql/pg_controls/vrprouting.control)

`vrprouting` — Vehicle Routing Problems on the Database. Use it for the corresponding spatial data or geospatial workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION vrprouting;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.5.0`.
- Install the confirmed extension dependencies first: `plpgsql`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
