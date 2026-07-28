## 用法

来源：

- [官方上游 README](https://github.com/open-gpdb/heapy/blob/3b16e36b362be564492a61d1f01e66086167e453/README.HOT)
- [官方扩展控制文件 (heapy.control)](https://github.com/open-gpdb/heapy/blob/3b16e36b362be564492a61d1f01e66086167e453/heapy.control)
- [官方扩展 SQL (heapy--1.0.sql)](https://github.com/open-gpdb/heapy/blob/3b16e36b362be564492a61d1f01e66086167e453/heapy--1.0.sql)

`heapy` — Heap Only Tuple (HOT) 特性消除了冗余的索引条目，并允许在删除或过时的更新元组后重用空间，而无需进行全表真空。它通过允许单页真空，也称为“碎片整理”来实现这一点。请在相应的分析或存储工作流中使用此特性。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION heapy;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `heapy_define_relation_offload_policy_internal(reloid OID)` 是一个扩展函数，返回 `void`。
- `heapy_define_relation_offload_policy_internal_seg(reloid OID)` 是一个扩展函数，返回 `void`。
- `heapy_tableam_handler(internal)` 是一个扩展函数，返回 `table_am_handler`。
- `heapy` 是一个扩展定义的访问方法。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，与固定源进行比对。
