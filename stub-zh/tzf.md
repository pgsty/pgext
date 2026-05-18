## 用法

来源：[README](https://github.com/ringsaturn/pg-tzf/blob/main/README.md), [releases](https://github.com/ringsaturn/pg-tzf/releases)

`tzf` 是一个 PostgreSQL 扩展，用于根据经纬度坐标快速查找 timezone。pgext catalog 将 package `pg_tzf` 映射到 extension `tzf`，并记录版本 `0.2.4`，覆盖 PostgreSQL 14-18。

### 创建扩展

```sql
CREATE EXTENSION tzf;
```

上游项目按 PostgreSQL major version 提供一个构建产物。其 release page 现在在 `v0.2.4` 之后列出 `v0.3.0`；此 stub 仍让版本和 package 名与 `db/extension.csv` 对齐。

### 函数

坐标查询：

```sql
SELECT tzf_tzname(116.3883, 39.9289) AS timezone;
```

批量坐标查询：

```sql
SELECT unnest(
  tzf_tzname_batch(
    ARRAY[-74.0060, -118.2437, 139.6917],
    ARRAY[40.7128, 34.0522, 35.6895]
  )
) AS timezones;
```

Point 查询：

```sql
SELECT tzf_tzname_point(point(-74.0060, 40.7128)) AS timezone;
```

批量 point 查询：

```sql
SELECT unnest(
  tzf_tzname_batch_points(
    ARRAY[
      point(-74.0060, 40.7128),
      point(-118.2437, 34.0522),
      point(139.6917, 35.6895)
    ]
  )
) AS timezones;
```

### 说明

- 上游 README 记录支持 PostgreSQL 14 到 18 builds。
- 预构建 release tarballs 包含 `tzf.so`、`tzf.control` 和 `tzf--<version>.sql`。
- 当前 README 仍指向 `sql/tzf.sql` 中的完整 schema，并包含上面四个查询函数的 benchmark 数字。
