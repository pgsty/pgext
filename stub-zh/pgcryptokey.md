## 用法

来源：

- [Official 0.85 source archive](https://momjian.us/download/pgcryptokey/pgcryptokey-0.85.tar.gz)
- [Official project directory](https://momjian.us/download/pgcryptokey/)

`pgcryptokey` 管理由访问密码包裹的数据加密密钥。它把包裹后的密钥保存在数据库表中，并结合 `pgcrypto` 完成加密、密钥轮换和重新加密。

### 安装和解锁

```sql
CREATE EXTENSION pgcryptokey CASCADE;
```

依赖扩展是 `pgcrypto`。源码发布版 0.85 使用 SQL 扩展版本 1.0，安装需要超级用户。在客户端模式中，应先按照上游流程，使用 `get_shared_key()` 和 `set_session_access_password(encrypted_password)` 建立会话访问密码。共享密钥交换仅支持 SSL 或 Unix 域套接字连接；加密密码参数使用十六进制编码。

启动模式则预加载 `pgcryptokey_acpass`，执行受保护的服务器端密码获取脚本，并需要重启。这种模式让访问密码在整个服务器范围内生效且只读。应选择一种模式；服务器运行期间不能混用启动和客户端模式。

### 创建和使用密钥

解锁密钥访问后：

```sql
SELECT create_cryptokey('app-key', 32);
SELECT set_cryptokey('app-key');

CREATE TEMP TABLE secrets(ciphertext bytea);
INSERT INTO secrets VALUES
  (pgp_sym_encrypt('example', get_cryptokey('app-key')));
SELECT pgp_sym_decrypt(ciphertext, get_cryptokey('app-key'))
FROM secrets;
```

密钥长度单位为字节。可以按名称或整数密钥 ID 选择密钥；按名称查找只定位当前未被替代的密钥。

### 轮换和重新加密

`supersede_cryptokey(name, byte_len)` 及其密钥 ID 重载创建替代密钥并返回 ID。新旧密钥最初使用相同的访问密码。修改包裹密码时，应使用整数密钥 ID 重载 `change_key_access_password(key_id, new_encrypted_password)`。会话必须已经设置共享密钥和当前访问密码；新密码必须用共享密钥加密，并使用十六进制编码。

源码发布版 0.85 的名称重载调用了未定义的 `change_access_password` 函数，无法完成密码修改。应使用上面的整数重载。

`reencrypt_data(data, old_key_id, new_key_id)` 和 `reencrypt_data_bytea(data, old_key_id, new_key_id)` 迁移加密值。应在密文旁保留原密钥 ID，并在调用 `drop_cryptokey(name)` 或密钥 ID 重载前验证重新加密结果；删除密钥可能让残留密文无法读取。

### 安全边界

应保护密钥表、函数授权、访问密码获取脚本和备份。上游指出，所有用户都能查看启动模式的 `pgcryptokey.access_password`；使用包裹密钥仍需要表权限。`get_cryptokey(name)` 返回原始密钥材料，不应通过普通查询、日志或应用追踪暴露其结果。这种设计不隔离可信数据库管理员对密钥的访问。
