## 用法

来源：

- [官方 README](https://github.com/gleu/autovacuum_informations/blob/0dde21773a3929e0b0bacca018607d911fa71f00/README.md)
- [官方扩展 SQL](https://github.com/gleu/autovacuum_informations/blob/0dde21773a3929e0b0bacca018607d911fa71f00/autovacuum_informations--1.0.sql)
- [官方实现](https://github.com/gleu/autovacuum_informations/blob/0dde21773a3929e0b0bacca018607d911fa71f00/autovacuum_informations.c)

`autovacuum_informations` 1.0 通过两个 C 函数暴露 PostgreSQL 内存中的 autovacuum launcher 与 worker 状态。它是与 PostgreSQL 内部实现紧密耦合的诊断原型；上游明确说明该项目仍处于开发中，禁止用于生产环境。

### 核心流程

安装扩展后，可以读取 launcher PID，或展开每个活动 worker 返回的记录：

```sql
CREATE EXTENSION autovacuum_informations;

SELECT get_autovacuum_launcher_pid();

SELECT *
FROM get_autovacuum_workers_infos() AS w(
  idx integer,
  datid oid,
  relid oid,
  pid integer,
  launchtime timestamptz,
  dobalance boolean,
  cost_delay integer,
  cost_limit integer,
  cost_limit_base integer
);
```

launcher 函数返回 `bigint` 类型的 PID；没有可用 PID 时返回 null。worker 函数声明为 `SETOF record`，因此调用方必须像上例一样提供九列记录定义。

### 重要对象

- `get_autovacuum_launcher_pid()` 从 autovacuum 共享内存读取 launcher 进程 ID。
- `get_autovacuum_workers_infos()` 返回当前 worker 序号、数据库与关系 OID、PID、启动时间、负载均衡标志以及成本参数。

### 运维说明

控制文件将扩展固定在 `pg_catalog` 中，并标记为不可迁移。它没有 GUC 或预加载要求，但 C 实现直接使用 PostgreSQL 私有的 autovacuum 结构和锁，因此必须针对精确的服务端构建验证兼容性。所有结果都应视为瞬时诊断状态，并严格遵守上游的非生产警告。
