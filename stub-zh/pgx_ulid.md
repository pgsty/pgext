

## 用法

> [pgx_ulid: PostgreSQL 的 ULID 类型与方法](https://github.com/pksunkara/pgx_ulid)

```sql
CREATE EXTENSION pgx_ulid;
```

### ULID 类型

该扩展提供了原生的 `ulid` 类型——一种 26 个字符、按字典序可排序的标识符，以二进制形式存储。

### 函数

| 函数 | 描述 |
|---|---|
| `gen_ulid()` | 生成一个新的 ULID |
| `gen_monotonic_ulid()` | 生成单调递增的 ULID（需要 `shared_preload_libraries`） |

### 类型转换

- `ulid::timestamp` -- 从 ULID 中提取创建时间
- `timestamp::ulid` -- 从时间戳生成 ULID（随机部分为零）
- `ulid::uuid` / `uuid::ulid` -- 在 ULID 和 UUID 之间互相转换

### 示例

```sql
-- 将 ULID 用作主键
CREATE TABLE users (
  id ulid NOT NULL DEFAULT gen_ulid() PRIMARY KEY,
  name text NOT NULL
);

-- 通过文本表示进行查询
SELECT * FROM users WHERE id = '01ARZ3NDEKTSV4RRFFQ69G5FAV';

-- 从 ULID 中提取时间戳
ALTER TABLE users
ADD COLUMN created_at timestamp GENERATED ALWAYS AS (id::timestamp) STORED;

-- 按日期范围查询
SELECT * FROM users
WHERE id BETWEEN '2023-09-15'::timestamp::ulid AND '2023-09-16'::timestamp::ulid;
```
