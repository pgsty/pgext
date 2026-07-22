## 用法

来源：

- [官方上游 README](https://github.com/zdk/pg-search-thai/blob/e943164a3640a65c5714028447189f9a3a7b4f82/README.md)
- [官方扩展 SQL](https://github.com/zdk/pg-search-thai/blob/e943164a3640a65c5714028447189f9a3a7b4f82/thai_parser/sql/thai_parser--1.0.0.sql)
- [官方扩展控制文件](https://github.com/zdk/pg-search-thai/blob/e943164a3640a65c5714028447189f9a3a7b4f82/thai_parser/thai_parser.control)

`thai_parser` 为 PostgreSQL 全文检索提供泰语分词能力，适用于空格不能可靠标记词边界的泰语文本。它创建解析器 `thai_parser` 与开箱即用的配置 `thaicfg`；该扩展依赖 `libthai`，且数据库必须使用 UTF-8 编码。

### 核心流程

创建扩展，生成持久化检索向量，并用 GIN 建立索引：

```sql
CREATE EXTENSION thai_parser;

CREATE TABLE articles (
    id bigserial PRIMARY KEY,
    body text NOT NULL,
    search tsvector GENERATED ALWAYS AS
        (to_tsvector('thaicfg', body)) STORED
);

CREATE INDEX articles_search_idx ON articles USING GIN (search);

INSERT INTO articles (body) VALUES
    ('ส้มตำกับข้าวเหนียว'),
    ('ต้มยำกุ้ง');

SELECT id, body
FROM articles, to_tsquery('thaicfg', 'ส้มตำ') AS query
WHERE search @@ query;
```

应用输入应使用绑定参数及合适的查询构造器。`to_tsquery` 要求完整的全文检索语法；如果不希望用户提供运算符，请使用 `plainto_tsquery`。

### 重要对象

- `thai_parser` 既是扩展名，也是安装后的文本检索解析器。
- `thaicfg` 是扩展创建的默认文本检索配置。
- `ts_parse` 可显示分词器输出以便诊断。词元编号 97、98、99 分别表示泰语词、ASCII 词与空格。
- 默认 `thaicfg` 映射只索引泰语词元。混合泰语和英语的内容需要被检索时，应在独立的文本检索配置中加入并测试 `asciiword` 映射。

### 前置条件与限制

服务器构建需要解析器所使用的 `libthai` 开发和运行库，数据库编码必须为 UTF-8。这些属于主机级依赖，因此仅有数据库转储不足以在新服务器上完成恢复。

上游仓库还包含可选的 Hunspell 词典资源，但只使用解析器时并不需要。泰语分词与词典选择会改变索引词元；扩展、词典或映射变更后，应重建受影响的全文索引，并针对有代表性的拼写、复合词、混合语言文本、排序及事务内检索行为进行回归测试。
