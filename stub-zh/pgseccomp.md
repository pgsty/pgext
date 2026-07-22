## 用法

来源：

- [官方 README](https://github.com/CrunchyData/pgseccomp/blob/3ed6d0780dc5d9a5204e57ce7628544a3610e218/README.md)
- [1.0 版 SQL API](https://github.com/CrunchyData/pgseccomp/blob/3ed6d0780dc5d9a5204e57ce7628544a3610e218/pgseccomp--1.0.sql)
- [过滤器与 GUC 实现](https://github.com/CrunchyData/pgseccomp/blob/3ed6d0780dc5d9a5204e57ce7628544a3610e218/pgseccomp.c)

`pgseccomp` 将 Linux seccomp 系统调用过滤器应用到 PostgreSQL postmaster 与客户端后端。它可以允许选定系统调用、记录审计日志、返回错误或终止进程，并由集群级基础策略叠加更严格的会话策略。

### 核心流程

该库要求 Linux 内核 3.5 或更新版本，以及 libseccomp 2.4 或更新版本。把它加入 `shared_preload_libraries`，先使用日志记录或宽松的默认动作，重启 PostgreSQL，再创建扩展并检查实际生效的策略。

```sql
CREATE EXTENSION pgseccomp;
SELECT * FROM seccomp_filter() ORDER BY context, syscallnum;
```

集群设置包括 `pgseccomp.global_syscall_allow`、`pgseccomp.global_syscall_log`、`pgseccomp.global_syscall_error`、`pgseccomp.global_syscall_kill` 与 `pgseccomp.global_syscall_default`。对应的 `pgseccomp.session_syscall_*` 设置会叠加到客户端后端；`pgseccomp.session_roles` 用来选择角色专属列表。

`set_client_filter(default_action, allow_list, log_list, error_list, kill_list)` 可在一个会话中执行一次，施加相同或更严格的叠加策略。NULL 列表参数采用已配置的会话列表，而 `*` 会继承该动作的全局列表。安装脚本会撤销 `PUBLIC` 对两个 SQL 函数的执行权限。

### 安全与恢复

错误规则可能使 PostgreSQL 无法打开文件、分配内存、通信、派生工作进程或正常关停；`kill` 动作会发送 `SIGSYS`。应为准确的 PostgreSQL 构建与工作负载建立已观测系统调用清单，先引入 `log`，再考虑 `error` 或 `kill`，保留控制台级恢复入口，并测试启动、认证、备份、复制、扩展、维护及故障切换。

Seccomp 策略加载后只能变得更严格；会话不能通过 SQL 函数放宽继承的过滤器。审计事件通常从操作系统审计日志读取，而数据库管理员可能没有读取权限。应与主机安全团队协同制定策略与诊断流程，不能把该扩展当作完整沙箱。
