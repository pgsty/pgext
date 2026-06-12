来源：[jsonschema v0.1.9 README](https://github.com/theory/pg-jsonschema-boon/blob/v0.1.9/README.md)、[documentation](https://github.com/theory/pg-jsonschema-boon/blob/v0.1.9/doc/jsonschema.md)、[control file](https://github.com/theory/pg-jsonschema-boon/blob/v0.1.9/jsonschema.control)、[SQL definition](https://github.com/theory/pg-jsonschema-boon/blob/v0.1.9/sql/jsonschema--0.1.9.sql)。

## 用法

`jsonschema` 在 PostgreSQL 内根据 JSON Schema 校验 JSON 和 JSONB 值。它是 `theory/pg-jsonschema-boon` 扩展，不同于 Supabase `pg_jsonschema`，但提供了名为 `json_matches_schema()` 和 `jsonb_matches_schema()` 的兼容包装函数。

该扩展通过 Rust `boon` 校验器支持 JSON Schema draft 4、draft 6、draft 7、draft 2019-09 和 draft 2020-12。运行时除了 PostgreSQL 之外没有其他依赖。

### 校验 Schema 与文档

```sql
CREATE EXTENSION IF NOT EXISTS jsonschema;

SELECT jsonschema_is_valid(
  '{
     "type": "object",
     "required": ["name", "email"],
     "properties": {
       "name":  { "type": "string" },
       "age":   { "type": "number", "minimum": 0 },
       "email": { "type": "string", "format": "email" }
     }
   }'::json
);

SELECT jsonschema_validates(
  '{"name":"Amos Burton","email":"amos@rocinante.ship"}'::json,
  '{
     "type": "object",
     "required": ["name", "email"],
     "properties": {
       "name":  { "type": "string" },
       "email": { "type": "string", "format": "email" }
     }
   }'::json
);
```

`jsonschema_is_valid(schema)` 返回 schema 本身是否能够编译，并能按所选 draft 通过校验。`jsonschema_validates(data, schema)` 返回 JSON/JSONB 值是否满足该 schema。

### CHECK 约束

```sql
CREATE TABLE customer_profile (
  id       bigserial PRIMARY KEY,
  profile  jsonb NOT NULL,
  CHECK (
    jsonschema_validates(
      profile,
      '{
         "type": "object",
         "required": ["email"],
         "properties": {
           "email": { "type": "string", "format": "email" },
           "tags":  {
             "type": "array",
             "items": { "type": "string", "maxLength": 16 }
           }
         }
       }'::jsonb
    )
  )
);
```

当数据库需要在写入时拒绝格式错误的 JSON 文档时，可使用约束。

### 组合式 Schema

```sql
SELECT jsonschema_validates(
  jsonb_build_object(
    'first_name', 'Naomi',
    'last_name', 'Nagata',
    'shipping_address', jsonb_build_object(
      'street_address', '1 Rocinante Way',
      'city', 'Ceres Station',
      'state', 'The Belt'
    )
  ),
  'https://example.com/schemas/customer',
  '{
     "$id": "https://example.com/schemas/address",
     "type": "object",
     "required": ["street_address", "city", "state"],
     "properties": {
       "street_address": { "type": "string" },
       "city": { "type": "string" },
       "state": { "type": "string" }
     }
   }'::jsonb,
  '{
     "$id": "https://example.com/schemas/customer",
     "type": "object",
     "required": ["first_name", "last_name", "shipping_address"],
     "properties": {
       "first_name": { "type": "string" },
       "last_name": { "type": "string" },
       "shipping_address": { "$ref": "/schemas/address" }
     }
   }'::jsonb
);
```

带 `id` 的重载允许多个 schema 通过 `$id` 相互引用，适合组件化的 JSON Schema 定义。

### 兼容函数

```sql
SELECT json_matches_schema(
  '{"type":"string","maxLength":4}'::json,
  '"1234"'::json
);

SELECT jsonb_matches_schema(
  '{"type":"object","required":["id"]}'::json,
  '{"id":42}'::jsonb
);
```

这些包装函数沿用常见的 `pg_jsonschema` 参数顺序：schema 在前，instance 在后。

### Draft 版本选择与注意事项

```sql
SET jsonschema.default_draft = 'V2020';
SET jsonschema.default_draft = 'V7';
```

如果 schema 省略 `$schema`，`jsonschema.default_draft` 会控制默认 draft。支持值为 `V4`、`V6`、`V7`、`V2019` 和 `V2020`。

- 如果任一参数为 NULL，`jsonschema_validates(data, schema)` 返回 NULL。
- 无效或无法编译的 schema 可能在校验调用中抛出错误；文档校验失败会返回 `false`，并在 `INFO` 级别记录详情。
- `jsonschema_is_valid(id, VARIADIC schemas)` 和 `jsonschema_validates(data, id, VARIADIC schemas)` 需要匹配的 `$id` 值，才能可靠解析组合式 schema。
