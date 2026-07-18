## 用法

来源：

- [上游安装与基本用法](https://github.com/no0p/plotpg/blob/20b7b3515c0de5c214d1ebb32c9063d9724386b0/README.md)
- [查询绘图实现](https://github.com/no0p/plotpg/blob/20b7b3515c0de5c214d1ebb32c9063d9724386b0/plotpg.c)
- [直接调用 gnuplot 的实现](https://github.com/no0p/plotpg/blob/20b7b3515c0de5c214d1ebb32c9063d9724386b0/gnuplot.c)

`plotpg` 通过服务器主机上的 `gnuplot` 可执行程序渲染 PostgreSQL 查询结果。`plot` 函数经 SPI 执行传入查询，写入临时数据和脚本，调用外部程序并返回输出。

```sql
CREATE EXTENSION plotpg;
SELECT plot('SELECT price, weight FROM widgets');
```

该实现会执行调用者提供的 SQL 与 gnuplot 命令，并在 `/tmp` 下使用可预测的进程名文件。应把这些函数视为操作系统命令能力：撤销 `PUBLIC` 的执行权限，只授予可信管理员，限制数据库与操作系统权限，绝不能传入用户输入。该项目是较早的实验，建议改在应用层或受维护的可视化服务中渲染查询结果。
