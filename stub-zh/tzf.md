


## 用法

来源：[README](https://github.com/ringsaturn/pg-tzf/blob/v0.3.0/README.md)、[v0.3.0 release](https://github.com/ringsaturn/pg-tzf/releases/tag/v0.3.0)

`tzf` 是一个 PostgreSQL 扩展，用于根据经纬度坐标快速查找 timezone。pgext catalog 将 package `pg_tzf` 映射到 extension `tzf`，并记录版本 `0.3.0`，覆盖 PostgreSQL 14-18。

### 创建扩展

```sql
CREATE EXTENSION tzf;
```

上游项目按 PostgreSQL major version 提供一个构建产物。

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
- 上游 v0.3.0 增加基于 city-json 的 timezone lookup query 测试和依赖更新；SQL API 仍是上面四个 lookup functions。
