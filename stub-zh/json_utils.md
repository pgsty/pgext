## 用法

来源：

- [上游 README](https://github.com/asotolongo/pg_json_utils/blob/3910a970506f1f9978f8925322e0a75e391b8962/README.md)
- [扩展 SQL](https://github.com/asotolongo/pg_json_utils/blob/3910a970506f1f9978f8925322e0a75e391b8962/json_utils--0.1.sql)
- [扩展 control 文件](https://github.com/asotolongo/pg_json_utils/blob/3910a970506f1f9978f8925322e0a75e391b8962/json_utils.control)
- [PGXN 发行包](https://pgxn.org/dist/json_utils/)

`json_utils` 是针对 PostgreSQL 旧式 `json` 类型的 SQL/PLpgSQL 扩展，提供键存在、键值包含和相等比较运算符：

```sql
CREATE EXTENSION json_utils;

SELECT '{"field1":"value1","field2":341}'::json ? 'field2';
SELECT '{"field1":"value1","field2":341}'::json @> '"field2":341';
SELECT '{"field1":"value1"}'::json = '{"field1":"value1"}'::json;
```

### 版本与注意事项

PGXN 包标记为 `0.1.1`，但其 control 文件与安装脚本公开的扩展版本为 `0.1`。该实现采用文本匹配，而不是 PostgreSQL 的结构化 `jsonb` 语义；上游 README 还明确警告已发布版本存在已知缺陷。新应用应优先使用内置 `jsonb` 运算符；只有在使用代表性输入验证过其旧式行为后，才应采用 `json_utils`。
