## Usage

Sources:

- [Official extension control file (json_model.control)](https://api.pgxn.org/src/json_model/json_model-2.0.0-alpha3/json_model.control)
- [Official extension SQL (json_model--2.0.sql)](https://api.pgxn.org/src/json_model/json_model-2.0.0-alpha3/json_model--2.0.sql)

`json_model` — JSON Model PL/pgSQL Runtime - JSON value validation. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION json_model;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `jm_array_is_unique(val JSONB, path TEXT[], rep jm_report_entry[])` is an extension function and returns `BOOLEAN`.
- `jm_call(fun TEXT, val JSONB, path TEXT[], rep jm_report_entry[])` is an extension function and returns `BOOLEAN`.
- `jm_check_constraint(val JSONB, op TEXT, cst ANYELEMENT, path TEXT[], rep jm_report_entry[])` is an extension function and returns `BOOLEAN`.
- `jm_cmap_get(TEXT, JSONB)` is an extension function and returns `TEXT`.
- `jm_is_valid_date(val TEXT, path TEXT[], rep jm_report_entry[])` is an extension function and returns `BOOLEAN`.
- `jm_is_valid_datetime(val TEXT, path TEXT[], rep jm_report_entry[])` is an extension function and returns `BOOLEAN`.
- `jm_is_valid_email(val TEXT, path TEXT[], rep jm_report_entry[])` is an extension function and returns `BOOLEAN`.
- `jm_is_valid_extreg(val TEXT, path TEXT[], rep jm_report_entry[])` is an extension function and returns `BOOLEAN`.
- `jm_is_valid_regex(val TEXT, path TEXT[], rep jm_report_entry[])` is an extension function and returns `BOOLEAN`.
- `jm_is_valid_time(val TEXT, path TEXT[], rep jm_report_entry[])` is an extension function and returns `BOOLEAN`.
- `jm_is_valid_url(val TEXT, path TEXT[], rep jm_report_entry[])` is an extension function and returns `BOOLEAN`.
- `jm_is_valid_uuid(val TEXT, path TEXT[], rep jm_report_entry[])` is an extension function and returns `BOOLEAN`.
- `jm_object_size(val JSONB)` is an extension function and returns `INT`.
- `jm_report_add_entry` is an extension procedure.

### Requirements and Caveats

- The reviewed control file declares default version `2.0`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
