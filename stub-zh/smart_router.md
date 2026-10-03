## 用法

来源：

- [官方 setup.sh](https://github.com/mradul-001/dbisFinal/blob/b1aa6af716ce5c906c0d18766258277d2a64fd64/setup.sh)
- [官方 smart_router.control](https://github.com/mradul-001/dbisFinal/blob/b1aa6af716ce5c906c0d18766258277d2a64fd64/dbis-project/contrib/smart_router/smart_router.control)
- [官方 Makefile](https://github.com/mradul-001/dbisFinal/blob/b1aa6af716ce5c906c0d18766258277d2a64fd64/dbis-project/contrib/smart_router/Makefile)
- [官方 smart_router.c](https://github.com/mradul-001/dbisFinal/blob/b1aa6af716ce5c906c0d18766258277d2a64fd64/dbis-project/contrib/smart_router/smart_router.c)
- [官方 smart_router--1.0.sql](https://github.com/mradul-001/dbisFinal/blob/b1aa6af716ce5c906c0d18766258277d2a64fd64/dbis-project/contrib/smart_router/smart_router--1.0.sql)

`smart_router` 1.0 是面向固定双集群缓存演示的研究性预加载模块，在该仓库自带的 PostgreSQL 源码树中构建。安装 SQL 不创建对象；官方配置通过预加载激活模块，没有数据库级扩展创建步骤。

### 启用边界

```conf
shared_preload_libraries = 'smart_router'
```

重启后会注册执行器钩子与模式同步后台进程。上游脚本在端口 6000 启动远端集群及 `company_remote` 数据库，在端口 6001 启动本地集群。源码将远端登录角色硬编码为 `mradul`，后台进程的本地数据库固定为 `postgres`；这些是演示环境假设，并非可配置的路由 API。

### 行为与限制

模块通过 libpq 获取远端表定义与数据，填充本地表，并用 LRU 策略跟踪五个表条目。执行器钩子检查查询文本，尝试使用远端或缓存处理。上游没有提供调整固定连接和缓存布局的管理函数或 GUC。

仅用于可丢弃的演示集群和可信输入。淘汰过程会对本地表执行 DELETE，远端值未经转义就拼入 SQL，查询识别器也不是通用 SQL 解析器。不能让模块操作本地权威数据表，不能假定其具备透明路由、事务一致性或安全的多用户行为。该模块未声明许可或原版 PostgreSQL 版本兼容范围。
