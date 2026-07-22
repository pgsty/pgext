## 用法

来源：

- [官方 README](https://github.com/postgrespro/pg_tsparser/blob/11ed232e83dfe7e900de084f09219041f880d1a2/README.md)
- [官方扩展 SQL](https://github.com/postgrespro/pg_tsparser/blob/11ed232e83dfe7e900de084f09219041f880d1a2/pg_tsparser--1.0.sql)
- [官方解析器实现](https://github.com/postgrespro/pg_tsparser/blob/11ed232e83dfe7e900de084f09219041f880d1a2/tsparser.c)

`pg_tsparser` 提供全文检索解析器 `tsparser`，它基于修改过的 PostgreSQL 9.6 默认解析器，会额外把下划线连接词和字母/数字混合的连字符词作为完整词元输出。它能让 `pg_trgm` 或 `123-abc` 等标识符在保留组成部分的同时作为整体检索。

### 创建配置

安装扩展只会创建解析器，不会提供完整的文本搜索配置。必须显式把词元类型映射到词典：

```sql
CREATE EXTENSION pg_tsparser;

CREATE TEXT SEARCH CONFIGURATION english_ts (
  PARSER = tsparser
);

ALTER TEXT SEARCH CONFIGURATION english_ts
  ADD MAPPING FOR email, file, float, host, hword_numpart, int,
  numhword, numword, sfloat, uint, url, url_path, version
  WITH simple;

ALTER TEXT SEARCH CONFIGURATION english_ts
  ADD MAPPING FOR asciiword, asciihword, hword_asciipart,
  word, hword, hword_part
  WITH english_stem;
```

采用配置前应比较分词结果：

```sql
SELECT to_tsvector('english', 'pg_trgm 123-abc') AS default_tokens,
       to_tsvector('english_ts', 'pg_trgm 123-abc') AS linked_tokens;
```

### 对象索引

- `tsparser` 是文本搜索解析器。
- `tsparser_start()`、`tsparser_nexttoken()`、`tsparser_end()`、`tsparser_lextype()` 和 `tsparser_headline()` 实现解析器回调。

### 注意事项

实现来自 PostgreSQL 9.6 解析器内部代码，因此每个目标大版本都必须验证二进制/源码兼容性。额外的完整词元会改变位置、排序、短语行为、索引大小和查询结果；修改配置后要重建相关 `tsvector` 列及索引。替换既有配置前，应使用代表性语料测试标点、版本字符串、路径、URL、邮件地址、非 ASCII 文本、词干、词典、标题和短语查询。
