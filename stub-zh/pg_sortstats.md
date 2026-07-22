## 用法

来源：

- [官方 README](https://github.com/powa-team/pg_sortstats/blob/b9228cca5947b010c8114ecd5091e76d57ef28c3/README.md)
- [扩展 SQL API](https://github.com/powa-team/pg_sortstats/blob/b9228cca5947b010c8114ecd5091e76d57ef28c3/pg_sortstats--0.0.1.sql)
- [共享内存与 GUC 实现](https://github.com/powa-team/pg_sortstats/blob/b9228cca5947b010c8114ecd5091e76d57ef28c3/pg_sortstats.c)

`pg_sortstats` 0.0.1 按查询 ID、用户、数据库和文本排序键收集累计排序节点统计。它区分 top-N 堆排序、快速排序、外部排序/合并、磁盘与内存空间、并行 worker，并估算让已观察排序留在内存中所需的 `work_mem`。上游明确标记为开发中、尚未达到生产就绪状态。

### 预加载与配置

两个依赖都必须在启动时加载，并重启服务器：

```conf
shared_preload_libraries = 'pg_stat_statements,pg_sortstats'
pg_sortstats.max = 10000
pg_sortstats.save = on
```

随后在需要报告对象的数据库中安装：

```sql
CREATE EXTENSION pg_sortstats;

SELECT queryid, userid, dbid, sort_keys,
       external_sorts, space_disk, work_mems
FROM pg_sortstats
ORDER BY space_disk DESC;

SELECT * FROM pg_sortstats(false);

SELECT pg_sortstats_reset();
```

`pg_sortstats` 是 `pg_sortstats(true)` 上的视图。当不希望披露排序键文本或承担其收集成本时，可以传 false。计数器是累计总数，而不是每次执行值；应关联 `pg_stat_statements` 中的 `queryid`，并比较测量区间内的变化。

### GUC 与报告字段

`pg_sortstats.enabled` 是用户可设置开关，默认开启。`pg_sortstats.max` 是启动时最大跟踪条目数，默认 10,000，并影响共享内存。`pg_sortstats.save` 默认开启，控制是否跨正常关机/重启保存统计。

除身份与 `sort_keys` 外，函数还报告 `nb_keys`、输入 `lines`、`lines_to_sort`、估算 `work_mems`、各算法计数、`nb_tapes`、`space_disk`、`space_memory`、`non_parallels` 和 `nb_workers`。

### 运维边界

收集会增加执行器 hook 和共享内存开销。排序键文本与跨用户查询标识可能暴露工作负载结构，而 `pg_sortstats_reset()` 会清除共享统计；应撤销不应拥有这些能力的角色的报告/重置执行权限。查询 ID 依赖 `pg_stat_statements` 和适当的查询 ID 计算设置。

应把估算值作为调优证据，而不是直接设为集群级 `work_mem`：`work_mem` 会乘以并发操作和会话数量。部署前应在精确 PostgreSQL 大版本上验证扩展兼容性与计数准确性。
