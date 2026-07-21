## 用法

来源：

- [odbc_fdw 0.6.1 README](https://github.com/devrimgunduz/odbc_fdw/blob/0.6.1/README.md)
- [odbc_fdw changelog](https://github.com/devrimgunduz/odbc_fdw/blob/0.6.1/NEWS.md)
- [Extension control file](https://github.com/devrimgunduz/odbc_fdw/blob/0.6.1/odbc_fdw.control)
- [0.6.0 to 0.6.1 comparison](https://github.com/devrimgunduz/odbc_fdw/compare/0.6.0...0.6.1)

`odbc_fdw` 将 ODBC 数据源中的表或驱动程序特定查询暴露为 PostgreSQL 外部表。它主要是一个跨异构系统的读取/查询桥梁；在依赖其进行生产查询之前，请验证数据类型转换和远程驱动行为。

### 核心工作流程

```sql
CREATE EXTENSION odbc_fdw;

-- In 0.6.1 a superuser must set the server-level dsn or driver option.
CREATE SERVER warehouse_odbc
  FOREIGN DATA WRAPPER odbc_fdw
  OPTIONS (dsn 'warehouse');

CREATE USER MAPPING FOR analyst
  SERVER warehouse_odbc
  OPTIONS (odbc_UID 'reporter', odbc_PWD 'secret');

CREATE FOREIGN TABLE remote_customer (
  id bigint,
  name text,
  created_at timestamp
) SERVER warehouse_odbc
  OPTIONS (schema 'sales', table 'customer');

SELECT * FROM remote_customer WHERE id = 42;
```

使用 `driver` 替代 `dsn` 进行无 DSN 连接。其他驱动程序属性使用 `odbc_` 前缀，并可放置在服务器、用户映射或外部表上。将凭据放在用户映射中。对大小写敏感的属性名称进行引号引用，必要时用花括号包裹包含 `=` 或 `;` 的值。

### 查询和导入

`sql_query` 覆盖 `table`；当 FDW 需要显式行数查询时，请与 `sql_count` 结合使用：

```sql
CREATE FOREIGN TABLE active_customer (
  id bigint,
  name text
) SERVER warehouse_odbc
  OPTIONS (
    sql_query 'SELECT id, name FROM sales.customer WHERE active = 1',
    sql_count 'SELECT count(*) FROM sales.customer WHERE active = 1'
  );

IMPORT FOREIGN SCHEMA sales
  FROM SERVER warehouse_odbc
  INTO imported
  OPTIONS (prefix 'odbc_');
```

### 重要对象和选项

- `dsn` 或 `driver` 选择 ODBC 数据源；0.6.1 版本限制这些服务器选项仅对超级用户可用，因为驱动程序管理器加载共享库。
- `schema`、`table`、`sql_query` 和 `sql_count` 选择远程关系或查询。
- `prefix` 改变由 `IMPORT FOREIGN SCHEMA` 创建的本地名称。
- `ODBCTablesList(server_name, ...)` 列出可见的远程表。
- `ODBCTableSize(server_name, table_name)` 和 `ODBCQuerySize(server_name, query)` 返回远程行数。

0.6.0 版本恢复了兼容性并修复了最近 PostgreSQL 发行版上的崩溃。0.6.1 版本转义了远程字面量和标识符以防止 SQL 注入，限制了驱动程序选择，并在调试连接字符串中屏蔽了常见的凭据属性。升级前请允许委派 FDW 使用，同时保留正常的服务器所有权和用户映射控制。

仅支持上游 README 中列出的 ODBC 类型。标识符长度、驱动程序 SQL 语法、编码、空值处理和二进制值可能会有所不同。源/包版本为 0.6.1，而控制文件和安装 SQL 继续声明扩展版本 0.5.2。
