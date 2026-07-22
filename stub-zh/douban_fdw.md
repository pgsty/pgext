## 用法

来源：

- [官方 douban_fdw README](https://github.com/xiaowing/douban_fdw/blob/8576d03007de63e22a2a114d6617756dca4412c5/README.md)
- [扩展 SQL](https://github.com/xiaowing/douban_fdw/blob/8576d03007de63e22a2a114d6617756dca4412c5/douban_fdw--1.0.sql)
- [豆瓣排行榜客户端源码](https://github.com/xiaowing/douban_fdw/blob/8576d03007de63e22a2a114d6617756dca4412c5/douban_rank.go)

`douban_fdw` 是只读外部数据包装器，将两个历史豆瓣电影排行榜接口暴露为外部表。它面向旧版 PostgreSQL，以 Go 和 cgo 编写，并依赖旧的无认证豆瓣 API；即使只用于实验，也应先确认远端接口仍可访问。

### 创建排行榜外部表

外部服务器没有选项。每张外部表都必须指定 `rank_name`，实现支持的值为 `top250` 和 `us_box`：

```sql
CREATE EXTENSION douban_fdw;
CREATE SERVER douban_api FOREIGN DATA WRAPPER douban_fdw;

CREATE FOREIGN TABLE douban_top250 (
    id           bigint,
    title        text,
    originname   text,
    rating       real,
    year         text,
    genres       text,
    directors    text,
    casts        text,
    collectcount integer,
    subtype      text,
    url          text
)
SERVER douban_api
OPTIONS (rank_name 'top250');

SELECT id, title, rating
FROM douban_top250;
```

可识别的列为 `casts`、`collectcount`、`directors`、`genres`、`id`、`originname`、`rating`、`subtype`、`title`、`url` 和 `year`。源码明确拒绝拼写为 `imags` 的列。上游 JSON 响应中没有的数据会按包装器转换逻辑变为空值或空内容。

### 运行边界

每次扫描都会通过硬编码的明文 HTTP `/v2/movie/` 端点下载所选排行榜的全部数据。它没有认证、超时、本地缓存、谓词下推或写入能力。上游 README 还记载了每个 IP 每小时 40 次 API 调用的历史限制，并要求数据库使用 UTF-8；这些只能视为历史信息，不能当作当前服务承诺。

由于 API 接口和 PostgreSQL 兼容性说明都很陈旧，应固定源码版本，测试网络故障和响应结构变化，且不要把此包装器用作持久的生产集成。若数据具有运维重要性，应改用受维护且带认证的采集任务。
