## 用法

来源：

- [官方 pg_comparator 手册入口](https://github.com/zx80/pg_comparator/blob/01cd47b2fae021b3320eea2f0885297a320d6ec4/README.pg_comparator)
- [命令行手册源码](https://github.com/zx80/pg_comparator/blob/01cd47b2fae021b3320eea2f0885297a320d6ec4/pg_comparator.pl)
- [3.1 版本扩展 SQL](https://github.com/zx80/pg_comparator/blob/01cd47b2fae021b3320eea2f0885297a320d6ec4/pgcmp--3.1.sql)

`pgcmp` 是外部 `pg_comparator` 命令行程序在 PostgreSQL 端的支撑扩展。两者组合通过层次校验和比对带键表，并可选地把第二张表同步为第一张表的内容。

### 核心流程

在每个参与的 PostgreSQL 数据库中安装 `pgcmp`，然后从操作系统运行只读比对。URL 后缀标识表、键列与被比较列。

```sql
CREATE EXTENSION pgcmp;
```

```sh
pg_comparator --ask-pass \
  /family/calvin?id:c1,c2 \
  /family/hobbes
```

命令会报告使第二张表与第一张表一致所需的 `UPDATE`、`INSERT` 和 `DELETE` 键。单独使用 `--synchronize` 是试运行；再加上 `--do-it` 才会实际写入。应从纯比对开始，检查每个差异，备份目标，并在同时启用两个选项前演练恢复。

### 扩展界面

- `cksum2(text)`、`cksum4(text)` 和 `cksum8(text)` 提供默认校验和函数族。
- `fnv2(text)`、`fnv4(text)` 和 `fnv8(text)` 提供 FNV 校验和函数族。
- 扩展还安装该命令在 PostgreSQL 比较路径中使用的 XOR 聚合与类型转换。

### 运维边界

命令要求非空键，并可能扫描、加锁、创建临时摘要表，以及在同步模式中写入目标。应使用 `--ask-pass` 或受保护的环境处理，避免把密码放进进程参数。异构比较可能因类型、编码、排序规则、NULL 处理或类型转换而不一致，校验和冲突还可能造成假阴性。版本 `3.1` 已废弃；当前系统应优先使用受维护的复制或对账工具，并在触碰生产数据前于副本上验证负载、事务、超时和回滚行为。
