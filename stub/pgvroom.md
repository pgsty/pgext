## Usage

Sources:

- [Official upstream README](https://github.com/pgrouting/pgvroom/blob/d38a14de6b8b98fafd7949d427b427e51fb65f2d/README.md)
- [Official extension control file (pgvroom.control)](https://github.com/pgrouting/pgvroom/blob/d38a14de6b8b98fafd7949d427b427e51fb65f2d/sql/pg_controls/pgvroom.control)

`pgvroom` — VROOM functionality reached from the database. Use it for the corresponding spatial data or geospatial workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pgvroom;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- Install the confirmed extension dependencies first: `plpgsql`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
