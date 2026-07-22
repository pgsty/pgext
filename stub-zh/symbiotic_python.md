## 用法

来源：

- [已复核 commit 的官方 README](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/README.md)
- [安装时 SQL 脚本](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/sql/symbiotic_python.sql)
- [Rust 实现](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/src/lib.rs)
- [Python 通知监听器](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/scripts/symbiotic.py)
- [构建元数据](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/Cargo.toml)

`symbiotic_python` 是一个已废弃的 pgrx 原型，尝试为每个数据库创建 Python 虚拟环境并启动异步通知监听器。它不能安全安装：安装扩展本身就会以 PostgreSQL 服务账户执行操作系统、文件系统、软件下载与进程启动等副作用。

### 预期流程

SQL API 包括 `create_venv(name)`、`drop_venv(name)`、`venv_path(name)` 和 `run_symbiotic(venv, channel)`，另有无参数包装器。安装脚本会自动为当前数据库调用环境创建。它在服务账户的 `$HOME/.symbiotic` 下创建文件，用 `pip` 安装 `asyncpg` 与 `asyncpg-listen`，写入 shell 启动器并启动分离的监听进程。

监听器重新连接本地数据库、订阅通知通道、把载荷追加到日志，并在收到关闭载荷时退出。扩展没有可靠的 SQL 状态或停止接口、包锁定文件、进程监管机制，也没有用于这些外部副作用的清理事务。

### 关键安全边界

`create_venv(name)` 和 `drop_venv(name)` 把调用者控制的名称直接插入 shell 命令，且未加引用。这允许攻击者以 PostgreSQL 操作系统账户注入 shell 命令；删除路径还会构造未引用的 `rm -rf`。把函数放在 `symbiotic` 模式并不能使其安全，默认函数权限也可能让暴露范围大于预期。

不要在任何包含数据、凭据或网络访问的主机上安装或加载 `symbiotic_python`。即使试用也不要执行其 SQL。唯一适当的评估方式是离线源码检查，或在无秘密且无法访问其他系统的一次性沙箱中检查。其旧 pgrx 依赖只覆盖历史 PostgreSQL 大版本，无法缓解设计层面的命令执行缺陷。
