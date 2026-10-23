## 用法

来源：

- [Official source](https://github.com/soliyaanurag/Bloom-Filter-based-semi-join/blob/009b0dca08d240b773efeb92c2b57e3d4d8f88e7/README.md)
- [Official source](https://github.com/soliyaanurag/Bloom-Filter-based-semi-join/blob/009b0dca08d240b773efeb92c2b57e3d4d8f88e7/sjbloom/sjbloom.control)
- [Official source](https://github.com/soliyaanurag/Bloom-Filter-based-semi-join/blob/009b0dca08d240b773efeb92c2b57e3d4d8f88e7/sjbloom/sjbloom--1.0.sql)

`sjbloom` 1.0 为半连接过滤研究实现提供布隆、精确键集合和直接通过三种过滤器。随附的自动规划集成修改了 PostgreSQL 18.6 的 `postgres_fdw`；仅安装本扩展并不会启用该规划行为。

### 过滤函数

```sql
CREATE EXTENSION sjbloom;
SELECT sjb_contains(sjb_build(ARRAY[1,2,3], 1024, 3), 2);
SELECT sjb_info(sjb_build(ARRAY[1,2,3], 1024, 3));
```

`sjb_build` 接收整数键、位数和哈希次数。`sjb_contains` 测试整数是否通过字节数组过滤器。布隆匹配可能是假阳性，最终必须通过精确连接排除；NULL 输入仍返回 NULL。

### 自动过滤与限制

`sjb_build_auto` 根据基数、行宽、网络速度和标定成本选择过滤器。`sjb_choice_report` 展示模型的候选方案，`sjb_best_bits` 估算过滤器大小。默认成本来自上游测试机器，不能视为普遍的性能保证。

分布式使用时，两端需安装匹配函数并遵循修改后的 FDW 配置。独立函数不要求预加载。扩展可迁移模式，安装需要超级用户。这是课程原型，未验证跨当前主版本的可移植性或生产成熟度，仓库也没有可核实的许可文件。
