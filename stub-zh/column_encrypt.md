


## 用法

来源：[README](https://github.com/vibhorkum/column_encrypt/blob/master/README.md)，[v4.0 release](https://github.com/vibhorkum/column_encrypt/releases/tag/v4.0)，[SQL objects](https://github.com/vibhorkum/column_encrypt/blob/master/column_encrypt--4.0.sql)

`column_encrypt` 为 PostgreSQL 提供透明列级加密。它定义 `encrypted_text` 和 `encrypted_bytea` 类型，通过类型输入函数加密写入值，通过输出函数解密读取值，并在 `encrypt` schema 中管理数据加密密钥。

### 启用

先在服务器启动时加载共享库，重启 PostgreSQL，然后创建 schema 和扩展：

```conf
shared_preload_libraries = 'column_encrypt'
```

```sql
CREATE EXTENSION pgcrypto;
CREATE SCHEMA IF NOT EXISTS encrypt;
CREATE EXTENSION column_encrypt;
```

可以把 `encrypt` 加入 `search_path`，也可以显式使用 schema 前缀。

### 注册与加载密钥

```sql
SELECT encrypt.register_key('my-secret-data-key', 'my-master-passphrase');
SELECT encrypt.load_key('my-master-passphrase');

SELECT * FROM encrypt.keys();
SELECT * FROM encrypt.status();
```

扩展使用两层密钥模型，包括密钥加密密钥和数据加密密钥。密文携带 key version 头部，因此轮换密钥后仍可解密旧值。

### 加密列

```sql
CREATE TABLE secure_data (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  ssn encrypt.encrypted_text,
  payload encrypt.encrypted_bytea
);

INSERT INTO secure_data (ssn, payload)
VALUES ('888-999-2045', decode('aabbcc', 'hex'));

SELECT id, ssn FROM secure_data;
```

当前会话没有加载密钥时，读取加密值会报解密错误。

### 密钥操作

常用函数包括 `encrypt.activate_key`、`encrypt.revoke_key`、`encrypt.rotate`、`encrypt.verify`、`encrypt.unload_key`、`encrypt.loaded_cipher_key_versions`、`encrypt.blind_index`。

不能直接暴露明文值的查找场景，可以使用 blind index：

```sql
SELECT encrypt.blind_index('888-999-2045', 'lookup-hmac-key');
```

### 注意事项

扩展有意拒绝加密值的 binary send/receive。相等和哈希语义基于解密后的明文；不支持范围排序。从旧的密文哈希语义升级后，需要重建加密列上的哈希索引。
