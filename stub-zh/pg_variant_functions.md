## 用法

来源：

- [官方扩展控制文件](https://github.com/rlichtenwalter/pg_variant_functions/blob/8ad88883b840ee334aeb4d549aacfa102dceae11/pg_variant_functions.control)
- [官方上游文档](https://github.com/rlichtenwalter/pg_variant_functions/blob/8ad88883b840ee334aeb4d549aacfa102dceae11/README.md)

`pg_variant_functions` — pg_variant_functions：函数工具类 PostgreSQL 扩展；上游说明为“变异分析函数”。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`tinyint`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_variant_functions";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_variant_functions';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
