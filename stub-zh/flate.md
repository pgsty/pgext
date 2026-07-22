## 用法

来源：

- [官方 README](https://github.com/grahamedgecombe/pgflate/blob/a6880b7f5c2304b286e944945482efc7226ec18f/README.markdown)
- [官方扩展 SQL](https://github.com/grahamedgecombe/pgflate/blob/a6880b7f5c2304b286e944945482efc7226ec18f/flate--1.0.1.sql)
- [官方 C 实现](https://github.com/grahamedgecombe/pgflate/blob/a6880b7f5c2304b286e944945482efc7226ec18f/flate.c)

`flate` 为 PostgreSQL `bytea` 值提供 zlib 原始 DEFLATE 压缩与解压。版本 `1.0.1` 适合明确需要原始流的应用；其输出不包含 zlib 或 gzip 包装头。

### 核心流程

`deflate(bytea, bytea, integer)` 接受可选的预置词典与压缩级别参数。压缩时使用词典后，`inflate(bytea, bytea)` 必须接收同一个词典：

```sql
CREATE EXTENSION flate;

SELECT deflate(
  convert_to('hello hello hello hello', 'UTF8'),
  convert_to('hello', 'UTF8'),
  9
);

SELECT convert_from(
  inflate(E'\\xcb00110a182400'::bytea, convert_to('hello', 'UTF8')),
  'UTF8'
);
```

NULL 级别使用 zlib 默认值；显式级别遵循 zlib 的 0 到 9 范围。两个函数均声明为不可变。

### 运维注意事项

压缩值中不会嵌入词典。词典缺失或不一致、输入损坏，或流包装格式错误时，解压会失败或产生不可用数据。

压缩输入可能膨胀为大得多的结果。调用 `inflate(bytea, bytea)` 前应限制不受信任输入的大小，并考虑后端内存与语句超时。应随数据保存编码、包装方式和词典版本，以便其他客户端复现操作。
