## 用法

来源：

- [官方源码 1](https://github.com/SriPrarabdha/pgsmith/blob/75bddc282d216f2154dd30f20cd1b67a08ebe4a1/README.md)
- [官方源码 2](https://github.com/SriPrarabdha/pgsmith/blob/75bddc282d216f2154dd30f20cd1b67a08ebe4a1/extension/pgsmith.control)
- [官方源码 3](https://github.com/SriPrarabdha/pgsmith/blob/75bddc282d216f2154dd30f20cd1b67a08ebe4a1/extension/build.py)
- [官方源码 4](https://github.com/SriPrarabdha/pgsmith/blob/75bddc282d216f2154dd30f20cd1b67a08ebe4a1/docs/install.md)
- [官方源码 5](https://github.com/SriPrarabdha/pgsmith/blob/75bddc282d216f2154dd30f20cd1b67a08ebe4a1/docs/sql-reference.md)
- [官方源码 6](https://github.com/SriPrarabdha/pgsmith/blob/75bddc282d216f2154dd30f20cd1b67a08ebe4a1/docs/sync.md)
- [官方源码 7](https://github.com/SriPrarabdha/pgsmith/blob/75bddc282d216f2154dd30f20cd1b67a08ebe4a1/docs/operations.md)
- [官方源码 8](https://github.com/SriPrarabdha/pgsmith/blob/75bddc282d216f2154dd30f20cd1b67a08ebe4a1/SECURITY.md)

`pgsmith` 0.1.0 将 Flagsmith 配置镜像到 PostgreSQL，并在 SQL 中计算功能开关。计算使用本地快照，新配置由独立同步服务或可选的数据库内集成获取。这是早期源码版本，使用 SQL/PL/pgSQL，不含需要编译的 PostgreSQL 库。

### 启用与同步

上游面向 PostgreSQL 14–18，要求数据库使用 UTF-8。在已安装扩展文件的服务器上执行：

```sql
CREATE EXTENSION pgsmith;
GRANT flagsmith_reader, flagsmith_tracker TO app_user;
GRANT flagsmith_loader TO sync_user;
```

控制文件将它标为可信扩展且不可迁移模式。它创建 `flagsmith` 与 `flagsmith_c` 模式；常规安装不需要预加载或重启。只有安装者具有创建角色的权限时，才会创建四个辅助角色，因此需要核对角色及授权。

`pgsmith` 同步命令通过私有配置获取具备 `flagsmith_loader` 权限的数据库连接、Flagsmith API 地址和服务端环境密钥：

```bash
pgsmith sync --once
```

另一种安装方式使用 `pg_tle` 安装器，需要该扩展及相应管理权限。RDS/Aurora 需要在预加载配置中启用 `pg_tle`，这是该安装方式的要求。可选的数据库内同步依赖 Supabase Vault、`http` 或 `pg_net`，调度还需要 `pg_cron`；本地开关计算本身不依赖这些组件。

### 计算功能开关

至少同步一个环境后：

```sql
SELECT * FROM flagsmith.flags('user-42', '{"plan":"pro","country":"IN"}'::jsonb);
SELECT flagsmith.is_enabled('new_checkout', 'user-42', '{"plan":"pro"}'::jsonb);
SELECT * FROM flagsmith.health();
```

`flagsmith.flags` 返回 `feature`、`enabled`、`value`、`variant`、`reason` 和 `metadata`。`flagsmith.flag` 获取单个开关；`flagsmith.value` 与 `flagsmith.variant` 返回其值或实验分组。`flagsmith.get_flags` 通过解释器计算指定环境或快照。若整条查询应共享一次开关计算，使用 `MATERIALIZED` CTE。

### 权限与运维

- `flagsmith_reader` 计算开关，`flagsmith_tracker` 记录曝光与事件，`flagsmith_planner` 增加审查权限，`flagsmith_loader` 改变正在使用的配置。默认不向 `PUBLIC` 授权。
- 读取角色可以计算任意身份及特征。若用于行安全决策，不应直接将该角色授予 Supabase 的 `anon` 或 `authenticated`；应提供绑定当前认证用户的受限包装函数。
- `flagsmith.pin` 与 `flagsmith.unpin` 控制正在使用的快照。`flagsmith.health`、`flagsmith.sync_status` 和 `flagsmith.event_status` 用于检查陈旧配置、同步及事件问题。同步失败时继续使用最后接受的配置。
- 备份不包含派生快照与编译函数。恢复后运行 `pgsmith sync --once` 或 `flagsmith.sync_tick` 重新生成。上游 0.1.0 尚无升级路径；按文档重装后需要重新同步。
- 仅所有者能添加远程同步来源，因为来源地址决定服务端密钥发往哪里。选择 `pg_net` 传输前应检查队列表权限，队列可能包含凭据。
