## 用法

来源：

- [官方扩展控制文件](https://github.com/BlueTreble/pg_thud/blob/4d2aba37d20c942f9c05412a6f045f328106cfb3/pg_thud.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/pg_thud/)

`pg_thud` — 扩展 pgtap 的数据库单元测试框架，提供脚手架与测试数据辅助功能。

已复核目录快照记录的版本为 `0.1.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`pgtap`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_thud";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_thud';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
