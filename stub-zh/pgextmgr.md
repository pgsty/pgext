## 用法

来源：

- [官方仓库 README](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/README.md)
- [扩展控制文件](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgextmgr/pgextmgr.control)
- [管理器实现](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgextmgr/src/lib.rs)

`pgextmgr` 是一个实验性的 pgrx hook 管理器，允许经过专门插桩的 PostgreSQL 扩展插件注册有序的规划器、执行器与查询重写回调。它可以检查这些注册项，并在运行时启用或禁用插件。这是面向兼容插件的原型框架，并非任意已安装扩展的通用管理器。

### 核心流程

上游测试配置会预加载管理器及经过插桩的插件，使其 hook 在服务器启动期间完成安装：

```conf
shared_preload_libraries = 'pgextmgr,pgext_pg_stat_statements,pgext_pg_hint_plan,pgext_pg_poop'
```

修改 `shared_preload_libraries` 后重启 PostgreSQL，再创建 `pgextmgr` 以及针对其 ABI 构建的插件。更改状态前先检查插件状态和 hook 顺序：

```sql
CREATE EXTENSION pgextmgr;

SELECT * FROM pgextmgr.all();
SELECT * FROM pgextmgr.hooks();

SELECT pgextmgr.disable('pgext_pg_poop');
SELECT pgextmgr.enable('pgext_pg_poop');
```

`pgextmgr.all()` 报告每个已注册插件的顺序及启用/禁用状态。`pgextmgr.hooks()` 报告各受支持 hook 的有序插件链。修改状态的函数返回受影响的注册项数量。

### 函数索引

- `pgextmgr.all()` 列出已注册插件及其当前状态。
- `pgextmgr.hooks()` 列出管理器中的规划器、执行器与重写器回调顺序。
- `pgextmgr.enable(text)` 与 `pgextmgr.disable(text)` 修改一个已知插件的状态。
- `pgextmgr.enable_all()` 与 `pgextmgr.disable_all()` 修改全部已注册插件的状态。

### 运维说明

`pgextmgr` 安装到固定的 `pgextmgr` 模式，创建时需要超级用户权限。审阅的源码面向 PostgreSQL 15 时期的 pgrx 环境。其插件快速入门使用定制链接参数以及框架的 `__pgext_before_init`、`__pgext_after_init` 符号；未针对该 ABI 构建的普通扩展不会成为可管理插件。

启用/禁用状态保存在后端进程内存中，而不是持久目录。应确认重新连接和重启后的行为，并使用实际插件集合测试 hook 顺序与故障处理。上游 README 将该框架描述为尚在开发中的工作，因此它更适合受控研究，不能假定其具备生产兼容性或稳定 API。
