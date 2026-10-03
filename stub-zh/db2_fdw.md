## 用法

来源：

- [18.2.0 README](https://github.com/pg-fdw/db2_fdw/blob/18.2.0/README.md)
- [18.2.0 control file](https://github.com/pg-fdw/db2_fdw/blob/18.2.0/db2_fdw.control)
- [18.2.0 SQL API](https://github.com/pg-fdw/db2_fdw/blob/18.2.0/sql/db2_fdw--18.2.0.sql)

`db2_fdw` 18.2.0 通过 PostgreSQL 外部表查询和修改 IBM Db2 表，能够下推支持的过滤条件以及所需列。上游要求 PostgreSQL 10.1 或更高版本，以及与 PostgreSQL 架构相同的 IBM Db2 客户端 11.1 或更高版本。服务进程必须能够加载该客户端并访问指定数据库；安装扩展并不会提供外部客户端或 Db2 服务。

### 连接并导入表

示例假设 Db2 客户端可以连接已编目的 SAMPLE 数据库，且存在 DB2INST1.EMPLOYEE 表。由超级用户安装扩展，然后创建服务和指定角色的用户映射；应使用真实 Db2 凭据替换示例密码：

```sql
CREATE EXTENSION db2_fdw;
CREATE SERVER db2srv FOREIGN DATA WRAPPER db2_fdw
  OPTIONS (dbserver 'SAMPLE');
CREATE USER MAPPING FOR CURRENT_USER SERVER db2srv
  OPTIONS (user 'db2inst1', password 'change-me');
CREATE SCHEMA db2_remote;
IMPORT FOREIGN SCHEMA "DB2INST1" LIMIT TO ("EMPLOYEE")
  FROM SERVER db2srv INTO db2_remote;
SELECT empno, firstname, lastname, salary
FROM db2_remote.employee
WHERE empno = '000010';
```

导入会取得封装器所需的 Db2 列元数据；默认的智能大小写折叠会将全大写名称转为小写。独立的应用角色需要外部服务 USAGE 权限、自己的用户映射以及相应模式和表权限。凭据只供单个角色使用时，不应建立 PUBLIC 映射。用户名和密码均为空字符串时会选择上游的外部认证路径，结果取决于 Db2 客户端环境。

### 选项与写入

- 服务选项 `dbserver` 指定 Db2 连接，`no_encoding_error` 控制编码转换错误处理；本版本的 `batch_size` 为保留项，不能视为已实现批量插入。
- 表选项 `schema` 和 `table` 指定远端对象，`readonly` 禁止修改。`prefetch` 默认为 100，允许 0–1024；`fetch_size` 虽可设置，目前仍固定为 1。`sample_percent` 控制 ANALYZE 采样。
- 为进行 `UPDATE` 和 `DELETE`，列选项 `key` 必须标记全部远端主键列。导入的元数据包括 `db2type`、`db2size`、`db2bytes`、`db2chars`、`db2scale`、`db2null` 和 `db2ccsid`；修改导入表时应保留这些信息。
- IMPORT FOREIGN SCHEMA 的 `case` 和 `readonly` 选项控制名称折叠以及导入表能否写入。

INSERT、UPDATE 和 DELETE 还需要 Db2 侧的权限。开放写入前，应检查列映射和远端键定义；PostgreSQL 声明本身不会创建远端主键。不支持下推的过滤会在本地执行，应检查 EXPLAIN 后再判断是否已经下推。

### 诊断与事务

```sql
SELECT db2_diag();
SELECT db2_diag('db2srv');
SELECT db2_close_connections();
```

`db2_diag()` 返回本地客户端与构建信息，并可选地返回远端服务诊断；`db2_close_connections()` 关闭当前后端缓存的连接，不应在已经修改 Db2 数据的事务中调用。长时间运行的 PostgreSQL 会话可能持续占用远端连接和事务资源。

### 类型与维护

常见映射包括字符类型到文本或字符类型、BLOB 到 bytea、整数类型到 PostgreSQL 整数类型，以及 DATE、TIMESTAMP、TIME 到对应类型。声明的 PostgreSQL 类型与宽度必须容纳远端值，转换错误会在查询时出现。控制版本为 18.2.0，允许重定位；普通使用无需共享预加载。IBM 客户端的环境和连接要求应以对应版本 README 为准，并区分软件包升级与对实际外部数据库的验证。
