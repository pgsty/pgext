## 用法

来源：

- [pg_pii_vault.control](https://github.com/g0ddest/pg_pii_vault/blob/306482c1134317c3e503ba9c9cf8ec1f4b332fe1/pg_pii_vault.control)
- [Cargo.toml](https://github.com/g0ddest/pg_pii_vault/blob/306482c1134317c3e503ba9c9cf8ec1f4b332fe1/Cargo.toml)
- [README.md](https://github.com/g0ddest/pg_pii_vault/blob/306482c1134317c3e503ba9c9cf8ec1f4b332fe1/README.md)
- [UPGRADING.md](https://github.com/g0ddest/pg_pii_vault/blob/306482c1134317c3e503ba9c9cf8ec1f4b332fe1/UPGRADING.md)
- [sql/pg_pii_vault--0.0.0--0.1.0.sql](https://github.com/g0ddest/pg_pii_vault/blob/306482c1134317c3e503ba9c9cf8ec1f4b332fe1/sql/pg_pii_vault--0.0.0--0.1.0.sql)
- [sql/pg_pii_vault--0.1.0--0.1.1.sql](https://github.com/g0ddest/pg_pii_vault/blob/306482c1134317c3e503ba9c9cf8ec1f4b332fe1/sql/pg_pii_vault--0.1.0--0.1.1.sql)

`pg_pii_vault` 0.1.1 提供 `piitext` 类型，使用 HashiCorp Vault Transit 中按数据主体划分的密钥显式加密列值。只有 SELECT 权限的角色看到密文，解密需要单独的函数授权。支持 PostgreSQL 14–18，安装需要超级用户。

### 核心用法

```ini
shared_preload_libraries = 'pg_pii_vault'
pii_vault.url = 'https://vault.example.internal:8200'
pii_vault.token_file = '/run/vault-agent/pg_pii_vault.token'
pii_vault.mount = 'transit'
```

```sql
CREATE EXTENSION pg_pii_vault;
SELECT check_name, ok, required FROM piitext_vault_check();
SELECT piitext_encrypt('test value', int4send(123))::text;
```

### 运行边界

先配置最小权限的 Vault 令牌文件与 TLS，预加载 `pg_pii_vault` 并重启。预加载提供集群级密钥缓存失效，并保护令牌配置。示例假定 Transit 挂载点已初始化，使用测试主体密钥。全部 `pii_vault.*` 设置均仅限超级用户。默认 export 模式将可导出的密钥带入后端内存，transit 模式则使用不可导出密钥，将加解密交给 Vault。`piitext_shred` 不可逆地删除主体密钥；`piitext_cache_invalidate`、`piitext_stats`、`piitext_vault_check`、`piitext_key_version` 和 `piitext_reencrypt` 辅助运维。网络或权限故障产生明确错误，只有确认密钥／版本不存在才返回删除标记。设置超时，并同时保护 Vault 备份和加密数据库数据。

0.1 改变旧应用约定：移除隐式明文写入和隐式解密，撤销 PUBLIC 的加密／解密／管理函数执行权，而已持久化的解密表达式仍可能保存明文。从 0.0.0 升级前，按完整上游流程盘点并备份，调整查询，将秘密移出角色／数据库设置，配置预加载并重启，然后在各受影响数据库运行 ALTER EXTENSION。替换可能暴露的旧令牌，显式加密既有暂存明文，并按指引移除或重新设计含明文的索引、生成列、物化视图与统计对象。历史 WAL、转储、副本和 Vault 快照需要分别处理保留策略。暂存值迁移完成后，可用 `pii_vault.allow_staging = off` 拒绝它们。从 0.1.0 升至 0.1.1 需要匹配文件、重启和 SQL 升级，但不重写值。没有自动降级路径，旧库不能读取新存储格式。
