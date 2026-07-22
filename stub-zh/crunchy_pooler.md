## 用法

来源：

- [Crunchy Bridge PgBouncer 文档](https://docs.crunchybridge.com/how-to/pgbouncer)
- [Crunchy Bridge 产品页面](https://www.crunchydata.com/products/crunchy-bridge)

`crunchy_pooler` 是 Crunchy Bridge 托管的辅助扩展，使该服务的 PgBouncer 能够通过 PostgreSQL 校验数据库角色。它是 Crunchy Bridge 专用组件，并非可移植的社区软件包；每个需要通过连接池访问的数据库都要单独启用它。

### 启用连接池访问

以超级用户连接每个目标数据库并安装扩展：

```sql
CREATE EXTENSION crunchy_pooler;
```

扩展会创建 `crunchy_pooler` 角色，并仅授予它访问服务商 `user_lookup` 认证辅助对象的权限。删除扩展会移除此认证路径，使 PgBouncer 无法再接受到该数据库的连接。

将集群原有连接串的端口改为 `5431` 即可使用 PgBouncer；端口 `5432` 仍然直连 PostgreSQL。用户名、密码、数据库名和 TLS 设置保持不变。

### 池化模式与限制

Crunchy Bridge 默认使用事务池化，因此不能假设会话级状态能跨事务保留。服务在 PgBouncer 1.22 及以后版本中支持预备语句，前提是 `max_prepared_statements` 大于零；文档中的服务默认值为 `250`。需要固定后端会话时可选择会话池化；语句池化则无法执行多语句事务。

超级用户和复制角色会被 PgBouncer 明确拒绝，必须通过端口 `5432` 直连。如果所选数据库没有安装此扩展，客户端会收到 `FATAL: bouncer config error`。可用性、升级和实现细节均由 Crunchy Data 控制，依赖特定 PgBouncer 版本或设置前应查阅当前服务文档。
