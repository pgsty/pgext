

## 用法

> [sqlite_fdw: SQLite 外部数据包装器](https://github.com/pgspider/sqlite_fdw)

### 创建服务器

```sql
CREATE EXTENSION sqlite_fdw;

CREATE SERVER sqlite_server FOREIGN DATA WRAPPER sqlite_fdw
  OPTIONS (database '/path/to/database.db');
```

**服务器选项：** `database`（必填，SQLite 文件路径）、`updatable`（默认 `true`）、`truncatable`（默认 `false`）、`keep_connections`（默认 `true`）、`batch_size`（默认 1）、`force_readonly`（默认 `false`）。

由于 SQLite 没有认证模型，无需 `CREATE USER MAPPING`。

### 创建外部表

```sql
CREATE FOREIGN TABLE remote_data (
  id integer OPTIONS (key 'true'),
  name text,
  created timestamp OPTIONS (column_type 'INT'),
  data bytea
)
SERVER sqlite_server
OPTIONS (table 'data_table');
```

**表选项：** `table`（SQLite 表名，如与 PostgreSQL 名称不同）、`updatable`、`truncatable`、`batch_size`。

**列选项：** `column_name`（映射到不同的 SQLite 列名）、`column_type`（SQLite 亲和类型：`INT` 用于时间戳纪元，`BLOB` 用于 UUID）、`key`（标记为主键，UPDATE/DELETE 必需）。

### CRUD 操作

```sql
SELECT * FROM remote_data WHERE id > 100;
INSERT INTO remote_data (id, name) VALUES (1, 'test');
UPDATE remote_data SET name = 'updated' WHERE id = 1;
DELETE FROM remote_data WHERE id = 1;
```

### 导入外部模式

```sql
IMPORT FOREIGN SCHEMA public
  FROM SERVER sqlite_server INTO local_schema;
```

**导入选项：** `import_default`（默认 `false`）、`import_not_null`（默认 `true`）。

### 数据类型映射

| SQLite 类型 | PostgreSQL 类型 |
|-------------|-----------------|
| int | bigint |
| text, char, clob | text |
| blob | bytea |
| real, float, double | double precision |
| datetime | timestamp |
| uuid | uuid |
| json, jsonb | json, jsonb |

时间戳可以存储为 TEXT（ISO 格式）或 INT（Unix 纪元，使用 `column_type 'INT'`）。UUID 可以存储为 TEXT（36 字符）或 BLOB（16 字节）。SQLite 数据库文件必须对 PostgreSQL 操作系统用户可读（DML 操作还需可写）。
