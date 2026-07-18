## 用法

来源：

- [上游 README](https://github.com/IgKh/pg_hspell/blob/89de5bffea99a918bc285096e70fcc4c5f102b34/README.md)
- [扩展 SQL](https://github.com/IgKh/pg_hspell/blob/89de5bffea99a918bc285096e70fcc4c5f102b34/pg_hspell--1.0.sql)
- [扩展 control 文件](https://github.com/IgKh/pg_hspell/blob/89de5bffea99a918bc285096e70fcc4c5f102b34/pg_hspell.control)
- [PostgreSQL 文本搜索配置文档](https://www.postgresql.org/docs/current/textsearch-configuration.html)

`pg_hspell` 提供名为 `hspell` 的文本搜索词典模板，用于对希伯来语单词做词形还原。它要求 GNU/Linux，以及使用 `--enable-linginfo` 构建的外部 `libhspell`。

扩展会创建词典，但不会创建完整的文本搜索配置，因此需要显式映射到配置中：

```sql
CREATE EXTENSION pg_hspell;

CREATE TEXT SEARCH CONFIGURATION hebrew (COPY = simple);
ALTER TEXT SEARCH CONFIGURATION hebrew
  ALTER MAPPING FOR word, hword, hword_part
  WITH hspell, simple;

SELECT to_tsvector('hebrew', 'הרכבת');
```

### 语言学限制

版本 `1.0` 记录支持 PostgreSQL 9.6 及以上版本。被识别的单词会输出所有可能词元，而不进行上下文消歧，因此提高召回率的同时会降低精确率。包含 Niqqud 的词元无法被 `libhspell` 识别，PostgreSQL 默认解析器还可能在引号处拆分希伯来语缩写。进入该词典前应去除 Niqqud，并使用代表性文本测试解析器映射。
