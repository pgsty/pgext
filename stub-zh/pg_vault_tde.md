## 用法

来源：

- [pg_vault_tde 1.7.0 README](https://api.pgxn.org/src/pg_vault_tde/pg_vault_tde-1.7.0/README.md)
- [pg_vault_tde v1.7.0 发行版](https://github.com/labmiriade/pg_vault_tde/releases/tag/v1.7.0)
- [pg_vault_tde 1.7 控制文件](https://api.pgxn.org/src/pg_vault_tde/pg_vault_tde-1.7.0/pg_vault_tde.control)
- [pg_vault_tde 运维文档](https://api.pgxn.org/src/pg_vault_tde/pg_vault_tde-1.7.0/doc/pg_vault_tde.md)

`pg_vault_tde` 通过 `encrypted_heap` 表访问方法，为 PostgreSQL 17 和 18 提供透明的元组加密。它会在存储前使用 AES-256-GCM 加密用户列数据，并通过 HashiCorp Vault/OpenBao、本地 PKCS#12 钱包，或 v1.7 新增的 PKCS#11 HSM 管理每个关系的数据加密密钥。MVCC 元组头仍为明文。

### 配置与安装

```conf
shared_preload_libraries = 'pg_vault_tde'
pg_vault_tde.kms_provider = 'vault'
pg_vault_tde.vault_url = 'https://vault.example.com:8200'
pg_vault_tde.vault_transit_mount = 'transit'
pg_vault_tde.vault_key_name = 'pg-tde-dek'
pg_vault_tde.vault_ca_cert = '/etc/ssl/vault/ca.pem'
```

按照文档通过令牌、AppRole 或 Kubernetes 设置配置 Vault 身份验证，不要将机密写入 PostgreSQL 配置。重启 PostgreSQL，然后创建扩展：

```sql
CREATE EXTENSION pg_vault_tde;
SELECT * FROM pg_vault_tde_health_check();
```

`kms_provider` 没有可用的默认值，必须显式设置。除了 PostgreSQL 服务器文件，扩展还要求 OpenSSL 3 和 libcurl。

### 创建加密表

```sql
CREATE TABLE customer_secrets (
  id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
  email text,
  ssn text
) USING encrypted_heap;
```

加密以表为单位：普通 `heap` 表不受影响。元组值、TOAST 数据和 WAL 表示会被加密；MVCC 所需的元组头仍然可见。

### 索引

使用 `tde_btree` 进行等值查找，同时避免存储明文键：

```sql
CREATE UNIQUE INDEX customer_secrets_id_tde_idx
ON customer_secrets USING tde_btree (id);

CREATE INDEX customer_secrets_email_tde_idx
ON customer_secrets USING tde_btree (email);
```

`tde_btree` 使用确定性的 AES-256-SIV，支持等值比较，但不支持范围排序或仅索引扫描。默认情况下，`encrypted_heap` 表会拒绝其他访问方法，因为它们会写入明文索引键。`PRIMARY KEY` 和 `UNIQUE` 表约束仍会创建原生 btree 索引并产生警告；定义它们前请判断这种暴露是否可以接受。

### 完整性与轮换

```sql
SELECT * FROM pg_vault_tde_verify_integrity('customer_secrets');
SELECT * FROM pg_vault_tde_encrypted_size('customer_secrets');

SELECT pg_vault_tde_rotate_online('customer_secrets', 1000);
SELECT * FROM pg_vault_tde_get_rotation_status('customer_secrets');

SELECT pg_vault_tde_rotate_kek();
```

在线 DEK 轮换会分批重新加密表，并重建其 `tde_btree` 索引。KEK 轮换只会重新包装每张表的 DEK，不会重写元组。请限制这些操作的权限、监控完成状态，并避免并发恢复密钥目录。

### 提供方与备份边界

- 本地钱包默认位于 `PGDATA` 之外；应单独复制并保护，因为普通 `pg_basebackup` 不会包含它。
- 版本 1.7 新增 `pkcs11` 提供方和 `pg_vault_tde_pkcs11_keygen()`。独立工具 `pg_dump_tde` 与 `pg_restore_tde` 在此版本中不支持 PKCS#11。
- 普通 `pg_dump` 和 `COPY ... TO` 会读取解密后的行，因此会在没有警告的情况下生成明文。在支持的场景中，应使用随附的加密逻辑备份工具。
- 物理备份包含加密的关系字节和已包装的 DEK，但不包含 KEK。应预先配置 Vault/HSM 访问权限，或单独复制本地钱包。密钥封装函数与 `pg_basebackup_tde` 包装器可为物理备份附带一份能检出篡改的 DEK 包。

### 关键注意事项

- 当 `encrypted_heap` 表中存在以另一种设置写入的行时，绝不要切换 `pg_vault_tde.enabled`。扩展不会重写已有行，混合格式可能被静默误判为损坏。
- 普通索引、统计信息、日志、查询结果、客户端流量、临时工作数据和备份都可能在加密堆之外暴露明文。TDE 只是存储层控制，并非端到端加密。
- `tde_btree` 禁用范围语义，并且当前设计下加密表会禁用 HOT 更新；请对更新密集型工作负载和索引维护进行基准测试。
- 应对 KMS 凭据、钱包口令、HSM PIN、KEK、密封包和恢复流程实施相互独立的访问控制。缺少匹配密钥路径的备份将无法恢复。
- 软件包发行版本 1.7.0 安装的 SQL 扩展版本为 `1.7`；该扩展不可重定位，需要预加载和重启，并且仅支持 PostgreSQL 17-18。
