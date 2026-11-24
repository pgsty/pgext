## 用法

来源：

- [Release 1.2.2](https://github.com/duckdb/pg_duckdb/releases/tag/v1.2.2)
- [Control 1.2.2](https://github.com/duckdb/pg_duckdb/blob/v1.2.2/pg_duckdb.control)

- [README.md](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/README.md)
- [pg_duckdb.control](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/pg_duckdb.control)
- [settings.md](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/docs/settings.md)
- [functions.md](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/docs/functions.md)
- [transactions.md](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/docs/transactions.md)
- [gotchas_and_syntax.md](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/docs/gotchas_and_syntax.md)
- [v1.2.1](https://github.com/duckdb/pg_duckdb/releases/tag/v1.2.1)

`pg_duckdb` 将 DuckDB 分析引擎嵌入 PostgreSQL，可查询普通 PostgreSQL 表和外部文件。1.2.1 内嵌 DuckDB 2.0.1，上游列出支持 PostgreSQL 14–18。扩展不可重定位，需要实例预加载。

### 启用并比较查询

将库加入现有设置，然后重启 PostgreSQL：

```ini
shared_preload_libraries = 'pg_duckdb'
```

以超级用户安装，检查分析查询：

```sql
CREATE EXTENSION pg_duckdb;
CREATE TABLE orders (category text, amount numeric);
INSERT INTO orders VALUES ('a', 10), ('a', 20), ('b', 5);
SET duckdb.force_execution = true;
EXPLAIN SELECT category, sum(amount) FROM orders GROUP BY category;
SELECT category, sum(amount) FROM orders GROUP BY category;
SET duckdb.force_execution = false;
```

普通查询默认使用 PostgreSQL，强制执行后才切换；DuckDB 专用函数或远程存储操作会自动选择 DuckDB。应检查计划并以代表性数据测量。`read_parquet`、`read_csv` 和 `read_json` 读取外部数据，`duckdb.create_simple_secret` 配置存储凭据。避免将密钥写入共享 SQL 历史，授权执行前先核验访问边界。

### 权限与配置

`duckdb.postgres_role` 指定可执行 DuckDB、管理密钥并使用 MotherDuck 数据的角色；默认仅超级用户可用。安装时可能创建所配置的角色。成员资格会开放共享的 DuckDB 密钥与数据，应有意授权。本地文件系统访问另外受 PostgreSQL 服务器文件角色与 DuckDB 文件系统设置约束。社区扩展加载由管理员显式选择。

### 事务与兼容性限制

同一事务不能同时写 PostgreSQL 和 DuckDB 表，也不能混用两种引擎的 DDL。应保持 `duckdb.unsafe_allow_mixed_transactions` 关闭，绕过限制可能只提交一侧，破坏一致性。默认也禁止在 PostgreSQL 函数内执行 DuckDB；相应的不安全设置存在已知崩溃风险。

1.2.1 支持嵌套定长数组，改善 numeric 与浮点转换，并修复并行工作进程的内存读取、取消清理和日志/统计中含密钥内容的处理。这些修复不能替代凭据权限控制及类型、查询兼容性检查。安装匹配动态库，遵循预加载重启要求，并在各数据库更新 SQL 扩展后再依赖新行为。

### 1.2.2 发行版

1.2.2 内嵌 DuckDB 2.0.1，增加 PostgreSQL 19 发布制品，并调整 CA 证书包加载。安装匹配服务器主版本的文件后，遵循扩展既有的预加载与 SQL 升级流程；远程对象凭据和 TLS 配置仍需单独管理。
