## 用法

来源：

- [README](https://github.com/danolivo/pg_prosupport/blob/fb410817eb4c6315572e873f6950ed62570661fd/README.md)
- [Control file / 控制文件](https://github.com/danolivo/pg_prosupport/blob/fb410817eb4c6315572e873f6950ed62570661fd/pg_prosupport.control)
- [pg_prosupport--1.0.sql](https://github.com/danolivo/pg_prosupport/blob/fb410817eb4c6315572e873f6950ed62570661fd/pg_prosupport--1.0.sql)

`pg_prosupport` 利用规划器支持函数重写符合条件的 numeric 聚合。当前 1.0 源码可用于标准 PostgreSQL 19；PostgreSQL 18 需要该项目额外提供的内核补丁，并显式加载共享库。

### PostgreSQL 19 工作流

```sql
CREATE EXTENSION pg_prosupport;
CREATE TABLE amounts (amount numeric(12,2));
INSERT INTO amounts VALUES (12.50), (8.25);
EXPLAIN (VERBOSE, COSTS OFF) SELECT sum(amount), avg(amount) FROM amounts;
```

### 重写与设置

安装会为内置 numeric 聚合绑定支持函数，并在固定的 `prosupport` 模式创建对象。符合条件的有界 numeric 表达式可使用专门的求和和平均状态；不支持的精度或表达式仍走普通路径，这并不是普遍加速保证。

`pg_prosupport.bounded_numeric_agg` 默认启用有界重写，`pg_prosupport.fold_const_sum` 控制常量求和重写且默认关闭。PostgreSQL 19 按需加载共享库；带补丁的 PostgreSQL 18 则需要 `LOAD` 或预加载，修改启动预加载配置后需要重启。

### 维护边界

安装和绑定支持函数需要超级用户权限。恢复后应按上游说明调用 `prosupport.pps_attach_support()` 重新绑定；改变规划行为时应重新连接或清除缓存计划。在 PostgreSQL 19 上删除扩展前，必须先解除绑定，以免留下悬空的目录引用：

```sql
SELECT prosupport.pps_detach_support();
DROP EXTENSION pg_prosupport;
```

该项目继承了早期 pg_numeric_agg_support 仓库的工作。应将其作为规划器实验扩展，并针对准确的服务器构建验证。
