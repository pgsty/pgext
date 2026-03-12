

## 用法

> [babelfishpg_tsql: SQL Server Transact SQL 兼容性](https://babelfishpg.org/)

`babelfishpg_tsql` 扩展作为 Babelfish 项目的一部分，为 PostgreSQL 提供 Microsoft SQL Server Transact-SQL（T-SQL）兼容性。为 SQL Server 编写的应用程序可以连接到 PostgreSQL 并运行查询，只需最小的更改。

### 启用

```sql
CREATE EXTENSION babelfishpg_tsql;
```

### 关键特性

- **T-SQL 语言支持**：理解 T-SQL 语法，包括存储过程、函数、触发器和批处理
- **SQL Server 线协议**：应用程序可以使用 TDS（表格数据流）协议在 1433 端口连接
- **系统过程**：常用的 `sp_` 系统存储过程可用
- **系统视图**：SQL Server 目录视图（如 `sys.tables`、`sys.columns`、`sys.objects`）
- **多数据库语义**：支持 SQL Server 风格的数据库/模式分离

### 支持的 T-SQL 特性

- `BEGIN...END` 块、`IF...ELSE`、`WHILE` 循环
- `TRY...CATCH` 错误处理
- 临时表（`#temp`、`##global_temp`）
- 表变量（`DECLARE @t TABLE (...)`）
- `IDENTITY` 列和 `@@IDENTITY` / `SCOPE_IDENTITY()`
- `TOP` 子句、`OUTPUT` 子句
- `MERGE` 语句
- 公共表表达式（CTE）
- 同一实例内的跨数据库查询
- `EXEC` / `EXECUTE` 动态 SQL
- SQL Server 风格字符串连接和 NULL 处理
- `PRINT` 和 `RAISERROR` 语句

### 通过 TDS 协议连接

应用程序可以使用 SQL Server 驱动程序（JDBC、ODBC、ADO.NET）连接到 TDS 监听端口（默认 1433）：

```
Server: hostname
Port: 1433
Database: mydb
```

### 注意事项

- 需要 `babelfishpg_common` 扩展
- Babelfish for PostgreSQL 项目的一部分（Apache 2.0 / PostgreSQL 许可证）
- 并非所有 T-SQL 特性都受支持；请查看 Babelfish 兼容性参考
