

## 用法

> [envvar: 从 PostgreSQL 访问环境变量](https://github.com/theory/pg-envvar)

提供一个用于读取数据库服务器环境变量的函数。

### 函数

```sql
get_env(name TEXT) RETURNS TEXT
```

返回数据库服务器上指定环境变量的值，如果未设置则返回 NULL。

### 示例

```sql
SELECT get_env('PGTZ');
-- UTC

SELECT get_env('HOME');
-- /var/lib/postgresql

SELECT get_env('PATH');
-- /usr/local/sbin:/usr/local/bin:...

SELECT get_env('NONEXISTENT');
-- NULL
```

### 模式支持

可以将扩展安装到指定的模式中：

```sql
CREATE SCHEMA env;
CREATE EXTENSION envvar SCHEMA env;

SELECT env.get_env('PGTZ');
```
