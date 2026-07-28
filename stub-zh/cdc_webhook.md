## 用法

来源：

- [官方上游 README](https://github.com/hasura/postgres-cdc-plugin/blob/e08af17fe6f2a71fb1f181d1647e9bcff27b7f58/README.md)
- [官方扩展控制文件 (cdc_webhook.control)](https://github.com/hasura/postgres-cdc-plugin/blob/e08af17fe6f2a71fb1f181d1647e9bcff27b7f58/cdc_webhook.control)
- [官方扩展 SQL (cdc_webhook--1.0.sql)](https://github.com/hasura/postgres-cdc-plugin/blob/e08af17fe6f2a71fb1f181d1647e9bcff27b7f58/cdc_webhook--1.0.sql)

`cdc_webhook` — 一个启用变更数据捕获（CDC）的 PostgreSQL 扩展，通过发送 webhook 通知来监控数据库更改。该扩展用 C 和 SQL 编写，允许实时监控指定表上的 INSERT、UPDATE 和 DELETE 操作。在从 PostgreSQL 移动、转换或集成相应数据时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION cdc_webhook;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `call_webhook(payload JSONB, webhook_url TEXT, headers JSONB, timeout INT, cancel_on_failure BOOLEAN, retry_number INT, retry_interval INT, retry_backoff TEXT)` 是一个扩展函数，返回 `void`。
- `cdc_webhook.credentials_audit_trigger()` 是一个扩展函数，返回 `TRIGGER`。
- `cdc_webhook.event_log_audit_trigger()` 是一个扩展函数，返回 `TRIGGER`。
- `create_event_trigger` 是一个扩展函数。
- `cdc_webhook.credentials` 是一个由扩展安装或管理的表。
- `cdc_webhook.event_log` 是一个由扩展安装或管理的表。
- `cdc_webhook` 是一个由扩展创建的模式。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
