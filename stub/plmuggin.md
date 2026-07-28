## Usage

Sources:

- [Official upstream README](https://github.com/tatut/plmuggin/blob/77a0c4d35fa1441e3b86dec06ae75ffaa610281d/README.md)
- [Official extension control file (plmuggin.control)](https://github.com/tatut/plmuggin/blob/77a0c4d35fa1441e3b86dec06ae75ffaa610281d/plmuggin.control)
- [Official extension SQL (plmuggin--0.1.sql)](https://github.com/tatut/plmuggin/blob/77a0c4d35fa1441e3b86dec06ae75ffaa610281d/plmuggin--0.1.sql)

`plmuggin` — PL/Muggin is a HTML template engine inspired by pug defined as a PostgreSQL language handler. Muggin templates are whitespace sensitive and look like this:. Use it when database code must run in or interoperate with this procedural language. Upstream explicitly says it is not production-ready.

### Core Workflow

```sql
CREATE EXTENSION plmuggin;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `plmuggin_call_handler()` is an extension function and returns `language_handler`.
- `plmuggin_get_metadata(template_name TEXT, meta_key TEXT)` is an extension function and returns `TEXT`.
- `plmuggin_templates()` is an extension function and returns `SETOF`.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- The control file marks the extension as trusted.
- Upstream explicitly says the project is not production-ready.
- Upstream describes the project as a work in progress.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
