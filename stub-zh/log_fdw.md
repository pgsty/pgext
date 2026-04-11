
## 用法

> [README](https://github.com/aws/postgresql-logfdw)

`log_fdw` 是一个用于通过 SQL 读取 PostgreSQL 日志文件的外部数据包装器。它提供辅助函数，用于列出服务器日志目录中的文件，以及为单个日志文件创建外部表。

### 核心函数

上游 README 定义了两个 SQL 入口：

```sql
create_foreign_table_for_log_file(table_name text, server_name text, log_file_name text)
list_postgres_log_files()
```

`list_postgres_log_files()` 是对 PostgreSQL 核心函数 `pg_ls_logdir()` 的兼容包装。

## 基本流程

先创建扩展和外部服务器：

```sql
CREATE EXTENSION log_fdw;
CREATE SERVER log_fdw_server FOREIGN DATA WRAPPER log_fdw;
```

列出 PostgreSQL 日志目录中的文件：

```sql
SELECT * FROM list_postgres_log_files() ORDER BY 1 DESC LIMIT 10;
```

为 CSV 日志或普通 `.log` 文件创建外部表：

```sql
SELECT * FROM create_foreign_table_for_log_file(
  'postgresql_2022_11_28_csv',
  'log_fdw_server',
  'postgresql-2022-11-28.csv'
);

SELECT * FROM create_foreign_table_for_log_file(
  'postgresql_2022_11_28_log',
  'log_fdw_server',
  'postgresql-2022-11-28.log'
);
```

## 查询

由普通日志文件创建的外部表通常只暴露单个日志行样式的列；CSV 日志文件则会暴露结构化列，例如 `log_time`、`error_severity`、`message` 以及会话元数据。

典型用法如下：

```sql
SELECT * FROM postgresql_2022_11_28_log LIMIT 2;

SELECT log_time, error_severity, message
FROM postgresql_2022_11_28_csv
WHERE error_severity = 'ERROR'
ORDER BY log_time DESC
LIMIT 20;
```

## 权限

只有超级用户可以创建该扩展。上游还说明，超级用户可以通过最小权限授予把访问能力委托给非超级用户，例如：

```sql
CREATE ROLE foo;
GRANT pg_monitor TO foo;
GRANT CREATE ON SCHEMA bar TO foo;
GRANT USAGE ON FOREIGN SERVER log_fdw_server TO foo;
```

当使用 `list_postgres_log_files()` 时，需要 `pg_monitor`，因为底层 `pg_ls_logdir()` 函数需要这个权限。

## 构建说明

该项目可以使用 PGXS 独立构建：

```bash
export USE_PGXS=1
make
make install
```

也可以把源码复制到 PostgreSQL 的 `contrib` 目录中，作为更大发行版的一部分进行构建。
