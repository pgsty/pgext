## Usage

Sources:

- [Official upstream README](https://github.com/hill/pg_tamagotchi/blob/bf37d7bd2fb6d269a3f7a8202c73bf3b0ba4bb15/README.md)
- [Official extension control file (pg_tamagotchi.control)](https://github.com/hill/pg_tamagotchi/blob/bf37d7bd2fb6d269a3f7a8202c73bf3b0ba4bb15/pg_tamagotchi.control)
- [Official extension SQL (pg_tamagotchi--0.1.0.sql)](https://github.com/hill/pg_tamagotchi/blob/bf37d7bd2fb6d269a3f7a8202c73bf3b0ba4bb15/pg_tamagotchi--0.1.0.sql)

`pg_tamagotchi` — A tamagotchi that lives in your Postgres database. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_tamagotchi;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `feed(food text DEFAULT NULL)` is an extension function and returns `text`.
- `hatch(name text DEFAULT NULL)` is an extension function and returns `text`.
- `status()` is an extension function and returns `text`.
- `talk(message text DEFAULT NULL)` is an extension function and returns `text`.
- `vitals` is an extension-defined view.
- `message` is a table installed or managed by the extension.
- `pet` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
