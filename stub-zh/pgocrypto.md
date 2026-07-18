## 用法

来源：

- [上游 README](https://github.com/johto/pgocrypto/blob/87fa6cc60064ad0c90cd023b5727433e6eeaf41a/README.md)
- [1.0 版安装 SQL](https://github.com/johto/pgocrypto/blob/87fa6cc60064ad0c90cd023b5727433e6eeaf41a/pgocrypto--1.0.sql)
- [Go 互操作实现](https://github.com/johto/pgocrypto/blob/87fa6cc60064ad0c90cd023b5727433e6eeaf41a/pgocrypto.go)

pgocrypto 为 PostgreSQL 和 Go 提供配对的对称加密函数。PostgreSQL 函数使用 AES-CBC 加密 bytea 或 text，并在前面添加随机的 16 字节初始化向量；字符串密文使用 Base64 编码。配套的 Go 代码使用相同的线上格式。

### 基本用法

该扩展以不带模式限定的方式调用 pgcrypto 函数，且没有声明依赖。应将两个扩展安装在同一个受控模式中，并设置安全的搜索路径：

```sql
CREATE SCHEMA crypto;
CREATE EXTENSION pgcrypto WITH SCHEMA crypto;
CREATE EXTENSION pgocrypto WITH SCHEMA crypto;

SET search_path = crypto, pg_catalog;

WITH k AS (
    SELECT decode(repeat('01', 32), 'hex') AS key
)
SELECT pgo_decrypt_string(
           pgo_encrypt_string('hello', key),
           key
       )
FROM k;
```

示例密钥是确定性测试数据。生产密钥必须来自专门的密钥管理设计，不得嵌入 SQL、日志、模式对象或应用源码。

### 注意事项

- 这里的 AES-CBC 加密没有认证。攻击者可以修改密文而不会触发内置完整性失败，因此该格式可被篡改，也不是 AEAD 封装。
- 该 API 接受原始密钥字节，不提供密码派生、密钥标识符、轮换元数据、版本化、访问控制或安全存储。
- 解密错误与应用行为可能形成填充预言机风险。不得向不可信调用者暴露可区分的密码学失败。
- 新建的安全敏感系统应优先使用维护良好的认证加密格式。如果互操作要求该遗留格式，应增加经独立验证的认证层与有文档的迁移计划。
- 应控制函数搜索路径，因为不带模式限定的 pgcrypto 调用否则依赖会话配置。应限制执行权，并避免记录明文或密钥参数。
