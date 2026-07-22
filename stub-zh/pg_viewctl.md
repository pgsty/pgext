## 用法

来源：

- [官方 README](https://github.com/boringsql/pg_viewctl/blob/adf2c5e2c4ec23148ca16d10c4de3ebe1a0c5044/README.md)
- [导出函数实现](https://github.com/boringsql/pg_viewctl/blob/adf2c5e2c4ec23148ca16d10c4de3ebe1a0c5044/src/functions.rs)
- [官方模式演进练习](https://github.com/boringsql/pg_viewctl/blob/adf2c5e2c4ec23148ca16d10c4de3ebe1a0c5044/try/exercise.sql)

`pg_viewctl` 是一个实验性 pgrx 扩展，用于分析列级视图依赖，并为模式变更生成有序 SQL 计划。它帮助审查替换视图、删除或修改表列以及重命名视图列的影响；生成的迁移不会自动执行。

### 变更前分析

安装 0.1.0 版，检查依赖，再让扩展生成计划：

```sql
CREATE EXTENSION pg_viewctl;

SELECT *
FROM pg_viewctl.get_column_dependencies('app', 'orders')
ORDER BY source_column, dependent_view;

SELECT *
FROM pg_viewctl.analyze_drop_column('app', 'orders', 'legacy_code');

SELECT step, operation, sql, target
FROM pg_viewctl.generate_drop_column('app', 'orders', 'legacy_code')
ORDER BY step;
```

应把返回行视为迁移草案。把 SQL 复制到迁移文件前，要检查对象所有权、安全选项、授权、注释、物化视图刷新、列引用、锁与数据丢失风险。

### 生成器与依赖对象

- `generate_replace_view` 规划依赖视图/物化视图的拆除、目标替换、重建、授权与刷新。
- `generate_drop_column` 规划删除表列时的依赖处理。
- `generate_alter_type` 规划列类型变更与依赖重建。
- `generate_rename_view_column` 规划视图列重命名，并标注需要手工修改的依赖定义。
- `pgvc_column_dependencies` 与 `get_column_dependencies` 暴露列级关系，`pgvc_dependency_order` 则生成递归视图顺序。

生成器返回 `step`、`operation`、`sql` 与 `target` 列。生成的注释或手工编辑标记要求人工处理，并不能证明迁移可安全执行。

### 弃用流程

扩展可以在移除前把视图列记录为弃用：

```sql
SELECT pg_viewctl.deprecate_column(
  'app', 'v_orders', 'legacy_code',
  'Use external_code instead',
  '2026-12-01'::date
);

SELECT * FROM pg_viewctl.get_deprecated_columns('app');
SELECT pg_viewctl.check_column_deprecated('app', 'v_orders', 'legacy_code');
```

`undeprecate_column` 删除标记，`pgvc_deprecated_with_dependents` 汇总仍存在的依赖视图。该元数据只起提示作用，不会阻止查询继续使用该列。

### 边界

官方 README 把项目标为开发中，并称接口很可能变化。控制文件将其标为不可重定位，不要求超级用户或预加载；源码特性目标为 PostgreSQL 13 至 18，默认构建版本为 PostgreSQL 18。

依赖发现基于 PostgreSQL 目录和已保存视图定义。动态 SQL、应用代码、外部消费者、函数体引用以及未体现在这些目录中的依赖都可能漏掉。每个生成计划都必须在接近生产的恢复库上测试，通过正常变更控制执行，并保留回滚路径。
