## 用法

来源：

- [已复核 revision 的上游 README](https://bitbucket.org/adunstan/pg-closed-ranges/src/9a736068119bb0f4fa5d59c7cb41198a248a2b04/README.md)
- [扩展 control 文件](https://bitbucket.org/adunstan/pg-closed-ranges/src/9a736068119bb0f4fa5d59c7cb41198a248a2b04/cranges.control)

`cranges` 为 PostgreSQL 的三个内置离散范围类型增加闭区间变体：`cdaterange`、`cint4range` 和 `cint8range`。其规范表示包含两个端点，不同于内置的半开 `[)` 表示。

```sql
CREATE EXTENSION cranges;
SELECT '[1,5]'::cint4range;
SELECT '[2026-01-01,2026-01-31]'::cdaterange;
```

这些类型在其他方面复用 PostgreSQL 范围行为，包括运算符与索引支持。双参数构造器仍被 PostgreSQL 固定为假设 `[)`，因此 `cint4range(1,5)` 会规范化为 `[1,4]`。上游建议使用显式三参数构造器或字面量语法，以避免这种意外。

项目没有当前发行版或 PostgreSQL 兼容矩阵，并托管在 Bitbucket 而非 GitHub。将类型用于持久 schema 前，应在目标 PostgreSQL 大版本上测试转换、运算符、索引、转储与升级。
