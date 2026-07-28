## Usage

Sources:

- [Official upstream README](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/README.md)
- [Official extension control file (pg_sequence.control)](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_sequence/pg_sequence.control)
- [Official implementation source](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_sequence/src/lib.rs)

`pg_sequence` — ERP document numbering with formatted, scoped, auto-incrementing sequences. Use it when SQL needs these specialized functions or aggregates. Upstream explicitly says it is not production-ready.

### Core Workflow

```sql
CREATE EXTENSION pg_sequence;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.2.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Upstream explicitly says the project is not production-ready.
- Upstream describes the project as a proof of concept.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
