## 用法

来源：

- [官方扩展控制文件](https://github.com/unionlabs/pg_ibc/blob/14abdfdfc6e0fb63500147ed16de14195b0610dc/pg_ibc_0_1.control)
- [官方 Rust 包清单](https://github.com/unionlabs/pg_ibc/blob/14abdfdfc6e0fb63500147ed16de14195b0610dc/Cargo.toml)
- [官方 SQL API 实现](https://github.com/unionlabs/pg_ibc/blob/14abdfdfc6e0fb63500147ed16de14195b0610dc/src/lib.rs)

`pg_ibc_0_1` 0.1.1 提供 Union 的 IBC 索引流水线使用的不可变 PostgreSQL 函数。它可把转账包和 UCS03 确认解码成 JSONB、计算数据包哈希、预测多条链上的包装代币地址，并格式化 EIP-55 地址。它不会连接区块链、提交交易或验证共识状态。

### 核心流程

安装带版本后缀的扩展名，然后解码索引器已经采集的字节。对于 Cosmos RPC 数据，输入是 JSON 字节；把扩展格式设为 JSON 时，会尽可能把字符串形式的 memo 转为嵌套 JSON。

```sql
CREATE EXTENSION pg_ibc_0_1;

SELECT decode_transfer_packet_0_1(
    convert_to('{"sender":"union1example","memo":"{\"forward\":{}}"}', 'UTF8'),
    'cosmos',
    false,
    'json'
);

SELECT erc55_to_checksum_0_1(
    decode('52908400098527886e0f7030069857d2e4169ee7', 'hex')
);
```

对于 EVM 转账包，应传入 ABI 编码的数据包字节、`evm`，并为扩展字段选择 `string` 或 `json`。当 `throws = false` 时，普通 EVM 或 Cosmos 解码失败返回 JSON null；当 `throws = true` 时则抛出错误。

### 数据包与确认函数

- `decode_transfer_packet_0_1(bytea, text, boolean, text)` 解码较早的 EVM 或 Cosmos 转账包结构。
- `decode_packet_0_1(bytea, text)` 解码通道版本 `ucs03-zkgm-0` 的数据包，并返回带标签的 JSONB 结果。
- `decode_ack_0_1(bytea, bytea, text)` 解码通道版本 `ucs03-zkgm-0` 的确认。
- `decode_packet_ack_0_1` 和 `decode_packet_ack_0_2` 对数据包字段求哈希并解码可选确认；后者以文本接收高度与时间戳，因此不受 SQL 有符号 64 位整数限制。
- `decode_packet_ack_0_3` 接收数据包数据和外部提供的 32 字节数据包哈希，返回解码数据或结构化的哈希与解码错误详情。

### 地址派生函数

- `erc55_to_checksum_0_1` 要求地址恰好为 20 字节，并返回 EIP-55 字符串。
- `create3_0_1` 预测 EVM CREATE3 包装代币地址。
- `instantiate2_0_1` 使用此版本内嵌的校验和预测 CosmWasm Instantiate2 地址。
- `predict_cosmos_wrapper_0_1`、`predict_osmosis_wrapper_0_1` 和 `predict_aptos_wrapper_0_1` 根据通道路径与代币字节派生特定链的包装标识符。

### 使用边界

该包实现的是 Union 协议格式的一个快照。数据包解码器目前只识别 `evm`、`cosmos` 和通道版本 `ucs03-zkgm-0`；后续协议版本需要新的扩展或函数版本。不支持的 RPC 类型会进入未实现的 Rust 分支，即使 `throws` 为 false 也会报错。

所有输入都是原始 `bytea` 与标识符；该扩展不会验证数据包、确认、部署者、创建者或通道是否真实存在于链上。应保留原始字节与链元数据，以便升级后重新计算。版本 0.1.1 只声明了 PostgreSQL 14、15 和 16 的构建特性。
