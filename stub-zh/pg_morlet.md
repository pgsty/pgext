## 用法

来源：

- [v0.3.0-alpha.1 MorletDB README](https://github.com/merkhet-io/morletdb/blob/e85a24d1775cc971b2c4b444fbf3863910f8582c/README.md)
- [pg_morlet README](https://github.com/merkhet-io/morletdb/blob/e85a24d1775cc971b2c4b444fbf3863910f8582c/extensions/pg_morlet/README.md)
- [控制文件](https://github.com/merkhet-io/morletdb/blob/e85a24d1775cc971b2c4b444fbf3863910f8582c/extensions/pg_morlet/pg_morlet.control)
- [SQL 升级脚本](https://github.com/merkhet-io/morletdb/blob/e85a24d1775cc971b2c4b444fbf3863910f8582c/extensions/pg_morlet/sql/pg_morlet--0.2.0--0.3.0.sql)

`pg_morlet` 0.3.0 保存波形记录，并使用有版本标识的 Morlet 特征比较兼容区间。0.3.0-alpha.1 是面向 PostgreSQL 17、使用 pgrx 0.16.1 的开发预览版，在同步事务存储模型上增加采集时间线和依上下文比较。

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

### 升级边界

SQL 版本为 0.3.0。测试过的升级脚本接受 0.2.0-alpha.1 的 SQL 0.2.0，并可继续从 0.1.0-alpha.3 的 SQL 0.1.0 链式升级。更早的 0.1.0-alpha.1 安装需要重新安装。该版本仍是开发预览，没有稳定的模式契约或生产容量保证。
