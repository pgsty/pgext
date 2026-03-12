

## 用法

> [pgsodium: 基于 libsodium 的 PostgreSQL 加密函数](https://github.com/michelp/pgsodium)

`pgsodium` 是一个使用 [libsodium](https://download.libsodium.org/doc/) 库的 PostgreSQL 加密扩展。它提供了 libsodium 的直接 SQL 接口、服务器管理的密钥派生和透明列加密（TCE）。

```sql
CREATE EXTENSION pgsodium;
```

### 生成随机数据

```sql
SELECT pgsodium.randombytes_random();
SELECT pgsodium.randombytes_buf(16);         -- 16个随机字节
SELECT pgsodium.randombytes_uniform(100);    -- 0-99之间的随机整数
```

### 对称密钥加密（带认证）

```sql
SELECT * FROM pgsodium.crypto_secretbox_keygen();
SELECT pgsodium.crypto_secretbox('message', nonce, key);
SELECT pgsodium.crypto_secretbox_open(ciphertext, nonce, key);
```

### 公钥加密

```sql
SELECT * FROM pgsodium.crypto_box_new_keypair();
SELECT pgsodium.crypto_box('message', nonce, public_key, secret_key);
SELECT pgsodium.crypto_box_open(ciphertext, nonce, public_key, secret_key);
```

### 公钥签名

```sql
SELECT * FROM pgsodium.crypto_sign_new_keypair();
SELECT pgsodium.crypto_sign('message', secret_key);
SELECT pgsodium.crypto_sign_open(signed_message, public_key);
```

### 密码哈希

```sql
SELECT pgsodium.crypto_pwhash_str('my_password');
SELECT pgsodium.crypto_pwhash_str_verify(hash, 'my_password');
```

### 哈希

```sql
SELECT pgsodium.crypto_generichash('data');
SELECT pgsodium.crypto_shorthash('data', key);
```

### 服务器密钥管理

pgsodium 可以将外部根密钥加载到内存中，该密钥永远无法通过 SQL 访问。子密钥通过密钥 ID 派生：

```sql
SELECT * FROM pgsodium.create_key();
-- 返回一个 UUID 密钥 ID，用于 TCE 或加密函数
```

### 透明列加密（TCE）

```sql
CREATE TABLE private.users (
    id bigserial PRIMARY KEY,
    secret text
);

SECURITY LABEL FOR pgsodium ON COLUMN private.users.secret
  IS 'ENCRYPT WITH KEY ID dfc44293-fa78-4a1a-9ef9-7e600e63e101';
```

加密数据存储在磁盘上，通过生成的视图自动解密。

### 安全角色

- `pgsodium_keyiduser` -- 较低权限，只能通过 UUID 访问密钥
- `pgsodium_keymaker` -- 较高权限，可以操作原始密钥
