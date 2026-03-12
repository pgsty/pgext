


## 用法

> [omni_sqlite: 内嵌 SQLite](https://docs.omnigres.org/omni_sqlite/sqlite/)

`omni_sqlite` 扩展将 SQLite 数据库作为 PostgreSQL 的一等数据类型。它是一个模板化扩展。

### 创建 SQLite 数据库

```sql
SELECT 'CREATE TABLE user_config (key text, value text)'::omni_sqlite.sqlite;
```

### 在列中存储 SQLite

```sql
CREATE TABLE customer (
    id   bigserial PRIMARY KEY,
    name text NOT NULL,
    data omni_sqlite.sqlite DEFAULT 'CREATE TABLE user_config (key text, value text);'
);
```

### 执行语句

```sql
UPDATE customer
SET data = omni_sqlite.sqlite_exec(data,
    $$INSERT INTO user_config VALUES ('color', 'blue')$$)
RETURNING *;
```

带绑定参数：

```sql
SELECT omni_sqlite.sqlite_exec('CREATE TABLE tab (val text)',
    'INSERT INTO tab VALUES ($1)', row('hello'));
```

### 查询 SQLite 数据

```sql
SELECT * FROM omni_sqlite.sqlite_query(
    (SELECT data FROM customer),
    'SELECT rowid, key, value FROM user_config')
AS (id bigint, key text, value text);
```

### 序列化

```sql
-- 序列化为 bytea：
SELECT omni_sqlite.sqlite_serialize('CREATE TABLE foo (x)');
-- 从 bytea 反序列化：
SELECT omni_sqlite.sqlite_deserialize(bytea_data);
```

适用于多租户（每行隔离的 SQLite 数据库）和对等同步场景。
