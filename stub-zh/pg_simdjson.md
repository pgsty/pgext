## 用法

来源：

- [上游 README](https://github.com/erthalion/pg_simdjson/blob/4e5a20eab6296d48e41da6f255c70c16bbb3d413/README.md)
- [版本 0.1 SQL 定义](https://github.com/erthalion/pg_simdjson/blob/4e5a20eab6296d48e41da6f255c70c16bbb3d413/pg_simdjson--0.1.sql)
- [解析器实现](https://github.com/erthalion/pg_simdjson/blob/4e5a20eab6296d48e41da6f255c70c16bbb3d413/pg_simdjson.cpp)

`pg_simdjson` 版本 `0.1` 是一个 SIMD JSON 解析原型。它接受文本形式的 JSON 并返回 PostgreSQL `jsonb`；只有在目标 CPU 和 PostgreSQL 版本上对实际打包构建完成基准测试与验证后才应使用。

### 核心流程

```sql
CREATE EXTENSION pg_simdjson;

SELECT simdjson_parse('{"id":42,"tags":["fast","json"]}');
SELECT simdjson_parse(payload_text) FROM incoming_events;
```

扩展只提供一个函数：`simdjson_parse(text) returns jsonb`。它被声明为不可变、严格且并行安全。有效的 JSON 对象、数组和标量会转换为普通 `jsonb`；无效输入会报错。

### 语义与注意事项

解析后遵循 PostgreSQL `jsonb` 语义：无意义空白和对象键顺序不会保留，值会按 `jsonb` 规范化。上游说明回归结果除错误消息外与 PostgreSQL JSONB 测试一致；应用不能依赖完全相同的诊断文本。

该仓库内嵌了较老的 simdjson 源码快照，并把项目称为原型。历史基准结果不是性能保证。它不要求预加载，但在非实验环境使用前，仍需验证 C++ ABI、编译器、CPU 指令支持、畸形输入行为与现代 PostgreSQL 兼容性。
