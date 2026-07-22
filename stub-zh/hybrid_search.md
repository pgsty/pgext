## 用法

来源：

- [官方 hybrid_search 软件包页面](https://database.dev/langchain/hybrid_search)
- [版本 1.1.1 软件包 SQL 与文档](https://github.com/supabase/dbdev/blob/8b3b966d97f87027a36c052e679ebb45f7e25736/supabase/migrations/20231207112129_langchain%40hybrid_search-1.1.1.sql)
- [官方 dbdev 项目 README](https://github.com/supabase/dbdev/blob/8b3b966d97f87027a36c052e679ebb45f7e25736/README.md)

`hybrid_search` 是 database.dev 软件包 `langchain@hybrid_search` 的目录名称。它安装一个小型 LangChain/Supabase 模式，在同一个 `documents` 表上结合 pgvector 余弦距离搜索与 PostgreSQL 全文排名。

### 核心流程

通过 `dbdev` 安装软件包，创建所需的 `vector` 依赖，并启用版本 `1.1.0` 的注册表扩展：

```sql
SELECT dbdev.install('langchain@hybrid_search');
CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION "langchain@hybrid_search"
    SCHEMA public
    VERSION '1.1.0';

INSERT INTO documents (content, metadata)
VALUES ('PostgreSQL supports full-text search', '{"topic":"postgres"}');

SELECT *
FROM kw_match_documents('full text', 10);
```

安装的对象包括 `documents(id, content, metadata, embedding)`、`match_documents(query_embedding vector(1536), match_count integer, filter jsonb)` 和 `kw_match_documents(query_text text, match_count integer)`。`match_documents` 按余弦距离排序并支持 JSONB 包含过滤器；`kw_match_documents` 对 `to_tsvector(content)` 的匹配结果进行排名。

### 运维说明

该软件包依赖 `vector`，且官方版本 `1.1.0` 说明依赖解析需要手工完成。其表和函数使用固定的非限定名称，embedding 列为 `vector(1536)`；应检查名称冲突，并相应适配所用的嵌入模型。两个函数分别返回结果集；合并、去重与重新排名由应用侧混合检索器完成，而不是由该软件包中的单一 SQL 函数完成。
