## Usage

Sources:

- [Official upstream README](https://github.com/tencent/tbase/blob/128d47502476f84392a4ac54603efe007d063285/contrib/README)
- [Official extension control file (tbase_gts_tools.control)](https://github.com/tencent/tbase/blob/128d47502476f84392a4ac54603efe007d063285/contrib/tbase_gts_tools/tbase_gts_tools.control)
- [Official extension SQL (tbase_gts_tools--1.0.sql)](https://github.com/tencent/tbase/blob/128d47502476f84392a4ac54603efe007d063285/contrib/tbase_gts_tools/tbase_gts_tools--1.0.sql)

`tbase_gts_tools` — TBase helpers for inspecting transaction global timestamps and heap-page tuple metadata. Use it when administering or automating the database behavior described above. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION tbase_gts_tools;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `heap_page_ids(IN page bytea, OUT lp smallint, OUT lp_off smallint, OUT lp_flags smallint, OUT lp_len smallint, OUT t_xmin xid, OUT t_xmax xid, OUT t_xmin_gts bigint, OUT t_xmax_gts bigint, OUT t_field3 int4, OUT t_ctid tid, OUT t_infomask2 integer, OUT t_infomask integer, O…)` is an extension function and returns `SETOF`.
- `heap_page_items_with_gts(IN page bytea, OUT lp smallint, OUT lp_off smallint, OUT lp_flags smallint, OUT lp_len smallint, OUT t_xmin xid, OUT t_xmax xid, OUT t_xmin_gts bigint, OUT t_xmax_gts bigint, OUT t_field3 int4, OUT t_ctid tid, OUT t_infomask2 integer, OUT t_infomask integer, O…)` is an extension function and returns `SETOF`.
- `heap_page_items_with_gts_log(IN page bytea, OUT lp smallint, OUT lp_off smallint, OUT lp_flags smallint, OUT lp_len smallint, OUT t_xmin xid, OUT t_xmax xid, OUT t_xmin_gts bigint, OUT t_xmax_gts bigint, OUT t_field3 int4, OUT t_ctid tid, OUT t_infomask2 integer, OUT t_infomask integer, O…)` is an extension function and returns `SETOF`.
- `heap_page_items_without_data(IN page bytea, OUT lp smallint, OUT lp_off smallint, OUT lp_flags smallint, OUT lp_len smallint, OUT t_xmin xid, OUT t_xmax xid, OUT t_xmin_gts bigint, OUT t_xmax_gts bigint, OUT t_field3 int4, OUT t_ctid tid, OUT t_infomask2 integer, OUT t_infomask integer, O…)` is an extension function and returns `SETOF`.
- `txid_gts(int)` is an extension function and returns `bigint`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
