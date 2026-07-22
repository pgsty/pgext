## 用法

来源：

- [官方扩展控制文件](https://github.com/bigsql/pgtsql/blob/3240980dff20736a38101eb0806c8580ff169d9c/pgtsql.control)
- [官方 README](https://github.com/bigsql/pgtsql/blob/3240980dff20736a38101eb0806c8580ff169d9c/README.md)
- [官方扩展 SQL](https://github.com/bigsql/pgtsql/blob/3240980dff20736a38101eb0806c8580ff169d9c/sql/pgtsql--3.0.sql)
- [官方回归示例](https://github.com/bigsql/pgtsql/blob/3240980dff20736a38101eb0806c8580ff169d9c/test/sql/pltsql_at_prefixed_vars.sql)

`pgtsql` 3.0 安装 `pltsql` 过程语言，以及一个面向 PostgreSQL 函数和过程的小型 Transact-SQL 风格兼容层。它适合移植使用 `@` 变量、T-SQL 风格控制流、`PRINT` 和本地 `#` 临时表的有限过程代码；它不是完整的 SQL Server 兼容层或线协议端点。

### 核心流程

安装 `pgtsql` 需要超级用户，因为它会注册 C 语言处理器和校验器。最小函数可以使用 T-SQL 风格的变量声明和返回语法：

```sql
CREATE EXTENSION pgtsql;

CREATE FUNCTION tsql_value() RETURNS integer AS $$
DECLARE @value int = 120
BEGIN
    RETURN @value;
END
$$ LANGUAGE pltsql;

SELECT tsql_value();
SELECT sys.getdate(), sys.isnull(NULL::integer, 0);
```

3.0 版专门面向 PostgreSQL 11，并新增了存储过程支持。尝试在更新的 PostgreSQL 版本上使用前，应固定该源码并运行其回归测试。

### 已安装接口

- 语言：`pltsql`，由 `pltsql_call_handler` 和 `pltsql_validator` 支持。
- 日期/时间辅助函数：`sys.sysdatetime`、`sys.sysdatetimeoffset`、`sys.sysutcdatetime`、`sys.getdate` 和 `sys.getutcdate`。
- 空值辅助函数：针对有限标量类型重载的 `sys.isnull`。
- 类目录视图：`sys.tables`、`sys.views` 和 `sys.objects` 以类似 SQL Server 的列展示部分 PostgreSQL 关系。

解析器支持上游文档和测试中的部分约定，包括可选分号、带 `@` 前缀的变量、T-SQL 风格 `IF`/`ELSE`、`PRINT` 以及 `#` 临时表重写。必须逐项测试迁移语句；不受支持的 T-SQL 语法和语义不会自动模拟。

### 兼容性与安全边界

安装过程会无条件创建 `sys` schema，因此可能与现有 schema 冲突。扩展向 `PUBLIC` 授予 `sys` 的使用权限、兼容视图的查询权限以及辅助函数的执行权限；安装后应复核这些授权。

兼容视图中的许多 SQL Server 字段只是占位值，例如 epoch 时间戳和常量标志，并非等价的服务器元数据。项目不提供 SQL Server 认证、TDS 连接、完整内置函数覆盖、SQL Server 事务行为或完整系统目录。应把 `pgtsql` 视为范围有限的 PostgreSQL 过程语言移植辅助工具，并独立验证应用行为。
