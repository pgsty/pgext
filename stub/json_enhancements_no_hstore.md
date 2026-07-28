## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/json_enhancements/json_enhancements-1.0.2/README.md)
- [Official extension control file (json_enhancements_no_hstore.control)](https://api.pgxn.org/src/json_enhancements/json_enhancements-1.0.2/json_enhancements_no_hstore.control)
- [Official extension SQL (json_enhancements_no_hstore.sql)](https://api.pgxn.org/src/json_enhancements/json_enhancements-1.0.2/sql/json_enhancements_no_hstore.sql)

`json_enhancements_no_hstore` — Json Enhancements for PostgreSQL 9.2 ====================================. Use it when an application needs this specific database capability. The reviewed upstream project is archived or no longer maintained.

### Core Workflow

```sql
CREATE EXTENSION json_enhancements_no_hstore;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `json_agg_finalfn(internal)` is an extension function and returns `json`.
- `json_agg_transfn(internal, anyelement)` is an extension function and returns `internal`.
- `json_array_element(the_json json, element integer)` is an extension function and returns `json`.
- `json_array_element_text(the_json json, element integer)` is an extension function and returns `text`.
- `json_array_elements(the_json json)` is an extension function and returns `TABLE`.
- `json_array_length(the_json json)` is an extension function and returns `int`.
- `json_each(the_json json, key out text, out value json)` is an extension function and returns `SETOF record`.
- `json_each_text(the_json json, key out text, value out text)` is an extension function and returns `SETOF record`.
- `json_extract_path(the_json json, variadic path_elements text[])` is an extension function and returns `json`.
- `json_extract_path_op(the_json json, path_elements text[])` is an extension function and returns `json`.
- `json_extract_path_text(the_json json, variadic path_elements text[])` is an extension function and returns `text`.
- `json_extract_path_text_op(the_json json, path_elements text[])` is an extension function and returns `text`.
- `json_object_field(json, text)` is an extension function and returns `json`.
- `json_object_field_text(json, text)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
