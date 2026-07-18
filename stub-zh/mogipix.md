## 用法

来源：

- [固定提交的上游 README](https://gitlab.com/marjorielptt/mogipix/-/blob/6e134fe704b3a94a833d9ead1b2ae540b305a8d1/README.md)
- [固定提交的 Cargo 元数据](https://gitlab.com/marjorielptt/mogipix/-/blob/6e134fe704b3a94a833d9ead1b2ae540b305a8d1/Cargo.toml)
- [固定提交的 HEALPix SQL 接口](https://gitlab.com/marjorielptt/mogipix/-/blob/6e134fe704b3a94a833d9ead1b2ae540b305a8d1/src/lib.rs)
- [固定提交的 MOC 实现](https://gitlab.com/marjorielptt/mogipix/-/blob/6e134fe704b3a94a833d9ead1b2ae540b305a8d1/src/moc.rs)

mogipix 0.0.1 版把 Rust CDS HEALPix、MOC、BMOC 和 sky-region 运算暴露给 PostgreSQL。其 SQL 函数统一使用 mgx_ 前缀，支持 cell hash、层级和邻域运算、MOC 构造与集合运算、覆盖测试，以及适合索引空间过滤的 range 形式。

### 示例

hash 函数接收以度为单位的经纬度，并在内部转换为弧度。range-MOC 辅助函数接收 PostgreSQL 整数范围：

```sql
CREATE EXTENSION mogipix;

SELECT mgx_hash(13, 36.80105218, 56.78028536);

SELECT mgx_moc_to_ascii(
    mgx_create_range_moc_psql(
        29,
        ARRAY[int8range(100, 200), int8range(300, 400)]
    )
);
```

该扩展还为 MOC 和 BMOC 值定义 PostgreSQL 类型与运算符。对于表查询，上游演示了把函数式 HEALPix 索引或整数 multirange 包含判断，与 mgx_ 转换和 coverage 函数组合使用。

### 验证与兼容性

审阅的构建使用 pgrx 0.15.0，默认目标为 PostgreSQL 17，并声明了 PostgreSQL 13 至 17 的构建 feature。安装需要超级用户。0.0.1 仍是早期接口，因此应分别测试生成的扩展 SQL 和每一个 PostgreSQL 主版本。

depth、coordinate、hash 与 range 参数会从 PostgreSQL 有符号类型跨越到 Rust 库的无符号类型。部分实现路径使用未检查的类型转换、unsafe vector 表示转换或未检查的 BMOC 构造。调用前必须验证天文坐标单位以及合法的 HEALPix depth/range 边界，并将边界和空集结果与底层 CDS 库进行对照。当前 GitLab 项目在 2025 年仍有活动，没有被标记为归档。
