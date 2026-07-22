## 用法

来源：

- [官方 README](https://github.com/zulip/tsearch_extras/blob/f566e9606bac00a22c4b62e9511b022c861db05c/README.md)
- [扩展 SQL](https://github.com/zulip/tsearch_extras/blob/f566e9606bac00a22c4b62e9511b022c861db05c/tsearch_extras--1.0.sql)
- [全文搜索实现](https://github.com/zulip/tsearch_extras/blob/f566e9606bac00a22c4b62e9511b022c861db05c/tsearch_extras.c)

`tsearch_extras` 版本 `1.0` 暴露 PostgreSQL 全文搜索的底层信息：查询在源文本中的匹配位置与长度，以及 `tsvector` 中存储的词位。

### 核心流程

```sql
CREATE EXTENSION tsearch_extras;

SELECT ts_match_locs_array(
  'The quick brown fox jumped. The jumping continued.',
  plainto_tsquery('jump')
);

SELECT tsvector_lexemes(
  to_tsvector('The quick brown fox jumped over the lazy dog')
);
```

当结果不能依赖会话默认值时，应使用显式文本搜索配置：

```sql
SELECT ts_match_locs_array(
  'english'::regconfig,
  'The quick brown fox jumped.',
  plainto_tsquery('english', 'jump')
);
```

### 对象

- `ts_match_locs_array(text, tsquery) returns int4[][2]`：使用当前文本搜索配置返回匹配位置/长度对。
- `ts_match_locs_array(regconfig, text, tsquery) returns int4[][2]`：使用指定配置返回相同结构。
- `tsvector_lexemes(tsvector) returns text[]`：返回向量中存储的词位，不包含位置信息或权重。

### 兼容性边界

匹配位置通过 PostgreSQL 的解析器与 headline 内部机制产生，因此应使用应用实际采用的配置、编码、词典与服务器版本校验偏移量。仓库已归档，C 源码依赖内部全文搜索头文件。应针对目标 PostgreSQL 版本确认编译与回归行为；不需要预加载。
