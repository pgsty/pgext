

## 用法

> [odbc_fdw: 通过 ODBC 访问远程数据库的外部数据包装器](https://github.com/CartoDB/odbc_fdw)

### 创建服务器

使用 ODBC 配置中定义的 DSN 连接：

```sql
CREATE EXTENSION odbc_fdw;

CREATE SERVER odbc_server
  FOREIGN DATA WRAPPER odbc_fdw
  OPTIONS (dsn 'test');
```

或者不使用 DSN，直接指定连接属性：

```sql
CREATE SERVER odbc_server
  FOREIGN DATA WRAPPER odbc_fdw
  OPTIONS (
    odbc_DRIVER 'MySQL',
    odbc_SERVER '192.168.1.17',
    encoding 'iso88591'
  );
```

**服务器选项：** `dsn`（ODBC 数据源名称）、`driver`（ODBC 驱动名称，无 DSN 时必填）、`odbc_*`（驱动特定属性）、`encoding`（远程数据库字符编码）。

驱动特定选项需加 `odbc_` 前缀。属性 DSN、DRIVER、UID 和 PWD 会自动转为大写。

### 创建用户映射

```sql
CREATE USER MAPPING FOR postgres
  SERVER odbc_server
  OPTIONS (odbc_UID 'root', odbc_PWD '');
```

### 创建外部表

```sql
CREATE FOREIGN TABLE odbc_table (
  id integer,
  name varchar(255),
  description text,
  users float4,
  created timestamp
)
SERVER odbc_server
OPTIONS (
  odbc_DATABASE 'mydb',
  schema 'test',
  sql_query 'SELECT id, name, description, created, users FROM test.mytable',
  sql_count 'SELECT count(id) FROM test.mytable'
);

SELECT * FROM odbc_table;
```

**表选项：** `schema`（远程模式）、`table`（远程表名）、`sql_query`（自定义 SQL 查询，覆盖 `table`）、`sql_count`（自定义计数 SQL）。

### 导入外部模式

```sql
IMPORT FOREIGN SCHEMA test
  FROM SERVER odbc_server
  INTO public
  OPTIONS (odbc_DATABASE 'mydb');
```
