## 用法

来源：

- [官方扩展控制文件](https://github.com/postgrespro/memstat/blob/bf3b5a3ac1d8f3f25d4483859b16eab7818848f7/memstat.control)
- [官方上游文档](https://github.com/postgrespro/memstat/blob/bf3b5a3ac1d8f3f25d4483859b16eab7818848f7/README.md)

`memstat` — memstat：统计观测类 PostgreSQL 扩展；报告本地与实例后端内存上下文统计。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "memstat";
SELECT extversion
FROM pg_extension
WHERE extname = 'memstat';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
