## 用法

来源：

- [官方扩展控制文件](https://github.com/tatut/pgprolog/blob/e3629cc2f5a56da98c22939436d9d803e2f9a6d5/pgprolog.control)
- [官方上游文档](https://github.com/tatut/pgprolog/blob/e3629cc2f5a56da98c22939436d9d803e2f9a6d5/README.md)
- [官方 Rust 包清单](https://github.com/tatut/pgprolog/blob/e3629cc2f5a56da98c22939436d9d803e2f9a6d5/Cargo.toml)

`pgprolog` — pgprolog 是一个基于 pgrx、内嵌 Scryer Prolog 的 PostgreSQL Prolog 过程语言处理器。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgprolog";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgprolog';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
