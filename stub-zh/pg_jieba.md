## 用法

来源：

- [官方v2.0.1 README](https://github.com/jaiminpan/pg_jieba/blob/v2.0.1/README.md)
- [扩展控制文件](https://github.com/jaiminpan/pg_jieba/blob/v2.0.1/pg_jieba.control)
- [SQL解析器和配置定义](https://github.com/jaiminpan/pg_jieba/blob/v2.0.1/pg_jieba.sql)

`pg_jieba` 将基于 Jieba 的中文分词功能添加到 PostgreSQL 全文搜索中。上游的 `v2.0.1` 源代码发布安装了 SQL 扩展版本 `1.1.0`，其控制文件记录了这一点。它提供了独立的文档和查询解析器以及现成可用的文字搜索配置。

### 核心工作流程

```sql
CREATE EXTENSION pg_jieba;

SELECT to_tsvector(
    'jiebacfg',
    '小明硕士毕业于中国科学院计算所，后在日本京都大学深造'
);

SELECT plainto_tsquery('jiebaqry', '云计算专家');
```

使用 `jiebacfg` 构建可搜索的文档向量，并使用 `jiebaqry` 分段用户查询：

```sql
ALTER TABLE articles
ADD COLUMN search_vector tsvector
GENERATED ALWAYS AS (to_tsvector('jiebacfg', body)) STORED;

CREATE INDEX articles_search_idx
ON articles USING GIN (search_vector);

SELECT title
FROM articles
WHERE search_vector @@ plainto_tsquery('jiebaqry', '中文全文检索');
```

### 对象索引

- `jieba`: 文档文本搜索解析器。
- `jiebaqry`: 查询导向的文字搜索解析器。
- `jiebacfg`: 使用 `jieba` 和 `jieba_stem` 的文档文字搜索配置。
- `jiebaqry`: 同名的查询解析器文字搜索配置。
- `jieba_stem`: 包含 Jieba 停用词的简单字典，用于解析器的标记类别。

### 自定义词典和注意事项

上游从 PostgreSQL 的 `jieba.user.dict.utf8` 目录读取一个名为 `tsearch_data` 的自定义词典。条目可能包含一个单词及其可选的词性标签：

```text
云计算
韩玉鉴赏
蓝翔 nz
```

- v2.x 源代码需要 C++11 兼容编译器，因为其捆绑了 `cppjieba` 依赖。
- 上游发布的兼容性测试已过时且有限。针对生产中使用的具体 PostgreSQL 主版本构建和回归测试该包。
- 更改词典会改变分词结果。当字典输出变化时，请重新计算存储的 `tsvector` 值并重建相关索引。
