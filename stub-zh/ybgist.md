## 用法

来源：

- [README.md](https://github.com/makeafide/yb-pggist/blob/0ecfd02df2eca84a1cbfbdc977789a12331680ca/README.md)
- [sql/ybgist--1.2.sql](https://github.com/makeafide/yb-pggist/blob/0ecfd02df2eca84a1cbfbdc977789a12331680ca/sql/ybgist--1.2.sql)
- [docs/INSTALL-RELEASE.md](https://github.com/makeafide/yb-pggist/blob/0ecfd02df2eca84a1cbfbdc977789a12331680ca/docs/INSTALL-RELEASE.md)
- [ybgist.control](https://github.com/makeafide/yb-pggist/blob/0ecfd02df2eca84a1cbfbdc977789a12331680ca/ybgist.control)

`ybgist` 为特定 YugabyteDB 分支内置的访问方法提供空间运算符类和检查函数，不能安装到原版 PostgreSQL 或原版 YugabyteDB。

### 核心用法

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION ybgist;
SELECT ybgist_tune_costs();
```

### 运行边界

核心分支、修改后的 PostGIS 和扩展必须成套兼容部署，C++ 运行时和 ABI 要求也是安装约束的一部分。先创建 PostGIS，再由超级用户创建 ybgist，固定模式为 public。先批量载入数据，再用独立语句创建索引。`ybgist_record_index` 记录构建设置，`ybgist_verify_index` 检查回填完整性，`ybgist_tune_costs` 调整谓词成本。

源码 control 版本为 1.2；已发布的 r2 组合包早于当前功能。文档描述了 geometry、geography 和标量/空间多列索引，不支持仅索引扫描和最近邻排序。修改映射或覆盖参数时必须遵循上游漂移检查与重建流程。检查的源码树中未找到项目许可证声明。
