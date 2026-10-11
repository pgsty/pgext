## 用法

来源：

- [extension/postgis_hex9/README.md](https://github.com/MrBenGriffin/libhex9/blob/1eb0be7d43255b30ecb0c46611fd56b10bf7d773/extension/postgis_hex9/README.md)
- [extension/postgis_hex9/postgis_hex9--2.3.0.sql](https://github.com/MrBenGriffin/libhex9/blob/1eb0be7d43255b30ecb0c46611fd56b10bf7d773/extension/postgis_hex9/postgis_hex9--2.3.0.sql)
- [README.md](https://github.com/MrBenGriffin/libhex9/blob/1eb0be7d43255b30ecb0c46611fd56b10bf7d773/README.md)
- [extension/postgis_hex9/postgis_hex9.control](https://github.com/MrBenGriffin/libhex9/blob/1eb0be7d43255b30ecb0c46611fd56b10bf7d773/extension/postgis_hex9/postgis_hex9.control)

`postgis_hex9` 为 PostGIS 增加 Hex9 地理网格，使用 UUID 标识单元并进行分层空间聚合。它依赖 `postgis` 与原生 libhex9。默认布局达到第 30 层，传统布局达到第 29 层。

### 核心用法

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_hex9;
SELECT h9_encode('SRID=4326;POINT(-3.19 55.95)'::geometry);
SELECT h9_bin(h9_encode('SRID=4326;POINT(-3.19 55.95)'::geometry), 8);
```

### 运行边界

`h9_encode()` 与 `h9_decode()` 转换点和完整地址，`h9_bin()` 返回聚合键。`h9_cell()` 与 `h9_grid()` 生成几何，邻域函数以完整地址为输入。无需预加载或显式加载动态库。聚合键不是完整点地址。WGS84 与球面输入产生不同地址，UUID 不记录基准或投影规则，应在数据集中另存这些元数据。1.x 地址升级到 2.x 时必须从原始几何重新计算，不能解码旧地址再编码，也不能混用投影规则。控制文件默认要求超级用户安装，上游未声明 PostgreSQL 主版本支持矩阵。
