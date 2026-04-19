## 用法

来源: [official README](https://github.com/supabase/wrappers/blob/main/README.md), [official docs](https://fdw.dev/), [v0.6.0 release](https://github.com/supabase/wrappers/releases/tag/v0.6.0)

`wrappers` 既是一个用 Rust 编写 PostgreSQL foreign data wrapper 的框架，也是 Supabase 维护的一组 FDW 打包集合。单个扩展会安装多种 wrapper 实现，然后每个 foreign server 再选择自己需要的具体 wrapper 类型。

```sql
CREATE EXTENSION wrappers;
```

### 典型流程

先为某个 wrapper 创建 server，再通过 foreign table 暴露远端数据：

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

上游提供了用于 BigQuery、ClickHouse、DuckDB、MySQL、Redis、S3、Stripe、Snowflake、Slack、Notion、OpenAPI、Infura 等数据库与服务的 wrappers。不同 wrapper 的读写支持差异很大，但 `WHERE`、`ORDER BY` 和 `LIMIT` 下推是框架的核心能力。

### 版本说明

`v0.6.0` 保持了相同的扩展模型，但扩展了 wrapper 目录和行为。官方发布说明特别提到：

- 新增 OpenAPI FDW 支持
- 新增 Infura FDW 支持
- Snowflake `timeout_secs` table option
- 多个 wrapper 的写入路径和扫描修复

### 注意事项

- 各 wrapper 的选项、支持的对象和写能力差异很大；使用时应查阅官方目录页，确认具体 FDW 的能力。
- 文档警告，如果 materialized view 依赖 foreign table，logical restore 可能失败，因此应避免这种模式，或依赖物理备份。
