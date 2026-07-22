## Usage

Sources:

- [Pinned official README](https://github.com/asotolongo/db_info/blob/cfd4110846932f74e403c987449f70011ccf61d4/README.md)
- [Pinned extension SQL](https://github.com/asotolongo/db_info/blob/cfd4110846932f74e403c987449f70011ccf61d4/db_info--0.1.0.sql)

`db_info` provides SQL-only inventory helpers for summarizing the current database: object counts, database size and ownership, roles, installed extensions and languages, used data types, and tablespaces. It is intended for quick inspection rather than as a stable monitoring API.

### Core Workflow

The extension creates a fixed `db_info` schema. Query its views for a compact overview and call its functions for text reports:

```sql
CREATE EXTENSION db_info;

SELECT * FROM db_info.db_details;
SELECT * FROM db_info.db_obj_count;
SELECT * FROM db_info.get_extension_installed();
SELECT * FROM db_info.get_datatype_used();
```

### Important Objects

- `db_info.db_details`: current database name, owner, encoding, collation/ctype, tablespace, and size.
- `db_info.db_obj_count`: counts of schemas, tables, partitions, sequences, functions/procedures, triggers, rules, constraints, indexes, and views.
- `db_info.get_db_roles()`: a textual role and object-ownership report.
- `db_info.get_extension_installed()`: installed extension names and versions as text.
- `db_info.get_language_installed()`: installed procedural languages.
- `db_info.get_datatype_used()`: data types referenced by user columns.
- `db_info.get_tb_names()`: tablespaces used by tables, indexes, and materialized views.

### Caveats

The functions return formatted `text` or `SETOF text` rather than normalized rows, so their output is primarily human-readable. The views rely directly on PostgreSQL catalogs and `information_schema`; object-kind semantics and catalog columns can change between server versions.

The upstream README says it was tested on PostgreSQL 11 and later and explicitly warns that bugs may remain. Version 0.2.0 only revises the `db_obj_count` view; the base object definitions remain from 0.1.0. Review results before using them for auditing, capacity decisions, or automated alerts. No preload or restart is required.
