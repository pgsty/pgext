## 用法

来源：

- [官方上游 README](https://github.com/dalibo/hackingpg/blob/a16eacb921fb5560a504b61c539d98b8919617f8/journee6/README)
- [官方扩展控制文件 (cpg.control)](https://github.com/dalibo/hackingpg/blob/a16eacb921fb5560a504b61c539d98b8919617f8/journee6/cpg.control)
- [官方实现源代码](https://github.com/dalibo/hackingpg/blob/a16eacb921fb5560a504b61c539d98b8919617f8/journee6/cpg.c)

`cpg` 是一个基于 Corosync 的 PostgreSQL 背景工作者实验。它交换主节点成员信息，更新 `primary_conninfo`，并可以请求 systemd 促进或停止节点；它是一个工作坊模块而非通用的高可用产品。

### 核心工作流

上游配置示例安装了该库并进行了如下配置：

```ini
shared_preload_libraries = 'cpg'
include_if_exists = 'cpg.auto.conf'
cpg.interval = 10
cpg.service = 'postgresql-16.service'
cpg.is_user_service = false
```

添加库后重启 PostgreSQL。工作者加入 Corosync 组，将发现的连接状态写入 `cpg.auto.conf`，并在故障转移路径中使用配置的 systemd 服务。

### 重要设置

- `cpg.interval` 是工作者唤醒间隔的最大时间（秒）。
- `cpg.service` 指定工作者可以管理的 systemd 服务名称。
- `cpg.is_user_service` 选择使用用户 bus 而不是系统 bus。

### 要求与注意事项

- 审查过的控制文件标识版本 `0.1`；没有扩展 SQL 或 `CREATE EXTENSION` 工作流程。
- 源代码需要 Corosync CPG 和 systemd 库。演示配置还安装了一个 polkit 规则，以便 PostgreSQL 账户可以管理其服务。
- 工作者可以修改复制配置并触发促进或关闭操作。在任何非实验室使用前，请测试隔离、失去多数、过时成员信息、systemd 权限以及分裂脑行为。
- 该代码来自一个 PostgreSQL 编程工作坊，应被视为教育原型。
