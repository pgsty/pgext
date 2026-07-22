## 用法

来源：

- [官方 README](https://github.com/g0ddest/pg_pii_vault/blob/34256629bf8f1af5c4a2437452387ed9d156ee3e/README.md)
- [官方 SQL API 实现](https://github.com/g0ddest/pg_pii_vault/blob/34256629bf8f1af5c4a2437452387ed9d156ee3e/src/lib.rs)
- [官方 Vault 集成](https://github.com/g0ddest/pg_pii_vault/blob/34256629bf8f1af5c4a2437452387ed9d156ee3e/src/vault.rs)

`pg_pii_vault` 是早期阶段的 `piitext` 类型，使用 AES-256-GCM 和逐记录密钥标识加密文本，并从 HashiCorp Vault Transit 获取可导出的加密密钥。上游明确说明它尚未准备好用于生产环境。

### 核心流程

该扩展不可信，只能由超级用户安装。配置 Vault Transit 端点和窄权限令牌，然后使用稳定的记录密钥标识显式加密：

```sql
CREATE EXTENSION pg_pii_vault;

SET pii_vault.url = 'https://vault.internal';
SET pii_vault.token = '<vault-token>';
SET pii_vault.mount = 'transit';

CREATE TABLE private_users (
  id integer PRIMARY KEY,
  secret piitext NOT NULL
);

INSERT INTO private_users(id, secret)
VALUES (
  123,
  piitext_encrypt('sensitive value', decode('0000007b', 'hex'))
);

SELECT id, piitext_out_text(secret)
FROM private_users;
```

`piitext_debug` 会暴露已存储结构，`piitext_raw` 会暴露原始 CBOR 字节，二者的执行权限都应严格限制。`piitext_encrypt_piitext` 会解密已有值，并使用新的密钥标识重新加密。

### 关键安全边界

- 从 `text` 到 `piitext` 的隐式转换会调用 `piitext_in_text` 并保存明文暂存值；加密不会自动发生。只有通过 `piitext_encrypt` 或 `piitext_encrypt_piitext` 处理的行才会密封。
- 代码要求 Vault Transit 把密钥创建为 `exportable`，导出原始加密密钥，并在 PostgreSQL 后端执行 AES-GCM。因此数据库进程及其内存进入密钥信任边界；这不是通常不导出密钥的 Transit 加解密模式。
- `pii_vault.url`、`pii_vault.token`、`pii_vault.mount` 和 `pii_vault.cache_ttl_sec` 都是 `USERSET`。如果不严格控制函数与 GUC 使用，调用者可以选择端点和凭据，并触发阻塞式服务器 HTTP 请求。默认应撤销公共执行权限和出站网络访问。
- 密钥缓存位于各后端进程内存，并非 PostgreSQL 共享内存。删除 Vault 密钥后，要等后端缓存过期才会生效；而服务中断、认证失败、密钥已删、密文损坏或认证标签失败都会让 `piitext_out_text` 返回同一个 `****` 标记。
- `piitext_encrypt` 与 `piitext_out_text` 被声明为 `IMMUTABLE`，但它们依赖随机 IV、GUC、网络或 Vault 状态以及缓存状态。不能安全依赖规划器常量折叠、表达式索引或生成列相关假设。

### 运维限制

- Vault 调用在 SQL 后端中同步执行。应从外部设置网络与语句超时、隔离故障，并测量连接池放大和缓存行为。
- AAD 根据类型名和密钥标识派生，并没有独立验证表主键。应用必须保持密钥标识与预期记录绑定，并阻止相同标识内的密文交换。
- 自动触发器、密钥轮换、后台缓存清理和生产集成测试仍是路线图项目。mock 模式使用固定全零密钥，只能用于测试。
- “密码学粉碎”本身不能证明符合法规的删除。删除设计必须覆盖缓存、Vault 版本与备份、数据库备份、副本、WAL、导出、日志和明文暂存行。
