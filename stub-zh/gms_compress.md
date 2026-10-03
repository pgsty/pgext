## 用法

来源：

- [Control 1.0](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/contrib/gms_compress/gms_compress.control)
- [SQL 1.0](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/contrib/gms_compress/gms_compress--1.0.sql)
- [Official examples](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/contrib/gms_compress/sql/gms_compress.sql)

`gms_compress` 1.0 在 openGauss 中通过 zlib 压缩和解压 `RAW` 与 `BLOB` 值。其 SQL 使用内核特有类型和 `PACKAGE` 语法，不能据此认为兼容原生 PostgreSQL。官方 openGauss 构建已包含该模块。

### 基本用法

使用有权创建扩展的角色，在目标 openGauss 数据库中安装：

```sql
CREATE EXTENSION gms_compress;
SELECT gms_compress.lz_compress('12ab56'::raw, 6);
SELECT gms_compress.lz_uncompress(
  gms_compress.lz_compress('12ab56'::raw, 6)
);
```

扩展创建固定的 `gms_compress` schema。压缩级别范围为 1 到 9，默认为 6。单次调用的函数重载返回压缩或解压结果，过程重载还提供输出参数。

### 流式接口

- `lz_compress_open` 创建上下文并返回整数句柄。
- `lz_compress_add` 输入 RAW 数据块，`lz_compress_close` 完成压缩后的 BLOB 并关闭上下文。
- `lz_uncompress_open` 接受压缩后的 BLOB，`lz_uncompress_extract` 提取解压后的 RAW 值，`lz_uncompress_close` 关闭上下文。
- `isopen` 检查句柄是否处于打开状态。这些函数和过程都位于扩展的 schema 中。

### 限制与维护

最多同时使用五个压缩或解压句柄，每个句柄处理的数据不能超过 1 GiB。完成操作后关闭上下文，不要将句柄作为持久化应用标识保存。当前实现中，流式压缩添加过程的目标参数仅为语法兼容保留，不用于传递状态。该模块处理 RAW/BLOB 值，不提供透明表压缩。升级时核对对应的 openGauss 发行版，不能从其内核来源推断 PostgreSQL 大版本兼容范围。
