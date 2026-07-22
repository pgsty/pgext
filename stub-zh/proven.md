## 用法

来源：

- [已复核 commit 的 PostgreSQL 绑定](https://github.com/hyperpolymath/proven/blob/302f2a473fb7baf522aecf59eaa60f4e7c4d966e/bindings/sql/postgresql/proven--0.5.0.sql)
- [PostgreSQL 示例](https://github.com/hyperpolymath/proven/blob/302f2a473fb7baf522aecf59eaa60f4e7c4d966e/bindings/sql/postgresql/proven_functions.sql)
- [项目文档](https://github.com/hyperpolymath/proven/blob/302f2a473fb7baf522aecf59eaa60f4e7c4d966e/README.adoc)

`proven` 为 libproven 的 Idris 2/Zig 安全例程提供了一小组 PostgreSQL C 绑定。0.5.0 版本包含受检 `BIGINT` 算术、简化的电子邮件/URL/IPv4/JSON 校验、HTML 实体转义、十六进制转换及校验和辅助函数。

```sql
CREATE EXTENSION proven;
SELECT proven_safe_add(9223372036854775807, 1);
SELECT proven_validate_email('user@example.com');
SELECT proven_hex_encode('hello');
```

算术函数遇到溢出、下溢、除零或其他错误时返回 `NULL`，而不是抛出异常。校验器只确认上游解析器的简化语法，不证明地址所有权、可投递性、URL 安全性、可路由性或应用策略合规。

尽管名称如此，`proven_hash_sha256(text)` 返回的是十六进制 CRC32 校验和，而非 SHA-256。不要将其用于密码、签名、对抗攻击者的完整性或内容寻址安全；应使用 `pgcrypto` 等受维护的密码学实现。HTML 转义取决于输出上下文，并非通用 XSS 清洗器。

应让原生 libproven ABI 与扩展构建保持固定，并测试无效 UTF-8、大输入、`NULL`、溢出和升级行为。核心例程的形式化验证并不覆盖 PostgreSQL C 包装器、FFI、构建链或周边应用。
