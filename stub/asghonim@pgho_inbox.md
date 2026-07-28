## Usage

Sources:

- [Official database.dev package page](https://database.dev/asghonim/pgho_inbox)

`asghonim@pgho_inbox` — Inbound contact and inbox framework with validation, rate limiting, spam scoring, workflow state, and notification events. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "asghonim@pgho_inbox";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `add_attachment(p_message_id uuid, p_storage_provider text, p_storage_key text, p_mime text DEFAULT NULL, p_size bigint DEFAULT NULL)` is an extension function and returns `uuid`.
- `add_note(p_message_id uuid, p_author text, p_body text)` is an extension function and returns `uuid`.
- `assign_message(p_message_id uuid, p_assignee text, p_author text DEFAULT NULL)` is an extension function and returns `void`.
- `check_rate_limit(p_ip inet, p_max_count integer DEFAULT 10)` is an extension function and returns `boolean`.
- `claim_notifications(p_limit integer DEFAULT 10, p_type text DEFAULT NULL)` is an extension function and returns `SETOF`.
- `close_message(p_message_id uuid, p_author text DEFAULT NULL, p_reason text DEFAULT NULL)` is an extension function and returns `void`.
- `create_channel(p_name text, p_description text DEFAULT NULL, p_settings jsonb DEFAULT '{}')` is an extension function and returns `uuid`.
- `mark_notification_failed(p_notification_id uuid, p_error text DEFAULT NULL, p_max_attempts integer DEFAULT 3)` is an extension function and returns `void`.
- `mark_notification_sent(p_notification_id uuid)` is an extension function and returns `void`.
- `mark_spam(p_message_id uuid, p_author text DEFAULT NULL)` is an extension function and returns `void`.
- `open_message(p_message_id uuid, p_author text DEFAULT NULL)` is an extension function and returns `void`.
- `register_hook(p_event text, p_function_name text)` is an extension function and returns `uuid`.
- `reopen_message(p_message_id uuid, p_author text DEFAULT NULL)` is an extension function and returns `void`.
- `requeue_stale_notifications(p_timeout interval DEFAULT '30 minutes')` is an extension function and returns `integer`.

### Requirements and Caveats

- The catalog records version `0.0.3`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
