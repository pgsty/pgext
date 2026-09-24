## 用法

来源：

- [PGXN 0.1.1 README](https://pgxn.org/dist/chdb/0.1.1/README.html)
- [chdb_hook 0.1.1 文档](https://pgxn.org/dist/chdb/0.1.1/doc/chdb_hook.html)
- [chdb 0.1.1 发行元数据](https://api.pgxn.org/src/chdb/chdb-0.1.1/META.json)
- [Apache 2.0 许可证](https://api.pgxn.org/src/chdb/chdb-0.1.1/LICENSE.md)
- [0.1.1 changelog](https://api.pgxn.org/src/chdb/chdb-0.1.1/CHANGELOG.md)

`chdb_hook` 是 `chdb` 发行包附带的无 SQL 模块。它拦截基于 URL 的 `COPY` 与部分 `CREATE TABLE` 语句，让 chDB 读写本地文件、HTTP 资源、S3、Google Cloud Storage、Azure Blob/ABFS 与 HDFS 格式。它没有控制文件，也不创建 SQL 对象。

### 会话级流程

```sql
LOAD 'chdb_hook';

CREATE TABLE times (
    id integer,
    months integer,
    days integer
);

COPY times
FROM 's3://datasets-documentation/my-test-bucket-768/some_prefix/some_file_1.csv';
```

可以为单个会话显式加载模块，也可以通过 `session_preload_libraries` 为指定角色或数据库加载，或通过 `shared_preload_libraries` 全局加载。只有最后一种方式要求重启 PostgreSQL。

### 支持路径与权限

钩子识别 `file`、`http`、`https`、`s3`、`gs`、`gcs`、`oss`、`az`、`azure`、`abfss`、`abfs` 与 `hdfs` URL 方案。它保留普通关系权限：`COPY TO` 需要 `SELECT`，`COPY FROM` 需要 `INSERT` 与读写事务。本地 `file` 访问还需要 `pg_read_server_files` 或 `pg_write_server_files`，并满足操作系统权限。

表选项 `structure_from` 与 `copy_from` 可以从 URL 推断列或填充新表。凭据、格式、压缩、超时、通配符与显式结构都会传入 chDB 路径；应尽量避免让秘密进入语句日志或目录可见的表选项。

### 边界

加载 `chdb_hook` 会改变当前会话中普通 PostgreSQL 命令的解析与执行。授予服务器文件角色后，也会通过钩子获得云端与 HTTP 路径访问能力，因此只能面向可信 SQL 调用者，并限制网络出站。格式桥接对 NULL、数组、JSON、几何、时间、Protobuf、Parquet、Arrow 与编码存在已记录的边界情况；用于迁移或备份前应测试代表性往返转换。

0.1.1 与 `chdb` 使用相同的辅助进程及 libchdb 26.7.0 或以上，并支持 Linux 或 macOS 上的 PostgreSQL 15 及以上。辅助进程失败会中止发起的操作。大型导入导出可能消耗大量内存、CPU、磁盘与网络带宽，因此应设置 chDB 资源限制并监控主机。

### 0.1.1 版本变化

这是二进制更新，`chdb` 控制版本仍为 `0.1`，无需执行扩展 SQL 升级。该版增加 PostgreSQL 15 支持，大型无符号整数及 128/256 位整数映射为 `numeric`，并增加 BFloat16 和时间间隔映射。字符串转换会验证数据库编码，时间戳文本导出使用 ISO-8601 并按文档处理时区。替换并重新加载二进制后应验证代表性数据的往返转换。
