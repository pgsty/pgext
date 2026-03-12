

## 用法

> [plsh: PL/sh Shell 过程语言](https://github.com/petere/plsh)

`plsh` 允许将 PostgreSQL 函数编写为 Shell 脚本。函数体必须以 shebang 行开头，指定解释器。

```sql
CREATE EXTENSION plsh;
```

### 创建 Shell 函数

```sql
CREATE FUNCTION concat(text, text) RETURNS text AS '
#!/bin/sh
echo "$1$2"
' LANGUAGE plsh;

SELECT concat('hello ', 'world');  -- 'hello world'
```

### 参数传递

函数参数作为位置 Shell 变量（`$1`、`$2` 等）传递：

```sql
CREATE FUNCTION file_line_count(filename text) RETURNS int AS '
#!/bin/sh
wc -l < "$1"
' LANGUAGE plsh;
```

### 返回值

- 标准输出成为返回值（尾部换行符被去除）
- 空输出返回 NULL
- 标准错误输出会导致函数中止，并将其作为错误消息
- 非零退出状态会触发错误

```sql
CREATE FUNCTION system_uptime() RETURNS text AS '
#!/bin/sh
uptime
' LANGUAGE plsh;
```

### 数据库访问

不支持直接 SPI 访问，但由于 libpq 环境变量已预配置，可以使用 `psql`：

```sql
CREATE FUNCTION query_db(x int) RETURNS text AS $$
#!/bin/sh
psql -At -c "SELECT name FROM users WHERE id = $1"
$$  LANGUAGE plsh;
```

### 触发器函数

触发器上下文通过环境变量提供：

| 变量 | 说明 |
|------|------|
| `PLSH_TG_NAME` | 触发器名称 |
| `PLSH_TG_WHEN` | `BEFORE`、`INSTEAD OF` 或 `AFTER` |
| `PLSH_TG_LEVEL` | `ROW` 或 `STATEMENT` |
| `PLSH_TG_OP` | `DELETE`、`INSERT`、`UPDATE` 或 `TRUNCATE` |
| `PLSH_TG_TABLE_NAME` | 目标表名 |
| `PLSH_TG_TABLE_SCHEMA` | 目标表所在模式 |

事件触发器变量：`PLSH_TG_EVENT`、`PLSH_TG_TAG`。

### 内联执行

```sql
DO E'#!/bin/sh\necho "running shell command"' LANGUAGE plsh;
```

### 安全性

`plsh` 不应声明为 `TRUSTED`，因为 Shell 脚本在 PostgreSQL 用户权限下具有完全的操作系统级访问权限。只有超级用户应创建 `plsh` 函数。
