## 用法

来源：

- [已复核 commit 的官方 README](https://github.com/adunstan/json_schema_validate/blob/f771374c64cd014fb5ba9373832e2109e6fd4e35/README.md)
- [0.1.0 版本 SQL API](https://github.com/adunstan/json_schema_validate/blob/f771374c64cd014fb5ba9373832e2109e6fd4e35/json_schema_validate--0.1.0.sql)
- [验证实现](https://github.com/adunstan/json_schema_validate/blob/f771374c64cd014fb5ba9373832e2109e6fd4e35/json_schema_validate.c)
- [扩展 control 文件](https://github.com/adunstan/json_schema_validate/blob/f771374c64cd014fb5ba9373832e2109e6fd4e35/json_schema_validate.control)

`json_schema_validate` 使用较完整但并非全覆盖的 JSON Schema 实现，验证 PostgreSQL `json` 与 `jsonb` 值。需要布尔结果时使用 `jsonschema_is_valid`，需要结构化错误时使用 `jsonschema_validate`，重复使用同一模式时则使用 `jsonschema_compile`。0.1.0 版本要求 PostgreSQL 14 或更高版本。

### 核心流程

```sql
CREATE EXTENSION json_schema_validate;

SELECT jsonschema_is_valid(
  '{"id":1,"name":"Ada"}'::jsonb,
  '{"type":"object","required":["id","name"]}'::jsonb
);

SELECT jsonschema_validate(
  '{"id":"wrong"}'::jsonb,
  '{"properties":{"id":{"type":"integer"}}}'::jsonb
);

CREATE TABLE validated_events (
  data jsonb CHECK (
    jsonschema_is_valid(
      data,
      jsonschema_compile('{"type":"object","required":["id"]}'::jsonb)
    )
  )
);
```

对于有效输入，已复核的 `jsonschema_validate` 实现返回空数组 `[]`；输入无效时返回含路径和消息的对象数组。部分源码注释错误地声称有效输入返回 SQL NULL，因此应以实际实现为准，并对打包版本运行回归测试。

### 主要对象

`jsonschema_is_valid` 与 `jsonschema_validate` 都有用于 `json`、`jsonb` 及已编译 `jsonschema_compiled` 类型的重载。`jsonschema_compile(jsonb)` 预先解析模式，并缓存正则表达式以供重复验证。函数声明为 immutable、strict 且 parallel-safe，因此常量模式适合用于 CHECK 约束。

### 模式范围与注意事项

实现覆盖常用的类型、对象、数组、字符串、数字、组合、条件、格式与本地引用关键字。它不支持元组式 `prefixItems`、依赖关键字、未求值属性或数组项关键字、外部引用及动态引用。只支持本地 JSON Pointer 引用，未知格式会被忽略。

被忽略或不支持的关键字可能让完整标准实现本应拒绝的文档通过。把扩展用作完整性或安全边界前，必须固定并测试确切的模式词汇；对于深层嵌套模式、大数组、复杂正则表达式或大量累计验证错误，也应设置资源边界。
