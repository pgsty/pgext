## 用法

来源：

- [README.md](https://github.com/benjaminsanborn/textsearch_intl/blob/d24405defbbfa249af420436b6b2442d8342246c/README.md)
- [icu_parser--1.1.sql](https://github.com/benjaminsanborn/textsearch_intl/blob/d24405defbbfa249af420436b6b2442d8342246c/icu_parser--1.1.sql)
- [LICENSE](https://github.com/benjaminsanborn/textsearch_intl/blob/d24405defbbfa249af420436b6b2442d8342246c/LICENSE)
- [icu_parser.control](https://github.com/benjaminsanborn/textsearch_intl/blob/d24405defbbfa249af420436b6b2442d8342246c/icu_parser.control)

`icu_parser` 使用 ICU 单词边界切分中文、日文和泰文等不以空格分词的文字。发行版本为 1.1.0，SQL/control 版本为 1.1。要求 PostgreSQL 14 或更新版本、UTF8 数据库、ICU 和 `unaccent`。

### 核心用法

```sql
CREATE EXTENSION icu_parser CASCADE;
SELECT alias, token FROM ts_debug('icu_simple', '東京駅前ラーメン店 ร้านกาแฟ Café');
SELECT to_tsvector('icu_search', 'Café') @@ phraseto_tsquery('icu_search', 'cafe');
SELECT icu_parser_icu_version();
```

### 运行边界

`icu` 解析器产生单词、数字、假名、表意文字和空白词元。`icu_simple` 转为小写但保留重音和字符宽度；`icu_search` 依次使用 `icu_normalize`、`unaccent` 与 `simple`，执行 NFKC 和撇号移除，同时保留相关泰文、老挝文元音。词序重要时优先使用 `phraseto_tsquery()`。`icu_parser_icu_version()` 返回运行时 ICU 版本。ICU 主版本改变后应重建相关索引，因为分词结果可能变化。词典分词并非日文形态分析或任意子串搜索，部分 CJK 词前缀可能无法匹配。原生安装未标记为可信，也不需要预加载。
