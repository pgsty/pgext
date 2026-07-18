## 用法

来源：

- [PostgreSQL 9.6 tsearch2 control 文件](https://github.com/postgres/postgres/blob/70cf253d18da103864037011bead1683a96502dc/contrib/tsearch2/tsearch2.control)
- [PostgreSQL 9.6 兼容 SQL](https://github.com/postgres/postgres/blob/70cf253d18da103864037011bead1683a96502dc/contrib/tsearch2/tsearch2--1.0.sql)
- [PostgreSQL 9.6 兼容实现](https://github.com/postgres/postgres/blob/70cf253d18da103864037011bead1683a96502dc/contrib/tsearch2/tsearch2.c)
- [Amazon RDS PostgreSQL 扩展历史](https://docs.aws.amazon.com/AmazonRDS/latest/PostgreSQLReleaseNotes/postgresql-extensions.html)

`tsearch2` 版本 `1.0` 是面向 PostgreSQL 8.3 之前全文检索 API 的向后兼容软件包。它会创建 `tsvector` 与 `tsquery` 兼容 domain、旧 parser 与 dictionary 入口、遗留 rank/headline 包装器，以及历史 `tsearch2()` 更新 trigger。其用途仅限于协助迁移旧应用。

### 遗留兼容

在仍提供该软件包的旧服务器或 Amazon RDS 引擎上：

```sql
CREATE EXTENSION tsearch2;

SELECT to_tsvector('english', 'fat cats ate rats');
SELECT to_tsquery('english', 'cat & rat');
```

这些文本参数重载会把旧调用适配到 PostgreSQL 内置文本检索实现。新代码应直接使用当前基于 `regconfig` 的全文检索函数。

### 注意事项

- PostgreSQL 把该扩展描述为 pre-8.3 API 的兼容层，而不是独立搜索引擎。它不可重定位，并会引入许多旧名称，可能遮蔽现代内置函数或造成混淆。
- Amazon RDS 把 `tsearch2` 标记为 PostgreSQL 10 上的 deprecated 扩展，并表示从 RDS PostgreSQL 11.1 开始移除。当前 RDS 版本必须使用内置全文检索。
- 不要再依赖兼容 domain、可变的当前 parser/dictionary 状态或 `tsearch2()` trigger。应把 schema 与 SQL 迁移到现代文本检索配置及 `tsvector_update_trigger` 替代方案。
- 版本 `1.0` 代表冻结的兼容软件包。托管服务是否可用取决于确切的历史引擎版本，而不是本目录项。
