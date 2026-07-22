## 用法

来源：

- [官方 README](https://github.com/adjust/pg-roleman/blob/fdd31900d8f56fa260cbeefa8568444f3c717420/README.md)
- [0.3.0 版 SQL](https://github.com/adjust/pg-roleman/blob/fdd31900d8f56fa260cbeefa8568444f3c717420/extension/roleman--0.3.0.sql)
- [官方 PGXN 发行页](https://pgxn.org/dist/roleman/0.3.0/)

`roleman` 0.3.0 是用于参数化创建和修改角色、更改密码、授权、设置角色参数及删除角色的纯 SQL API。它会引用标识符，并对白名单角色属性和权限关键字进行检查，帮助应用避免自行拼接原始工具 SQL。

### 创建与配置角色

扩展固定安装到 `roleman` 模式，且需要超级用户安装。其函数使用调用者权限：调用者仍需具备结果命令 `CREATE ROLE`、`ALTER ROLE`、`GRANT`、`REVOKE` 或 `DROP ROLE` 所要求的 PostgreSQL 权限。

```sql
CREATE EXTENSION roleman;

SELECT roleman.create_role(
    'app_reader',
    ARRAY['LOGIN', 'NOINHERIT']
);

SELECT roleman.grant_schema(
    'app_reader',
    'reporting',
    ARRAY['USAGE']
);

SELECT roleman.grant_table(
    'app_reader',
    'reporting.daily_sales'::regclass,
    ARRAY['SELECT']
);

SELECT roleman.set_connection_limit('app_reader', 20);
SELECT roleman.set_search_path('app_reader', ARRAY['reporting', 'public']);
```

### 主要函数组

- `roleman.create_role(text, text[])` 和 `roleman.alter_base(text, text[])` 应用白名单角色属性。
- `roleman.role_set_password(text, text, timestamp)` 设置密码和可选过期时间。
- `roleman.grant_database`、`roleman.grant_schema`、`roleman.grant_schema_all`、`roleman.grant_table`、`roleman.grant_seq` 和 `roleman.grant_function` 应用白名单权限。
- `roleman.role_grant_role(text, text)` 授予角色成员资格。
- `roleman.set_guc`、`roleman.set_guc_from_current`、`roleman.set_connection_limit` 和 `roleman.set_search_path` 修改角色默认值。
- `roleman.role_blank_perms(text)` 广泛撤销当前数据库可见的权限与成员关系；`roleman.drop_role(text)` 删除角色。

### 安全性与已知缺陷

密码值作为 SQL 函数参数传输，可能通过语句日志、活动视图、监控或错误上下文泄露。应使用加密连接，并采用不会捕获 SQL 文本的凭据流程。

`roleman.role_blank_perms(text)` 既具破坏性，也不能作为全局完整重置：它遍历当前数据库中的模式，撤销显式数据库和对象权限及成员关系，但不会转移所有权、删除默认权限或检查所有其他数据库。必须先清点依赖并测试准确的撤权计划。

0.3.0 实现中的 `roleman.reset_guc(text, text)` 错误执行了 `ALTER USER ... SET ... FROM CURRENT`，并没有重置设置。不要调用它；应使用经审查的原生 `ALTER ROLE ... RESET ...` 命令。项目已归档，历史上只测试到 PostgreSQL 9.4 至 9.6，PostgreSQL 10 也只是预期兼容而非已证实。在目标版本上必须回归测试其目录查询和生成的工具命令。
