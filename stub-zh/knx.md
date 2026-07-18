## 用法

来源：

- [官方扩展控制文件](https://github.com/mike-kfed/pgknx/blob/c485a440f7347a0bd840018054466735d8f5f9ce/knx.control)
- [官方上游文档](https://github.com/mike-kfed/pgknx/blob/c485a440f7347a0bd840018054466735d8f5f9ce/README.md)

`knx` — knx：为 PostgreSQL 提供用于 KNX 地址的数据类型。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "knx";
SELECT extversion
FROM pg_extension
WHERE extname = 'knx';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
