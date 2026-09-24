## 用法

来源：

- [README](https://github.com/trainstar/synchro/blob/3de1e69cfeb7a093b43d7e986ead8cdae8eab4f2/README.md)
- [Control file / 控制文件](https://github.com/trainstar/synchro/blob/3de1e69cfeb7a093b43d7e986ead8cdae8eab4f2/extensions/synchro-pg/synchro_pg.control)
- [extensions/synchro-pg/sql/synchro_pg--0.3.0.sql](https://github.com/trainstar/synchro/blob/3de1e69cfeb7a093b43d7e986ead8cdae8eab4f2/extensions/synchro-pg/sql/synchro_pg--0.3.0.sql)
- [extensions/synchro-pg/src/registry.rs](https://github.com/trainstar/synchro/blob/3de1e69cfeb7a093b43d7e986ead8cdae8eab4f2/extensions/synchro-pg/src/registry.rs)
- [docs/src/content/docs/operations/configuration.mdx](https://github.com/trainstar/synchro/blob/3de1e69cfeb7a093b43d7e986ead8cdae8eab4f2/docs/src/content/docs/operations/configuration.mdx)

`synchro_pg` 是 Synchro 离线优先同步系统的数据库组件。固定源码中的扩展版本为 0.3.0：PostgreSQL 捕获和物化按作用域组织的变更，认证主机适配层服务客户端，原生客户端维护本地 SQLite 状态。

### 启用与就绪检查

上游目标版本为 PostgreSQL 18。由超级用户安装，预加载 `synchro_pg`，启用逻辑 WAL 并配置复制槽。将 `synchro.database` 和 `synchro.worker_login` 设为目标数据库和专用工作登录角色；postmaster 参数变更后需要重启。按上游要求为该登录角色配置 `synchro_worker` 角色和复制权限。

```sql
CREATE EXTENSION synchro_pg;
SELECT synchro.synchro_readiness();
```

### 同步工作流

为应用表定义稳定主键，并提供把行映射到服务端控制作用域的确定性成员函数。通过 `synchro.synchro_register_table` 注册各表，指定关系、成员函数、组合方式、时间戳与删除标记列、同步列和受影响作用域。注册会检查所有权和函数依赖，应用迁移必须保持这一约定。

通过 Go 适配层或同一认证协议的其他实现提供服务。客户端连接后提交本地修改意图，使用不透明游标拉取作用域变更，并在必要时重建。扩展负责捕获和增删改语义，主机适配层负责 HTTP 和认证。可移植种子库只提供服务端声明的可移植数据，不会赋予访问权限。

### 运行与限制

`synchro.replication_slot` 与 `synchro.publication_name` 选择逻辑捕获对象。`synchro.max_worker_heartbeat_age_seconds`、`synchro.max_wal_lag_bytes` 和 `synchro.max_wal_lag_seconds` 设置必须为正数的就绪阈值。应监控就绪状态、保留 WAL 和源事务大小。每个源事务最多解码 16 MiB 和 10,000 条负载记录；超限事务会阻塞推进，直至得到处理。

扩展不可重定位，安装时创建专用角色并撤销对模式、表和函数的公共访问。应按用途授予上游定义的操作员、适配层、工作进程、监控和种子角色。仓库早期的 v0.1.x 发布不能视为该扩展版本。
