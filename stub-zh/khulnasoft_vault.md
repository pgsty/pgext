## 用法

来源：

- [官方上游 README](https://github.com/khulnasoft/vault/blob/9f472c76a7cd7d9847175a6718853c1510fc911c/README.md)
- [官方扩展控制文件 (khulnasoft_vault.control)](https://github.com/khulnasoft/vault/blob/9f472c76a7cd7d9847175a6718853c1510fc911c/khulnasoft_vault.control)
- [官方扩展 SQL (khulnasoft_vault--0.2.8.sql)](https://github.com/khulnasoft/vault/blob/9f472c76a7cd7d9847175a6718853c1510fc911c/sql/khulnasoft_vault--0.2.8.sql)

`khulnasoft_vault` — Khulnasoft 提供了一个名为 vault.secrets 的表，可以用来存储敏感信息，如 API 密钥。在实现相应的安全、审计或访问控制工作流时，请使用它。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION khulnasoft_vault;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `vault.create_secret(new_secret text, new_name text = NULL, new_description text = '', new_key_id uuid = NULL)` 是一个扩展函数，返回 `uuid`。
- `vault.update_secret(secret_id uuid, new_secret text = NULL, new_name text = NULL, new_description text = NULL, new_key_id uuid = NULL)` 是一个扩展函数，返回 `void`。
- `vault.secrets` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 控制文件声明默认版本为 `0.2.8`。
- 在生产使用之前，请先安装并验证确认的扩展依赖项：`pgsodium`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
