## 用法

来源：

- [README.md](https://github.com/smoothdvd/pg_gse/blob/54346cf5cae9620116aba8744eef1b9bdf808ef0/README.md)
- [Cargo.toml](https://github.com/smoothdvd/pg_gse/blob/54346cf5cae9620116aba8744eef1b9bdf808ef0/Cargo.toml)
- [src/lib.rs](https://github.com/smoothdvd/pg_gse/blob/54346cf5cae9620116aba8744eef1b9bdf808ef0/src/lib.rs)
- [sql/test_bm25.sql](https://github.com/smoothdvd/pg_gse/blob/54346cf5cae9620116aba8744eef1b9bdf808ef0/sql/test_bm25.sql)
- [pg_gse.control](https://github.com/smoothdvd/pg_gse/blob/54346cf5cae9620116aba8744eef1b9bdf808ef0/pg_gse.control)

`pg_gse` 0.1.0 使用 Jieba 提供中文全文检索解析器和配置。分词可独立使用，BM25 排序是与 `pg_textsearch` 的可选集成。上游报告已在 PostgreSQL 17 和 18 中验证运行。

### 核心用法

```sql
CREATE EXTENSION pg_gse;
SELECT * FROM gse_tokenize('自然语言处理');
SELECT * FROM gse_tokenize('自然语言处理', 'search');
SELECT to_tsvector('public.gse_chinese', '自然语言处理');
```

### 运行边界

`gse_tokenize()` 返回词元及字节偏移。双参数形式接受 `precise`、`search` 或 `all`；`gse_set_mode()` 与 `gse_get_mode()` 管理当前后端的模式。`gse` 解析器与 `gse_chinese` 配置接入 PostgreSQL 常规全文检索。安装需要超级用户，扩展不可重定位，本身无需预加载。使用 BM25 时须安装 `pg_textsearch`，按其文档预加载该库并重启，配置名应包含模式，例如 `public.gse_chinese`。虽然存在自定义词典表 `pg_gse_custom_word`，但此修订的 `gse_sync_custom_word()` 尚不会加载词典变更，不应将 README 的词典示例视为已实现功能。Cargo 提供 PG14–18 选项，明确的运行验证仅覆盖 17/18。
