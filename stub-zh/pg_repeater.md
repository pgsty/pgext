## 用法

来源：

- [官方 README](https://github.com/danolivo/pg_repeater/blob/master/README.md)
- [扩展 control 文件](https://github.com/danolivo/pg_repeater/blob/master/pg_repeater.control)
- [钩子实现](https://github.com/danolivo/pg_repeater/blob/master/pg_repeater.c)

`pg_repeater` 0.1 是一个 PostgreSQL 补丁加扩展原型，可将 utility 命令和序列化执行计划转发到远端服务器。它依赖 postgres_fdw、pg_execplan 以及兼容的补丁源码树，是研究代码，而不是复制或分布式事务系统。

### 核心流程

本地和远端服务器需要匹配的对象布局及配套执行支持。先配置由扩展 GUC 指定名称的 postgres_fdw 服务器，再预加载钩子库并重启 PostgreSQL。

```conf
shared_preload_libraries = 'pg_repeater'
repeater.fdwname = 'remoteserv'
```

```sql
CREATE EXTENSION postgres_fdw;
CREATE EXTENSION pg_execplan;
CREATE EXTENSION pg_repeater;

CREATE SERVER remoteserv
FOREIGN DATA WRAPPER postgres_fdw
OPTIONS (host '/tmp', dbname 'remote_compute');

CREATE USER MAPPING FOR CURRENT_USER
SERVER remoteserv
OPTIONS (user 'remote_executor');
```

### 执行模型

- ProcessUtility 和 executor 钩子拦截符合条件的语句；计划被序列化，对象 OID 转为可移植表示，再通过配置的 postgres_fdw 连接发送。
- 远端通过 pg_exec_plan 执行传入计划，因此两端必须具有兼容的 PostgreSQL 内部接口、schema、函数、类型和数据假设。
- 实现排除了 COPY、CREATE EXTENSION、EXPLAIN、FETCH 和 VACUUM 等若干 utility 类别。

### 关键限制

- 会话内 LOAD 只能为当前会话安装钩子；要一致地拦截全服务器会话必须预加载，更改预加载列表后需要重启。
- 本地与远端执行之间没有分布式提交协议或可靠的回滚耦合。错误、网络中断或拓扑漂移都可能使节点分叉。
- 序列化 PostgreSQL 计划树属于私有且版本敏感的内部接口。不能混用服务器构建或扩展 ABI，也不要把这个原型作为生产可用性或一致性机制。
