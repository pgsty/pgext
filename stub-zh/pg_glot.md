## 用法

来源：

- [extensions/pg_glot/pg_glot.control](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/extensions/pg_glot/pg_glot.control)
- [README.md](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/README.md)
- [Cargo.toml](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/Cargo.toml)
- [extensions/pg_glot/Cargo.toml](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/extensions/pg_glot/Cargo.toml)
- [extensions/pg_glot/src/lib.rs](https://github.com/ysys143/pg_glot/blob/7da690a27e1226ea736e8e0eacd0329e4e8ec7e1/extensions/pg_glot/src/lib.rs)

`pg_glot` 通过内嵌 Lindera 词典提供韩语、日语和中文 PostgreSQL 文本检索配置，并提供 `glot.rrf` 融合函数。混合检索组件可独立选装。

### 核心用法

```sql
CREATE EXTENSION pg_glot;
SELECT to_tsvector('japanese', '東京都に住む');
SELECT * FROM glot.rrf(ARRAY[10,20,30]::bigint[], ARRAY[20,40]::bigint[], 60);
```

### 运行边界

由超级用户安装。基础扩展没有其他扩展依赖，也不要求预加载。工作区默认使用 PostgreSQL 17，Cargo 功能开关本身不能证明完整版本矩阵经过测试。上游对韩语的验证最充分。文本检索采用 PostgreSQL 的匹配与排名机制，可使用应用创建的 GIN 索引。`glot.rrf` 合并有序标识符数组，不提供嵌入模型。
