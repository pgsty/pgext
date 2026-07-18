

## 用法

来源：

- [PGXN 0.3.0 发行页](https://pgxn.org/dist/hyperion_vault/0.3.0/)
- [REST API 参考](https://pgxn.org/dist/hyperion_vault/0.3.0/docs/API.html)
- [安全指南](https://pgxn.org/dist/hyperion_vault/0.3.0/docs/SECURITY.html)

`hyperion_vault` 是面向 PostgreSQL 18 的密钥库，只把密文、由 KMS 包装的数据加密密钥和随机数写入 PostgreSQL 与 WAL。Pigsty 打包的 `0.3.0` 版本在上游标记为不稳定，它由 pgrx 扩展与同机部署的 REST API 共同组成。

### 启用扩展

轮换调度器由后台工作进程实现，因此应先把 `hyperion_vault` 加入现有的 `shared_preload_libraries` 列表，重启 PostgreSQL 后再创建扩展：

```sql
SHOW shared_preload_libraries;

CREATE EXTENSION hyperion_vault;
SELECT extversion
FROM pg_extension
WHERE extname = 'hyperion_vault';
```

扩展会创建 `vault` 模式。REST API 使用的专用数据库登录角色可通过扩展提供的辅助函数授权：

```sql
CREATE ROLE vault_service LOGIN PASSWORD 'use-a-managed-secret';
SELECT vault.grant_service_role('vault_service');
```

扩展提供 `hyperion_vault.rotation_enabled`、`hyperion_vault.scan_interval_secs` 与 `hyperion_vault.database` 等 GUC。REST 服务在 `/v1/secrets` 下提供创建、读取、更新、删除、轮换和校验接口。

### 安全边界

生产模式需要 AWS KMS；本地主密钥模式仅供开发与测试。密钥读取受 `VAULT_ALLOWED_IPS` IPv4/CIDR 白名单控制，空白名单会默认拒绝全部读取；管理操作需要 Bearer Token。部署这个预览版本前，应完整阅读上游安全与威胁模型文档，为 API 配置 TLS 终止，并在远程 PostgreSQL 连接中启用证书校验。
