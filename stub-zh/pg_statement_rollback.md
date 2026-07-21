## 用法

来源：

- [pg_statement_rollback v1.6 README](https://github.com/lzlabs/pg_statement_rollback/blob/v1.6/README.md)
- [v1.5 到 v1.6 的变更](https://github.com/lzlabs/pg_statement_rollback/compare/v1.5...v1.6)

pg_statement_rollback 会在每条符合条件的语句之前自动创建保存点，使显式事务在某条语句报错后仍可继续使用。它模拟部分其他数据库的语句级回滚行为，但客户端在出错后仍必须执行 ROLLBACK TO SAVEPOINT。

该模块加载到后端进程中，不需要 CREATE EXTENSION。

### 加载模块

仅为当前会话加载：

    LOAD 'pg_statement_rollback.so';

如需针对指定角色或数据库启用，请将其加入 session_preload_libraries，然后重新连接：

    session_preload_libraries = 'pg_statement_rollback'

只有部署确实需要服务器级全局加载时才应使用 shared_preload_libraries；在服务器范围修改任一预加载列表，都必须遵守相应的重启或重新连接边界。

### 从失败语句中恢复

    BEGIN;
    INSERT INTO accounts(id, balance) VALUES (1, 100);
    INSERT INTO accounts(id, balance) VALUES (1, 200);
    -- duplicate-key error
    ROLLBACK TO SAVEPOINT "PgSLRAutoSvpt";
    UPDATE accounts SET balance = 150 WHERE id = 1;
    COMMIT;

保存点名称在使用引号时区分大小写。应用程序必须检测语句错误，并在继续执行前发送回滚命令。

### 配置索引

- pg_statement_rollback.enabled 为当前会话启用自动保存点。
- pg_statement_rollback.savepoint_name 更改自动保存点名称，并且只能由超级用户控制。
- pg_statement_rollback.enable_writeonly 将保存点创建限制在可能执行写入的语句。

### 1.6 版本行为

1.6 版增加 PostgreSQL 19 构建支持，并改进了只读事务检测。该模块不再在只读事务中创建自动保存点，并会在 SET TRANSACTION ... READ ONLY 之前释放初始保存点，避免干扰备份导出及其他只读会话。

### 注意事项

- 这不是透明重试机制：客户端必须显式回滚到自动保存点。
- 保存点会给每条受覆盖语句增加开销。大范围启用前，应先测量写密集型工作负载。
- 上游 README 警告该模块在启用断言的 PostgreSQL 构建上可能崩溃；未经测试，不要把开发构建的行为视为生产环境安全。
- 跨整个事务的错误、连接故障以及使会话失效的错误，都无法通过保存点修复。
