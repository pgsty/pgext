## 用法

来源：

- [MEST README](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/README.md)
- [Control](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/mest.control)
- [SQL 1.0](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/mest--1.0.sql)

`mest` 1.0 提供多条目 GiST 与 SP-GiST 索引访问方法。一个源值可以产生多个索引条目，适合多范围和路径数据。上游文档支持 PostgreSQL 17、18。

### 基本流程

安装与服务器 PostgreSQL 主版本匹配的库，再由超级用户启用：

```sql
CREATE EXTENSION mest;

CREATE TABLE time_windows (id bigint, spans tstzmultirange);
CREATE INDEX time_windows_mgist ON time_windows USING mgist (spans);
CREATE INDEX time_windows_mspgist ON time_windows USING mspgist (spans);
```

根据实际数据比较查询计划和索引大小后选择适合的方法；示例展示了两种可用方法。

### 主要对象与配套扩展

- `mgist`、`mspgist` 是索引访问方法。
- `multirange_mrtree_ops`、`multirange_mquadtree_ops` 为 `anymultirange` 提供默认操作符类。
- `path_mrtree_ops`、`path_mquadtree_ops` 支持 `path`。
- `ranges()` 为整数、日期和时间戳多范围提供重载。
- `postgis_mest` 增加几何操作符类；`mobilitydb_mest` 增加 MobilityDB 类型的操作符类。按需单独安装这些扩展。

### 运维边界

基础扩展的控制文件允许重定位，没有声明扩展依赖。安装脚本创建访问方法与 C 函数，因此安装需要相应特权。基础扩展没有文档要求的预加载或重启步骤；配套扩展各有库文件和预加载要求。

一行对应多个条目会改变索引大小及维护开销。查询必须使用所选操作符类支持的操作符；安装扩展不会让所有谓词都能使用索引。
