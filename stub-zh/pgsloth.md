## 用法

来源：

- [官方 README](https://github.com/jamessewell/pgsloth/blob/cf5fbf1f68a7527bbe75c3eaf4c28bad92a582a0/README.md)
- [执行器钩子实现](https://github.com/jamessewell/pgsloth/blob/cf5fbf1f68a7527bbe75c3eaf4c28bad92a582a0/src/lib.rs)
- [扩展控制文件](https://github.com/jamessewell/pgsloth/blob/cf5fbf1f68a7527bbe75c3eaf4c28bad92a582a0/pgsloth.control)

`pgsloth` 版本 `0.0.0` 是一个演示模块，会在每次查询执行器启动时故意随机延迟 100 到 9,999 毫秒。它是用于观察延迟行为的测试玩具，不是生产扩展。

### 启用方式

该库只安装执行器钩子，没有 SQL 对象，因此不要用 `CREATE EXTENSION` 激活它。把库加入服务器配置并重启 PostgreSQL：

```conf
shared_preload_libraries = 'pgsloth'
```

重启后，普通查询会先输出通知并休眠：

```sql
SELECT 1;
```

### 运维边界

每个加载该钩子的后端都会受到全局延迟。它没有 GUC、角色豁免、关系过滤或会话开关。README 中的宽松复制和索引功能明确标记为“待实现”；当前源码只实现查询休眠。

要禁用它，应从 `shared_preload_libraries` 删除 `pgsloth` 并重启。绝不要在生产或共享服务器上启用：它会给应用查询与管理工作增加不可预测的延迟，并可能放大连接积压与超时故障。
