

## 用法

> [file_fdw: 平面文件访问的外部数据包装器](https://www.postgresql.org/docs/current/file-fdw.html)

### 创建服务器

```sql
CREATE EXTENSION file_fdw;

CREATE SERVER file_server FOREIGN DATA WRAPPER file_fdw;
```

### 读取 CSV 文件

```sql
CREATE FOREIGN TABLE csv_data (
  id integer,
  name text,
  value numeric
)
SERVER file_server
OPTIONS (filename '/path/to/data.csv', format 'csv', header 'true');

SELECT * FROM csv_data;
```

### 读取 PostgreSQL CSV 日志

```sql
CREATE FOREIGN TABLE pglog (
  log_time timestamp(3) with time zone,
  user_name text,
  database_name text,
  process_id integer,
  connection_from text,
  session_id text,
  session_line_num bigint,
  command_tag text,
  session_start_time timestamp with time zone,
  virtual_transaction_id text,
  transaction_id bigint,
  error_severity text,
  sql_state_code text,
  message text,
  detail text,
  hint text,
  internal_query text,
  internal_query_pos integer,
  context text,
  query text,
  query_pos integer,
  location text,
  application_name text,
  backend_type text,
  leader_pid integer,
  query_id bigint
)
SERVER file_server
OPTIONS (filename 'log/pglog.csv', format 'csv');
```

### 读取程序输出

```sql
CREATE FOREIGN TABLE process_list (
  pid text,
  command text
)
SERVER file_server
OPTIONS (program 'ps aux | tail -n +2', format 'text', delimiter ' ');
```

### 表选项

| 选项 | 描述 |
|--------|-------------|
| `filename` | 文件路径（相对于数据目录）。除非使用 `program`，否则必填 |
| `program` | 读取其标准输出的 Shell 命令。除非使用 `filename`，否则必填 |
| `format` | 数据格式：`csv`、`text` 或 `binary`（与 COPY 相同） |
| `header` | 文件是否有标题行时为 `true` |
| `delimiter` | 列分隔符字符 |
| `quote` | 引号字符 |
| `escape` | 转义字符 |
| `null` | 表示 NULL 值的字符串 |
| `encoding` | 数据编码 |
| `on_error` | 类型转换时的错误处理 |
| `reject_limit` | 最大容忍错误数 |

### 列选项

| 选项 | 描述 |
|--------|-------------|
| `force_not_null` | 不将列值与空字符串匹配 |
| `force_null` | 将引号值与空字符串匹配并返回 NULL |

```sql
CREATE FOREIGN TABLE films (
  code char(5) NOT NULL,
  title text NOT NULL,
  rating text OPTIONS (force_null 'true')
)
SERVER file_server
OPTIONS (filename '/data/films.csv', format 'csv');
```

file_fdw 是只读的。更改表级选项需要超级用户权限或 `pg_read_server_files` / `pg_execute_server_program` 角色。
