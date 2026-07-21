## 用法

来源：

- [pgmnemo v0.13.0 README](https://github.com/pgmnemo/pgmnemo/blob/v0.13.0/README.md)
- [pgmnemo v0.13.0 使用指南](https://github.com/pgmnemo/pgmnemo/blob/v0.13.0/docs/USAGE.md)
- [pgmnemo v0.13.0 SQL 参考](https://github.com/pgmnemo/pgmnemo/blob/v0.13.0/docs/SQL_REFERENCE.md)
- [pgmnemo v0.13.0 发行说明](https://github.com/pgmnemo/pgmnemo/releases/tag/v0.13.0)
- [pgmnemo v0.13.0 控制文件](https://github.com/pgmnemo/pgmnemo/blob/v0.13.0/extension/pgmnemo.control)

pgmnemo 将代理内存存储在 PostgreSQL 中，并通过向量、BM25 样式的文本、图、元数据、时间序列、来源和置信度结果信号检索。它安装到 pgmnemo 模式中，需要 vector 扩展，并且当前 SQL API 需要 1024 维嵌入。

v0.13.0 版本默认将置信度转换为贝叶斯后验概率，记录召回的记忆是否实际被使用，并添加了最小置信度过滤器以进行召回。

### 安装

    CREATE EXTENSION IF NOT EXISTS vector;
    CREATE EXTENSION IF NOT EXISTS pgmnemo CASCADE;

    SELECT pgmnemo.version();
    SELECT * FROM pgmnemo.stats();

v0.13.0 控制文件在所需 vector 扩展可用时将 pgmnemo 标记为受信任且非超级用户可安装。

### 载入一课

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

当 pgmnemo.gate_strict 强制执行时，需要 commit_sha 或 artifact_hash 来源。warn 允许带有审计警告的未验证写入；off 禁用此门。

### 带置信度过滤器的召回

混合召回结合嵌入和文本信号：

    SELECT lesson_id, topic, score, match_confidence
    FROM pgmnemo.recall_hybrid(
      '<1024 维向量字面量>'::vector(1024),
      'JWT 旋转密钥泄露',
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

最后的 p_min_score 参数在应用 LIMIT 之前移除匹配置信度低于阈值的候选者。NULL 保留 v0.13.0 之前的默认行为。发行说明建议将 0.40 作为起点，而不是通用值；根据嵌入模型和反馈质量进行校准。

此相同的 p_min_score 概念在 recall_fast、recall_lessons 和聚合召回入口点中可用。recall_lessons 在同时提供文本和嵌入且 pgmnemo.disable_hybrid 关闭时路由到混合召回。

### 记录结果

    SELECT pgmnemo.reinforce(1001, 'success', true);
    SELECT pgmnemo.reinforce(
      ARRAY[1001, 1002]::bigint[],
      'failure',
      false
    );

第三个 p_used 参数记录召回的记忆是否实际被使用。true 或 NULL 增加 use_count；false 记录结果而不计入使用次数。请明确指定以使分析能够区分忽略的建议和已使用的建议。

在默认后验模式下，匹配置信度为：

    (success_count + alpha)
    / (success_count + failure_count + alpha + beta)

默认 Beta 先验参数为 alpha 1 和 beta 1。当需要其他先验时，在 0.01 到 100 之间设置 pgmnemo.confidence_prior_alpha 和 pgmnemo.confidence_prior_beta。

### 类型化记忆和导航

重要的写入辅助函数包括 remember_fact、remember_event、remember_relation、add_edge、reembed 和 recompute_content。remember_fact 覆盖实体/属性对的活跃事实；事件保持追加导向；关系也填充图表面。

使用 navigate_locate 或 navigate_locate_dispatch 选择字符预算内的候选 ID，然后使用 navigate_expand_typed 获取内容和邻近的图边。

### 配置索引

- pgmnemo.confidence_mode：默认为后验概率；additive 保留了遗留计算。
- pgmnemo.confidence_prior_alpha 和 pgmnemo.confidence_prior_beta：贝叶斯先验参数。
- pgmnemo.confidence_boost_weight：置信度对排名的贡献；默认值为 0，因此除非启用否则置信度不会影响排名。
- pgmnemo.gate_strict 和 pgmnemo.include_unverified：来源验证和检索。
- pgmnemo.disable_hybrid 和 pgmnemo.ef_search：召回策略和 HNSW 搜索宽度。
- pgmnemo.track_recall_recency：是否更新 recall_last_recalled_at 和 recall_count。
- pgmnemo.max_query_text_chars、pgmnemo.tenant_id 和 pgmnemo.test_project_floor：文本、租户控制以及可选测试项目控制。

较旧的置信度差值设置在后验模式下已弃用并被忽略。

### 注意事项

- 召回可以写入最近性元数据。禁用 pgmnemo.track_recall_recency 以进行只读分析。
- 置信度模型仅在其强化反馈评估后才可靠。避免将后验值视为经过校准的概率。
- HNSW、文本、图和元数据索引增加了写入和维护成本。
- 默认置信度增强权重为 0，这意味着 p_min_score 可以过滤结果而置信度对排名没有任何贡献。
