## 用法

来源：

- [官方资料](https://github.com/aotenjou/mem_collector/blob/ed19a382495eeffe279497820340a2eb8a78b978/README.md)
- [官方资料](https://github.com/aotenjou/mem_collector/blob/ed19a382495eeffe279497820340a2eb8a78b978/mem_collector.control)
- [官方资料](https://github.com/aotenjou/mem_collector/blob/ed19a382495eeffe279497820340a2eb8a78b978/mem_collector--1.0.sql)

`mem_collector` 1.0 为研究型内存调优流程采集查询内存特征与窗口汇总。上游针对 PostgreSQL 12，规划器和执行器钩子及共享内存缓冲区要求预加载并重启服务器。

### 配置与读取样本

配置工作进程写入样本的数据库：

```ini
shared_preload_libraries = 'mem_collector'
mem_collector.window_seconds = 30
mem_collector.enable_bgw = on
mem_collector.bgw_dbname = 'postgres'
```

重启后，由超级用户在该数据库中安装：

```sql
CREATE EXTENSION mem_collector;
SELECT * FROM mem_collector_features();
SELECT * FROM mem_collector_pending_samples();
SELECT mem_collector_flush_query_samples(1024);
SELECT * FROM mem_query_samples ORDER BY id DESC LIMIT 5;
SELECT * FROM mem_window_samples ORDER BY id DESC LIMIT 5;
```

对应版本的 SQL 将前两个读取函数实现为对近期持久化样本表的查询，不能根据函数名认定它们直接展示未刷出的环形缓冲区。`mem_collector_generate_window_sample` 显式生成窗口样本，`mem_collector_reset` 重置采集状态。应限制对查询和工作负载元数据的访问，并为持久化样本表制定保留策略。

### 结果解释与运维

采集特征覆盖排序、哈希算子和简化的并行比例。峰值内存及溢写信号是根据公开执行器结构得到的估计值。部分压力与内存池使用字段尚未实现，比例的分母采用主机物理内存，而非 PostgreSQL 内存池预算。它不能作为生产内存核算的权威依据，也不会自动限制内存。

共享内存环形缓冲区在重启后丢失，持久化样本表则保留。指向其他数据库的工作进程需要该库中也具备相应 SQL 对象。启用采集和升级后应确认样本持续推进，并监控表增长与采集开销。
