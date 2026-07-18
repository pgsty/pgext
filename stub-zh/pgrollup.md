## 用法

来源：

- [官方扩展控制文件](https://github.com/mikeizbicki/pgrollup/blob/904f19f5d02fb0901fdb6a7297eceefab549d7ae/pgrollup.control)
- [官方上游文档](https://github.com/mikeizbicki/pgrollup/blob/904f19f5d02fb0901fdb6a7297eceefab549d7ae/README.md)

`pgrollup` — 通过增量维护汇总表加速 PostgreSQL 聚合查询。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`, `plpython3u`。
整理后的兼容版本集合为 `12,13`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgrollup";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgrollup';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
