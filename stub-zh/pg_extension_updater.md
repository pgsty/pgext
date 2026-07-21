## 用法

来源：

- [pg_extension_updater 官方 README](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_updater/README.md)
- [3.4 版控制文件](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_updater/pg_extension_updater.control)
- [工作进程注册 SQL](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_updater/pg_extension_updater--1.1.sql)
- [更新器实现与配置](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_updater/src/pg_extension_updater.c)
- [pg_extension_base 预加载文档](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_base/README.md)

`pg_extension_updater` 会在数据库生命周期工作进程启动时，对已安装版本与可用默认版本不同的每个扩展执行 `ALTER EXTENSION ... UPDATE`。它用于减少部署新扩展文件后 SQL 与二进制版本不一致的问题；它不会安装缺失的扩展，也不能替代发布测试。

### 启用自动更新

更新器依赖 `pg_extension_base`，后者必须在整个集群中预加载：

```conf
shared_preload_libraries = 'pg_extension_base'
```

重启 PostgreSQL 后，在每个需要自动更新的数据库中创建更新器：

```sql
CREATE EXTENSION pg_extension_updater CASCADE;
```

在 `template1` 中创建它，会让从该模板克隆的新数据库包含更新器。工作进程不会更新模板数据库自身的扩展，但会在新克隆的数据库中启动。

### 运行时行为

- 更新器控制文件中的 `#!shared_preload_libraries` 指令允许 `pg_extension_base` 加载其库。
- 安装过程会注册内部生命周期工作进程 `extension_updater.main(internal)`。
- 工作进程启动时读取 `pg_available_extensions`，并更新其中 `installed_version` 与 `default_version` 不同的条目。
- `ALTER EXTENSION` 失败会记为警告，并且在该工作进程的本轮运行中只尝试一次。
- `pg_extension_updater.enable` 是 postmaster 参数，默认为 `on`；设为 `off` 会禁用工作进程更新。修改后必须重启。

### 注意事项

- 自动迁移可能执行任何已安装扩展提供的升级 SQL。在生产数据库启用前，应验证软件包及升级路径。
- 应独立审查扩展依赖变化，并按应用需求制作备份；某项警告不会回滚其他已经成功的扩展更新。
- `3.4` 版没有面向用户的强制更新函数，也没有逐扩展允许列表。
- 与 `3.3` 相比，`3.4` 没有改变更新器的 SQL API。
