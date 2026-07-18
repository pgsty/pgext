## 用法

来源：

- [官方上游文档](https://github.com/veridit/sql_saga/blob/9e39fcc7fc862fcb793f59d996eb01268cf9e822/README.md)

`sql_saga` — 用于有效时间表、时间外键、保留历史的视图及批量时间合并操作的 C 与 PL/pgSQL 扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`btree_gist`, `plpgsql`。
整理后的兼容版本集合为 `18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "sql_saga";
SELECT extversion
FROM pg_extension
WHERE extname = 'sql_saga';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
