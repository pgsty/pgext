## 用法

来源：

- [官方 kham-pg README](https://github.com/preedep/kham/blob/ec2394cc70ab9687b6e0bcd7c02c2f75a4cf9fbd/kham-pg/README.md)
- [版本 0.8.2 扩展 SQL](https://github.com/preedep/kham/blob/ec2394cc70ab9687b6e0bcd7c02c2f75a4cf9fbd/kham-pg/sql/kham_pg--0.8.2.sql)
- [PGXN 0.8.2 文档](https://pgxn.org/dist/kham_pg/0.8.2/README.html)

`kham_pg` 为 PostgreSQL 提供泰语全文搜索。其 `kham` 解析器无需依赖空格即可分词；字典则在标准 `tsvector` 与 `tsquery` 流程中加入规范化词、语音编码、RTGS 罗马化、数字规范化、词性词位及命名实体处理。

### 核心流程

```sql
CREATE EXTENSION kham_pg;

SELECT *
FROM ts_parse('kham', 'ทักษิณเดินทางไปกรุงเทพ');

SELECT kham_tsvector('กินข้าวกับปลา');
SELECT kham_tsquery('ปลา');

CREATE TABLE articles (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title text,
    body text NOT NULL
);

CREATE INDEX articles_kham_fts_idx
ON articles USING gin (kham_tsvector(body));

SELECT id, title
FROM articles
WHERE kham_tsvector(body) @@ kham_tsquery('ปลา');
```

`kham_tsvector(document)` 是 `to_tsvector('kham', document)` 的便捷封装，`kham_tsquery(query)` 则封装 `plainto_tsquery('kham', query)`。标准高亮同样可用：

```sql
SELECT ts_headline('kham', body, kham_tsquery('ปลา'))
FROM articles;
```

### 文本搜索对象

- 解析器与即用配置：`kham`。
- 默认字典：`kham_fts_dict`，使用 `lk82` 泰语 Soundex。
- 替代字典：`kham_fts_dict_udom83` 与 `kham_fts_dict_metasound`。
- 简单小写透传字典：`kham_dict`。
- Token 类型包括 `thai`、`latin`、`number`、`punct`、`emoji`、`unknown` 与 `named`。

泰语与命名实体 Token 可展开为多个同位置词位，从而让语音或拉丁文字查询匹配泰语文本。泰语停用词会被抑制，泰语数字会与 ASCII 形式一起规范化，标点与 emoji 则由默认配置丢弃。

### 注意事项

版本 `0.8.2` 不可重定位；已复核上游版本面向 PostgreSQL `14` 到 `18`。分词、语音展开、停用词移除与命名实体识别都会影响召回率和索引大小，应使用具体业务领域的泰语文本验证。GIN 表达式索引只有在查询使用相同表达式与配置时才能匹配。库安装完成后无需预加载或重启。
