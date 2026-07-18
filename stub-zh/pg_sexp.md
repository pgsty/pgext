## 用法

来源：

- [当前上游类型与操作符文档](https://github.com/gdiazlo/pg_sexp/blob/17e9ef42c856f04981e30a094b275fa3c0f3edee/README.md)
- [0.1.0 版本 SQL 定义](https://github.com/gdiazlo/pg_sexp/blob/17e9ef42c856f04981e30a094b275fa3c0f3edee/sql/pg_sexp--0.1.0.sql)
- [扩展控制文件](https://github.com/gdiazlo/pg_sexp/blob/17e9ef42c856f04981e30a094b275fa3c0f3edee/pg_sexp.control)

`pg_sexp` 添加用于 Lisp 风格 S 表达式的紧凑二进制 `sexp` 类型。它提供列表访问、类型检查、结构与键包含、通配符匹配、等值哈希和 GIN 操作符类。

```sql
CREATE EXTENSION pg_sexp;
SELECT car('(a b c)'::sexp), cdr('(a b c)'::sexp);
SELECT '(user (name "alice") (age 30))'::sexp @>> '(name "alice")';
CREATE INDEX expr_gin ON documents USING gin (expr sexp_gin_ops);
```

上游声明支持 PostgreSQL 14 及以上版本。依赖索引前，应测试解析器规范化、数值与字符串区别、重复键、嵌套深度、通配符复杂度以及操作符与索引结果一致性。深度嵌套或恶意表达式可能开销很大；由于格式仍处于 1.0 之前，应固定扩展版本。
