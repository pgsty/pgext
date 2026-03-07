
## 用法

- https://pgroonga.github.io/
- [新闻](https://pgroonga.github.io/news/)：发布版本更新信息。
- [概述](https://pgroonga.github.io/overview/)：介绍 PGroonga 的基本信息。
- [安装](https://pgroonga.github.io/install/)：说明如何安装 PGroonga。
- [升级](https://pgroonga.github.io/upgrade/)：说明如何升级 PGroonga。
- [卸载](https://pgroonga.github.io/uninstall/)：说明如何卸载 PGroonga。
- [教程](https://pgroonga.github.io/tutorial/)：逐步演示 PGroonga 的使用方法。
- [常见问题](https://pgroonga.github.io/faq/)：常见问题解答。
- [使用技巧](https://pgroonga.github.io/how-to/)：针对特定场景的实用信息。
- [参考手册](https://pgroonga.github.io/reference/)：详细介绍各项功能，包括选项、函数和运算符。
- [故障排查](https://pgroonga.github.io/troubleshooting/)：说明如何排查和解决问题。
- [社区](https://pgroonga.github.io/community/)：介绍 PGroonga 社区。
- [用户](https://pgroonga.github.io/users/)：列出 PGroonga 的用户。
- [开发](https://pgroonga.github.io/development/)：说明如何参与 PGroonga 的开发。

以下是一个关于如何使用 PGroonga 的快速[教程](https://pgroonga.github.io/tutorial/)：

```sql
CREATE EXTENSION IF NOT EXISTS pgroonga;

CREATE TABLE memos
(
    id      integer,
    content text
);

CREATE INDEX pgroonga_content_index ON memos USING pgroonga (content);

INSERT INTO memos VALUES (1, 'PostgreSQL is a relational database management system.');
INSERT INTO memos VALUES (2, 'Groonga is a fast full text search engine that supports all languages.');
INSERT INTO memos VALUES (3, 'PGroonga is a PostgreSQL extension that uses Groonga as index.');
INSERT INTO memos VALUES (4, 'There is groonga command.');

SET enable_seqscan = off;

-- 现在让我们使用 pgroonga 进行查询

SELECT * FROM memos WHERE content &@ 'engine';
--  id |                                content
-- ----+------------------------------------------------------------------------
--   2 | Groonga is a fast full text search engine that supports all languages.
-- (1 row)

SELECT * FROM memos WHERE content &@~ 'PGroonga OR PostgreSQL';
--  id |                            content
-- ----+----------------------------------------------------------------
--   3 | PGroonga is a PostgreSQL extension that uses Groonga as index.
--   1 | PostgreSQL is a relational database management system.
-- (2 rows)

SELECT * FROM memos WHERE content LIKE '%engine%';
--  id |                                content
-- ----+------------------------------------------------------------------------
--   2 | Groonga is a fast full text search engine that supports all languages.
-- (1 row)
```
