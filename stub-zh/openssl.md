## 用法

来源：

- [官方扩展 SQL](https://api.pgxn.org/src/openssl/openssl-0.0.1/sql/openssl.sql)
- [官方 C 实现](https://api.pgxn.org/src/openssl/openssl-0.0.1/src/pg_openssl.c)
- [官方扩展控制文件](https://api.pgxn.org/src/openssl/openssl-0.0.1/openssl.control)

`openssl` 版本 `0.0.1` 是一个早期 PGXN 实验，尝试列出服务器所链接 OpenSSL 库公开的算法。尽管名称如此，它并不提供任何 SQL 加密、解密、哈希、签名或密钥管理操作。

### 公开对象

扩展声明 `cipher_info` 与 `digest_info` 复合类型，以及两个返回集合的函数：

```sql
CREATE EXTENSION openssl;

SELECT * FROM list_ciphers();
SELECT * FROM list_digests();
```

`cipher_info` 预期报告算法名称、块大小、密钥长度、IV 长度、模式及 OpenSSL 标志等字段。`digest_info` 预期报告名称、块大小、摘要长度、公钥类型、标志和数字标识符。

### 兼容性与安全边界

已核验的 C 实现使用旧式 OpenSSL 枚举 API，并且固定最多存储 100 个算法。更严重的是，`list_digests()` 强制只输出一行，而格式化该行的代码被禁用，导致结果构造不完整。应将摘要输出视为不可靠，不能依赖此扩展作资产清点或合规决策。

源码还包含一个未暴露的密钥类型枚举函数，测试中使用的函数名也与实际安装的 SQL 接口不一致。只有在核验具体 PostgreSQL 与 OpenSSL 构建后，才能在可丢弃的兼容性实验环境中使用 `openssl`。生产密码学应使用有算法与密钥处理文档的受维护 PostgreSQL 功能或应用库。
