## 用法

来源：

- [0.0.1 版本 README](https://github.com/hydradatabase/pg_quack/blob/5a49c79fd363bb27dcf6205fb3dd5d788758382b/README.md)
- [扩展 SQL](https://github.com/hydradatabase/pg_quack/blob/5a49c79fd363bb27dcf6205fb3dd5d788758382b/quack--0.0.1.sql)
- [扩展 control 文件](https://github.com/hydradatabase/pg_quack/blob/5a49c79fd363bb27dcf6205fb3dd5d788758382b/quack.control)
- [当前 main 分支说明](https://github.com/hydradatabase/pg_quack/blob/6213d66399c2b0fb8efd6640cfff2f8c6976a1e5/README.md)

`quack` 版本 `0.0.1` 是在 PostgreSQL 中嵌入 DuckDB 的表访问方法概念验证。先确保 `quack.data_dir` 选定的目录存在，并允许 PostgreSQL 操作系统账户写入，然后使用该访问方法创建和查询表：

```sql
CREATE EXTENSION quack;

CREATE TABLE quack_events (
  id bigint,
  payload text
) USING quack;

INSERT INTO quack_events VALUES (1, 'first event');
SELECT * FROM quack_events;
```

### 概念验证限制

已发布分支只支持 PostgreSQL 14 和 15、基础数据类型，以及 `COPY`、`INSERT` 和 `SELECT`。同一时间只能有一个连接对 quack 表执行 `INSERT` 或 `SELECT`，查询也不能把 quack 表与堆表或其他存储方法结合。未归档的 main 分支只宣布未来版本，并将用户引回这一概念验证。它适合受控评估，不应当作通用生产表引擎。
