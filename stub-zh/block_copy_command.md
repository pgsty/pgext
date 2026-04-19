## 用法

- 来源：[README](https://github.com/rustwizard/block_copy_command/blob/master/README.md)

`block_copy_command` 安装一个 `ProcessUtility` hook 来拦截 `COPY` 语句。库一旦加载，该 hook 就会在整个集群范围内生效，而 `CREATE EXTENSION` 只是在某个数据库中登记元数据。

### 启用扩展

```conf
shared_preload_libraries = 'block_copy_command'
```

```sql
CREATE EXTENSION block_copy_command;
```

上游 README 标明支持 PostgreSQL 13-18。

### 阻断规则

默认情况下，非超级用户不能执行 `COPY TO` 或 `COPY FROM`：

```sql
COPY my_table TO STDOUT;
COPY my_table FROM STDIN;
COPY (SELECT * FROM my_table) TO '/tmp/out.csv';
```

文档给出的优先级如下：

- `block_copy_command.blocked_roles`：始终阻断，即使是超级用户。
- `block_copy_command.block_program = on`：对所有用户阻断 `COPY ... PROGRAM`。
- `block_copy_command.enabled = off`：允许不在 `blocked_roles` 中的角色执行 `COPY`。
- 其他情况下，超级用户可以绕过方向阻断。
- `block_copy_command.block_to` 和 `block_copy_command.block_from` 控制对非超级用户的导出/导入阻断。

### 主要配置

- `block_copy_command.enabled`：非超级用户阻断的总开关。
- `block_copy_command.block_to`：阻断 `COPY TO`。
- `block_copy_command.block_from`：阻断 `COPY FROM`。
- `block_copy_command.block_program`：对所有用户阻断 `COPY TO/FROM PROGRAM`。
- `block_copy_command.hint`：为被阻止的错误追加自定义 `HINT`。
- `block_copy_command.blocked_roles`：以逗号分隔的始终阻断角色列表。
- `block_copy_command.audit_log_enabled`：将拦截事件写入审计表。

### 审计与注意事项

无论命令最终被允许还是被阻止，扩展都会进行拦截；同时它还定义了 `block_copy_command.audit_log`，并会把被阻止事件写入服务器日志。README 特别说明了一个重要 caveat：被阻止事件的审计行是在抛错前插入的，因此会随着事务一起回滚。实际使用时，PostgreSQL 服务器日志才是被阻止 `COPY` 尝试的权威记录。
