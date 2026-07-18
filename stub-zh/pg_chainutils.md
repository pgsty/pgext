## 用法

来源：

- [上游用法文档](https://github.com/usagi-coffee/pg_chainutils/blob/98f534c3912aad5c0c23276baf50045325be7dfb/README.md)
- [版本与依赖清单](https://github.com/usagi-coffee/pg_chainutils/blob/98f534c3912aad5c0c23276baf50045325be7dfb/Cargo.toml)
- [扩展控制文件](https://github.com/usagi-coffee/pg_chainutils/blob/98f534c3912aad5c0c23276baf50045325be7dfb/pg_chainutils.control)

`pg_chainutils` 是在 PostgreSQL 内解析和转换区块链值的 pgrx 扩展。它提供以太坊风格的 `H160`、`H256`、`U256` 类型，以及处理事件 topic 与载荷的辅助函数，覆盖 ERC-20、ERC-721、Uniswap、Sushiswap 等格式。

```sql
CREATE EXTENSION pg_chainutils;
SELECT H256.keccak256('Sync(uint112,uint112)');
SELECT H160.from_H256('0x0000000000000000000000001111111111111111111111111111111111111111');
```

版本 `0.2.4` 使用 pgrx 构建，已复核清单默认面向 PostgreSQL 17。输入辅助函数要求准确的 ABI 编码；入库前必须校验异常长度、前缀、topic 顺序、符号、代币精度和链特定约定。应把解码值视为不可信数据，关键财务结果还应与独立解码器交叉核对。
