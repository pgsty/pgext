## 用法

来源：

- [官方仓库](https://github.com/darthunix/x5fix)
- [目录源码版本的扩展控制文件](https://github.com/darthunix/x5fix/blob/3c28e84d98ec17a0e2cf0da7be5abb3b571334e1/x5fix.control)
- [0.1 版扩展 SQL](https://github.com/darthunix/x5fix/blob/3c28e84d98ec17a0e2cf0da7be5abb3b571334e1/x5fix--0.1.sql)
- [0.1 版 C 实现](https://github.com/darthunix/x5fix/blob/3c28e84d98ec17a0e2cf0da7be5abb3b571334e1/x5fix.c)

x5fix 是一个已废弃、侵入性很强的旧版 Greenplum 持久文件空间修复原型。它不是普通 PostgreSQL 扩展：唯一的 SQL 函数会进入 Greenplum 共享内存与持久目录存储，修改主副本和镜像的数据库标识与路径。

### 不要在在线集群中调用

目录源码依赖 Greenplum 私有头文件、锁、结构体、关系标识符和全局变量。这些接口与版本绑定，在普通 PostgreSQL 中也不存在。针对其他 Greenplum 源码树构建的二进制可能无法装载、导致后端崩溃、破坏共享状态或损坏持久文件空间元数据。

该函数会修改内部存储，却被声明为不可变函数，而且没有 SQL 层权限保护。这些声明与实际行为不相容。C 实现还会把主路径写入持久元组的两个路径字段，而没有使用独立的镜像路径参数；这是固定源码中的阻断性正确性缺陷。

### 非破坏性审计

对于继承的集群，只能先清点扩展与函数定义，再与匹配旧版 Greenplum 的专家规划移除或取证分析。

```sql
SELECT extversion
FROM pg_extension
WHERE extname = 'x5fix';

SELECT to_regprocedure(
    'add_entry_gp_persistent_filespace_node(oid,smallint,text,smallint,text)'
);
```

不要把调用返回的函数作为诊断手段。即使返回成功，也不能证明文件系统一致性、镜像正确性或目录修复的持久性。

### 恢复边界

应使用与厂商和版本匹配的 Greenplum 恢复流程、经过验证的备份和离线目录检查，而不是该扩展。如果旧部署已经包含它，应撤销所有应用角色的执行权限，保留准确的服务器二进制与源码版本用于取证，并在一次性副本上测试任何补救措施。

控制文件没有声明预加载要求，但函数会在调用时初始化并修改共享内存结构。这种设计进一步说明，若没有完整的专家级重写与故障注入测试计划，就不应移植或部署 0.1 版。
