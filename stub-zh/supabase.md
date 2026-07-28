## 用法

来源：

- [官方上游 README](https://github.com/constructive-io/supabase-test-suite/blob/44ceff6be0b62e1d1a60524acc0c65c1b8343726/packages/supabase/README.md)
- [官方扩展控制文件 (supabase.control)](https://github.com/constructive-io/supabase-test-suite/blob/44ceff6be0b62e1d1a60524acc0c65c1b8343726/packages/supabase/supabase.control)
- [官方扩展 SQL (supabase--0.0.1.sql)](https://github.com/constructive-io/supabase-test-suite/blob/44ceff6be0b62e1d1a60524acc0c65c1b8343726/packages/supabase/sql/supabase--0.0.1.sql)

`supabase` — Supabase 专用的 SQL、测试和辅助工具，用于构建具有行级安全性的稳健应用程序。在实现相应的安全、审计或访问控制工作流时使用它。在集成到应用程序 SQL 中之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION supabase;
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `auth.email()` 是一个扩展函数，返回 `text`。
- `auth.role()` 是一个扩展函数，返回 `text`。
- `auth.uid()` 是一个扩展函数，返回 `uuid`。
- `extensions.grant_pg_cron_access()` 是一个扩展函数，返回 `event_trigger`。
- `extensions.grant_pg_graphql_access()` 是一个扩展函数，返回 `event_trigger`。
- `extensions.grant_pg_net_access()` 是一个扩展函数，返回 `event_trigger`。
- `extensions.pgrst_ddl_watch()` 是一个扩展函数，返回 `event_trigger`。
- `extensions.pgrst_drop_watch()` 是一个扩展函数，返回 `event_trigger`。
- `extensions.set_graphql_placeholder()` 是一个扩展函数，返回 `event_trigger`。
- `graphql_public.graphql("operationName" text default null, query text default null, variables jsonb default null, extensions jsonb default null)` 是一个扩展函数，返回 `jsonb`。
- `pgbouncer.get_auth(p_usename text)` 是一个扩展函数，返回 `TABLE`。
- `storage.add_prefixes(_bucket_id text, _name text)` 是一个扩展函数，返回 `void`。
- `storage.delete_prefix(_bucket_id text, _name text)` 是一个扩展函数，返回 `boolean`。
- `storage.delete_prefix_hierarchy_trigger()` 是一个扩展函数，返回 `trigger`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 先安装并验证确认的扩展依赖项：`plpgsql`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
