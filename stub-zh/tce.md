## 用法

来源：

- [官方项目文档](https://the-cryptic-elephant.readthedocs.io/en/latest/)
- [官方管理说明](https://gitlab.com/dalibo/the_cryptic_elephant/-/blob/latest/docs/admin.md)
- [官方驱动说明](https://gitlab.com/dalibo/the_cryptic_elephant/-/blob/latest/docs/drivers.md)

`tce`（The Cryptic Elephant）是早期阶段的透明列加密扩展。它使用由外部密钥管理系统支持的按角色信封密钥，对 PostgreSQL 内存和存储中的值加密；但上游明确要求不要在生产中使用这一 `0.10.0-dev` 版本线。

### 加载并创建扩展

创建扩展前必须先加载动态库，使其钩子、GUC 与安全标签提供器存在。

```sql
ALTER DATABASE demo SET session_preload_libraries TO 'tce';
-- Reconnect to demo before continuing.
CREATE EXTENSION tce;
```

仅在隔离测试中，允许从文件获取私钥、初始化测试超级用户的信封密钥，并把一个列转换为 `tcetext`。

```sql
CREATE ROLE alice LOGIN SUPERUSER;
ALTER DATABASE demo
  SET tce.allow_unsecured_key_fetch_methods TO true;

\connect demo alice

SECURITY LABEL FOR tce_user_key ON ROLE alice
  IS 'FETCHED FROM FILE /secure/alice-private.pem';

CREATE TABLE secrets (id bigint PRIMARY KEY, secret text);
INSERT INTO secrets VALUES (1, 'classified');

ALTER TABLE secrets
  ALTER COLUMN secret TYPE tcetext
  USING secret::tcetext;

SELECT secret, raw(secret) FROM secrets;
```

文件方法默认有意关闭，文档也限定其只用于测试。密钥加密密钥应使用受支持的 KMS 方案，并保护数据库服务器与超级用户边界。

### 主要对象与设置

- `tcetext` 与 `tceint` 存储加密的文本和整数值；`raw` 用于检查密文。
- `tce.transparent_column_encryption` 启用 DDL 重写，自动转换新建的普通列。
- `tce.data_encryption_method` 选择数据密码；本次复核的源码实现 ChaCha20-Poly1305。
- `tce.encrypted_data_display_mode` 在角色无法解密时选择 `error`、`raw` 或 `empty`。
- `tce.safe_lock` 防止数据加密密钥信封被意外删除。
- `tce.kms.type` 与服务商专用设置用于配置 AWS 或 OpenBao 密钥获取。
- `tce.autocast` 辅助二进制协议驱动，但其 GUC 定义警告它只用于测试。

### 安全与管理边界

该设计假设系统管理员和 PostgreSQL 超级用户可信。安全标签、GUC 与 `tce.envelope` 表可能出现在转储中。逻辑转储和逻辑复制会在值离开发布端前解密；物理备份保留加密存储。恢复目标需要预先加载扩展、提供 TCE 类型，并具有可用的角色密钥。

没有密钥的角色会得到配置的显示行为。密钥变更只对新会话生效，因此需要立即撤销时必须终止旧会话。应显式测试驱动：二进制协议客户端不启用 `tce.autocast` 就无法透明工作，而当前源码又把启用它本身标记为不可用于生产。
