## 用法

来源：

- [Official upstream documentation](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/using-the-index-adviser-extension-on-an-apsaradb-rds-for-postgresql-instance)
- [Official primary documentation](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`index_adviser` 2.0 是阿里云 RDS PostgreSQL 扩展，根据已规划或已执行的查询估算有用的 B-tree 索引。它属于云厂商托管软件，建议只是估算结果，执行生成的 DDL 前必须人工审查。

### 核心流程

创建扩展，然后选择在当前会话加载，或者配置 `shared_preload_libraries` 并重启实例，使其对所有会话生效。

```sql
CREATE EXTENSION index_adviser;
LOAD 'index_adviser';

EXPLAIN SELECT * FROM orders WHERE customer_id = 42;
SELECT * FROM index_advisory;
SELECT show_index_advisory(NULL);
```

普通 `EXPLAIN` 不执行查询，但会记录建议。正常执行的查询会在运行期间被分析。`index_advisory` 保存逐查询原始估算；`show_index_advisory(pid)` 将指定会话建议格式化成 `CREATE INDEX` 语句；`select_index_advisory` 汇总所有已记录会话。

### 要求与限制

当前厂商文档要求在创建或重建扩展时，RDS PostgreSQL 小版本至少为 20230830。不要在只读事务中使用。它只建议单列或复合 B-tree 索引，不支持 GIN、GiST 或哈希索引。应用建议前，应核验工作负载代表性、写放大、存储、重复索引、锁影响与维护成本。
