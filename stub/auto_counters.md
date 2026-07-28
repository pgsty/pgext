## Usage

Sources:

- [Official upstream README](https://github.com/mbcheikh/auto_counters/blob/918ef7b428697d563572b51e271a8593f74847f6/README.md)
- [Official extension control file (auto_counters.control)](https://github.com/mbcheikh/auto_counters/blob/918ef7b428697d563572b51e271a8593f74847f6/auto_counters.control)
- [Official extension SQL (auto_counters--1.0.sql)](https://github.com/mbcheikh/auto_counters/blob/918ef7b428697d563572b51e271a8593f74847f6/auto_counters--1.0.sql)

`auto_counters` — A powerful and flexible PostgreSQL extension for **automatic contextual numbering** based on multiple field combinations. Perfect for generating document numbers, invoice numbers, or any numbering system that depends on contextual information like year, department, or category. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION auto_counters;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `auto_counters_install()` is an extension function and returns `void`.
- `create_counter_def(p_counter_id VARCHAR, p_table_name VARCHAR, p_fields TEXT [], p_description TEXT DEFAULT NULL, p_is_active BOOLEAN DEFAULT TRUE)` is an extension function and returns `VOID`.
- `create_counter_trigger_on_def_insert()` is an extension function and returns `TRIGGER`.
- `delete_counter_def(p_counter_id VARCHAR, p_cascade BOOLEAN DEFAULT FALSE)` is an extension function and returns `VOID`.
- `generic_counter_trigger()` is an extension function and returns `TRIGGER`.
- `get_counter_def(p_counter_id VARCHAR DEFAULT NULL)` is an extension function and returns `TABLE`.
- `get_counter_status(p_counter_id VARCHAR DEFAULT NULL)` is an extension function and returns `TABLE`.
- `get_next_counter_value(p_counter_id VARCHAR, p_key_values TEXT [])` is an extension function and returns `INTEGER`.
- `set_field(record_data ANYELEMENT, field_name TEXT, field_value ANYELEMENT)` is an extension function and returns `ANYELEMENT`.
- `sync_all_counter_triggers()` is an extension function and returns `void`.
- `toggle_counter_def(p_counter_id VARCHAR, p_is_active BOOLEAN)` is an extension function and returns `VOID`.
- `update_counter_def(p_counter_id VARCHAR, p_table_name VARCHAR DEFAULT NULL, p_fields TEXT [] DEFAULT NULL, p_description TEXT DEFAULT NULL, p_is_active BOOLEAN DEFAULT NULL)` is an extension function and returns `VOID`.
- `update_counter_trigger_on_def_change()` is an extension function and returns `TRIGGER`.
- `vw_counter_status` is an extension-defined view.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
