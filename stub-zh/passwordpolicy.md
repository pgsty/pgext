


## 用法

来源：[README](https://github.com/fmbiete/passwordpolicy/blob/v2.0.5/README.md)、[v2.0.5 release](https://github.com/fmbiete/passwordpolicy/releases/tag/v2.0.5)、[control file](https://github.com/fmbiete/passwordpolicy/blob/v2.0.5/passwordpolicy.control)

`passwordpolicy` 是 PostgreSQL `passwordcheck` module 的可配置替代方案。它会在 `CREATE ROLE` 和 `ALTER ROLE` 时检查密码，可以强制执行 password history 和 validity rules，也能在重复登录失败后模拟 soft account lock。

### 启用 Hook

在其他 password-check modules 前加载该 module，然后重启 PostgreSQL：

```conf
shared_preload_libraries = 'passwordpolicy'
```

使用 account soft-lock 或 password-history 功能时，在 `postgres` 数据库中安装 SQL 扩展：

```sql
CREATE EXTENSION passwordpolicy;
```

### 密码复杂度

配置项是动态的，但新值只对新 session 生效：

```conf
password_policy.min_password_len = 15
password_policy.min_special_chars = 1
password_policy.min_numbers = 1
password_policy.min_uppercase_letter = 1
password_policy.min_lowercase_letter = 1
password_policy.require_validuntil = off
```

只有创建 dictionary file 之后，才启用 CrackLib dictionary checks：

```conf
password_policy.cracklib_dictpath = '/var/cache/cracklib/postgresql_dict'
password_policy.enable_dictionary_check = on
```

### 软账号锁定

Soft-locking 会跟踪登录失败次数，并在超过阈值后延迟或拒绝响应：

```conf
password_policy_lock.number_failures = 5
password_policy_lock.failure_delay = 5
password_policy_lock.auto_unlock = on
password_policy_lock.auto_unlock_after = 0
password_policy_lock.max_number_accounts = 100
```

查看并重置 lock state：

```sql
SELECT * FROM passwordpolicy.accounts_locked() ORDER BY usename;
SELECT passwordpolicy.account_locked_reset('app_user');
```

如果 `password_policy_lock.include_all = false`，只有列在 `passwordpolicy.accounts_lockable` 中的 roles 才参与 soft-lock。

### 密码历史

Password history 会在 `postgres` 数据库中存储最近的 password hashes，并检查新密码是否复用：

```conf
password_policy_history.max_password_history = 5
password_policy_history.max_number_accounts = 100
```

### 注意事项

- 版本 2.0.5 支持 PostgreSQL 14-18。
- 该 module 必须 preload；修改 `shared_preload_libraries` 需要重启。
- PostgreSQL 无法在认证发生前真正阻止认证，因此 soft-lock 通过延迟和返回错误来模拟 lock。它不能缓解 authentication DoS attacks。
- 合理设置 `password_policy_lock.max_number_accounts` 和 `password_policy_history.max_number_accounts`，避免浪费内存或漏掉账号。
