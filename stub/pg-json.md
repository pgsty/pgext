## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg-json/pg-json-0.0.1/README)
- [Official extension control file (pg-json.control)](https://api.pgxn.org/src/pg-json/pg-json-0.0.1/pg-json.control)
- [Official extension SQL (pg-json--0.0.1.sql)](https://api.pgxn.org/src/pg-json/pg-json-0.0.1/pg-json--0.0.1.sql)

`pg-json` — pg-json - JSON support for PostgreSQL =====================================. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "pg-json";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `json_equals(this json, that json)` is an extension function and returns `boolean`.
- `json_get_value(data json, path text)` is an extension function and returns `text`.
- `json_in(cstring)` is an extension function and returns `json`.
- `json_not_equals(this json, that json)` is an extension function and returns `boolean`.
- `json_out(json)` is an extension function and returns `cstring`.
- `json` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
