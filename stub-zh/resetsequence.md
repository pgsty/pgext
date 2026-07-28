## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/resetsequence/resetsequence-1.0.0/README.md)
- [官方扩展控制文件 (resetsequence.control)](https://api.pgxn.org/src/resetsequence/resetsequence-1.0.0/resetsequence.control)
- [官方扩展 SQL (resetsequence--1.0.0.sql)](https://api.pgxn.org/src/resetsequence/resetsequence-1.0.0/resetsequence--1.0.0.sql)

`resetsequence` 模块包含一个 PostgreSQL 扩展，提供用于维护序列值的简单实用函数。这些函数可以把序列重置为通过 DEFAULT 子句与其关联的表列中已经使用的最大值。它适用于管理或自动化上述数据库行为。请以上方链接的固定上游修订为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION resetsequence;
```

在目标数据库中安装扩展；如果上游提供了最小示例，请运行该示例，并在集成到应用 SQL 前验证安装版本和返回值。

### 重要对象

- `resetseq_reset_sequences_in_database()` 是扩展函数，返回 `setof`。
- `resetseq_reset_sequences_in_schema(name)` 是扩展函数，返回 `setof`。
- `resetseq_reset_sequences_in_table(regclass)` 是扩展函数，返回 `setof`。
- `resetseq_sequence_max_value(oid)` 是扩展函数，返回 `bigint`。
- `resetseq_report_type` 是扩展定义的类型。

### 要求与注意事项

- 审阅的控制文件声明默认版本为 `1.0.0`。
- 控制文件将该扩展标记为可重定位。
- 生产使用前，请根据固定版本源码确认权限、支持的 PostgreSQL 版本、升级行为和失败情形。
