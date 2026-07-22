## 用法

来源：

- [Pinned control file](https://github.com/NguyenLe1605/pgmolji/blob/42b9486e29319026df7c957b484eaf47fceabecd/pgmolji.control)
- [Pinned executor-hook implementation](https://github.com/NguyenLe1605/pgmolji/blob/42b9486e29319026df7c957b484eaf47fceabecd/src/lib.rs)

`pgmolji` 是实验性的 pgrx 执行器钩子扩展。在某个后端加载其共享库后，它会记录受支持顺序扫描计划的简化表示，并把返回的每个 `text` 或 `varchar` 字节替换为随机选择的 emoji。它会有意改变查询结果，只适合演示或隔离实验。

### 演示

请使用可丢弃的数据库和会话：

```sql
CREATE EXTENSION pgmolji;

CREATE TEMP TABLE demo (id integer, message text);
INSERT INTO demo VALUES (1, 'hello');

SELECT id, message FROM demo WHERE id = 1;
```

整数会原样传递，而 `message` 会替换成随机 emoji 输出。模块还会为这个简单顺序扫描计划输出一条 `[pgmolji]` 日志。

### 暴露对象与钩子行为

- `hello_pgmolji()` 返回文本问候语，但随后会被目标接收器转换。
- `_PG_init()` 在共享库载入某个后端时安装 `ExecutorRun` 钩子。
- 自定义目标接收器转换类型 OID 为 `text` 或 `varchar` 的输出属性；其他类型原样传递。
- 计划渲染只理解顶层顺序扫描和少量表达式节点/运算符。不支持的计划或表达式会记录为不支持/未知。

### 安全边界

钩子在后端中激活后，没有 GUC、角色过滤、关系过滤或受支持的退出机制。输出不确定，而且转换依据存储字节长度，而不是 Unicode 字符数。因此应用、测试、导出、监控查询和管理工具都可能收到被破坏的文本值。

实现版本为 `0.0.0`，安装需要超级用户，并且不可迁移。它包含低层执行器和元组目标代码，测试面很有限；空值、TOAST 值、嵌套计划、非顺序扫描以及变化的 PostgreSQL 内部接口都是显著风险。不要在集群范围预加载，也不要安装到生产数据库。仅在某个后端加载时不需要重启。
