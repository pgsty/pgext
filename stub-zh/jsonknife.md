## 用法

来源：

- [官方扩展控制文件](https://github.com/niquola/jsonknife/blob/ce6e65ad39a0344feab49754636daf8bd06c5c92/jsonknife.control)
- [官方上游文档](https://github.com/niquola/jsonknife/blob/ce6e65ad39a0344feab49754636daf8bd06c5c92/README.md)

`jsonknife` — jsonknife：函数工具类 PostgreSQL 扩展；上游说明为“实用的 JSONB 检查和转换函数”。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `11`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "jsonknife";
SELECT extversion
FROM pg_extension
WHERE extname = 'jsonknife';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
