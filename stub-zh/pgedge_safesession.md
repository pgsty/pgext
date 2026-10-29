## 用法

来源：

- [官方文档](https://github.com/pgEdge/pgedge-safesession/blob/b406cd1e82cf21f30959351691c56de046562678/docs/index.md)
- [官方 README](https://github.com/pgEdge/pgedge-safesession/blob/b406cd1e82cf21f30959351691c56de046562678/README.md)
- [控制文件](https://github.com/pgEdge/pgedge-safesession/blob/b406cd1e82cf21f30959351691c56de046562678/pgedge_safesession.control)
- [Hook 与 GUC 实现](https://github.com/pgEdge/pgedge-safesession/blob/b406cd1e82cf21f30959351691c56de046562678/src/pgedge_safesession.c)
- [可选 SQL 登记](https://github.com/pgEdge/pgedge-safesession/blob/b406cd1e82cf21f30959351691c56de046562678/sql/pgedge_safesession--1.0.sql)
- [PostgreSQL CI 矩阵](https://github.com/pgEdge/pgedge-safesession/blob/b406cd1e82cf21f30959351691c56de046562678/.github/workflows/ci.yml)

`pgedge_safesession` 使用 SQL 版本 1.0，为配置的 PostgreSQL 角色强制只读会话。本文对应固定的源码修订，包含 1.0 标签之后的修复。上游要求 PostgreSQL 14 或更高版本，CI 覆盖 14–18。它补充普通的最小权限授权；以超级用户登录的会话豁免。

### 预加载与配置

Hook 必须在服务器启动时加载：

```ini
shared_preload_libraries = 'pgedge_safesession'
```

安装匹配库后重启 PostgreSQL。由超级用户在服务器配置中设置受限角色，而不是只在单个会话中设置：

```sql
ALTER SYSTEM SET pgedge_safesession.roles =
  'readonly_user,reporting_role';
SELECT pg_reload_conf();
```

模块无需 SQL 登记即可工作；可选登记会增加扩展目录项及状态辅助函数 `pgedge_safesession_is_restricted()`。限制锚定到 `session_user` 与实际角色授权。`SET ROLE`、临时提权和超级用户拥有的 `SECURITY DEFINER` 函数，都不能让已受限的登录会话逃脱限制。使会话进入受限状态的变更会使缓存计划失效。

### 保护控制

全部设置都是 `SUSET`。`pgedge_safesession.block_dml`、`block_ddl` 与 `block_c_functions` 默认开启；`block_all_c_functions` 默认关闭，启用后也会阻止 stable 与 immutable 的原生扩展函数。每条受限语句都会无条件重新强制只读事务状态，本次源码没有 `force_read_only` 设置。

默认阻止写入、DDL、COPY FROM、COPY TO PROGRAM、强表锁、受保护只读 GUC 的变更、有副作用的 volatile 函数及非受信任过程调用。普通 SELECT、不写入的 EXPLAIN ANALYZE、安全事务与配置操作、连接池复位仍可使用。普通 EXPLAIN 只规划而不执行语句。受信任语言的只读函数可以运行，函数内部的写入则被拦截；保守的允许列表仅豁免已审查的无害 volatile 内置函数。

### 安全边界

它不进行行过滤、可读数据脱敏或用户认证，也不限制豁免超级用户。若未阻止全部原生函数，IMMUTABLE/STABLE 标记仍是信任边界。类型输入输出函数或索引运算符类直接发起的 `fmgr` 调用可能绕过表达式检查；受限角色不能创建这些对象，因此应审查已有对象及权限。与 Spock 等其他 Hook 扩展的交互尚未经穷尽测试。依赖该策略前，应验证实际服务器构建、连接池、预备语句、角色变更和故障切换。
