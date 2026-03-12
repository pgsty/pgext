

## 用法

> [sslutils: 通过 SQL 管理 SSL 证书](https://github.com/EnterpriseDB/sslutils)

`sslutils` 是一个通过 SQL 命令管理 SSL 证书的 PostgreSQL 扩展。它提供直接在数据库内生成、检查和管理 SSL/TLS 证书的函数。

```sql
CREATE EXTENSION sslutils;
```

### 函数

该扩展提供以下 SSL 证书管理的 SQL 函数：

| 函数 | 描述 |
|----------|-------------|
| `openssl_rsa_generate_key(bits int)` | 生成 RSA 私钥 |
| `openssl_rsa_key_to_csr(key text, cn text, ...)` | 生成证书签名请求 |
| `openssl_csr_to_crt(csr text, ca_key text, ca_crt text)` | 签署 CSR 生成证书 |
| `openssl_rsa_generate_crl(ca_key text, ca_crt text)` | 生成证书吊销列表 |
| `ssl_is_init_fn()` | 检查 SSL 是否已初始化 |
| `ssl_get_cipher_fn()` | 获取当前 SSL 密码套件 |
| `ssl_get_version_fn()` | 获取当前 SSL 版本 |

### 典型工作流程

```sql
-- 生成 CA 私钥
SELECT openssl_rsa_generate_key(2048);

-- 创建自签名 CA 证书
-- 生成服务器密钥和 CSR
-- 使用 CA 签署 CSR
```

该扩展适用于在托管 PostgreSQL 环境中自动化 SSL 证书配置。
