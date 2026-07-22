## 用法

来源：

- [官方 README](https://github.com/adjust/all_substrings_tokenizer/blob/f459a88981981659b7f21fab955ca1cbc802c219/README.md)
- [0.1 版本扩展 SQL](https://github.com/adjust/all_substrings_tokenizer/blob/f459a88981981659b7f21fab955ca1cbc802c219/sql/all_substrings_tokenizer--0.1.sql)
- [C 实现](https://github.com/adjust/all_substrings_tokenizer/blob/f459a88981981659b7f21fab955ca1cbc802c219/src/all_substrings_tokenizer.c)

`all_substrings_tokenizer` 会把文本展开为所有至少包含三个字符的连续子串。它是一个支持多字节字符的小型集合返回函数，可用于自定义搜索、匹配或演示流水线；它不是 PostgreSQL 全文搜索解析器，也不是可配置的分词器。

### 核心流程

```sql
CREATE EXTENSION all_substrings_tokenizer;

SELECT token
FROM all_substrings_set('1二湖楽a');
```

结果包含 `1二湖`、`1二湖楽`、`1二湖楽a`、`二湖楽`、`二湖楽a` 和 `湖楽a`。与其他数据联接时，使用复合结果的 `token` 字段：

```sql
SELECT d.id, s.token
FROM documents AS d
CROSS JOIN LATERAL all_substrings_set(d.title) AS s
WHERE char_length(s.token) = 3;
```

### 对象索引

- `all_substrings_set(text) -> setof __substrings_token` 输出各个子串。
- `__substrings_token` 是仅含一个 `token text` 属性的复合类型。

### 运维说明

最小长度被硬编码为三个字符，没有 SQL 参数或 GUC 可以修改。少于三个字符的输入不返回任何行。该函数为 `IMMUTABLE` 且 `STRICT`，因此 SQL `NULL` 也不返回任何行。

输出量随输入长度平方增长：包含 `n` 个字符的字符串会产生 `(n - 2) * (n - 1) / 2` 行。对用户可控文本调用前应限制输入长度。版本 `0.1` 可重定位且未声明预加载要求，但上游仓库已归档，应在目标 PostgreSQL 版本上进行兼容性测试。
