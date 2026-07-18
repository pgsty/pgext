## 用法

来源：

- [项目 README](https://github.com/txpipe/mumak/blob/4da43f3cb21d986280faf2ad373db954418dc888/README.md)
- [Rust 实现中的函数列表](https://github.com/txpipe/mumak/blob/4da43f3cb21d986280faf2ad373db954418dc888/extension/src/lib.rs)
- [函数文档](https://github.com/txpipe/mumak/blob/4da43f3cb21d986280faf2ad373db954418dc888/docs/FUNCTIONS.md)
- [扩展 control 文件](https://github.com/txpipe/mumak/blob/4da43f3cb21d986280faf2ad373db954418dc888/extension/mumak.control)
- [扩展 Cargo 清单](https://github.com/txpipe/mumak/blob/4da43f3cb21d986280faf2ad373db954418dc888/extension/Cargo.toml)

`mumak` 0.0.6 是一个用于在 PostgreSQL 内检查 Cardano CBOR 的 pgrx 扩展。索引器可以把区块、交易或 UTxO 保存在 `bytea` 中，再直接通过 SQL 提取哈希、地址、手续费、lovelace 与资产数量、元数据、Plutus 数据、epoch、slot 等字段，无须预先把所有结构完全规范化。

### 验证安装并查询 CBOR

```sql
CREATE EXTENSION mumak;
SELECT hello_extension();

SELECT tx_hash(tx_cbor), tx_fee(tx_cbor), tx_outputs_json(tx_cbor)
FROM cardano_transactions;
```

已实现的函数族包括 `block_tx_count(bytea)`、`block_number(bytea)`、`tx_hash(bytea)`、`tx_inputs(bytea)`、`tx_outputs(bytea)`、`tx_is_valid(bytea)`、`utxo_lovelace(integer,bytea)`，以及地址和 Bech32 转换函数。应以 Rust 实现为 API 准绳：顶层 README 中含有规划中的名称，并非都与 0.0.6 实际交付函数完全一致。

### 输入与部署注意事项

许多解码失败不会抛出异常，而是返回空值、false、零或 -1 等哨兵结果。在把结果当作区块链事实之前，必须先校验 CBOR 及其 era/network 上下文。control 文件要求由超级用户安装。已审查 Cargo feature 覆盖 PostgreSQL 11 至 16；项目只发布源码和有限的 Linux 构件说明，因此应针对实际服务器版本和架构完成构建与回归测试。
