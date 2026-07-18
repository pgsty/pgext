## 用法

来源：

- [已复核 commit 的 postgresql-migrations README](https://github.com/drewc/postgresql-migrations/blob/3da0fe400bfa37f92a78e8854f08522f3989079d/README.org)
- [已复核 commit 的 migration.control](https://github.com/drewc/postgresql-migrations/blob/3da0fe400bfa37f92a78e8854f08522f3989079d/migration.control)
- [版本 0.0.1 的安装 SQL](https://github.com/drewc/postgresql-migrations/blob/3da0fe400bfa37f92a78e8854f08522f3989079d/migration--0.0.1.sql)
- [上游迁移设计与示例](https://github.com/drewc/postgresql-migrations/blob/3da0fe400bfa37f92a78e8854f08522f3989079d/doc/migration.org)

`migration` 是一个纯 SQL 迁移运行器。它把编号后的向上/向下脚本路径记录在 `migration_registry` 中，从数据库服务器读取这些文件并执行其内容，再把结果写入 `migration_log` 和 `migration_error`。

### 登记并运行脚本

```sql
CREATE EXTENSION migration;

SELECT migration_setenv('ROOT_DIRECTORY', '/srv/postgresql/migrations');
SELECT register_migration(0, '000-up.sql', '000-down.sql');
SELECT register_migration(1, '001-up.sql', '001-down.sql');

SELECT migration_migrate_database(1);
SELECT migration_current_number();
SELECT migration_migrate_database(0);
```

路径在服务器端而不是客户端解析。已复核 SQL 暴露的是 `migration_migrate_database`；README 中较旧的 `migration.migrate_database()` 写法与已安装 API 不一致。

### 注意事项

- `migration_read_file` 通过动态 `COPY` 实现服务器端文件读取，`migration_eval` 会把所得文本作为 SQL 执行。应把函数权限和脚本目录限制给可信数据库运维人员。
- 向上和向下脚本以调用者的数据库权限运行，可任意修改模式或数据。执行前必须审查、校验、备份并演练。
- 版本 `0.0.1` 是 2018 年代码，且没有当前兼容性矩阵。其递归运行器和自定义错误日志不能替代发布编排、锁管理或外部审计轨迹。
