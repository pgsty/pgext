

## 用法

> [pgsmcrypto: PostgreSQL 国密算法扩展](https://github.com/zhuobie/pgsmcrypto)

`pgsmcrypto` 为 PostgreSQL 提供中国国密（SM 系列）算法，包括 SM3 哈希、SM2 非对称加密/签名和 SM4 对称加密。

```sql
CREATE EXTENSION pgsmcrypto;
```

### SM3 消息摘要

```sql
SELECT sm3_hash_string('abc');              -- 返回64字符十六进制字符串（32字节）
SELECT sm3_hash('abc'::bytea);              -- 对 bytea 输入求哈希
SELECT sm3_hash(E'\\x616263');              -- 对原始十六进制输入求哈希
```

### SM2 非对称加密

#### 密钥生成

```sql
SELECT sm2_gen_keypair();                   -- 返回 {私钥, 公钥} 数组
SELECT sm2_privkey_valid('f774...');        -- 验证私钥（1=有效）
SELECT sm2_pubkey_valid('8093...');         -- 验证公钥（1=有效）
SELECT sm2_pk_from_sk('f774...');           -- 从私钥派生公钥
```

#### 密钥导出/导入（PEM）

```sql
SELECT sm2_keypair_to_pem_bytes('f774...');       -- 私钥转 PEM
SELECT sm2_pubkey_to_pem_bytes('8093...');        -- 公钥转 PEM
SELECT sm2_keypair_from_pem_bytes(pem_bytes);     -- 从 PEM 导入
SELECT sm2_pubkey_from_pem_bytes(pem_bytes);      -- 从 PEM 导入公钥
```

#### 签名与验签

```sql
-- 裸签名/验签（直接签名消息）
WITH s AS (
    SELECT sm2_sign_raw('abc'::bytea, 'f774...') AS sig
)
SELECT sm2_verify_raw('abc'::bytea, sig, '8093...') FROM s;

-- 标准签名/验签（SM2规范，带 id + SM3 摘要）
WITH s AS (
    SELECT sm2_sign('myid'::bytea, 'abc'::bytea, 'f774...') AS sig
)
SELECT sm2_verify('myid'::bytea, 'abc'::bytea, sig, '8093...') FROM s;
```

#### 加密与解密

```sql
-- 标准加密/解密
WITH c AS (
    SELECT sm2_encrypt('abc'::bytea, '8093...') AS enc
)
SELECT sm2_decrypt(enc, 'f774...') FROM c;

-- 还提供：sm2_encrypt_c1c2c3、sm2_encrypt_asna1、sm2_encrypt_hex、sm2_encrypt_base64
-- 以及对应的解密变体
```

### SM4 对称加密

```sql
-- ECB 模式（密钥必须为16字节）
SELECT sm4_encrypt_ecb('abc'::bytea, '1234567812345678'::bytea);
SELECT sm4_decrypt_ecb(encrypted, '1234567812345678'::bytea);

-- CBC 模式（密钥和 IV 必须为16字节）
SELECT sm4_encrypt_cbc('abc'::bytea, '1234567812345678'::bytea, '0000000000000000'::bytea);
SELECT sm4_decrypt_cbc(encrypted, '1234567812345678'::bytea, '0000000000000000'::bytea);
```
