## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/vibhorkum/edb_block_commands/blob/9cf36b955c54b47578b303bb27bae16718352863/README.md)
- [扩展 control 文件](https://github.com/vibhorkum/edb_block_commands/blob/9cf36b955c54b47578b303bb27bae16718352863/edb_block_commands.control)
- [C 实现](https://github.com/vibhorkum/edb_block_commands/blob/9cf36b955c54b47578b303bb27bae16718352863/edb_block_commands.c)

`edb_block_commands` 是仅适用于 EDB Advanced Server 的策略钩子，可限制超级用户执行工具命令、DML 与读取命令。已复核的上游说明只声称在 EPAS 10 及更高版本测试过；它不是可移植的 PostgreSQL 扩展。通过 `session_preload_libraries` 为指定角色加载，并重新建立角色会话。

```sql
CREATE EXTENSION edb_block_commands;
ALTER ROLE guarded_admin
  SET session_preload_libraries = '$libdir/edb_block_commands';
```

命令级白名单包括 `edb_block_commands.su_read_whitelist`、`edb_block_commands.su_write_whitelist` 与 `edb_block_commands.su_alter_system_whitelist`；`edb_block_commands.su_whitelist` 会让列出的超级用户绕过全部限制。所有白名单变更都应视为特权安全策略变更，并在隔离的 EPAS 实例验证。

函数 `edb_switch_user(text)`，尤其是 `edb_switch_user_u(text)`，会跨越角色边界；上游将后者描述为把普通用户切换成超级用户。不要向不可信角色授予执行权。该项目属于实验性质且没有当前支持或兼容矩阵，部署前必须审计其 C 级权限检查，并继续以原生角色隔离作为主要控制。
