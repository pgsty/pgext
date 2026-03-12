

## 用法

> [pg_uuidv7: 在 PostgreSQL 中创建有效的版本 7 UUID](https://github.com/fboulnois/pg_uuidv7)

```sql
CREATE EXTENSION pg_uuidv7;
```

### 函数

| 函数 | 描述 |
|---|---|
| `uuid_generate_v7()` | 生成一个新的 UUIDv7 |
| `uuid_v7_to_timestamptz(uuid)` | 从 UUIDv7 中提取时间戳 |
| `uuid_timestamptz_to_v7(timestamptz [, bool])` | 将时间戳转换为 UUIDv7（第二个参数设为 `true` 则将随机位清零） |

### 示例

```sql
-- 生成一个 UUIDv7
SELECT uuid_generate_v7();
-- 018570bb-4a7d-7c7e-8df4-6d47afd8c8fc

-- 从 UUIDv7 中提取时间戳
SELECT uuid_v7_to_timestamptz('018570bb-4a7d-7c7e-8df4-6d47afd8c8fc');
-- 2023-01-02 04:26:40.637+00

-- 将时间戳转换为 UUIDv7
SELECT uuid_timestamptz_to_v7('2023-01-02 04:26:40.637+00');
-- 018570bb-4a7d-7630-a5c4-89b795024c5d

-- 用于日期范围查询时，将随机位清零
SELECT uuid_timestamptz_to_v7('2023-01-02 04:26:40.637+00', true);
-- 018570bb-4a7d-7000-8000-000000000000

-- 用作主键
CREATE TABLE events (
  id uuid NOT NULL DEFAULT uuid_generate_v7() PRIMARY KEY,
  data text
);
```
