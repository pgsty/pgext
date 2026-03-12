

## 用法

> [pguecc: PostgreSQL 的椭圆曲线加密（micro-ecc 绑定）](https://github.com/ameensol/pg-ecdsa)

需要 `pgcrypto` 扩展。

### 生成密钥对

```sql
SELECT ecdsa_make_key('secp256k1');
-- (public_key_hex, private_key_hex)
```

### 签名数据

```sql
SELECT ecdsa_sign(
    '000000000000000000000000000000000000000000',  -- 私钥（十六进制）
    '1234',                                          -- 要签名的数据
    'sha256',                                        -- 哈希函数
    'secp160r1'                                      -- 曲线名称
);
```

### 验证签名

```sql
SELECT ecdsa_verify(
    '<public_key_hex>',
    'hello, world',
    '<signature_hex>',
    'sha256',
    'secp256k1'
);
-- t
```

### 密钥验证

```sql
SELECT ecdsa_is_valid_public_key('<public_key_hex>', 'secp256k1');
SELECT ecdsa_is_valid_private_key('<private_key_hex>', 'secp256k1');
SELECT ecdsa_is_valid_curve('secp256k1');  -- true
```

### 支持的曲线

`secp160r1`、`secp192r1`、`secp224r1`、`secp256r1`、`secp256k1`

### 原始 API

直接签名（不进行哈希处理，请谨慎使用）：

```sql
SELECT ecdsa_sign_raw(private_key_bytea, hash_bytea, 'secp256k1');
SELECT ecdsa_verify_raw(public_key_bytea, hash_bytea, signature_bytea, 'secp256k1');
```
