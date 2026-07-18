## 用法

来源：

- [官方扩展控制文件](https://github.com/NguyenLe1605/pgmolji/blob/42b9486e29319026df7c957b484eaf47fceabecd/pgmolji.control)

`pgmolji` — pgmolji 是一个 pgrx 扩展，通过执行器钩子将查询结果中的文本与 varchar 值改写为随机 Emoji。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgmolji";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgmolji';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
