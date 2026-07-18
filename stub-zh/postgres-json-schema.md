## 用法

来源：

- [上游 README](https://github.com/gavinwahl/postgres-json-schema/blob/e25c758abb6395a5457bd1124da2983a48dc2aac/README.md)
- [扩展 control 文件](https://github.com/gavinwahl/postgres-json-schema/blob/e25c758abb6395a5457bd1124da2983a48dc2aac/postgres-json-schema.control)

`postgres-json-schema` `0.1.1` 提供纯 PL/pgSQL 函数，用于按 JSON Schema draft 4 校验 PostgreSQL `jsonb` 值。典型用途是通过 check constraint 强制约束文档结构。

扩展名包含连字符，必须加引号：

```sql
CREATE EXTENSION "postgres-json-schema";

CREATE TABLE events (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    payload jsonb NOT NULL,
    CONSTRAINT payload_is_object
        CHECK (validate_json_schema('{"type":"object"}', payload))
);

INSERT INTO events (payload) VALUES ('{"kind":"created"}');
```

校验器覆盖 JSON Schema draft 4，但不支持远程 HTTP reference。应把引用的 schema 保存在本地，不要依赖网络获取。该扩展不安装动态库，也无需配置预加载。
