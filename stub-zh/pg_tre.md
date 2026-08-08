## 用法

来源：

- [pg_tre v3.0.2 README](https://codeberg.org/gregburd/pg_tre/src/tag/v3.0.2/README.md)
- [pg_tre v3.0.2 控制文件](https://codeberg.org/gregburd/pg_tre/src/tag/v3.0.2/pg_tre.control)
- [pg_tre v3.0.2 扩展 SQL](https://codeberg.org/gregburd/pg_tre/src/tag/v3.0.2/sql/pg_tre--3.0.2.sql)
- [pg_tre v3.0.2 变更日志](https://codeberg.org/gregburd/pg_tre/src/tag/v3.0.2/CHANGELOG.md)
- [生产容量规划与限制](https://codeberg.org/gregburd/pg_tre/src/tag/v3.0.2/LIMITATIONS.md)
- [pg_tre v3.0.2 PGXN 元数据](https://codeberg.org/gregburd/pg_tre/src/tag/v3.0.2/META.json)
- [PG17 与 PG18 CI 矩阵](https://codeberg.org/gregburd/pg_tre/src/tag/v3.0.2/.github/workflows/ci.yml)
- [官方回归测试](https://codeberg.org/gregburd/pg_tre/src/tag/v3.0.2/test/sql)

`pg_tre` 3.0.2 提供 `tre` 原生 PostgreSQL 索引访问方法，用于在 `text` 上执行近似正则表达式匹配。它通过码点三元组和压缩 posting 缩小候选范围，再由 TRE 正则引擎对堆元组做权威复查。适用于需要真实插入、删除、替换成本的容错标识符、日志消息、错误码或其他词法模式；它不能替代语言学全文检索或语义向量检索。

### 启用扩展

索引写入路径使用自定义 WAL 资源管理器，因此创建或修改 `tre` 索引前，必须预加载该库并完整重启 PostgreSQL；仅重载配置并不足够。

```ini
shared_preload_libraries = 'pg_tre'
```

随后在每个需要它的数据库中安装扩展。v3.0.2 控制文件设置了 `superuser=true`、`trusted=false` 和 `relocatable=false`，因此创建扩展需要超级用户，且扩展不可迁移模式。

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

SELECT id, body,
       body <@> tre_pattern('database', 1) AS distance
FROM documents
WHERE body %~~ tre_pattern('database', 1)
ORDER BY distance ASC NULLS LAST, id;
```

未预加载时按需载入，`tre_amatch` 和 `tre_version` 等兼容函数仍可使用，但索引变更会因自定义资源管理器未注册而失败。应把 `pg_tre` 加入现有的逗号分隔预加载列表，而不是覆盖其他库，然后重启服务器。

带标签的 v3.0.2 构建元数据要求 PostgreSQL 17 或更高版本，Nix flake 分别导出 `pg17` 和 `pg18` 构建，回归 CI 矩阵也在 PostgreSQL 17 与 18 上运行。因此，这两个大版本构成该发行版有证据支持的范围。README 与用户指南标题仍写 PostgreSQL 18+，相对于带标签的构建和测试配置已经过时。

### 模式、操作符与函数

`tre_pattern(text)` 使用 `pg_tre.default_max_cost` 创建模式；`tre_pattern(text, int4)` 显式指定最大成本；`tre_pattern(text, int4, int4, int4, int4)` 还可分别设置插入、删除和替换成本。模式语言为 POSIX 扩展正则表达式，并加入 TRE 的局部近似预算语法，例如 `{~1}`。

- `text %~~ tre_pattern` 返回近似正则匹配布尔值，是主要的可索引谓词。
- `text <@> tre_pattern` 返回最佳编辑成本；若模式预算内不存在可行对齐则返回 `NULL`。`tre` 操作符族支持按升序直接从索引为 top-N 查询排序；合适时仍应加入选择性 `%~~` 谓词，并使用确定性的并列排序键。
- `tre_distance` 和 `tre_similarity` 以函数形式提供距离与归一化相似度。`tre_amatch`、`tre_amatch_cost` 和 `tre_amatch_detail` 是直接调用 TRE、不会使用索引的兼容函数。
- `LIKE`、`ILIKE`、`~`、`~*` 和 `=` 在 v3.0.2 操作符族中绑定为有损候选过滤器；PostgreSQL 始终会在堆上复查其原生语义。
- `tre_trgm_similarity`、`tre_trgm_distance`、`tre_word_similarity` 和 `tre_strict_word_similarity` 以不同名称提供类似 pg_trgm 的计算。自 3.0.0 起，`pg_tre` 刻意不再创建裸 `%`、`<->`、`<%`、`<<->`、`<<%` 或 `<<<->` 操作符，从而能够与 `pg_trgm` 共存。

`tre` 索引仅支持单列，接受 `text` 操作符类 `tre_text_ops`，并且是有损索引：它不提供仅索引扫描，每一条最终结果都由 TRE 复查决定。表达式索引和部分索引仍可通过 PostgreSQL 的常规机制使用。

### 选择性、成本与安全限制

索引加速依赖可用的字面量三元组。过短模式、以非字面量正则结构为主的模式、较大的编辑预算，或超过 `pg_tre.max_extraction_fanout=4096` 的提取扇出，都可能退化为全有损候选位图。正确性不受影响，但规划器可能合理地选择顺序扫描。应在代表性数据上使用 `EXPLAIN (ANALYZE, BUFFERS)`；生产建议使用 0-2 的编辑预算，3 必须结合负载测试，普通文本通常不应使用大于 3 的值，因为 TRE 复查成本会快速上升。

重要安全默认值包括 `pg_tre.max_nfa_states=10000`、`pg_tre.compile_timeout_ms=1000` 和 `pg_tre.match_timeout_ms=1000`。处理不可信模式时应保持这些边界。版本 3.0.1 在 `%~~` 候选构建与复查路径中加入了协作式中断检查，因此 v3.0.2 的宽范围扫描可以被普通 `statement_timeout` 或客户端取消终止。

支持的逐索引存储参数为 `fastupdate`、`pending_list_limit`、`range_size_blocks` 和 `q`；`q` 必须保持为 3。默认值为 `fastupdate=true`、`pending_list_limit=4096` KiB 和 `range_size_blocks=128`。启用快速更新时，插入先进入 pending list，后续维护再将其合并；对于写入密集的索引，应监控并执行 vacuum，避免 pending list 无限增长。

版本 3.0.0 删除了逐元组位置 bloom 载荷，以及 `tuple_bloom_enable`、`bloom_tuple_bits` GUC 和索引选项。仍指定其中任一选项的旧索引定义会失败，必须更新。range bloom 仍会构建，但 v3.0.2 变更日志明确记录扫描尚未探测它；因此，当前过滤应理解为基于 posting 的候选缩减加上强制 TRE 复查，而不是旧 README 文案仍展示的三个活跃层级。

### 构建、维护与升级边界

索引构建使用 PostgreSQL tuplesort，内存主要受 `maintenance_work_mem` 约束；临时磁盘消耗仍随发出的三元组元组数增长。大型构建应先估算，并在在线表上优先使用并发操作：

```sql
SELECT *
FROM tre_estimate_index_build('documents'::regclass, 2);

CREATE INDEX CONCURRENTLY documents_body_tre_live
ON documents USING tre (body);

REINDEX INDEX CONCURRENTLY documents_body_tre_live;
```

`tre_estimate_index_build` 最多采样 2000 行，报告估算行数、三元组元组数、临时磁盘和最终索引大小。`pg_tre.build_max_entries_mb=0` 表示不启用临时磁盘保护；临时表空间受限时，应根据测量设置非零上限。`pg_tre.min_trigram_freq=1` 默认保留所有 posting；调高它可以缩小索引，但会增加有损回退工作，因此应针对顺序扫描基线验证查询计划和等价召回正确性。

版本 3.0.2 是打包与文档发行版：既不改变 C 行为、SQL 接口、WAL，也不改变磁盘格式。从 3.0.1 升级只需执行版本更新 `ALTER EXTENSION pg_tre UPDATE TO '3.0.2'`，无需 `REINDEX`。跨越更早版本升级时，应先阅读所有中间版本说明；1.6.0 之前的格式边界可能要求重建索引，而 3.0.0 的 v9 格式仍可向后读取 v6-v8 索引。
