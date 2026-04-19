## 用法

来源：[README](https://github.com/ringsaturn/pg-tzf/blob/main/README.md)，[releases](https://github.com/ringsaturn/pg-tzf/releases)，[Cargo.toml](https://github.com/ringsaturn/pg-tzf/blob/main/Cargo.toml)

`tzf` 是一个 PostgreSQL 扩展，用于根据经纬度快速查找时区名称。

### 创建扩展

```sql
CREATE EXTENSION tzf;
```

上游项目按 PostgreSQL major version 分别打包构建产物。当前源码元数据声明的扩展版本是 `0.2.4`。

### 函数

#### 坐标查找

```sql
SELECT tzf_tzname(116.3883, 39.9289) AS timezone;
```

#### 批量坐标查找

```sql
SELECT unnest(
  tzf_tzname_batch(
    ARRAY[-74.0060, -118.2437, 139.6917],
    ARRAY[40.7128, 34.0522, 35.6895]
  )
) AS timezones;
```

#### 点查找

```sql
SELECT tzf_tzname_point(point(-74.0060, 40.7128)) AS timezone;
```

#### 批量点查找

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

- 上游 README 记录其支持 PostgreSQL 14 到 18 的构建。
- 预编译发布包内包含 `tzf.so`、`tzf.control` 和 `tzf--<version>.sql`。
- 当前 README 仍指向 `sql/tzf.sql` 中的完整 schema，并附带这四个查找函数的 benchmark 数据。
