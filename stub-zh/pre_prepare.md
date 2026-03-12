

## 用法

> [pre_prepare: 在服务器端预准备你的语句](https://github.com/dimitri/preprepare)

pre_prepare 在连接建立时自动准备 SQL 语句，使客户端可以直接使用 `EXECUTE` 而无需先调用 `PREPARE`。

### 设置

在 `postgresql.conf` 中配置：

```
preprepare.relation = 'preprepare.statements'
preprepare.at_init = on    -- 连接时自动准备（需要 local_preload_libraries）
```

创建存储语句的表：

```sql
CREATE TABLE preprepare.statements (name text PRIMARY KEY, statement text);
```

插入语句（包含完整的 `PREPARE` 语法）：

```sql
INSERT INTO preprepare.statements VALUES ('test', 'prepare test as select 1');
```

### 函数

**`prepare_all()`** -- 从配置的关系中准备所有语句：

```sql
SELECT prepare_all();
```

**`prepare_all('schema.table')`** -- 从指定表准备语句：

```sql
SELECT prepare_all('public.expensive_planning');
```

**`discard()`** -- 类似 `DISCARD ALL` 但不执行 `DEALLOCATE ALL`（保留已准备的语句）：

```sql
SELECT discard();
```

### 配合 PgBouncer 使用

设置 `connect_query` 在每次服务器连接时自动准备：

```ini
[databases]
foo = port=5432 connect_query='SELECT prepare_all();'
```

避免使用 `DISCARD ALL` 作为 `reset_query`（它会释放已准备的语句）。
