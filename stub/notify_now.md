## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/notify_now/notify_now-1.0.0/README.md)
- [Official extension control file (notify_now.control)](https://api.pgxn.org/src/notify_now/notify_now-1.0.0/notify_now.control)
- [Official extension SQL (notify_now--1.0.sql)](https://api.pgxn.org/src/notify_now/notify_now-1.0.0/notify_now--1.0.sql)

`notify_now` — This simple extension allows you to return multiple responses from a single query using the built-in PostgreSQL NOTIFY API. There are no additional dependencies. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION notify_now;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `notify_now(text, text)` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
