## 用法

来源：

- [Official README](https://github.com/Edwards-R/nonoms/blob/4ef6c5fcf05f0cf03a6ed852a20d2f791f539c5c/readme.md)
- [Version 1.2.0 SQL implementation](https://github.com/Edwards-R/nonoms/blob/4ef6c5fcf05f0cf03a6ed852a20d2f791f539c5c/nonoms--1.2.0.sql)
- [Initialization documentation](https://github.com/Edwards-R/nonoms/blob/4ef6c5fcf05f0cf03a6ed852a20d2f791f539c5c/docs/procedure/init_nonoms.md)
- [Rank-creation documentation](https://github.com/Edwards-R/nonoms/blob/4ef6c5fcf05f0cf03a6ed852a20d2f791f539c5c/docs/procedure/insert_rank.md)

`nonoms` 1.2.0 按 NoNomS 协议实现规范化生物命名模型。它创建可配置的等级层次，并记录当前理解、异名、组成、合并和拆分。上游将项目标为 beta；必须根据消费应用的命名规则验证其模型。

### 初始化新命名体系

使用名为 `nomenclature` 的专用空模式。1.2.0 的拆分实现包含字面量 `nomenclature` 模式引用，因此实际并不能完整迁移到其他扩展模式。

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

`init_nonoms` 创建 `rank` 目录和 capstone 等级。载入名称之前必须按父级优先顺序创建全部等级；上游说明不支持在层次中间插入等级，也不支持在等级已经填充后再新增等级。

### 主要过程与类型

- `insert_rank`：创建一个等级表、组成表、索引、外键和触发器。
- `insert_current_understanding` 与 `insert_synonym_understanding`：添加当前名称和异名。
- `merge_understandings`：用合并后的理解替换多个当前理解，并传播子级。
- `split_understanding`：通过 `split_result[]` 创建目标及组成记录。
- `create_binomial_view`：等级结构完成后创建跨等级二名法视图。
- `rank`、各等级表和各等级 `_composition` 表：持久命名状态。

### 迁移与完整性边界

初始化会拒绝已填充模式，除非传入 `override => true`；没有完整对象审查和可恢复备份时不要使用该覆盖参数。等级创建和合并/拆分会执行动态 DDL 或递归数据更新，并创建约束、触发器、索引和临时表。应把它们作为受控迁移执行，而不是放进不可信应用请求。

已复核仓库的 1.2.0 源码没有当前 PostgreSQL 兼容矩阵，也没有声明的许可证文件；`split_understanding` 还存在硬编码模式假设。应固定源码，在真实层次数据副本上测试每个过程，并在生产采用前验证当前/异名指针、组成关系、子级传播、回滚、备份导出和恢复。
