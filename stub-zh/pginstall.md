## 用法

来源：

- [官方扩展控制文件](https://github.com/dimitri/pginstall/blob/66e4274fc544c48be03c719f8d3f03955649a994/src/client/pginstall.control)

`pginstall` — PostgreSQL 扩展安装器，通过 local_preload_libraries 钩住 CREATE EXTENSION，从仓库下载并安装扩展归档。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "pginstall";
SELECT extversion
FROM pg_extension
WHERE extname = 'pginstall';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
