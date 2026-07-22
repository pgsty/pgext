## 用法

来源：

- [官方类型与操作符文档](https://github.com/gdiazlo/pg_sexp/blob/17e9ef42c856f04981e30a094b275fa3c0f3edee/README.md)
- [0.1.0 版本 SQL 定义](https://github.com/gdiazlo/pg_sexp/blob/17e9ef42c856f04981e30a094b275fa3c0f3edee/sql/pg_sexp--0.1.0.sql)
- [扩展控制文件](https://github.com/gdiazlo/pg_sexp/blob/17e9ef42c856f04981e30a094b275fa3c0f3edee/pg_sexp.control)

`pg_sexp` 添加用于 Lisp 风格 S 表达式的紧凑二进制 `sexp` 类型。当 PostgreSQL 需要存储、检查、匹配和索引符号列表，而不希望把它们展平成文本时，可使用该扩展。

### 核心流程

```sql
CREATE EXTENSION pg_sexp;

CREATE TABLE document (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    expr sexp NOT NULL
);

INSERT INTO document (expr)
VALUES ('(user (name "alice") (age 30))');

SELECT id, car(expr), cdr(expr)
FROM document
WHERE expr @>> '(name "alice")';

CREATE INDEX document_expr_gin
ON document USING gin (expr sexp_gin_ops);
```

### 对象与操作符

- `car(sexp)`、`cdr(sexp)`、`nth(sexp, integer)`、`head(sexp)` 和 `sexp_length(sexp)` 用于检查列表。
- `sexp_typeof(sexp)` 以及 `is_list`、`is_atom`、`is_symbol`、`is_string`、`is_number` 和 `is_nil` 谓词用于检查值类型。
- `@>` 执行结构包含；`@>>` 执行按键且与顺序无关的包含；`@~` 与 `sexp_find` 支持通配符、剩余项和捕获模式。
- `sexp_gin_ops` 支持 GIN 包含搜索，等值还支持哈希索引。

### 兼容性与注意事项

上游要求 PostgreSQL 14 或更高版本。版本 `0.1.0` 仍处于 1.0 之前，因此在存储持久数据前应固定二进制格式和扩展版本。应测试解析器规范化、符号与字符串、整数与浮点数的区别、重复键、嵌套深度、通配符复杂度，以及索引与顺序扫描的结果一致性。深度嵌套或恶意表达式与广泛通配搜索可能开销很大；应在应用边界限制输入和查询。
