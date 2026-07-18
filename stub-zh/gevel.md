## 用法

来源：

- [当前 HEAD 的上游 README](http://sigaev.ru/git/gitweb.cgi?p=gevel.git;a=blob;f=README.gevel;hb=cb9416a0fae4b90eaaefeaa85948c50b481dc2a0)
- [当前扩展控制文件](http://sigaev.ru/git/gitweb.cgi?p=gevel.git;a=blob;f=gevel.control;hb=cb9416a0fae4b90eaaefeaa85948c50b481dc2a0)
- [当前安装 SQL](http://sigaev.ru/git/gitweb.cgi?p=gevel.git;a=blob;f=gevel.sql;hb=cb9416a0fae4b90eaaefeaa85948c50b481dc2a0)
- [当前索引页实现](http://sigaev.ru/git/gitweb.cgi?p=gevel.git;a=blob;f=gevel.c;hb=cb9416a0fae4b90eaaefeaa85948c50b481dc2a0)
- [项目官方文档](http://www.sai.msu.su/~megera/oddmuse/index.cgi/Gevel)

gevel 1.1 是面向高级索引开发者的底层诊断模块。它遍历 GiST、GIN 与 SP-GiST 页面，报告树结构、页面统计、已存键值和 GIN 匹配数量估计。当前源码也包含 B-tree 与 BRIN 诊断，但其构建与安装条件在不同 PostgreSQL 版本间并不一致。

### 诊断示例

如果所用发行包提供了可工作的版本化扩展脚本，可在测试索引上执行：

```sql
CREATE EXTENSION gevel;

SELECT gist_stat('public.demo_gist_idx');
SELECT gist_tree('public.demo_gist_idx', 1);

SELECT gin_count_estimate(
    'public.demo_gin_idx',
    'postgres'::tsquery
);
```

各类打印函数返回 setof record，调用方必须按索引操作符类提供精确的输出列类型。上游 README 给出了 GiST、GIN、SP-GiST、B-tree 与 BRIN 的具体例子。

### 锁、权限与打包风险

该代码通过 PostgreSQL 私有访问方法结构读取原始索引页。若干 GiST、B-tree 与 BRIN 路径会取得 AccessExclusiveLock 并递归扫描整棵索引；大型或损坏索引可能长时间阻塞业务，甚至使后端崩溃。打印出的索引键还可能泄露被索引数据。SQL 默认允许公共执行，而 C 路径没有实现类似表 SELECT 的权限边界。应撤销 PUBLIC 对全部函数的执行权，只向诊断角色开放经过审查的窄封装。

当前 HEAD 停留在 2020 年。Makefile 安装旧式 gevel.sql，且未声明常规的 EXTENSION 与版本化 SQL PGXS 变量，尽管 gevel.control 声明版本 1.1；B-tree 与 BRIN SQL 还只在 PostgreSQL 12 条件下安装。发行包可能做过补丁，因此必须核对实际安装文件与函数清单，不能假定上游 CREATE EXTENSION 原样可用。

只能在副本或可丢弃副本上使用，按精确服务器源码编译，设置语句与锁超时，先测试一个小索引，并且不要把页面输出当作稳定支持的 API。
