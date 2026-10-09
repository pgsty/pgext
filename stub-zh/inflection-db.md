## 用法

来源：

- [packages/inflection-db/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/inflection-db/README.md)
- [packages/inflection-db/sql/inflection-db--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/inflection-db/sql/inflection-db--0.47.0.sql)
- [packages/inflection-db/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/inflection-db/Makefile)
- [packages/inflection-db/inflection-db.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/inflection-db/inflection-db.control)

`inflection-db` 在 PGPM 名称变换库之上提供数据库命名辅助函数。`inflection_db.get_table_name` 将名称复数化，其他函数生成字段、索引、约束、模式和命名空间的名称。

### 核心用法

```sql
CREATE EXTENSION "inflection-db";
SELECT inflection_db.get_table_name('BlogPost');
SELECT inflection_db.get_index_name('BlogPost', ARRAY['created_at']);
```

### 运行边界

固定版本 SQL 中的 `inflection_db.get_identifier` 截取前 63 个字符，不保证名称唯一；多字节标识符仍须检查 PostgreSQL 的字节长度限制。结果只是文本，不会执行 DDL，拼接 SQL 时必须正确引用标识符。`inflection_db.get_index_name` 与 `inflection_db.get_unique_index_name` 接受字段数组。 依赖为 `pgpm-inflection`, `pgpm-verify`, `plpgsql`。 这是纯 SQL 模块，无需预加载动态库或重启；该修订未公布 PostgreSQL 主版本支持矩阵。模式布局固定，不可迁移。启用前须部署相匹配的完整依赖。
