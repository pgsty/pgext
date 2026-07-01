


## 用法

- 来源：[pg_uuid_v8 README](https://github.com/ineron/pg_uuid_v8)，[SQL definitions](https://github.com/ineron/pg_uuid_v8/blob/main/pg_uuid_v8--1.0.sql)，[control file](https://github.com/ineron/pg_uuid_v8/blob/main/pg_uuid_v8.control)

`pg_uuid_v8` 生成外观类似 UUID v4 的 UUID，同时在其中嵌入加密后的微秒级时间戳，便于提取、排序和范围谓词查询。SQL 文件同时暴露了 `uuid_stego_*` 名称和 `uuid_v8_*` 便捷别名。

### 生成 UUID

```sql
CREATE EXTENSION pg_uuid_v8;

SELECT uuid_v8_set_seed('replace-with-a-secret-seed');
SELECT uuid_v8_generate();
```

等价的底层生成函数是：

```sql
SELECT uuid_stego_generate();
```

插入事件时，可以把它用作默认表达式：

```sql
CREATE TABLE events (
  id uuid PRIMARY KEY DEFAULT uuid_v8_generate(),
  data jsonb,
  created_at timestamptz DEFAULT now()
);
```

### 提取并查询隐藏时间戳

提取嵌入的时间戳，返回 Unix epoch 以来的微秒数：

```sql
SELECT uuid_v8_extract_timestamp(id)
FROM events
ORDER BY uuid_v8_extract_timestamp(id)
LIMIT 10;
```

README 建议为基于时间的查询创建函数索引：

```sql
CREATE INDEX events_uuid_v8_time_idx
ON events USING btree (uuid_v8_extract_timestamp(id));

SELECT *
FROM events
WHERE uuid_v8_extract_timestamp(id)
      BETWEEN timestamp_to_stego_time('2026-01-01'::timestamptz)
          AND timestamp_to_stego_time(now())
ORDER BY uuid_v8_extract_timestamp(id);
```

辅助函数可在时间戳和整数时间戳格式之间转换：

```sql
SELECT timestamp_to_stego_time(now());
SELECT stego_time_to_timestamp(uuid_v8_extract_timestamp(id))
FROM events;
```

### 范围辅助函数与操作符

SQL 定义包含直接的范围辅助函数：

```sql
SELECT *
FROM events
WHERE uuid_stego_in_range(
  id,
  now() - interval '24 hours',
  now()
);
```

它还为 `uuid` 定义了感知时间戳的比较函数和操作符：

- `uuid_stego_compare(uuid, uuid)` 和 `uuid_v8_compare(uuid, uuid)`。
- `uuid_stego_lt`、`uuid_stego_le`、`uuid_stego_gt`、`uuid_stego_ge`。
- 操作符 `<`、`<=`、`>` 和 `>=` 会按隐藏时间戳比较 UUID。

### Seed 与加密模式

设置并查看 seed：

```sql
SELECT uuid_v8_set_seed('replace-with-a-secret-seed');
SELECT uuid_v8_get_seed();
```

可用加密模式为 `XOR`、`AES128` 和 `AES256`：

```sql
SELECT uuid_v8_get_encryption_mode();
SELECT uuid_v8_set_encryption_mode('AES128');
SELECT uuid_v8_set_encryption_mode('XOR');
```

如需持久化默认值，README 记录了 `uuid_v8.encryption_mode` GUC：

```sql
ALTER SYSTEM SET uuid_v8.encryption_mode = 'AES128';
SELECT pg_reload_conf();
```

### 注意事项

- seed 必须保密；解释隐藏时间戳时需要用到它。
- 使用某个 seed 和加密模式生成的 UUID，必须用相同设置解码。
- 基于提取时间戳的函数索引会增加存储和更新开销，但这是高效时间范围谓词的预期路径。
- 本地 Pigsty 元数据将该扩展固定在 `public` schema，使 UUID 比较操作符的 commutator 能在 PostgreSQL 17 和 18 上解析；如果在非 Pigsty 构建中使用其他 schema，应显式测试这些操作符。
