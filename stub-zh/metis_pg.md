## 用法

来源：

- [官方资料](https://github.com/DBHammer/Metis/blob/e3eed5d90f1a896258ff55c0602a92199bc9dee4/docs/POSTGRESQL_INTEGRATION.md)
- [官方资料](https://github.com/DBHammer/Metis/blob/e3eed5d90f1a896258ff55c0602a92199bc9dee4/pg_metis/metis_pg.control)
- [官方资料](https://github.com/DBHammer/Metis/blob/e3eed5d90f1a896258ff55c0602a92199bc9dee4/pg_metis/metis_pg--0.1.sql)

`metis_pg` 0.1 是研究型基数估计扩展，通过 PostgreSQL 规划器钩子提供工作负载训练得到的估计值。它可以在服务器中读取预先生成的 JSON 直方图，也可以调用外部 Java 估计服务。官方集成指南针对 PostgreSQL 14。

### 核心工作流

由超级用户安装，为目标关系准备统计数据，并在每个会话中加载动态库：

```sql
CREATE EXTENSION metis_pg;
LOAD 'metis_pg';
SELECT metis_pg_version();
SET metis.enable = on;
SET metis.estimator = 'c';
SET metis.stats_dir = '/work/metis_stats/dmv/current';
SELECT metis_reset_stats_cache();
SELECT metis_timing_stats();
```

指定目录必须包含针对实际关系训练的直方图，并允许 PostgreSQL 读取。训练需要 Java 21、Maven，以及单独安装并获得许可的 Gurobi 优化器；仓库不提供这一商业依赖。C 估计器读取训练生成的 JSON 文件，规划期间不需要 Java 服务。

### 外部估计与限制

选择 `metis.estimator = 'http'` 使用外部服务，并通过 `metis.host`、`metis.port` 和 `metis.timeout_ms` 配置正确端点及请求预算。请求会将关系和谓词信息发送到后端之外，应保护服务及其统计数据。

`metis_estimate_selectivity` 和 `metis_estimate_selectivity_http` 提供直接估计接口，`metis_reset_timing_stats` 清理计时统计。规划器支持的谓词为单表列与常量之间的等值或包含边界的比较，并以逻辑与组合。其他谓词保留 PostgreSQL 估计。会话加载无需重启服务器，也不会激活其他会话的钩子。应使用代表性数据验证估计值和实际计划，不能假设工作负载训练模型能够普遍泛化。
