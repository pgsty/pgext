## 用法

来源：

- [扩展 control 文件](https://github.com/LukMRVC/tree_similarity_extension/blob/4b5b0753dd00504d7df3cbca8e22b67b55017fa6/tree_similarity_extension.control)
- [Rust 软件包元数据](https://github.com/LukMRVC/tree_similarity_extension/blob/4b5b0753dd00504d7df3cbca8e22b67b55017fa6/Cargo.toml)
- [扩展实现与测试](https://github.com/LukMRVC/tree_similarity_extension/blob/4b5b0753dd00504d7df3cbca8e22b67b55017fa6/src/lib.rs)

尽管名字如此，`tree_similarity_extension` 0.0.0 并没有实现树相似度。已审查源码只是最小化的 pgrx 演示，安装一个问候函数及其对应的单元测试。

### 调用已实现的函数

```sql
CREATE EXTENSION tree_similarity_extension;

SELECT hello_tree_similarity_extension();
```

该函数返回固定文本 `Hello, from my own extension 2!`。仓库没有定义树类型、解析器、距离或相似度函数、运算符、索引、表或配置参数。

### 原型边界

应用需要树相似度算法时，不要仅根据目录名称选择此扩展。应把 0.0.0 版视为脚手架，并在形成依赖前核实实际安装的对象列表。

软件包固定使用 pgrx 0.11.4，默认启用 PostgreSQL 13 构建 feature，声明了 PostgreSQL 11 至 16 的 feature，但没有提供超出这些构建声明的兼容性承诺。control 文件将安装标为需要超级用户且不可迁移。仓库没有 README、release、SQL 升级路径或使用契约；只能把它用作已审查的演示源码，不能作为生产相似度服务。
