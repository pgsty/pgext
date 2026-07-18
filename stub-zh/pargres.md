## 用法

来源：

- [官方扩展控制文件](https://github.com/danolivo/pargres/blob/d62ebdff2f9758f1b444639dc5cf65fd1edc7e4f/pargres.control)
- [官方上游文档](https://github.com/danolivo/pargres/blob/d62ebdff2f9758f1b444639dc5cf65fd1edc7e4f/README.md)

`pargres` — pargres：用于共享无集群并行查询执行的 PostgreSQL 原型扩展，需要 shared_preload_libraries 预加载。

已复核目录快照记录的版本为 `0.1`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "pargres";
SELECT extversion
FROM pg_extension
WHERE extname = 'pargres';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
