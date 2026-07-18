## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/README.md)
- [Rust 实现](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/src/lib.rs)
- [安装时 SQL 脚本](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/sql/symbiotic_python.sql)
- [扩展 control 文件](https://github.com/MarchLiu/symbiotic_python/blob/519af89947c885de1f445d40cd2e5b0d5a5bf9dc/symbiotic_python.control)

`symbiotic_python` 是一个已废弃、只开发了一天的 pgrx 原型，不是可以安全安装的扩展。其安装脚本会立即调用 `create_venv`，在 PostgreSQL 服务账号的主目录下创建文件，通过 `pip` 下载 `asyncpg` 与 `asyncpg-listen`，写入启动脚本，并启动连接当前数据库的后台 Python 监听器。

### 严重安全警告

公开 SQL 函数把调用者提供的环境名称直接插入 `create_venv` 和 `drop_venv` 使用的 `sh -c` 命令，且没有 shell 引号。能够调用这些函数的用户可注入 shell 语法，以 PostgreSQL 服务账号身份执行操作系统命令；`drop_venv` 还会构造未加引号的 `rm -rf` 命令。

不要执行 `CREATE EXTENSION symbiotic_python`，不要暴露其共享库，也不要在包含数据或凭据的主机上测试。唯一适当用途是在隔离构建环境中检查源码。仓库没有运维文档、持续维护记录，也没有缓解这些安装时副作用的措施。
