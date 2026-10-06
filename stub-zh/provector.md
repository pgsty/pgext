## 用法

来源：

- [README.md](https://github.com/grussdorian/provector/blob/4b20fa81fbf2ca24c720542640158b7d995396f2/README.md)
- [provector--1.0.sql](https://github.com/grussdorian/provector/blob/4b20fa81fbf2ca24c720542640158b7d995396f2/provector--1.0.sql)
- [test/test.sql](https://github.com/grussdorian/provector/blob/4b20fa81fbf2ca24c720542640158b7d995396f2/test/test.sql)
- [provector.cpp](https://github.com/grussdorian/provector/blob/4b20fa81fbf2ca24c720542640158b7d995396f2/provector.cpp)
- [provector.control](https://github.com/grussdorian/provector/blob/4b20fa81fbf2ca24c720542640158b7d995396f2/provector.control)

`provector` 1.0 是早期 C++ 扩展示例。核对的 SQL 接口只有一个文本转换函数；项目名及 pgvector 依赖并不能证明其实现了向量搜索。

### 核心用法

```sql
CREATE EXTENSION vector;
CREATE EXTENSION provector;
SELECT provector_demo('a title');
```

### 运行边界

`provector_demo` 接收文本，转换为标题式大小写并返回文本，声明为 strict 和 volatile。控制文件依赖 `vector`，但这个示例函数不接收向量参数。

安装库及依赖后由超级用户创建扩展。核对的源码不要求预加载或重启。构建链接区域设置及 gettext 设施，非 ASCII 输入应在目标区域设置下验证。上游未发布 PostgreSQL 主版本兼容矩阵，应将其视为教学原型而非生产搜索组件。
