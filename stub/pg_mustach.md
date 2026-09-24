## Usage

Sources:

- [README](https://api.pgxn.org/src/pg_mustach/pg_mustach-2.0.0/README.md)
- [Control file](https://api.pgxn.org/src/pg_mustach/pg_mustach-2.0.0/pg_mustach.control)
- [SQL](https://api.pgxn.org/src/pg_mustach/pg_mustach-2.0.0/pg_mustach--2.0.sql)

`pg_mustach` 2.0 renders Mustache templates from `jsonb` through libmustach 2 or newer. It adds prepared templates and replaces the older JSON-library-specific adapters with a unified interface.

### Core Workflow

```sql
CREATE EXTENSION pg_mustach;
SELECT mustach('{"name":"PostgreSQL"}'::jsonb, 'Hello {{name}}!');
BEGIN;
SELECT mustach_template('Hello {{name}}!', 'greeting');
SELECT mustach_json('{"name":"Ada"}'::jsonb, tplname := 'greeting');
SELECT mustach_free('greeting');
COMMIT;
```

### Templates and Flags

`mustach` renders directly. `mustach_template` prepares a named template or the unnamed slot; `mustach_json` renders it, and `mustach_free` releases it. Use the named argument `tplname` explicitly: a positional second text argument can resolve to the file-writing overload instead.

`pg_mustach.transaction` defaults to true, so prepared templates are cleared at transaction end. A preparation in one autocommit statement will not survive to the next statement. Set it false only when intentionally managing session-local lifetime, especially with connection pools. `pg_mustach.flags` and `mustach_set_flags` control rendering flags; `mustach_with_*` helpers return flag values.

### Upgrade and Access

The control version is `2.0`, while the PGXN distribution is 2.0.0. Review the shipped upgrade SQL before migrating applications from `json` adapters to `jsonb`. File-output overloads write on the server and require superuser privileges. The extension is relocatable and needs no preload. Bound template size and output volume for untrusted inputs.

Local-file partials are separately controlled by the superuser-only `pg_mustach.whitelist` prefix list. Ordinary roles need an explicit matching grant; a nonempty list also restricts superusers. The whitelist does not authorize file-output overloads.
