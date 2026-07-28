## 用法

来源：

- [官方上游 README](https://github.com/joshmeek/otlet/blob/5bac507a37862fb305c3a9d5a731d90f146ce946/README.md)
- [官方扩展控制文件 (otlet.control)](https://github.com/joshmeek/otlet/blob/5bac507a37862fb305c3a9d5a731d90f146ce946/crates/otlet_pg/otlet.control)
- [官方实现源代码](https://github.com/joshmeek/otlet/blob/5bac507a37862fb305c3a9d5a731d90f146ce946/crates/otlet_pg/src/lib.rs)

`otlet` — 在 PostgreSQL 中使用后台工作者进行本地 LLM 推理，通过队列任务、收据和语义数据操作。将其用于相应的向量、模型或检索工作流。在目标 PostgreSQL 构建中使用上述链接的上游固定版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION otlet;

SELECT output
FROM otlet.ask(
  'qwen35_4b',
  'Summarize these customer notes in one sentence.',
  (SELECT jsonb_agg(to_jsonb(n))
   FROM customer_notes n WHERE customer = 'Riverline Labs'),
  '{"type":"object","required":["summary"],"additionalProperties":false,"properties":{"summary":{"type":"string"}}}'
);
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
