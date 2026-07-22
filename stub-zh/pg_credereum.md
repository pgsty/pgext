## 用法

来源：

- [官方 README](https://github.com/postgrespro/pg_credereum/blob/master/README.md)
- [官方 control 文件](https://github.com/postgrespro/pg_credereum/blob/master/pg_credereum.control)

`pg_credereum` 是面向 PostgreSQL 10 的可加密验证行变更历史原型。客户端签署事务变更集，后台工作进程把变更打包成 Merkle 树区块，还可选择把区块哈希发布到 Ethereum 智能合约。上游明确说明它尚未达到生产可用状态。

### 核心流程

预加载该库，选择唯一受管理数据库及扩展模式，然后重启服务器：

```conf
shared_preload_libraries = 'pg_credereum'
pg_credereum.database = 'appdb'
pg_credereum.schema = 'public'
pg_credereum.block_period = 1000
```

创建扩展，并把累加器触发器挂到主键恰好为单个名为 `id` 的 `bigint` 字段的表上：

```sql
CREATE EXTENSION pg_credereum;

CREATE TABLE t (id bigint PRIMARY KEY, value integer NOT NULL);
CREATE TRIGGER t_after
AFTER INSERT OR UPDATE OR DELETE ON t
FOR EACH ROW EXECUTE PROCEDURE credereum_acc_trigger();
REVOKE TRUNCATE ON t FROM PUBLIC;
```

在一个事务中执行 DML，调用 `credereum_get_changeset()` 并验证返回的证明，然后在提交前用 `credereum_sign_transaction(pubkey text, sign bytea)` 签名。

### API 与配置

- `credereum_get_changeset()`：返回当前事务变更的 Merkle 证明行。
- `credereum_sign_transaction(text, bytea)`：记录客户端公钥与签名。
- `credereum_merkle_proof(varbit[])`：返回所选 Merkle 键的历史与证明数据。
- `credereum_tx_log` 与 `credereum_block`：事务和区块历史。
- `pg_credereum.block_retry_period`：重试间隔，默认 `5000` 毫秒。
- `pg_credereum.eth_end_point`、`pg_credereum.eth_source_addr`、`pg_credereum.eth_contract_addr`：可选 Ethereum 发布设置。

### 注意事项

扩展需要 PostgreSQL 10，以及文档所列的 Ethereum、CURL 和 Jansson 构建依赖。它只管理 `pg_credereum.database` 指定的数据库。`TRUNCATE` 不会被跟踪，同一行在一个区块内不能修改两次；第二次修改会触发唯一约束错误。缩短 `block_period` 可减小冲突时间窗口。

如果向 Ethereum 发布区块哈希失败，工作进程会跳过该哈希并尝试下一个区块，因此可信存储中可能合理地存在缺口。应把它视为研究原型，独立验证每条证明和签名路径，不能仅凭安装此扩展就认定数据未被篡改。
