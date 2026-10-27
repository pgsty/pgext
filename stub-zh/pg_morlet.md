## 用法

来源：

- [README](https://github.com/merkhet-io/morletdb/blob/8b9e0af38ec2dffef6c6b5ae69982a29920a352a/extensions/pg_morlet/README.md)
- [Control](https://github.com/merkhet-io/morletdb/blob/8b9e0af38ec2dffef6c6b5ae69982a29920a352a/extensions/pg_morlet/pg_morlet.control)
- [Changelog](https://github.com/merkhet-io/morletdb/blob/8b9e0af38ec2dffef6c6b5ae69982a29920a352a/CHANGELOG.md)
- [Durable ingestion](https://github.com/merkhet-io/morletdb/blob/8b9e0af38ec2dffef6c6b5ae69982a29920a352a/docs/durable-ingestion.md)
- [SQL upgrade](https://github.com/merkhet-io/morletdb/blob/8b9e0af38ec2dffef6c6b5ae69982a29920a352a/extensions/pg_morlet/sql/pg_morlet--0.1.0--0.2.0.sql)

`pg_morlet` 0.2.0 保存等间隔采样的波形记录，利用有版本标识的 Morlet 特征检索相似区间。这一开发预览版面向 PostgreSQL 17，使用 pgrx 0.16.1，安装到 `public`，需要超级用户安装，不要求共享预加载或后台工作进程。

### 记录与搜索

```sql
CREATE EXTENSION pg_morlet;
SELECT morlet_create_stream_collection('public', 'signals');
SELECT morlet_start_recording(
  'signals', 'motor-001', 4096, 'vibration', 'g', '{"asset":"motor-1"}'
);
SELECT morlet_append_chunk(
  'signals', 'motor-001', 0,
  ARRAY(SELECT sin(i::float8 / 8) FROM generate_series(0,2047) i)
);
SELECT * FROM morlet_search_streams(
  'signals',
  ARRAY(SELECT sin(i::float8 / 8) FROM generate_series(0,511) i),
  4096, 'vibration', 'g', '{"asset":"motor-1"}'
);
```

### 一致性与权限

`morlet_append_chunk` 在同一事务发布原始块、特征和提取检查点。相同重试验证已提交输入；发生变化的重试、缺口和重叠会报错。`morlet_search_interval` 用已存区间检索，`morlet_read_interval` 重建样本，`morlet_stream_status` 报告摄取进度。集合表遵循调用者的普通权限和行安全策略，直接修改受管理表可能破坏跨表约束。

### 限制与升级

检索只比较采样率、通道、单位、窗口及提取器版本兼容的数据，精确扫描候选，默认预算为 4096 个段，返回左闭右开的样本区间。每块接受 1–32768 个有限样本，检索窗口为 512–4096 个样本。目前没有重采样、异步提取、ANN 索引或同步多通道检索；相似度是调查线索，不是经过校准的置信度。

SQL 版本 0.2.0 对应发行版 0.2.0-alpha.1。经过测试的升级脚本接受 0.1.0-alpha.3 所带的 0.1.0 SQL；0.1.0-alpha.1 安装需重新安装。已有 float8[] 集合保持原样；将其重新编码为带校验和的无损块需要单独传输数据，不属于 ALTER EXTENSION 的工作。


### 0.2.0 预览版存储与上下文

新的流集合采用无损编码块并验证 CRC-32C。纯原始记录、物化包络和依上下文比较属于可选工作流；相似度不是经过校准的故障概率。该版本仍是 PostgreSQL 17 开发预览，生产容量和并发负载尚未得到验证。
