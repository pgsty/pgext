## Usage

Sources:

- [Official upstream README](https://github.com/yugabyte/yugabyte-db/blob/0f345ba55edef263a2f67c44438dbdffcbb44754/src/postgres/third-party-extensions/mage/README.md)
- [Official extension control file (mage.control)](https://github.com/yugabyte/yugabyte-db/blob/0f345ba55edef263a2f67c44438dbdffcbb44754/src/postgres/third-party-extensions/mage/mage.control)
- [Official extension SQL (mage--1.5.0--1.6.0.sql)](https://github.com/yugabyte/yugabyte-db/blob/0f345ba55edef263a2f67c44438dbdffcbb44754/src/postgres/third-party-extensions/mage/mage--1.5.0--1.6.0.sql)

`mage` — Since AGE is based on the powerful PostgreSQL RDBMS, it is robust and fully featured. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION mage;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `mag_catalog.age_graph_stats(agtype)` is an extension function and returns `agtype`.
- `mag_catalog.age_is_valid_label_name(agtype)` is an extension function and returns `boolean`.
- `mag_catalog.age_tostring("any")` is an extension function and returns `agtype`.
- `mag_catalog.agtype_array_to_agtype(agtype[])` is an extension function and returns `agtype`.
- `mag_catalog.agtype_contained_by_top_level(agtype, agtype)` is an extension function and returns `boolean`.
- `mag_catalog.agtype_contains_top_level(agtype, agtype)` is an extension function and returns `boolean`.
- `mag_catalog.agtype_to_json(agtype)` is an extension function and returns `json`.
- `mag_catalog.create_elabel(graph_name cstring, label_name cstring)` is an extension function and returns `void`.
- `mag_catalog.create_vlabel(graph_name cstring, label_name cstring)` is an extension function and returns `void`.
- `mag_catalog.graph_exists(graph_name name)` is an extension function and returns `agtype`.
- `mag_catalog.load_edges_from_file(graph_name name, label_name name, file_path text, load_as_agtype bool default false)` is an extension function and returns `void`.
- `mag_catalog.load_labels_from_file(graph_name name, label_name name, file_path text, id_field_exists bool default true, load_as_agtype bool default false)` is an extension function and returns `void`.
- `mag_catalog.gin_agtype_ops` is an extension-defined operator class.

### Requirements and Caveats

- The reviewed control file declares default version `1.6.0`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
