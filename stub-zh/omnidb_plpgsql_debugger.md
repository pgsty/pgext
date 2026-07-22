## 用法

来源：

- [官方 README](https://github.com/omnidb/plpgsql_debugger/blob/223f9150306e00ef7706119a71f0c52385125a99/README.md)
- [1.0.0 扩展 SQL](https://github.com/omnidb/plpgsql_debugger/blob/223f9150306e00ef7706119a71f0c52385125a99/omnidb_plpgsql_debugger--1.0.0.sql)
- [OmniDB 调试器文档](https://omnidb.readthedocs.io/en/latest/en/13_writing_and_debugging_plpgsql_functions.html)

`omnidb_plpgsql_debugger` 是通过 OmniDB 图形界面调试 PL/pgSQL 函数的 PostgreSQL 服务端组件。它不是独立的交互式调试器：服务端模块必须与 OmniDB 客户端配合，调试过程还会创建额外的本地数据库连接。

### 核心流程

先预加载模块并重启 PostgreSQL，再在每一个需要调试的数据库中创建扩展。

```conf
shared_preload_libraries = 'omnidb_plpgsql_debugger'
```

```sql
CREATE EXTENSION omnidb_plpgsql_debugger;

GRANT ALL ON SCHEMA omnidb TO debugger_user;
GRANT ALL ON ALL TABLES IN SCHEMA omnidb TO debugger_user;
```

使用调试角色连接 OmniDB，在函数编辑器中打开 PL/pgSQL 函数、设置断点，然后从图形界面启动调试器。

### 对象与连接要求

- `omnidb.omnidb_enable_debugger(character varying)`：初始化服务端调试会话。
- `omnidb.contexts`、`omnidb.variables` 与 `omnidb.statistics`：保存客户端使用的调试器状态。

创建扩展和授予调试访问权限都属于超级用户任务。额外的调试器连接必须在 `pg_hba.conf` 中允许同一数据库与角色。上游文档给出了本地 trust 规则，或配合 `.pgpass` 的密码认证；应优先采用经过认证的本地访问，并将规则限制到最小范围。

预加载需要重启服务器。对 `omnidb` 模式授予全部权限会暴露可修改的调试状态，因此只能授予可信开发者，并只在允许交互式调试的数据库中启用。
