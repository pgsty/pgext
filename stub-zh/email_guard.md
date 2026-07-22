## 用法

来源：

- [官方 0.35.0 版 README](https://github.com/mansueli/tle/blob/919de6247f4126dd11cd20e36602e0c5baf6fcaf/email_guard/README.md)
- [官方 0.35.0 版 control 文件](https://github.com/mansueli/tle/blob/919de6247f4126dd11cd20e36602e0c5baf6fcaf/email_guard/email_guard.control)
- [官方 0.35.0 版 SQL](https://github.com/mansueli/tle/blob/919de6247f4126dd11cd20e36602e0c5baf6fcaf/email_guard/email_guard--0.35.0.sql)
- [官方软件包页面](https://database.dev/mansueli/email_guard)

`email_guard` 提供邮箱规范化、一次性域名检查，以及 Supabase Auth 的用户创建前钩子。版本 `0.35.0` 可拦截一次性邮箱，并阻止仅在点号、加号标签或 `googlemail.com` 别名上不同的重复 Gmail 身份。

### 核心流程

把扩展安装到固定模式，然后可以直接调用其纯辅助函数：

```sql
CREATE EXTENSION email_guard SCHEMA extensions;

SELECT extensions.normalize_email('First.Last+news@gmail.com');
SELECT extensions.is_disposable_email('person@example.invalid');
```

在 Supabase Auth 中，将项目的用户创建前 Postgres Function 钩子配置为调用：

```text
extensions.hook_prevent_disposable_and_enforce_gmail_uniqueness
```

钩子接收 `jsonb` 格式的 Auth 事件。允许注册时返回空 JSON 对象；一次性域名返回 HTTP 风格的 `403` 错误；规范化后的 Gmail 身份已存在于 `auth.users` 时返回 `409`。没有邮箱的事件（例如手机号注册）会直接放行。

### API 与数据

- `normalize_email(text)`：邮箱转小写、去首尾空白，并规范化 Gmail/Googlemail 本地部分。
- `is_disposable_email_domain(text)`：检查预置黑名单，也会匹配父域名。
- `is_disposable_email(text)`：提取域名并执行黑名单检查。
- `hook_prevent_disposable_and_enforce_gmail_uniqueness(jsonb)`：Supabase Auth 钩子。
- `disposable_email_domains`：随版本发布的域名黑名单表。

### 注意事项

钩子会查询 `auth.users`，因此仅适用于 Supabase Auth 部署，并需要相应权限。Gmail 唯一性检查只覆盖钩子执行时已经存在的用户；应在真实注册路径中验证并发与错误处理。应通过升级扩展更新预置域名数据，不要依赖未公开的内部修改方式。

本文记录的目录目标是 `0.35.0`。上游 `main` 此后已更新到 `0.36.0`；如需复现这里的准确函数和黑名单内容，应固定已复核版本。
