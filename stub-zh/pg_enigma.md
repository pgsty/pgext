


## 用法

> [pg_enigma: 使用 PGP 和 RSA 密钥的 PostgreSQL 加密数据类型](https://github.com/SoftwareLibreMx/pg_enigma)

`pg_enigma` 为 PostgreSQL 提供 `Enigma` 加密数据类型，使用 PGP 或 OpenSSL RSA 密钥对静态数据进行加密。数据以加密形式存储，仅在私钥加载到内存中时才进行解密。

```sql
CREATE EXTENSION IF NOT EXISTS pg_enigma;
```

### PGP 密钥加密

```sql
-- 创建带加密列的表（密钥槽 2）
CREATE TABLE test_pgp (
    id SERIAL,
    val Enigma(2)
);

-- 加载公钥用于加密
SELECT set_public_key_from_file(2, '/path/to/public-key.asc');

-- 插入数据（自动使用公钥加密）
INSERT INTO test_pgp (val) VALUES ('A secret value'::Text);

-- 没有私钥时，SELECT 返回加密的 PGP 消息
SELECT * FROM test_pgp;

-- 加载私钥以启用解密
SELECT set_private_key_from_file(2, '/path/to/private-key.asc', 'passphrase');

-- 现在 SELECT 返回解密的明文
SELECT * FROM test_pgp;
-- id |      val
-- ----+----------------
--   1 | A secret value

-- 从内存中移除私钥
SELECT forget_private_key(2);
-- 后续 SELECT 再次返回加密数据
```

### RSA 密钥加密

```sql
CREATE TABLE test_rsa (
    id SERIAL,
    val Enigma(3)
);

SELECT set_public_key_from_file(3, '/path/to/alice_public.pem');
INSERT INTO test_rsa (val) VALUES ('Another secret value'::Text);

SELECT set_private_key_from_file(3, '/path/to/alice_private.pem', 'passphrase');
SELECT * FROM test_rsa;

SELECT forget_private_key(3);
```

### 函数

| 函数 | 描述 |
|------|------|
| `set_public_key_from_file(slot, path)` | 加载公钥用于加密 |
| `set_private_key_from_file(slot, path, passphrase)` | 加载私钥用于解密 |
| `forget_private_key(slot)` | 从内存中移除私钥 |
