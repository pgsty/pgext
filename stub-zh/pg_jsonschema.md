


## 用法

> [pg_jsonschema: 提供 JSON Schema 校验的 PostgreSQL 扩展](https://github.com/supabase/pg_jsonschema)

`pg_jsonschema` 为 PostgreSQL 添加了 JSON Schema 校验函数，可通过检查约束对 JSON/JSONB 列实施模式验证。

### 函数

| 函数 | 描述 |
|------|------|
| `json_matches_schema(schema json, instance json)` | 验证 JSON 实例是否符合 schema，返回布尔值 |
| `jsonb_matches_schema(schema json, instance jsonb)` | 验证 JSONB 实例是否符合 schema，返回布尔值 |
| `jsonschema_is_valid(schema json)` | 检查 JSON schema 本身是否合法 |
| `jsonschema_validation_errors(schema json, instance json)` | 返回校验错误消息数组 |

### 表约束

使用检查约束强制文档结构：

```sql
CREATE TABLE customer (
    id serial PRIMARY KEY,
    metadata json,
    CHECK (
        json_matches_schema(
            '{
                "type": "object",
                "properties": {
                    "tags": {
                        "type": "array",
                        "items": {
                            "type": "string",
                            "maxLength": 16
                        }
                    }
                }
            }',
            metadata
        )
    )
);

-- 有效插入（通过检查约束）
INSERT INTO customer(metadata) VALUES ('{"tags": ["vip", "darkmode-ui"]}');

-- 无效插入（被检查约束拒绝）
INSERT INTO customer(metadata) VALUES ('{"tags": [1, 3]}');
-- ERROR: new row violates check constraint
```

### 错误检查

获取详细的校验错误信息：

```sql
SELECT jsonschema_validation_errors('{"maxLength": 4}', '"123456789"');
-- 返回: {"\"123456789\" is longer than 4 characters"}
```

### 模式校验

使用前验证 schema 是否格式正确：

```sql
SELECT jsonschema_is_valid('{
    "type": "object",
    "properties": {
        "name": {"type": "string"},
        "age":  {"type": "integer", "minimum": 0}
    },
    "required": ["name"]
}');
-- 返回: true
```
