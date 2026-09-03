## 用法

来源：

- [官方文档](https://github.com/codybrom/pg_fedi/blob/2616f1f1e9566934bf53261d16fcf4c051d4ce0d/README.md)
- [扩展控制文件](https://github.com/codybrom/pg_fedi/blob/2616f1f1e9566934bf53261d16fcf4c051d4ce0d/pg_fedi.control)
- [官方仓库](https://github.com/codybrom/pg_fedi)

`pg_fedi` 面向 PostgreSQL 与 pg_tle 环境的纯 SQL ActivityPub 联邦功能。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_fedi`：

```sql
CREATE EXTENSION pg_fedi CASCADE;
```

经审查的控制文件或官方流程要求 `pgcrypto`。只有服务器已安装这些扩展文件时，`CASCADE` 才会成功。

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- Create a local actor (app layer generates RSA keypair, passes it in)
SELECT ap_create_local_actor('alice', 'Alice', 'Hello!', public_pem, private_pem);

-- Post a note (queues delivery to followers)
SELECT ap_create_note('alice', '<p>Hello, fediverse!</p>');

-- Process an inbound activity
SELECT ap_process_inbox_activity('{"type":"Follow", ...}'::jsonb);

-- WebFinger lookup
SELECT ap_webfinger('acct:alice@myinstance.social');

-- Actor profile as ActivityStreams JSON-LD
SELECT ap_serialize_actor('alice');

-- Timelines
SELECT * FROM ap_public_timeline;
SELECT * FROM ap_home_timeline('alice');
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `ap_set_setting` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ap_process_inbox_activity` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ap_serialize_actor` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ap_webfinger` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ap_create_local_actor` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ap_create_note` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ap_deliveries` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `ap_delivery_failure` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 应将扩展升级视为数据库变更：先审查上游升级路径、权限、锁以及备份恢复行为。
