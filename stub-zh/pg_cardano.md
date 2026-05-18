## 用法

来源：[README](https://github.com/Fell-x27/pg_cardano/blob/master/README.md), [Cargo.toml version 1.2.0](https://github.com/Fell-x27/pg_cardano/blob/master/Cargo.toml), [releases page](https://github.com/Fell-x27/pg_cardano/releases)

`pg_cardano` 是一个 Rust 扩展，用于在 PostgreSQL 内提供面向 Cardano 的二进制和加密工具。上游仓库没有发布 GitHub releases 或 tags，但默认分支上的当前 crate 版本是 `1.2.0`。

```sql
CREATE EXTENSION pg_cardano;
```

### 核心能力

- Base58 编码/解码。
- Bech32 编码/解码。
- CBOR 与 `jsonb` 互转，包含简单管线和扩展管线。
- Blake2b 哈希。
- Ed25519 签名和签名验证。
- CIP-105 与 CIP-129 的 dRep ID 辅助函数。
- Shelley 地址构造和解析。
- Asset-name 解码，以及 CIP-88 pool key registration 验证。

### 常见模式

Base58 和 Bech32：

```sql
SELECT cardano.base58_encode('Cardano'::bytea);
SELECT cardano.base58_decode('3Z6ioYHE3x');
SELECT cardano.bech32_encode('ada', 'is amazing'::bytea);
SELECT cardano.bech32_decode_prefix('ada1d9ejqctdv9axjmn8dypl4d');
SELECT cardano.bech32_decode_data('ada1d9ejqctdv9axjmn8dypl4d');
```

CBOR 转换：

```sql
SELECT cardano.cbor_decode_jsonb('\xa3636164616b...'::bytea);
SELECT cardano.cbor_encode_jsonb('{"ada": "is amazing!", "bytes": "\\xdeadbeef"}'::jsonb);
SELECT cardano.cbor_decode_jsonb_ext('\xa3636164616b...'::bytea);
SELECT cardano.cbor_encode_jsonb_ext('{"type":"map","value":[...]}'::jsonb);
```

哈希和签名：

```sql
SELECT cardano.blake2b_hash('Cardano is amazing!'::bytea, 32);
SELECT cardano.ed25519_verify_signature(
  '\x432753BD...'::bytea,
  'Cardano is amazing!'::bytea,
  '\x74265f96...'::bytea
);
```

地址和治理辅助函数：

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

- 上游文档说明通过 `pgrx` 支持 PostgreSQL 12+ 和 Linux 构建，但当前 crate features 以及本仓库的包矩阵面向 PostgreSQL 15 到 18，默认 feature 是 `pg18`。
- 本仓库使用 `pgrx` 0.17.0 打包 `pg_cardano` 1.2.0，并带有 `extension.csv` 中记录的 PG18 修复。
- 简单 CBOR 辅助函数对某些 CBOR 构造有意是有损的；需要精确保留结构往返时请使用 `*_ext` 函数。
