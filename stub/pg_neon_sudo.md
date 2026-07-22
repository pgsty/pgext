## Usage

Sources:

- [Extension control file](https://github.com/neondatabase/pg_neon_sudo/blob/2fea160a4d68788a0ad57fc3c87d8af24ef2367a/pg_neon_sudo.control)
- [SQL function definitions](https://github.com/neondatabase/pg_neon_sudo/blob/2fea160a4d68788a0ad57fc3c87d8af24ef2367a/pg_neon_sudo--0.0.1.sql)
- [C implementation](https://github.com/neondatabase/pg_neon_sudo/blob/2fea160a4d68788a0ad57fc3c87d8af24ef2367a/pg_neon_sudo.c)

`pg_neon_sudo` 0.0.1 is a Neon-specific privilege bridge for starting and stopping PostgreSQL Anonymizer dynamic masking. Its two functions temporarily execute the corresponding `anon` routine as PostgreSQL's bootstrap superuser. It is specialized infrastructure code, not a general-purpose `sudo` facility.

### Core Workflow

The extension must be installed by a superuser and creates a fixed `sudo` schema:

```sql
CREATE EXTENSION pg_neon_sudo;

SELECT sudo.anon_start_dynamic_masking();
SELECT sudo.anon_stop_dynamic_masking();
```

Both calls return `operation successful` on supported PostgreSQL builds after invoking `anon.start_dynamic_masking()` or `anon.stop_dynamic_masking()`.

### Function Index

- `sudo.anon_start_dynamic_masking()` starts dynamic masking through PostgreSQL Anonymizer.
- `sudo.anon_stop_dynamic_masking()` stops dynamic masking.

### Requirements and Security Boundary

The implementation supports PostgreSQL 16 and later; earlier versions raise a feature-not-supported error. The target `anon` functions must already exist, although the extension control file does not declare PostgreSQL Anonymizer as a dependency. No preload setting is used.

These functions deliberately cross a superuser boundary. Restrict `EXECUTE` to the exact roles that administer masking and audit every grant. The SQL declarations mark the functions `IMMUTABLE`, but they have operational side effects, so do not use them in indexes, generated columns, or planner-foldable expressions. This extension is intended for Neon's controlled environment; portability and security assumptions must be re-evaluated elsewhere.
