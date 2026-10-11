## 用法

来源：

- [README.md](https://api.pgxn.org/src/pg_grammar_guard/pg_grammar_guard-0.4.8/README.md)
- [pg_grammar_guard.control](https://api.pgxn.org/src/pg_grammar_guard/pg_grammar_guard-0.4.8/pg_grammar_guard.control)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_grammar_guard/pg_grammar_guard-0.4.8/CHANGELOG.md)

`pg_grammar_guard` 0.4.8 根据目录标识符生成 GBNF 或 JSON Schema，并监测已批准语法的漂移。它以纯 SQL 实现，要求 `pg_living_assertions` 0.5.9 或更新版本。

### 核心用法

```sql
CREATE EXTENSION pg_grammar_guard CASCADE;
CREATE TABLE public.grammar_demo (id integer, label text);
SELECT grammar_guard.grammar_for_json(ARRAY[
  ROW('column', 'enum',
      grammar_guard.catalog_columns('public.grammar_demo'), true)
]::grammar_guard.grammar_field[]);
```

### 运行边界

`grammar_guard.catalog_tables()` 与 `grammar_guard.catalog_columns()` 提供目录名称，`grammar_guard.grammar_for_json()` 组合带类型的字段、嵌套结构与有界数组。`grammar_guard.watch()` 保存重建语法的 SQL，`grammar_guard.check_grammar()` 与认可状态比较，结果见 `living_assertions.status`。生成查询使用作者的搜索路径并以作者身份执行。NULL 输入、空或重复字段名、空 pivot 及无效边界会被拒绝，`max_items` 上限为 1000。0.4.8 改变了规范化指纹，有序目录列表也可能使已有监测首次检查失效，应核实后重新批准。语法只约束合法词元，不保证语义正确或访问授权；仅登记可信 SQL。无需预加载。Apache-2.0 从 0.4.4 开始适用，更早版本保留原许可证。
