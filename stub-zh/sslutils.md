## 用法

来源：

- [1.4.1 README](https://github.com/EnterpriseDB/sslutils/blob/v1.4.1/README.sslutils)
- [1.4.1 SQL API](https://github.com/EnterpriseDB/sslutils/blob/v1.4.1/sslutils--1.4.1.sql)
- [1.4.1 control file](https://github.com/EnterpriseDB/sslutils/blob/v1.4.1/sslutils.control)
- [File-access checks](https://github.com/EnterpriseDB/sslutils/blob/v1.4.1/sslutils.c)

`sslutils` 为 Postgres Enterprise Manager 及受控管理流程提供 SSL 证书生成和吊销辅助函数。部分函数会读取或修改数据库服务器上的文件。

### 安装和检查

```sql
CREATE EXTENSION sslutils;
SELECT sslutils_version();
SELECT openssl_get_crt_expiry_date('server.crt');
```

这个不可重定位的扩展需要超级用户安装。调用函数无需预加载或重启 PostgreSQL；部署生成的证书是另外的服务器配置操作。

### 主要函数

| 函数 | 用途 |
|---|---|
| `openssl_rsa_generate_key(bits)` | 以文本返回 RSA 私钥 |
| `openssl_rsa_key_to_csr(key, cn, country, state, location, unit, email)` | 返回 CSR |
| `openssl_csr_to_crt(csr, ca_cert_path, private_key_path, days)` | 使用服务器证书和私钥文件签署 CSR |
| `openssl_rsa_generate_crl(ca_cert_path, ca_key_path, days)` | 返回证书吊销列表 |
| `openssl_is_crt_expire_on(cert_path, at_time)` | 检查证书到期情况，返回 1、-1 或 0 |
| `openssl_get_crt_expiry_date(cert_path)` | 返回到期时间戳 |
| `openssl_revoke_certificate(cert_pem, crl_path, days)` | 吊销证书并重新生成 CRL |

生成自签名证书时，签署函数的证书路径参数为 NULL，私钥路径指定签名密钥。可选有效期默认是 3650 天。文件路径指向 PostgreSQL 服务器，而非 SQL 客户端机器。

### 保护密钥和服务器文件

在 PostgreSQL 11 及以上，证书检查函数需要 `pg_read_server_files` 成员资格；吊销操作同时检查该角色和 `pg_write_server_files`。这些角色具有广泛的服务器文件访问能力，应仅授予可信管理员，并按需限制函数执行权限。

以 SQL 文本返回的私钥可能进入查询结果、日志、命令历史或应用追踪。应保护它们的传递和存储。这个扩展不提供 SSL 会话检查函数，也不会自动配置 PostgreSQL 的 TLS 设置。

### 1.4.1 的文件访问变化

证书检查和 CA 签名路径会检查是否位于服务器数据目录内；应将输入文件放在该目录，并在适当时使用服务器相对路径。吊销要求设置 `sslutils.revoke_certificate_crl_paths`，它是在 postgresql.conf 中配置、重载后生效的 CRL 输出路径前缀逗号分隔允许列表。吊销检查仅比较字符串前缀，不解析规范化路径来验证目录包含关系；应配置范围较窄的前缀，并保护输出位置。

吊销实现的第一个参数接收 PEM 证书文本，尽管较早 README 和 SQL 注释描述为证书路径。它读取服务器上的 CA 文件 `ca_certificate.crt`、`ca_key.key`，并创建或追加 `revoke_cert.db`。操作前应准备好 CA 文件。不要认为广泛的服务器文件角色能够绕过扩展自身的检查。
