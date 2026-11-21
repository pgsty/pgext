## 用法

来源：

- [Beta release](https://github.com/merkhet-io/morletdb/releases/tag/v0.4.0-beta.1)

- [Official source](https://github.com/merkhet-io/morletdb/blob/v0.4.0-beta.1/extensions/pg_morlet/pg_morlet.control)
- [Official source](https://github.com/merkhet-io/morletdb/blob/v0.4.0-beta.1/docs/compatibility.md)
- [Official source](https://github.com/merkhet-io/morletdb/blob/v0.4.0-beta.1/extensions/pg_morlet/sql/pg_morlet--0.3.11--0.3.12.sql)
- [Official source](https://github.com/merkhet-io/morletdb/blob/v0.4.0-beta.1/extensions/pg_morlet/sql/pg_morlet--0.3.0--0.3.1.sql)
- [Official source](https://github.com/merkhet-io/morletdb/blob/v0.4.0-beta.1/extensions/pg_morlet/README.md)
- [Official source](https://github.com/merkhet-io/morletdb/blob/v0.4.0-beta.1/docs/recording-workflow.md)

`pg_morlet` 0.3.12 保存波形记录，并使用有版本标识的 Morlet 特征比较兼容区间。v0.4.0-beta.1 源码标签实际包含该 control/SQL 版本。PostgreSQL 17 是参考平台，18 有功能 CI 覆盖，19 的验证基于候选发布版。该扩展仍属实验性质，使用 pgrx 0.19.3。

### 记录与搜索

```sql
CREATE EXTENSION pg_morlet;
SELECT morlet_create_stream_collection('public', 'signals');
SELECT morlet_enable_acquisitions('signals'::regclass);
SELECT morlet_create_acquisition('signals'::regclass, 'run-001', 'motor-1');
```

采集模型在整数纳秒时间线上对齐通道，并记录偏移、时钟来源和不确定度。`morlet_build_acquisition_segments`、`morlet_read_acquisition` 与 `morlet_compare_acquisition` 提供新的对齐工作流。版本化 SQL 文件定义了采集设置与函数签名；执行比较还需要已有记录及构建好的上下文分段。扩展不会重采样通道，也不估算通道间时钟偏移。

### 一致性与权限

摄取和特征发布仍在同一事务中完成。集合表遵循调用者权限和行安全策略，直接修改可能破坏跨表不变量。搜索只比较采样率、通道、单位、窗口及提取器版本兼容的数据。相似度是调查信号，不是经过校准的概率。

### 连续记录与维护

当前流程使用 `morlet_start_recording`、`morlet_append_chunk` / `morlet_append_chunks`、`morlet_read_interval` 和 `morlet_search_interval`。区间读取按调用者权限重组数据块；不兼容或部分不可见的区间会明确报错。可选包络探针需要 `vector`，普通存储和精确搜索不需要。

0.3.11 到 0.3.12 的 SQL 更新修正了 `morlet_build_envelopes` 对已压缩片段的计数，以及 `morlet_collection_status` 的积压计数。更新前先安装匹配的库与 SQL 文件：

```sql
ALTER EXTENSION pg_morlet UPDATE TO '0.3.12';
```

当前标签提供版本化升级链。上游只在 PostgreSQL 17 测试过从旧 alpha 版本升级，18/19 使用全新安装。部分 README 段落仍描述没有升级路径的旧 0.1.0 框架，应以本修订的 control、SQL 升级链和兼容文档为准。保留记录和备份，不要为遵循旧演示而删除已有数据的扩展。

### Beta 运行限制

0.4.0-beta.1 发行版保留 SQL 版本 0.3.12。`morlet_latest` 按匹配的记录各返回一行，而非整个数据流只返回一行。上游持续摄取试验保留了核对的样本，但没有达到稳定运行与搜索目标，包络和验证任务出现积压。应监控积压，并针对实际负载验证恢复和查询延迟。Beta 版的 API 与数据块格式承诺不代表已经具备生产容量。
