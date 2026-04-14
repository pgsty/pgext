## 用法
- GitHub 仓库: [`pierreforstmann/pg_query_rewrite`](https://github.com/pierreforstmann/pg_query_rewrite)
- README: [pierreforstmann/pg_query_rewrite/blob/master/README.md](https://github.com/pierreforstmann/pg_query_rewrite/blob/master/README.md)

`pg_query_rewrite` 是一个 PostgreSQL 扩展，可以把一条精确匹配的源 SQL 语句改写为另一条预定义 SQL 语句。它必须先通过 `shared_preload_libraries` 在服务器级加载，然后再在每个数据库中安装。

README 将它描述为一个由共享内存规则驱动的简易语句改写引擎。

### 安装

```conf
shared_preload_libraries = 'pg_query_rewrite'
pg_query_rewrite.max_rules = 10
```

```sql
CREATE EXTENSION pg_query_rewrite;
```

README 说明该扩展已在 PostgreSQL 9.5 到 18 之间通过测试。

### 管理规则

使用辅助函数管理改写规则：

```sql
SELECT pgqr_add_rule('select 10;', 'select 11;');
SELECT pgqr_remove_rule('select 10;');
SELECT pgqr_truncate();
SELECT pgqr_rules();
```

README 中的示例会把 `select 10;` 改写为 `select 11;`，然后展示插入后的规则列表。

### 行为

- 匹配是精确匹配，对大小写、空格和分号都敏感。
- 不支持参数化 SQL 语句。
- SQL 语句长度上限硬编码为 32K。
- 规则只存在于共享内存中；扩展本身不会持久化这些设置。
- `pg_query_rewrite.max_rules` 限制可改写的 SQL 语句数量，未设置时默认值为 `10`。

### 范围

上游 README 已足够说明这个扩展的用途、服务器级加载方式、规则管理、具体改写示例和主要限制，因此不需要额外的文档页或主页。
