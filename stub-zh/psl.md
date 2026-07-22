## 用法

来源：

- [上游 README](https://github.com/wttw/pgpsl/blob/8def780438e263506a8ae4beb9c754a31f9ee341/README.md)
- [版本 1.0.0 SQL 定义](https://github.com/wttw/pgpsl/blob/8def780438e263506a8ae4beb9c754a31f9ee341/psl--1.0.0.sql)
- [扩展控制文件](https://github.com/wttw/pgpsl/blob/8def780438e263506a8ae4beb9c754a31f9ee341/psl.control)

软件包 `pgpsl` 安装扩展 `psl` 版本 `1.0.0`。它提供一个函数，使用编译进动态库的 Public Suffix List 快照查找主机名的可注册域名。

### 核心流程

```sql
CREATE EXTENSION psl;

SELECT registered_domain('foo.bar.blighty.com');
SELECT registered_domain('www.blighty.co.uk');
SELECT registered_domain('co.uk');
```

`registered_domain(text)` 会把结果转换为小写。对于主机名，它返回所属的可注册域；对于已经是注册域的输入，返回原值；对于公共后缀或不含点的名称，返回 NULL。若顶级后缀看似有效但不在列表中，则退回到最后两个组件。

### 数据新鲜度与注意事项

Public Suffix List 在构建时固定，不能动态刷新。后缀准确性重要时，必须使用经审核的列表快照重新构建扩展；仅变更软件包不会更新已经加载的动态库。

上游明确警告其列表预处理代码已损坏，且尝试修复曾产生损坏的内嵌列表和崩溃。使用前应测试代表性后缀、国际化名称、无效输入与未知后缀。不要把这个无人维护的函数作为 Cookie、租户隔离或授权的唯一安全边界；过期或损坏的列表可能误判域名所有权。
