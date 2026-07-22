## 用法

来源：

- [官方 WIP 说明](https://github.com/3manuek/twitter_parser/blob/1fafaae28bd7dc2c0d01617ddf992a19027f2c43/README)
- [1.0 版本扩展 SQL](https://github.com/3manuek/twitter_parser/blob/1fafaae28bd7dc2c0d01617ddf992a19027f2c43/twitter_parser--1.0.sql)
- [C 解析器实现](https://github.com/3manuek/twitter_parser/blob/1fafaae28bd7dc2c0d01617ddf992a19027f2c43/twitter_parser.c)

`twitter_parser` 是一个未完成的全文搜索解析器原型，用于识别 Twitter 风格的 `@mention` 和 `#hashtag` 词元。创建扩展会注册名为 `twitterparser` 的解析器；它不会创建可直接使用的文本搜索配置。

### 核心流程

在考虑创建配置前，先使用解析器自省函数评估它：

```sql
CREATE EXTENSION twitter_parser;

SELECT * FROM ts_token_type('twitterparser');
SELECT * FROM ts_parse('twitterparser', '#postgres');
SELECT * FROM ts_parse('twitterparser', '@postgres');
```

解析器声明词元 ID `24` 为 `mention`，词元 ID `25` 为 `hashtag`。文本搜索配置还需要为这些别名设置词典映射，但这个原型尚不完整，不适合生产索引。

### 对象索引

- 文本搜索解析器 `twitterparser` 提供启动、取词元、结束、摘要和词元类型回调。
- 词元类型 `mention`（`24`）用于以 `@` 开头的文本。
- 词元类型 `hashtag`（`25`）用于以 `#` 开头的文本。
- `twitterprs_start`、`twitterprs_getlexeme`、`twitterprs_end` 和 `twitterprs_lextype` 是底层解析器支持函数，不是应用 API。

### 运维说明

版本 `1.0` 可重定位，且未声明依赖或预加载要求。上游将其标记为 WIP。实现只会把标签或提及读取到 ASCII 空格为止，不校验标点或 Unicode 规则。更重要的是，它不会跳过空格或普通文本，并且在这类输入上没有设置词元类型，因此混合句子可能提前停止或出现不可预测行为。

在修复并测试状态机之前，不要用这个解析器构建持久全文索引。应在隔离数据库中测试非法输入、多词元、标点、多字节文本、摘要生成、升级和重建索引。生产环境应使用有维护的解析器，或把标签与提及规范化到单独的索引列。
