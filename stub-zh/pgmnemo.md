## 用法

来源：

- [pgmnemo v0.20.0 README](https://github.com/pgmnemo/pgmnemo/blob/v0.20.0/README.md)
- [pgmnemo v0.20.0 发行说明](https://github.com/pgmnemo/pgmnemo/releases/tag/v0.20.0)
- [pgmnemo v0.20.0 使用指南](https://github.com/pgmnemo/pgmnemo/blob/v0.20.0/docs/USAGE.md)
- [pgmnemo v0.20.0 SQL 参考](https://github.com/pgmnemo/pgmnemo/blob/v0.20.0/docs/SQL_REFERENCE.md)
- [pgmnemo v0.20.0 变更日志](https://github.com/pgmnemo/pgmnemo/blob/v0.20.0/CHANGELOG.md)
- [pgmnemo v0.20.0 控制文件](https://github.com/pgmnemo/pgmnemo/blob/v0.20.0/extension/pgmnemo.control)
- [v0.19.1 到 v0.20.0 升级 SQL](https://github.com/pgmnemo/pgmnemo/blob/v0.20.0/extension/pgmnemo--0.19.1--0.20.0.sql)

pgmnemo 将智能体记忆存储在 PostgreSQL 中，并通过向量、BM25 风格文本、图、元数据、时间、来源和结果置信度等信号进行检索。它安装在 pgmnemo 模式中，依赖 vector 扩展，当前 SQL API 要求使用 1024 维嵌入。

版本 0.20.0 保留了语料库维护、情境与实体召回接口，并为 `recall_hybrid()` 新增两个可选候选集扩展器：因果边广度优先扩展和实体键 GIN 扩展。两者默认均关闭，返回结果现在会标明每个候选项的检索路径。

### 安装

    CREATE EXTENSION IF NOT EXISTS vector;
    CREATE EXTENSION IF NOT EXISTS pgmnemo CASCADE;

    SELECT pgmnemo.version();
    SELECT * FROM pgmnemo.stats();

v0.20.0 控制文件将 pgmnemo 标记为受信任，将其安装到模式 `pgmnemo`，要求 `vector`，且不可重定位。

### 写入一条经验

    SELECT pgmnemo.ingest(
      p_role        := 'developer',
      p_project_id  := 1,
      p_topic       := 'security',
      p_lesson_text := 'Rotate signing keys after a compromise.',
      p_importance  := 4,
      p_embedding   := NULL,
      p_commit_sha  := 'abc1234',
      p_metadata    := '{"source":"incident-runbook"}'::jsonb
    );

当 pgmnemo.gate_strict 为 enforce 时，必须提供 commit_sha 或 artifact_hash 来源信息。warn 允许写入未经验证的记录，但会产生审计警告；off 则禁用该门控。

### 带置信度过滤的召回

混合召回会结合嵌入与文本信号：

    SELECT lesson_id, topic, score, match_confidence, retrieval_source
    FROM pgmnemo.recall_hybrid(
      '<1024-dimensional vector literal>'::vector(1024),
      'JWT rotation key compromise',
      10,
      'developer',
      1,
      0.4,
      0.4,
      60,
      'dag-2026-abc',
      ARRAY['note', 'fact'],
      0.40
    );

0.13.0 新增的最后一个 p_min_score 参数会在应用 LIMIT 之前，剔除 match_confidence 低于阈值的候选项。传入 NULL 可保留 0.13 之前的行为。发行说明建议将 0.40 作为起点，而非通用值；应针对嵌入模型与反馈质量进行校准。

recall_fast、recall_lessons 和池化召回入口同样支持 p_min_score 概念。当同时提供文本与嵌入，且 pgmnemo.disable_hybrid 为 off 时，recall_lessons 会路由到混合召回。

### 记录结果

    SELECT pgmnemo.reinforce(1001, 'success', true);
    SELECT pgmnemo.reinforce(
      ARRAY[1001, 1002]::bigint[],
      'failure',
      false
    );

第三个 p_used 参数记录召回的记忆是否实际被采用。true 或 NULL 会增加 use_count；false 会记录结果，但不计入使用次数。建议显式传值，让分析能够区分被忽略的建议与实际采用的建议。

在默认 posterior 模式下，匹配置信度为：

    (success_count + alpha)
    / (success_count + failure_count + alpha + beta)

默认 Beta 先验为 alpha 1 和 beta 1。只有在有充分理由采用其他先验时，才应将 pgmnemo.confidence_prior_alpha 和 pgmnemo.confidence_prior_beta 设置为 0.01 到 100 之间的值。

### 类型化记忆与导航

重要的写入辅助函数包括 remember_fact、remember_event、remember_relation、add_edge、reembed 和 recompute_content。remember_fact 会取代同一实体/属性对的当前有效事实；事件保持追加式；关系也会填充图接口。

使用 navigate_locate 或 navigate_locate_dispatch 在字符预算内选择候选 ID，再用 navigate_expand_typed 获取内容及相邻图边。

### 情境与实体召回

0.15 系列新增了确定性的情境指纹与专用召回路径：

```sql
SELECT pgmnemo.extract_sit_fp(
  'security',
  'failure_class=KEY_ROTATION outcome=COMPLETED'
);
SELECT *
FROM pgmnemo.recall_situation(
  pgmnemo.extract_sit_fp(
    'security',
    'failure_class=KEY_ROTATION outcome=COMPLETED'
  ),
  1,
  'developer',
  10
);
```

从 0.15.1 开始，`recall_situation` 默认返回已经验证的记忆。仅当调用方明确接受没有来源验证的记忆时，才设置 `pgmnemo.include_unverified = on`。

0.16 系列还会在写入期间将稳定的实体键提取到 `metadata.entity_keys`，并提供以实体为中心的召回：

```sql
SELECT pgmnemo.extract_entity_keys('The run failed with INFRA_FAILURE.');
SELECT * FROM pgmnemo.recall_entity('failure:INFRA_FAILURE', 10);
```

这些提取器是确定性分类器，而不是语义实体解析。应规范化应用词汇并检查生成的键，再决定是否将它们用于租户隔离或授权判断。

### 可选的图与实体候选集扩展

版本 0.20.0 可以引入原 ANN 与 BM25 候选集之外的条目。图扩展从排名靠前的 ANN 锚点出发，只沿 `causal` 边遍历，并受深度上限和每节点枢纽上限限制。实体扩展使用 `metadata.entity_keys` 中已建 GIN 索引的键。两个主权重的默认值均为 `0.0`，因此升级不会启用任何扩展路径。

评估新路径时，建议使用事务局部设置：

```sql
BEGIN;
SET LOCAL pgmnemo.graph_expand_weight = '0.15';
SET LOCAL pgmnemo.graph_expand_depth = '1';
SET LOCAL pgmnemo.graph_expand_ann_k = '15';
SET LOCAL pgmnemo.graph_expand_per_node = '10';
SET LOCAL pgmnemo.graph_entity_expand_weight = '0.10';
SET LOCAL pgmnemo.graph_entity_min_overlap = '1';
SET LOCAL pgmnemo.graph_entity_max_expansion = '50';

SELECT lesson_id, score, retrieval_source
FROM pgmnemo.recall_hybrid(
  query_embedding := '<1024-dimensional vector literal>'::vector(1024),
  query_text := 'JWT rotation key compromise',
  k := 10
);
ROLLBACK;
```

`retrieval_source` 是第 18 个输出列，取值为 `ann`、`graph` 或 `entity`。图深度限于 1 或 2，ANN 锚点超采样范围为 10-50，每节点枢纽上限范围为 3-50。实体扩展至少要求一个重叠键，并使用配置的每键候选项上限。应将这些权重视为与工作负载相关的排序控制：v0.20.0 发行版刻意不对这些可选路径作通用召回率或延迟保证。

### 配置索引

- pgmnemo.confidence_mode：默认为 posterior；additive 保留旧版计算方式。
- pgmnemo.confidence_prior_alpha 和 pgmnemo.confidence_prior_beta：贝叶斯先验参数。
- pgmnemo.confidence_boost_weight：置信度对排名的贡献；默认值为 0，因此除非启用，否则置信度不会改变排名。
- pgmnemo.gate_strict 和 pgmnemo.include_unverified：来源强制要求与检索控制。
- pgmnemo.disable_hybrid 和 pgmnemo.ef_search：召回策略与 HNSW 搜索宽度。
- pgmnemo.track_recall_recency：召回是否更新 last_recalled_at 和 recall_count。
- pgmnemo.max_query_text_chars、pgmnemo.tenant_id 和 pgmnemo.test_project_floor：文本、租户和可选测试项目控制。
- pgmnemo.graph_expand_weight、pgmnemo.graph_expand_depth、pgmnemo.graph_expand_ann_k 和 pgmnemo.graph_expand_per_node：因果边候选集扩展、深度、ANN 锚点与每节点枢纽上限。
- pgmnemo.graph_entity_expand_weight、pgmnemo.graph_entity_min_overlap 和 pgmnemo.graph_entity_max_expansion：实体键 GIN 扩展、最小重叠数与每键候选项上限。

旧版 confidence-delta 设置已弃用，在 posterior 模式下会被忽略。

### 升级到 0.20.0

从 0.19.1 升级时，使用软件包提供的扩展更新路径：

```sql
ALTER EXTENSION pgmnemo UPDATE TO '0.20.0';

SELECT extversion
FROM pg_extension
WHERE extname = 'pgmnemo';
```

升级脚本会删除并重建 11 参数的 `recall_hybrid()`，因为它增加 `retrieval_source` 后，返回结构从 17 列变为 18 列。它还会删除并重建 `stats()`，后者增加七个图扩展设置，返回结构从 19 列增至 26 列。升级前应检查位置式行映射、wrapper、prepared consumer 和精确列数断言。按名选取稳定列的调用方不受影响。

### 注意事项

- pgmnemo 0.20.0 应使用 PostgreSQL 17 或 18。带标签的变更日志指出，0.10 系列引入的语法使旧有的 PostgreSQL 14-16 兼容性声明不再准确；当前 Pigsty 软件包面向 17-18。

语料库维护操作默认为只读：

```sql
SELECT * FROM pgmnemo.reclassify_corpus();
SELECT * FROM pgmnemo.consolidate(
  p_threshold := 0.92,
  p_dry_run := true,
  p_role := NULL,
  p_limit := 100
);
SELECT * FROM pgmnemo.undo_consolidate(
  p_canonical_id := 42,
  p_dry_run := true
);
```

只有在事务内审阅结果之后，才能设置 `p_dry_run := false`。在 0.14.2 中，重新分类只处理类型为空或由分类器拥有的条目，并保留 event、relation 等由整理者拥有的类型。合并会把非规范经验标记为已取代、写入边并累积证据计数；`undo_consolidate` 使用这些边恢复选定的簇。

- 召回可能写入最近访问元数据。进行只读分析时应禁用 pgmnemo.track_recall_recency。
- 置信度模型的可靠性取决于强化反馈质量。未经评估，不应将 posterior 值视为经过校准的概率。
- HNSW、文本、图和元数据索引会增加写入与维护成本。
- confidence_boost_weight 的默认值为 0，这意味着 p_min_score 可以过滤结果，而置信度依然完全不参与排名。
- 分类采用确定性的关键词与正则表达式启发式规则，而不是语义审查。在应用语料库变更之前，务必检查试运行分布和建议的重复簇。
