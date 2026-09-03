## Usage

Sources:

- [PGXN 2.6.0 README](https://pgxn.org/dist/plphp/2.6.0/README.html)
- [PL/php language reference](https://pgxn.org/dist/plphp/2.6.0/doc/plphp.html)
- [plphp control file](https://api.pgxn.org/src/plphp/plphp-2.6.0/plphp.control)
- [PL/php 2.6.0 changelog](https://api.pgxn.org/src/plphp/plphp-2.6.0/CHANGELOG.md)

`plphp` embeds PHP as an untrusted PostgreSQL procedural language for functions, procedures, triggers, event triggers, and anonymous blocks. Use it only when PHP APIs are worth running with the operating-system privileges of the PostgreSQL server process; it is not a sandboxed language.

### Core Workflow

```sql
CREATE EXTENSION plphp;

CREATE FUNCTION php_square(integer)
RETURNS integer
LANGUAGE plphp
AS $$
    return $args[0] * $args[0];
$$;

SELECT php_square(12);
```

PL/php 2.6 supports PHP 8.1 through 8.4 using the embed SAPI without ZTS and PostgreSQL 11 through 18. Installation and creation require a superuser because the language is deliberately untrusted.

### Main Interfaces

- `spi_exec(...)`, `spi_prepare(...)`, `spi_exec_prepared(...)`, `spi_query(...)`, and cursor helpers access PostgreSQL through SPI.
- `return_next()` emits rows from set-returning functions.
- `spi_commit()` and `spi_rollback()` provide transaction control inside procedures.
- `subtransaction(...)` creates a catchable subtransaction boundary.
- `$_TD` exposes trigger and event-trigger context.
- `$_SHARED` is session-global storage, while `$_SD` is private per-function session storage.
- `plphp.on_init`, `plphp.start_proc`, and `plphp_modules` initialize a session interpreter.

The optional `jsonb_plphp`, `hstore_plphp`, and `bytea_plphp` extensions add native transforms. A function must declare `TRANSFORM FOR TYPE ...` to use a transform; installing a companion does not change every PL/php function automatically.

### Security and Operations

A PL/php function can read or write files, open network connections, invoke PHP facilities, and affect anything accessible to the PostgreSQL operating-system account. Only grant function ownership and creation rights to roles trusted like a server administrator. Apply normal `SECURITY DEFINER` and `search_path` hardening when a privileged wrapper is unavoidable.

SPI query text, dynamic identifiers, session initialization code, loaded modules, and persistent interpreter state all require review. Long-running PHP code blocks a PostgreSQL backend, and memory retained in `$_SHARED` or `$_SD` lasts for the session. Use prepared statements for untrusted values, bound resource use, test cancellation/error paths, and recycle sessions when application code retains excessive state.

Version 2.6 adds per-function `$_SD` and the binary-safe `bytea_plphp` transform. PostgreSQL tracks the control version as `2.6` while the source distribution is `2.6.0`.
