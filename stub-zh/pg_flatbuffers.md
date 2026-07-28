## 用法

来源：

- [官方上游 README](https://github.com/youyuanwu/pg_flatbuffers/blob/a09d4f88c02d0709620abfa0106ceb0c1fd578fa/README.md)
- [官方扩展控制文件 (pg_flatbuffers.control)](https://github.com/youyuanwu/pg_flatbuffers/blob/a09d4f88c02d0709620abfa0106ceb0c1fd578fa/crates/pg_flatbuffers/pg_flatbuffers.control)
- [官方实现源代码](https://github.com/youyuanwu/pg_flatbuffers/blob/a09d4f88c02d0709620abfa0106ceb0c1fd578fa/crates/pg_flatbuffers/src/lib.rs)

`pg_flatbuffers` — pg_flatbuffers: 查询和转换字节列中的 FlatBuffers 载荷。在相应的 SQL 或数据库实用程序工作流中使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_flatbuffers;

-- Register a schema produced by `flatc -b --schema orders.fbs`
INSERT INTO flatbuffers_schemas (name, bfbs)
VALUES ('default', pg_read_binary_file('/tmp/orders.bfbs'));

-- Single-value extraction
SELECT flatbuffers_query('myco.orders.Order:customer.email', payload)
FROM   orders_raw
WHERE  id = 42;

-- Vector fan-out as rows (suitable for joins)
SELECT o.id, sku
FROM   orders_raw o,
       LATERAL flatbuffers_query_multi(
         'myco.orders.Order:items[*].sku', o.payload) AS sku;

-- JSON round-trip
SELECT flatbuffers_to_json('myco.orders.Order', payload) -> 'customer'
FROM   orders_raw;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `flatbuffers_extension_version()` 是一个扩展函数。

### 要求与注意事项

- 元组记录扩展版本 `0.0.1`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
