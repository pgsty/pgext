## 用法

来源：

- [官方 README](https://github.com/hooopo/pg-fuzzywuzzy/blob/b02f84ac42f7cc1c5bc51f4daf0c58a68752bf56/README.md)
- [0.0.2 版本安装 SQL](https://github.com/hooopo/pg-fuzzywuzzy/blob/b02f84ac42f7cc1c5bc51f4daf0c58a68752bf56/sql/fuzzywuzzy.sql)
- [扩展控制文件](https://github.com/hooopo/pg-fuzzywuzzy/blob/b02f84ac42f7cc1c5bc51f4daf0c58a68752bf56/fuzzywuzzy.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/fuzzywuzzy/)

`fuzzywuzzy` 为短文本提供一个 SQL 层的相似度分数。它封装 `fuzzystrmatch` 的 Levenshtein 距离，适用于不需要完整搜索系统时的简单模糊排序或过滤。

### 核心流程

安装扩展后，调用 `ratio(text, text)` 可得到 0 至 100 的整数分数：

```sql
CREATE EXTENSION fuzzywuzzy;

SELECT name, ratio(name, '人棉锦纶') AS score
FROM products
ORDER BY score DESC
LIMIT 3;
```

如果 `fuzzystrmatch` 尚未存在，安装脚本会自动创建它。如果目标环境限制创建扩展，应显式预置该依赖。

### 对象与语义

- `ratio(text, text)` 对相同的非空字符串返回 100；任一参数为 NULL 或空字符串时返回 0。
- 实现根据 `levenshtein(text, text)` 、两个输入长度与整数返回值计算分数；它受 fuzzywuzzy 启发，但没有暴露 Python fuzzywuzzy 的更广 API。
- 该函数声明为 `IMMUTABLE`，版本 `0.0.2` 可迁移。

项目规模很小，且没有说明 Unicode 规范化、排序规则、长度上限或索引契约。应明确地规范化文本，测试多字节和长输入，不要把该分数当成经过校准的概率。
