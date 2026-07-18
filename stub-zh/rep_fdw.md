## 用法

来源：

- [官方项目或服务商页面](https://anthonysotolongo.wordpress.com/)

`rep_fdw` — 通过 postgres_fdw 将表变更复制到另一台 PostgreSQL 服务器

已复核目录快照记录的版本为 `0.1.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`, `postgres_fdw`。

```sql
CREATE EXTENSION "rep_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'rep_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
