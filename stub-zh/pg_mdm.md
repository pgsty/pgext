## 用法

来源：

- [Official source](https://github.com/trickle-labs/pg-mdm/blob/4942af8fb472d1276405b69f032506c7976c906f/pg_mdm.control)
- [Official source](https://github.com/trickle-labs/pg-mdm/blob/4942af8fb472d1276405b69f032506c7976c906f/Cargo.toml)
- [Official source](https://github.com/trickle-labs/pg-mdm/blob/4942af8fb472d1276405b69f032506c7976c906f/sql/archive/pg_mdm--0.15.0.sql)
- [Official source](https://github.com/trickle-labs/pg-mdm/blob/4942af8fb472d1276405b69f032506c7976c906f/docs/tutorials/first-entity.md)
- [Official source](https://github.com/trickle-labs/pg-mdm/blob/4942af8fb472d1276405b69f032506c7976c906f/docs/how-to/install.md)
- [Official source](https://github.com/trickle-labs/pg-mdm/blob/4942af8fb472d1276405b69f032506c7976c906f/docs/explanation/security.md)
- [Official source](https://github.com/trickle-labs/pg-mdm/blob/4942af8fb472d1276405b69f032506c7976c906f/docs/reference/compatibility.md)
- [Official source](https://github.com/trickle-labs/pg-mdm/blob/4942af8fb472d1276405b69f032506c7976c906f/sql/configure_helper.sql)

`pg_mdm` 0.15.0 是仅供合成数据演示的预览版，用于定义实体、匹配记录并发布主记录。上游明确排除真实工作负载和实时控制，发布资格检查仍处于阻塞状态。提供可下载的软件包并不表示已经具备生产使用资格。

### 准备预览环境

已验证的配置为 Linux amd64 上的 PostgreSQL 18，配合明确固定的 `pg_trickle` 0.108.4 源码构建。按官方流程，两者均须预加载并重启 PostgreSQL。扩展依赖 `pg_trickle`，安装需要超级用户，使用固定模式。

```ini
shared_preload_libraries = 'pg_trickle,pg_mdm'
```

```sql
CREATE EXTENSION pg_trickle;
CREATE EXTENSION pg_mdm;
SELECT extname, extversion FROM pg_extension
WHERE extname IN ('pg_trickle', 'pg_mdm');
```

不要随意替换生产端构建：集成层会检查图与增量能力，不兼容时会拒绝执行。配置时保留已有的其他预加载项。

### 定义并发布实体

先按首个实体教程创建源表，再通过 `mdm.create(mdm.entity(...))` 登记定义。定义包含 `mdm.source`、`mdm.field`、`mdm.match` 和 `mdm.golden_value`。下例假定教程中的客户定义和受限调用者已配置完成：

```sql
SELECT session_user, current_user;
SELECT mdm_admin.verify_installation();
SELECT mdm.refresh('customer', 'ALLOW');
SELECT * FROM mdm_out.customer;
SELECT * FROM mdm_out.customer_members;
SELECT * FROM mdm_out.customer_review;
```

`mdm.refresh` 发布值表、成员表和待审核表。仅定义实体不会持续刷新这些公共输出，需要显式调度刷新。`ALLOW` 允许教程使用全量刷新路径。使用 `prefer_source` 时，优先级由规范源名称排序确定；`sources` 数组只筛选可用来源，其排列顺序不决定优先级。

### 角色与恢复

使用随附的 `sql/configure_helper.sql`，将受保护对象绑定到独立的辅助角色。该角色必须为 NOLOGIN、NOSUPERUSER、NOBYPASSRLS，且没有角色成员关系。将其与操作角色和应用登录角色分离，只授予必要的模式与函数权限、源表 SELECT 和 MAINTAIN 权限，以及明确的输出访问权。系统同时检查 `session_user` 与 `current_user`；在超级用户会话中仅执行 `SET ROLE` 不够。

升级或恢复后应重新绑定并验证辅助角色。重启、恢复或克隆后须有意识地重新建立策略运行时激活状态；缺少预加载时无法激活。自动策略动作仅限 ASSIGN_QUEUE、ESCALATE 和 SET_DUE_AT，不会自动合并记录或选定主值。即使部分解析采用增量方式，完整发布、摘要与快照处理仍可能涉及全局工作。
