## 用法

来源：

- [Official database.dev 包页面](https://database.dev/asghonim/pgho_inbox)

`asghonim@pgho_inbox` — 入站联系和收件箱框架，包含验证、速率限制、垃圾邮件评分、工作流状态和通知事件。当应用程序需要此特定数据库功能时，请使用它。在目标 PostgreSQL 构建中测试上游版本链接中的固定修订版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION "asghonim@pgho_inbox";
```

在目标数据库中安装扩展，当可用时运行上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `add_attachment(p_message_id uuid, p_storage_provider text, p_storage_key text, p_mime text DEFAULT NULL, p_size bigint DEFAULT NULL)` 是一个扩展函数，返回 `uuid`。
- `add_note(p_message_id uuid, p_author text, p_body text)` 是一个扩展函数，返回 `uuid`。
- `assign_message(p_message_id uuid, p_assignee text, p_author text DEFAULT NULL)` 是一个扩展函数，返回 `void`。
- `check_rate_limit(p_ip inet, p_max_count integer DEFAULT 10)` 是一个扩展函数，返回 `boolean`。
- `claim_notifications(p_limit integer DEFAULT 10, p_type text DEFAULT NULL)` 是一个扩展函数，返回 `SETOF`。
- `close_message(p_message_id uuid, p_author text DEFAULT NULL, p_reason text DEFAULT NULL)` 是一个扩展函数，返回 `void`。
- `create_channel(p_name text, p_description text DEFAULT NULL, p_settings jsonb DEFAULT '{}')` 是一个扩展函数，返回 `uuid`。
- `mark_notification_failed(p_notification_id uuid, p_error text DEFAULT NULL, p_max_attempts integer DEFAULT 3)` 是一个扩展函数，返回 `void`。
- `mark_notification_sent(p_notification_id uuid)` 是一个扩展函数，返回 `void`。
- `mark_spam(p_message_id uuid, p_author text DEFAULT NULL)` 是一个扩展函数，返回 `void`。
- `open_message(p_message_id uuid, p_author text DEFAULT NULL)` 是一个扩展函数，返回 `void`。
- `register_hook(p_event text, p_function_name text)` 是一个扩展函数，返回 `uuid`。
- `reopen_message(p_message_id uuid, p_author text DEFAULT NULL)` 是一个扩展函数，返回 `void`。
- `requeue_stale_notifications(p_timeout interval DEFAULT '30 minutes')` 是一个扩展函数，返回 `integer`。

### 要求与注意事项

- 该目录记录版本 `0.0.3`。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 身份之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行验证。
