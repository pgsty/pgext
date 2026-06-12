## 用法

来源：[upstream README](https://github.com/ibarwick/firebird_fdw)、[1.4.2 README](https://github.com/ibarwick/firebird_fdw/blob/REL_1_4_STABLE/README.md)、[1.4.2 tag](https://github.com/ibarwick/firebird_fdw/tree/1.4.2)。

`firebird_fdw` 通过外部数据包装器 API 将 PostgreSQL 连接到 Firebird 数据库。它支持读取、写入、`IMPORT FOREIGN SCHEMA`、受支持表达式的谓词下推、连接缓存，以及 PostgreSQL 14+ 上的外部表 `TRUNCATE`。

### 创建服务器

```sql
CREATE EXTENSION firebird_fdw;

CREATE SERVER firebird_server FOREIGN DATA WRAPPER firebird_fdw
  OPTIONS (address 'localhost', database '/path/to/database.fdb');
```

服务器选项包括：

- `address`，默认 `localhost`。
- `port`，默认 `3050`。
- `database`，Firebird 数据库名称或路径。
- `updatable`，默认 `true`；表级设置可以覆盖它。
- `disable_pushdowns`，用于调试或基准测试。
- `quote_identifiers`，默认引用表名和列名标识符。
- `implicit_bool_type`，用于由整数承载的 Firebird 布尔列。
- `batch_size`，用于 PostgreSQL 14+ 上的多行插入。

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

使用列名映射：

```sql
CREATE FOREIGN TABLE fb_mapped (
  id smallint OPTIONS (column_name 'test_id'),
  val varchar(2048) OPTIONS (column_name 'test_val')
)
SERVER firebird_server
OPTIONS (table_name 'fdw_test');
```

使用自定义查询时，外部表是只读的：

```sql
CREATE FOREIGN TABLE fb_query (
  id smallint,
  val varchar(2048)
)
SERVER firebird_server
OPTIONS (query $$ SELECT id, val FROM fdw_test WHERE id > 10 $$);
```

表选项包括 `table_name`、`query`、`updatable`、`estimated_row_count`、`quote_identifier` 和 `batch_size`。列选项包括 `column_name`、`quote_identifier` 和 `implicit_bool_type`。

### 导入外部模式

```sql
IMPORT FOREIGN SCHEMA someschema
  FROM SERVER firebird_server
  INTO public
  OPTIONS (import_views 'true', verbose 'true');
```

导入选项包括 `import_not_null`、`import_views`、`updatable` 和 `verbose`。模式名在 Firebird 中没有特殊含义，可以是 PostgreSQL `IMPORT FOREIGN SCHEMA` 语法接受的任意值。

### CRUD 操作

```sql
SELECT * FROM fb_test WHERE id > 5;
INSERT INTO fb_test VALUES (10, 'new record');
UPDATE fb_test SET val = 'updated' WHERE id = 10;
DELETE FROM fb_test WHERE id = 10;
TRUNCATE fb_test;
```

`UPDATE` 和 `DELETE` 使用 Firebird 的 `RDB$DB_KEY`。由于 Firebird 没有原生 `TRUNCATE` 语句，`TRUNCATE` 会实现为不带条件的 Firebird `DELETE`。

### 工具函数

- `firebird_fdw_version()` 返回 FDW 版本整数。
- `firebird_fdw_close_connections()` 关闭当前 PostgreSQL 会话中缓存的 Firebird 连接。
- `firebird_fdw_server_options(servername text)` 显示有效服务器选项值，以及每个选项是否由用户显式提供。
- `firebird_fdw_diag()` 返回诊断键值数据，包括 FDW 和 `libfq` 版本。
- `firebird_version()` 报告已配置外部服务器的 Firebird 服务端版本；为此它可能会打开连接。

### 注意事项

- Pigsty 为 PostgreSQL 14-18 打包 `firebird_fdw` 1.4.2。上游文档声明兼容 PostgreSQL 10-19，其中较新的 FDW API 功能只在较新的 PostgreSQL 版本中可用。
- 上游支持 Firebird 2.5 及更高版本；更旧的 Firebird 版本不是已测试目标。
- `firebird_fdw` 与 `libfq` 共同开发，因此包兼容性取决于这些库是否匹配。
- 基于自定义查询的外部表不能更新。
- `implicit_bool_type` 功能仍属实验性质，面向使用整数列表示布尔值的场景。
