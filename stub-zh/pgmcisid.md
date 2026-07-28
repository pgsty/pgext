## 用法

来源：

- [官方扩展控制文件 (pgmcisid.control)](https://github.com/morkato/mcisid/blob/5b67978b152806ffc6943b7d351f4f36a1a2d3ba/pgmcisid/pgmcisid.control)
- [官方扩展 SQL (pgmcisid--1.0.sql)](https://github.com/morkato/mcisid/blob/5b67978b152806ffc6943b7d351f4f36a1a2d3ba/pgmcisid/pgmcisid--1.0.sql)

`pgmcisid` — MCIS ID V1 生成器。当 SQL 需要这些特殊函数或聚合时使用它。请使用上方链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgmcisid;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `mcisidv1_created_at(mcisidv1)` 是一个扩展函数，返回 `TIMESTAMP`。
- `mcisidv1_gen` 是一个扩展函数。
- `mcisidv1_get_epoch()` 是一个扩展函数，返回 `TIMESTAMP`。
- `mcisidv1_get_instant(mcisidv1)` 是一个扩展函数，返回 `BIGINT`。
- `mcisidv1_instant_sequence(mcisidv1)` 是一个扩展函数，返回 `BIGINT`。
- `mcisidv1_origin_model(mcisidv1)` 是一个扩展函数，返回 `SMALLINT`。
- `mcisidv1_type_cmp(mcisidv1, mcisidv1)` 是一个扩展函数，返回 `INTEGER`。
- `mcisidv1_type_eq(mcisidv1, mcisidv1)` 是一个扩展函数，返回 `BOOLEAN`。
- `mcisidv1_type_gt(mcisidv1, mcisidv1)` 是一个扩展函数，返回 `BOOLEAN`。
- `mcisidv1_type_in(CSTRING)` 是一个扩展函数，返回 `mcisidv1`。
- `mcisidv1_type_lt(mcisidv1, mcisidv1)` 是一个扩展函数，返回 `BOOLEAN`。
- `mcisidv1_type_out(mcisidv1)` 是一个扩展函数，返回 `CSTRING`。
- `text_to_mcisidv1(TEXT)` 是一个扩展函数，返回 `mcisidv1`。
- `mcisidv1` 是一个扩展定义的类型。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
