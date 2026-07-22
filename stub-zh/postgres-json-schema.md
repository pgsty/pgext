## 用法

来源：

- [官方 README](https://github.com/gavinwahl/postgres-json-schema/blob/e25c758abb6395a5457bd1124da2983a48dc2aac/README.md)
- [官方扩展 SQL](https://github.com/gavinwahl/postgres-json-schema/blob/e25c758abb6395a5457bd1124da2983a48dc2aac/postgres-json-schema--0.1.1.sql)
- [官方控制文件](https://github.com/gavinwahl/postgres-json-schema/blob/e25c758abb6395a5457bd1124da2983a48dc2aac/postgres-json-schema.control)

`postgres-json-schema` 提供一个 PL/pgSQL 验证器，用于按照 JSON Schema draft v4 检查 JSONB 值。版本 `0.1.1` 适合绑定该历史草案的数据库约束，但不实现后续 JSON Schema 草案或远程 HTTP 引用。

### 核心流程

安装扩展，并在检查约束中使用 `validate_json_schema`：

```sql
CREATE EXTENSION "postgres-json-schema";

CREATE TABLE events (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  payload jsonb NOT NULL,
  CONSTRAINT events_payload_schema CHECK (
    validate_json_schema(
      '{"type":"object","required":["kind"],"properties":{"kind":{"type":"string"}}}'::jsonb,
      payload
    )
  )
);
```

有效记录返回 true；无效记录会让检查约束拒绝写入。

### 函数与支持特性

`validate_json_schema(schema jsonb, data jsonb, root_schema jsonb DEFAULT NULL)` 返回布尔值，并递归计算 draft-v4 关键字。源码覆盖类型、属性、必需键、数组项、数值与长度边界、组合关键字、枚举、唯一性、依赖、模式、格式和根锚定的本地 `$ref` 路径。第三个参数在递归时用于保留根模式，常规调用应省略。

### 运维注意事项

README 明确不支持远程 HTTP 引用。应把模式视为可信配置：畸形或意外的模式结构可能抛出 SQL 错误，而不是简单返回 false。约束覆盖的每次插入或更新都会在 PL/pgSQL 中同步执行验证，因此应为大文档和高写入速率做基准测试。修改字面模式需要重新定义约束，并验证已有记录。应用契约应明确固定 draft v4 语义；需要较新草案、远程解析或详细错误路径的应用，应使用数据库外仍在维护的验证器或更新的扩展。
