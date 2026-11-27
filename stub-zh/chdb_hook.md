## 用法

来源：

- [doc/chdb_hook.md](https://api.pgxn.org/src/chdb/chdb-0.1.4/doc/chdb_hook.md)
- [chdb_hook.control](https://api.pgxn.org/src/chdb/chdb-0.1.4/chdb_hook.control)
- [sql/chdb_hook.sql](https://api.pgxn.org/src/chdb/chdb-0.1.4/sql/chdb_hook.sql)
- [Makefile](https://api.pgxn.org/src/chdb/chdb-0.1.4/Makefile)
- [CHANGELOG.md](https://api.pgxn.org/src/chdb/chdb-0.1.4/CHANGELOG.md)
- [LICENSE.md](https://api.pgxn.org/src/chdb/chdb-0.1.4/LICENSE.md)

`chdb_hook` 拦截基于 URL 的 `COPY` 和扩展 `CREATE TABLE` 命令，通过 chDB 搬运数据。发行版 0.1.4 新增控制文件和 SQL 版本 0.1，提供 `chdb_hook_version()`。仍可直接加载模块，因此钩子流程不强制创建扩展。

### 加载与使用

超级用户可在单个会话中启用：

```sql
LOAD 'chdb_hook';
CREATE TABLE times (id integer, months integer, days integer);
COPY times
FROM 's3://datasets-documentation/my-test-bucket-768/some_prefix/some_file_1.csv';
```

新的可选 SQL 包装可按以下方式安装：

`CREATE EXTENSION chdb_hook` / `SELECT chdb_hook_version()`

`CREATE EXTENSION` 只在执行命令的当前会话加载模块。后续会话仍需使用 `LOAD` 或调用 `chdb_hook_version()`，除非已配置预加载。`session_preload_libraries` 可面向指定数据库或角色，`shared_preload_libraries` 则作用于整个集群并要求重启服务器。SQL 包装不受信任且可重定位。

### 数据源与权限

支持的 URL 方案包括 `file`、`http`、`https`、`s3`、`gs`、`gcs`、`oss`、`az`、`azure`、`abfss`、`abfs` 和 `hdfs`。普通关系权限继续生效：`COPY TO` 需要 `SELECT`，`COPY FROM` 需要 `INSERT` 和读写事务。本地文件还需 `pg_read_server_files` 或 `pg_write_server_files` 及操作系统权限。构建时可通过 `NO_FILE_SCHEME` 禁用本地 URL。

表选项 `structure_from` 和 `copy_from` 支持模式推断及数据装载。格式、结构、压缩、超时和凭据通过显式 COPY 选项传入。`encoding_check` 可取 `fail`、`replace`、`remove` 或 `truncate`；基于 URL 的 `COPY FROM` 不允许带 `WHERE`。

### 主机凭据与资源限制

`chdb_hook.host_credentials_roles` 由超级用户配置，内容为已有角色列表，默认为空。只有这些角色的成员才能在未提供显式密钥时使用主机 AWS 凭据。S3 的 `role_arn` 选项可用这些凭据代入 IAM 角色，要求同样的角色成员资格，且仅在未提供 `access_key` 时生效。

`chdb_hook.max_memory` 默认改为 4 GB，`0` 表示不限额。`chdb_hook.max_threads` 和 `chdb_hook.max_parsing_threads` 控制并发，均需超级用户权限。大型导入应设置明确的资源上限。

### 维护边界

该模块在 Linux 或 macOS 上支持 PostgreSQL 15 及以上，需要匹配的辅助程序及 libchdb。受支持的转换会保留声明的时间戳精度和会话 `TimeZone`。应验证复杂类型及编码的往返转换，并避免凭据进入 SQL 日志，尤其是调试输出。已有直接加载方式可继续不创建 SQL 对象，版本函数则需要安装新的 SQL 包装。
