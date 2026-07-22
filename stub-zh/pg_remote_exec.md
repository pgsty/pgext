## 用法

来源：

- [Official README](https://github.com/cybertec-postgresql/pg_remote_exec/blob/1ac5d03725b86a1bebcf754b2e7f0212a46218c5/README.md)
- [Extension SQL](https://github.com/cybertec-postgresql/pg_remote_exec/blob/1ac5d03725b86a1bebcf754b2e7f0212a46218c5/pg_remote_exec--1.0.sql)
- [C implementation](https://github.com/cybertec-postgresql/pg_remote_exec/blob/1ac5d03725b86a1bebcf754b2e7f0212a46218c5/pg_remote_exec.c)

pg_remote_exec 从 PostgreSQL 后端执行操作系统 shell 命令。尽管名称带有 remote，命令实际运行在数据库服务器主机，而不是另行配置的远程主机。这是从 SQL 逃逸到操作系统的高危特权能力，通常不应安装。

### 核心流程

只有超级用户或 `pg_execute_server_program` 成员可以调用这些函数。如果经过审计的管理流程确实没有更安全的替代方案，应撤销默认访问，并包装固定命令，而不是接收用户提供的文本：

```sql
CREATE EXTENSION pg_remote_exec;

REVOKE ALL ON FUNCTION pg_remote_exec(text) FROM PUBLIC;
REVOKE ALL ON FUNCTION pg_remote_exec_fetch(text, boolean) FROM PUBLIC;

SELECT pg_remote_exec_fetch('/usr/local/sbin/report-db-health', false);
```

扩展会把文本交给主机 shell。使用 SQL 参数并不能令 shell 元字符变安全，命令产生的外部副作用也不受事务保护。

### 重要对象

- `pg_remote_exec(text)` 调用 C 库 `system()` 函数并返回其原始整数状态。
- `pg_remote_exec_fetch(text, boolean)` 调用 `popen()`，把标准输出按每行一条记录返回，并根据 boolean 决定对非零关闭状态报错还是忽略。
- `pg_execute_server_program` 是执行时检查的 PostgreSQL 预定义角色。

### 安全与运维说明

授予执行权限实际上等同于允许以数据库服务端操作系统账户运行任意命令。它可以读取文件、修改 PostgreSQL 之外的数据、建立网络连接、耗尽资源或植入恶意程序。fetch 函数不捕获 stderr，取消与资源清理并不完整，长时间命令还会占用后端。应优先使用具备独立认证、允许列表、超时、资源限制和审计日志的作业执行器。
