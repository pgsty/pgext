## 用法

来源：

- [官方 README](https://github.com/tvondra/empty/blob/ac8d6da58df3763a462a5fdd1f781269cff7d67d/README.md)
- [官方扩展控制文件](https://github.com/tvondra/empty/blob/ac8d6da58df3763a462a5fdd1f781269cff7d67d/empty.control)
- [官方扩展 SQL](https://github.com/tvondra/empty/blob/ac8d6da58df3763a462a5fdd1f781269cff7d67d/empty--1.0.sql)
- [官方 C 实现](https://github.com/tvondra/empty/blob/ac8d6da58df3763a462a5fdd1f781269cff7d67d/empty.c)

`empty` 1.0 是已归档的 PostgreSQL 扩展开发沙箱，而不是一个完整一致的用户功能。它汇集了 C 函数、固定 2×2 `matrix` 类型、示例 FDW、逻辑解码回调、表扫描、共享内存及错误日志 hook 等演示代码，只适合源码学习或受控实验。

### 最小实验

库会在 `_PG_init` 中申请共享内存和 LWLock tranche，因此必须通过 `shared_preload_libraries` 加载并重启 PostgreSQL，然后再创建扩展。下面的独立小示例会运行矩阵代码：

```conf
shared_preload_libraries = 'empty'
```

```sql
CREATE EXTENSION empty;

SELECT '(1 2 3 4)'::matrix + '(5 6 7 8)'::matrix;
SELECT '(1 2 3 4)'::matrix * '(5 6 7 8)'::matrix;

SELECT *
FROM matrix_powers('(1 1 1 0)'::matrix, 4);
```

解析器只接受上述括号格式中的四个整数。`matrix_powers(matrix, count)` 会输出连续幂及从零开始的序号。

### 演示接口

- 算术示例：`int4_plus`、`numeric_plus` 和 `random_string`。
- 矩阵示例：`matrix`、运算符 `+` 与 `*`，以及 `matrix_powers`。
- 检查示例：`read_table(regclass)` 扫描表并把每个值写入服务器 warning，而不是返回结果行。
- FDW 示例：`empty_fdw` 包含没有稳定文档契约的实验性文件/扫描代码。
- 同一个共享库还导出逻辑解码回调，并通过全局 emit-log hook 统计部分服务器消息级别。

这些部分是互不相关的教学实验；极简 README 没有定义受支持行为或兼容策略。

### 安全与维护边界

`read_table` 会通过服务器日志泄漏完整行内容；预加载后，全局 hook 会影响整个实例。应限制安装和执行权限，绝不能让示例读取敏感数据，并测试它与其他预加载模块的 hook 交互。

项目在 2019 年提交后归档，并使用会随 PostgreSQL 版本变化的服务器内部接口。部分声明和实现选择只是演示，并非生产级设计。不要把 `empty` 当成受支持的 FDW、监控扩展、随机生成器、矩阵库或解码插件；只复制并审查确实需要的具体模式。
