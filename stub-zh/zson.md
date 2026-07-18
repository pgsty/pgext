## 用法

来源：

- [官方扩展控制文件](https://github.com/postgrespro/zson/blob/e214c79711b9eb92786280756a549b9e4cb2a215/zson.control)

`zson` — 基于训练字典压缩、支持透明类型转换与 JSONB 操作符的 JSONB 兼容类型。

已复核目录快照记录的版本为 `1.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "zson";
SELECT extversion
FROM pg_extension
WHERE extname = 'zson';
```

该上游项目与 `Postgres Professional` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
