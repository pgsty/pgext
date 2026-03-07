

## 用法

> https://github.com/frectonz/pglite-fusion/blob/main/README.md


以下是一些使用示例。

```sql
-- 加载 PG 扩展
CREATE EXTENSION pglite_fusion;

-- 创建一个包含 SQLite 列的表
CREATE TABLE people (
                        name     TEXT NOT NULL,
                        database SQLITE DEFAULT init_sqlite('CREATE TABLE todos (task TEXT)')
);

-- 向 people 表中插入一行数据
INSERT INTO people VALUES ('frectonz');

-- 为 "frectonz" 创建一条待办事项
UPDATE people
SET database = execute_sqlite(
        database,
        'INSERT INTO todos VALUES (''solve multitenancy'')'
               )
WHERE name = 'frectonz';

-- 为 "frectonz" 创建一条待办事项
UPDATE people
SET database = execute_sqlite(
        database,
        'INSERT INTO todos VALUES (''buy milk'')'
               )
WHERE name = 'frectonz';

-- 查询 frectonz 的信息
SELECT
    name,
    (
        SELECT json_agg(get_sqlite_text(sqlite_row, 0))
        FROM query_sqlite(
                database,
                'SELECT * FROM todos'
             )
    ) AS todos
FROM
    people
WHERE
    name = 'frectonz';
```


--------

## API 文档


### `empty_sqlite`

创建一个空的 SQLite 数据库，并将其作为二进制对象返回。可用于在 PostgreSQL 列中初始化一个空的 SQLite 数据库。

#### 使用示例：

```sql
SELECT empty_sqlite();
```

-----

### `query_sqlite`

对存储为二进制对象的 SQLite 数据库执行 SQL 查询，并将结果以 JSON 编码的行集合形式返回。该函数适用于查询存储在 PostgreSQL 列中的 SQLite 数据库。

#### 参数：

- `sqlite`：要查询的 SQLite 数据库，以二进制对象形式存储。
- `query`：要在 SQLite 数据库上执行的 SQL 查询字符串。

#### 使用示例：

```sql
SELECT * FROM query_sqlite(
        database,
        'SELECT * FROM todos'
              );
```

-----

### `execute_sqlite`

对存储为二进制对象的 SQLite 数据库执行 SQL 语句（如 `INSERT`、`UPDATE` 或 `DELETE`）。更新后的 SQLite 数据库将作为二进制对象返回，以便后续继续操作。

#### 参数：

- `sqlite`：要执行 SQL 语句的 SQLite 数据库，以二进制对象形式存储。
- `query`：要在 SQLite 数据库上执行的 SQL 语句。

##### 使用示例：

```sql
UPDATE people
SET database = execute_sqlite(
        database,
        'INSERT INTO todos VALUES (''solve multitenancy'')'
               )
WHERE name = 'frectonz';
```

-----

### `init_sqlite`

创建一个 SQLite 数据库，并在其上执行指定的初始化查询。可用于初始化一个已包含所需表结构的 SQLite 数据库。

#### 参数：

- `query`：要在 SQLite 数据库上执行的 SQL 语句。

##### 使用示例：

```sql

CREATE TABLE people (
                        name     TEXT NOT NULL,
                        database SQLITE DEFAULT init_sqlite('CREATE TABLE todos (task TEXT)')
);
```

-----

### `get_sqlite_text`
从 `query_sqlite` 返回的行中提取指定列的文本值。使用此函数可从查询结果中获取文本类型的值。

#### 参数：

- `sqlite_row`：`query_sqlite` 返回的结果行。
- `index`：要从行中提取的列索引。

#### 使用示例：

```sql
SELECT get_sqlite_text(sqlite_row, 0)
FROM query_sqlite(database, 'SELECT * FROM todos');
```

----

### `get_sqlite_integer`

从 `query_sqlite` 返回的行中提取指定列的整数值。使用此函数可从查询结果中获取整数类型的值。

#### 参数：

- `sqlite_row`：`query_sqlite` 返回的结果行。
- `index`：要从行中提取的列索引。

#### 使用示例：

```sql
SELECT get_sqlite_integer(sqlite_row, 1)
FROM query_sqlite(database, 'SELECT * FROM todos');
```

----

### `get_sqlite_real`

从 `query_sqlite` 返回的行中提取指定列的实数（浮点数）值。使用此函数可从查询结果中获取实数类型的值。

#### 参数：

- `sqlite_row`：`query_sqlite` 返回的结果行。
- `index`：要从行中提取的列索引。

#### 使用示例：

```sql
SELECT get_sqlite_real(sqlite_row, 2)
FROM query_sqlite(database, 'SELECT * FROM todos');
```