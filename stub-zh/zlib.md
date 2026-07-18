## 用法

来源：

- [官方扩展控制文件](https://github.com/overplumbum/postgresql-zlib/blob/74f0e5126266105fea05dcd248400e73bc7bacc2/zlib.control)

`zlib` — PostgreSQL zlib 解压函数。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "zlib";
SELECT extversion
FROM pg_extension
WHERE extname = 'zlib';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
