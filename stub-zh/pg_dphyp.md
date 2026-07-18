## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/TantorLabs/pg_dphyp/blob/f76770ca001acca74f80ff6111d224aa44f415ff/README.md)
- [扩展 control 文件](https://github.com/TantorLabs/pg_dphyp/blob/f76770ca001acca74f80ff6111d224aa44f415ff/pg_dphyp.control)
- [连接搜索钩子实现](https://github.com/TantorLabs/pg_dphyp/blob/f76770ca001acca74f80ff6111d224aa44f415ff/src/pg_dphyp.c)

`pg_dphyp` 是一个无 SQL 对象的规划器模块，实现基于超图的动态规划连接枚举。它通过加载库启用；目录没有记录 SQL 扩展对象。

```conf
shared_preload_libraries = 'pg_dphyp'
pg_dphyp.enabled = on
pg_dphyp.min_relations = 8
pg_dphyp.max_relations = 16
pg_dphyp.cj_strategy = 'pass'
```

该模块可缩小部分多表连接查询的搜索空间，但上游提醒它找到的连接顺序不一定最优。交叉连接与不连通关系集合需要特殊处理；超过配置阈值后，会回退到 PostgreSQL 动态规划或 GEQO。其他控制项包括 `pg_dphyp.count_cc` 与 `pg_dphyp.geqo_cc_threshold`。

规划器钩子和内部结构对大版本很敏感。已复核源码声称兼容 PostgreSQL 12 至 17，但也说明没有真正的测试套件，计划输出测试依赖具体版本。生产启用前，应分别基准规划与执行时间，对比禁用模块后的计划，并测试外连接、横向引用、不连通图、预备语句、分区及升级。
