## 用法

来源：

- [官方上游 README](https://github.com/telkomdev/pgcrypsi/blob/ee909d7ad315aa5c5178687db542e5fc73959d84/README.md)
- [官方扩展控制文件 (pgcrypsi.control)](https://github.com/telkomdev/pgcrypsi/blob/ee909d7ad315aa5c5178687db542e5fc73959d84/pgcrypsi.control)
- [官方扩展 SQL (pgcrypsi--0.0.1.sql)](https://github.com/telkomdev/pgcrypsi/blob/ee909d7ad315aa5c5178687db542e5fc73959d84/pgcrypsi--0.0.1.sql)

`pgcrypsi` — C Crypsi PostgreSQL 扩展。在实现相应的安全、审计或访问控制工作流时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgcrypsi;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `pgcrypsi_aes_128_gcm_decrypt(text, text)` 是一个扩展函数，返回 `text`。
- `pgcrypsi_aes_128_gcm_encrypt(text, text)` 是一个扩展函数，返回 `text`。
- `pgcrypsi_aes_192_gcm_decrypt(text, text)` 是一个扩展函数，返回 `text`。
- `pgcrypsi_aes_192_gcm_encrypt(text, text)` 是一个扩展函数，返回 `text`。
- `pgcrypsi_aes_256_gcm_decrypt(text, text)` 是一个扩展函数，返回 `text`。
- `pgcrypsi_aes_256_gcm_encrypt(text, text)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
