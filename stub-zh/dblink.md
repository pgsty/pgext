

## 用法

> [dblink: 从数据库内部连接到其他 PostgreSQL 数据库](https://www.postgresql.org/docs/current/dblink.html)

### 连接到远程数据库

```sql
CREATE EXTENSION dblink;

-- 未命名连接（只允许一个）
SELECT dblink_connect('dbname=remotedb host=remotehost options=-csearch_path=');

-- 命名连接（允许多个）
SELECT dblink_connect('myconn', 'dbname=remotedb host=remotehost');
```

或通过外部服务器定义连接：

```sql
CREATE SERVER remote_srv FOREIGN DATA WRAPPER dblink_fdw
  OPTIONS (hostaddr '192.168.1.10', dbname 'remotedb');
CREATE USER MAPPING FOR local_user SERVER remote_srv
  OPTIONS (user 'remote_user', password 'secret');

SELECT dblink_connect('myconn', 'remote_srv');
```

### 查询远程数据库

```sql
-- 临时连接
SELECT * FROM dblink(
  'dbname=remotedb host=remotehost',
  'SELECT id, name, value FROM remote_table'
) AS t(id int, name text, value numeric);

-- 使用命名连接
SELECT * FROM dblink(
  'myconn',
  'SELECT id, name FROM remote_table WHERE status = 1'
) AS t(id int, name text);
```

必须始终在 `AS` 子句中指定列定义列表。

### 执行命令（无结果集）

```sql
-- 在远程数据库上执行 INSERT、UPDATE、DELETE、DDL
SELECT dblink_exec('myconn', 'INSERT INTO remote_table VALUES (1, ''test'', 42)');
SELECT dblink_exec('myconn', 'UPDATE remote_table SET value = 100 WHERE id = 1');
SELECT dblink_exec('myconn', 'DELETE FROM remote_table WHERE id = 1');
```

返回命令状态字符串（例如 `INSERT 0 1`）。

### 基于游标的访问

```sql
SELECT dblink_open('myconn', 'mycursor', 'SELECT * FROM large_table');
SELECT * FROM dblink_fetch('myconn', 'mycursor', 100) AS t(id int, data text);
SELECT dblink_close('myconn', 'mycursor');
```

### 连接管理

```sql
SELECT dblink_get_connections();    -- 列出打开的命名连接
SELECT dblink_disconnect('myconn'); -- 关闭命名连接
```

### 创建视图以方便使用

```sql
CREATE VIEW remote_data AS
  SELECT * FROM dblink(
    'dbname=remotedb host=remotehost',
    'SELECT id, name, value FROM data_table'
  ) AS t(id int, name text, value numeric);

SELECT * FROM remote_data WHERE value > 100;
```

### 关键函数

| 函数 | 描述 |
|----------|-------------|
| `dblink_connect(connstr)` | 打开未命名持久连接 |
| `dblink_connect(name, connstr)` | 打开命名持久连接 |
| `dblink_disconnect(name)` | 关闭命名连接 |
| `dblink(connstr, sql)` | 执行查询，返回行 |
| `dblink_exec(connstr, sql)` | 执行命令，返回状态 |
| `dblink_open(name, cursor, sql)` | 在远程数据库上打开游标 |
| `dblink_fetch(name, cursor, count)` | 从远程游标获取行 |
| `dblink_close(name, cursor)` | 关闭远程游标 |
| `dblink_get_connections()` | 列出所有打开的命名连接 |
