## 用法

来源：

- [当前复核的研究限制](https://github.com/gburd/ra/blob/8166ea89e91bfef98e1c31fe140b9525d37a1d33/README.md)
- [Ra v0.6.0 README](https://github.com/gburd/ra/blob/v0.6.0/README.md)
- [扩展控制文件](https://github.com/gburd/ra/blob/v0.6.0/crates/ra-pg-extension/pg_ra_planner.control)
- [扩展 Cargo 清单](https://github.com/gburd/ra/blob/v0.6.0/crates/ra-pg-extension/Cargo.toml)
- [扩展 SQL](https://github.com/gburd/ra/blob/v0.6.0/crates/ra-pg-extension/sql/pg_ra_planner--0.1.0.sql)

`pg_ra_planner` 是 Ra 关系代数优化器接入 PostgreSQL 规划器钩子的实验性扩展。仓库标签为 v0.6.0，但控制文件和 SQL 版本为 0.1.0。这是研究预览版，上游 README 记录了已知的错误答案缺陷，不应在生产正确性场景使用。

### 启用预览版

```sql
CREATE EXTENSION pg_ra_planner;
```

服务器启动前必须通过 `shared_preload_libraries` 加载动态库，因此修改后需要重启。版本化引导 SQL 检查规划器 GUC 是否注册，Rust crate 还导出诊断函数。源码清单使用开发者本机的 pgrx 路径，本轮未证明可重现的可移植构建。`parser_hook` 特性还需要带有上游钩子补丁的 PostgreSQL 构建。

### 规划器控制项

Cargo 特性标记覆盖 PostgreSQL 13–19；这属于源码特性列表，并非本轮独立验证的兼容矩阵。

SQL 脚本会检查动态库是否注册了设置。文档列出的控制项包括 `ra_planner.enabled`、`ra_planner.min_confidence`、`ra_planner.log_decisions` 和 `ra_planner.max_relations`。

```sql
SET ra_planner.enabled = on;
SET ra_planner.min_confidence = 0.9;
SET ra_planner.log_decisions = on;
```

应在一次性测试集群上运行上游验证语料，再开始实验。保留 PostgreSQL 原生规划器回退，复核执行计划和结果，并将学习模型与计划输出视为实验性结果。
