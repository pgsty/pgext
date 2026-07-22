## 用法

来源：

- [官方扩展 SQL](https://github.com/choplin/pg_msgpack/blob/7cf18817053f4a42df80c34d816003505de8b1c1/pg_msgpack--0.0.1.sql)
- [官方转换实现](https://github.com/choplin/pg_msgpack/blob/7cf18817053f4a42df80c34d816003505de8b1c1/pg_msgpack.c)
- [官方访问操作符实现](https://github.com/choplin/pg_msgpack/blob/7cf18817053f4a42df80c34d816003505de8b1c1/pg_msgpack_op.c)

`pg_msgpack` 增加了处理传统 MessagePack 数据的 `msgpack` 值类型，并提供类 JSON 文本输入输出、类型转换和字段/元素访问。它可用于检查 PostgreSQL 中已有的紧凑载荷，但这个 2013 年实现早于现代 MessagePack 与 PostgreSQL `jsonb` 设施。

### 核心流程

```sql
CREATE EXTENSION pg_msgpack;

SELECT '{"name":"Ada","scores":[10,20]}'::msgpack -> 'name';
SELECT ('{"name":"Ada","scores":[10,20]}'::msgpack -> 'scores') -> 1;
SELECT '{"a":1}'::json::msgpack::bytea;
```

数组位置从零开始。键不存在、容器类型不兼容或解包失败时，通常返回 SQL NULL 而不是报错。

### SQL 接口

- `msgpack` 类型接受并输出类 JSON 文本。
- `msgpack` 与 `json` 之间的转换使用 `WITH INOUT`。
- `msgpack` 与 `bytea` 之间的赋值转换使用 `WITHOUT FUNCTION`，不经校验或转换便暴露其存储二进制表示。
- 对象 `->` 操作符接受文本键，数组 `->` 操作符接受整数位置；两者都返回另一个 `msgpack` 值。

该扩展没有提供 `jsonb` 转换、变更 API 或 `GIN`/B-tree 操作符类。如果索引、当前 JSON 语义或丰富操作符比 MessagePack 线格式兼容性更重要，应使用原生 JSON 类型。

### 安全与兼容性

绝不能给整数访问操作符传入负数组位置。核验的 C 代码只检查上界，随后把带符号位置直接用于指针偏移；负值可能访问非法内存并破坏后端稳定性。查询进入扩展前必须校验所有索引。

传统格式化器和 raw-string 处理无法保证所有值都能完成现代 MessagePack 到 JSON 的完整往返。应把转换和文本输出视为检查工具，而不是规范交换编码。

本次核验的上游仓库只有一个 2013 年代的修订，并使用旧 PostgreSQL JSON 解析器与 messagepack-c 接口。部署前应针对确切服务端和库版本完成构建与模糊测试，并让生产后端隔离不可信载荷。
