## 用法

来源：

- [官方扩展控制文件](https://github.com/ycui1984/pg_thrift/blob/1de265a661ab3c61aa593d7e99d3a313024170fc/pg_thrift.control)
- [官方 README](https://github.com/ycui1984/pg_thrift/blob/1de265a661ab3c61aa593d7e99d3a313024170fc/README.md)
- [官方扩展 SQL](https://github.com/ycui1984/pg_thrift/blob/1de265a661ab3c61aa593d7e99d3a313024170fc/pg_thrift--1.0.sql)

`pg_thrift` 1.0 可以从存为 `bytea` 的 Apache Thrift Binary Protocol 或 Compact Protocol 数据中读取指定字段。它适合检查已知的 Thrift 载荷，并针对标量字段建立表达式索引；它不会加载 IDL 模式、生成 Thrift 消息，也不提供 Thrift RPC 客户端或服务器。

### 核心流程

安装扩展，把编码后的结构体存为 `bytea`，然后使用数字形式的 Thrift 字段 ID 调用对应协议的 getter。容器 getter 会把编码元素返回为 `bytea[]`；再将这些元素传给匹配的 `parse_thrift_binary_*` 或 `parse_thrift_compact_*` 函数。

```sql
CREATE EXTENSION pg_thrift;

CREATE TABLE thrift_events (
    payload bytea NOT NULL
);

CREATE INDEX thrift_events_name_idx
    ON thrift_events (thrift_binary_get_string(payload, 1));

SELECT thrift_binary_get_int32(payload, 2) AS event_id,
       thrift_binary_get_string(payload, 1) AS event_name
FROM thrift_events
WHERE thrift_binary_get_string(payload, 1) = 'created';
```

扩展还定义了 `thrift_binary` 类型。它的文本输入是类似 `{"type":"int16","value":60}` 的小型 JSON 对象；这个便捷类型只支持 Binary Protocol。

### 对象索引

- Binary 结构体 getter：`thrift_binary_get_bool`、`thrift_binary_get_byte`、`thrift_binary_get_double`、`thrift_binary_get_int16`、`thrift_binary_get_int32`、`thrift_binary_get_int64`、`thrift_binary_get_string`，以及结构体/列表/集合/映射 getter。
- Compact 结构体 getter：对应的 `thrift_compact_get_*` 函数族。
- 标量/容器解析器：处理容器 getter 所返回编码切片的 `parse_thrift_binary_*` 与 `parse_thrift_compact_*` 函数族。
- 便捷转换：`jsonb_to_thrift_binary(jsonb)` 返回 Binary Protocol 字节；`get_thrift_binary_type` 与 `get_thrift_binary_value` 检查自定义类型。

字段 ID 和线类型必须与应用的 IDL 一致。扩展没有模式注册表，也不能验证某个字段编号是否具有预期的业务含义。

### 兼容性与注意事项

上游代码停留在 2018 年，目录适用范围是 PostgreSQL 10–11。在更新的服务器上采用前，必须固定并测试确切源码。C 解析器直接处理输入字节；接受不可信载荷之前，应验证协议、消息大小、嵌套层次及畸形输入行为。

所有 SQL 函数都标记为 immutable，因此可以建立表达式索引；但 IDL 或线编码变化后，已有索引表达式的语义可能过期，届时应重建相关索引。`pg_thrift` 是解码和查询辅助工具，不能替代应用层 Thrift 兼容性测试。
