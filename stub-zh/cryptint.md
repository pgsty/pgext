

## 用法

> [cryptint: 使用 SKIP32 和 XTEA 加解密整数](https://github.com/dverite/cryptint)

### SKIP32（32 位整数加密）

基于 Skipjack 的 24 轮分组密码。使用 80 位（10 字节）密钥加密 32 位值。

```sql
SELECT skip32_encrypt(42, '\x00010203040506070809'::bytea);
-- 加密后的 int4

SELECT skip32_decrypt(175586429, '\x00010203040506070809'::bytea);
-- 0
```

往返示例：

```sql
SELECT i, skip32_encrypt(i, '\x00010203040506070809'::bytea) AS encrypted,
       skip32_decrypt(skip32_encrypt(i, '\x00010203040506070809'::bytea), '\x00010203040506070809'::bytea) AS decrypted
FROM generate_series(-3, 3) AS i;
```

### XTEA（64 位整数加密）

64 轮分组密码。使用 128 位（16 字节）密钥加密 64 位值。

```sql
SELECT xtea_encrypt(42::bigint, '\x000102030405060708090a0b0c0d0e0f'::bytea);
-- 加密后的 bigint

SELECT xtea_decrypt(103200416458222088::bigint, '\x000102030405060708090a0b0c0d0e0f'::bytea);
-- 1
```

### 函数参考

| 函数 | 描述 |
|------|------|
| `skip32_encrypt(int, bytea)` | 使用 10 字节密钥加密 int4 |
| `skip32_decrypt(int, bytea)` | 使用 10 字节密钥解密 int4 |
| `xtea_encrypt(bigint, bytea)` | 使用 16 字节密钥加密 bigint |
| `xtea_decrypt(bigint, bytea)` | 使用 16 字节密钥解密 bigint |
