## 用法

来源：

- [官方上游 README](https://github.com/bisoftbilgi/bisoft-postgresql-toolkit/blob/9a49035d2d370bdd8daefe7e8bdf67d29e0f648d/password_profile/README.md)
- [官方扩展控制文件 (password_profile.control)](https://github.com/bisoftbilgi/bisoft-postgresql-toolkit/blob/9a49035d2d370bdd8daefe7e8bdf67d29e0f648d/password_profile/password_profile.control)
- [官方实现源代码](https://github.com/bisoftbilgi/bisoft-postgresql-toolkit/blob/9a49035d2d370bdd8daefe7e8bdf67d29e0f648d/password_profile/src/lib.rs)

`password_profile` — 企业级 PostgreSQL 密码策略和身份验证加固。使用 Rust + pgrx (v0.16.1) 构建。已在 PostgreSQL 16-18 上测试。在实现相应的安全、审计或访问控制工作流时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION password_profile;
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `add_to_blacklist` 是一个扩展函数。
- `check_password` 是一个扩展函数。
- `check_password_expiry` 是一个扩展函数。
- `check_user_access` 是一个扩展函数。
- `clear_login_attempts` 是一个扩展函数。
- `get_lock_cache_stats()` 是一个扩展函数。
- `get_password_stats` 是一个扩展函数。
- `init_login_attempts_table()` 是一个扩展函数。
- `is_user_locked` 是一个扩展函数。
- `load_blacklist_from_file` 是一个扩展函数。
- `record_failed_login` 是一个扩展函数。
- `record_password_change` 是一个扩展函数。
- `remove_from_blacklist` 是一个扩展函数。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件标记该扩展为不受信任。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
