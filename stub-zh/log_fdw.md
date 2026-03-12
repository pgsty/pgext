

## 用法

> [log_fdw: 访问 Postgres 日志文件的外部数据包装器](https://github.com/aws/postgresql-logfdw)

### 创建服务器

```sql
CREATE EXTENSION log_fdw;

CREATE SERVER log_fdw_server FOREIGN DATA WRAPPER log_fdw;
```

### 列出可用日志文件

```sql
SELECT * FROM list_postgres_log_files();
```

返回 PostgreSQL 日志目录中每个日志文件的文件名和大小。

### 为 CSV 日志创建外部表

```sql
SELECT * FROM create_foreign_table_for_log_file(
  'postgresql_2024_01_15_csv',   -- 外部表名
  'log_fdw_server',               -- 服务器名
  'postgresql-2024-01-15.csv'     -- 日志文件名
);
```

### 为纯文本日志创建外部表

```sql
SELECT * FROM create_foreign_table_for_log_file(
  'postgresql_2024_01_15_log',
  'log_fdw_server',
  'postgresql-2024-01-15.log'
);
```

### 查询日志数据

```sql
-- 查询 CSV 格式日志（结构化列）
SELECT log_time, error_severity, message
FROM postgresql_2024_01_15_csv
WHERE error_severity = 'ERROR'
ORDER BY log_time DESC
LIMIT 20;

-- 查询纯文本日志
SELECT * FROM postgresql_2024_01_15_log LIMIT 10;
```

### 授权非超级用户访问

只有超级用户可以创建该扩展，但可以授予访问权限：

```sql
GRANT pg_monitor TO monitoring_user;
GRANT CREATE ON SCHEMA public TO monitoring_user;
GRANT USAGE ON FOREIGN SERVER log_fdw_server TO monitoring_user;
```

### 函数参考

| 函数 | 描述 |
|----------|-------------|
| `list_postgres_log_files()` | 列出可用日志文件及其大小 |
| `create_foreign_table_for_log_file(table_name, server_name, file_name)` | 从日志文件创建外部表 |
