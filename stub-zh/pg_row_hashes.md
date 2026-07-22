## 用法

来源：

- [官方 README](https://github.com/ibotty/pg-row-hashes/blob/main/README.md)
- [官方 0.3.2 版实现](https://github.com/ibotty/pg-row-hashes/blob/main/src/lib.rs)
- [官方 UUID XOR 聚合](https://github.com/ibotty/pg-row-hashes/blob/main/src/xor_agg.rs)
- [官方 control 文件](https://github.com/ibotty/pg-row-hashes/blob/main/pg_row_hashes.control)

`pg_row_hashes` 为有序标识符和类似映射的行校验提供确定性、非密码学指纹。版本 `0.3.2` 提供以 `uuid` 返回的 128 位 FarmHash、以 `bigint` 返回的 64 位 SeaHash、下划线拼接的 MD5 标识辅助函数，以及 UUID XOR 聚合。

### 核心流程

参数顺序属于标识符时使用 `id_*` 函数；键值对校验使用 `checksum_*`：

```sql
CREATE EXTENSION pg_row_hashes;

SELECT id_farmhash('tenant-7', 'order-42');
SELECT id_seahash('tenant-7', 'order-42');

SELECT checksum_farmhash(
  'status', 'paid',
  'amount', '19.95'
);

SELECT checksum_farmhash_extendable(
  'status', 'paid',
  'optional_note', NULL
);
```

Extendable 变体会跳过值为 `NULL` 的键值对，因此给模式新增可空字段时，旧行指纹不会改变，直到该字段有值。非 extendable 变体会保留该键的存在信息。

### API

- `id_farmhash(VARIADIC text[])` / `id_seahash(VARIADIC text[])`：有序多段指纹。
- `id_farmhash_bytea(bytea)` / `id_seahash_bytea(bytea)`：原始字节指纹。
- `id_underscore_md5(VARIADIC text[])`：对以 `_` 拼接的参数计算 MD5，并返回 `uuid`。
- `checksum_farmhash` / `checksum_seahash`：顺序规范化、保留空值键的键值校验。
- `checksum_farmhash_extendable` / `checksum_seahash_extendable`：跳过空值键值对。
- `bit_xor(uuid)`：parallel-safe 的 UUID XOR 聚合。

### 注意事项

这里的 FarmHash、SeaHash、MD5 与 XOR 是指纹工具，不是认证、防篡改或抗碰撞密码学。应把源业务键与指纹一同保存，并处理可能的碰撞。XOR 还会抵消成对出现的相同 UUID，无法单独证明多重集合内容。

Checksum 参数是文本键值对，数量必须为偶数。调用方必须规定稳定的类型转文本、区域、时区、数值与 JSON 规范化规则；格式变化就会改变指纹。版本 `0.3.2` 为 PostgreSQL 17/18 构建，应确认软件包 feature 与目标服务器一致。
