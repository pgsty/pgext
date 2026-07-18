## 用法

来源：

- [官方扩展控制文件](https://github.com/troels/pghooks/blob/f82688238e25bb220478404eeeb4d6a825f014a1/pglogical_output_nzhooks.control)

`pglogical_output_nzhooks` — 仅保留硬编码表名变更的 pglogical 输出插件行过滤钩子示例

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`pglogical`。

```sql
CREATE EXTENSION "pglogical_output_nzhooks";
SELECT extversion
FROM pg_extension
WHERE extname = 'pglogical_output_nzhooks';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
