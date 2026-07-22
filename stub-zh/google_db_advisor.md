## 用法

来源：

- [AlloyDB 索引顾问官方文档](https://docs.cloud.google.com/alloydb/docs/use-index-advisor)

`google_db_advisor` 是 AlloyDB 托管的索引顾问扩展。它分析已记录的工作负载语句，并推荐普通索引或向量索引；它不会自动创建建议的索引。

### 核心流程

按照 AlloyDB 服务流程启用扩展，运行有代表性的工作负载，然后按需发起分析：

```sql
SELECT google_db_advisor_recommend_indexes();

SELECT *
FROM google_db_advisor_recommended_indexes;
```

应用任何 DDL 前，应检查推荐所依据的证据：

```sql
SELECT * FROM google_db_advisor_workload_report;
SELECT * FROM google_db_advisor_workload_statements;
```

保存结果后，或在开始新的独立观测窗口前，重置顾问状态：

```sql
SELECT google_db_advisor_reset();
```

### 重要对象

- `google_db_advisor_recommend_indexes()` 发起一次按需分析。
- `google_db_advisor_recommended_indexes` 展示当前角色可见的索引 DDL 建议与详细信息。
- `google_db_advisor_workload_report` 汇总被分析的工作负载。
- `google_db_advisor_workload_statements` 展示顾问使用的语句。
- `google_db_advisor_reset()` 清除顾问收集的工作负载和推荐。

### 服务与安全边界

- 这是 AlloyDB 服务组件，并非可移植的社区软件包。可用性、启用方式、版本和权限均以所选 AlloyDB 实例及其当前文档为准。
- 应在收集窗口中覆盖全部重要查询形态。不完整或缺乏代表性的工作负载会产生不完整建议。
- 生成的 DDL 只是建议：执行前需检查重复或重叠索引、写放大、存储、构建时间、锁影响以及后续删除成本。
- 结果会按查询角色过滤。应使用明确的审查角色，不要把某个用户看到的结果当作全实例清单。
- 向量索引建议还有 Google 文档所列的服务开关与预览限制；用于生产环境前应核对当前服务要求。
