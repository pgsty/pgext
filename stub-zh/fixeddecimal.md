## 用法

来源：

- [官方扩展控制文件](https://github.com/2ndQuadrant/fixeddecimal/blob/970f56b53a2eef74130527778190597dbd9a6a30/fixeddecimal.control)
- [官方上游文档](https://github.com/2ndQuadrant/fixeddecimal/blob/970f56b53a2eef74130527778190597dbd9a6a30/README.md)

`fixeddecimal` — fixeddecimal：提供高性能固定精度 decimal 类型的 PostgreSQL 扩展。

已复核目录快照记录的版本为 `1.1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "fixeddecimal";
SELECT extversion
FROM pg_extension
WHERE extname = 'fixeddecimal';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
