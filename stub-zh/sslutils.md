## 用法

来源：

- [1.4.2 README](https://github.com/EnterpriseDB/sslutils/blob/v1.4.2/README.sslutils)
- [1.4.2 SQL API](https://github.com/EnterpriseDB/sslutils/blob/v1.4.2/sslutils--1.4.2.sql)
- [1.4.2 control file](https://github.com/EnterpriseDB/sslutils/blob/v1.4.2/sslutils.control)
- [File-access checks](https://github.com/EnterpriseDB/sslutils/blob/v1.4.2/sslutils.c)
- [1.4.1 to 1.4.2 upgrade](https://github.com/EnterpriseDB/sslutils/blob/v1.4.2/sslutils--1.4.1--1.4.2.sql)
- [PEM agent grants](https://github.com/EnterpriseDB/sslutils/blob/v1.4.2/sslutils-pemagent.sql.in)

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

在 1.4.2 中，全部七个 `openssl_*` 函数采用 `SECURITY DEFINER`，并将 `search_path` 固定为 `pg_catalog, pg_temp`。其所有者 `sslutils_owner` 是集群级的 `INHERIT`、`NOLOGIN` 角色。在 PostgreSQL 11 及以上，它拥有 `pg_read_server_files` 和 `pg_write_server_files`，但不得拥有成员、管理权限属性或其他角色的成员资格。已有所有者角色不符合这些要求时，安装会失败。

`PUBLIC` 没有执行权限。安装或升级时，若已存在 `pem_agent`，则授予它 `EXECUTE`；已有的显式函数授权会保留。如果之后才创建代理角色，超级用户需要在扩展所在模式中执行已安装的 `sslutils-pemagent.sql`。使用非 public 模式的代理还需要该模式的 `USAGE`，但不应获得 `CREATE`。调用者需要相应函数权限，无需直接拥有服务器文件角色；绝不能授予它们 `sslutils_owner` 的成员资格。

仅向可信调用者授权，使其能够在允许路径内检查证书和使用 CA 密钥。限制在数据目录内并不能隔离不同代理或不同 CA。

以 SQL 文本返回的私钥可能进入查询结果、日志、命令历史或应用追踪。应保护它们的传递和存储。这个扩展不提供 SSL 会话检查函数，也不会自动配置 PostgreSQL 的 TLS 设置。

### 升级和已有授权

以超级用户升级已有安装：

```sql
ALTER EXTENSION sslutils UPDATE TO '1.4.2';
```

迁移调用者后，撤销代理原有的广泛文件角色授权：

```sql
REVOKE pg_read_server_files, pg_write_server_files FROM pem_agent;
```

还要检查各登录角色直接获得和继承的授权，包括 `pg_execute_server_program` 与超级用户权限。单个数据库内的扩展升级不会自动撤销这些集群级授权。删除扩展也会保留共享的 `sslutils_owner` 角色；只有在所有数据库中卸载扩展并撤销其文件角色成员资格之后，才能移除该角色。

### 文件访问边界

证书检查和 CA 签名路径解析后必须位于服务器数据目录内。吊销还要求设置 `sslutils.revoke_certificate_crl_paths`，这是由超级用户在 postgresql.conf 中配置、重载后生效的 CRL 输出目录逗号分隔允许列表。在 1.4.2 中，检查会解析路径并验证目录边界；新输出文件通过其已有父目录进行检查，已有输出文件本身也会接受检查。

使用由管理员控制的专用 CRL 输出目录。不要允许整个数据目录或存放配置文件的目录，因为获得授权的调用者可以选择允许目录内的输出文件名。证书和私钥路径也必须由管理员控制。

吊销实现的第一个参数接收 PEM 证书文本，尽管较早 README 和 SQL 注释描述为证书路径。它读取服务器上的 CA 文件 `ca_certificate.crt`、`ca_key.key`，并创建或追加 `revoke_cert.db`。操作前应准备好 CA 文件。不要认为广泛的服务器文件角色能够绕过扩展自身的检查。
