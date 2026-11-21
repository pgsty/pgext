## 用法

来源：

- [README.md](https://github.com/pierreforstmann/pg_start_sql/blob/a94782661c3f94a7f779a3cd61200d771fdba04a/README.md)
- [pg_start_sql.control](https://github.com/pierreforstmann/pg_start_sql/blob/a94782661c3f94a7f779a3cd61200d771fdba04a/pg_start_sql.control)
- [pg_start_sql--1.0.0.sql](https://github.com/pierreforstmann/pg_start_sql/blob/a94782661c3f94a7f779a3cd61200d771fdba04a/pg_start_sql--1.0.0.sql)

`pg_start_sql` 在 PostgreSQL 实例启动时执行配置的 SQL。所核验控制文件和安装 SQL 的版本为 1.0.0，上游报告已测试 PostgreSQL 9.5–19。全部启动语句均以超级用户执行，没有配置可改为低权限角色。

### 配置启动 SQL

将库加入现有预加载列表，并配置目标数据库及语句：

```ini
shared_preload_libraries = 'pg_start_sql'
pg_start_sql.dbname = 'postgres'
pg_start_sql.stmt = 'SELECT 1'
```

重启 PostgreSQL 使预加载生效，并按上游流程创建扩展：

```sql
CREATE EXTENSION pg_start_sql;
```

后台任务作用于整个实例，无需在每个数据库创建扩展。`pg_start_sql.dbname` 默认为 postgres。必须至少配置 `pg_start_sql.stmt` 或 `pg_start_sql.file` 之一，后者指定包含启动语句的文件。

### 运行边界

配置和 SQL 文件的写入权限应限于可信管理员，因为语句以超级用户执行。启动任务应允许在多次启动中安全重复，并核验执行结果和服务器日志。保留其他预加载库。控制文件允许 SQL 对象重定位，但这不改变实例级预加载和重启要求。
