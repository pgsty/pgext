## 用法

来源：

- [contrib/alohadb_temporal/alohadb_temporal--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_temporal/alohadb_temporal--1.0.sql)
- [contrib/alohadb_temporal/alohadb_temporal.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_temporal/alohadb_temporal.c)
- [README.md](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/README.md)
- [contrib/alohadb_temporal/alohadb_temporal.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_temporal/alohadb_temporal.control)

`alohadb_temporal` 是 AlohaDB 的系统版本管理模块。它为基础表关联历史表并安装行触发器，在更新和删除时保存旧行版本。

### 核心用法

```sql
CREATE EXTENSION alohadb_temporal;
CREATE TABLE temporal_demo (id integer PRIMARY KEY, label text);
SELECT alohadb_enable_system_versioning('temporal_demo'::regclass);
INSERT INTO temporal_demo (id, label) VALUES (1, 'first');
UPDATE temporal_demo SET label = 'second' WHERE id = 1;
SELECT * FROM alohadb_temporal_status();
```

### 运行边界

该模块属于基于 PostgreSQL 18 的 AlohaDB 分支，尚未验证标准 PostgreSQL 兼容性。原生函数需要管理员安装。`alohadb_enable_system_versioning` 会修改基础表，增加时间列、历史存储和触发器，因此调用者必须能修改目标表。`alohadb_temporal_status` 列出注册信息。`alohadb_as_of` 和 `alohadb_versions_between` 返回记录，调用时需要声明列类型。历史数据随写入增长，应规划保留策略。调用 `alohadb_disable_system_versioning` 前应审查实现，它会移除版本管理设施和时间列。此源码原型不构成覆盖整个应用的审计或防篡改保证。
