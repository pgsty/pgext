## 用法

来源：

- [官方资料](https://github.com/brauliobo/pg_sift/blob/v4.1.4/README.md)
- [官方资料](https://github.com/brauliobo/pg_sift/blob/v4.1.4/docs/UPGRADING.md)
- [官方资料](https://github.com/brauliobo/pg_sift/blob/v4.1.4/pg_sift.control)
- [官方资料](https://github.com/brauliobo/pg_sift/blob/v4.1.4/sql/pg_sift--3.5.sql)
- [用法与安全](https://github.com/brauliobo/pg_sift/blob/v4.1.4/docs/USAGE.md)
- [支持版本](https://github.com/brauliobo/pg_sift/blob/v4.1.4/docs/SUPPORT.md)
- [发布说明](https://github.com/brauliobo/pg_sift/releases/tag/v4.1.4)
- [4.0.3 to 4.1 SQL](https://github.com/brauliobo/pg_sift/blob/v4.1.4/sql/pg_sift--4.0.3--4.1.sql)
- [Options](https://github.com/brauliobo/pg_sift/blob/v4.1.4/docs/OPTIONS.md)
- [4.3.2 release](https://github.com/brauliobo/pg_sift/releases/tag/v4.3.2)
- [4.3.2 upgrade guide](https://github.com/brauliobo/pg_sift/blob/v4.3.2/docs/UPGRADING.md)
- [4.3.2 changes](https://github.com/brauliobo/pg_sift/blob/v4.3.2/CHANGELOG.md)

`pg_sift` 发行版 4.1.4 安装的 SQL 扩展版本为 4.1，同时提供 BM25 文本检索、压缩列索引与列式表访问方法。上游支持 PostgreSQL 17–19，但这不代表每个大版本都有对应的软件包。首次安装需要超级用户，普通用法不强制要求预加载。

### 文本检索

```sql
CREATE EXTENSION pg_sift;
CREATE TABLE articles (id bigint PRIMARY KEY, body text);
INSERT INTO articles VALUES
  (1, 'PostgreSQL extensions add full text search'),
  (2, 'PostgreSQL indexing and query planning');
CREATE INDEX articles_search ON articles USING sift (body);
SELECT id, sift.score(ctid) AS score
FROM articles
WHERE body @@> 'postgresql AND search'
ORDER BY score DESC LIMIT 10;
```

`sift` 索引处理文本列或表达式，支持布尔、短语、前缀、通配符、正则与模糊查询。`sift.score(ctid)` 可用于扫描查询的输出或排序；也支持在过滤条件中与常量分数阈值比较。`sift.tokenize`、`sift.highlight` 与 `sift.facets` 分别提供分析、高亮和分面能力。`sift.sections` 与 `sift.field_query` 支持文档所述的字段权重流程，但不代表可以创建多列 BM25 索引。

### 列式分析

```sql
CREATE TABLE orders (id bigint PRIMARY KEY, category text, amount numeric(10,2));
CREATE INDEX orders_col ON orders USING sift_col (category, amount);
SELECT category, count(*), sum(amount) FROM orders GROUP BY category;
CREATE TABLE events (id bigint, kind text, amount numeric(10,2)) USING sift_table;
```

`sift_col` 为普通表增加索引，`sift_table` 则直接以列格式存储表。符合条件的聚合与过滤使用 `Custom Scan (SiftAgg)`，其他情况保留普通执行计划，应检查实际计划。`sift.enable_agg = off` 可关闭聚合下推以作比较。3.3 扩展了 top-N 排序、点查与小表连接，新建列索引默认启用预聚合。可选 Vulkan 路径仍属实验功能，默认关闭，上游软件包也未启用。

### 安全与维护

受行级安全策略约束的角色默认不使用排名扫描。超级用户可通过 `sift.rls_scores = on` 显式启用；返回行仍遵循策略，但评分依赖整个索引的统计信息，包括隐藏行。启用前应评估这一信息边界。安全屏障及其他不支持的查询结构仍受文档所述限制。

3.3 的安装与升级会将 `sift` 模式的使用权限授予 `PUBLIC`，各函数仍检查自身的表或索引访问权限。`sift.verify`、`sift.segments`、`sift.merge` 等维护操作需要表所有权或 MAINTAIN 权限。修改分析器或评分选项需要重建受影响的检索索引。

`sift.col_auto_merge` 与 `sift.table_auto_defrag` 现在默认开启。后台维护可能清理表并重写列索引，从而产生 WAL。列合并预算由 `sift.col_auto_merge_wal_pct` 控制，`sift.col_maintenance_status` 与 `sift.col_index_status` 可查看维护状态。关闭这两个自动维护开关可恢复旧版行为。

`write_buffer_limit` 替代旧名 `mutable_limit`，旧名保留为别名。备库回放索引截断时可能取消查询。通过逻辑复制发布列式表需要遵循上游预加载配置；基本文本检索和列式分析不要求预加载。

### 从 3.3 之前版本升级

先运行上游升级预检查，安装匹配文件并更新每个数据库，再重启已加载旧动态库的进程，使所有读取方使用同一发布版本：

```sql
ALTER EXTENSION pg_sift UPDATE TO '3.4';
SELECT * FROM sift.upgrade_commands();
```

除使用 `analyzer = keyword` 的索引外，必须重建所有在 3.3 之前创建的文本索引：超过 255 字节的词的处理方式已改变，未重建的索引可能漏掉匹配行。升级过程会打印命令，并发索引重建应在事务块外逐条执行。建议的列索引重建可能改善性能，而建议的列式表重写可能持有排他锁；执行前需逐项评估。

3.3 默认的 `sift.col_format_minor = 15` 无法被 3.2 动态库读取。在写入新格式前，应升级所有读取方，包括备库。如果必须保留 3.2 兼容性，应遵循对应版本的降级指南，并在升级前设置服务器级格式上限；会话级限制无法影响后台工作进程。可选的点查 TID 目录需要格式 16。数据文件不会自动降级。

3.3 的 SQL 迁移修复了旧版 DROP 事件触发器权限缺陷，之前的手动修正不再是单独的安装步骤。它也处理新后端首条带绑定参数的查询，扩展了 3.2.3 对普通查询的修复。

### 3.4 升级与规划变化

发行版 3.4.2 使用 SQL 扩展版本 3.4；安装匹配文件后执行所提供的 3.3 到 3.4 SQL 升级。新版本加入外键辅助的列扫描，调整 IN 列表与扫描启动成本；不能证明等价或成本更高的查询仍使用 PostgreSQL 普通计划。修复涵盖聚合路径排序元数据、达到有符号 128 位上限的 numeric 求和和文本字典处理。已有索引及预加载进程应按该发行版升级指南处理，基准测试耗时不能作为实际负载保证。

### 3.4.2 维护修复

3.4.2 继续使用 SQL 扩展版本 3.4，修复在 `ALTER TABLE` 添加列后进行 `ORDER BY ... LIMIT` 时的崩溃，改进分组聚合及范围/IN 扫描，并修复字典与列维护的边界问题。所有读取方都应先安装匹配文件，再写入更新后的列格式；3.3 到 3.4 的升级边界保持不变。

### 3.5 发行版

3.5.1 继续支持 PostgreSQL 17–19，SQL 升级边界为 `3.5`。3.5 系列增加列格式小版本以及规划器和维护行为变化；写入新格式前，应先在所有读取方安装匹配动态库。3.5.1 修复了无法启动并行 worker 时并行聚合扫描导致的后端崩溃。执行 `ALTER EXTENSION pg_sift UPDATE TO '3.5'` 前请遵循该发行版的升级指南。

### 3.5.2–3.5.3 正确性修复

3.5.2 修复列聚合读取分块词典及多个宽列时的崩溃和错误分组。3.5.3 修复通用计划中 ANY/ALL 测试使用 NULL 或紧凑文本数组参数时的崩溃或错误结果。应在所有读取节点安装 3.5.3，并重启服务，或重新连接所有已加载旧库的连接池与工作进程。SQL 仍为 3.5，这两个补丁版本无需额外 SQL 更新或 REINDEX。替换二进制前，`sift.enable_agg = off` 可使受影响语句恢复使用 PostgreSQL 自身计划。此前 3.5 的格式与升级边界仍然适用。

### 4.0.3 升级与变化跟踪

先运行对应版本的升级预检查，为每个读取方部署匹配库，并重启或替换加载旧版库的进程。遵循所提供的 SQL 链，经过 3.5、3.7、4.0 升级到 4.0.3。即使从 4.0.2 到 4.0.3，也需要更新 `sift.index_choices` 和 `sift.help` 的 SQL 定义；仅该补丁不改变格式，无需 REINDEX。更早版本的文本索引修复及列格式读取限制仍适用。

新建或重建的列索引默认使用 `track_changes = 'auto'`，通过 UPDATE 与 DELETE 两个语句级触发器跟踪符合条件的单列主键。这会增加写入工作和一次基线堆表扫描。仅执行 ALTER EXTENSION 不会为已有索引启用此功能，`track_changes = 'off'` 可关闭它。分区、继承或并发构建等情况下跟踪不完整时，查询回退到堆表检查；修改表拓扑前应阅读升级指南。文本索引的 `merge_fanin` 默认值现为 4，合并会更频繁。

### 4.1 规划与维护

4.1.0 扩展了连接、EXISTS/NOT EXISTS、IN/NOT IN、标量子查询与文本键 top-N 的可用计划，并改进超出内存预算的分组扫描。`sift.agg_probe_rewrite` 和 `sift.agg_probe_direct` 控制相应探测优化；不满足语义或成本条件的查询保留普通计划。

新建或重建的 `key_cut` 索引默认按键范围构建，由 `sift.col_build_key_ranges` 控制。自动合并的每轮工作受 `sift.col_auto_merge_pass_mb` 限制，默认 1024 MB；设为 0 时一次完成，不代表关闭合并。

```sql
ALTER EXTENSION pg_sift UPDATE TO '4.1';
```

执行前须安装匹配文件并重启所有读取节点，包括备库。从 4.0.3 升至 4.1 只推进 SQL 版本，不改 SQL 对象或磁盘格式，也不新增 REINDEX 要求。更早版本受损索引的重建要求仍适用；重建已有 key_cut 索引是采用新构建方式的可选操作。

### 4.1.1–4.1.4 的正确性与维护修复

4.1.4 修复了受影响的 `NOT NULL` 列上 JSONB 包含条件取反时的错误结果，以及连接或相关查询重新扫描 `sift_col` 索引时重复返回写缓冲行的问题。升级前可用 `sift.enable_agg = off` 避开前者，用 `sift.enable_col_scan = off` 避开后者；仅关闭聚合不能解决重复行问题。

4.1.1 阻止新的跟踪行计数错误；4.1.2 还会对旧跟踪日志重新检查堆表并修复过期状态。4.1.3 修复并发重建索引时的表锁与写入事务死锁。应在所有读取节点安装 4.1.4 文件并重启服务器，包括备库。SQL 仍为 4.1：已有 4.1 安装无需再次执行 `ALTER EXTENSION`，这些补丁也不要求转换格式或强制 `REINDEX`。

在 `pg_upgrade`、禁用触发器的恢复或并发索引重建之后，跟踪列索引可能仍返回正确结果，但使用堆表检查。可执行普通 `REINDEX INDEX`，或为每个索引重新设置原有 `track_changes` 值，恢复优化路径。更早版本的升级要求及受损索引处理要求仍然适用。

### 4.2–4.3.2 的升级边界

先在所有读取节点（包括备库）安装匹配动态库并重启服务器，再更新每个数据库：

```sql
ALTER EXTENSION pg_sift UPDATE TO '4.3.2';
```

4.3.2 修复了聚合扫描将相关派生表的结果复用于后续外层行、导致计数或求和错误的问题。升级前可用 `sift.enable_agg = off` 避开该扫描。4.3.1 还修复了跟踪 UPDATE/DELETE 语句期间的读取，并为 `sift.table_option_values` 设置复制标识，使发布删除操作的全表 publication 不再阻止关系删除。该目录修复需要执行 SQL 更新，仅替换动态库并不足够。

4.2/4.3 系列改变了索引格式。4.3 新建或合并的列段可能使用旧库无法读取的格式小版本 44。开始写入前必须升级所有读取方；若需回退，应先遵循对应版本的格式上限流程。已有索引在合并或重建时采用新布局。4.3.1 至 4.3.2 的补丁本身不改变格式或对象，也不新增 REINDEX 要求。

运行上游升级预检查，并移除废弃设置 `sift.enable_col` 和 `sift.table_auto_defrag_dead_pct`。`sift.eval_methods` 只保留 `penalty_maintenance`，应移除已废弃的方法名。关系选项 `rollup` 和 `rollup_dims` 更名为 `rollup_leaves` 和 `rollup_columns`，4.x 仍接受旧拼写。更早版本的升级要求与受损索引重建要求继续适用。
