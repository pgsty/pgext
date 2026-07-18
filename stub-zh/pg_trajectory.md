## 用法

来源：

- [项目 README](https://github.com/ahmetkucuk/pg-trajectory/blob/1b73624f752149f5c2cf25b1e2aabb616d967b27/README.md)
- [扩展 control 文件](https://github.com/ahmetkucuk/pg-trajectory/blob/1b73624f752149f5c2cf25b1e2aabb616d967b27/pg_trajectory.control)
- [0.0.1 版安装 SQL](https://github.com/ahmetkucuk/pg-trajectory/blob/1b73624f752149f5c2cf25b1e2aabb616d967b27/pg_trajectory--0.0.1.sql)
- [轨迹数据模型](https://github.com/ahmetkucuk/pg-trajectory/blob/1b73624f752149f5c2cf25b1e2aabb616d967b27/model/dataTypes.sql)

`pg_trajectory` 0.0.1 是实验性的纯 SQL PostgreSQL/PostGIS 时空轨迹模型。它把 `tg_pair` 定义为时间戳与几何对象的组合，把 `trajectory` 定义为起止时间、几何类型、包围几何对象和样本数组。源码还包含构造、修改、距离和相似度例程。

### 安装边界

control 文件未声明 PostGIS 依赖，但安装 SQL 引用了 `geometry`。应先安装 PostGIS，并在准确的 PostgreSQL 与 PostGIS 大版本上测试完整脚本：

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pg_trajectory;

SELECT pg_typeof(NULL::tg_pair), pg_typeof(NULL::trajectory);
```

样本时间使用不带时区的 `timestamp`。导入数据前应确定一种时区约定，并在应用边界要求几何类型和 SRID 一致。

### 原型缺陷与兼容性

审查到的 0.0.1 SQL 存在内部不一致。`trajectory` 复合类型没有 `sampling_interval` 字段，但 `_trajectory()` 会为该字段赋值；其他例程还混用或调用未定义的函数名。PL/pgSQL 延迟编译可能使这些路径直到第一次调用才失败。`CREATE EXTENSION` 成功不能视为功能验收完成。

仓库没有当前 release、升级链、兼容矩阵或现代自动化测试证据。使用前必须覆盖每个所需构造器和操作符、错误与空数组、混合几何类型、时间戳、SRID、null 和大型轨迹。在转储恢复和扩展升级行为得到验证之前，不要用其自定义复合布局长期保存数据。生产时空工作负载应优先使用受维护的 PostGIS 原生模式，除非已在本地审计并修复此原型。
