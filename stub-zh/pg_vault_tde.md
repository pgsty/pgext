## 用法

来源：

- [pg_vault_tde.control](https://github.com/labmiriade/pg_vault_tde/blob/2520a0d80fc3c4d905f57e47e2d871ed7e6b0e7d/pg_vault_tde.control)
- [README.md](https://github.com/labmiriade/pg_vault_tde/blob/2520a0d80fc3c4d905f57e47e2d871ed7e6b0e7d/README.md)
- [doc/pg_vault_tde.md](https://github.com/labmiriade/pg_vault_tde/blob/2520a0d80fc3c4d905f57e47e2d871ed7e6b0e7d/doc/pg_vault_tde.md)

`pg_vault_tde` 发行版 1.7.2 通过 `encrypted_heap` 使用 AES-256-GCM 加密表值，支持 Vault/OpenBao、本地 PKCS#12 钱包或 PKCS#11。SQL/control 版本仍为 1.7。需要 PostgreSQL 17–18、OpenSSL 3 与 libcurl。先配置密钥提供方和认证，预加载库并重启，再由超级用户创建扩展。

### 核心用法

```ini
shared_preload_libraries = 'pg_vault_tde'
```

```sql
CREATE EXTENSION pg_vault_tde;
SELECT * FROM pg_vault_tde_health_check();
CREATE TABLE customer_secrets (id bigint, secret text) USING encrypted_heap;
CREATE INDEX customer_secrets_id_idx ON customer_secrets USING tde_btree (id);
```

### 运行边界

示例假定密钥提供方已配置。`pg_vault_tde_health_check`、`pg_vault_tde_verify_integrity`、`pg_vault_tde_get_rotation_status` 和加密大小辅助函数提供运行状态。`tde_btree` 支持加密等值查询，不支持排序／范围或仅索引扫描。不支持 numeric 和非确定性排序规则的加密索引；应审计旧有唯一／排除约束，它们的检查可能无效。普通索引可能暴露键值，pg_dump/COPY 会输出解密数据。保护外部钱包／KMS，并在支持时使用匹配的加密备份工具。元组头、统计、日志和查询结果不在其保护范围内。

替换旧库或重启之前，必须按上游恢复清单检查：旧版轮换后的表可能依赖仅存于内存的密钥；并发轮换可能要求趁数据仍可读时导出，轮换过的行外值需要指定的重写步骤。1.7.0 的 TOAST 升级还另有导出要求。检查轮换／部分索引及钱包创建者。升级到 1.7.2 后，所有既有加密表需执行 VACUUM FULL 迁移元组布局，应预留维护窗口和额外磁盘空间。启用 `pg_vault_tde.toast_custom_rmgr` 时，WAL ID 从 128 改为 161，要求干净停机、主备协调升级、复制槽追平和新的基础备份，不能跨这两种 WAL 格式滚动升级。

密钥管理操作必须逐一执行。表轮换保留读取，但在单个事务期间阻塞写入；逻辑槽必须在重启或再次轮换前完成解码。1.7.2 收紧调用者权限，但上游仍记录 SECURITY DEFINER 包装器的限制，应遵循列出的授权／撤权步骤。不要在已有数据的加密表上切换 `pg_vault_tde.enabled`。本次目录更新不执行实际升级，也不改变 Pigsty 软件包基线。
