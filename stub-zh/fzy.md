## 用法

来源：

- [官方 README](https://github.com/romgrk/pg_fzy/blob/0c0aab00004c738e86b031d618c215d900f1f5f4/README.md)
- [0.0.1 扩展 SQL](https://github.com/romgrk/pg_fzy/blob/0c0aab00004c738e86b031d618c215d900f1f5f4/fzy--0.0.1.sql)

`fzy` 将 fzy 模糊匹配评分暴露给 SQL。它适合用缩写或尚未输入完整的查询词，对数量适中的候选字符串进行排序；它只是评分函数，不提供索引，也不是完整的搜索引擎。

### 核心流程

先在数据库中安装扩展，再为每个候选值计算分数并排序。第二个参数必须转换为 PostgreSQL 的 C 字符串伪类型。

```sql
CREATE EXTENSION fzy;

SELECT name,
       fzy('Stev', name::cstring) AS score
FROM users
ORDER BY score DESC;
```

分数越高表示匹配越好。0.0.1 版本只实现了 fzy 算法的大小写不敏感形式。

### SQL 对象

- `fzy(cstring, cstring) RETURNS integer`：计算匹配分数。

该函数声明为 strict，因此任一参数为空时结果为空。扩展 SQL 还将其声明为 volatile，这可能限制优化器可采用的优化。扩展没有提供索引操作符或访问方法，因此 PostgreSQL 会对查询选中的每一个候选行执行该函数。
