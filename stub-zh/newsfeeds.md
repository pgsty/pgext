## 用法

来源：

- [官方仓库 README](https://github.com/ttfkam/pg_newsfeeds/blob/02cb001324498999d7d055453b62dac93a4c4e9e/README.md)
- [1.0.0 版本扩展 SQL](https://github.com/ttfkam/pg_newsfeeds/blob/02cb001324498999d7d055453b62dac93a4c4e9e/newsfeeds--1.0.0.sql)
- [扩展控制文件](https://github.com/ttfkam/pg_newsfeeds/blob/02cb001324498999d7d055453b62dac93a4c4e9e/newsfeeds.control)

`newsfeeds` 发布了一个特定新闻采集应用的数据库模式与查询辅助函数。它保存抓取规则和已收集标题，构造带权全文检索向量，向外部爬虫返回待处理源，并以行或 JSON 形式提供标题搜索。它本身不会抓取网页。

### 核心流程

预期流程是在应用模式中创建扩展、登记新闻源，再由外部工作进程消费 `pending_feeds()` 并填充 `headlines`：

```sql
CREATE SCHEMA aggregator;
CREATE EXTENSION newsfeeds WITH SCHEMA aggregator;

INSERT INTO aggregator.newsfeeds
    (url, entries, title_selector, link_selector, feedname)
VALUES
    ('https://example.com/news', '.entry', '.title', 'a', 'Example News');

SELECT aggregator.pending_feeds();
SELECT * FROM aggregator.headlines('7 days', '', 0.1, 100, 0);
SELECT aggregator.headlines_as_json('7 days', '', 0.1, 100, 0);
```

### 对象索引

- `newsfeeds` 保存源 URL、CSS 选择器、刷新间隔和爬虫状态。
- `headlines` 保存已收集元数据、正文、标签、归档数据和加权 `tsvector`，并在 `fts` 上建立 GIN 索引。
- `clean_query(text)` 重写应用的搜索语法。
- `pending_feeds()` 以 JSON 返回爬虫任务。
- `headlines(...)` 与 `headlines_as_json(...)` 浏览或搜索已收集项目。
- `fts()` 是预期用于维护加权搜索向量的触发器函数。

### 运维说明

已发布的 `1.0.0` SQL 针对特定应用，作为独立扩展并不完整。它含有可疑的列级 `FOREIGN KEY (newsfeed)` 声明，引用了同一脚本未创建的 `title`、`description`、`rank_modifier` 和 `reify_url` 对象，而且尽管控制文件声明扩展可重定位，辅助函数仍硬编码了 `aggregator` 模式。尝试全新安装前，应验证并修正准确的 SQL。

上游没有包含爬虫实现或受支持的部署契约。选择器执行、URL 抓取、内容清理、保留策略与网络安全均应视为外部工作进程的责任。仓库只有一行 README，也没有发行或兼容性说明。
