## 用法

来源：

- [官方上游 README](https://github.com/viveknathani/pg_nowhere/blob/0d3ae6a044a4954bd2e65cee01d47fc03394e5b6/README.md)
- [官方扩展控制文件 (pg_nowhere.control)](https://github.com/viveknathani/pg_nowhere/blob/0d3ae6a044a4954bd2e65cee01d47fc03394e5b6/pg_nowhere.control)
- [官方扩展 SQL (pg_nowhere--0.1.sql)](https://github.com/viveknathani/pg_nowhere/blob/0d3ae6a044a4954bd2e65cee01d47fc03394e5b6/pg_nowhere--0.1.sql)

`pg_nowhere` — 一个 PostgreSQL 扩展，用于禁止没有 WHERE 子句的 UPDATE 和 DELETE 查询。在管理或自动化上述数据库行为时使用它。请使用链接中的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_nowhere;

-- Create test table
CREATE TABLE users (id SERIAL, name TEXT);
INSERT INTO users (name) VALUES ('Alice'), ('Bob');

-- This should work
UPDATE users SET name = 'Charlie' WHERE id = 1;

-- This should fail
UPDATE users SET name = 'Dave';
DELETE FROM users;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
