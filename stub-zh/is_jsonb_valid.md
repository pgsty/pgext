## 用法

来源：

- [上游 README](https://github.com/furstenheim/is_jsonb_valid/blob/3afa28d70500791a233dc17c919d6136b83cb2ef/README.md)
- [扩展 SQL](https://github.com/furstenheim/is_jsonb_valid/blob/3afa28d70500791a233dc17c919d6136b83cb2ef/is_jsonb_valid--0.1.4.sql)
- [扩展 control 文件](https://github.com/furstenheim/is_jsonb_valid/blob/3afa28d70500791a233dc17c919d6136b83cb2ef/is_jsonb_valid.control)
- [PGXN 发行包](https://pgxn.org/dist/is_jsonb_valid/)

`is_jsonb_valid` 根据 JSON Schema 校验 JSONB 值并返回布尔结果。使用 `is_jsonb_valid(schema, data)` 校验 draft 4，或使用 `is_jsonb_valid_draft_v7(schema, data)` 校验 draft 7：

```sql
CREATE EXTENSION is_jsonb_valid;

SELECT is_jsonb_valid(
  '{"type":"object","required":["name"]}'::jsonb,
  '{"name":"Ada"}'::jsonb
);

SELECT is_jsonb_valid_draft_v7(
  '{"type":"integer","minimum":0}'::jsonb,
  '7'::jsonb
);
```

### 模式支持范围

上游版本 `0.1.4` 记录支持 PostgreSQL 9.6 及以上版本。它不支持远程引用，本地引用仅限模式根节点中的 definitions，并且未实现 JSON Schema 的 format 校验。对于不支持的关键字与引用布局，应将其视为应用层校验缺口，不能假定扩展完整兼容相应 draft。
