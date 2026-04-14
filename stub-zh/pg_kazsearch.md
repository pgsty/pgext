## 用法
> 来源: [README](https://github.com/darkhanakh/pg-kazsearch/blob/main/README.md) 和 [项目仓库](https://github.com/darkhanakh/pg-kazsearch)。

`pg_kazsearch` 是一个面向哈萨克语的 PostgreSQL 全文检索扩展。
上游 README 将其描述为使用 `pgrx` 构建的 Rust 扩展，它接入 PostgreSQL 的文本检索流水线。

它会创建一个可直接使用的配置 `kazakh_cfg`，以及配套词典 `pg_kazsearch_dict`。

### 快速开始

```sql
CREATE EXTENSION pg_kazsearch;

SELECT to_tsvector('kazakh_cfg', 'президенттің жарлығы');
-- 'жарлық':2 'президент':1

SELECT ts_lexize('pg_kazsearch_dict', 'алмаларымыздағы');
-- {алма}
```

### 使用场景

README 展示了以下典型用法：

- 对单个哈萨克语词语做词干提取
- 使用 `to_tsvector('kazakh_cfg', ...)` 构建 `tsvector`
- 为表添加生成列类型的 `tsvector`
- 用 GIN 索引这些列
- 使用 `websearch_to_tsquery('kazakh_cfg', ...)` 进行检索

示例表工作流：

```sql
ALTER TABLE articles ADD COLUMN fts tsvector
    GENERATED ALWAYS AS (
        setweight(to_tsvector('kazakh_cfg', title), 'A') ||
        setweight(to_tsvector('kazakh_cfg', body), 'B')
    ) STORED;

CREATE INDEX idx_fts ON articles USING GIN (fts);

SELECT title FROM articles
WHERE fts @@ websearch_to_tsquery('kazakh_cfg', 'президенттің жарлығы')
ORDER BY ts_rank_cd(fts, websearch_to_tsquery('kazakh_cfg', 'президенттің жарлығы')) DESC
LIMIT 10;
```

### 调优

可以在运行时调整惩罚权重：

```sql
ALTER TEXT SEARCH DICTIONARY pg_kazsearch_dict (w_deriv = 3.5, w_short_char = 100.0);
```

### 部署

README 列出了三种支持的安装路径：

- 预编译的 Debian/Ubuntu 软件包
- 基于 `ghcr.io/darkhanakh/pg-kazsearch` 的 Docker 镜像
- 使用 `cargo pgrx install` 从源码构建

本项目的仓库元数据对应 PostgreSQL 16-18。
