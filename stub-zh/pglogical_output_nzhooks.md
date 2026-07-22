## 用法

来源：

- [扩展控制文件](https://github.com/troels/pghooks/blob/f82688238e25bb220478404eeeb4d6a825f014a1/pglogical_output_nzhooks.control)
- [1.0 版本 SQL 对象](https://github.com/troels/pghooks/blob/f82688238e25bb220478404eeeb4d6a825f014a1/pglogical_output_nzhooks--1.0.sql)
- [钩子实现](https://github.com/troels/pghooks/blob/f82688238e25bb220478404eeeb4d6a825f014a1/pglogical_output_nzhooks.c)
- [构建集成](https://github.com/troels/pghooks/blob/f82688238e25bb220478404eeeb4d6a825f014a1/Makefile)

`pglogical_output_nzhooks` 是一个极简的 `pglogical_output` 行过滤钩子示例。它不是通用过滤扩展：已复核版本只输出关系名恰好为 `nzmagic` 的行，并对检查的每个关系记录警告。

### 集成边界

```sql
CREATE EXTENSION pglogical_output_nzhooks;
```

扩展安装 `pglo_nzhooks_setup_fn(internal)`。`internal` 参数是由匹配的 `pglogical_output` 钩子 ABI 提供的指针；用户不能从普通 SQL 安全调用此函数。应通过钩子机制配置兼容的 `pglogical_output` 构建来调用该 setup 符号。仓库没有记录稳定的终端用户配置语法，因此应从实际构建的 pglogical 版本确认，而不是猜测选项。

构建时，Makefile 期望在相邻路径 `../pglogical/pglogical_output` 中找到 pglogical 输出插件头文件。它没有声明 SQL 扩展依赖或版本检查。应从兼容源码构建两个组件，并测试其私有 `PGLogicalHooks` 与 `PGLogicalRowFilterArgs` 布局一致。

### 仅限示例的行为

setup 函数会以警告级别输出 `HELLO THERE!`。行过滤器会对每个变更关系记录名称警告，只比较未限定关系名与 `nzmagic`，并对其他所有表返回 false。模式名被忽略，因此不同模式中的同名表都会通过。

不要原样部署此示例：它可能静默丢弃几乎所有逻辑变更，并淹没服务器日志。应派生代码以实现明确的模式限定策略，删除逐行警告日志，为插入/更新/删除与分区添加回归测试，并验证复制槽重启、故障转移、二进制兼容性与升级行为。
