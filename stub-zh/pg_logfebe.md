## 用法

来源：

- [官方协议与配置文档](https://github.com/logplex/pg_logfebe/blob/bac15ef2386fc70197ba889e23027bffe00d0f5b/README.rst)
- [扩展 control 文件](https://github.com/logplex/pg_logfebe/blob/bac15ef2386fc70197ba889e23027bffe00d0f5b/pg_logfebe.control)

`pg_logfebe` 是一个仅供预加载的日志模块，它使用 PostgreSQL 前端/后端帧格式，通过 Unix 域套接字把结构化日志记录发送给另一个进程。它没有扩展 SQL 对象，因此不要为它执行 `CREATE EXTENSION`。

### 核心流程

先在受保护的 Unix 套接字上运行兼容接收器，再配置 PostgreSQL 并重启服务器以加载模块。

```conf
shared_preload_libraries = 'pg_logfebe'
logfebe.unix_socket = '/run/pg_logfebe/log.sock'
logfebe.identity = 'primary-db'
```

接收器首先收到版本与身份消息，随后持续收到结构化日志记录。身份值是运维人员自定义的标签，使同一个接收器能够区分不同 PostgreSQL 实例。

### 协议与运维边界

- `logfebe.unix_socket`：指定接收端 Unix 域套接字。
- `logfebe.identity`：提供启动阶段发送的实例标签。
- `shared_preload_libraries`：必需，因为模块会在服务器启动时安装日志钩子。

其帧格式类似 PostgreSQL FE/BE 协议，但文档描述的日志流是独立的版本化协议，需要专用接收器。应保护套接字目录和下游日志，因为结构化记录可能包含查询文本、错误详情等敏感数据库信息。修改预加载列表需要重启；它没有按数据库执行的 DDL 或 SQL API。
