## 用法

来源：

- [Official extension control file](https://github.com/choplin/pg-jsonpath/blob/968eff51fcb958b2641fd81cb5739b3c033a7979/jsonpath.control)
- [Official upstream README](https://github.com/choplin/pg-jsonpath/blob/968eff51fcb958b2641fd81cb5739b3c033a7979/README.md)

`jsonpath` 0.1 是 Stefan Goessner 早期 JavaScript JSONPath 实现的 PL/v8 封装。它早于 PostgreSQL 内置 SQL/JSON 路径功能，实现的是该 JavaScript 库的语法，而不是 SQL 标准 `jsonpath` 类型。

### 核心流程

先安装 `plv8`，再创建扩展并调用 `jsonPath(json, text)`。

```sql
CREATE EXTENSION plv8;
CREATE EXTENSION jsonpath;

SELECT jsonPath('{"x":{"a":1,"b":2}}'::json, '$.x.[a,b]');
```

函数返回包含全部匹配项的 PostgreSQL `json[]`。应使用内嵌 JavaScript 实现所接受的表达式；它们不能与围绕 PostgreSQL 原生 `jsonpath` 类型构建的操作符和函数互换。

### 注意事项

该项目规模小且年代较早，执行发生在 PL/v8 内部。除非必须兼容这一特定 API，否则在受支持的 PostgreSQL 版本上应优先使用原生 SQL/JSON 路径功能。应验证表达式语义、限制输入规模，并在不可信路径表达式或文档可能消耗过多后端 CPU 或内存时限制执行权限。
