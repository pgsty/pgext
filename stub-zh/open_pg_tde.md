## 用法

来源：

- [官方文档](https://commandprompt.github.io/open_pg_tde/)
- [官方2.4.0 README](https://github.com/commandprompt/open_pg_tde/blob/2.4.0/README.md)
- [官方扩展控制文件](https://github.com/commandprompt/open_pg_tde/blob/2.4.0/open_pg_tde.control)

`open_pg_tde` 2.4.0 通过 `tde_heap` 表访问方法为上游 PostgreSQL 提供透明的静态加密。它在允许加密表和平文表共存的同时，对表数据、索引、TOAST、WAL 和临时溢出文件进行加密。

### 核心工作流

构建与 `open_pg_tde` 核心补丁匹配的 PostgreSQL，安装扩展，将 `open_pg_tde` 添加到 `shared_preload_libraries` 中，然后重启服务器。在创建加密表之前，配置密钥提供者和主密钥：

```sql
CREATE EXTENSION open_pg_tde;

SELECT open_pg_tde_add_database_key_provider_file(
  'file-keyring',
  '/var/lib/postgresql/open_pg_tde.keyring'
);
SELECT open_pg_tde_create_key_using_database_key_provider(
  'primary-key',
  'file-keyring'
);
SELECT open_pg_tde_set_key_using_database_key_provider(
  'primary-key',
  'file-keyring'
);

CREATE TABLE secret_data (
  id bigint PRIMARY KEY,
  payload text
) USING tde_heap;

SELECT open_pg_tde_is_encrypted('secret_data');
```

当威胁模型要求与数据库主机分离时，请使用生产密钥管理提供商而不是文件密钥环示例。

### 重要对象与设置

- `tde_heap` 是加密表访问方法。
- `open_pg_tde.data_cipher` 选择用于关系数据的 AES-128-XTS 或 AES-256-XTS。
- 数据库和服务器密钥提供者函数注册密钥环、KMIP 兼容或 OpenBao 提供商，并选择主密钥。
- `open_pg_tde_is_encrypted(regclass)` 检查关系是否使用加密存储。

### 要求与注意事项

- 版本 2.4.0 支持上游 PostgreSQL 16、17 和 18，根据其标记的 README 文件。
- 该扩展需要与补丁匹配的 PostgreSQL 源树；未打补丁的原版服务器不提供所需的存储管理器和 WAL 挂钩。
- 更改 `shared_preload_libraries` 需要重启服务器。
- 系统目录和统计信息在声明的加密覆盖范围之外。请查阅官方威胁模型，了解备份、密钥、内存、日志和主机妥协的情况。
- 在生产使用之前，请单独备份密钥材料并测试恢复、轮换、故障转移、WAL 重放和提供者丢失的行为。
