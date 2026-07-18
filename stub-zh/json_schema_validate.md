## 用法

来源：

- [上游 README](https://github.com/adunstan/json_schema_validate/blob/f771374c64cd014fb5ba9373832e2109e6fd4e35/README.md)
- [扩展 SQL](https://github.com/adunstan/json_schema_validate/blob/f771374c64cd014fb5ba9373832e2109e6fd4e35/json_schema_validate--0.1.0.sql)
- [扩展 control 文件](https://github.com/adunstan/json_schema_validate/blob/f771374c64cd014fb5ba9373832e2109e6fd4e35/json_schema_validate.control)

`json_schema_validate` 根据 JSON Schema 校验 `json` 或 `jsonb` 数据。`jsonschema_is_valid` 返回布尔值；`jsonschema_validate` 对合法数据返回空数组，对非法数据返回结构化错误对象：

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
    jsonschema_is_valid(data, '{"type":"object","required":["id"]}'::jsonschema_compiled)
  )
);
```

### 支持范围

版本 `0.1.0` 要求 PostgreSQL 14 或更高版本。它支持许多 Draft 7 及后续版本关键字、本地 JSON Pointer 引用和预编译模式，但未实现元组形式的 `prefixItems`、依赖关键字、未求值条目或属性关键字、外部引用以及动态引用标识符。未知 format 会被忽略。将该扩展用作安全或数据完整性边界前，应核对上游关键字列表。
