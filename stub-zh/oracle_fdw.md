

## 用法

> [oracle_fdw: 访问 Oracle 的外部数据包装器](https://github.com/laurenz/oracle_fdw)

### 创建服务器

```sql
CREATE EXTENSION oracle_fdw;

CREATE SERVER oradb FOREIGN DATA WRAPPER oracle_fdw
  OPTIONS (dbserver '//dbserver.mydomain.com:1521/ORADB');
```

**服务器选项：** `dbserver`（必填，Oracle 连接字符串）、`isolation_level`（`serializable`、`read_committed` 或 `read_only`，默认 `serializable`）、`nchar`（昂贵的字符转换，默认 `off`）、`set_timezone`（与 Oracle 同步时区，默认 `off`）。

### 创建用户映射

```sql
CREATE USER MAPPING FOR pguser SERVER oradb
  OPTIONS (user 'orauser', password 'orapwd');
```

`user` 使用空字符串可启用外部（OS）认证。

### 创建外部表

```sql
CREATE FOREIGN TABLE oratab (
  id integer OPTIONS (key 'true') NOT NULL,
  text character varying(30),
  floating double precision NOT NULL
)
SERVER oradb
OPTIONS (schema 'ORAUSER', table 'ORATAB');
```

**表选项：** `table`（必填，Oracle 表名，大写）、`schema`（表所有者）、`dblink`（Oracle DB link）、`max_long`（LONG 列最大长度，默认 32767）、`readonly`（默认 `false`）、`sample_percent`（ANALYZE 采样率，默认 100）、`prefetch`（每次往返获取行数，默认 50）。

**列选项：** `key`（标记为主键，UPDATE/DELETE 必需）、`strip_zeros`（从字符串中移除 ASCII 0）。

也可以用查询代替表名，将其放在括号中：

```sql
CREATE FOREIGN TABLE oraquery (
  id integer,
  text character varying(30)
)
SERVER oradb
OPTIONS (table '(SELECT id, text FROM ORAUSER.ORATAB WHERE id > 10)');
```

### 导入外部模式

```sql
IMPORT FOREIGN SCHEMA "ORAUSER"
  FROM SERVER oradb INTO local_schema;
```

**导入选项：** `case`（`keep`、`lower` 或 `smart`，默认 `smart`）、`readonly`、`skip_tables`、`skip_views`、`skip_matviews`、`max_long`、`sample_percent`、`prefetch`。

### 工具函数

```sql
SELECT oracle_diag();              -- 显示版本和环境信息
SELECT oracle_diag('oradb');       -- 包含 Oracle 服务器版本
SELECT oracle_close_connections(); -- 关闭所有缓存的 Oracle 连接
SELECT oracle_execute('oradb', 'TRUNCATE TABLE ORAUSER.ORATAB'); -- 执行任意 Oracle SQL
```

### 数据类型映射

| Oracle 类型 | PostgreSQL 类型 |
|---|---|
| CHAR, VARCHAR2, NVARCHAR2 | char, varchar, text |
| CLOB, NCLOB | text, json |
| NUMBER | numeric, float4, float8, int2, int4, int8, boolean |
| DATE, TIMESTAMP | date, timestamp, timestamptz |
| BLOB, LONG RAW | bytea |
| XMLTYPE | xml, text |
| SDO_GEOMETRY | geometry (PostGIS) |
