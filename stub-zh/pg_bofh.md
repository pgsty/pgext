## 用法

来源：

- [上游 README](https://github.com/rdunklau/pg_bofh/blob/master/README.md)
- [扩展 control 文件](https://github.com/rdunklau/pg_bofh/blob/master/pg_bofh.control)

`pg_bofh` 是一个小型 planner hook 示例，用于拒绝不带 `WHERE` 子句的查询。其上游版本为 `0.0.1`。

这是一个 headless 模块：仓库中没有带版本号的扩展 SQL，因此不存在 SQL 激活步骤。编译并安装动态库后，将其加入服务器配置并重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_bofh'
```

模块加载后，不带 `WHERE` 子句的语句会报错而不执行。该全局规则影响的不只是误操作的全表变更，因此应把它视为 planner hook 演示，而不是完整的生产级授权或查询安全系统。在集群中启用前，应先测试应用与维护工作负载。
