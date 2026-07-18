## 用法

来源：

- [官方扩展控制文件](https://github.com/darold/pgtt-rsl/blob/f48f1a90bc09a80cb6f9d596535277960f658be9/pgtt_rsl.control)
- [官方上游文档](https://github.com/darold/pgtt-rsl/blob/f48f1a90bc09a80cb6f9d596535277960f658be9/README.md)

`pgtt_rsl` — 使用非日志表、行级安全、视图与清理后台进程提供 Oracle/DB2 风格的全局临时表

已复核目录快照记录的版本为 `2.0.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgtt_rsl";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgtt_rsl';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
