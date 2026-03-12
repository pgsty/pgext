

## 用法

> [pg_tde: PostgreSQL 透明数据加密](https://github.com/percona/pg_tde)

`pg_tde` 在文件级别提供透明数据加密（TDE），对元组、WAL 和索引进行加密。它使用 `tde_heap` 访问方法，支持基于文件的密钥环和外部密钥管理系统（KMS）。

```sql
CREATE EXTENSION pg_tde;
```

### 配置

添加到 `postgresql.conf`：

```ini
shared_preload_libraries = 'pg_tde'
```

### 设置密钥提供者

```sql
-- 基于文件的密钥提供者（数据库级别）
SELECT pg_tde_add_database_key_provider_file('file_keyring', '/path/to/keyring');

-- 或全局级别的密钥提供者
SELECT pg_tde_add_global_key_provider_file('file_keyring', '/path/to/keyring');

-- 使用数据库密钥提供者设置加密密钥
SELECT pg_tde_set_key_using_database_key_provider('my_key', 'file_keyring');

-- 或使用全局密钥提供者
SELECT pg_tde_set_key_using_global_key_provider('my_key', 'file_keyring');
```

### 创建加密表

```sql
CREATE TABLE sensitive_data (
    id serial PRIMARY KEY,
    secret text
) USING tde_heap;
```

使用 `USING tde_heap` 创建的表中的所有数据都会在磁盘上透明加密。

### 检查加密状态

```sql
SELECT pg_tde_is_encrypted('sensitive_data');
```

### 附加函数

| 函数 | 描述 |
|----------|-------------|
| `pg_tde_add_database_key_provider_file(name, path)` | 添加基于文件的数据库密钥提供者 |
| `pg_tde_add_global_key_provider_file(name, path)` | 添加基于文件的全局密钥提供者 |
| `pg_tde_add_database_key_provider_vault_v2(...)` | 添加 HashiCorp Vault 数据库密钥提供者 |
| `pg_tde_add_global_key_provider_vault_v2(...)` | 添加 HashiCorp Vault 全局密钥提供者 |
| `pg_tde_set_key_using_database_key_provider(key, provider)` | 通过数据库提供者设置加密密钥 |
| `pg_tde_set_key_using_global_key_provider(key, provider)` | 通过全局提供者设置加密密钥 |
| `pg_tde_is_encrypted(table)` | 检查表是否已加密 |

### 注意事项

- 仅适用于 Percona Server for PostgreSQL 17+
- 加密元组、WAL 和索引
- 尚不支持加密临时文件和统计信息
