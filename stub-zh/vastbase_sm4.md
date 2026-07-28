## 用法

来源：

- [官方上游 README](https://github.com/pangpang20/sm4prj/blob/b55e4567fd7b64cca2da238f10c9bd518021237a/vastbase-sm4-extension/README.md)
- [官方扩展控制文件 (vastbase_sm4.control)](https://github.com/pangpang20/sm4prj/blob/b55e4567fd7b64cca2da238f10c9bd518021237a/vastbase-sm4-extension/vastbase_sm4.control)
- [官方扩展 SQL (vastbase_sm4--1.0.sql)](https://github.com/pangpang20/sm4prj/blob/b55e4567fd7b64cca2da238f10c9bd518021237a/vastbase-sm4-extension/vastbase_sm4--1.0.sql)

`vastbase_sm4` — SM4 加密/解密函数，使用 Java 实现并通过 JNI 集成。在实现相应的安全、审计或访问控制工作流时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION vastbase_sm4;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `sm4_decrypt(cipher_text text, hex_key text)` 是一个扩展函数，返回 `text`。
- `sm4_decrypt_base64(cipher_text text, base64_key text)` 是一个扩展函数，返回 `text`。
- `sm4_encrypt(plain_text text, hex_key text)` 是一个扩展函数，返回 `text`。
- `sm4_encrypt_base64(plain_text text, base64_key text)` 是一个扩展函数，返回 `text`。
- `sm4_generate_key()` 是一个扩展函数，返回 `text`。
- `sm4_extension_info` 是一个扩展定义的视图。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
