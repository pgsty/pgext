## 用法

来源：

- [官方 README](https://github.com/blkchain/pg_blkchain/blob/d90127e749a72c918fe3499cfc4d26a731ba30d5/README.md)
- [官方扩展 SQL](https://github.com/blkchain/pg_blkchain/blob/d90127e749a72c918fe3499cfc4d26a731ba30d5/pg_blkchain--0.0.1.sql)
- [官方 C 实现](https://github.com/blkchain/pg_blkchain/blob/d90127e749a72c918fe3499cfc4d26a731ba30d5/pg_blkchain.c)

`pg_blkchain` 是一个实验性 C 扩展，用于在 PostgreSQL 中解码和构造旧式 Bitcoin 区块、交易、输入、输出与脚本。版本 `0.0.1` 被明确标为开发中，复核的源码停留在 2017 年；不能把其中的解析或签名验证视为当前共识或安全实现。

### 核心流程

安装扩展，将原始序列化交易保存为 `bytea`，然后展开交易输出进行分析：

```sql
CREATE EXTENSION pg_blkchain;

SELECT r.id, (get_vout(r.tx)).*
FROM raw_transactions AS r
WHERE r.id = 37898;
```

脚本可以展开为操作码记录：

```sql
SELECT op_sym, op, encode(data, 'hex') AS data_hex
FROM parse_script(transaction_output_script);
```

第二个示例假定调用方提供的 `transaction_output_script` 是一个 `bytea` 值。

### 重要对象

- `get_block(bytea)` 返回 `CBlock`；`get_tx(bytea)` 返回 `CTx`。
- `get_vin(bytea)` 与 `get_vout(bytea)` 返回 `CTxIn` 和 `CTxOut` 记录集。
- `parse_script(bytea)` 返回 `CScriptOp` 记录，包含符号操作码、数值操作码与压入的数据。
- `verify_sig(txTo bytea, txFrom bytea, int)` 尝试验证一个交易输入签名。
- `build_vin(...)` 与 `build_vout(...)` 是用于序列化输入和输出组成部分的聚合。
- `get_vin_outpt_arr(bytea)`、`get_vin_outpt_jsonb(bytea)` 与 `get_vin_outpt_bytea(bytea)` 以多种表示返回前序输出引用。

### 运维注意事项

上游警告明确要求自行承担使用风险。代码早于多次 Bitcoin 与 PostgreSQL 重大变化，会在数据库后端中解析不受信任的二进制数据，也没有仍在维护的兼容性或共识验证保证。畸形输入或 C 缺陷可能终止后端。README 中的签名示例还仅在双 SHA256 查找查询中需要 `pgcrypto`，它并未声明为扩展依赖。验证工作应使用 PostgreSQL 外部仍在维护的 Bitcoin 库，并将此扩展限制在隔离的历史数据实验中。
