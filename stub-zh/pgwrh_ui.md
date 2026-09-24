## 用法

来源：

- [Console guide](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_ui/README.md)
- [Control file](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_ui/pgwrh_ui.control)
- [Viewer role setup](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_ui/readonly.sql)
- [Operator role setup](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_ui/operator.sql)
- [PostgREST configuration](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_ui/postgrest.conf)

`pgwrh_ui` 1.0.0-alpha1 为 PostgreSQL 18 上的 pgwrh 控制节点提供浏览器控制台，展示分片放置、副本状态和变更阻碍，并可启用管理操作。它只应安装在控制节点数据库中。独立的 PostgREST 进程负责提供 SQL 端点及内嵌浏览器资源。

### 启用控制台

数据库需要 `pgwrh` 1.0.0-alpha1 和 PL/pgSQL。控制台不依赖 `pgwrh_wait`，也不需要它的预加载。PostgREST 必须支持自定义媒体处理器；上游 Compose 示例使用 14.16。以受信任的管理员身份连接控制节点后执行：

```sql
CREATE EXTENSION pgwrh_ui;
```

通过随附的 `readonly.sql` 创建部署角色，需要管理功能时再执行 `operator.sql`。这些角色不会随扩展自动创建或删除。在源码目录中，将 `CONTROLLER_ADMIN_URI` 设置为管理员连接后执行：

```sh
psql -X -v ON_ERROR_STOP=1 "$CONTROLLER_ADMIN_URI" -f pgwrh_ui/readonly.sql
```

```sql
CREATE ROLE pgwrh_ui_authenticator LOGIN NOINHERIT;
GRANT pgwrh_ui_viewer TO pgwrh_ui_authenticator;
```

按照当前部署的常规方式配置该登录角色的认证，然后使用随附配置启动外部服务：

```sh
export PGRST_DB_URI='postgresql://pgwrh_ui_authenticator@localhost/controller_database'
postgrest pgwrh_ui/postgrest.conf
```

随附配置以查看者权限在本地 `/rpc/index` 路径提供控制台。只应暴露 `pgwrh_ui` 模式。扩展或函数发生变化后，重新加载 PostgREST 的模式缓存：

```sql
NOTIFY pgrst, 'reload schema';
```

### 访问与管理

`pgwrh_ui_viewer` 可以查看控制节点上的所有组。`pgwrh_ui_operator` 还可修改副本权重与路由状态、登记副本，以及启动、提交或回滚变更。应在执行操作员设置脚本并配置认证后，才授予操作员角色成员资格。控制台没有内置登录页或令牌管理器；远程访问需要经过认证的 PostgREST 角色或认证反向代理。

副本准备、复制凭据、源表授权、复制组创建和表策略仍通过 SQL 管理。每个组共用一份草稿；基于过期状态提交的表单会返回 HTTP 409，需要重新审查。维护模式的路由变更不会删除数据或停止已分配的复制。

### 监控边界

就绪报告不提供每副本的心跳时间戳或持久历史记录。阻碍数为零不能证明副本当前可达。显示的 WAL 延迟是字节距离，不是经过时间，也不构成读取一致性保证。控制节点故障会同时导致控制台不可用。此 alpha 仅支持全新安装；安装软件包不会启动 PostgREST。
