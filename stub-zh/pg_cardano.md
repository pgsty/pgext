## 用法

来源：[README](https://github.com/Fell-x27/pg_cardano/blob/master/README.md), [Cargo.toml version 1.2.0](https://github.com/Fell-x27/pg_cardano/blob/master/Cargo.toml)

`pg_cardano` 是一个 Rust 扩展，为 PostgreSQL 提供面向 Cardano 的二进制与加密工具。上游仓库没有发布 GitHub release 或 tag，但默认分支当前 crate 版本是 `1.2.0`。

```sql
CREATE EXTENSION pg_cardano;
```

### 核心能力

- Base58 编码与解码。
- Bech32 编码与解码。
- CBOR 与 `jsonb` 之间的双向转换，提供简单与扩展两套管线。
- Blake2b 哈希。
- Ed25519 签名与签名验证。
- CIP-105 与 CIP-129 的 dRep ID 辅助函数。
- Shelley 地址构造与解析。
- 资产名称解码与 CIP-88 pool key registration 验证。

### 常见模式

#### Base58 与 Bech32

```sql
SELECT cardano.base58_encode('Cardano'::bytea);
SELECT cardano.base58_decode('3Z6ioYHE3x');
SELECT cardano.bech32_encode('ada', 'is amazing'::bytea);
SELECT cardano.bech32_decode_prefix('ada1d9ejqctdv9axjmn8dypl4d');
SELECT cardano.bech32_decode_data('ada1d9ejqctdv9axjmn8dypl4d');
```

#### CBOR 转换

```sql
SELECT cardano.cbor_decode_jsonb('\xa3636164616b...'::bytea);
SELECT cardano.cbor_encode_jsonb('{"ada": "is amazing!", "bytes": "\\xdeadbeef"}'::jsonb);
SELECT cardano.cbor_decode_jsonb_ext('\xa3636164616b...'::bytea);
SELECT cardano.cbor_encode_jsonb_ext('{"type":"map","value":[...]}'::jsonb);
```

#### 哈希与签名

```sql
SELECT cardano.blake2b_hash('Cardano is amazing!'::bytea, 32);
SELECT cardano.ed25519_verify_signature(
  '\x432753BD...'::bytea,
  'Cardano is amazing!'::bytea,
  '\x74265f96...'::bytea
);
```

#### 地址与治理辅助函数

```sql
SELECT cardano.tools_drep_id_encode_cip105('\x28111ae1...'::bytea, FALSE);
SELECT cardano.tools_drep_id_encode_cip129('\x28111ae1...'::bytea, TRUE);
SELECT cardano.tools_shelley_address_build(
  '\x7415251f...'::bytea, FALSE,
  '\x7c3ae2f2...'::bytea, FALSE,
  0
);
SELECT cardano.tools_shelley_addr_get_type('addr_test1vp6p2fgl...');
```

### 注意事项

- 上游文档要求 PostgreSQL 12+，并通过 `pgrx` 在 Linux 上构建；crate feature 当前面向 PostgreSQL 15 到 18，默认 feature 是 `pg18`。
- 简单 CBOR 辅助函数会有意丢失部分 CBOR 结构信息；如果需要精确的结构往返，请使用 `*_ext` 函数。
