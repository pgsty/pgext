## 用法

来源：

- [官方 README](https://github.com/dgillis/pg_vartype/blob/536d1581db6c9e6651f014797bf6391a5231e9be/README.md)
- [官方扩展 SQL](https://github.com/dgillis/pg_vartype/blob/536d1581db6c9e6651f014797bf6391a5231e9be/sql/pg_vartype--0.4.sql)
- [官方比较实现](https://github.com/dgillis/pg_vartype/blob/536d1581db6c9e6651f014797bf6391a5231e9be/src/pg_vartype.c)

`pg_vartype` 提供可在同一列中保存整数、浮点数、布尔值、字符串、日期或时间戳的 `vartype` 类型。它早于 PostgreSQL `jsonb`；上游也指出，对于许多当前应用，原生 JSON 存储更合适。

### 核心流程

```sql
CREATE EXTENSION pg_vartype;

CREATE TABLE attributes (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  value vartype
);

INSERT INTO attributes (value) VALUES
  (42::bigint::vartype),
  ('blue'::text::vartype),
  (DATE '2026-01-01'::vartype),
  (false::boolean::vartype);

SELECT id, value, vartype_type(value), vartype_len(value)
FROM attributes;
```

使用无类型字面量时，字符串本身必须带双引号，例如 `'"blue"'::vartype`；否则 `'false'::vartype` 这样的 token 会被解析为布尔值。应用参数通过 `text` 转换会更明确。

### 类型、转换与操作符

- `vartype_type` 是描述存储类别的枚举；`vartype_type(vartype)` 与 `vartype_len(vartype)` 辅助函数分别暴露类别和存储长度。
- 显式转换覆盖常见整数与浮点类型、`boolean`、`text`/字符类型、`date`、`timestamp` 和 `timestamptz`，并在已定义的位置支持双向转换。
- 比较操作符和默认 `vartype_ops` B-tree 操作符类支持排序与索引。

不同类别组的值按照扩展内部优先级排序。整数与浮点值设计为数值比较，日期/时间戳值通过转换比较；这些规则属于扩展特有语义，而不是 PostgreSQL 普通单一类型语义。

### 关键比较风险

核验的 `0.4` C 比较器对混合数值操作数的处理不对称：整数对浮点数的一个分支会把浮点数转换成整数，反向分支却使用浮点比较。因此小数值可能产生不一致的比较结果，并违反 `vartype_ops` 所要求的排序约定。

在该行为得到修复，或针对选定构建被确定复现为安全之前，不要把混合整数/浮点 `vartype` 值用于 B-tree 索引、唯一约束、有序聚合或范围谓词。让索引值保持在同一类别可避开这条特定路径，但不能替代完整验证。

该扩展不可重定位；本次核验的代码库最后修改于 2017 年，文档主要面向 PostgreSQL 9.6。生产使用前应测试输入解析、转换、时区行为、比较律、备份恢复以及确切 PostgreSQL 版本。
