## 用法

来源：

- [官方 README](https://github.com/DblBee/json_validator_ext/blob/81c2935930072a1c03f0626a4239c2cb180396d9/README.md)
- [Rust 实现](https://github.com/DblBee/json_validator_ext/blob/81c2935930072a1c03f0626a4239c2cb180396d9/src/lib.rs)
- [扩展控制文件](https://github.com/DblBee/json_validator_ext/blob/81c2935930072a1c03f0626a4239c2cb180396d9/json_validator_ext.control)
- [Rust 包清单](https://github.com/DblBee/json_validator_ext/blob/81c2935930072a1c03f0626a4239c2cb180396d9/Cargo.toml)

`json_validator_ext` 根据 JSON Schema 校验 PostgreSQL `jsonb` 值。它提供适用于约束的简洁操作符、等价函数、详细错误报告，以及每个后端独立的已编译 schema LRU 缓存。

### 核心流程

安装扩展后，在 `@@` 左侧放 JSON 实例，右侧放对应 schema：

```sql
CREATE EXTENSION json_validator_ext;

CREATE TABLE support_ticket (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    details jsonb NOT NULL,
    CONSTRAINT details_valid CHECK (
        details @@ '{
          "required": ["email"],
          "properties": {"email": {"type": "string", "format": "email"}}
        }'::jsonb
    )
);

SELECT *
FROM unnest(explain_json_schema_errors(
    '{"type":"integer"}'::jsonb,
    '"not-an-integer"'::jsonb
));
```

### 对象与配置

- `jsonb @@ jsonb` 返回左侧实例是否满足右侧 schema。
- `validate_json_schema(jsonb, jsonb)` 先接收 schema，再接收实例。
- `explain_json_schema_errors(jsonb, jsonb)` 返回包含实例路径和校验消息的文本数组。
- `json_validator.cache_size` 控制每个后端缓存的已编译 schema 数量。默认值为 100，取值范围为 1 至 1000，并且是可由超级用户在运行时设置的参数。

无效 schema 定义会让布尔 API 返回 false；应使用诊断函数区分 schema 无效与实例无效。控制文件将扩展标记为仅超级用户可安装、不可信且不可迁移。规范扩展和包名是 `json_validator_ext`，但 README 标题把项目称为 `pg_json_schema`；版本 `0.0.0` 仍属早期 API，用于完整性约束前应固定已复核构建，并回归测试 JSON Schema 行为。
