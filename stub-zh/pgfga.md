## 用法

来源：

- [官方 README](https://github.com/craigpastro/pgfga/blob/main/README.md)
- [面向 SQL 的实现](https://github.com/craigpastro/pgfga/blob/main/src/lib.rs)
- [授权求值器](https://github.com/craigpastro/pgfga/blob/main/src/check.rs)
- [存储实现](https://github.com/craigpastro/pgfga/blob/main/src/storage.rs)

`pgfga` 0.1.0 是一个受 Zanzibar 启发、用 pgrx 编写的实验性关系型访问控制实现。它存储授权 schema 与关系元组，然后通过固定的 pgfga schema 执行权限检查。上游将其标为缺少校验的进行中项目；不能把它作为唯一的生产授权边界。

### 核心流程

```sql
CREATE EXTENSION pgfga;

SELECT pgfga.create_schema(
  '{"namespaces":{"document":{"relations":{"viewer":[{"namespace":"user"}]},"permissions":{"can_view":{"computedUserset":"viewer"}}},"user":{"relations":{},"permissions":{}}}}'::json
) AS schema_id \gset

SELECT pgfga.create_tuple(
    :'schema_id'::uuid,
    'document', 'policy-1', 'viewer',
    'user', 'anya', ''
);

SELECT pgfga.check(
    :'schema_id'::uuid,
    'document', 'policy-1', 'can_view',
    'user', 'anya', ''
);
```

示例中的 psql 命令把创建 schema 返回的 UUID 保存到变量中，供后续调用使用。

### 函数索引

- `pgfga.create_schema(json)` 保存命名空间、关系和权限模型，并返回其 UUID。
- `pgfga.read_schema(uuid)` 和 `pgfga.read_schemas()` 用于查看已保存模型。
- `pgfga.create_tuple(...)`、`pgfga.read_tuples(...)` 和 `pgfga.delete_tuple(...)` 管理关系元组；读取过滤器中的空字符串匹配任意值。
- `pgfga.check(...)` 判断主体是否对资源具有某个关系或权限；求值器支持 computed userset、tuple-to-userset 遍历、并集、交集和排除。

### 关键限制

- 上游明确说明没有 schema/元组校验，其路线图也仍计划增加适当索引和基于迭代器的读取。因此无效元组可能被持久化，大型检查也可能效率很低。
- 扩展仅允许超级用户安装、不可迁移，并固定到 pgfga schema。所有变更函数应与读取/检查权限分开保护。
- 授权数据和应用数据不会自动由约束或分布式策略事务耦合。应用必须定义生命周期清理、审计、备份和失败关闭行为。
- 上游仓库已归档，README 也将项目称为实验性。如果要评估，必须冻结并审查确切源码版本。
