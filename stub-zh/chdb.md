## 用法

来源：

- [PGXN 0.1.1 README](https://pgxn.org/dist/chdb/0.1.1/README.html)
- [chdb 0.1 文档](https://pgxn.org/dist/chdb/0.1.1/doc/chdb.html)
- [chdb 控制文件](https://api.pgxn.org/src/chdb/chdb-0.1.1/chdb.control)
- [Apache 2.0 许可证](https://api.pgxn.org/src/chdb/chdb-0.1.1/LICENSE.md)
- [0.1.1 changelog](https://api.pgxn.org/src/chdb/chdb-0.1.1/CHANGELOG.md)

`chdb` 通过隔离的辅助进程从 PostgreSQL 执行 chDB 查询。若能显式声明返回行结构，可用它进行偶发的 ClickHouse 格式分析或读取 chDB 表函数；它不是持久化的嵌入式 ClickHouse 数据库。

### 核心流程

```sql
CREATE EXTENSION chdb;

SELECT *
FROM chdb_query('SELECT number, number * number FROM numbers(5)')
AS result(number bigint, square bigint);
```

`chdb_query(text)` 返回 `record`，因此每次调用都需要与 chDB 结果匹配的 `AS` 列定义列表。默认不向任何角色授予 `EXECUTE`；管理员必须显式授权该函数。

```sql
GRANT EXECUTE ON FUNCTION chdb_query(text) TO analytics_role;
SELECT pgchdb_version();
```

### 资源控制

- `chdb.max_memory` 映射到 chDB 最大内存设置；`0` 表示不限制。
- `chdb.max_threads` 限制查询处理线程。
- `chdb.max_parsing_threads` 限制受支持输入格式的并行解析。

这些设置需要超级用户权限。向可能执行大扫描的工作负载开放 `chdb_query(text)` 前应设置有限值，因为不受约束的辅助进程会与 PostgreSQL 竞争内存、CPU、磁盘和网络带宽。

### 架构与边界

每次调用都会启动由辅助进程承载的临时 chDB 数据库，并在查询结束后删除。一个 `chdb_query(text)` 调用创建的对象无法被下一次调用看到。辅助进程崩溃会让发起请求的后端失败，但不会共享 PostgreSQL 内存；它仍可能消耗主机资源并访问 chDB 允许的数据源。

0.1.1 要求 PostgreSQL 15 或以上，以及 Linux 或 macOS 上的 libchdb 26.7.0 或以上。控制文件要求超级用户、不受信任、可重定位；扩展版本为 `0.1`，库/发行包版本为 `0.1.1`。应谨慎审查远程凭据与 URL：查询文本能够读取 chDB 支持的网络或本地数据源，结果还会经过显式类型转换边界进入 PostgreSQL。

### 0.1.1 版本变化

这是二进制更新，`chdb` 控制版本仍为 `0.1`，无需执行扩展 SQL 升级。该版增加 PostgreSQL 15 支持，大型无符号整数及 128/256 位整数映射为 `numeric`，并增加 BFloat16 和时间间隔映射。字符串转换会验证数据库编码，时间戳文本导出使用 ISO-8601 并按文档处理时区。替换并重新加载二进制后应验证代表性数据的往返转换。
