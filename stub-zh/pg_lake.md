## 用法

来源：

- [官方 pg_lake README](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/README.md)
- [版本 3.4 控制文件](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake/pg_lake.control)
- [官方构建和启动指南](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/building-from-source.md)
- [官方项目文档索引](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/README.md)
- [DuckDB Secret 管理器](https://duckdb.org/docs/stable/configuration/secrets_manager.html)

`pg_lake` 是 Snowflake PostgreSQL 湖仓套件的顶层扩展。它会安装查询对象存储文件和创建事务型 Iceberg 表所需的表访问、Iceberg、COPY、查询引擎、扩展基础设施和映射组件。PostgreSQL 扩展负责查询规划与事务协调，独立的本地 `pgduck_server` 进程则使用 DuckDB 执行向量化计算。

### 启动打包版服务栈

版本 `3.4` 支持 PostgreSQL 16 至 18。PIGSTY 的 RPM 与 DEB 包会安装扩展文件和带版本路径的 `pgduck_server` 二进制，但目前不会安装或自动启动 `systemd` 服务；执行 `CREATE EXTENSION` 也不会拉起 `pgduck_server`。

将 `pg_extension_base` 加入 `shared_preload_libraries`，然后重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_extension_base'
```

`pgduck_server` 默认监听 `/tmp/.s.PGSQL.5332`，socket 权限为 `0770`。应使用 PostgreSQL 的操作系统用户启动它，确保 PostgreSQL 可以访问该 socket；不要直接使用无关的登录用户运行裸命令。

```shell
# Debian/Ubuntu with PostgreSQL 18; use 16 or 17 as appropriate.
PG_LAKE_SERVER=/usr/lib/postgresql/18/bin/pgduck_server
# RHEL-compatible systems use /usr/pgsql-18/bin/pgduck_server.

sudo install -d -o postgres -g postgres -m 0700 \
  /var/lib/pg_lake /var/lib/pg_lake/extensions
sudo install -d -o postgres -g postgres -m 0750 /var/cache/pg_lake
sudo -u postgres -H "$PG_LAKE_SERVER" \
  --duckdb_database_file_path /var/lib/pg_lake/pgduck_server.db \
  --extensions_dir /var/lib/pg_lake/extensions \
  --cache_dir /var/cache/pg_lake
```

该命令会以前台方式持续运行，不能退出；生产环境应交由服务管理器监管。如果使用独立服务账号，应将其加入 `postgres` 用户组，并使用 `--unix_socket_group postgres --unix_socket_permissions 0770` 启动服务。

在另一个终端中先验证查询引擎，再创建扩展：

```shell
sudo -u postgres psql -X \
  "host=/tmp port=5332 dbname=postgres connect_timeout=2" \
  -c 'SELECT version();'
```

然后在目标数据库中创建完整的依赖树：

```sql
CREATE EXTENSION pg_lake CASCADE;
SELECT lake.version();
```

### 配置对象存储访问

对象存储凭据由 `pgduck_server` 解析，而不是由 PostgreSQL 后端解析。AWS 与 GCP 可以使用各自的标准凭据链。对于 MinIO 等本地 S3 兼容服务，先创建存储桶，再直接连接 `pgduck_server` 创建持久化的 DuckDB Secret：

```shell
sudo -u postgres psql -X -h /tmp -p 5332 -d postgres
```

```sql
CREATE PERSISTENT SECRET pglake_object_store (
    TYPE S3,
    KEY_ID 'access-key',
    SECRET 'secret-key',
    REGION 'us-east-1',
    ENDPOINT 'minio.example.com:9000',
    SCOPE 's3://analytics-bucket',
    URL_STYLE 'path',
    USE_SSL false
);
```

连接 PostgreSQL，并在创建表的同一会话中指定托管 Iceberg 数据的位置：

```sql
SET pg_lake_iceberg.default_location_prefix =
    's3://analytics-bucket/warehouse';
```

### 核心工作流

创建和修改事务型 Iceberg 表：

```sql
CREATE TABLE measurements (
    station_name text NOT NULL,
    measured_at timestamptz NOT NULL,
    value double precision
) USING iceberg;

INSERT INTO measurements VALUES
    ('Istanbul', now(), 18.5),
    ('Haarlem', now(), 9.3);
```

通过 `COPY` 导入或导出 Parquet、CSV 或换行符分隔的 JSON 文件：

```sql
COPY (SELECT * FROM measurements)
TO 's3://analytics-bucket/export/measurements.parquet';

COPY measurements
FROM 's3://analytics-bucket/import/measurements.parquet';
```

直接查询文件，无需先将其载入 PostgreSQL：

```sql
CREATE FOREIGN TABLE external_events ()
SERVER pg_lake
OPTIONS (path 's3://analytics-bucket/events/*.parquet');

SELECT count(*) FROM external_events;
```

### 组件索引

- `pg_lake`：元扩展，并提供 `lake.version()`。
- `pg_lake_table`：数据湖 FDW、Iceberg 表语法、文件工具和表目录。
- `pg_lake_iceberg`：Iceberg 元数据、快照、清单和目录集成。
- `pg_lake_copy`：拦截针对对象存储文件和湖格式的 `COPY`。
- `pg_lake_engine`：共享查询重写、类型转换、清理以及 `pgduck_server` 客户端层。
- `pg_extension_base`：预加载和生命周期工作进程基础设施。
- `pg_map`：为嵌套湖数据生成 PostgreSQL 映射类型。

### 运营注意事项

- 每台可能执行湖查询的 PostgreSQL 主机都必须运行 `pgduck_server`。应使用服务管理器监管进程，并在承载流量前验证本地 socket。
- 默认 socket 权限为 `0770`，其属主和属组取决于启动 `pgduck_server` 的账号。服务用户不匹配时会出现 `ERROR: could not start query engine`。
- S3 及兼容存储的凭据由 DuckDB Secret/凭据链解析。仅授予工作负载所需的存储桶权限。
- 首次启动可能下载 DuckDB Spatial 扩展。请确保服务账号具备必要的网络权限，并能写入状态目录和缓存目录。
- `CREATE PERSISTENT SECRET` 创建的凭据可跨服务重启保留，但 DuckDB 会将其以未加密形式存放在 `~/.duckdb/stored_secrets`。应保持服务账号及其主目录不变，限制文件权限，并按凭据进行保护。
- 默认内存上限为系统内存的 80%。当 PostgreSQL 与 `pgduck_server` 共用生产主机时，应显式设置 `--memory_limit`。
- Iceberg 写入会按语句创建 Parquet 文件。应批量插入并定期运行 `VACUUM`，避免产生大量小文件。
- PostgreSQL 扩展、`pgduck_server`、对象存储数据和 Iceberg 目录共同构成一个部署单元。单独创建扩展并不能证明外部服务可用；备份和升级时也应分别核验这些组件。
