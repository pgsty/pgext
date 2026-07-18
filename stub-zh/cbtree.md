## 用法

来源：

- [官方扩展控制文件](https://github.com/yanliu8/cbtree/blob/871f57358cb297f22be6f419a8e52dc86567ab7c/cbtree.control)

`cbtree` — cbtree：通用类 PostgreSQL 扩展；上游说明为“带计数功能的 B-tree 访问方法”。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "cbtree";
SELECT extversion
FROM pg_extension
WHERE extname = 'cbtree';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
