## 用法

来源：

- [官方 README](https://github.com/nkhorman/json_fdw/blob/65b8e7d4dc39bb2844879bb3c0199588e9cfa8a2/README.md)
- [扩展 SQL](https://github.com/nkhorman/json_fdw/blob/65b8e7d4dc39bb2844879bb3c0199588e9cfa8a2/json_fdw--1.0.sql)
- [FDW 实现](https://github.com/nkhorman/json_fdw/blob/65b8e7d4dc39bb2844879bb3c0199588e9cfa8a2/json_fdw.c)

此仓库中的 `json_fdw` 是 `json_fdw2` 分支：一个面向 PostgreSQL 9.4 时代的外部数据包装器，用于读取本地、gzip 压缩或通过 HTTP 获取的逐行 JSON。它支持嵌套字段和数组投影；其实验性远程写映射不能当作事务数据库接口。

### 核心读取流程

物理文件的每一行都必须是一个完整 JSON 文档。嵌套成员使用带引号的点分列名映射。缺失成员和不兼容值可能显示为空值；只有在明确接受坏文档时才应设置 `max_error_count`。

```sql
CREATE EXTENSION json_fdw;
CREATE SERVER json_server FOREIGN DATA WRAPPER json_fdw;

CREATE FOREIGN TABLE customer_reviews (
    customer_id text,
    "review.date" date,
    "review.rating" integer,
    "product.id" text,
    "product.similar_ids" text[]
)
SERVER json_server
OPTIONS (
    filename '/srv/json/customer_reviews.ndjson.gz',
    max_error_count '0'
);

ANALYZE customer_reviews;

SELECT customer_id, "review.rating"
FROM customer_reviews
WHERE "review.date" >= date '2026-01-01';
```

`filename` 也可以是 HTTP URL。内容会缓存在本地，并在每次查询时使用 ETag 重新验证；`http_post_vars` 提供表单式 POST 参数。HTTPS 行为取决于所链接的 libcurl 构建，上游并未测试。

### 远程操作映射

执行远程 API 操作时，互斥的 `rom_url` 与 `rom_path` 选项会选择版本 2 的远程操作映射文档。实现把读取映射为 GET，把插入或更新映射为 PUT；DELETE 仍未实现。

远程写入是外部 HTTP 副作用，不会随 PostgreSQL 事务回滚。实现也不会可靠地把每个 libcurl 或 HTTP 失败转换为 SQL 错误，因此 SQL 完成并不能证明远端资源已改变。即使只做测试，也应使用幂等键、应用层响应校验、审计日志和对账。

### 安全与兼容性

授予服务器使用权可能通过 `filename` 暴露服务端可读文件，或允许请求内网端点。应限制 FDW 和服务器的所有权及 `USAGE`，在扩展外实施路径与目标白名单，并禁止不可信用户定义表选项。

该分支依赖特定的 `nkhorman/yajl` `json_path` 分支、libcurl 和 zlib，上游明确只支持 PostgreSQL 9.4。它没有维护中的现代兼容矩阵。生产负载应选择仍在维护的 FDW 或独立摄取服务。
