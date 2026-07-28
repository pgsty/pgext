## Usage

Sources:

- [Official upstream README](https://github.com/fljdin/mysql_migrator/blob/3c15bd7d7b9a2183d3c9753b3ec8fd238b5ca2cc/README.md)
- [Official extension control file (mysql_migrator.control)](https://github.com/fljdin/mysql_migrator/blob/3c15bd7d7b9a2183d3c9753b3ec8fd238b5ca2cc/mysql_migrator.control)
- [Official extension SQL (mysql_migrator--0.3.0.sql)](https://github.com/fljdin/mysql_migrator/blob/3c15bd7d7b9a2183d3c9753b3ec8fd238b5ca2cc/mysql_migrator--0.3.0.sql)

`mysql_migrator` — MySQL/MariaDB to PostgreSQL migration tools ===========================================. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION mysql_migrator;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `db_migrator_callback(OUT create_metadata_views_fun regprocedure, OUT translate_datatype_fun regprocedure, OUT translate_identifier_fun regprocedure, OUT translate_expression_fun regprocedure, OUT create_foreign_table_fun regprocedure)` is an extension function and returns `record`.
- `mysql_create_catalog(server name, schema name DEFAULT NAME 'public', options jsonb DEFAULT NULL)` is an extension function and returns `void`.
- `mysql_migrate_identity(pgstage_schema name DEFAULT NAME 'pgsql_stage')` is an extension function and returns `integer`.
- `mysql_mkforeign(server name, schema name, table_name name, orig_schema text, orig_table text, column_names name[], column_options jsonb[], orig_columns text[], data_types text[], nullable boolean[], options jsonb)` is an extension function and returns `text`.
- `mysql_translate_datatype(v_type text, v_length integer, v_precision integer, v_scale integer)` is an extension function and returns `text`.
- `mysql_translate_expression(s text)` is an extension function and returns `text`.
- `mysql_translate_identifier_noop(text)` is an extension function and returns `name`.

### Requirements and Caveats

- The reviewed control file declares default version `0.3.0`.
- Install the confirmed extension dependencies first: `mysql_fdw`, `db_migrator`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
