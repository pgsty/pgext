## 用法

来源：

- [Official alohadb_jsonschema.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_jsonschema/alohadb_jsonschema.control)
- [Official alohadb_jsonschema--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_jsonschema/alohadb_jsonschema--1.0.sql)
- [Official jsonschema_validator.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_jsonschema/jsonschema_validator.c)

`alohadb_jsonschema` 1.0 按 JSON Schema Draft-07 的一个子集校验 JSONB 文档，并提供可选的命名模式注册表。源码属于 AlohaDB 内核，尚未确认与原生 PostgreSQL 兼容。

### 校验文档

```sql
CREATE EXTENSION alohadb_jsonschema;
SELECT jsonschema_is_valid('{"age":30}'::jsonb,
  '{"type":"object","required":["age"],"properties":{"age":{"type":"integer","minimum":0}}}'::jsonb);
SELECT * FROM jsonschema_validate('{"age":-1}'::jsonb,
  '{"properties":{"age":{"minimum":0}}}'::jsonb);
SELECT jsonschema_register('person', '{"type":"object","required":["age"]}'::jsonb);
SELECT jsonschema_validate_named('{"age":30}'::jsonb, 'person');
```

### 对象与约束

`jsonschema_is_valid()` 返回布尔值，标记为不可变、严格且并行安全，适合使用固定模式的 CHECK。`jsonschema_validate()` 返回有效性、错误路径与消息。`alohadb_jsonschema_registry` 保存名称、JSONB 模式和创建时间。命名校验需要读取注册表，因此标记为稳定而非不可变。通过数据库权限控制注册表写入；修改已注册模式不会自动重新校验应用中的已有数据。

### 支持边界

支持的校验包括类型、属性、必需字段、数字/字符串/数组/对象边界、模式匹配、枚举、常量与逻辑组合。`$ref` 仅解析内部 JSON 指针，不应假定支持远程引用或全部 Draft-07 关键字。安装需要管理员权限；控制文件标记扩展不可重定位且不是可信扩展。此模块没有共享预加载要求。
