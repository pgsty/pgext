## 用法

来源：

- [0.3.3 setup and upgrades](https://github.com/pgsty/pg_mooncake/blob/v0.3.3/README.md)
- [0.3.3 SQL procedures](https://github.com/pgsty/pg_mooncake/blob/v0.3.3/src/functions.rs)
- [0.3.3 dependency guard](https://github.com/pgsty/pg_mooncake/blob/v0.3.3/src/sql/bootstrap.sql)
- [0.3.3 release](https://github.com/pgsty/pg_mooncake/releases/tag/v0.3.3)

`pg_mooncake` 0.3.3 通过逻辑复制，将 PostgreSQL 表维护为 Arrow 与 Parquet/Iceberg 存储中的只读列存镜像。此 PGSTY 版本支持 PostgreSQL 15–19，严格要求 `pg_duckdb` 1.2.1、DuckDB 2.0.1 与 Mooncake scanner 0.0.5。

### 配置与启用

管理员应保留已有预加载项，按以下顺序加入两个库并重启 PostgreSQL：

```ini
shared_preload_libraries = 'pg_duckdb,pg_mooncake'
wal_level = logical
duckdb.allow_unsigned_extensions = true
mooncake.extension_repository = '/usr/pgsql-17/share/extension/pg_mooncake'
```

实际路径应取目标服务器的 `pg_config --sharedir` 并追加 `/extension/pg_mooncake`；示例使用 EL 的 PostgreSQL 17 路径。本地 scanner 可在禁用社区扩展时加载。在新会话中以超级用户安装：

```sql
CREATE EXTENSION pg_mooncake CASCADE;
```

安装会检查依赖版本。须配置足够的复制槽、WAL sender 和 worker；源连接须具备认证、表读取以及复制和发布所需权限。

### 创建与检查镜像

源表 DDL 提交后，在可写会话中以独立命令创建镜像：

```sql
CREATE TABLE trades (id bigint PRIMARY KEY, symbol text, price real);
CALL mooncake.create_table('trades_analytics', 'trades');
INSERT INTO trades VALUES (1, 'AMZN', 207), (2, 'AMZN', 210);
SELECT symbol, avg(price) FROM trades_analytics GROUP BY symbol;
SELECT * FROM mooncake.list_tables();
CALL mooncake.create_snapshot('trades_analytics');
CALL mooncake.optimize_table('trades_analytics', 'full');
```

写入应发往源表，镜像不可写。远程源使用 `src_uri` 参数；新鲜读取与快照以源服务器的 WAL 位置同步。外部 Iceberg 读取者只看到已发布快照，本地查询还可读取缓冲数据。

### 持久性与升级

0.3.3 修复了快照可能在 Iceberg 持久化源提交前就确认完成的问题。从 0.3.1 或 0.3.2 升级时，安装匹配的库与 scanner 并重启，从已配置仓库刷新缓存中的旧 scanner，然后重新连接再更新：

```sql
SELECT duckdb.raw_query(format('FORCE INSTALL mooncake FROM %L', current_setting('mooncake.extension_repository')));
ALTER EXTENSION pg_mooncake UPDATE TO '0.3.3';
```

用于备份或迁移前，须为每个镜像重新创建快照。从基于 WAL 的 0.3.0 升级时，先停止源写入，在旧二进制下逐一比较镜像与源并验证快照，再干净停止 PostgreSQL；启动新 worker 前须完整备份存储，并离线保留旧 `_wal`。应遵循 tag 中逐库升级步骤；新快照无法补回已经丢失的旧 WAL。无法验证的镜像应从权威源重建。

### 存储与生命周期

采用物理备份或流复制备库时，应在创建镜像前将 `mooncake.data_directory` 放在 `PGDATA` 外，为每个实例使用私有存储并单独备份。逻辑转储不会保存 Moonlink 数据或复制状态；源连接字符串和元数据文件也须妥善保护。

重命名或删除含镜像的数据库前，应删除镜像并提交，再在 worker 可用时执行 `SELECT mooncake._reconcile_drops()`。初始化会为已有叶分区设置完整副本标识，后续分区变更须另行处理。应监控复制槽延迟与保留 WAL，并在丢弃源数据前验证清理结果。
