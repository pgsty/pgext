## 用法

来源：

- [官方 README](https://github.com/pramsey/pgsql-addressing-dictionary/blob/master/README.md)
- [扩展控制文件](https://github.com/pramsey/pgsql-addressing-dictionary/blob/master/addressing_dictionary.control)
- [1.1 版扩展 SQL](https://github.com/pramsey/pgsql-addressing-dictionary/blob/master/addressing_dictionary--1.1.sql)
- [PostgreSQL 文本搜索词典文档](https://www.postgresql.org/docs/current/textsearch-dictionaries.html)

addressing_dictionary 安装面向地址的全文搜索配置与词典。它适合在本地地址近似检索中，把街道类型、方向和序数的不同写法归一为可比较的词位；它不是地理编码器，也不能替代地址解析与校验。

### 核心流程

在数据库中安装一次扩展，用所提供的配置归一化地址文本，并让索引表达式与查询表达式保持一致。

```sql
CREATE EXTENSION addressing_dictionary;

SELECT to_tsvector('addressing_en', '1234 n main st');

CREATE INDEX address_search_idx
ON address_book
USING gin (to_tsvector('addressing_en', full_address));

SELECT *
FROM address_book
WHERE to_tsvector('addressing_en', full_address)
      @@ plainto_tsquery('addressing_en', '1234 north main street');
```

1.1 版会保留方向缩写，同时产生展开形式，因此单字母的北向标记可以输出两种形式。这种设计刻意保留歧义，而不会假定每次出现都只有一种含义。

### 已安装对象

- 英文配置把地址标记的同义词词典、停用词词典和叙词表组合起来。
- 法文配置安装对应的法文文本搜索对象。
- 两套配置都复制自 PostgreSQL 的简单配置，因此通过扩展的数据文件归一化，而不是使用通用语言词干提取器。

### 运维边界

该扩展可重定位，不需要预加载或重启。所有可能执行查询或恢复扩展的服务器都必须安装其词典、同义词、停用词和叙词表文件。

全文匹配属于启发式方法。缩写可能有歧义，全国范围的数据也可能产生大量误匹配或重复候选。应把它用于改进候选检索，尤其是范围有限的城市或地区；对正确性有要求时，仍需另外使用结构化地址或地理空间检查。
