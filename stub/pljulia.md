## Usage

Sources:

- [Official README](https://gitlab.com/pljulia/pljulia/-/blob/343b46ef81bfa97f76e6e0b9acbe06ed0dae5e24/README.md)
- [Version 0.8 extension SQL](https://gitlab.com/pljulia/pljulia/-/blob/343b46ef81bfa97f76e6e0b9acbe06ed0dae5e24/pljulia--0.8.sql)
- [Official regression SQL](https://gitlab.com/pljulia/pljulia/-/tree/343b46ef81bfa97f76e6e0b9acbe06ed0dae5e24/sql)

`pljulia` embeds Julia as an untrusted PostgreSQL procedural language. Version 0.8 supports ordinary functions, anonymous blocks, common scalar/array/composite conversions, set-returning functions, session-local shared data, and SPI access. Code runs inside the database backend, so only trusted administrators should define Julia functions.

### Core Workflow

After building against the exact PostgreSQL and Julia installations, create the language and define a function:

```sql
CREATE EXTENSION pljulia;

CREATE FUNCTION julia_filter_nulls(x integer[])
RETURNS integer[]
AS $$
return filter!(el -> el != nothing, x)
$$ LANGUAGE pljulia;

SELECT julia_filter_nulls(ARRAY[1, 2, NULL, 4]);
```

SQL null maps to Julia `nothing`. Booleans, integers, floating values, numeric values, text, arrays, and composites have documented conversions; a composite is represented as a dictionary on input and may be returned as a dictionary or positional tuple.

### Sets, SPI, and Session State

- `return_next` emits rows from a set-returning function.
- `spi_exec(query, limit)` returns a bounded result array or affected-row count.
- `spi_exec(query)` returns a cursor consumed with `spi_fetchrow`; close an incompletely consumed cursor with `spi_cursor_close`.
- `spi_prepare` and `spi_exec_prepared` support parameterized plans.
- `GD` is a Julia dictionary shared by PL/Julia calls only for the lifetime of the current PostgreSQL session.

```sql
CREATE FUNCTION julia_count_rows()
RETURNS bigint
AS $$
rows = spi_exec("SELECT count(*) AS n FROM app.events", 1)
return Int64(rows[1]["n"])
$$ LANGUAGE pljulia;
```

Prefer prepared parameters over interpolating SQL text. SPI work participates in the caller's PostgreSQL transaction.

### Experimental Boundaries

The README labels the project work in progress. It documents basic `DO` support and describes triggers/event triggers, but still tells readers to use a separate branch for regular trigger support while the current main tree contains trigger regression SQL. Treat triggers as experimental and verify the exact checked-out revision rather than assuming the prose and source expose the same feature set.

The control file installs the language in `pg_catalog`, is non-relocatable, and does not require preload. The language is created without `TRUSTED`, so it has the security boundary of an untrusted procedural language. Julia libraries, allocations, native code, crashes, and blocking calls share the backend process.

Version 0.8 requires Julia and its headers/libraries to match the build and runtime host. Pin Julia, PostgreSQL, PL/Julia, packages, and native dependencies; test backend lifecycle, parallelism, memory growth, exceptions, encoding, upgrades, and dump/restore. Do not expose function creation or mutable Julia package loading to untrusted roles.
