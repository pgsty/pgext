

## 用法

> [decoder_raw: 原始 SQL 格式的逻辑复制输出插件](https://github.com/michaelpq/pg_plugins/blob/main/decoder_raw/)

一个逻辑解码输出插件，将 WAL 变更转换为原始 SQL 语句。是 Michael Paquier 的 pg_plugins 集合的一部分。

### 配置

在 `postgresql.conf` 中：

```ini
wal_level = logical
max_replication_slots = 10
max_wal_senders = 10
```

### 使用 SQL 函数

```sql
-- 创建逻辑复制槽
SELECT * FROM pg_create_logical_replication_slot('raw_slot', 'decoder_raw');

-- 执行 DML 操作
INSERT INTO my_table VALUES (1, 'hello');
UPDATE my_table SET val = 'world' WHERE id = 1;
DELETE FROM my_table WHERE id = 1;

-- 获取原始 SQL 变更
SELECT data FROM pg_logical_slot_get_changes('raw_slot', NULL, NULL);
-- 输出：
-- INSERT INTO public.my_table (id, val) VALUES (1, 'hello');
-- UPDATE public.my_table SET val = 'world' WHERE id = 1;
-- DELETE FROM public.my_table WHERE id = 1;

-- 删除槽
SELECT pg_drop_replication_slot('raw_slot');
```

### 使用 pg_recvlogical

```bash
# 创建槽
pg_recvlogical -d postgres --slot raw_slot --create-slot -P decoder_raw

# 流式传输 SQL 语句变更
pg_recvlogical -d postgres --slot raw_slot --start -f -

# 删除槽
pg_recvlogical -d postgres --slot raw_slot --drop-slot
```

### 关键特性

- 输出变更为可执行的 SQL 语句（INSERT、UPDATE、DELETE）
- 适用于调试逻辑解码或在另一个数据库上重放变更
- 表应设置 REPLICA IDENTITY 以获得正确的 UPDATE/DELETE 输出
- 设计为自定义逻辑解码插件的轻量级模板
