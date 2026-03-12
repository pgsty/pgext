

## 用法

> [wrappers: Supabase 开发的外部数据包装器](https://github.com/supabase/wrappers)

Supabase Wrappers 是一个构建 PostgreSQL 外部数据包装器（FDW）的框架，提供 30 多个预构建的连接器，用于云服务、数据库和 API。支持 WHERE、ORDER BY 和 LIMIT 下推，部分包装器还支持数据修改（INSERT/UPDATE/DELETE）。

### 可用包装器

| 类别 | 包装器 |
|----------|----------|
| **数据库** | ClickHouse、BigQuery、Snowflake、DuckDB、SQL Server、Redis、PostgreSQL |
| **存储** | AWS S3、Cloudflare D1、Apache Iceberg |
| **SaaS/API** | Stripe、Firebase、Airtable、Auth0、Notion、Slack、HubSpot、Shopify |
| **认证** | AWS Cognito、Clerk、Auth0 |
| **其他** | OpenAPI、Logflare、Calendly、Cal.com、Paddle、Orb、Infura、Gravatar |

### 示例（Stripe）

```sql
CREATE EXTENSION wrappers;

CREATE SERVER stripe_server
  FOREIGN DATA WRAPPER stripe_wrapper
  OPTIONS (
    api_key_id '<key_ID>',
    api_url 'https://api.stripe.com/v1/',
    api_version '2024-06-20'
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

SELECT id, email, name FROM stripe_customers WHERE email LIKE '%@example.com';
```

`rowid_column` 选项是 INSERT/UPDATE/DELETE 支持所必需的。`attrs` 列提供以 JSON 形式访问额外元数据的能力。

每个包装器使用自己的 `FOREIGN DATA WRAPPER` 名称（例如 `stripe_wrapper`、`firebase_wrapper`、`clickhouse_wrapper`），但它们都通过单一的 `wrappers` 扩展安装。请参阅[文档](https://fdw.dev)了解各包装器的特定选项和支持的对象。
