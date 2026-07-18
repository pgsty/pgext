


## 用法

来源：[README](https://github.com/wolfgangbrandl/db2_fdw), [current upstream README](https://github.com/pg-fdw/db2_fdw/blob/master/README.md)

`db2_fdw` 是一个 PostgreSQL foreign data wrapper，可从 PostgreSQL 查询和修改 IBM Db2 表。它会尽可能下推所需列和 `WHERE` 条件，并提供连接清理和诊断辅助函数。

### 创建服务器

```sql
CREATE EXTENSION db2_fdw;

CREATE SERVER db2srv FOREIGN DATA WRAPPER db2_fdw
  OPTIONS (dbserver 'SAMPLE');
```

**服务器选项：** `dbserver`（必填的 Db2 连接字符串）、`batch_size`（当前保留给未来批处理行为）、`no_encoding_error`（`ON`、`OFF`、`YES`、`NO`、`TRUE` 或 `FALSE`）。

### 创建用户映射

```sql
CREATE USER MAPPING FOR PUBLIC SERVER db2srv
  OPTIONS (user 'db2inst1', password 'secret');
```

将 `user` 和 `password` 设为空字符串，可通过 Db2 客户端环境启用外部认证。

### 创建外部表

```sql
CREATE FOREIGN TABLE employee (
  empno char(6) OPTIONS (key 'true'),
  firstname varchar(12),
  lastname varchar(15),
  salary numeric
)
SERVER db2srv
OPTIONS (schema 'DB2INST1', table 'EMPLOYEE');
```

**表选项：** `table`（必填，Db2 表名或简单查询，区分大小写，通常为大写）、`schema`（表所有者）、`readonly`（默认 `false`）、`sample_percent`（`ANALYZE` 采样）、`prefetch`（每次往返获取的行数，默认 `100`，范围 `0`-`1024`）、`fetch_size`（可接受但当前固定为 `1`）、`batch_size`、`no_encoding_error`。上游文档将 `max_long` 标为 deprecated，且不再使用。

**列选项：** `key`（所有主键列都要设为 `true`，`UPDATE` 和 `DELETE` 必需），以及导入表上的 Db2 元数据选项，例如 `db2type`、`db2size`、`db2bytes`、`db2chars`、`db2scale`、`db2null`、`db2ccsid`。

### 导入外部模式

```sql
IMPORT FOREIGN SCHEMA "DB2INST1" FROM SERVER db2srv INTO public;
```

**导入选项：** `case`（`keep`、`lower` 或 `smart`，默认 `smart`）、`readonly`。

### CRUD 操作

```sql
SELECT * FROM employee WHERE empno = '000010';
INSERT INTO employee (empno, firstname, lastname, salary) VALUES ('999999', 'John', 'Doe', 50000);
UPDATE employee SET salary = 55000 WHERE empno = '999999';
DELETE FROM employee WHERE empno = '999999';
```

### 连接辅助函数

```sql
SELECT db2_close_connections();
SELECT db2_diag();
SELECT db2_diag('db2srv');
```

`db2_close_connections()` 会关闭当前会话中缓存的 Db2 连接。`db2_diag()` 会报告 db2_fdw、PostgreSQL、Db2 client，以及可选远端服务器的诊断详情。

### 数据类型映射

| Db2 类型 | PostgreSQL 类型 |
|----------|------------------|
| CHAR | char |
| VARCHAR | varchar |
| CLOB | text |
| VARGRAPHIC, GRAPHIC | text |
| BLOB | bytea |
| SMALLINT, INTEGER, BIGINT | smallint, integer, bigint |
| DOUBLE | numeric, float |
| DATE | date |
| TIMESTAMP | timestamp |
| TIME | time |

`WHERE` 条件和列投影会下推到 DB2，以减少数据传输。
