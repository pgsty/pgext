## 用法

来源：

- [官方仓库 README](https://github.com/xaneets/pg-debyte/blob/e0b6e4d4c98781acc0cfbc458fc95de9e0fbb1af/README.md)
- [官方扩展控制文件](https://github.com/xaneets/pg-debyte/blob/e0b6e4d4c98781acc0cfbc458fc95de9e0fbb1af/pg_debyte_ext/pg_debyte_ext.control)
- [官方示例扩展实现](https://github.com/xaneets/pg-debyte/blob/e0b6e4d4c98781acc0cfbc458fc95de9e0fbb1af/pg_debyte_ext/src/lib.rs)

`pg_debyte_ext` 0.1.0 是 `pg-debyte` Rust workspace 中的 PostgreSQL 示例扩展。它把已注册的 Bincode `bytea` 载荷解码为 `jsonb`，既可以显式指定类型标识与 schema 版本，也可以使用自描述 envelope。上游希望开发者构建自己的 registry-backed 扩展，而不是把这个演示 schema 当成通用解码器。

### 核心流程

固定示例以 UUID `11111111-1111-1111-1111-111111111111` 注册一个演示类型，并提供第二个已知 schema 解码器：

```sql
CREATE EXTENSION pg_debyte_ext;

SELECT bytea_to_json_by_id(
  decode('010464656d6f', 'hex'),
  '11111111-1111-1111-1111-111111111111'::uuid,
  1::smallint
);

SELECT bytea_to_json_demo_record_second(
  decode('01067365636f6e6401', 'hex')
);
```

第一个调用返回包含整数 ID 和文本 label 的 JSON 对象。第二个调用解码示例中固定的 `DemoRecordSecond` schema。`bytea_to_json_auto(bytea)` 只接受仓库定义的 envelope 格式，并根据其中编码的类型、schema 版本、codec 与 action 标识进行分派。

### 重要对象

- `bytea_to_json_by_id(bytea, uuid, smallint)` 使用显式选择的已注册解码器处理原始字节。
- `bytea_to_json_auto(bytea)` 解析 envelope 并自动分派。
- `bytea_to_json_demo_record_second(bytea)` 是为第二个演示记录生成的已知 schema 函数。
- `pg_debyte.max_input_bytes`、`pg_debyte.max_output_bytes` 和 `pg_debyte.max_json_bytes` 是用户可设置的安全限制；固定源码中三者默认都是 950 MiB。

### 扩展构建边界

固定提交中的可复用 workspace crate 版本为 0.2.1，而示例扩展控制文件仍为 0.1.0。自定义扩展需要定义可序列化 Rust 类型、codec、UUID/schema-version key、可选 action（例如有界 Zstandard 解压）以及静态 registry。原始载荷必须与所选 Bincode schema 精确匹配；envelope 必须引用已注册的 codec、action 与 decoder ID。

### 运维说明

示例为 PostgreSQL 15 和 17 提供 pgrx 构建 feature，必须使用与目标大版本匹配的构建。它没有预加载或重启要求：库初始化会注册 GUC 和演示 registry。解码错误、未知类型、错误 envelope、输入输出超限以及 Rust panic 都会转换为 PostgreSQL 错误，但仍应把限制下调到适合实际负载的值。未经资源测试，不要用 950 MiB 默认值处理不可信的大型载荷。
