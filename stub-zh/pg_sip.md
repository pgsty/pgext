## 用法

来源：

- [扩展 SQL](https://github.com/quentusrex/pg_sip/blob/master/pg_sip--0.0.1.sql)
- [版本函数实现](https://github.com/quentusrex/pg_sip/blob/master/src/pg_sip_version.c)
- [SIP 头部原型](https://github.com/quentusrex/pg_sip/blob/master/src/pg_sip_header.c)
- [官方仓库](https://github.com/quentusrex/pg_sip)

`pg_sip` 0.0.1 是一个把 SIP 消息解析能力暴露给 PostgreSQL 的未完成原型。扩展实际安装的 SQL 接口只有版本函数；源码树虽有提取头部的 C 函数，但扩展安装脚本没有声明它。

### 已安装工作流

```sql
CREATE EXTENSION pg_sip;

SELECT pg_sip_version();
-- 0.0.1

SELECT p.oid::regprocedure
FROM pg_proc AS p
JOIN pg_depend AS d
  ON d.classid = 'pg_proc'::regclass
 AND d.objid = p.oid
JOIN pg_extension AS e
  ON e.oid = d.refobjid
WHERE e.extname = 'pg_sip';
```

### API 边界

- `pg_sip_version()` 为 immutable 和 strict，返回扩展版本字符串。
- 源码中的首个匹配 SIP 头读取函数只在回归脚本里被手工声明后测试；它不属于 CREATE EXTENSION，因此不是受支持的已安装 API。
- 校验和更多 SIP 访问器等路线图条目不是已经实现的接口。

### 构建与安全说明

- 解析器原型通过仓库子模块链接捆绑的 libre 库；未初始化子模块的源码归档并不完整。
- SIP 消息属于不可信网络输入。若没有模糊测试、内存安全审查、长度限制和目标大版本测试，不要在生产中手工暴露未声明的 C 符号。
- 每次构建后都要检查打包的扩展 SQL；编译产物中存在某个符号，并不意味着它是扩展管理的数据库对象。
