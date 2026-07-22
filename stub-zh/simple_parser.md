## 用法

来源：

- [官方 README](https://github.com/fairingrey/pg_simple_parser/blob/cdeaa0f1ed3b89d1961312c01e529f1c7c068084/README.md)
- [扩展控制文件](https://github.com/fairingrey/pg_simple_parser/blob/cdeaa0f1ed3b89d1961312c01e529f1c7c068084/simple_parser.control)
- [1.0 版扩展 SQL](https://github.com/fairingrey/pg_simple_parser/blob/cdeaa0f1ed3b89d1961312c01e529f1c7c068084/simple_parser--1.0.sql)
- [C 分词实现](https://github.com/fairingrey/pg_simple_parser/blob/cdeaa0f1ed3b89d1961312c01e529f1c7c068084/simple_parser.c)

`simple_parser` 1.0 安装全文检索解析器 `simpleparser`。它把 ASCII 空格字符之间的每段字节都视为一个 `word`，因此下划线、连字符等标点会保留在词元内部。

### 创建文本检索配置

扩展只安装解析器，不提供现成的文本检索配置。应创建配置，并把 `word` 词元类型映射到字典：

```sql
CREATE EXTENSION simple_parser;

CREATE TEXT SEARCH CONFIGURATION simple_space (
  PARSER = simpleparser
);

ALTER TEXT SEARCH CONFIGURATION simple_space
ADD MAPPING FOR word WITH simple;

SELECT to_tsvector('simple_space', 'alpha_beta gamma-delta');
```

解析器还会发出 `blank` 词元类型，使 PostgreSQL 标准 headline 函数能够保留间距，但它通常不需要字典映射。

### 已安装对象

- `simpleparser` 是公开的文本检索解析器。
- `simpleprs_start`、`simpleprs_getlexeme`、`simpleprs_end` 和 `simpleprs_lextype` 是内部 C 解析器回调；应调用 PostgreSQL 文本检索函数，而不是直接调用它们。
- 解析器复用 PostgreSQL 词元 ID 3（`word`）和 12（`blank`），并把 headline 生成委托给 `pg_catalog.prsd_headline`。

### 分词注意事项

只有字节值为 ASCII 空格（`0x20`）时才会分隔。制表符、换行符、不换行空格和其他 Unicode 空白都保留在 `word` 内，标点和控制字节也是如此。解析器不会执行语言学分词、Unicode 规范化、大小写折叠、词干提取或停用词删除；这些行为只能由映射的字典提供。

索引应用数据前应检查实际词元：

```sql
SELECT t.alias, p.token
FROM ts_parse('simpleparser', E'alpha_beta\tgamma delta') AS p
JOIN ts_token_type('simpleparser') AS t USING (tokid);
```

构建 `tsvector` 索引后再修改配置或字典，需要重新生成存储的向量并重建受影响索引。只有当字面 ASCII 空格边界符合应用搜索语义时，才应使用此解析器。
