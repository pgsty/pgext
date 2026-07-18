## 用法

来源：

- [已复核 commit 的 nonoms README](https://github.com/Edwards-R/nonoms/blob/4ef6c5fcf05f0cf03a6ed852a20d2f791f539c5c/readme.md)
- [已复核 commit 的 nonoms 1.2.0 安装 SQL](https://github.com/Edwards-R/nonoms/blob/4ef6c5fcf05f0cf03a6ed852a20d2f791f539c5c/nonoms--1.2.0.sql)
- [已复核 commit 的 nonoms 初始化过程](https://github.com/Edwards-R/nonoms/blob/4ef6c5fcf05f0cf03a6ed852a20d2f791f539c5c/procedure/init_nonoms.sql)

`nonoms` 1.2.0 按 NoNomS 协议实现规范化的生物命名模型。它创建各分类等级专用的表，并跟踪当前名称、异名、组合、合并与拆分。应安装到专用空模式，先初始化顶级分类，再按父级优先顺序添加等级。

```sql
CREATE SCHEMA nomenclature;
CREATE EXTENSION nonoms WITH SCHEMA nomenclature;

CALL nomenclature.init_nonoms('Example scheme', 2026);
CALL nomenclature.insert_rank('kingdom', 1, 'Kingdom');
CALL nomenclature.insert_current_understanding(
  2, 1, 'Animalia', 'Linnaeus', 1758
);

SELECT * FROM nomenclature.rank ORDER BY id;
SELECT * FROM nomenclature.kingdom ORDER BY id;
```

`init_nonoms` 创建 `rank` 表和初始顶级分类；之后的过程会为每个等级动态创建一张命名表和一张组成表。

### 注意事项

- 上游明确称该项目处于 beta 阶段；应按应用领域模型验证其命名规则和迁移行为。
- 初始化面向全新的扩展模式；除非使用覆盖参数，否则会拒绝已填充的模式。未复核所有既有对象前，不要启用覆盖。
- 等级操作执行动态 DDL，并创建触发器、外键与索引。初始化和结构变更应作为受控迁移执行。
- 已复核仓库没有声明许可证，也没有当前 PostgreSQL 大版本兼容矩阵。
