## 用法

来源：

- [官方 README](https://github.com/Houston-Inc/dict_voikko/blob/746eeef7f40cadbe5a527ccee4250a2f8daf2a94/README.md)
- [官方扩展 SQL](https://github.com/Houston-Inc/dict_voikko/blob/746eeef7f40cadbe5a527ccee4250a2f8daf2a94/dict_voikko--1.0.sql)
- [官方 C 实现](https://github.com/Houston-Inc/dict_voikko/blob/746eeef7f40cadbe5a527ccee4250a2f8daf2a94/dict_voikko.c)
- [官方 control 文件](https://github.com/Houston-Inc/dict_voikko/blob/746eeef7f40cadbe5a527ccee4250a2f8daf2a94/dict_voikko.control)

`dict_voikko` 增加 PostgreSQL 全文检索词典，使用 Voikko 形态分析把芬兰语单词和复合词还原为基础 lexeme。它只是词典组件，并非完整文本搜索配置。

### 核心流程

```sql
CREATE EXTENSION dict_voikko;

SELECT ts_lexize('voikko', 'kissoja');
SELECT ts_lexize('voikko', 'kerrostalollekohan');

CREATE TEXT SEARCH DICTIONARY finnish_voikko_stop (
  TEMPLATE = voikko_template,
  StopWords = finnish
);

CREATE TEXT SEARCH CONFIGURATION finnish_voikko (
  COPY = pg_catalog.finnish
);

ALTER TEXT SEARCH CONFIGURATION finnish_voikko
  ALTER MAPPING FOR asciiword, asciihword, hword_asciipart,
                    word, hword, hword_part
  WITH finnish_voikko_stop, finnish_stem;
```

默认 `voikko` 词典随 `voikko_template` 一起安装。`StopWords` 是模板唯一接受的选项。

### 词典链行为

识别到的屈折单词返回基本形式；文档中的复合词示例产生 `kerros` 与 `talo`。停用词返回空 lexeme 结果，无法识别的单词返回 SQL NULL，使映射中的下一个词典可以继续处理。应在 Voikko 词典后保留 `finnish_stem` 等后备词典。

实现只使用 Voikko 返回的第一项形态分析，并用固定正则表达式提取词干。应验证领域词汇、复合词、专名、编码和期望规范化，不能把它当成完整芬兰语语言分析器。

### 依赖与兼容性

服务端库必须链接 `libvoikko`，并安装芬兰语形态词典数据。C 源码打开语言 `fi-x-morpho`，旧 README 则描述名为 `mor-morpho` 的词典；应确认已安装 Voikko 数据实际暴露的语言标签，而不是盲目使用旧名称。

每个初始化文本搜索词典的后端都会加载 Voikko 资源。应计入其进程内存，并确保每个主库、备库或池化服务器都安装相同词典。

版本 `1.0` 使用旧代码库中的 PostgreSQL 与 Voikko API，最后修改于 2018 年，没有当前兼容矩阵。生产索引前应针对确切 PostgreSQL、`libvoikko` 和形态数据版本构建测试；分析器数据变化会改变生成的 `tsvector`，可能需要重建索引/数据。
