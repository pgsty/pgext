## 用法
- GitHub 仓库: [`rustwizard/block_copy_command`](https://github.com/rustwizard/block_copy_command)
- README: [rustwizard/block_copy_command/blob/master/README.md](https://github.com/rustwizard/block_copy_command/blob/master/README.md)

`block_copy_command` 通过安装 `ProcessUtility` hook，在整个集群范围内拦截 `COPY` 命令。它通过 `shared_preload_libraries` 加载，而 `CREATE EXTENSION` 只是在每个数据库中登记扩展元数据。

这个扩展适用于希望默认禁止非超级用户执行 `COPY TO` 和 `COPY FROM`，同时仍可通过 GUC 和审计表进行更细粒度控制的部署场景。

### 安装与启用

```conf
shared_preload_libraries = 'block_copy_command'
```

```sql
CREATE EXTENSION block_copy_command;
```

README 指出，一旦库被加载，这个 hook 就会对整个集群生效。

### 拦截规则

默认情况下，非超级用户无法执行 `COPY`。

```sql
COPY my_table TO STDOUT;
COPY my_table FROM STDIN;
COPY (SELECT * FROM my_table) TO '/tmp/out.csv';
```

超级用户默认可以绕过拦截，除非它们被列入 `block_copy_command.blocked_roles`，或者启用了 `block_copy_command.block_program`。`COPY ... PROGRAM` 默认对所有用户都被阻止。

### 配置项

- `block_copy_command.enabled` 控制是否拦截非超级用户。
- `block_copy_command.block_to` 控制是否阻止 `COPY TO`。
- `block_copy_command.block_from` 控制是否阻止 `COPY FROM`。
- `block_copy_command.block_program` 阻止所有用户执行 `COPY TO/FROM PROGRAM`。
- `block_copy_command.hint` 会为被阻止的命令附加自定义 `HINT:`。
- `block_copy_command.blocked_roles` 会永久阻止指定角色，包括超级用户。
- `block_copy_command.audit_log_enabled` 控制是否将拦截到的 `COPY` 事件写入 `block_copy_command.audit_log`。

### 审计日志

扩展会把被拦截的 `COPY` 活动记录到 `block_copy_command.audit_log`，并将被阻止事件以 `LOG` 级别写入 PostgreSQL 服务器日志。

README 中常见的监控查询包括查看最近事件、筛选被阻止事件，以及按用户分组统计。

### 范围

上游 README 已经覆盖了需求、启用方式、拦截行为、主要 GUC、审计表和测试覆盖，因此这个 stub 不需要额外的项目主页或文档站内容。
