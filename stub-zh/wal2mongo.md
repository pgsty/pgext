

## 用法

> [wal2mongo: MongoDB 的 PostgreSQL 逻辑解码输出插件](https://github.com/HighgoSoftware/wal2mongo)

一个逻辑解码输出插件，将 PostgreSQL WAL 变更格式化为 MongoDB 兼容命令，实现从 PostgreSQL 到 MongoDB 的数据复制。

### 配置

在 `postgresql.conf` 中：

```ini
wal_level = logical
max_replication_slots = 10
```

### 使用 SQL 函数

```sql
-- 创建逻辑复制槽
SELECT * FROM pg_create_logical_replication_slot('w2m_slot', 'wal2mongo');

-- 执行 DML 操作
CREATE TABLE books (id SERIAL PRIMARY KEY, title VARCHAR(100), author VARCHAR(100));
INSERT INTO books (id, title, author) VALUES (123, 'My Book', 'Author');

-- 查看变更（MongoDB 格式）
SELECT * FROM pg_logical_slot_peek_changes('w2m_slot', NULL, NULL);
-- 输出：db.books.insertOne( { id:123, title:"My Book", author:"Author" } )

-- 消费变更
SELECT * FROM pg_logical_slot_get_changes('w2m_slot', NULL, NULL);

-- 删除槽
SELECT pg_drop_replication_slot('w2m_slot');
```

### 使用 pg_recvlogical

```bash
pg_recvlogical -d postgres --slot w2m_slot --create-slot -P wal2mongo
pg_recvlogical -d postgres --slot w2m_slot --start -f -
```

### 复制到 MongoDB

输出可以直接在 MongoDB shell 中应用：

```javascript
// 复制 pg_logical_slot_get_changes 的输出
db.books.insertOne( { id:123, title:"My Book", author:"Author" } )
```

或保存为 `.js` 文件并导入：

```bash
mongo < changes.js
```

### 输出格式

- **INSERT**：`db.<table>.insertOne( { <columns> } )`
- **UPDATE**：`db.<table>.updateOne( { <key> }, { $set: { <changes> } } )`
- **DELETE**：`db.<table>.deleteOne( { <key> } )`

表需要主键或副本标识才能捕获 UPDATE/DELETE 操作。
