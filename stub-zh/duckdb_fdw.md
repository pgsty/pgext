## 用法

来源：

- [2.0.2 README](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/README.md)
- [2.0.2 控制文件](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/duckdb_fdw.control)
- [2.0.2 SQL 函数与权限](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/duckdb_fdw--2.0.2.sql)
- [服务器、映射与表选项](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/option.c)
- [扫描、写入与诊断函数](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/duckdb_fdw.c)
- [连接与事务处理](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/connection.c)
- [模式导入实现](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/import.c)
- [2.0.1 至 2.0.2 升级脚本](https://github.com/alitrack/duckdb_fdw/blob/4c444d19f80eace50bc09fa3ec766e8ca2cf667d/duckdb_fdw--2.0.1--2.0.2.sql)

`duckdb_fdw` 2.0.2 通过 DuckDB 原生 C API，将 DuckDB 表和数据源映射为 PostgreSQL 外部表。可用于查询 DuckDB 数据库、读取 Parquet，或访问可选的 DuckDB 目录。外部表支持读取和插入，不支持通过 FDW 执行 PostgreSQL UPDATE 和 DELETE。

### 启用并映射 DuckDB 表

以超级用户创建扩展。全新安装无需预加载，也无需重启 PostgreSQL。持久化 DuckDB 文件应放在 PostgreSQL 服务账户可访问的目录中；以下示例假定目录已经存在。

```sql
CREATE EXTENSION duckdb_fdw;
CREATE SERVER duckdb_server
  FOREIGN DATA WRAPPER duckdb_fdw
  OPTIONS (database '/var/lib/postgresql/duckdb/analytics.duckdb');

SELECT duckdb_execute(
  'duckdb_server',
  'CREATE TABLE measurements (sensor_id INTEGER, reading DOUBLE)'
);

CREATE FOREIGN TABLE measurements (
  sensor_id integer,
  reading double precision
) SERVER duckdb_server OPTIONS (table 'measurements');

INSERT INTO measurements VALUES (1, 21.5), (2, 19.8);
SELECT * FROM measurements WHERE reading > 20;

SELECT extversion FROM pg_extension WHERE extname = 'duckdb_fdw';
SELECT duckdb_fdw_version();
```

系统目录查询返回扩展版本；`duckdb_fdw_version()` 虽然使用这个名称，实际返回的是所链接的 DuckDB 引擎版本。无需凭据的本地文件可以不创建用户映射。数据库文件访问使用 PostgreSQL 操作系统账户的权限。

### 导入表与读取 Parquet

```sql
CREATE SCHEMA duckdb_import;
IMPORT FOREIGN SCHEMA main
  FROM SERVER duckdb_server INTO duckdb_import;
```

这里应使用 DuckDB 的模式名，通常是 `main`，而非 PostgreSQL 的 `public`。读取已有 Parquet 文件时，应声明与文件匹配的 PostgreSQL 列类型：

```sql
CREATE FOREIGN TABLE parquet_measurements (
  sensor_id integer,
  reading double precision
) SERVER duckdb_server
  OPTIONS (table '/var/lib/postgresql/duckdb/measurements.parquet');

SELECT * FROM parquet_measurements WHERE sensor_id = 1;
```

### 选项与权限

| 对象 | 用途 |
| --- | --- |
| 服务器 `database` | DuckDB 文件路径，或仅在当前连接内存在的内存数据库 |
| 服务器 `force_readonly` | 以只读方式打开文件库，并拒绝 FDW 写入及辅助函数中的非只读语句 |
| 服务器 `extensions`、`attach_catalogs` | 加载可选 DuckDB 扩展并附加目录 |
| 表 `table` | 映射 DuckDB 表或 Parquet 路径 |
| 列 `column_name` | 将 PostgreSQL 列映射到不同名称的 DuckDB 列 |
| `duckdb_execute(server, statement)` | 执行 DuckDB SQL，返回 void，不返回结果表 |
| `duckdb_create_s3_secret(...)` | 在 DuckDB 连接中创建 S3 密钥对象 |
| `duckdb_fdw_preflight()` | 使用 DuckDB 前检查运行时兼容性 |

只向应用角色授予必需的外部服务器、模式和外部表权限。安装时会撤销 PUBLIC 对任意 SQL 和密钥辅助函数的 EXECUTE 权限；仅应向可信角色授予这些权限。

S3、MotherDuck 和 Quack 凭据应优先存入用户映射。服务器选项可从 PostgreSQL 系统目录中读取，不适合存放秘密。远程数据源还依赖相应 DuckDB 扩展、网络访问和凭据，安装 FDW 本身不代表这些条件已经满足。

### 连接与事务边界

- 默认在事务结束时关闭连接。使用 `database ':memory:'` 时，应在同一个显式事务中完成初始化和查询，否则临时数据库会消失。
- `duckdb_fdw.keep_connection` 默认关闭。开启后会在提交后保留连接，但可写文件连接也会跨事务持有 DuckDB 文件锁，可能阻止其他 PostgreSQL 会话打开该文件。
- 外部表插入跟随最外层 PostgreSQL 事务提交或回滚。保存点回滚不会单独撤销 DuckDB 写入：如果外层事务提交，已中止子事务中的写入仍可能保留。
- DuckDB 不参与 PostgreSQL 两阶段提交。不要使用预备事务，也不要假定 PostgreSQL 与 DuckDB 能在崩溃时保证原子提交。经辅助函数发送的任意 SQL 也不会自动获得 PostgreSQL 的事务语义。

### 兼容性与升级

Pigsty 2.0.2 软件包面向 PostgreSQL 15-18。上游 README 声称支持 PostgreSQL 13-18，范围大于这批软件包。2.0.2 接受 DuckDB 1.5.x 和 2.0.x 运行时，相比 2.0.1 不改变 SQL 函数集合。

运行时保护会阻止在同一后端中加载 `pg_duckdb` 后执行 DuckDB 操作。应将两个扩展放在不同 PostgreSQL 实例中。从预加载 `pg_duckdb` 的配置切换时，需移除其预加载项并重启 PostgreSQL；替换文件不会卸载已有进程中的库。

```sql
SELECT duckdb_fdw_runtime_compatibility_status();
SELECT duckdb_fdw_runtime_fingerprint();
SELECT duckdb_fdw_preflight();
```

安装匹配文件后，使用新的后端会话，并逐库更新：

```sql
ALTER EXTENSION duckdb_fdw UPDATE TO '2.0.2';
```
