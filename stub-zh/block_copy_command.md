## 用法

- 来源：[README](https://github.com/rustwizard/block_copy_command/blob/master/README.md)

`block_copy_command` 安装一个会拦截 `COPY` 语句的 `ProcessUtility` hook。库加载后，该 hook 在集群范围内生效；`CREATE EXTENSION` 只是在某个数据库中登记元数据。

### 启用扩展

```conf
shared_preload_libraries = 'block_copy_command'
```

```sql
CREATE EXTENSION block_copy_command;
```

### 阻断规则

默认情况下，非超级用户不能执行 `COPY TO` 或 `COPY FROM`：

```sql
COPY my_table TO STDOUT;
COPY my_table FROM STDIN;
COPY (SELECT * FROM my_table) TO '/tmp/out.csv';
```

文档中的优先级如下：

- `block_copy_command.blocked_roles`：始终阻断，即使是超级用户。
- `block_copy_command.block_program = on`：对所有用户阻断 `COPY ... PROGRAM`。
- `block_copy_command.enabled = off`：允许不在 `blocked_roles` 中的角色执行 `COPY`。
- 其他情况下，超级用户会绕过方向阻断。
- `block_copy_command.block_to` 和 `block_copy_command.block_from` 控制对非超级用户的导出/导入阻断。

### 主要配置

- `block_copy_command.enabled`：非超级用户阻断的总开关。
- `block_copy_command.block_to`：阻断 `COPY TO`。
- `block_copy_command.block_from`：阻断 `COPY FROM`。
- `block_copy_command.block_program`：对所有用户阻断 `COPY TO/FROM PROGRAM`。
- `block_copy_command.hint`：给被阻断命令的错误追加自定义 `HINT`。
- `block_copy_command.blocked_roles`：逗号分隔的始终阻断角色列表。
- `block_copy_command.audit_log_enabled`：将拦截事件写入审计表。

### 审计与注意事项

允许和被阻断的尝试都会被拦截；扩展定义了 `block_copy_command.audit_log`，并会把被阻断事件写入服务器日志。README 提到一个重要限制：被阻断的审计行是在抛出错误前插入的，因此会随事务一起回滚。实际使用中，PostgreSQL 服务器日志才是被阻断 `COPY` 尝试的权威记录。
