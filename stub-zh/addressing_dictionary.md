## 用法

来源：

- [上游 README](https://github.com/pramsey/pgsql-addressing-dictionary/blob/bf3590e038d820c8eb5c79135e5a78abf5fc9f2f/README.md)
- [扩展 SQL](https://github.com/pramsey/pgsql-addressing-dictionary/blob/bf3590e038d820c8eb5c79135e5a78abf5fc9f2f/addressing_dictionary--1.1.sql)
- [扩展 control 文件](https://github.com/pramsey/pgsql-addressing-dictionary/blob/bf3590e038d820c8eb5c79135e5a78abf5fc9f2f/addressing_dictionary.control)

`addressing_dictionary` 为邮政地址安装全文检索词典与配置。文档中的 `addressing_en` 配置会规范化常见街道类型、方向和数字形式，使不同写法生成可匹配的词元。

先在 PostgreSQL 服务器上安装上游词典文件，再启用并使用扩展：

```sql
CREATE EXTENSION addressing_dictionary;

SELECT to_tsvector('addressing_en', '1234 n main st');
SELECT to_tsquery('addressing_en', 'north & main & street');
```

### 运维说明

版本 `1.1` 是可迁移的纯 SQL 扩展，但其文本搜索词典依赖上游构建过程安装的文件。安装 SQL 还会定义 `addressing_fr`；README 的示例和语言说明集中于 `addressing_en`，因此在依赖法语配置前应使用本地地址数据验证。地址规范化可能产生误匹配或多个候选结果，不能替代地理编码或地址校验。
