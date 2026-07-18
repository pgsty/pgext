## 用法

来源：

- [官方扩展控制文件](https://github.com/sangli00/hash_fdw/blob/9c685317f6a08381b426706c235ac515505cac51/hash_fdw.control)
- [官方上游文档](https://github.com/sangli00/hash_fdw/blob/9c685317f6a08381b426706c235ac515505cac51/README)

`hash_fdw` — 基于共享内存哈希表的外部数据包装器，可按配置的键存储和查询行。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "hash_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'hash_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
