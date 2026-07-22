## 用法

来源：

- [固定提交的上游 README](https://github.com/KaiserKarel/pg_bech32/blob/a4eb4aaae249540b9839cd8735cb450bad1ee8eb/README.md)
- [固定提交的 Rust SQL 实现](https://github.com/KaiserKarel/pg_bech32/blob/a4eb4aaae249540b9839cd8735cb450bad1ee8eb/src/lib.rs)
- [固定提交的 Rust 包清单](https://github.com/KaiserKarel/pg_bech32/blob/a4eb4aaae249540b9839cd8735cb450bad1ee8eb/Cargo.toml)
- [固定提交的扩展 control 模板](https://github.com/KaiserKarel/pg_bech32/blob/a4eb4aaae249540b9839cd8735cb450bad1ee8eb/pg_bech32.control)

`pg_bech32` `0.0.0` 为 PostgreSQL 增加 Bech32、Bech32m 和无校验和 Bech32 编码，以及带校验 Bech32 字符串的解码。它以 `bytea` 接收二进制载荷，并把解码数据与人类可读前缀一同返回。

### 核心流程

```sql
CREATE EXTENSION pg_bech32;

SELECT bech32_encode(
  'union',
  decode('644a2606654a7c0e70bf343ae6b828d3fe448447', 'hex'),
  'bech32'
);

SELECT (bech32_decode(
  'union1v39zvpn9ff7quu9lxsawdwpg60lyfpz8pmhfey'
)).*;

SELECT bech32_encode_lower(
  'example',
  convert_to('payload', 'UTF8'),
  'bech32m'
);
```

### 重要对象

- `bech32_encode(hrp, bytea, mode)`：以 `bech32`、`bech32m` 或 `nochecksum` 模式编码。
- `bech32_encode_lower(hrp, bytea, mode)`：使用相同模式生成小写编码字符串。
- `bech32_decode(text)`：验证并解码带校验和的 Bech32 或 Bech32m 字符串。
- 组合类型 `bech32`：包含 `hrp text` 与 `data bytea` 字段；上游未加引号的类型声明会被 PostgreSQL 折叠为小写。

### 运维说明

解码器不会返回使用了哪种校验和变体，文档化的解码路径面向带校验和的 Bech32/Bech32m，而不是 `nochecksum`；协议或版本上下文应保存在值之外。非法 HRP、字符串、校验和或模式名会由 Rust 实现抛出错误。control 模板要求超级用户安装，并且不可重定位。

Cargo feature 矩阵使用 pgrx `0.11.3`，面向 PostgreSQL 11–16；该版本没有确认更高 PostgreSQL 版本。`0.0.0` 版本号和极少的上游文档表明项目接口仍处早期；把编码值用作持久标识符前，应根据权威测试向量验证规范大小写、最大长度、协议特定 witness/version 规则和往返结果。
