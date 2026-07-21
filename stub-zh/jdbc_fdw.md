## 用法

来源：

- [jdbc_fdw v0.5.0 README](https://github.com/pgspider/jdbc_fdw/blob/v0.5.0/README.md)
- [Extension control file](https://github.com/pgspider/jdbc_fdw/blob/v0.5.0/jdbc_fdw.control)
- [jdbc_fdw v0.5.0 release](https://github.com/pgspider/jdbc_fdw/releases/tag/v0.5.0)

`jdbc_fdw` 将 JDBC 数据源暴露为 PostgreSQL 外部表，并可通过辅助函数执行远程 SQL。当存在合适的 JDBC 驱动但没有更专业的 FDW 时使用它；JVM、驱动 JAR、凭据和远程查询行为都在一个 PostgreSQL 后端进程中运行。

### 核心工作流程

```sql
CREATE EXTENSION jdbc_fdw;

CREATE SERVER reporting_jdbc
  FOREIGN DATA WRAPPER jdbc_fdw
  OPTIONS (
    drivername 'org.postgresql.Driver',
    url 'jdbc:postgresql://db.example/reporting',
    jarfile '/opt/jdbc/postgresql.jar',
    querytimeout '10',
    maxheapsize '256'
  );

CREATE USER MAPPING FOR app_user
  SERVER reporting_jdbc
  OPTIONS (username 'reader', password 'secret');

CREATE FOREIGN TABLE remote_orders (
  id bigint OPTIONS (key 'true'),
  created_at timestamptz,
  total numeric
) SERVER reporting_jdbc;

SELECT * FROM remote_orders WHERE id = 42;
```

v0.5.0 中没有表级选项。外部列按名称映射。当 `UPDATE` 或 `DELETE` 需要行标识时，请用 `OPTIONS (key 'true')` 标记远程主键列。

### 导入和直接 SQL

```sql
IMPORT FOREIGN SCHEMA public
  FROM SERVER reporting_jdbc
  INTO jdbc_import
  OPTIONS (recreate 'true');

SELECT *
FROM jdbc_exec('reporting_jdbc', 'SELECT id, name FROM customer')
  AS t(id bigint, name text);
```

上游 README 中提到，`IMPORT FOREIGN SCHEMA` 当前仅适用于 GridDB。`jdbc_exec` 返回 `record`，因此返回多列的查询需要定义列列表。

### 重要选项与限制

- 服务器选项：必需的 `drivername` 和 `url`，绝对路径的 `jarfile`，以及 `querytimeout` 和 JVM 的 `maxheapsize`。
- 用户映射选项：`username` 和 `password`。
- 列选项：`key = true` 标识可写操作中的主键列。
- `jdbc_exec(connname, sql)` 执行特定于驱动的 SQL 语句，并可以返回定义的记录集。

0.5.0 版本支持上游项目的谓词、列和聚合下推，但不支持远程 `RETURNING`、`GROUP BY`、`ORDER BY`、类型转换或事务控制语句。数组和外部表 `TRUNCATE` 未实现。请使用选定驱动程序测试数据类型的转换和写入语义。

保护 JAR 路径和服务器定义免受不可信用户的影响，将密码保存在用户映射中，并限制 JVM 堆大小和远程查询时间。源/包发行版为 0.5.0，而 `jdbc_fdw.control` 继续声明 SQL 扩展版本 1.2；请使用 `pg_extension.extversion` 而不是假设这些版本空间是相同的。
