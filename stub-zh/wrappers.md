


## 用法

来源：[official README](https://github.com/supabase/wrappers/blob/v0.6.1/README.md)、[official docs](https://fdw.dev/)、[v0.6.1 release](https://github.com/supabase/wrappers/releases/tag/v0.6.1)

`wrappers` 既是一个用 Rust 编写 PostgreSQL 外部数据包装器的框架，也是 Supabase 维护的一组 FDW 打包集合。单个扩展会安装多种包装器实现，然后每个外部服务器再选择自己需要的具体包装器类型。

```sql
CREATE EXTENSION wrappers;
```

### 典型流程

先为某个包装器创建服务器，再通过外部表暴露远端数据：

```sql
CREATE SERVER stripe_server
  FOREIGN DATA WRAPPER stripe_wrapper
  OPTIONS (
    api_key_id 'stripe_api_key',
    api_url 'https://api.stripe.com/v1/'
  );

CREATE FOREIGN TABLE stripe_customers (
  id text,
  email text,
  name text,
  description text,
  created timestamp,
  attrs jsonb
)
  SERVER stripe_server
  OPTIONS (
    object 'customers',
    rowid_column 'id'
  );
```

### 覆盖范围

上游提供了用于 BigQuery、ClickHouse、DuckDB、DynamoDB、MySQL/Doris、Redis、S3、S3 Vectors、Stripe、Snowflake、Slack、Notion、OpenAPI、Infura 等数据库与服务的包装器。不同包装器的读写支持差异很大，但 `WHERE`、`ORDER BY` 和 `LIMIT` 下推是框架的核心能力。

### 版本说明

`v0.6.1` 保持了相同的扩展模型，但扩展了包装器目录和行为。官方发布说明特别提到：

- 新增 DynamoDB FDW 支持
- 通过 `mysql_fdw` 支持 MySQL/Doris
- `iceberg_fdw` 支持模式演进
- `_id` 选项可按名称查找保险库密钥
- 支持 COUNT、SUM、AVG、MIN 和 MAX 聚合下推，包括 MySQL FDW
- 修复参数状态刷新与重新扫描，并更新依赖和安全性

### 注意事项

- 各包装器的选项、支持的对象和写能力差异很大；使用时应查阅官方目录页，确认具体 FDW 的能力。
- 文档警告，如果物化视图依赖外部表，逻辑恢复可能失败，因此应避免这种模式，或依赖物理备份。
