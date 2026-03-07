

> [!WARNING] 此扩展目前存在问题，与 pg_duckdb 和 pg_mooncake 存在冲突

## 用法

### 创建扩展

安装 `duckdb_fdw` yum 软件包后，可以在 PostgreSQL 数据库中创建该扩展：

```sql
-- 创建扩展
CREATE EXTENSION duckdb_fdw;

-- 创建 duckdb_fdw 外部服务器
CREATE SERVER duckdb_server FOREIGN DATA WRAPPER duckdb_fdw OPTIONS (database '/tmp/duck.db');

-- 创建用户映射 [可选]
-- GRANT USAGE ON FOREIGN SERVER duckdb_server TO PUBLIC;

SELECT duckdb_fdw_version();

-- 可以使用 `duckdb_execute` 执行 duckdb 命令，例如在 duckdb 中创建一张表：
-- 在 duckdb 中创建表
SELECT duckdb_execute('duckdb_server', 'CREATE TABLE t1 (a integer,b varchar);');

-- 创建映射到 duckdb 表的外部表
CREATE FOREIGN TABLE t1 (
    a integer,
    b text
) SERVER duckdb_server OPTIONS (
    table 't1'
);

-- 写入数据并读取
INSERT INTO t1 SELECT i AS a,i::text AS b FROM generate_series(1,10) i;
SELECT * FROM t1;
```


也可以从 duckdb 服务器导入外部模式，例如，先使用 duckdb 命令行工具创建一张表：

```bash
duckdb /tmp/duck.db

CREATE TABLE t1 (
  a integer,
  b text
);

INSERT INTO t1 VALUES (1, 'a'), (2 , 'b'), (3, 'c');
SELECT * FROM t1;
```

然后将该模式导入 PostgreSQL：

```sql
IMPORT FOREIGN SCHEMA public FROM SERVER duckdb_server INTO public;
```

### 其他资源

- [DuckDB 官网](https://duckdb.org/)
- [GitHub: duckdb_fdw](https://github.com/alitrack/duckdb_fdw/)
- [构建 libduckdb](https://github.com/digoal/blog/blob/master/202401/20240124_01.md)
