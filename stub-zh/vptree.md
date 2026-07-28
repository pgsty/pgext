## 用法

来源：

- [官方上游 README](https://github.com/fake-name/pg-spgist_hamming/blob/9fa70b08e0f0108de6a6673ce095c86a987d261d/README.md)
- [官方扩展控制文件 (vptree.control)](https://github.com/fake-name/pg-spgist_hamming/blob/9fa70b08e0f0108de6a6673ce095c86a987d261d/vptree/vptree.control)
- [官方扩展 SQL (vptree--1.0.sql)](https://github.com/fake-name/pg-spgist_hamming/blob/9fa70b08e0f0108de6a6673ce095c86a987d261d/vptree/vptree--1.0.sql)

`vptree` — VP-tree 实现。使用它来处理相应的向量、模型或检索工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION vptree;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `vptree_area_match(int8, vptree_area)` 是一个扩展函数，返回 `boolean`。
- `vptree_choose(internal, internal)` 是一个扩展函数，返回 `void`。
- `vptree_config(internal, internal)` 是一个扩展函数，返回 `void`。
- `vptree_eq_match(int8, int8)` 是一个扩展函数，返回 `boolean`。
- `vptree_get_distance(int8, int8)` 是一个扩展函数，返回 `float8`。
- `vptree_inner_consistent(internal, internal)` 是一个扩展函数，返回 `void`。
- `vptree_leaf_consistent(internal, internal)` 是一个扩展函数，返回 `boolean`。
- `vptree_picksplit(internal, internal)` 是一个扩展函数，返回 `void`。
- `vptree_area` 是一个扩展定义的类型。
- `vptree_ops` 是一个扩展定义的操作类。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
