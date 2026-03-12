

## 用法

> [sequential_uuids: 顺序 UUID 生成器，提供更好的索引局部性](https://github.com/tvondra/sequential-uuids)

生成具有顺序模式的 UUID，减少索引中的随机 I/O，同时保持足够的随机性以避免冲突。

```sql
CREATE EXTENSION sequential_uuids;
```

### 函数

| 函数 | 描述 |
|---|---|
| `uuid_sequence_nextval(sequence regclass, block_size int DEFAULT 65536, block_count int DEFAULT 65536)` | 基于序列生成顺序 UUID |
| `uuid_time_nextval(interval_length int DEFAULT 60, interval_count int DEFAULT 65536)` | 基于当前时间戳生成顺序 UUID |

### 示例

```sql
CREATE SEQUENCE my_seq;

-- 基于序列的 UUID 生成
SELECT uuid_sequence_nextval('my_seq'::regclass);

-- 基于时间的 UUID 生成（默认参数下约每 45 天循环一次）
SELECT uuid_time_nextval();

-- 用作列的默认值
CREATE TABLE orders (
  id uuid DEFAULT uuid_time_nextval() PRIMARY KEY,
  data text
);

-- 自定义块大小和数量
SELECT uuid_sequence_nextval('my_seq', 256, 65536);
SELECT uuid_time_nextval(120, 32768);
```
