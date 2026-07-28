## 用法

来源：

- [官方上游 README](https://github.com/misachi/pg_wal_ext/blob/b8160bbd5956a7346bd2b61d4b426395881a1798/README.md)
- [官方扩展控制文件 (pg_wal_ext.control)](https://github.com/misachi/pg_wal_ext/blob/b8160bbd5956a7346bd2b61d4b426395881a1798/pg_wal_ext.control)
- [官方扩展 SQL (pg_wal_ext--1.0.sql)](https://github.com/misachi/pg_wal_ext/blob/b8160bbd5956a7346bd2b61d4b426395881a1798/pg_wal_ext--1.0.sql)

`pg_wal_ext` — 读取 WAL 文件并生成 SQL。在移动、转换或集成相应的数据时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_wal_ext;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pg_xlog_records(IN xlog_file_path text, OUT page_num int4, OUT txn_id xid, OUT xlog_type text, OUT commit_ts timestamptz, OUT generated_sql text)` 是一个扩展函数，返回 `SETOF`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
