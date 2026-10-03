## 用法

来源：

- [production/fdw/README.md](https://github.com/gaia-platform/GaiaPlatform/blob/d0473e6e5710b3dbfc4ab674dda9ebaee8a5a5d9/production/fdw/README.md)
- [production/CMakeLists.txt](https://github.com/gaia-platform/GaiaPlatform/blob/d0473e6e5710b3dbfc4ab674dda9ebaee8a5a5d9/production/CMakeLists.txt)
- [production/fdw/CMakeLists.txt](https://github.com/gaia-platform/GaiaPlatform/blob/d0473e6e5710b3dbfc4ab674dda9ebaee8a5a5d9/production/fdw/CMakeLists.txt)
- [production/fdw/src/gaia_fdw.control](https://github.com/gaia-platform/GaiaPlatform/blob/d0473e6e5710b3dbfc4ab674dda9ebaee8a5a5d9/production/fdw/src/gaia_fdw.control)
- [production/fdw/src/gaia_fdw.sql](https://github.com/gaia-platform/GaiaPlatform/blob/d0473e6e5710b3dbfc4ab674dda9ebaee8a5a5d9/production/fdw/src/gaia_fdw.sql)
- [production/fdw/src/CMakeLists.txt](https://github.com/gaia-platform/GaiaPlatform/blob/d0473e6e5710b3dbfc4ab674dda9ebaee8a5a5d9/production/fdw/src/CMakeLists.txt)
- [production/fdw/tests/airport_setup.sql](https://github.com/gaia-platform/GaiaPlatform/blob/d0473e6e5710b3dbfc4ab674dda9ebaee8a5a5d9/production/fdw/tests/airport_setup.sql)

`gaia_fdw` 0.6 将 Gaia Platform 数据库暴露为 PostgreSQL 外部表。CMake 从 production 0.6.0 的主、次版本生成扩展版本。所引用源码来自 2022 年，不能据此认定它兼容当前 PostgreSQL 主版本。

### 核心工作流

启动 Gaia 数据库并准备好其中的机场数据集后，安装包装器并导入表：

```sql
CREATE EXTENSION gaia_fdw;
CREATE SERVER gaia FOREIGN DATA WRAPPER gaia_fdw;
CREATE SCHEMA airport_fdw;
IMPORT FOREIGN SCHEMA airport_fdw FROM SERVER gaia INTO airport_fdw;
```

### 前提与对象

安装默认需要超级用户权限。可重定位的控制文件加载 `gaia_fdw-0.6`，构建时链接 Gaia 客户端、目录与载荷库。SQL 创建 `gaia_fdw_handler`、`gaia_fdw_validator` 和 `gaia_fdw` 外部数据包装器，没有声明扩展自身的预加载要求。访问表之前，Gaia 客户端必须能连接外部数据库。

### 数据语义

包装器支持读写 Gaia 数据，但 NULL 语义与 PostgreSQL 不同。已有的 NULL 字符串可以读为 NULL；向非引用字段写入 NULL 会被忽略，可能保留空字符串、零等序列化默认值。向引用字段写入 NULL 会移除引用。`gaia_id` 不可更新。依赖 SQL NULL 语义的应用在使用导入表之前，应先核对这些行为。
