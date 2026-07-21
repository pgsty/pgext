## 用法

来源：

- [pg_tde 2.2 setup](https://github.com/percona/pg_tde/blob/2.2.0/documentation/docs/setup.md)
- [Key-provider and key-management functions](https://github.com/percona/pg_tde/blob/2.2.0/documentation/docs/functions.md)
- [pg_tde 2.2.0 release notes](https://github.com/percona/pg_tde/blob/2.2.0/documentation/docs/release-notes/release-notes-v2.2.0.md)
- [Transparent data encryption limitations](https://github.com/percona/pg_tde/blob/2.2.0/documentation/docs/index/tde-limitations.md)
- [TDE table access method](https://github.com/percona/pg_tde/blob/2.2.0/documentation/docs/index/table-access-method.md)
- [WAL encryption](https://github.com/percona/pg_tde/blob/2.2.0/documentation/docs/wal-encryption.md)

pg_tde 为 Percona Server for PostgreSQL 提供透明数据加密功能。它通过 tde_heap 访问方法对表数据进行加密，并可以加密 WAL，密钥由文件、HashiCorp Vault 或 KMIP 提供商管理。它并不是可直接用于社区版 PostgreSQL 的即插即用扩展。

### 预加载和创建扩展

添加库并重启服务器：

    shared_preload_libraries = 'pg_tde'

然后在将使用加密表的每个数据库中启用 pg_tde：

    CREATE EXTENSION pg_tde;

以超级用户或具有适当权限的数据库所有者身份运行设置。上游 pg_tde 2.2 版本与兼容的 Percona Server for PostgreSQL 17 或 18 版本绑定；2.2.0 发行说明警告称，它不适用于比 17.10 和 18.4 更早版本的 Percona Distribution。

### 配置密钥提供者

注册一个提供者，然后设置主密钥。本地文件提供者适合用于评估：

    SELECT pg_tde_add_database_key_provider_file(
      'local-file',
      '/secure/path/pg_tde_keys'
    );

    SELECT pg_tde_set_principal_key(
      'app-principal-key',
      'local-file'
    );

对于生产环境，上游建议使用 Vault 或 KMIP 等外部提供者而不是本地文件提供者。独立于数据库文件保护提供者的凭据、密钥文件、备份和恢复程序。

提供者管理包括针对文件、Vault 和 KMIP 提供者的数据库范围和服务器全局变体，以及列出、更改和删除提供者、检查或轮换主密钥的函数。

### 创建并转换加密表

使用加密访问方法创建一个表：

    CREATE TABLE customer_secrets (
      id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
      payload jsonb NOT NULL
    ) USING tde_heap;

仅在测试过锁、表重写、磁盘空间和备份影响后，才转换现有表：

    ALTER TABLE customer_secrets SET ACCESS METHOD tde_heap;

更改表访问方法会重写整张表。请规划维护窗口，并在预演副本上确认索引、备库、备份与恢复流程。

### 启用 WAL 加密

WAL 加密是一个单独的服务器配置项：

    pg_tde.wal_encrypt = on

更改它需要重启。在启用之前，请确保每个主库、备库、备份、归档和恢复主机都具有所需的提供者配置和密钥访问权限。

### 对象索引

- tde_heap：加密表访问方法。
- pg_tde_add_database_key_provider_file/vault/kmip：数据库范围的提供者注册。
- pg_tde_add_global_key_provider_file/vault/kmip：服务器全局的提供者注册。
- pg_tde_set_principal_key 和 pg_tde_set_server_principal_key：选择用于保护数据加密密钥的主密钥。
- pg_tde_list_all_key_providers：检查已注册的提供者。
- pg_tde_change_key_provider_* 和 pg_tde_delete_key_provider：管理提供者的定义。
- pg_tde.wal_encrypt：启用写前日志记录的加密。
- pg_tde_upgrade：在 2.2 系列中引入的升级助手。

### 安全性和恢复边界

- pg_tde 会加密受支持的用户表存储，但不会加密所有 PostgreSQL 产物；系统目录、规划器统计信息和临时溢出文件属于文档明确列出的例外。
- 上游警告称，在已发生分歧的节点之间使用 pg_rewind 或 pg_tde_rewind 可能会损坏集群。请遵循文档中的重建/恢复路径，不要假设常规 rewind 是安全的。
- 在未预加载 pg_tde 的情况下开始恢复可能会损坏加密数据。确保在库和密钥存在的情况下验证灾难恢复自动化。
- Percona 文档中提到了与 Citus 和 TimescaleDB 的不兼容性以及对某些 WAL 检查和恢复工具的限制。
- 加密并不能替代 SQL 权限、TLS、主机加固、审计日志或经过测试的备份。丢失密钥可能会使原本完好的备份无法恢复。
