## 用法

来源：

- [官方 README](https://github.com/hexiaoting/rasterdb/blob/7241bd185509371b08651f9c04876144da59772c/README.md)
- [官方控制文件](https://github.com/hexiaoting/rasterdb/blob/7241bd185509371b08651f9c04876144da59772c/rasterdb.control)
- [官方安装 SQL](https://github.com/hexiaoting/rasterdb/blob/7241bd185509371b08651f9c04876144da59772c/rasterdb--0.1.sql)
- [官方 FDW 实现](https://github.com/hexiaoting/rasterdb/blob/7241bd185509371b08651f9c04876144da59772c/rasterdb.c)
- [官方栅格加载器](https://github.com/hexiaoting/rasterdb/blob/7241bd185509371b08651f9c04876144da59772c/load_raster.c)

`rasterdb` 是实验性的只读外部数据包装器：它打开 PostgreSQL 服务器上的栅格文件，通过 GDAL 与 PostGIS 栅格内部接口转换，并为每个切片返回一个 `raster` 值。它并不访问远程数据库；数据源是服务器本地文件或目录。

### 核心流程

安装 PostGIS 与该扩展。在服务器端创建包含绝对位置的配置文件，再让只有一列的外部表引用该文件。

```text
tile_size=100x100
location=/srv/postgresql/rasters
batchsize=50
```

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION rasterdb;

CREATE SERVER raster_server
  FOREIGN DATA WRAPPER rasterdb_fdw;

CREATE FOREIGN TABLE raster_tiles (rast raster)
  SERVER raster_server
  OPTIONS (conf_file '/etc/postgresql/rasterdb.conf');

SELECT count(*) FROM raster_tiles;
```

`tile_size` 选择切片尺寸；省略时，每个源栅格作为一个切片返回。`location` 指定单个栅格文件，或指定一个目录以扫描 GDAL 能识别的文件。`batchsize` 控制 FDW 每批物化的切片数量。

### 对象与查询行为

扩展会创建 `rasterdb_fdw` 及其处理器和验证器。外部表必须提供 `conf_file` 选项，而且必须只有一个 PostGIS `raster` 列。扫描时按需读取并转换栅格文件；FDW 没有写回调、重新扫描支持、统计信息收集或谓词下推。

### 安全与兼容性

PostgreSQL 服务器进程会打开配置文件以及 `location` 选择的每个路径。应限制创建或修改这些外部表的权限，把配置文件放在应用不可写的目录，并只授予数据库操作系统账户所需的文件系统权限。

已复核的 `0.1` 代码在进程范围内启用所有 PostGIS GDAL 驱动，假设输入栅格只有一个波段，在空操作验证器中接受任意选项，并使用 2018 年的 PostgreSQL 与 PostGIS 内部 API。它还存在固定长度文件名处理和不完整的规划器/执行器回调。应把它视作隔离数据转换原型，而不是仍受维护的生产 FDW；运行前必须测试具体文件格式、路径、内存用量与服务器版本兼容性。
