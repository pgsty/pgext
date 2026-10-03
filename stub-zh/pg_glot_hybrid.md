## 用法

来源：

- [extensions/pg_glot_hybrid/pg_glot_hybrid.control](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/extensions/pg_glot_hybrid/pg_glot_hybrid.control)
- [README.md](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/README.md)
- [Cargo.toml](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/Cargo.toml)
- [extensions/pg_glot_hybrid/Cargo.toml](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/extensions/pg_glot_hybrid/Cargo.toml)
- [extensions/pg_glot_hybrid/src/lib.rs](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/extensions/pg_glot_hybrid/src/lib.rs)
- [extensions/pg_glot_hybrid/src/customscan.rs](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/extensions/pg_glot_hybrid/src/customscan.rs)

`pg_glot_hybrid` 将 `pg_glot` 的中日韩配置、`pg_textsearch` 的 BM25 与 `vector` 的稠密向量结合起来，提供 `glot.rank` 规划器集成和显式返回集合的 `glot.hybrid` 接口。

### 核心用法

```conf
shared_preload_libraries = 'pg_textsearch,pg_glot_hybrid'
```

```sql
CREATE EXTENSION pg_glot_hybrid CASCADE;
CREATE TABLE glot_documents (id bigint PRIMARY KEY, body text, emb vector(3));
CREATE INDEX ON glot_documents USING bm25(body) WITH (text_config = 'public.korean');
CREATE INDEX ON glot_documents USING hnsw(emb vector_cosine_ops);
SELECT id, score FROM glot.hybrid('glot_documents', 'id', 'body', 'emb',
  '형태소 분석', '[1,0,0]'::vector, 60, 60, 10);
```

### 运行边界

由超级用户安装。要使用预期的规划器路径，需要预加载下面两个库并重启。自定义扫描要求查询字面值、普通 ORDER BY/LIMIT 和真实列引用。缺少钩子时，`glot.rank` 会回退到另一种非 RRF 分数。显式 `glot.hybrid` 接口不依赖该钩子，但要求匹配的 BM25 索引、向量索引和 bigint 键。嵌入由应用提供。所核对默认目标为 PostgreSQL 17，各依赖扩展仍有自己的版本要求。
