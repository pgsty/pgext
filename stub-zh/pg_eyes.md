## 用法

来源：

- [官方 README](https://github.com/alexandersamoylov/pg_eyes/blob/31c21bb1a307235bfda250727254491d505e0ee3/README.md)
- [1.4 版 SQL](https://github.com/alexandersamoylov/pg_eyes/blob/31c21bb1a307235bfda250727254491d505e0ee3/sql/pg_eyes--1.4.sql)
- [扩展控制文件](https://github.com/alexandersamoylov/pg_eyes/blob/31c21bb1a307235bfda250727254491d505e0ee3/pg_eyes.control)

`pg_eyes` 1.4 是固定安装在 `eyes` 模式中的纯 SQL 监控工具包。它把活动、锁、容量、目录、设置、复制和 `pg_stat_statements` 数据暴露为可直接采集的指标与视图。上游 1.x 分支面向 PostgreSQL 9.6，因此在新目录上使用前必须验证每个对象。

### 设置与基础查询

`pg_stat_statements` 必须正确安装和加载，通常要放入 `shared_preload_libraries`，其统计信息才有用。扩展本身需要超级用户安装，且不可迁移模式。

```conf
shared_preload_libraries = 'pg_stat_statements'
```

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pg_eyes;

SELECT stat_name, stat_value
FROM eyes.get_activity();

SELECT * FROM eyes.active_connections;
SELECT * FROM eyes.active_queryes;
SELECT * FROM eyes.blocked_queryes;
SELECT * FROM eyes.top_queryes LIMIT 20;
```

拼写 `active_queryes`、`blocked_queryes` 和 `top_queryes` 是上游 SQL API 的一部分。

### 重要对象

- `eyes.get_activity()` 返回会话、事务、数据库和后台写入器计数器、WAL、复制及响应时间的指标名和值。
- `eyes.get_activity(text)` 为选定分组执行表 `eyes.get_activity` 中保存的自定义标量指标 SQL。
- `eyes.get_pg_stat_activity()` 和 `eyes.get_pg_stat_statements()` 是由 `postgres` 拥有的 `SECURITY DEFINER` 包装器。
- `eyes.db_objects`、`eyes.db_tables`、`eyes.db_tables_size`、`eyes.db_indexes`、`eyes.db_indexes_size`、`eyes.db_attributes`、`eyes.db_functions` 和 `eyes.db_settings` 暴露对象清单及配置细节。

许多计数器是累计值，监控系统必须把它们转换为速率。根据上游 README，以 `_time` 结尾的指标使用毫秒。

### 安全与兼容性

定义者权限包装器有意向可执行它们的非超级用户暴露完整活动和语句文本。查询文本、客户端地址、设置文件路径和函数定义都可能敏感；安装后应显式审计并收窄 `EXECUTE` 与模式权限。

`eyes.get_activity(text)` 以函数所有者身份执行 `stat_query` 列。绝不能给不可信角色授予 `eyes.get_activity` 表的写权限，插入前也要审查每条自定义指标。旧升级路径还可能保留定义者权限管理函数；应使用 `\df+ eyes.*` 检查实际权限。

1.4 版 SQL 引用了 PostgreSQL 9.6 时代的目录列，例如历史上的 `pg_stat_statements.total_time`。这不能证明它兼容所有后续大版本。生产部署前，应针对准确的 PostgreSQL 和 `pg_stat_statements` 版本测试安装及每个选定视图，或先移植 SQL。
