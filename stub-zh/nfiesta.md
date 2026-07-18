## 用法

来源：

- [已复核版本的 README](https://gitlab.com/nfiesta/nfiesta_pg/-/blob/5bf6158866a327db97a1c8513f58a06228573ecb/README.md)
- [nFIESTA 项目 Wiki](https://gitlab.com/nfiesta/nfiesta_pg/-/wikis/home)
- [SQL 源码树](https://gitlab.com/nfiesta/nfiesta_pg/-/tree/5bf6158866a327db97a1c8513f58a06228573ecb/functions)

`nfiesta` 是森林资源清查估算与分析系统 nFIESTA 的 PostgreSQL 估算引擎。它把国家森林清查野外样本与辅助数据（通常是遥感产品）结合，实现单阶段与两阶段估算、配置、ETL、结果和可加性工作流。

```sql
CREATE EXTENSION nfiesta;
SELECT e.extname, e.extversion
FROM pg_extension AS e
WHERE e.extname = 'nfiesta';
```

扩展安装的是大型领域模式，而非一个独立函数。应按照精确版本的上游安装指南和最小工作示例，加载有文档的测试数据，配置估算期与模型，再把总量、比率、方差和置信结果与独立计算的参考值比较。

统计正确性取决于抽样设计、分层、权重、辅助变量、缺失数据策略、模型配置和具体版本 SQL。应保留输入血缘与配置，隔离 ETL 和 GUI 角色，审计 security-definer 与动态 SQL 函数，并固定扩展版本以保证报告可复现。升级可能改变大量函数、视图、配置表和结果代次，必须在副本上测试迁移与回滚。
