

## 用法

> [jdbc_fdw: 通过 JDBC 访问远程服务器的外部数据包装器](https://github.com/pgspider/jdbc_fdw)

### 创建服务器

```sql
CREATE EXTENSION jdbc_fdw;

CREATE SERVER jdbc_server FOREIGN DATA WRAPPER jdbc_fdw
  OPTIONS (
    drivername 'org.postgresql.Driver',
    url 'jdbc:postgresql://remotehost:5432/mydb',
    jarfile '/usr/share/java/postgresql.jar',
    maxheapsize '256'
  );
```

**服务器选项：** `drivername`（必填，JDBC 驱动类）、`url`（必填，JDBC 连接 URL）、`jarfile`（必填，JDBC 驱动 JAR 的绝对路径）、`querytimeout`（查询超时秒数）、`maxheapsize`（JVM 堆大小 MB，最小 1）。

### 创建用户映射

```sql
CREATE USER MAPPING FOR CURRENT_USER SERVER jdbc_server
  OPTIONS (username 'dbuser', password 'dbpass');
```

### 创建外部表

```sql
CREATE FOREIGN TABLE remote_table (
  id integer OPTIONS (key 'true'),
  name text,
  value numeric
)
SERVER jdbc_server
OPTIONS (table_name 'schema.tablename');
```

在主键列上设置 `key 'true'` 以启用 UPDATE 和 DELETE 操作。

### 查询远程数据

```sql
SELECT * FROM remote_table WHERE id > 100;
```

### 使用 jdbc_exec 执行任意 SQL

`jdbc_exec` 函数在远程数据库上执行 SQL 并返回结果集：

```sql
SELECT * FROM jdbc_exec('jdbc_server', 'SELECT id, name FROM remote_schema.remote_table WHERE status = 1')
  AS t(id integer, name text);
```

这对于执行超出外部表定义范围的查询很有用，包括在远程服务器上执行 DDL 或复杂查询。
