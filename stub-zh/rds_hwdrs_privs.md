## 用法

来源：

- [华为云 RDS for PostgreSQL rds_hwdrs_privs 文档](https://support.huaweicloud.com/intl/en-us/usermanual-rds-pg/rds_09_0068.html)

`rds_hwdrs_privs` 是华为云 RDS for PostgreSQL 托管扩展，从服务管理的 `root` 角色委派有限的复制和目录权限。它面向 `root` 尚不能直接执行这些操作的旧 RDS 引擎；并不是可安装的社区扩展。

### 检查与安装

华为文档将其用于 PostgreSQL 9.5、9.6、10 和 11.5 或更早版本。更高版本允许 `root` 直接执行相关操作。应检查精确实例，并使用服务商控制函数，而不是原生 `CREATE EXTENSION`：

```sql
SELECT *
FROM pg_available_extension_versions
WHERE name = 'rds_hwdrs_privs';

SELECT control_extension('create', 'rds_hwdrs_privs');
```

只有 `root` 或 `root` 成员可以使用扩展函数。

### 委派操作

```sql
SELECT control_select_on_pg_authid('grant', 'drs_sync');
SELECT control_user_privilege('replication', 'drs_sync');
SELECT control_user_privilege('bypassrls', 'drs_sync');

SELECT create_publication_for_all_tables(
  'drs_all_tables',
  'insert, update, delete'
);

SELECT exec_pg_replication_origin_func(
  'pg_replication_origin_create',
  'drs_origin'
);
```

`control_select_on_pg_authid` 接受 `grant` 或 `revoke`。`control_user_privilege` 接受 `BYPASSRLS`、`NOBYPASSRLS`、`REPLICATION` 或 `NOREPLICATION`。发布辅助函数会创建由 `root` 所有的 `FOR ALL TABLES` 发布；PostgreSQL 10 支持 insert/update/delete 发布操作，PostgreSQL 11 还支持 truncate。复制源包装器只允许服务商列出的 `pg_replication_origin_*` 操作，并且仅当所选函数需要参数时才提供第二个参数。

### 安全与生命周期边界

这些权限会暴露密码角色元数据、绕过行级安全、允许复制连接，并影响全表逻辑发布。应只授予迁移/DRS 账户所需能力，使用专用角色，并在任务结束后撤销。授予 `root` 成员关系比调用辅助函数宽泛得多。

删除扩展使用 `SELECT control_extension('drop', 'rds_hwdrs_privs')`。应先有计划地删除委派权限和复制对象，不能假定卸载会反转运行状态。可用性与语义属于华为托管引擎，应与服务商和目标 DRS 工作流协调。
