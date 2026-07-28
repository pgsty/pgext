## 用法

来源：

- [官方上游 README](https://github.com/timvaillancourt/pg_mysql/blob/d79179caf3013967faf220af08207adbf9c790e5/README.md)
- [官方扩展控制文件 (pg_mysql.control)](https://github.com/timvaillancourt/pg_mysql/blob/d79179caf3013967faf220af08207adbf9c790e5/pg_mysql.control)
- [官方扩展 SQL (pg_mysql--1.0.sql)](https://github.com/timvaillancourt/pg_mysql/blob/d79179caf3013967faf220af08207adbf9c790e5/pg_mysql--1.0.sql)

`pg_mysql` — 为 PostgreSQL 提供企业级复制 - 终于来了！在移植或模拟相应的数据库 API 时使用它。使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_mysql;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `mysql.delete(port integer)` 是一个扩展函数，返回 `text`。
- `mysql.query(sql text, port integer DEFAULT 0)` 是一个扩展函数，返回 `text`。
- `mysql.start(port integer DEFAULT 0, semi_sync boolean DEFAULT false)` 是一个扩展函数，返回 `mysql`。
- `mysql.start_replica(source_host text, source_port integer, port integer DEFAULT 0, semi_sync boolean DEFAULT false)` 是一个扩展函数，返回 `mysql`。
- `mysql.status(port integer DEFAULT 0)` 是一个扩展函数，返回 `SETOF`。
- `mysql.stop(port integer)` 是一个扩展函数，返回 `text`。
- `mysql.start_result` 是一个扩展定义的类型。
- `mysql.status_info` 是一个扩展定义的类型。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
