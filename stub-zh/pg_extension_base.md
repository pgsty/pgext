## 用法

来源：

- [官方 pg_extension_base README](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/pg_extension_base/README.md)
- [版本 3.5 控制文件](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/pg_extension_base/pg_extension_base.control)
- [SQL API 定义](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/pg_extension_base/pg_extension_base--1.6.sql)
- [3.4 至 3.5 工作进程状态升级](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/pg_extension_base/pg_extension_base--3.4--3.5.sql)
- [pg_lake 构建和预加载指南](https://github.com/Snowflake-Labs/pg_lake/blob/v3.5.3/docs/building-from-source.md)

`pg_extension_base` 是 Snowflake 为其他 PostgreSQL 扩展提供的基础设施扩展。它提供了基于控制文件的库预加载、数据库级生命周期工作进程、依赖感知型更新以及短期附加工作进程。应用程序用户通常将其作为 `pg_lake` 的依赖项安装；扩展开发人员可以直接使用其 C 和 SQL API。

pg_lake 发布与软件包版本为 `3.5.3`，SQL 扩展版本为 `3.5`。动态库与查询服务器应使用同一发布版本。

### 启用基础设施

在其他扩展可以使用 `pg_extension_base` 的启动机制之前，必须预加载它：

```conf
shared_preload_libraries = 'pg_extension_base'
```

重启 PostgreSQL，然后在一个托管依赖扩展的数据库中创建它：

```sql
CREATE EXTENSION pg_extension_base;

SELECT * FROM extension_base.list_preload_libraries();
SELECT * FROM extension_base.list_base_workers();
```

### 扩展开发人员工作流

一个依赖扩展可以使用控制文件指令请求启动加载：

```text
requires = 'pg_extension_base'
module_pathname = '$libdir/my_extension'
#!shared_preload_libraries
```

要附加数据库生命周期工作者，定义一个内部 C 函数并在扩展安装期间注册它：

```sql
CREATE FUNCTION my_extension.main_worker(internal)
RETURNS internal
LANGUAGE C
AS 'MODULE_PATHNAME', 'my_extension_main_worker';

SELECT extension_base.register_worker(
    'my_extension_main',
    'my_extension.main_worker'
);
```

基础设施在服务器启动、`CREATE EXTENSION` 或从模板创建数据库后会启动该工作者，并尝试在 `DROP EXTENSION` 或 `DROP DATABASE` 时停止它。

### SQL API 索引

- `extension_base.list_preload_libraries()`: 报告发现用于启动加载的扩展/库对。
- `extension_base.register_worker(name, regproc)` 和 `deregister_worker(name)`: 管理生命周期工作者注册；公共执行被撤销。
- `extension_base.list_base_workers()`: 报告基础工作者的数据库、扩展、PID 及重启状态。
- `extension_base.list_database_starters()`: 报告每个数据库的启动进程。
- `extension_base.run_attached(command, dbname)`: 在短生命周期工作者中运行命令，可选地在另一个数据库中，并返回命令 ID 和标签。

### 运维边界

- 生命周期工作者函数以超级用户身份运行且在事务之外。它们必须开始自己的事务、检查中断并避免信任用户控制的 SQL。
- 关于失败 `DROP` 操作的工作者终止是尽力而为；工作者代码必须容忍扩展短暂消失或停止被撤销。
- `run_attached` 用于独立工作的有界任务，可能独立于调用者提交。它不是长时间运行任务的分离作业队列。
- 版本 `3.5` 在 `extension_base.list_base_workers()` 中新增 `failure_count`，便于观察工作进程的重启退避。升级会删除并重建该函数；迁移前应检查依赖对象及显式授权。
