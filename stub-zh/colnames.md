## 用法

来源：

- [已复核 commit 的当前上游文档](https://github.com/theory/colnames/blob/00bdd0f4becdde261c0395ebfbc7943b49e5d519/doc/colnames.md)
- [1.7.0 版本 SQL 对象](https://github.com/theory/colnames/blob/00bdd0f4becdde261c0395ebfbc7943b49e5d519/sql/colnames.sql)
- [C 实现](https://github.com/theory/colnames/blob/00bdd0f4becdde261c0395ebfbc7943b49e5d519/src/colnames.c)

`colnames` 只提供一个函数 `colnames(record)`，以 `text[]` 返回输入记录的字段名。通用触发器或元编程代码需要复合值结构时可使用它。

```sql
CREATE EXTENSION colnames;
SELECT colnames(ROW(1, 'foo', 458.0));
SELECT colnames(t)
FROM (SELECT 1 AS id, 'Ada'::text AS name) AS t;
```

匿名行会生成 `f1` 等名称，具名子查询或表记录则保留属性名。如果输出用于生成 SQL，应测试已删除列、重复别名、domain 和不断变化的复合类型。

PostgreSQL 9.3 及以后可通过 `row_to_json` 与 `json_object_keys` 在不安装 C 扩展时取得名称。不方便部署扩展时应优先使用纯 SQL 方案。若把名称插入动态 SQL，必须使用 PostgreSQL 标识符引用函数并对值保持参数化；发现字段名并不能自动保证 SQL 构造安全。
