## 用法

来源：

- [README.rst](https://github.com/dvarrazzo/jsonb_parser/blob/78128f3094e976a55adb04c9fe287b7ac2045e60/README.rst)
- [pg_ubjson/ubjson--0.0.1.sql](https://github.com/dvarrazzo/jsonb_parser/blob/78128f3094e976a55adb04c9fe287b7ac2045e60/pg_ubjson/ubjson--0.0.1.sql)
- [pg_ubjson/Makefile](https://github.com/dvarrazzo/jsonb_parser/blob/78128f3094e976a55adb04c9fe287b7ac2045e60/pg_ubjson/Makefile)
- [pg_ubjson/ubjson.control](https://github.com/dvarrazzo/jsonb_parser/blob/78128f3094e976a55adb04c9fe287b7ac2045e60/pg_ubjson/ubjson.control)

`ubjson` 是 jsonb format parser 项目的 PostgreSQL 组件。它沿用 JSONB 的文本输入输出与存储，使用自己的二进制发送和接收函数与客户端交换数据。

### 核心用法

```sql
CREATE EXTENSION ubjson;
SELECT '{"answer":42}'::jsonb::ubjson;
SELECT ('{"answer":42}'::jsonb::ubjson)::jsonb;
```

### 运行边界

该类型采用 JSONB 的文本表示，并提供双向无复制赋值转换。二进制表示是单独的协议，客户端须使用对应解码器，不能当成普通 PostgreSQL JSONB 线协议数据。它没有新增通用 JSON 查询语言。

创建 C 扩展要求超级用户，上游未要求预加载。这是实验源码，未声明 PostgreSQL 主版本支持矩阵；升级客户端或服务器前应验证二进制兼容性。
