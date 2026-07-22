## 用法

来源：

- [官方 README](https://github.com/2ndQuadrant/pglog/blob/5d57809e4e73913a06b4aaa6553f53f43d010a39/README.asciidoc)
- [官方控制文件](https://github.com/2ndQuadrant/pglog/blob/5d57809e4e73913a06b4aaa6553f53f43d010a39/pglog.control)
- [官方安装 SQL](https://github.com/2ndQuadrant/pglog/blob/5d57809e4e73913a06b4aaa6553f53f43d010a39/pglog--1.0.sql)
- [官方日志实现](https://github.com/2ndQuadrant/pglog/blob/5d57809e4e73913a06b4aaa6553f53f43d010a39/pglog.c)
- [官方 spool 读取器](https://github.com/2ndQuadrant/pglog/blob/5d57809e4e73913a06b4aaa6553f53f43d010a39/pglog_spool.c)

`pglog` 是 PostgreSQL 9.3 时代的概念验证项目：它通过错误钩子捕获消息，追加到自己的 spool 文件，再通过只读外部表公开这些文件。它与 PostgreSQL 的常规日志收集器相互独立，不能替代受支持的日志传输或可观测性管道。

### 核心流程

预加载动态库，选择数据库服务器本地的 spool 目录，并在安装 SQL 对象前重启 PostgreSQL。

```conf
shared_preload_libraries = 'pglog'
pglog.directory = 'pglog_spool'
pglog.min_messages = 'WARNING'
pglog.rotation_age = '1d'
```

```sql
CREATE EXTENSION pglog;

SELECT log_time, database_name, error_severity, message
FROM pglog
ORDER BY log_time DESC
LIMIT 100;
```

`pglog` 外部表采用类似 CSV 日志的结构，包含连接、语句、错误和应用字段。查询会扫描配置目录中的 spool 文件；该包装器既不下推过滤条件，也不收集外部表统计信息。

### 配置与对象

- `pglog.directory` 在未使用绝对路径时相对于 `PGDATA`，默认值为 `pglog_spool`。
- `pglog.min_messages` 设置捕获的最低严重级别，默认值为 `WARNING`。
- `pglog.rotation_age` 以分钟控制按时间轮转；`0` 表示禁用轮转。
- `pglog_server` 与 `pglog` 外部表由扩展创建。
- `pglog_severity` 是严重级别所使用的枚举。

### 运维边界

已复核的 `1.0` 源码没有在 PostgreSQL 后端之间串行化追加操作，因此并发消息可能交错或破坏记录。轮转只创建新文件，不删除旧文件。读取器最多按文件系统顺序打开 16 个目录项，也不会隔离当前数据库。

日志行可能包含来自多个数据库的语句、详情、用户名、主机数据和应用文本。应严格限制访问权限，并把 spool 目录视作敏感数据保护。该代码面向已经过时的 PostgreSQL API，且没有当前兼容性声明，因此只能在隔离测试构建中验证；运维场景应优先使用仍受维护的 PostgreSQL 日志与采集机制。
