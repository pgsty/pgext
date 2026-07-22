## 用法

来源：

- [官方 0.1.7 版 README](https://github.com/kaznak/pgx_uuidv7/blob/3b3828ae89d634c2e828cc872a1ac96394e6cf60/README.md)
- [官方 control 文件](https://github.com/kaznak/pgx_uuidv7/blob/3b3828ae89d634c2e828cc872a1ac96394e6cf60/pgx_uuidv7.control)
- [官方 0.1.7 发布页](https://github.com/kaznak/pgx_uuidv7/releases/tag/v0.1.7)

`pgx_uuidv7` 生成按时间有序的 UUID 版本 7 值，并提取其中嵌入的时间戳或版本号。它适合作为 UUID 主键，相比随机 UUIDv4 具有更好的插入局部性，同时仍使用标准 `uuid` 存储类型。版本 `0.1.7` 面向 PostgreSQL 17 与 18。

### 核心流程

创建扩展，把当前时间生成器用作字段默认值：

```sql
CREATE EXTENSION pgx_uuidv7;

CREATE TABLE events (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v7_now(),
  payload jsonb NOT NULL
);

INSERT INTO events (payload) VALUES ('{"kind":"signup"}');

SELECT id, uuid_get_version(id), uuid_to_timestamptz(id)
FROM events;
```

UUID 值还可以直接转换为 `timestamptz`，用于时间过滤。

### API

- `uuid_generate_v7_now()`：按当前时间生成 UUIDv7。
- `uuid_generate_v7(timestamptz)`：按指定时间戳生成。
- `uuid_generate_v7_at_interval(interval)`：按当前时间的偏移量生成。
- `uuid_get_version(uuid)`：提取 UUID 版本。
- `uuid_to_timestamptz(uuid)`：提取 UUIDv7 时间戳。

在 PostgreSQL 18 之前的构建中，还提供兼容别名 `uuidv7()`、`uuidv7(interval)`、`uuid_extract_version(uuid)`、`uuid_extract_timestamp(uuid)`。PostgreSQL 18 构建会刻意省略它们，以避免与原生函数冲突。

### 注意事项

UUIDv7 中嵌入的时间戳是可观察的；不要把该标识符当作隐藏创建时间的手段。时间顺序可改善局部性，但不能替代业务时间戳，也不保证跨不同时钟的全局生成顺序。从 PostgreSQL 17 迁移到 18 时，应验证类型转换与别名，因为 PostgreSQL 18 原生接口会取代这些兼容名称。
