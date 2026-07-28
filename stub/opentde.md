## Usage

Sources:

- [Official extension control file (opentde.control)](https://github.com/ddlifter/cuddly-meme/blob/4740cdde27bf12b026279860af2559db780e1972/opentde.control)
- [Official extension SQL (opentde--1.0.sql)](https://github.com/ddlifter/cuddly-meme/blob/4740cdde27bf12b026279860af2559db780e1972/opentde--1.0.sql)

`opentde` — Open Transparent Data Encryption for PostgreSQL. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION opentde;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `opentde_blind_index(text)` is an extension function and returns `bytea`.
- `opentde_debug_keys()` is an extension function and returns `text`.
- `opentde_get_dek_hex(oid)` is an extension function and returns `text`.
- `opentde_rotate_master_key(bytea)` is an extension function and returns `int4`.
- `opentde_rotate_table_dek(oid)` is an extension function and returns `int4`.
- `opentde_set_master_key(bytea)` is an extension function and returns `void`.
- `opentde_tableam_handler(internal)` is an extension function and returns `table_am_handler`.
- `set_column_level_encryption(enable boolean)` is an extension function and returns `void`.
- `opentde` is an extension-defined access method.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
