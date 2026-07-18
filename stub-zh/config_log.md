## 用法

来源：

- [官方上游文档](https://github.com/ibarwick/config_log/blob/db66842eec9398701bd045fad7e6a09f6b77eeb5/README.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/config_log/)

`config_log` — config_log：统计观测类 PostgreSQL 扩展；上游说明为“使用自定义后台工作进程，将 postgresql.conf 的变更记录到表中”。

已复核目录快照记录的版本为 `0.1.7`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "config_log";
SELECT extversion
FROM pg_extension
WHERE extname = 'config_log';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
