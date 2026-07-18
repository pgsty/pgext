## 用法

来源：

- [官方扩展控制文件](https://github.com/preedep/kham/blob/ec2394cc70ab9687b6e0bcd7c02c2f75a4cf9fbd/kham-pg/kham_pg.control)
- [官方上游文档](https://pgxn.org/dist/kham_pg/0.8.2/README.html)
- [官方 Rust 包清单](https://github.com/preedep/kham/blob/ec2394cc70ab9687b6e0bcd7c02c2f75a4cf9fbd/kham-pg/Cargo.toml)

`kham_pg` — 泰语全文搜索解析器，支持分词、Soundex、罗马化与命名实体标注

已复核目录快照记录的版本为 `0.8.2`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "kham_pg";
SELECT extversion
FROM pg_extension
WHERE extname = 'kham_pg';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
