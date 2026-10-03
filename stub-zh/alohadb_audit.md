## 用法

来源：

- [contrib/alohadb_audit/alohadb_audit.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_audit/alohadb_audit.control)
- [contrib/alohadb_audit/alohadb_audit.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_audit/alohadb_audit.c)
- [contrib/alohadb_audit/alohadb_audit--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_audit/alohadb_audit--1.0.sql)
- [contrib/alohadb_audit/alohadb_audit--1.0--1.1.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_audit/alohadb_audit--1.0--1.1.sql)

`alohadb_audit` 1.1 在所引用的 AlohaDB 源码中记录 DML 与 DDL 活动。该扩展与定制内核关联，本文不声称其兼容普通 PostgreSQL。

### 启用

超级用户必须预加载该库并重启服务器。将其并入已有预加载列表，再在固定的 `public` 模式中创建扩展。

```ini
shared_preload_libraries = 'alohadb_audit'
alohadb.audit_enabled = on
alohadb.audit_log_format = 'json'
```

```sql
CREATE EXTENSION alohadb_audit;
SELECT * FROM audit_log_status();
```

### 配置与运维

`alohadb.audit_log_directory` 指定服务端可写的日志目录。`alohadb.audit_databases` 与 `alohadb.audit_operations` 筛选事件；`alohadb.audit_log_query_text` 控制是否记录 SQL 文本。`audit_log_status()` 返回当前配置。可选加密使用 OpenSSL AES-256-GCM，密钥通过 64 位十六进制字符串形式的 `alohadb.audit_encryption_key` 配置。`audit_decrypt_log(line, key, redact)` 解密指定日志行，并可按模式隐藏密码。

应保护日志文件、密钥配置与查询文本；将密钥作为 SQL 字面量传入也可能使其进入日志。脱敏依赖模式匹配，不能保证去除全部机密。日志保留需另行安排，并应在所用的具体 AlohaDB 构建中验证钩子与加密行为，再依赖这些审计记录。
