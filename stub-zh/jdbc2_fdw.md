## 用法

来源：

- [官方上游文档](https://github.com/heimir-sverrisson/jdbc2_fdw/blob/758caf50c674577144b644f36432c4736ec99627/README.md)
- [官方选项实现](https://github.com/heimir-sverrisson/jdbc2_fdw/blob/758caf50c674577144b644f36432c4736ec99627/option.c)
- [官方扩展 SQL](https://github.com/heimir-sverrisson/jdbc2_fdw/blob/758caf50c674577144b644f36432c4736ec99627/jdbc2_fdw--1.0.sql)

`jdbc2_fdw` 是通过 JDBC 查询数据库的实验性外部数据包装器。它在 PostgreSQL 后端进程中嵌入 Java 虚拟机，并使用基于 PostgreSQL 9.4 时代 `postgres_fdw` 的反解析器，将符合条件的过滤条件下推到远端数据库。公开项目只说明了只读查询；不要依赖远端 `INSERT`、`UPDATE` 或 `DELETE` 支持。

### 核心流程

安装扩展，在外部服务器上定义 JDBC 驱动和连接，将凭据保存在用户映射中，再映射远端表：

```sql
CREATE EXTENSION jdbc2_fdw;

CREATE SERVER jdbc_remote
  FOREIGN DATA WRAPPER jdbc2_fdw
  OPTIONS (
    drivername 'org.postgresql.Driver',
    url 'jdbc:postgresql://db.example.com:5432/app',
    jarfile '/opt/jdbc/postgresql.jar',
    querytimeout '30'
  );

CREATE USER MAPPING FOR app_user
  SERVER jdbc_remote
  OPTIONS (username 'remote_user', password 'secret');

CREATE FOREIGN TABLE public.remote_orders (
  id bigint,
  status text
)
SERVER jdbc_remote
OPTIONS (schema_name 'public', table_name 'orders');

SELECT * FROM public.remote_orders WHERE status = 'open';
```

所有使用该服务器的 PostgreSQL 后端都必须能够读取并加载 JDBC 驱动 JAR 与嵌入式 JVM 库。

### 选项

服务器选项包括 `drivername`、`url`、`querytimeout`、`jarfile`、`maxheapsize`、`use_remote_estimate`、`fdw_startup_cost`、`fdw_tuple_cost` 和 `updatable`。用户映射接受 `username` 与 `password`。外部表接受 `schema_name`、`table_name`、`updatable` 和 `use_remote_estimate`；列可以用 `column_name` 覆盖远端名称。

### 兼容性与安全

代码基于较早的 PostgreSQL 和 JNI 接口。使用前应验证准确的 PostgreSQL、JDK、体系结构与 JDBC 驱动组合；扩展安装成功并不能证明后端可以创建 JVM。密码会存储在 PostgreSQL 元数据中，因此应限制目录可见性和用户映射管理权限，在 JDBC URL 中要求 TLS，并设置有界的查询超时和 JVM 堆。测试类型转换、空值、条件下推、取消、连接失败及后端内存占用。除非选定分支明确记录并测试了其他能力，否则应将它视为只读实验性集成。
