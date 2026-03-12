

## 用法

> [db2_fdw: 访问 DB2 的外部数据包装器](https://github.com/wolfgangbrandl/db2_fdw)

### 创建服务器

```sql
CREATE EXTENSION db2_fdw;

CREATE SERVER db2srv FOREIGN DATA WRAPPER db2_fdw
  OPTIONS (dbserver 'SAMPLE');
```

**服务器选项：** `dbserver`（必填，DB2 数据库连接字符串）。

### 创建用户映射

```sql
CREATE USER MAPPING FOR PUBLIC SERVER db2srv
  OPTIONS (user 'db2inst1', password 'secret');
```

`user` 和 `password` 使用空字符串可启用外部认证。

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

**表选项：** `table`（必填，DB2 表名，区分大小写，通常大写）、`schema`（表所有者）、`readonly`（默认 `false`）、`prefetch`（每次往返获取行数，默认 200，范围 0-10240）、`max_long`（LONG 列最大长度，默认 32767）。

**列选项：** `key`（设为 `true` 标记主键列，UPDATE/DELETE 必需）。

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

### 数据类型映射

| DB2 类型 | PostgreSQL 类型 |
|----------|------------------|
| CHAR | char |
| VARCHAR | varchar |
| CLOB | text |
| BLOB | bytea |
| SMALLINT, INTEGER, BIGINT | smallint, integer, bigint |
| DOUBLE | numeric, float |
| DATE | date |
| TIMESTAMP | timestamp |
| TIME | time |

WHERE 条件和列投影会下推到 DB2 以最小化数据传输。
