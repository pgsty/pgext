## 用法

来源：

- [已复核 commit 的 x5fix control 文件](https://github.com/darthunix/x5fix/blob/3c28e84d98ec17a0e2cf0da7be5abb3b571334e1/x5fix.control)
- [已复核 commit 的 x5fix 0.1 安装 SQL](https://github.com/darthunix/x5fix/blob/3c28e84d98ec17a0e2cf0da7be5abb3b571334e1/x5fix--0.1.sql)
- [已复核 commit 的 x5fix C 实现](https://github.com/darthunix/x5fix/blob/3c28e84d98ec17a0e2cf0da7be5abb3b571334e1/x5fix.c)

`x5fix` 0.1 是已废弃且无文档的旧 Greenplum 持久文件空间修复代码。它只暴露 `add_entry_gp_persistent_filespace_node(oid, smallint, text, smallint, text)`。该函数没有安全的通用调用方式；在完全匹配的取证环境安装后，也只应检查登记的签名：

```sql
CREATE EXTENSION x5fix;

SELECT 'add_entry_gp_persistent_filespace_node(oid,smallint,text,smallint,text)'::regprocedure;
```

### 注意事项

- 不要从普通维护手册调用修复函数。它会获取 Greenplum 内部锁，并修改持久目录及共享内存状态。
- 安装 SQL 把这个有修改行为的函数声明为 `IMMUTABLE`，该标记不正确，可能让规划器作出不安全假设。
- C 源码在替换元组时禁用 WAL 刷新，而且把主路径写入镜像路径字段，而不是使用传入的镜像路径，表明存在严重实现缺陷。
- 仓库没有 README、测试、许可证、发布或受支持版本矩阵；它依赖旧 Greenplum 私有头文件，不能当作 PostgreSQL 兼容扩展。
- 任何恢复用途都要求二进制与准确的 Greenplum 服务器构建匹配，并具备已验证备份、离线复现以及供应商资质的源码复核。
