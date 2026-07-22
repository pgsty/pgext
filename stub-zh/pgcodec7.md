## 用法

来源：

- [官方 README](https://github.com/AbdulYadi/pgcodec7/blob/b54cdcbadb03042396a65b8db7f9913fb0686fc1/README.md)
- [官方扩展 SQL](https://github.com/AbdulYadi/pgcodec7/blob/b54cdcbadb03042396a65b8db7f9913fb0686fc1/pgcodec7--1.0.sql)
- [官方 C 实现](https://github.com/AbdulYadi/pgcodec7/blob/b54cdcbadb03042396a65b8db7f9913fb0686fc1/pgcodec7.c)

`pgcodec7` 1.0 把 `bytea` 转换成一组 6 位或 7 位文本分块，并从 `text[]` 还原原始字节。它可以为特定传输减少十六进制编码膨胀，但既不是压缩，也不是加密。

### 核心流程

实际 SQL API 与旧 README 示例不同：`pgencode7(integer, bytea)` 接受字节而不是文件名，`pgdecode7(integer, text[])` 需要分块数组而不是单个文本值。聚合时必须保持分块顺序。

```sql
CREATE EXTENSION pgcodec7;

SELECT convert_from(
           pgdecode7(6, array_agg(chunk ORDER BY ord)),
           'UTF8'
       ) AS decoded
FROM pgencode7(6, convert_to('hello', 'UTF8'))
     WITH ORDINALITY AS encoded(chunk, ord);
```

`pgencode7` 会返回多行，因为每个输出分块被限制在约 1024 个字符以内。如果分块要经过不能保持行序的表、队列或传输层，请显式保存序号。

### 函数

- `pgencode7(i_encodelen integer, by_data bytea) RETURNS SETOF text` 只接受模式 6 或 7。
- `pgdecode7(i_decodelen integer, t_encodes text[]) RETURNS bytea` 连接并解码有序的分块数组。

两个函数都安装在 `public` 中，扩展 SQL 将执行权限授予 `PUBLIC`。

### 编码与安全边界

6 位输出使用更便于传输的字母表。7 位字母表包含 0x80–0xAF 范围的字节，它们不是独立有效的 UTF-8 字符；许多文本客户端、JSON 编码器、排序规则和网关可能拒绝或改写这些字节。只有在经过端到端测试且能保留原始字节的路径上才能使用 7 位模式。

解码器的 C 实现并不会在执行算术前安全拒绝字母表之外的每个字节。不要把不可信或被改写的分块直接交给 `pgdecode7`；应先验证模式、字母表、分块数量、总大小和顺序。畸形或在传输中受损的输入可能导致错误或后端不稳定。

上游项目针对 PostgreSQL 11 构建，自 2019 年后未再活跃。依赖它之前，应在准确的目标 PostgreSQL 版本上测试构建兼容性和往返结果。
