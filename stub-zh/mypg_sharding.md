## 用法

来源：

- [官方扩展控制文件](https://github.com/tongxin/mypg_sharding/blob/567526b4292bc4b613b7d7c18ad0b423da1985d6/mypg_sharding.control)
- [官方上游文档](https://github.com/tongxin/mypg_sharding/blob/567526b4292bc4b613b7d7c18ad0b423da1985d6/README.md)

`mypg_sharding` — mypg_sharding：基于 postgres_fdw 的 PostgreSQL 分片工具扩展。

已复核目录快照记录的版本为 `0.0.1`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`, `postgres_fdw`。

```sql
CREATE EXTENSION "mypg_sharding";
SELECT extversion
FROM pg_extension
WHERE extname = 'mypg_sharding';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
