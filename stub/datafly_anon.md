## Usage

Sources:

- [Official extension control file (datafly_anon.control)](https://api.pgxn.org/src/datafly_anon/datafly_anon-1.0.30/datafly_anon.control)
- [Official extension SQL (datafly_anon--1.25.sql)](https://api.pgxn.org/src/datafly_anon/datafly_anon-1.0.30/sql/datafly_anon--1.25.sql)

`datafly_anon` — Datafly anonymizer. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION datafly_anon;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `add_level_generalization(sch_name varchar, attribute_name varchar,tbl_name varchar,generalization_rule varchar, new_level integer,function varchar, target_sch_name varchar, target_view varchar, re_init_anon bool)` is an extension function and returns `void`.
- `check_if_generalization_rule_exists(attribute_name varchar, schema_name varchar, table_name varchar, rule varchar, target_schema_name varchar, target_table_name varchar)` is an extension function and returns `boolean`.
- `check_if_level_exists(attribute_name varchar, schema_name varchar,table_name varchar, new_level integer, target_schema_name varchar, target_table_name varchar)` is an extension function and returns `boolean`.
- `configure_plugin(json_config json)` is an extension function and returns `text`.
- `does_column_exist_in_table(attribute_name varchar, sch_name varchar,tbl_name varchar)` is an extension function and returns `boolean`.
- `does_table_exist(schema_name varchar, tbl_name varchar)` is an extension function and returns `boolean`.
- `generalize(attribute_name varchar, target_sch_name varchar, target_view varchar, sch_name varchar, tbl_name varchar)` is an extension function and returns `void`.
- `generalize_daterange(val DATE, step TEXT)` is an extension function and returns `DATERANGE`.
- `generalize_numrange(val NUMERIC, step VARCHAR)` is an extension function and returns `NUMRANGE`.
- `generate_init_view(sch_name varchar, tbl_name varchar, target_sch_name varchar, target_view varchar, test_mode bool, is_triggered bool)` is an extension function and returns `varchar`.
- `generate_triggers(k_param integer,schema_name varchar, table_name varchar, target_schema_name varchar, target_table_name varchar)` is an extension function and returns `void`.
- `init_datafly(k integer, sch_name varchar, tbl_name varchar, target_sch_name varchar, target_view varchar, test_mode bool, is_triggered bool default false)` is an extension function and returns `text`.
- `init_datafly_tg()` is an extension function and returns `TRIGGER`.
- `remove_level_generalization(attribute_name varchar, tbl_name varchar, sch_name varchar, target_sch_name varchar, target_view varchar, generalization_lvl integer,re_init_anon bool)` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `1.25`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
