


## 用法

> [sslinfo: SSL 证书信息函数](https://www.postgresql.org/docs/current/sslinfo.html)

sslinfo 提供函数来访问当前数据库连接中使用的 SSL 证书信息。

### 函数

```sql
-- 检查当前连接是否使用 SSL
SELECT ssl_is_used();

-- SSL/TLS 协议版本（TLSv1.2、TLSv1.3 等）
SELECT ssl_version();

-- 密码套件名称（例如 DHE-RSA-AES256-SHA）
SELECT ssl_cipher();

-- 检查客户端是否提供了有效证书
SELECT ssl_client_cert_present();

-- 客户端证书序列号
SELECT ssl_client_serial();

-- 客户端证书主题（完整 DN）
SELECT ssl_client_dn();
-- 例如 /CN=Somebody /C=Some country/O=Some organization

-- 证书颁发者（完整 DN）
SELECT ssl_issuer_dn();

-- 客户端证书主题的特定字段
SELECT ssl_client_dn_field('CN');
SELECT ssl_client_dn_field('O');

-- 证书颁发者的特定字段
SELECT ssl_issuer_field('CN');

-- 客户端证书扩展
SELECT * FROM ssl_extension_info();
-- 返回：name、value、critical
```

### 注意事项

- 如果连接未使用 SSL，大多数函数返回 NULL
- 需要 PostgreSQL 使用 OpenSSL 支持编译
- `ssl_client_serial()` 和 `ssl_issuer_dn()` 的组合可以唯一标识一个证书
