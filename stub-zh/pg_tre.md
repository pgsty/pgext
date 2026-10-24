## 用法

来源：

- [README.md](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.2/README.md)
- [pg_tre.control](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.2/pg_tre.control)
- [sql/pg_tre--4.5.2.sql](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.2/sql/pg_tre--4.5.2.sql)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.2/CHANGELOG.md)
- [LIMITATIONS.md](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.2/LIMITATIONS.md)
- [META.json](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.2/META.json)
- [.github/workflows/ci.yml](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.2/.github/workflows/ci.yml)
- [库初始化](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.2/src/module.c)

`pg_tre` 4.5.2 提供 `tre` 原生索引访问方法，对文本执行近似正则表达式匹配。上游支持 PostgreSQL 17–19；Pigsty 的 4.5.2 软件包面向 PostgreSQL 17–18。它利用码点三元组与倒排表缩小候选范围，再到堆表核对匹配。适合需要显式编辑代价的容错标识符、日志消息与词法模式检索。

### 启用与查询

索引写入使用 PostgreSQL 的通用 WAL 机制。从 4.0.0 起，执行 `CREATE EXTENSION` 即可启用，新安装无需设置 `shared_preload_libraries` 或重启。创建扩展要求超级用户，控制文件不标记为可信，也不允许重定位。

```sql
CREATE EXTENSION pg_tre;
CREATE TABLE documents (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  body text NOT NULL
);
INSERT INTO documents (body) VALUES
  ('The PostgreSQL database system'),
  ('The Postgres databse system');
CREATE INDEX documents_body_tre ON documents USING tre (body);
SELECT id, body, body <@> tre_pattern('database', 1) AS distance
FROM documents
WHERE body %~~ tre_pattern('database', 1)
ORDER BY distance ASC NULLS LAST, id;
```

PostgreSQL 19 使用内核正则字符分类器；PostgreSQL 17–18 的 ICU 字符类边界行为可能与原生 PostgreSQL 正则语义不同，具体见上游说明。

### 模式与对象

- `tre_pattern(text)` 使用 `pg_tre.default_max_cost`；`tre_pattern(text, int4)` 显式指定最大代价；五参数形式还可指定插入、删除与替换代价。模式采用 POSIX 扩展正则表达式和 TRE 近似匹配预算。
- `text %~~ tre_pattern` 是可索引的近似匹配谓词；`text <@> tre_pattern` 返回编辑代价，预算内无匹配时返回 NULL。
- `tre_distance`、`tre_similarity`、`tre_amatch`、`tre_amatch_cost` 与 `tre_amatch_detail` 提供匹配及兼容函数。
- 索引操作符族也支持 LIKE、ILIKE、正则与等值谓词，并进行堆表复核。独立命名的三元组函数可与 `pg_trgm` 共存；`pg_tre` 不再安装 3.0 中删除的冲突三元组操作符。

索引通过 `tre_text_ops` 接受一个文本列，支持表达式和部分索引，不支持仅索引扫描。短模式、较大编辑预算与缺少字面的正则可能产生大量候选；应检查实际执行计划，并保留取消机制和查询时间限制。

### 维护与限制

`fastupdate` 必须为 true；关闭它现在会被拒绝，因为直接写入路径无法保持索引格式。已有关闭设置的索引仍可打开，但插入会走待合并列表。`pending_list_limit` 单位为 KiB，默认 0 表示不限制。正值会在越过阈值时触发同步合并，由当前插入后端承担工作；已有合并正在执行时，列表可能暂时超过该值。

`pg_tre.min_trigram_freq` 必须保持为 1。以前大于 1 的值会省略稀有倒排条目，可能漏掉结果；现在这类设置会被拒绝。不可信模式应保留有限的正则状态数、编译和匹配时间。索引构建使用 tuplesort 和临时磁盘，内存由 maintenance_work_mem 控制；安排大型构建前应先估算规模。

```sql
SELECT * FROM tre_estimate_index_build('documents'::regclass, 2);
SELECT * FROM pg_tre_index_pending('documents_body_tre'::regclass);
SELECT * FROM pg_tre_index_check('documents_body_tre'::regclass, 2);
```

`pg_tre_index_check` 比较索引与顺序扫描结果。每个样本会执行多次顺序扫描，大表上须限制样本数量。4.5.1 修复了它的搜索路径处理，固定可信查找路径并限定调用名称；以高权限用户运行此辅助函数前应使用修复版本。

### 升级边界

安装匹配文件，确保会话使用新库，再逐库更新：

```sql
ALTER EXTENSION pg_tre UPDATE TO '4.5.2';
```

从 3.x 跨至 4.x 时，索引页格式虽然不变，WAL 记录类型却不兼容。替换库前必须干净停库；物理备库不能与主库运行不同主版本的扩展，须在主库升级后从新的基础备份重建，此边界不支持滚动升级。还须移除旧的 `wal_consistency_checking = 'pg_tre'` 设置，否则 4.x 无法启动；如需一致性检查，上游建议使用 `wal_consistency_checking = 'all'`。

4.4 写入格式 11，并可读取格式 6–10。从这些格式升级无须例行重建；可选的 REINDEX 用于获得新的密集倒排表示。经过 4.4+ 写入后降级到 4.3，需要按旧格式重建受影响索引。更早版本的升级可能跨越其他边界，应检查更新日志。

所有以大于 1 的 min_trigram_freq 构建的索引都必须 REINDEX；安装修复无法补回已省略的倒排条目。4.5 还会在此前受影响版本的索引上关闭过期的 SuRF 加速，立即恢复正确性；如需重新获得该加速，可选择重建。4.5.0 升至 4.5.1 不改变格式或普通查询行为，无须额外 REINDEX。含已删除元组 bloom 选项的旧索引定义，重建前仍须修正。

### 4.5.2 正则表达式修复

4.5.2 修复换行模式 `(?n)` 中的锚点和否定字符类，不改变索引格式；从 4.5.1 升级不需要额外 REINDEX。从更早的受影响版本升级时，仍须遵守上面的 WAL 迁移及遗漏倒排项重建要求。

```sql
SELECT E'a\nfoo' %~~ tre_pattern('(?n)^foo', 0) AS line_start,
       E'foo\n' %~~ tre_pattern('(?n)foo$', 0) AS line_end,
       E'\n' %~~ tre_pattern('(?n)[^x]', 0) AS negated_class;
```

预期结果依次为 true、true、false。修复同时作用于匹配函数和索引复核；未启用换行模式的模式保持原有行为。
