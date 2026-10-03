## 用法

来源：

- [Official README](https://bitbucket.org/adunstan/blackhole_fdw/src/b97c35672aa8618fd84dae6bb11714fdf23acef2/README.md)
- [Extension control file](https://bitbucket.org/adunstan/blackhole_fdw/src/b97c35672aa8618fd84dae6bb11714fdf23acef2/blackhole_fdw.control)
- [Official SQL example](https://bitbucket.org/adunstan/blackhole_fdw/src/b97c35672aa8618fd84dae6bb11714fdf23acef2/sql/blackhole_fdw.sql)

`blackhole_fdw` 是 Andrew Dunstan 提供的独立 FDW 教学骨架，也可用作刻意丢弃数据的接收端。插入行会被丢弃，更新和删除不产生实际存储效果，查询返回空结果集。绝不可用它保存需要留存的数据。

### 基本用法

安装库与扩展 SQL 后，管理员可创建用于丢弃数据的外部表。插入成功并不表示数据已被保存。

```sql
CREATE EXTENSION blackhole_fdw;
CREATE SERVER sink FOREIGN DATA WRAPPER blackhole_fdw;
CREATE FOREIGN TABLE discarded (id integer, payload text) SERVER sink;
INSERT INTO discarded VALUES (1, 'intentionally discarded');
SELECT * FROM discarded;
```

### 开发边界

作者将其作为 FDW 开发起点，包含扫描和修改回调。control 版本为 `0.0.1`，SQL 对象可迁移，上游未要求预加载。源码未提供当前 PostgreSQL 大版本测试矩阵。规范上游位于 Bitbucket，GitHub 仓库为导入镜像。应将其视为刻意丢弃数据的演示设施，而非归档或复制机制。
