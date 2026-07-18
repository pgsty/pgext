## 用法

来源：

- [官方扩展控制文件](https://github.com/gleu/autovacuum_informations/blob/0dde21773a3929e0b0bacca018607d911fa71f00/autovacuum_informations.control)
- [官方上游文档](https://github.com/gleu/autovacuum_informations/blob/0dde21773a3929e0b0bacca018607d911fa71f00/README.md)

`autovacuum_informations` — 提供 autovacuum launcher PID 与 autovacuum worker 信息查询函数的 PostgreSQL 扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "autovacuum_informations";
SELECT extversion
FROM pg_extension
WHERE extname = 'autovacuum_informations';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
