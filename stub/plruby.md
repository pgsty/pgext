## Usage

Sources:

- [PL/Ruby v2.5.0 README](https://github.com/commandprompt/plruby/blob/v2.5.0/README.md)
- [PL/Ruby language reference](https://github.com/commandprompt/plruby/blob/v2.5.0/doc/plruby.md)
- [PL/Ruby cookbook](https://github.com/commandprompt/plruby/blob/v2.5.0/doc/cookbook.md)
- [PL/Ruby v2.5.0 control file](https://github.com/commandprompt/plruby/blob/v2.5.0/plruby.control)
- [PL/Ruby changelog](https://github.com/commandprompt/plruby/blob/v2.5.0/CHANGELOG.md)

`plruby` is the maintained Command Prompt procedural-language extension that embeds Ruby 3 in PostgreSQL. Package release 2.5.0 installs SQL extension version `2.5`. It supports scalar and set-returning functions, triggers, event triggers, procedures, anonymous `DO` blocks, SPI queries, cursors, and prepared plans.

### Create a Function

```sql
CREATE EXTENSION plruby;

CREATE FUNCTION ruby_add(integer, integer)
RETURNS integer
LANGUAGE plruby
AS $$
  args[0] + args[1]
$$;

SELECT ruby_add(2, 3);
```

Arguments are exposed through `args`; Ruby's final expression becomes the SQL return value. PostgreSQL scalar, array, composite, and record conversion rules are documented in the language reference.

### Set-Returning Functions

Use `return_next` to emit rows from a set-returning function:

```sql
CREATE FUNCTION ruby_series(integer)
RETURNS SETOF integer
LANGUAGE plruby
AS $$
  1.upto(args[0]) { |n| return_next(n) }
$$;

SELECT * FROM ruby_series(3);
```

### SPI and Database Work

PL/Ruby exposes PostgreSQL's Server Programming Interface for SQL execution, prepared plans, and cursors. Keep SQL values in parameters rather than interpolating them into command text, and release long-lived cursors or prepared state when the session no longer needs them.

Procedures can use the documented transaction-control surface where PostgreSQL permits `COMMIT` or `ROLLBACK`. Functions and triggers remain subject to PostgreSQL's normal transactional restrictions.

### Triggers and Session State

Trigger functions receive trigger metadata through `$_TD` and return the row action documented by PL/Ruby. Event triggers, anonymous `DO` blocks, backend-local session data, and shared data are also available. These features run inside the database backend, so an exception, blocking call, or memory leak directly affects that backend.

### Version 2.5.0

- `bytea` now maps to a raw, NUL-safe Ruby `String` with `ASCII-8BIT` encoding instead of PostgreSQL hex text. This is a breaking conversion change: audit functions that parse or construct `\x...` strings and build bytes explicitly, for example with `Array#pack`.
- `$_SD` adds per-function state that persists across calls in one session and resets when the function is recompiled. `$_SHARED` remains session-wide across PL/Ruby functions.
- `spi_colnames`, `spi_coltypes`, and `spi_coltypmods` expose result-column metadata, and `ltree_plruby` adds the opt-in `ltree` transform.
- After installing the 2.5.0 shared library and SQL files, run `ALTER EXTENSION plruby UPDATE` in each database that already has the extension.

### Security and Requirements

- `plruby` is an untrusted language. Ruby 3 provides no safe in-process sandbox, so creating PL/Ruby functions is restricted to superusers and code executes with the PostgreSQL server process's operating-system authority.
- Review all PL/Ruby source as privileged server code. Never allow tenants or ordinary application roles to submit arbitrary Ruby.
- Upstream v2.5.0 supports PostgreSQL 11-18 and Ruby 3.x. Current Pigsty packages target PostgreSQL 14-18.
- No `shared_preload_libraries` setting is required. Existing sessions must reconnect after server-side library replacement before assuming a new runtime is active.
- `jsonb_plruby`, `hstore_plruby`, and `ltree_plruby` are companion transforms. A function must explicitly declare `TRANSFORM FOR TYPE ...` to receive native Ruby structures instead of the normal datum wrapper/conversion path.
