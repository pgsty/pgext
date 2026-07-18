## 用法

来源：

- [官方扩展控制文件](https://github.com/rogerioacp/oblivpg_ftw/blob/5a26f0ffa02242aa187d22bde8ca8dc3bc05dde0/oblivpg_fdw.control)
- [官方上游文档](https://github.com/rogerioacp/oblivpg_ftw/blob/5a26f0ffa02242aa187d22bde8ca8dc3bc05dde0/README.md)

`oblivpg_fdw` — oblivpg_fdw：基于可信硬件与 ORAM 的隐私保护外部数据包装器。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "oblivpg_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'oblivpg_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
