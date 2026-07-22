## 用法

来源：

- [官方 README](https://github.com/madebykite/pgunicoll/blob/af4e0c71fa3fe109ecca9acb2b254af1d545c2c8/README.md)
- [扩展 SQL](https://github.com/madebykite/pgunicoll/blob/af4e0c71fa3fe109ecca9acb2b254af1d545c2c8/sql/pgunicoll.sql)
- [ICU 后端实现](https://github.com/madebykite/pgunicoll/blob/af4e0c71fa3fe109ecca9acb2b254af1d545c2c8/pgunicoll.c)

`pgunicoll` 以 `bytea` 生成 ICU Unicode Collation Algorithm 排序键。应用需要对 PostgreSQL 文本表达式显式使用某个 ICU locale 排序时可以使用它，但它不会创建 PostgreSQL collation 对象，也不会自动影响普通 `ORDER BY` 表达式。

### 核心流程

创建扩展后，按 `pgunicoll(text, text)` 返回的排序键排列行。

```sql
CREATE EXTENSION pgunicoll;

SELECT name
FROM person
ORDER BY pgunicoll(name, 'de_DE');
```

locale 参数默认为空字符串。空字符串与 `root` 都会请求 ICU root collator。实现启用了 ICU 数字排序，因此连续数字按数值比较，而不仅是逐字节文本顺序。

对于重复搜索，可建立表达式索引；查询必须使用完全相同的 locale 表达式。

```sql
CREATE INDEX person_name_de_idx
ON person (pgunicoll(name, 'de_DE'));

SELECT name
FROM person
ORDER BY pgunicoll(name, 'de_DE');
```

### 重要对象

- `pgunicoll(text, text DEFAULT '') RETURNS bytea`：返回 ICU 排序键；声明为 STRICT 与 IMMUTABLE。

### 稳定性与兼容性

无论数据库编码为何，该函数始终把输入字节解释为 UTF-8。因此它不适用于非 UTF-8 数据库。locale 可用性与排序结果取决于链接的 ICU 库。

升级 ICU 可能改变排序键。升级后应重建所有表达式索引，并刷新任何已存储或物化的键；主库、备库与恢复目标应保持兼容 ICU 版本。函数虽然为索引使用声明成 IMMUTABLE，但只有 ICU 版本和 locale 数据固定时结果才稳定。应把原始文本而非生成键作为事实来源。
