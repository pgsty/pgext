## Usage

Sources:

- [Official upstream README](https://github.com/tyrchen/rust-training/blob/dd08bcb50cffdc4f1d715afdaf83af05eb76705c/README.md)
- [Official extension control file (db_assist.control)](https://github.com/tyrchen/rust-training/blob/dd08bcb50cffdc4f1d715afdaf83af05eb76705c/live_coding/db_assist/db_assist.control)
- [Official implementation source](https://github.com/tyrchen/rust-training/blob/dd08bcb50cffdc4f1d715afdaf83af05eb76705c/live_coding/db_assist/src/lib.rs)

`db_assist` — Minimal pgrx training extension exposing a hello_db_assist function. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION db_assist;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hello_db_assist()` is an extension function.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
