## 用法

来源：

- [PGXN 0.4.1](https://pgxn.org/dist/pg_grammar_guard/0.4.1/)

`pg_grammar_guard` 根据目录中的标识符生成 GBNF 或 JSON Schema，并检测已批准语法的漂移。它是无需预加载的纯 SQL 扩展，本包支持 PostgreSQL 14–18。control 文件明确依赖 `pg_living_assertions`。

### 生成语法

```sql
CREATE EXTENSION pg_grammar_guard CASCADE;
CREATE TABLE public.grammar_demo (id integer, label text);
SELECT grammar_guard.grammar_for_json(ARRAY[
  ROW('column', 'enum',
      grammar_guard.catalog_columns('public.grammar_demo'), true)
]::grammar_guard.grammar_field[]);
```

目录辅助函数枚举实际存在的表、列和枚举标签。生成器支持嵌套对象和有界数组；任意 SQL、文件路径等开放集合仍需单独验证。

### 检测变化

`grammar_guard.watch()` 保存用于重建语法的查询，`grammar_guard.check_grammar()` 依据实时目录执行检查。可通过 `living_assertions.status` 查看结果及其年龄。

只有受信任的管理员才应登记基线 SQL。语法能限制合法标识符，无法证明选择的表、关联或答案在语义上正确。旧 0.2 系列基线缺少原始生成查询，升级时需要人工重新批准。
