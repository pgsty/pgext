## 用法

来源：

- [官方扩展控制文件](https://github.com/sroeschus/loginfo/blob/e78fb0312856ff0c4fc4000fa49f75be24ad3d67/loginfo.control)
- [官方上游文档](https://github.com/sroeschus/loginfo/blob/e78fb0312856ff0c4fc4000fa49f75be24ad3d67/README.md)

`loginfo` — 查询 PostgreSQL 日志配置、日志文件列表及普通或 CSV 日志内容。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "loginfo";
SELECT extversion
FROM pg_extension
WHERE extname = 'loginfo';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
