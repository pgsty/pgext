## 用法

来源：

- [PGXN 官方文档](https://pgxn.org/dist/pg_icu_parser/1.0.0/)

`pg_icu_parser` 提供采用 ICU Unicode 单词边界分析的 `icu_parser` 全文搜索解析器。它适用于中文、日文、韩文和希伯来文等语言，因为 PostgreSQL 默认解析器以空白和模式为中心，可能无法产生合适的单词边界。

### 核心流程

服务器必须是 PostgreSQL 10 或更高版本，并且构建时已经启用 ICU。创建扩展后，建立使用该解析器的文本搜索配置，再映射所需的词元类型：

```sql
CREATE EXTENSION pg_icu_parser;

CREATE TEXT SEARCH CONFIGURATION sample (PARSER = icu_parser);
ALTER TEXT SEARCH CONFIGURATION sample
  ADD MAPPING FOR word WITH english_stem;
ALTER TEXT SEARCH CONFIGURATION sample
  ADD MAPPING FOR number, ideo, kana WITH simple;

SELECT to_tsvector('sample', '東京でPostgreSQLを使う');
```

应为实际语言选择词典；上面的 `english_stem` 只用于演示映射语法。

### 解析器接口

- `word`、`number`、`kana`、`ideo` 和 `blank` 是 `icu_parser` 输出的词元类型。
- `pg_icu_parser.locale` 选择 ICU locale，默认值为 `en`。上游指出单词边界规则通常不依赖 locale，因此只有在验证了具体需求后才应修改。
- 任意 PostgreSQL 文本搜索词典都可以映射到这些输出词元类型。

### 取舍

- 解析器要求 PostgreSQL 构建已经包含 ICU；创建扩展无法为未启用 ICU 的服务器补上该能力。
- 与默认解析器相比，它输出的词元类型更少、更粗，也不会把 URL、电子邮件等特殊模式识别为独立类别。
- 当前审阅的版本不支持 `ts_headline`。
- 上游报告其速度慢于默认解析器，而数据库编码为 UTF-8 时转换开销会较低。应在目标系统上用有代表性的文档和查询形态进行基准测试。
- 修改解析器映射或 locale 行为后，应重建已存储的 `tsvector` 值和依赖索引，避免新旧行采用不同分词规则。
