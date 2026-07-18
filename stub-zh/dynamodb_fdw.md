## 用法

来源：

- [上游 README 与选项参考](https://github.com/pgspider/dynamodb_fdw/blob/1d79f84b455705a1fb778a0c3ae35483ebd0d40c/README.md)
- [扩展 control 文件](https://github.com/pgspider/dynamodb_fdw/blob/1d79f84b455705a1fb778a0c3ae35483ebd0d40c/dynamodb_fdw.control)
- [权威选项校验实现](https://github.com/pgspider/dynamodb_fdw/blob/1d79f84b455705a1fb778a0c3ae35483ebd0d40c/option.c)
- [发行元数据](https://github.com/pgspider/dynamodb_fdw/blob/1d79f84b455705a1fb778a0c3ae35483ebd0d40c/META.json)

`dynamodb_fdw` 把 Amazon DynamoDB 表映射为 PostgreSQL 外部表。它支持 `SELECT`、`INSERT`、`UPDATE` 和 `DELETE`，可下推受支持的 `WHERE` 表达式及 DynamoDB `SIZE` 操作，并能用 `->` 和 `->>` 访问嵌套 map/list 值。

### 定义外部表

```sql
CREATE EXTENSION dynamodb_fdw;

CREATE SERVER dynamodb_local
  FOREIGN DATA WRAPPER dynamodb_fdw
  OPTIONS (endpoint 'http://localhost:8000');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER dynamodb_local
  OPTIONS (user 'local-access-key', password 'local-secret-key');

CREATE FOREIGN TABLE public.orders (
  order_id text,
  created_at numeric,
  payload jsonb
)
SERVER dynamodb_local
OPTIONS (table_name 'orders', partition_key 'order_id', sort_key 'created_at');

SELECT order_id, payload->>'status'
FROM public.orders
WHERE order_id = 'A-100';
```

服务端选项为 `endpoint`；用户映射选项为 `user` 和 `password`；表选项为 `partition_key`、`sort_key` 和 `table_name`；列选项为 `column_name`。README 的一处示例使用了 `username`，但校验器实际接受的是上例中的 `user`。用户映射包含凭据，必须保护其访问权限。

该 FDW 不支持 `IMPORT FOREIGN SCHEMA`、外部表 `TRUNCATE` 或 `COPY FROM`。DynamoDB set 会映射为数组，但顺序可能改变；set 中的空值会替换为类型默认值，缺失的 map 属性也不满足文档所述的 `IS NULL` 下推。README 面向 PostgreSQL 13 至 17，并要求 AWS C++ SDK。目录/control 版本是 `1.1`，而同一提交的发行元数据写的是 `2.1.1`；应把它们视为不同版本通道，不要假定可以互换。
