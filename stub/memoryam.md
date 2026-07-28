## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/memoryam/memoryam-0.0.1/README.md)
- [Official extension control file (memoryam.control)](https://api.pgxn.org/src/memoryam/memoryam-0.0.1/memoryam.control)
- [Official extension SQL (memoryam--0.0.1.sql)](https://api.pgxn.org/src/memoryam/memoryam-0.0.1/memoryam--0.0.1.sql)

`memoryam` — MemoryAM is an in-memory TEMPORARY TABLE implementation in C++ of a PostgreSQL storage method. Its mission is to be a simple implementation of a TableAM storage system. As such, we store all changes in memory, and only allow access from a single connection. Use it when an application needs this specific database capability. Upstream explicitly says it is not production-ready.

### Core Workflow

```sql
CREATE EXTENSION memoryam;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `memoryam_relation_details(IN regclass, OUT row_number bigint, OUT xmin integer, OUT xmax integer, OUT is_deleted bool)` is an extension function and returns `SETOF`.
- `memoryam_storage_details(OUT table_name text, OUT row_count bigint, OUT deleted_count bigint, OUT transaction_count bigint)` is an extension function and returns `SETOF`.
- `memoryam_tableam_handler(internal)` is an extension function and returns `table_am_handler`.
- `memoryam` is an extension-defined access method.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Upstream explicitly says the project is not production-ready.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
