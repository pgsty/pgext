## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/adjust/wltree/blob/f8fb773edfc793c6bd323beaa6cb8b1b98ce366d/README.md)
- [2.0 版本 SQL 对象](https://github.com/adjust/wltree/blob/f8fb773edfc793c6bd323beaa6cb8b1b98ce366d/wltree--2.0.sql)
- [扩展 control 文件](https://github.com/adjust/wltree/blob/f8fb773edfc793c6bd323beaa6cb8b1b98ce366d/wltree.control)

`wltree` 是 PostgreSQL `ltree` 模块的补丁副本。它使用 `::` 而不是 `.` 作为标签分隔符，并允许花括号、感叹号、星号和百分号等特殊字符，在查询中需转义。

```sql
CREATE EXTENSION wltree;
SELECT '!foo::{bar}::baz%'::ltree
       ~ '\!foo::\{bar\}::baz\%'::lquery;
```

该扩展故意创建熟悉的 `ltree`、`lquery` 及相关对象名，但使用不同于核心扩展的语法和语义。未经完整对象冲突审查，不要在同一数据库与核心 `ltree` 一起安装，也不要假定两者的值可以互换。

上游说明全文 `ltxtquery` 查询不支持转义字符。所有节点与客户端都应固定同一实现，测试解析、转义、操作符、GiST 索引、转储恢复、逻辑复制和升级，并在语义变化后重建依赖索引。外部提供的路径或查询必须独立验证；转义规则不能替代参数化 SQL。
