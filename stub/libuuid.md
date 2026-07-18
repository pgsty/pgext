## Usage

Sources:

- [Project README](https://github.com/petere/pglibuuid/blob/e05be1a67b06163c71571b5a9f926c20d70d11d7/README.md)
- [Extension control file](https://github.com/petere/pglibuuid/blob/e05be1a67b06163c71571b5a9f926c20d70d11d7/libuuid.control)
- [Version 1.0 SQL API](https://github.com/petere/pglibuuid/blob/e05be1a67b06163c71571b5a9f926c20d70d11d7/libuuid--1.0.sql)
- [libuuid wrapper implementation](https://github.com/petere/pglibuuid/blob/e05be1a67b06163c71571b5a9f926c20d70d11d7/libuuid.c)

`libuuid` 1.0 is an archived C wrapper around the operating system's libuuid library. It adds `uuid_generate()`, `uuid_generate_random()`, and `uuid_generate_time()`, each returning PostgreSQL's built-in `uuid` type.

### Generate values

```sql
CREATE EXTENSION libuuid;

SELECT uuid_generate_random();
SELECT uuid_generate_time();
SELECT uuid_generate();
```

`uuid_generate_random()` requests a random UUID, while `uuid_generate_time()` requests a time-based UUID. The generic `uuid_generate()` delegates the selection strategy to libuuid, so applications that require a specific UUID version should call an explicit function and verify the result.

### Prefer current PostgreSQL facilities

Time-based UUIDs expose time and node-related information and are more predictable than random UUIDs; do not use them as secrets. Random UUID collision resistance also depends on the host library and entropy source. Generation does not replace a unique constraint.

The repository is archived, has no upgrade path after 1.0, and uses PostgreSQL server internals plus a platform library. Modern PostgreSQL provides `gen_random_uuid()` in core, and newer majors add more built-in UUID facilities. Prefer built-ins for new schemas. Before retiring this extension, compare UUID-version requirements and defaults, change application and column defaults, then verify that no function dependency remains.
