## 用法

来源：

- [项目 README](https://github.com/melanieplageman/format/blob/837b6dcc957d2190105c819181eb19c0b1ad9b76/README.md)
- [扩展 control 文件](https://github.com/melanieplageman/format/blob/837b6dcc957d2190105c819181eb19c0b1ad9b76/format.control)
- [SQL 与 C 实现](https://github.com/melanieplageman/format/tree/837b6dcc957d2190105c819181eb19c0b1ad9b76/sql)
- [回归示例](https://github.com/melanieplageman/format/blob/837b6dcc957d2190105c819181eb19c0b1ad9b76/test/sql/base.sql)

`format` 0.0.1 是一个已弃用的 C 扩展，它重载 `format(text, hstore)`，用 `hstore` 中的命名值执行替换。上游要求用户改用后继项目 `format_x`；新代码应优先使用该后继项目或 PostgreSQL 内置格式化功能。

### 命名替换

安装本扩展前先安装所需的 `hstore` 依赖：

```sql
CREATE EXTENSION hstore;
CREATE EXTENSION format;

SELECT format(
  '%(name)I = %(value)L',
  hstore(ARRAY['name', 'status', 'value', 'ready'])
);
```

替换项使用 `%(key)s`、`%(key)I` 或 `%(key)L`，分别表示普通字符串、SQL 标识符和 SQL 字面量。格式化器也实现宽度和对齐标志。构造 SQL 时应使用标识符或字面量模式；普通字符串替换不会执行 SQL 转义。

### 兼容性与迁移

该函数标记为 strict 和 immutable。键缺失或值为 null 时可能产生警告；null 标识符会报错，而字面量格式化可以产生 `NULL`。依赖它之前，应对每个模板测试缺失值和 null 输入。

上游仓库实际上已停止维护，并明确将此实现标为弃用。其重载名称在审查和迁移时还容易与 PostgreSQL 内置 `format()` 函数混淆。不要新增对 0.0.1 版的依赖；应把现有模板迁移到 `format_x` 或明确的内置格式化方式，并在删除扩展前比较转义与 null 行为。
