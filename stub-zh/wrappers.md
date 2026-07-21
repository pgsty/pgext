## 用法

来源：

- [Wrappers v0.6.2 README](https://github.com/supabase/wrappers/blob/v0.6.2/README.md)
- [官方FDW文档](https://fdw.dev/)
- [v0.6.2版本发布说明](https://github.com/supabase/wrappers/releases/tag/v0.6.2)
- [MongoDB FDW文档](https://fdw.dev/catalog/mongodb/)
- [安全指导](https://fdw.dev/guides/security/)

`wrappers` 是一个用 Rust 编写的 PostgreSQL 外部数据封装框架，同时也是一系列由 Supabase 维护的 FDWs 的打包集合。安装一个扩展即可实现多种封装器的实现，每个外部服务器可以选择它需要的具体封装器类型。

```sql
CREATE EXTENSION wrappers;
```

### 通常的工作流程

为一个封装器创建一个服务器，然后通过外部表暴露远程数据：

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

### 包含的内容

上游版本提供了对 BigQuery、ClickHouse、DuckDB、DynamoDB、MySQL/Doris、Redis、S3、S3 Vectors、Stripe、Snowflake、Slack、Notion、OpenAPI、Infura 等数据库和服务的封装器。读写支持因封装器而异，但 `WHERE`、`ORDER BY` 和 `LIMIT` 的下推是框架的核心功能。

### 版本 0.6.2

`v0.6.2`版本保持了相同的扩展模型，并增加了以下内容：

- MongoDB FDW 支持读写操作
- 在 WASM 封装器中使用会话变量凭证进行每次请求的身份验证
- OpenAPI FDW 的 RFC 8288 `Link` 头部分页支持
- 运行时、依赖和封装器特定的修复，详情参见发布说明

每个封装器的具体页面仍然是服务器选项、外部表列、下推和支持写操作的权威来源。

### 注意事项

- 封装器特定的选项、支持的对象以及写操作因封装器而异；请查阅官方目录页面以获取您使用的具体 FDW 的确切信息。
- 文档警告说，当物化视图依赖外部表时，逻辑恢复可能会失败，请避免这种模式或依赖于物理备份。
- 外部表本身不提供安全边界。请将它们保留在私有模式中，并谨慎授予访问权限，使用可用的最小特权远程凭证，并对暴露或缓存远程数据的本地表应用行级安全性。
- 请将 API 密钥和令牌存储在支持的秘密存储中或每次请求的凭据机制中，而不是将其嵌入到版本控制系统中的 SQL 中。
