## 用法

来源：

- [上游 README](https://github.com/junhkang/kor_search/blob/f10c6ff5f474b6cb455de56f6fb5b5af0dbf6c96/README.md)
- [扩展 SQL](https://github.com/junhkang/kor_search/blob/f10c6ff5f474b6cb455de56f6fb5b5af0dbf6c96/sql/kor_search--1.0.0.sql)
- [C 实现](https://github.com/junhkang/kor_search/blob/f10c6ff5f474b6cb455de56f6fb5b5af0dbf6c96/src/kor_search.c)
- [PGXN 文档](https://pgxn.org/dist/kor_search/1.0.0/doc/kor_search.html)

`kor_search` 使用数据库内的关键字和同义词表进行韩语与英语单词匹配，不调用外部翻译器或形态分析器。应先安装 `pg_trgm`，因为版本 `1.0.0` 的 SQL 会创建 `gin_trgm_ops` 索引，但 control 文件没有声明这一依赖：

```sql
CREATE EXTENSION pg_trgm;
CREATE EXTENSION kor_search;

SELECT kor_search_like('이것은 엘지 제품입니다', 'lg');
SELECT kor_search_similar('I eat rice', '밥 먹다');
SELECT kor_search_tsvector('데이터 과학은 중요합니다', 'data science');
SELECT kor_search_regex('자바와 파이썬', 'java|python');
```

### API 与安全说明

README 示例使用了过期名称 `kor_like` 和 `kor_regex_search`；已发布的安装 SQL 实际定义的是上例采用的 `kor_search_like` 与 `kor_search_regex`。搜索质量和开销取决于 `kor_search_word_transform` 与 `kor_search_word_synonyms` 的内容和规模。C 实现把调用者提供的文本未经 SQL 转义便插入内部 SPI 语句，因此不能让不可信输入调用这些函数；此类使用前应撤销 `EXECUTE` 权限或修补实现。
