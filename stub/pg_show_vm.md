## Usage

Sources:

- [Official upstream README](https://github.com/s-hironobu/pg_show_vm/blob/61b6600ae9af4c1527e1157f58642880da70b584/README.md)
- [Official extension control file (pg_show_vm.control)](https://github.com/s-hironobu/pg_show_vm/blob/61b6600ae9af4c1527e1157f58642880da70b584/pg_show_vm.control)
- [Official extension SQL (pg_show_vm--1.0.sql)](https://github.com/s-hironobu/pg_show_vm/blob/61b6600ae9af4c1527e1157f58642880da70b584/pg_show_vm--1.0.sql)

`pg_show_vm` — This extension supports PostgreSQL versions 16 and 17. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_show_vm;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_show_rel_vm(IN relname text, IN index bool, IN partition bool, OUT relid int, OUT relpages int, OUT all_visible int, OUT all_frozen int, OUT type int)` is an extension function and returns `SETOF`.
- `pg_show_vm(IN relid oid, OUT relid int, OUT relpages int, OUT all_visible int, OUT all_frozen int, OUT type int)` is an extension function and returns `SETOF`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
