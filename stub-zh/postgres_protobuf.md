## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/mpartel/postgres-protobuf/blob/dd1abe9f1d915656eb149eee1f84490137ec7d1d/README.md)
- [0.2 版本 SQL 对象](https://github.com/mpartel/postgres-protobuf/blob/dd1abe9f1d915656eb149eee1f84490137ec7d1d/postgres_protobuf--0.2.sql)
- [C++ 实现](https://github.com/mpartel/postgres-protobuf/blob/dd1abe9f1d915656eb149eee1f84490137ec7d1d/postgres_protobuf.cpp)

`postgres_protobuf` 存储 Protocol Buffer 模式，并查询以 `bytea` 保存的 protobuf 值。它可以选择单个字段、数组或行集，还能在消息与其 JSON 表示之间转换。

```sql
CREATE EXTENSION postgres_protobuf;
INSERT INTO protobuf_file_descriptor_sets (name, file_descriptor_set)
VALUES ('default', $1::bytea);
SELECT protobuf_query(
  'example.User:profile.email',
  protobuf_payload
)
FROM event_store;
```

应使用 `protoc` 生成包含导入项的描述符集，并在查询前把它插入且提交。相关函数包括 `protobuf_query_array`、`protobuf_query_multi`、`protobuf_to_json_text` 和 `protobuf_from_json_text`。

上游警告不要接收不可信查询或未经验证的 protobuf 字节。大型或递归 map 密集消息可能在 PostgreSQL 普通内存核算之外消耗大量 C++ 堆；每次查询都会扫描消息，模式只缓存一个事务，依赖模式的查询函数也不能安全用作索引表达式。Proto3 缺失/默认字段不会返回结果，已弃用 group 不受支持，已复核实现仅测试 AMD64。应限制输入，在数据库外验证并重新序列化，限制描述符变更，并基准内存和延迟。
