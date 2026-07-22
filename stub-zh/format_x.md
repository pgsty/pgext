## 用法

来源：

- [官方 format_x 文档](https://github.com/melanieplageman/format_x/blob/e65255bcb23aa17a4e00fab66d9e77908fc2b69b/doc/format_x.md)
- [0.0.1 版本 SQL 对象](https://github.com/melanieplageman/format_x/blob/e65255bcb23aa17a4e00fab66d9e77908fc2b69b/sql/format_x--0.0.1.sql)
- [官方 PGXN 发行页面](https://pgxn.org/dist/format_x/)

`format_x` 为 PostgreSQL `format()` 增加了对复合值、`json`、`jsonb` 与 `hstore` 的命名及链式查找。需要格式化结构化值、但不想先逐个提取字段时可以使用它。

### 命名查找

```sql
CREATE EXTENSION format_x;

SELECT format_x(
  'Hello %(name)s from %(address.city)s',
  '{"name":"Ada","address":{"city":"London"}}'::jsonb
);
```

查找形式为 `%(key)s`，点号用于遍历嵌套值。还可以把位置参数选择与查找结合，例如 `%2(name)s`。标准的 `%s`、`%I`、`%L` 转换、位置字段、宽度与精度也受到支持。

```sql
SELECT format_x(
  'SELECT * FROM %1(name)I WHERE code = %1(code)L',
  '{"name":"places","code":"GB"}'::jsonb
);
```

### 输入与注意事项

安装的重载是 `format_x(text)` 与 `format_x(text, VARIADIC "any")`。复合值、JSON 和 JSONB 可直接查找。只有传入 `hstore` 值时才需要 `hstore`，它不是声明的扩展依赖。

缺失键、不支持的输入类型以及部分 `NULL` 查找会报错。`%I` 会引用标识符并拒绝 `NULL`；`%L` 按 `quote_nullable` 语义引用字面量。引用字段并不能让任意格式模板自动安全：模板本身应可信，执行生成 SQL 时仍应使用参数。此版本使用 PostgreSQL 内部格式化与类型 API，因此部署前应在准确的 PostgreSQL 版本上验证构建。
