## 用法

来源：[README](https://github.com/pksunkara/pgx_ulid/blob/master/README.md)，[releases](https://github.com/pksunkara/pgx_ulid/releases)

`pgx_ulid` 提供原生 `ulid` 类型、生成函数，以及与 `timestamp` 和 `uuid` 之间的类型转换。README 说明它以二进制形式存储，并支持 monotonic ULID。

### 启用扩展

```sql
CREATE EXTENSION ulid;
-- or CREATE EXTENSION pgx_ulid; if installed manually under that name
```

### 生成 ULID

```sql
SELECT gen_ulid();
SELECT gen_monotonic_ulid();
```

`gen_monotonic_ulid()` 需要：

```conf
shared_preload_libraries = 'pgx_ulid'
```

README 明确说明，这个 preload 要求只影响 `gen_monotonic_ulid()`；扩展的其余部分无需预加载即可使用。

### 将 `ulid` 用作主键

```sql
CREATE TABLE users (
  id ulid NOT NULL DEFAULT gen_ulid() PRIMARY KEY,
  name text NOT NULL
);

SELECT * FROM users
WHERE id = '01ARZ3NDEKTSV4RRFFQ69G5FAV';
```

### 类型转换与范围查询

```sql
ALTER TABLE users
ADD COLUMN created_at timestamp GENERATED ALWAYS AS (id::timestamp) STORED;

SELECT * FROM users
WHERE id BETWEEN '2023-09-15'::timestamp::ulid
            AND '2023-09-16'::timestamp::ulid;
```

README 还记录了 `ulid` 与 `uuid` 之间的双向 cast。

### 注意事项

- Monotonic ULID 通过 shared memory 和 LWLock 维护最近一次生成的值。
- README 提到 monotonic 生成在理论上可能发生溢出并报错，但实践中可视为极小概率事件。
- 截至 2026-04-19，上游当前版本为 `v0.2.3`，但没有单独发布该版本的面向用户说明。
