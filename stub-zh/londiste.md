## 用法

来源：

- [锁定提交的上游 README](https://github.com/pgq/londiste-sql/blob/55085fc4e8e6cb00e2b6255427fd86c50cd6fef3/README.rst)
- [锁定提交的扩展控制文件](https://github.com/pgq/londiste-sql/blob/55085fc4e8e6cb00e2b6255427fd86c50cd6fef3/londiste.control)
- [锁定提交的模式安装脚本](https://github.com/pgq/londiste-sql/blob/55085fc4e8e6cb00e2b6255427fd86c50cd6fef3/londiste--3.8.sql)
- [锁定提交的函数源码](https://github.com/pgq/londiste-sql/tree/55085fc4e8e6cb00e2b6255427fd86c50cd6fef3/functions)
- [锁定提交的权限授予](https://github.com/pgq/londiste-sql/blob/55085fc4e8e6cb00e2b6255427fd86c50cd6fef3/structure/grants.sql)

londiste 3.8 提供 Londiste 逻辑复制的 PostgreSQL 服务器端支持模式。它本身不是独立复制守护进程。其表和函数跟踪已注册表与序列、合并状态、待处理外键、分区维护、复制触发器和复制 DDL 执行。控制文件依赖 pgq_node，且扩展必须由超级用户安装。

### 安装与检查

应先从匹配的软件包安装 PGQ 组件，再安装 Londiste，随后由 Londiste 管理进程注册队列和表：

```sql
CREATE EXTENSION pgq;
CREATE EXTENSION pgq_node;
CREATE EXTENSION londiste;

SELECT extname, extversion
FROM pg_extension
WHERE extname IN ('pgq', 'pgq_node', 'londiste')
ORDER BY extname;

SELECT queue_name, table_name, local, merge_state
FROM londiste.table_info
ORDER BY queue_name, table_name;
```

SQL 仓库只包含服务器端支持代码。队列拓扑、提供者/订阅者设置、守护进程配置、复制、追赶、故障转移和维护，都必须遵循匹配 Londiste/SkyTools 客户端软件包的文档。

### 破坏性 DDL 与权限边界

该 API 会有意创建和删除复制触发器、删除和恢复外键、修改 session_replication_role、删除过期分区表，以及记录或发出 DDL 脚本。这些操作可能绕过常规触发器和约束，或直接破坏表。应使用专用数据库所有者和守护进程角色，将复制拓扑纳入配置管理，并在副本上测试修复流程。

已复核的 set_session_replication_role 函数是 SECURITY DEFINER，且没有受保护的 search_path。PostgreSQL 函数默认可由 PUBLIC 执行，而随附授权脚本只向公共用户授予模式使用权和若干状态表读权限，没有撤销函数执行权。安装后应立即从 PUBLIC 撤销每个 londiste 函数的执行权，再只向可信守护进程角色授予实际需要的管理入口。允许不可信用户进入数据库前，必须复核所有 SECURITY DEFINER 函数和模式所有权。

Londiste 状态必须与 PGQ 事件和外部守护进程保持一致。不应手工编辑 table_info、seq_info、pending_fkeys 或 applied_execute。应执行协调备份，监控队列延迟和失败事件，并在每个节点上演练升级。扩展固定安装到 pg_catalog 模式且不可重定位，它的 SQL 必须与已安装 pgq_node 版本匹配；混用独立来源的版本可能留下不可用函数或复制元数据。
