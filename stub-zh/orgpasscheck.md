## 用法

来源：

- [官方上游 README](https://github.com/mbpcore/orgpasscheck/blob/4246a0f2f390a8a8f056758c09f05af5042ad23c/orgpasscheck_v5/README.md)
- [官方扩展控制文件 (orgpasscheck.control)](https://github.com/mbpcore/orgpasscheck/blob/4246a0f2f390a8a8f056758c09f05af5042ad23c/orgpasscheck_v5/orgpasscheck.control)
- [官方扩展 SQL (orgpasscheck--5.0.sql)](https://github.com/mbpcore/orgpasscheck/blob/4246a0f2f390a8a8f056758c09f05af5042ad23c/orgpasscheck_v5/orgpasscheck--5.0.sql)

`orgpasscheck` — **PostgreSQL 16+ 的企业级密码策略执行扩展**。在实现相应的安全、审计或访问控制工作流时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION orgpasscheck;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `orgpasscheck.add_blacklist(p_pattern TEXT, p_reason TEXT DEFAULT NULL)` 是一个扩展函数，返回 `VOID`。
- `orgpasscheck.add_expiry_exemption(p_username TEXT, p_reason TEXT DEFAULT NULL, p_expires_at TIMESTAMPTZ DEFAULT NULL)` 是一个扩展函数，返回 `VOID`。
- `orgpasscheck.change_password(p_username TEXT, p_password TEXT, p_expiry_days INTEGER DEFAULT NULL)` 是一个扩展函数，返回 `VOID`。
- `orgpasscheck.create_user(p_username TEXT, p_password TEXT, p_login BOOLEAN DEFAULT true, p_superuser BOOLEAN DEFAULT false, p_expiry_days INTEGER DEFAULT NULL)` 是一个扩展函数，返回 `VOID`。
- `orgpasscheck.list_expiry_exemptions()` 是一个扩展函数，返回 `TABLE`。
- `orgpasscheck.purge_audit_log(p_older_than INTERVAL DEFAULT '1 year')` 是一个扩展函数，返回 `INTEGER`。
- `orgpasscheck.purge_old_history()` 是一个扩展函数，返回 `INTEGER`。
- `orgpasscheck.purge_user_history(p_username TEXT)` 是一个扩展函数，返回 `INTEGER`。
- `orgpasscheck.record_password_history(p_username TEXT, p_password TEXT)` 是一个扩展函数，返回 `VOID`。
- `orgpasscheck.remove_blacklist(p_pattern TEXT)` 是一个扩展函数，返回 `VOID`。
- `orgpasscheck.remove_expiry_exemption(p_username TEXT)` 是一个扩展函数，返回 `VOID`。
- `orgpasscheck.verify_password_hash(p_password TEXT, p_salt TEXT, p_stored TEXT)` 是一个扩展函数，返回 `BOOLEAN`。
- `orgpasscheck.expired_passwords` 是一个扩展定义视图。
- `orgpasscheck.policy_summary` 是一个扩展定义视图。

### 要求与注意事项

- 控制文件声明默认版本为 `5.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
