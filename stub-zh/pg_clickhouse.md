
## 用法

> [pg_clickhouse: PostgreSQL 的 ClickHouse 集成](https://github.com/ClickHouse/pg_clickhouse)

`pg_clickhouse` 允许直接从 PostgreSQL 向 ClickHouse 执行分析查询，而无需重写 SQL。它支持 PostgreSQL 13+ 和 ClickHouse v23+。

## 入门

上游建议的起步方式主要有两种：

- 使用已发布的 Docker 镜像 `ghcr.io/clickhouse/pg_clickhouse:18`
- 通过 `make` / `make install` 从源码构建，或从 PGXN 安装

安装完成后启用扩展：

```sql
CREATE EXTENSION pg_clickhouse;
```

也可以安装到指定 schema：

```sql
CREATE SCHEMA ch;
CREATE EXTENSION pg_clickhouse WITH SCHEMA ch;
```

## 连接 ClickHouse

参考文档展示的标准流程如下：

```sql
CREATE SERVER taxi_srv
FOREIGN DATA WRAPPER clickhouse_fdw
OPTIONS (driver 'binary', host 'localhost', dbname 'taxi');

CREATE USER MAPPING FOR CURRENT_USER
SERVER taxi_srv
OPTIONS (user 'default');

CREATE SCHEMA taxi;
IMPORT FOREIGN SCHEMA taxi FROM SERVER taxi_srv INTO taxi;
```

文档中列出的服务器选项包括：

- `driver`，必填，可选 `binary` 或 `http`
- `dbname`
- `fetch_size`
- `host`
- `port`

## 文档重点

README 将 pg_clickhouse 的核心定位为面向分析工作负载的透明下推：

- 教程会带你把 PostgreSQL 连接到 ClickHouse 示例数据库，并查询导入后的表
- 参考文档会说明扩展生命周期命令、外部服务器选项以及扩展暴露的 SQL 对象

项目 README 还给出了 TPC-H 基准示例，说明在什么情况下查询下推能显著缩短耗时。

## 运行说明

参考文档把版本分成两层：

- 库版本，可通过 `pgch_version()` 或 `pg_get_loaded_modules()` 查看
- 扩展版本，由 PostgreSQL 系统目录和扩展升级脚本跟踪

小版本和大版本升级时，可能需要执行 `ALTER EXTENSION pg_clickhouse UPDATE`。
