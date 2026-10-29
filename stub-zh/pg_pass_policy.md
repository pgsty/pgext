## 用法

来源：

- [README.md](https://github.com/gliatti/pg_pass_policy/blob/ad47a9b8902ad6375d6afa20aea38d383dfc37c8/README.md)
- [pg_pass_policy.control](https://github.com/gliatti/pg_pass_policy/blob/ad47a9b8902ad6375d6afa20aea38d383dfc37c8/pg_pass_policy.control)
- [pg_pass_policy--1.0.sql](https://github.com/gliatti/pg_pass_policy/blob/ad47a9b8902ad6375d6afa20aea38d383dfc37c8/sql/pg_pass_policy--1.0.sql)
- [pg_pass_policy.c](https://github.com/gliatti/pg_pass_policy/blob/ad47a9b8902ad6375d6afa20aea38d383dfc37c8/src/pg_pass_policy.c)

`pg_pass_policy` 管理 PostgreSQL 角色密码的复杂度、复用历史、到期时间、轮换和认证锁定。本次源码快照安装 SQL 版本 1.0，上游文档支持 PostgreSQL 14–18。它通过钩子影响整个集群的认证，但只在一个指定数据库保存策略。

### 启用

启用前先审查管理账户豁免和救援通道，将动态库加入已有预加载列表并重启 PostgreSQL：

```ini
shared_preload_libraries = 'pg_pass_policy'
pg_pass_policy.database = 'postgres'
pg_pass_policy.exempt_roles = 'postgres'
```

仅在指定数据库，以超级用户连接：

```sql
CREATE EXTENSION pg_pass_policy;
SELECT pass_policy.worker_status();
SELECT * FROM pass_policy.v_role_status;
```

启动期间，在工作进程载入锁定状态之前，非豁免连接可能收到 57P03，trust/peer 连接也不例外。单独配置的救援通道仅允许名单内的超级用户通过 peer 认证。工作进程名额不足或数据库不存在都可能阻止初始化；未验证管理恢复通道之前，不应启用。

### 策略配置与管理

```sql
SELECT pass_policy.create_profile('application',
  '{"min_length":24,"max_age":"30 days","min_age":"0"}');
SELECT pass_policy.assign_profile('app_user', 'application');
SELECT * FROM pass_policy.effective_profile('app_user');
SELECT * FROM pass_policy.v_expiring;
SELECT * FROM pass_policy.v_locked;
SELECT * FROM pass_policy.v_recent_audit;
```

默认策略组合了复杂度、历史、到期和认证失败阈值，默认要求 SCRAM-SHA-256。密码变更应在 `pg_pass_policy.database` 中执行；开启强制检查后，在其他数据库修改会被拒绝。

安装会创建固定的 `pass_policy` 模式及 `pass_policy_viewer`、`pass_policy_admin` 两个角色。管理函数使用 SECURITY DEFINER，并检查 `session_user`：应直接以有权限的成员连接，而非使用 SET ROLE。管理角色能够轮换非特权用户的凭据，应严格控制其成员资格。

### 轮换与安全边界

`pass_policy.rotate_password` 生成并设置密码，`pass_policy.switch_rotation_pair` 支持成对角色轮换。两者都会向调用者一次性返回明文。应使用校验证书的 TLS，并避免结果进入客户端表格、日志或终端历史；服务端的处理方式不能保护客户端副本。

`pass_policy.generate_password`、`pass_policy.check_password`、`pass_policy.failed_attempts` 和 `pass_policy.worker_status` 分别提供生成、预检查及运行状态。扩展不能锁定或使豁免角色过期。删除扩展后模式和管理角色仍会保留，应审查依赖后再单独清理。本目录尚未将该源码快照视为已有标签发布或已完成生产验收的版本。
