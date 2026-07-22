## 用法

来源：

- [官方 README](https://github.com/extio1/pg_isolate/blob/main/README.md)
- [Postmaster 钩子实现](https://github.com/extio1/pg_isolate/blob/main/src/main.c)
- [配置解析器](https://github.com/extio1/pg_isolate/blob/main/src/config/group_config.c)
- [Setuid 辅助程序实现](https://github.com/extio1/pg_isolate/blob/main/src/cgroup_initialize.c)

`pg_isolate` 1.0 是仅限 Linux 的预加载原型，按数据库和用户规则把 PostgreSQL 后端移入 cgroup v2 目录。它不安装 SQL 扩展对象。当前实现要求 setuid-root 辅助程序，并包含不安全、绑定特定主机的代码路径；未经完整安全重构，不应部署到生产服务器。

### 预期配置

模块读取相对路径下的 `pg_isolate.json`，选择第一个匹配数据库和用户的组，并可配置默认组：

```json
{
  "reporting": {
    "database": "analytics",
    "user": "reporter"
  },
  "default": {}
}
```

模块在 postmaster 启动时加载：

```conf
shared_preload_libraries = 'pg_isolate'
```

不存在 `CREATE EXTENSION` 步骤。上游构建目标会把 `cgroup_initialize` 编译到 PostgreSQL 二进制目录，将所有者改为 root，并在服务器启动前设置 setuid 位。

### 实际行为

启动时，辅助程序创建硬编码的 `/sys/fs/cgroup/postgres` 层级，启用内存和 CPU 控制器，并把 postmaster 移入实例子树。认证钩子随后移动匹配的后端并记录连接元数据。虽然 JSON 解析器接受 `memory_hard`、`memory_soft` 和 `cpu_weight`，但已检查代码只创建目录，并不会应用这些资源值。

### 严重风险

解析器会把无长度限制的 JSON 字符串复制到很小的固定缓冲区。辅助程序在开发者专用主目录下写调试文件，后台工作进程钩子还包含硬编码的 `test_module` 分支：它会取消共享挂载命名空间，并尝试对 `/home/extio1/alpine` 执行 `pivot_root`。虽然 README 写的是 PGDATA，但配置查找实际相对于 postmaster 工作目录。这些行为，加上 setuid 辅助程序、对版本敏感的钩子、cgroup v2 假设和详细认证日志，使当前源码不适合生产隔离。
