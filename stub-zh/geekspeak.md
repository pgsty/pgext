## 用法

来源：

- [官方扩展控制文件](https://github.com/ttfkam/pg_geekspeak/blob/8d176e395f7f8a98a3c4dd2d48a1ee1e359b4e47/geekspeak.control)
- [官方上游文档](https://github.com/ttfkam/pg_geekspeak/blob/8d176e395f7f8a98a3c4dd2d48a1ee1e359b4e47/README.md)

`geekspeak` — GeekSpeak 播客网站的数据库模式扩展。

已复核目录快照记录的版本为 `1.0.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`btree_gist`, `isn`, `multicorn`, `pgcrypto`, `plpgsql`。
整理后的兼容版本集合为 `10`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "geekspeak";
SELECT extversion
FROM pg_extension
WHERE extname = 'geekspeak';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
