## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/zxdvd/pghelper/blob/04e91fb5722396cfe772cee9be9edc07a67f035f/README.md)
- [扩展 SQL 目录](https://github.com/zxdvd/pghelper/tree/04e91fb5722396cfe772cee9be9edc07a67f035f/sql)
- [扩展 control 文件](https://github.com/zxdvd/pghelper/blob/04e91fb5722396cfe772cee9be9edc07a67f035f/myhelper.control)

`myhelper` 是收集在 `helper` 模式中的个人纯 SQL 工具箱。已复核目录包含数组辅助函数、安全转换、URL 解码、日期工具、权限检查、表描述、锁与膨胀报告，以及表切换脚本。

```sql
CREATE EXTENSION myhelper;
SELECT helper.array_uniq(ARRAY[3, 1, 3, 2]);
SELECT * FROM helper.unnest_with_index(ARRAY['a', 'b']);
SELECT helper.get_days_of_month(current_timestamp);
```

示例只能用于发现，应通过 `\df helper.*` 检查已安装签名。这是异构个人工具箱，不是稳定版本化公共 API；管理与表切换模块可能检查或修改数据库对象，并可能使用动态 SQL。

应先在可丢弃数据库安装，对比固定 SQL 中的每个创建对象，从 `PUBLIC` 撤销执行权，只逐项授予只读辅助函数。应测试搜索路径、标识符引用、空值与异常、权限检查、并发 DDL 和版本兼容性。膨胀、锁、权限和在线表变更应优先使用持续维护的专用工具；未经逐语句审计，不要把组合 SQL 导入托管服务。
