## 用法

来源：

- [官方 README](https://github.com/KaiserKarel/postgreth/blob/c86650deaaf1f86f7ae9931c339e24e44143a0b0/README.md)
- [官方 Rust 实现](https://github.com/KaiserKarel/postgreth/blob/c86650deaaf1f86f7ae9931c339e24e44143a0b0/src/lib.rs)
- [官方包与 PostgreSQL 特性矩阵](https://github.com/KaiserKarel/postgreth/blob/c86650deaaf1f86f7ae9931c339e24e44143a0b0/Cargo.toml)
- [官方 control 文件](https://github.com/KaiserKarel/postgreth/blob/c86650deaaf1f86f7ae9931c339e24e44143a0b0/postgreth.control)

`postgreth` 是实验性的 pgrx 扩展，用于解码 Ethereum ABI 值与日志，以及处理 Ethereum bloom filter。它不会连接 Ethereum 节点、验证链数据或同步表；调用者必须提供已经获取的字节、ABI JSON 与日志 JSON。

### 核心流程

解码单个 ABI 标量时，提供 Solidity 类型及其字节：

```sql
CREATE EXTENSION postgreth;

SELECT item_to_jsonb(
  'uint256',
  decode(repeat('00', 31) || '2a', 'hex'),
  false
);
```

`item_to_jsonb(text,bytea,boolean)` 返回解码后的 `jsonb`。最后一个标志会为需要它的 tuple 类值添加 ABI 偏移字；它不是修复非法输入的通用选项。

### 类型与函数

- `Bloom`、`Address` 与 `B256` 分别表示 256 字节 bloom filter、20 字节地址和 32 字节值。
- `covers(Bloom,Bloom)`、`contains_input(Bloom,bytea)` 与 `contains_input_hashed(Bloom,bytea)` 检查 bloom 的覆盖/成员关系。
- `m3_2048(Bloom,bytea)` 与 `m3_2048_hashed(Bloom,bytea)` 把原始或预哈希输入加入 bloom filter。
- `contains_raw_log(Bloom,Address,B256[])` 检查地址与 topics 是否匹配 bloom filter。
- `log_to_jsonb(text,text)` 使用事件 ABI 解码 receipt-log JSON，并返回带键的 `jsonb`。

bloom 成员判断具有概率性：阳性结果可能是假阳性，不能代替精确日志匹配或链验证。

### 安全与兼容性

多个解析路径使用 Rust `expect`/`unwrap`。非法 ABI JSON、错误日志 JSON、无法解码的字节、非法 Solidity 类型，或长度不恰好为 32 字节的哈希，都可能通过扩展边界报错或 panic。调用前必须校验不可信数据，并在目标后端上模糊测试失败行为。

control 文件要求超级用户安装，并将扩展标记为不可重定位。核验的 crate 版本是 `0.0.1`，使用 pgrx `0.11.3`，默认 PostgreSQL 13，并定义 PostgreSQL 11–16 的构建 feature。这些 feature 声明并不等同于发行认证；应针对确切服务端构建并测试所选 feature。

上游明确说明该项目不面向公众使用，并缺少广泛测试、CI/CD 与发行流程。应把它用于受控实验或经过内部审计的分支，而不是生产级区块链信任边界。
