## 用法

来源：

- [官方上游 README](https://github.com/vexdb-thu/vexdb-lite/blob/315dd536ca1446f1e563594687254647a06790c7/README.md)
- [官方扩展控制文件 (vexdb_lite.control)](https://github.com/vexdb-thu/vexdb-lite/blob/315dd536ca1446f1e563594687254647a06790c7/vexdb_pg/vexdb_lite.control)
- [官方扩展 SQL (vexdb_lite--1.0.sql)](https://github.com/vexdb-thu/vexdb-lite/blob/315dd536ca1446f1e563594687254647a06790c7/vexdb_pg/sql/vexdb_lite--1.0.sql)

`vexdb_lite` — 一个跨平台向量数据库，可以作为插件集成到现有数据库中。使用它来进行相应的向量、模型或检索工作流。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION vexdb_lite;

CREATE TABLE items (
    id  BIGSERIAL PRIMARY KEY,
    vec floatvector(128)
);

INSERT INTO items (vec) VALUES
    ('[0.10, 0.20, 0.30]'),
    ('[0.40, 0.50, 0.60]');
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `array_to_floatvector(double precision[], integer, boolean)` 是一个扩展函数，返回 `floatvector`。
- `array_to_floatvector(integer[], integer, boolean)` 是一个扩展函数，返回 `floatvector`。
- `array_to_floatvector(numeric[], integer, boolean)` 是一个扩展函数，返回 `floatvector`。
- `array_to_floatvector(real[], integer, boolean)` 是一个扩展函数，返回 `floatvector`。
- `cosine_distance(floatvector, floatvector)` 是一个扩展函数，返回 `float8`。
- `floatvector(floatvector, integer, boolean)` 是一个扩展函数，返回 `floatvector`。
- `floatvector_add(floatvector, floatvector)` 是一个扩展函数，返回 `floatvector`。
- `floatvector_cmp(floatvector, floatvector)` 是一个扩展函数，返回 `int4`。
- `floatvector_eq(floatvector, floatvector)` 是一个扩展函数。
- `floatvector_ge(floatvector, floatvector)` 是一个扩展函数。
- `floatvector_gt(floatvector, floatvector)` 是一个扩展函数。
- `floatvector_in(cstring, oid, integer)` 是一个扩展函数，返回 `floatvector`。
- `floatvector_l2_squared_distance(floatvector, floatvector)` 是一个扩展函数，返回 `float8`。
- `floatvector_le(floatvector, floatvector)` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
