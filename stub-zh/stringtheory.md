## 用法

来源：

- [官方上游文档](https://github.com/hydradatabase/pg_stringtheory/blob/249f4cd83bdb6557ffb0143ed4f3407f71598224/README.md)
- [官方扩展 SQL](https://github.com/hydradatabase/pg_stringtheory/blob/249f4cd83bdb6557ffb0143ed4f3407f71598224/stringtheory--1.0.2.sql)
- [官方实现](https://github.com/hydradatabase/pg_stringtheory/blob/249f4cd83bdb6557ffb0143ed4f3407f71598224/src/text.cpp)

`stringtheory` 提供面向 SIMD 的 C++ 精确文本相等和子串搜索函数。它在 x86_64 上使用 SSE4.2，在 aarch64/arm64 上使用 SWAR64。函数位于 `stringtheory` 模式中，适合显式求值比较，而不是可直接替换的操作符或索引访问方法。

### 核心流程

创建扩展并调用带模式限定的函数：

```sql
CREATE EXTENSION stringtheory;

SELECT stringtheory.equals('postgres', 'postgres');
SELECT stringtheory.strstr('postgresql', 'sql');
```

第一个查询返回 true。第二个查询返回 `8`，因为 `stringtheory.strstr(text, text)` 报告从零开始的字节偏移；没有匹配时返回 `-1`。

### 重要对象

- `stringtheory.equals(text, text)` 返回精确字节相等的布尔结果。
- `stringtheory.strstr(text, text)` 返回首个从零开始的字节偏移，或 `-1`。

两个函数都声明为 `STRICT`、`IMMUTABLE` 和 parallel safe。因此参数为 NULL 时会直接返回 NULL，而不会进入 C++ 函数。

### 语义与性能注意事项

已复核的实现会在 `stringtheory.equals(text, text)` 任一输入为空时返回 false，包括 `equals('', '')`。`stringtheory.strstr(text, text)` 在待搜索文本或搜索词为空时返回 `-1`。这些行为不同于普通相等和常见子串语义。偏移量以字节而不是 Unicode 字符计数，因此多字节文本需要特别处理。

该扩展既没有定义操作符，也没有定义索引操作符类。调用这些函数的谓词会作为函数表达式求值，不能自动使用普通 text B-tree 或 trigram 索引。上游基准数字只来自一台 AMD 主机，不能作为普遍性能保证。替换内置运算前，应以代表性数据验证 CPU 指令支持、准确的 PostgreSQL ABI、空字符串与 Unicode 行为、选择性和端到端查询计划。
