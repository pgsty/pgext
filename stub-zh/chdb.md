## 用法

来源：

- [doc/chdb.md](https://api.pgxn.org/src/chdb/chdb-0.1.4/doc/chdb.md)
- [doc/chdb_hook.md](https://api.pgxn.org/src/chdb/chdb-0.1.4/doc/chdb_hook.md)
- [chdb.control](https://api.pgxn.org/src/chdb/chdb-0.1.4/chdb.control)
- [sql/chdb.sql](https://api.pgxn.org/src/chdb/chdb-0.1.4/sql/chdb.sql)
- [CHANGELOG.md](https://api.pgxn.org/src/chdb/chdb-0.1.4/CHANGELOG.md)
- [LICENSE.md](https://api.pgxn.org/src/chdb/chdb-0.1.4/LICENSE.md)

`chdb` 在隔离的辅助进程中执行 chDB 查询。发行版 0.1.4 保留 SQL/control 版本 0.1，在 Linux 或 macOS 上支持 PostgreSQL 15 及以上，依赖配套辅助程序和兼容的 libchdb。

### 查询流程

该扩展不受信任、可重定位，需要超级用户安装，无需共享预加载：

```sql
CREATE EXTENSION chdb;
SELECT *
FROM chdb_query('SELECT number, number * number FROM numbers(5)')
AS result(number bigint, square bigint);
SELECT pgchdb_version();
```

`chdb_query(text)` 返回记录，必须显式声明与查询结果匹配的列。普通角色默认没有执行权限，应仅向已存在且获准的角色授予 `chdb_query(text)` 的 `EXECUTE` 权限。每次调用都会创建并删除临时 chDB 数据库，因此一次调用创建的对象不会保留到下一次调用。

### 资源与凭据设置

超级用户设置 `chdb.max_memory`、`chdb.max_threads` 和 `chdb.max_parsing_threads` 控制辅助进程资源。默认内存上限改为 4 GB，`0` 表示取消该限制。两个线程设置默认都是 `0`，由 chDB 决定线程数；大型工作负载应设置明确上限。

0.1.4 改变了主机云凭据的使用规则。`chdb.host_credentials_roles` 默认为空列表，只有超级用户列出的已有 PostgreSQL 角色的成员，才能使用主机 AWS 凭据；其他调用者未传入密钥时会发起匿名请求。该设置不会授予查询函数执行权限。列出的角色成员能使用主机凭据本身具有的访问权限。

### 数据与升级边界

查询能访问 chDB 可达的数据源，应审查端点权限、显式凭据和返回类型。复杂数组、映射、嵌套值及编码转换需要匹配 PostgreSQL 类型，并用代表性数据检验往返转换。辅助进程的内存与 PostgreSQL 隔离，但仍消耗主机 CPU、内存、磁盘和网络资源。

`pgchdb_version()` 返回完整动态库版本。已有 SQL 版本 0.1 在更新到 0.1.4 动态库时无需 SQL 升级；安装匹配的库及辅助程序后，使用新会话加载新二进制。本版还为 `chdb_hook` 增加了可选扩展包装，详见其独立文档。
