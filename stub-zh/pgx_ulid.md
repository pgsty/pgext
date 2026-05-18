## 用法

来源：[README](https://github.com/pksunkara/pgx_ulid/blob/master/README.md), [releases](https://github.com/pksunkara/pgx_ulid/releases)

`pgx_ulid` 提供原生 `ulid` 类型、生成器，以及与 `timestamp` 和 `uuid` 的双向 casts。README 记录了二进制存储和 monotonic ULID 支持。pgext catalog 记录 package 和 extension 名均为 `pgx_ulid`，版本 `0.2.3`，覆盖 PostgreSQL 14-18。

### 启用扩展

```sql
CREATE EXTENSION pgx_ulid;
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

README 明确说明该 preload 要求只影响 `gen_monotonic_ulid()`；扩展其他功能无需 preload。

### 将 `ulid` 用作主键

```sql
CREATE TABLE users (
  id ulid NOT NULL DEFAULT gen_ulid() PRIMARY KEY,
  name text NOT NULL
);

SELECT * FROM users
WHERE id = '01ARZ3NDEKTSV4RRFFQ69G5FAV';
```

### Casts 与范围查询

```sql
ALTER TABLE users
ADD COLUMN created_at timestamp GENERATED ALWAYS AS (id::timestamp) STORED;

SELECT * FROM users
WHERE id BETWEEN '2023-09-15'::timestamp::ulid
            AND '2023-09-16'::timestamp::ulid;
```

README 还记录了 `ulid` 与 `uuid` 之间的 casts。

### 注意事项

- Monotonic ULIDs 使用 shared memory 和 LWLock 来保存上一次生成的值。
- README 提到 monotonic generation 理论上可能 overflow 并抛出错误，虽然它认为实践中可忽略。
- 上游 README 也展示 `CREATE EXTENSION ulid`；此 stub 遵循 `db/extension.csv`，其中 package 和 lead extension 都是 `pgx_ulid`。
- GitHub release page 将 `v0.2.3` 列为最新，且只标为 `Release 0.2.3`，没有单独用户侧 release notes。
