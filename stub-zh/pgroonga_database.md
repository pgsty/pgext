

## 用法

> [PGroonga 文档](https://pgroonga.github.io/) | [GitHub: pgroonga/pgroonga](https://github.com/pgroonga/pgroonga)

`pgroonga_database` 是 [PGroonga](https://pgroonga.github.io/) 项目的子扩展。它为 PGroonga 提供数据库管理功能，PGroonga 使 PostgreSQL 成为支持所有语言的快速全文搜索平台。

PGroonga 是一个全面的全文搜索解决方案，以 [Groonga](https://groonga.org/) 作为后端。它开箱即用地支持所有语言（包括中日韩 CJK），并提供丰富功能：

- 支持所有语言的快速全文搜索
- 丰富的查询语法（查询语言、脚本语法）
- JSON 搜索
- 感知 HTML/XML 标签的高亮
- 相似搜索
- 同义词扩展
- 自动补全
- 查询日志分析

PGroonga 文档非常详尽，涵盖数百页。详细用法、API 参考、运算符、函数和调优指南请参见官方文档：

- [PGroonga 官方文档](https://pgroonga.github.io/)
- [入门指南](https://pgroonga.github.io/install/)
- [教程](https://pgroonga.github.io/tutorial/)
- [使用指南](https://pgroonga.github.io/how-to/)
- [参考手册](https://pgroonga.github.io/reference/)

## 快速开始

```sql
CREATE EXTENSION pgroonga_database;
CREATE EXTENSION pgroonga;

-- 创建包含文本内容的表
CREATE TABLE memos (
  id integer,
  content text
);

-- 创建 PGroonga 索引
CREATE INDEX pgroonga_content_index ON memos USING pgroonga (content);

-- 插入数据
INSERT INTO memos VALUES (1, 'PostgreSQL is a relational database management system.');
INSERT INTO memos VALUES (2, 'Groonga is a fast full text search engine that supports all languages.');
INSERT INTO memos VALUES (3, 'PGroonga is a PostgreSQL extension that uses Groonga as its backend.');

-- 全文搜索
SELECT * FROM memos WHERE content &@~ 'PostgreSQL OR Groonga';
```
