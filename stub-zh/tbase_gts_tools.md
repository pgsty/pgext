## 用法

来源：

- [官方上游 README](https://github.com/tencent/tbase/blob/128d47502476f84392a4ac54603efe007d063285/contrib/README)
- [官方扩展控制文件 (tbase_gts_tools.control)](https://github.com/tencent/tbase/blob/128d47502476f84392a4ac54603efe007d063285/contrib/tbase_gts_tools/tbase_gts_tools.control)
- [官方扩展 SQL (tbase_gts_tools--1.0.sql)](https://github.com/tencent/tbase/blob/128d47502476f84392a4ac54603efe007d063285/contrib/tbase_gts_tools/tbase_gts_tools--1.0.sql)

`tbase_gts_tools` — TBase 帮助程序用于检查事务全局时间戳和堆页元组元数据。在管理或自动化上述数据库行为时使用它。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION tbase_gts_tools;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `heap_page_ids(IN page bytea, OUT lp smallint, OUT lp_off smallint, OUT lp_flags smallint, OUT lp_len smallint, OUT t_xmin xid, OUT t_xmax xid, OUT t_xmin_gts bigint, OUT t_xmax_gts bigint, OUT t_field3 int4, OUT t_ctid tid, OUT t_infomask2 integer, OUT t_infomask integer, O…)` 是一个扩展函数，返回 `SETOF`。
- `heap_page_items_with_gts(IN page bytea, OUT lp smallint, OUT lp_off smallint, OUT lp_flags smallint, OUT lp_len smallint, OUT t_xmin xid, OUT t_xmax xid, OUT t_xmin_gts bigint, OUT t_xmax_gts bigint, OUT t_field3 int4, OUT t_ctid tid, OUT t_infomask2 integer, OUT t_infomask integer, O…)` 是一个扩展函数，返回 `SETOF`。
- `heap_page_items_with_gts_log(IN page bytea, OUT lp smallint, OUT lp_off smallint, OUT lp_flags smallint, OUT lp_len smallint, OUT t_xmin xid, OUT t_xmax xid, OUT t_xmin_gts bigint, OUT t_xmax_gts bigint, OUT t_field3 int4, OUT t_ctid tid, OUT t_infomask2 integer, OUT t_infomask integer, O…)` 是一个扩展函数，返回 `SETOF`。
- `heap_page_items_without_data(IN page bytea, OUT lp smallint, OUT lp_off smallint, OUT lp_flags smallint, OUT lp_len smallint, OUT t_xmin xid, OUT t_xmax xid, OUT t_xmin_gts bigint, OUT t_xmax_gts bigint, OUT t_field3 int4, OUT t_ctid tid, OUT t_infomask2 integer, OUT t_infomask integer, O…)` 是一个扩展函数，返回 `SETOF`。
- `txid_gts(int)` 是一个扩展函数，返回 `bigint`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
