## 用法

来源：

- [目录源码版本的官方 README](https://github.com/rjuju/pg_dropbuffers/blob/8328aed1a7ffc3c1a322fa2decf1be7c63024842/README.md)
- [目录源码版本的扩展控制文件](https://github.com/rjuju/pg_dropbuffers/blob/8328aed1a7ffc3c1a322fa2decf1be7c63024842/pg_dropbuffers.control)
- [0.0.1 版扩展 SQL](https://github.com/rjuju/pg_dropbuffers/blob/8328aed1a7ffc3c1a322fa2decf1be7c63024842/pg_dropbuffers--0.0.1.sql)
- [0.0.1 版 C 实现](https://github.com/rjuju/pg_dropbuffers/blob/8328aed1a7ffc3c1a322fa2decf1be7c63024842/pg_dropbuffers.c)

pg_dropbuffers 是破坏性的基准测试辅助工具，可驱逐当前数据库的缓冲区，还能请求操作系统丢弃页缓存。上游明确限定它只用于测试，并警告可能丢失数据和严重影响性能；不要在生产系统中安装或运行。

### 隔离测试流程

只能在一次性集群上使用，并应先停止应用流量、保存所有重要数据。每次只处理一层缓存，才能在基准结果中明确记录发生了什么变化。

```sql
CREATE EXTENSION pg_dropbuffers;

SELECT pg_drop_current_db_buffers();
SELECT pg_drop_system_cache();
SELECT pg_drop_caches();
```

第一个函数刷新并丢弃属于当前数据库的共享缓冲区。第二个函数调用 C 源码中配置的操作系统命令。包装函数会按这个顺序调用两者。

### 操作系统要求

0.0.1 版的系统缓存驱逐功能只适用于 Linux：它会执行一个固定的 sudo 命令来设置内核缓存。PostgreSQL 服务账号必须被允许无密码执行这条准确命令，否则函数会报错。这是影响其他数据库与进程的主机级特权操作，并非只影响当前会话。

### 安全边界

上游 README 警告，删除前可能无法安全保留脏缓冲区。每次调用都应视为可能丢失数据或破坏测试集群稳定性。绝不能向不可信 SQL 用户暴露这些函数；如果确实必须保留扩展，应撤销默认执行权限，只授权专用基准测试操作员。

缓存驱逐本身不能保证可重复的冷缓存基准：并发活动可能立即重新填充缓存，主机级缓存调用也会扰动无关工作负载。每个结果都应记录集群状态、调用顺序、工作负载时序和操作系统条件。

扩展没有声明预加载要求，但使用 PostgreSQL 内部缓冲区接口和硬编码主机命令。在任何一次性测试之前，都必须确认源码与二进制和准确的服务器、操作系统兼容。
