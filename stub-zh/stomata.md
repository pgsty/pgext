## 用法

来源：

- [STOMATA 0.1.0 README](https://github.com/CrystallineCore/Stomata/blob/0193b0cdc90a6d9866553f0788a51a462a8fb7ec/README.md)
- [Control](https://github.com/CrystallineCore/Stomata/blob/0193b0cdc90a6d9866553f0788a51a462a8fb7ec/stomata.control)
- [SQL 0.1.0](https://github.com/CrystallineCore/Stomata/blob/0193b0cdc90a6d9866553f0788a51a462a8fb7ec/sql/stomata--0.1.0.sql)

`stomata` 0.1.0 使用分段自动机为 `LIKE` 和 `ILIKE` 建立索引，适用于带首尾锚点、短字面量或下划线通配符的模式。上游将首个版本标为测试阶段；要求 PostgreSQL 16 或更高版本，并说明已在 16、17、18 上测试。

### 基本用法

```sql
CREATE EXTENSION stomata;
CREATE TABLE contacts (id bigint, email text);
CREATE INDEX contacts_email_stomata ON contacts USING stomata (email);

SELECT * FROM contacts WHERE email LIKE 'jo%@example.com';
SELECT * FROM contacts WHERE email ILIKE 'JO%';
SELECT * FROM stomata_index_info('contacts_email_stomata');
SELECT stomata_verify('contacts_email_stomata');
```

使用超级用户安装；control 文件未将扩展标为可信。每个索引只支持一个 text 或 varchar 列或表达式，支持表达式索引和部分索引。执行器会重新检查候选记录是否满足模式。

### 对象与调优

- `stomata_index_info`、`stomata_runs`、`stomata_key_stats`：检查数据段、待合并记录、大小和键类别。
- `stomata_keys`、`stomata_pattern_keys`、`stomata_pattern_stats`：检查值或模式生成的键。
- `stomata_candidate_pages`、`stomata_estimate`：检查候选页面与规划器估计。
- `stomata_verify`：统计索引未覆盖的可见记录，正常结果应为零。
- `stomata_merge_pending`、`stomata_compact`：合并待处理记录，或将所有数据段整合。
- 索引选项包括 `k`（段宽，默认 3）、`exact`（逐行记录，默认开启）、`exact_bigrams`（默认关闭）、`rollup`（3）、`skip_depth`（4）和 `pending_limit`（4096 kB）。最后一项设为零时，由 vacuum 或显式维护执行合并。
- `stomata.estimate_budget` 限制规划阶段的估计工作量。通过 `ALTER INDEX` 修改索引选项后，执行 `REINDEX` 重建。

### 维护与限制

插入可能触发合并，增加语句延迟。Vacuum 和压缩整合可能重写数据段，产生较多 WAL；完整整合期间可能临时占用约两倍于有效索引的空间。释放的空间会被复用；要归还给操作系统，需要重建索引。

不支持仅索引扫描、排序、多列索引或并行构建。不为否定模式、正则表达式及相似度运算符建立索引。ASCII 大小写折叠会降低非 ASCII 不区分大小写查询的筛选精度；非确定性排序规则会退化为扫描所有页面。备库的页面回收冲突可能取消查询。1.0 之前，上游磁盘格式变化可能要求重建索引。
