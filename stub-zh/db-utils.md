## 用法

来源：

- [packages/db-utils/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/db-utils/README.md)
- [packages/db-utils/sql/db-utils--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/db-utils/sql/db-utils--0.47.0.sql)
- [packages/db-utils/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/db-utils/Makefile)
- [packages/db-utils/db-utils.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/db-utils/db-utils.control)

`db-utils` 在固定模式中提供递归 JSONB 工具。`db_utils.jsonb_deep_merge` 递归合并对象，其他冲突由补丁值覆盖；`db_utils.jsonb_set_deep` 设置嵌套路径时可创建中间对象。

### 核心用法

```sql
CREATE EXTENSION "db-utils";
SELECT db_utils.jsonb_deep_merge('{"a":{"x":1}}', '{"a":{"y":2}}');
SELECT db_utils.jsonb_set_deep('{}', ARRAY['a','b'], '3'::jsonb);
```

### 运行边界

`db_utils.get_column_smart_comment` 以 `SECURITY DEFINER` 方式读取列注释。`db_utils.timestamps` 会修改指定表：补充时间戳列、替换时间戳触发器并创建缺少的索引。该辅助函数需要 DDL 权限，也会取得常规 DDL 锁，并非只读的时间戳查询函数。 依赖为 `inflection-db`, `pgpm-inflection`, `pgpm-stamps`, `pgpm-verify`, `plpgsql`。 这是纯 SQL 模块，无需预加载动态库或重启；该修订未公布 PostgreSQL 主版本支持矩阵。模式布局固定，不可迁移。启用前须部署相匹配的完整依赖。
