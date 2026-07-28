## Usage

Sources:

- [Official upstream README](https://gitlab.com/byard1/pg_kalman/-/blob/main/README.md)
- [Official extension control file](https://gitlab.com/byard1/pg_kalman/-/blob/main/pg_kalman.control)
- [Official project page](https://gitlab.com/byard1/pg_kalman)

`pg_kalman` — A simple Kalman filter extension for PostgreSQL. Mostly just for fun. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_kalman;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
