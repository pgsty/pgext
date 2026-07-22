## 用法

来源：

- [官方 README](https://github.com/paroga/pg-cbor/blob/23afc682e6168e59c4d27d64fb10e8c5e624722a/README.md)
- [0.1 版本 SQL 对象](https://github.com/paroga/pg-cbor/blob/23afc682e6168e59c4d27d64fb10e8c5e624722a/sql/cbor.sql)
- [CBOR 输入输出实现](https://github.com/paroga/pg-cbor/blob/23afc682e6168e59c4d27d64fb10e8c5e624722a/src/cbor_io.c)

`cbor` 增加了基于 RFC 7049 的简明二进制对象表示类型。它适合数据库需要存储、交换 CBOR 字节，同时保留 SQL 文本表示的场景；它不能替代 `jsonb` 的查询与索引能力。

### 编码与解码

```sql
CREATE EXTENSION cbor;

SELECT '{"sensor":"alpha","values":[1,2,3]}'::cbor;

WITH encoded AS (
  SELECT cbor_encode('{"ok":true,"count":3}'::cbor) AS payload
)
SELECT cbor_decode(payload)
FROM encoded;
```

`cbor_encode(cbor)` 以 `bytea` 返回二进制表示；`cbor_decode(bytea)` 解析二进制 CBOR。该类型的文本输入输出覆盖整数、字节串、文本串、数组、映射、标签、浮点值与简单值。

### 操作符与索引

扩展提供 `=`、`<>`、`<`、`<=`、`>` 和 `>=`，以及默认 B-tree 与哈希操作符类。它还暴露 `cbor_contains(cbor, cbor)` 与 `cbor_contained(cbor, cbor)` 函数，但 `0.1.0` 版本没有定义包含操作符。

```sql
CREATE TABLE messages (id bigint GENERATED ALWAYS AS IDENTITY, body cbor);
CREATE INDEX messages_body_btree ON messages (body);
```

这个旧版本面向原始 RFC 7049，而不是后来的 RFC 8949。与其他实现交换数据前，应测试标签、浮点值、重复映射键、排序、规范编码与相等语义。如果要求逐字节一致，应保留原始 `bytea`，因为解码再编码未必能保留外部编码器的精确表示。
