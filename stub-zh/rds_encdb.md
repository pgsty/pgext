## 用法

来源：

- [阿里云 rds_encdb 文档](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-rds-encdb-extension-to-encrypt-sensitive-columns)
- [阿里云 RDS PostgreSQL 概述](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/what-is-apsaradb-rds-for-postgresql)

`rds_encdb` 是阿里云 RDS PostgreSQL 在查询结果离开数据库时加密指定列的功能。它是服务商管理的结果集脱敏，而不是可移植的静态列加密；阿里云也没有为它发布社区源码包或控制文件。

### 启用与核心流程

官方流程仅适用于小版本 `20250228` 或更高的 RDS PostgreSQL 16 实例。通常不支持直接自行安装：必须取得服务商授权，启用实例参数，并使用 RDS 指定的高权限账户。

```sql
CREATE EXTENSION rds_encdb;

CREATE TABLE test(a text, b text, c text);
INSERT INTO test VALUES ('foo', 'bar', 'hello world');

INSERT INTO rds_encdb.encryption_rule
VALUES
  (9, 'rule1', 'test', '1'),
  (10, 'rule1', 'test', '2');

SELECT * FROM rds_encdb.rules;
```

安装前必须把实例参数 `rds_encdb.enable_encryption` 设为 `on`。写入 `rds_encdb.encryption_rule` 的新规则对现有会话生效。文档给出的默认算法为 AES-256-GCM。

### 规则与角色授权

`rds_encdb.encryption_rule` 通过 `attrelid` 与 `attnum` 标识受保护的表和列，并以 `rule_name` 分组。`rds_encdb.rules` 视图按规则与表汇总这些条目。

账户默认使用 `RESTRICTED ACCESS`，受保护的结果列返回密文；`FULL ACCESS` 返回明文。高权限账户通过 `rds_encdb.setup_encryption_role(account, permission, expiration)` 与 `rds_encdb.remove_encryption_role(account)` 管理该状态，记录可在 `rds_encdb.encryption_role_auth` 中查看。可选的 EncJDBC 客户端建立密钥材料后可以解密受限结果。

### 限制与安全边界

服务商文档不支持函数返回的结果集、游标和 prepare/execute 等非 `SELECT` 操作，以及使用 CTE 或 `UNION` 的查询。预备语句是许多客户端的默认行为，因此必须测试实际 ORM 与连接池行为。

加密根据账户策略作用于返回值；它不能阻止权限足够高的账户读取明文，也不能替代存储加密、传输加密、审计或最小权限。可用性、升级、密钥行为、授权与支持均处于阿里云 AliPG 服务边界内。
