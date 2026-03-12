

## 用法

> [ddl_historization: 跟踪 PostgreSQL 数据库中的所有 DDL 变更](https://github.com/rodo/pg_ddl_historization)

将数据库上所有的 DDL 变更（CREATE、ALTER、DROP 等）记录到历史化表中，用于审计和跟踪。

### 安装

```sql
CREATE EXTENSION ddl_historization;
```

该扩展安装事件触发器，自动捕获 DDL 语句并将其存储在历史化表中。

### 查询 DDL 历史

安装后，所有 DDL 变更会自动记录。查询历史表以查看已发生的变更：

```sql
SELECT * FROM ddl_history ORDER BY ddl_date DESC;
```

### 与 pg_tle 集成

在 AWS RDS 环境中，可以通过 `pg_tle` 部署该扩展：

```sql
-- 构建 pg_tle 部署文件
-- $ make pgtle
-- 然后在实例上执行 pgtle.ddl_historization-0.3.sql
```

### 注意事项

- DDL 语句通过 PostgreSQL 事件触发器捕获
- 支持 `CREATE`、`ALTER`、`DROP` 及其他 DDL 命令
- 被 `schedoc` 扩展用作依赖项，用于模式文档化
