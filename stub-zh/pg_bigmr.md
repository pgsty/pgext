## 用法

来源：

- [Official extension control file](https://github.com/shinyaaa/pg_bigmr/blob/cb85c4eab90a8a41aa81492751713b683e20e084/pg_bigmr.control)
- [Official upstream documentation](https://github.com/shinyaaa/pg_bigmr/blob/cb85c4eab90a8a41aa81492751713b683e20e084/README.md)
- [Official SQL definition](https://github.com/shinyaaa/pg_bigmr/blob/cb85c4eab90a8a41aa81492751713b683e20e084/sql/pg_bigmr--0.1.0.sql)
- [Official Rust implementation](https://github.com/shinyaaa/pg_bigmr/blob/cb85c4eab90a8a41aa81492751713b683e20e084/src/lib.rs)

`pg_bigmr` 0.1.0 是 `pg_bigm` 的实验性 Rust 与 pgrx 重实现。它提供基于二元组的文本相似度以及用于加速 `LIKE` 和相似度搜索的 GIN 操作符类，但上游明确标记为开发中项目。

### 核心流程

把模块加入 `shared_preload_libraries`，重启 PostgreSQL，创建扩展，再用 `gin_bigm_ops` 构建 GIN 索引。

```conf
shared_preload_libraries = 'pg_bigmr'
```

```sql
CREATE EXTENSION pg_bigmr;

CREATE INDEX documents_body_bigm_idx
  ON documents USING gin (body gin_bigm_ops);

SELECT * FROM documents WHERE body LIKE '%database%';
SELECT bigm_similarity('database', 'databass');
```

`show_bigm(text)` 显示提取的二元组，`likequery(text)` 把字面值转义成包含式 `LIKE` 模式，`pg_gin_pending_stats(oid)` 报告 GIN 索引待处理列表的页数和元组数。`=%` 操作符使用相似度阈值。

### 参数与注意事项

`pg_bigmr.similarity_limit` 默认 0.3，`pg_bigmr.gin_key_limit` 默认不限制，`pg_bigmr.enable_recheck` 默认开启。改变限制或关闭堆表复核，可能用正确性或选择性换取速度。它并不保证与 `pg_bigm` 即插即用兼容；应针对真实文本做基准测试，验证 Unicode 与通配符行为，且未经工作负载专项验证不要投入生产。
