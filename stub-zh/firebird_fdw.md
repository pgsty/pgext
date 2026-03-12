

## 用法

> [firebird_fdw: Firebird 外部数据包装器](https://github.com/ibarwick/firebird_fdw)

### 创建服务器

```sql
CREATE EXTENSION firebird_fdw;

CREATE SERVER firebird_server FOREIGN DATA WRAPPER firebird_fdw
  OPTIONS (address 'localhost', database '/path/to/database.fdb');
```

**服务器选项：** `address`（默认 `localhost`）、`port`（默认 `3050`）、`database`（必填，Firebird 数据库文件路径）、`updatable`（默认 `true`）、`disable_pushdowns`（禁用 WHERE 子句下推）、`quote_identifiers`、`implicit_bool_type`（启用整数到布尔值的转换）、`batch_size`（PostgreSQL 14+）。

### 创建用户映射

```sql
CREATE USER MAPPING FOR CURRENT_USER SERVER firebird_server
  OPTIONS (username 'sysdba', password 'masterke');
```

### 创建外部表

```sql
CREATE FOREIGN TABLE fb_test (
  id smallint,
  val varchar(2048)
)
SERVER firebird_server
OPTIONS (table_name 'fdw_test');
```

列名映射：

```sql
CREATE FOREIGN TABLE fb_mapped (
  id smallint OPTIONS (column_name 'test_id'),
  val varchar(2048) OPTIONS (column_name 'test_val')
)
SERVER firebird_server
OPTIONS (table_name 'fdw_test');
```

自定义查询（只读）：

```sql
CREATE FOREIGN TABLE fb_query (
  id smallint,
  val varchar(2048)
)
SERVER firebird_server
OPTIONS (query $$ SELECT id, val FROM fdw_test WHERE id > 10 $$);
```

**表选项：** `table_name`、`query`（与 `table_name` 互斥，只读）、`updatable`、`estimated_row_count`、`quote_identifier`、`batch_size`。

**列选项：** `column_name`、`quote_identifier`、`implicit_bool_type`。

### 导入外部模式

```sql
IMPORT FOREIGN SCHEMA someschema
  FROM SERVER firebird_server
  INTO public
  OPTIONS (import_views 'true', verbose 'true');
```

**导入选项：** `import_not_null`（默认 `true`）、`import_views`（默认 `true`）、`updatable`、`verbose`。

schema 参数在 Firebird 中没有特殊含义，可以设置为任意值。

### CRUD 操作

```sql
SELECT * FROM fb_test WHERE id > 5;
INSERT INTO fb_test VALUES (10, 'new record');
UPDATE fb_test SET val = 'updated' WHERE id = 10;
DELETE FROM fb_test WHERE id = 10;
```
