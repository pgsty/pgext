## 用法

来源：

- [目录版本对应的官方 README](https://github.com/drewc/postgresql-migrations/blob/3da0fe400bfa37f92a78e8854f08522f3989079d/README.org)
- [目录版本对应的扩展 SQL](https://github.com/drewc/postgresql-migrations/blob/3da0fe400bfa37f92a78e8854f08522f3989079d/migration--0.0.1.sql)

`migration` 0.0.1 是运行在数据库内的迁移执行器。它登记带编号的向上/向下 SQL 文件，从 PostgreSQL 服务器文件系统读取并动态执行，再记录状态。它是较老的管理工具，不是面向不可信用户的安全迁移服务。

### 核心流程

把审核过的 SQL 文件放到数据库服务器，选定稳定的根目录，按顺序登记，并用单个管理会话执行迁移。

```sql
CREATE EXTENSION migration;

SELECT migration_set_env('ROOT_DIRECTORY', '/srv/db-migrations');
SELECT register_migration(1, '001_create.up.sql', '001_create.down.sql');
SELECT register_migration(2, '002_index.up.sql',  '002_index.down.sql');

-- Apply registered migrations through version 2.
SELECT migration_migrate_database(2);

-- Inspect the registry and recorded errors.
TABLE migration_registry;
TABLE migration_error;
```

### 对象索引

- `migration_registry` 保存迁移编号、向上/向下文件路径和应用状态。
- `migration_set_env(text, text)` 及相关环境辅助函数管理服务器端根目录等值。
- `register_migration(...)` 只登记路径，不会冻结文件内容或计算校验和。
- `migration_run(...)` 执行单个方向，`migration_migrate_database(numeric)` 则沿注册表移动到目标版本。

### 安全与运维

- 文件通过服务器端 COPY 从 PostgreSQL 主机读取，文件内容会作为 SQL 执行。扩展及其函数只能授权给可信管理员。
- 已登记文件应保持不可变、经过审核并有备份。注册表存储路径而非内容哈希，替换文件会改变后续实际执行的内容。
- 同一时刻只运行一个迁移器；实现没有为并发执行者提供 advisory lock 协议。
- 应在一次性副本上同时测试向上和向下路径，并提前备份数据库。错误记录有助于诊断，但不能替代事务审查和恢复规划。
