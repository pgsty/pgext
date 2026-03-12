

## 用法

> [pgcryptokey: PostgreSQL 加密密钥管理](https://momjian.us/download/pgcryptokey/)

`pgcryptokey` 在 PostgreSQL 内部管理加密数据密钥。密钥以加密方式存储，通过访问密码保护，支持系统级和会话级密钥访问。

```sql
CREATE EXTENSION pgcryptokey;
```

### 密钥管理函数

| 函数 | 描述 |
|----------|-------------|
| `create_cryptokey(name, byte_len)` | 生成新的加密密钥 |
| `set_cryptokey(name)` | 设置当前活动密钥 |
| `get_cryptokey(name)` | 获取密钥材料 |
| `drop_cryptokey(name)` | 删除密钥 |
| `supersede_cryptokey()` | 轮换到新密钥（相同访问密码） |
| `change_key_access_password()` | 更新密钥认证凭据 |
| `reencrypt_data()` | 使用不同密钥重新加密数据 |

### 会话控制

| 函数 | 描述 |
|----------|-------------|
| `get_shared_key()` | 建立客户端/服务器共享密钥（仅 SSL/Unix） |
| `set_session_access_password()` | 客户端提供的密码认证 |

### 典型工作流程

```sql
-- 创建密钥
SELECT create_cryptokey('mykey', 32);

-- 设置活动密钥
SELECT set_cryptokey('mykey');

-- 使用 pgcrypto 函数和托管密钥加密数据
UPDATE secrets SET data = pgp_sym_encrypt(plaintext, get_cryptokey('mykey'));

-- 解密数据
SELECT pgp_sym_decrypt(data, get_cryptokey('mykey')) FROM secrets;

-- 轮换密钥
SELECT supersede_cryptokey();
```

访问密码可以在数据库启动时配置以实现系统级访问，也可以由各个客户端按会话配置以实现细粒度安全控制。
