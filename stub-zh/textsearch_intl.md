## 用法

来源：

- [README.md](https://github.com/benjaminsanborn/textsearch_intl/blob/3cc6bf9a9c22a3946802e95ff472678bb1c2c60d/README.md)
- [textsearch_intl--1.0.sql](https://github.com/benjaminsanborn/textsearch_intl/blob/3cc6bf9a9c22a3946802e95ff472678bb1c2c60d/textsearch_intl--1.0.sql)
- [textsearch_intl.control](https://github.com/benjaminsanborn/textsearch_intl/blob/3cc6bf9a9c22a3946802e95ff472678bb1c2c60d/textsearch_intl.control)
- [META.json](https://github.com/benjaminsanborn/textsearch_intl/blob/3cc6bf9a9c22a3946802e95ff472678bb1c2c60d/META.json)
- [LICENSE](https://github.com/benjaminsanborn/textsearch_intl/blob/3cc6bf9a9c22a3946802e95ff472678bb1c2c60d/LICENSE)
- [改名提交](https://github.com/benjaminsanborn/textsearch_intl/commit/b023b27aa16b82fc3c1b80ef0928e9f185e7f43e)
- [PGXN icu_parser 1.1.0](https://pgxn.org/dist/icu_parser/1.1.0/)

`textsearch_intl` 将 ICU 分词、规范化和 n-gram 匹配组合为多语言全文检索功能。它要求 PostgreSQL 14 或更新版本、UTF8 数据库、ICU 和 `unaccent`；上游 CI 覆盖 PostgreSQL 14–18。

### 名称与版本

项目于 2026-10-09 将 `icu_parser` 改名为 `textsearch_intl`，并重新开始版本编号：当前源码的 SQL/control 版本为 `1.0`，META 版本为 `1.0.0`。GitHub/PGXN 最新已发布版本仍是旧名 `icu_parser` 的发行包 `1.1.0`，其 SQL 版本为 `1.1`。新名称尚无发布标签，也没有从旧名称升级的 SQL 脚本。

### 核心用法

```sql
CREATE EXTENSION textsearch_intl CASCADE;
SELECT alias, token FROM ts_debug('intl_simple', '東京駅前ラーメン店 ร้านกาแฟ Café');
SELECT to_tsvector('intl_search', 'Café') @@ phraseto_tsquery('intl_search', 'cafe');
SELECT to_tsvector('intl_search_ngram', '행복한편의점')
       @@ phraseto_tsquery('intl_search_ngram', '편의점');
SELECT intl_icu_version();
```

### 检索配置

- `intl_simple` 使用 `intl` 解析器，保留重音和字符宽度。
- `intl_search` 组合 `intl_normalize`、`unaccent` 与 `intl_join`，进行规范化的整词检索。词语拼接默认关闭。
- `intl_search_ngram` 使用 `intl_runs` 与 `intl_ngram`，匹配韩文、中日文词元内部，以及不带空格的泰文、老挝文、高棉文和缅甸文片段。N-gram 可能增加误匹配和索引体积。

### 运行边界

使用超级用户安装；扩展可迁移，未标记为可信，也不需要预加载。词序重要时优先使用 `phraseto_tsquery()`。`intl_icu_version()` 返回已加载的 ICU 版本；ICU 主版本改变后应重建相关索引。ICU 分词不是日文形态分析，部分 CJK 词前缀查询仍不可靠。
