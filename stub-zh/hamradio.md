## 用法

来源：

- [官方扩展控制文件](https://github.com/df7cb/postgresql-hamradio/blob/4e2181d68b62d224b120ce1bb38d2c23ee63e937/hamradio.control)

`hamradio` — hamradio：为 PostgreSQL 提供 ham radio 扩展。

已复核目录快照记录的版本为 `1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`postgis`, `uctext`。

```sql
CREATE EXTENSION "hamradio";
SELECT extversion
FROM pg_extension
WHERE extname = 'hamradio';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
