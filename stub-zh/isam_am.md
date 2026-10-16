## 用法

来源：

- [README.md](https://github.com/lMafa1813l/BD2-Gestor-isam-postgres/blob/b6b2d9fec5abadbafeaa01a6d17bca4b1597d1b0/README.md)
- [isam_am/isam_am--1.0.sql](https://github.com/lMafa1813l/BD2-Gestor-isam-postgres/blob/b6b2d9fec5abadbafeaa01a6d17bca4b1597d1b0/isam_am/isam_am--1.0.sql)
- [isam_am/isam_am--1.0--1.1.sql](https://github.com/lMafa1813l/BD2-Gestor-isam-postgres/blob/b6b2d9fec5abadbafeaa01a6d17bca4b1597d1b0/isam_am/isam_am--1.0--1.1.sql)
- [isam_am/Makefile](https://github.com/lMafa1813l/BD2-Gestor-isam-postgres/blob/b6b2d9fec5abadbafeaa01a6d17bca4b1597d1b0/isam_am/Makefile)
- [isam_am/isam_am.control](https://github.com/lMafa1813l/BD2-Gestor-isam-postgres/blob/b6b2d9fec5abadbafeaa01a6d17bca4b1597d1b0/isam_am/isam_am.control)

`isam_am` 1.1 是面向 PostgreSQL 18.6 课程项目的 ISAM 索引实现。安装路径组合 1.0 的 SQL 定义与 1.0 到 1.1 的升级脚本。

### 核心用法

```sql
CREATE EXTENSION isam_am;
CREATE TABLE isam_example (k integer);
SET isam_am.fill_pct = 80;
CREATE INDEX isam_example_idx ON isam_example USING isam_am (k);
SELECT * FROM isam_index_stats('isam_example_idx');
SELECT isam_reorg_needed('isam_example_idx', 25);
```

### 运行边界

默认操作符类覆盖 `int4` 比较，示例使用 `isam_am.fill_pct` 控制填充设置。`isam_index_stats(regclass)` 返回键、数据块、溢出和链统计；`isam_reorg_needed(regclass, float8)` 比较溢出比例与阈值，默认阈值为 25。重组使用 REINDEX，会产生相应的锁与存储开销。默认仅超级用户可安装。上游测试属于项目证据，不等于生产支持承诺；仓库中未找到许可证。
