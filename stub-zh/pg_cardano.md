


## 用法

> [pg_cardano: PostgreSQL 的 Cardano 区块链数据类型和加密函数](https://github.com/Fell-x27/pg_cardano)

`pg_cardano` 扩展为在 PostgreSQL 中处理 Cardano 区块链数据提供加密和实用函数，包括 Base58、Bech32、CBOR、Blake2b、Ed25519 和 Shelley 地址工具。

```sql
CREATE EXTENSION pg_cardano;
```

### Base58 编码/解码

```sql
SELECT cardano.base58_encode('Cardano'::bytea);
-- '3Z6ioYHE3x'

SELECT cardano.base58_decode('3Z6ioYHE3x');
-- '\x43617264616e6f'
```

### Bech32 编码/解码

```sql
SELECT cardano.bech32_encode('ada', 'is amazing'::bytea);
-- 'ada1d9ejqctdv9axjmn8dypl4d'

SELECT cardano.bech32_decode_prefix('ada1d9ejqctdv9axjmn8dypl4d');
-- 'ada'

SELECT cardano.bech32_decode_data('ada1d9ejqctdv9axjmn8dypl4d');
-- '\x697320616d617a696e67'
```

### CBOR 编码/解码

简单 CBOR（轻量级，满足大多数 Cardano 场景）：

```sql
SELECT cardano.cbor_decode_jsonb('\xa3636164616b...'::bytea);
SELECT cardano.cbor_encode_jsonb('{"ada": "is amazing!", "bytes": "\\xdeadbeef"}'::jsonb);
```

扩展 CBOR（完整支持，无限制）：

```sql
SELECT cardano.cbor_decode_jsonb_ext('\xa3636164616b...'::bytea);
SELECT cardano.cbor_encode_jsonb_ext('{"type":"map","value":[...]}'::jsonb);
```

### Blake2b 哈希

```sql
SELECT cardano.blake2b_hash('Cardano is amazing!'::bytea, 32);
-- '\x2244d5c9699fa93b...'
```

### Ed25519 签名与验证

```sql
SELECT cardano.ed25519_sign_message(
    '\x43D68AEC...'::bytea,       -- 签名密钥
    'Cardano is amazing!'::bytea   -- 消息
);

SELECT cardano.ed25519_verify_signature(
    '\x432753BD...'::bytea,        -- 验证密钥
    'Cardano is amazing!'::bytea,  -- 消息
    '\x74265f96...'::bytea         -- 签名
);
-- true
```

### dRep 视图 ID 构建器 (CIP-105/CIP-129)

```sql
SELECT cardano.tools_drep_id_encode_cip105('\x28111ae1...'::bytea, FALSE);
-- 'drep_vkh19qg34ctllr7lh48nnj4akfc978qzqzuwzkgsdu6r3zc42lnl6a0'

SELECT cardano.tools_drep_id_encode_cip129('\x28111ae1...'::bytea, TRUE);
-- 'drep1yv5pzxhp0lu0m7757ww2hke8qhcuqgqt3c2ezphngwytz4gj324g7'
```

### Shelley 地址工具

```sql
-- 构建基础地址
SELECT cardano.tools_shelley_address_build(
    '\x7415251f...'::bytea, FALSE,  -- 支付凭证, 是否为脚本
    '\x7c3ae2f2...'::bytea, FALSE,  -- 权益凭证, 是否为脚本
    0                                -- 网络 ID
);

-- 提取支付/权益凭证
SELECT cardano.tools_shelley_addr_extract_payment_cred('addr_test1qp6p2fgl...');
SELECT cardano.tools_shelley_addr_extract_stake_cred('addr_test1qp6p2fgl...');

-- 获取地址类型
SELECT cardano.tools_shelley_addr_get_type('addr_test1vp6p2fgl...');
-- 'PMT_KEY:NONE'
```

### 资产名称转换

```sql
SELECT cardano.tools_read_asset_name('hello'::bytea);
-- 'hello'（有效 UTF-8 作为字符串返回）

SELECT cardano.tools_read_asset_name('\xdeadbeef'::bytea);
-- 'deadbeef'（非 UTF-8 以十六进制返回）
```

### CIP-88 矿池密钥注册验证

```sql
SELECT cardano.tools_verify_cip88_pool_key_registration('\xa119...'::bytea);
-- true
```
