## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/json_accessors/json_accessors-1.3.6/README.md)
- [Official extension control file (json_accessors.control)](https://api.pgxn.org/src/json_accessors/json_accessors-1.3.6/json_accessors.control)
- [Official extension SQL (json_accessors.sql)](https://api.pgxn.org/src/json_accessors/json_accessors-1.3.6/sql/json_accessors.sql)

`json_accessors` — JSON accessor functions for PostgreSQL ======================================. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION json_accessors;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `json_array_to_bigint_array(text)` is an extension function and returns `bigint[]`.
- `json_array_to_boolean_array(text)` is an extension function and returns `boolean[]`.
- `json_array_to_int_array(text)` is an extension function and returns `int[]`.
- `json_array_to_numeric_array(text)` is an extension function and returns `numeric[]`.
- `json_array_to_object_array(text)` is an extension function and returns `text[]`.
- `json_array_to_text_array(text)` is an extension function and returns `text[]`.
- `json_array_to_timestamp_array(text)` is an extension function and returns `timestamp`.
- `json_get_bigint(text, text)` is an extension function and returns `bigint`.
- `json_get_bigint_array(text, text)` is an extension function and returns `bigint[]`.
- `json_get_boolean(text, text)` is an extension function and returns `boolean`.
- `json_get_boolean_array(text, text)` is an extension function and returns `boolean[]`.
- `json_get_int(text, text)` is an extension function and returns `int`.
- `json_get_int_array(text, text)` is an extension function and returns `int[]`.
- `json_get_numeric(text, text)` is an extension function and returns `numeric`.

### Requirements and Caveats

- The reviewed control file declares default version `1.3.6`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
