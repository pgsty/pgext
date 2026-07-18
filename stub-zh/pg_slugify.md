## 用法

来源：

- [官方扩展控制文件](https://github.com/pindlebot/pg_slugify/blob/fab77da2df8306eac3f8010ff2c7cad8c3dfb75e/pg_slugify.control)

`pg_slugify` — 使用 unaccent 将文本规范化为适合 URL 的 slug

已复核目录快照记录的版本为 `0.0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`unaccent`。

```sql
CREATE EXTENSION "pg_slugify";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_slugify';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
