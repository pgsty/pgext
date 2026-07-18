## 用法

来源：

- [官方扩展控制文件](https://github.com/CrunchyData/pgseccomp/blob/3ed6d0780dc5d9a5204e57ce7628544a3610e218/pgseccomp.control)
- [官方上游文档](https://github.com/CrunchyData/pgseccomp/blob/3ed6d0780dc5d9a5204e57ce7628544a3610e218/README.md)

`pgseccomp` — 为 PostgreSQL 提供 seccomp 系统调用过滤，降低内核攻击面。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "pgseccomp";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgseccomp';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
