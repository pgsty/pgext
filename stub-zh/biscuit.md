

## 用法

> [GitHub: CrystallineCore/pg_biscuit](https://github.com/CrystallineCore/pg_biscuit)

`biscuit`（pg_biscuit）是一个 PostgreSQL 扩展，提供类似 IAM 的模式匹配和位图索引。它使用专用位图索引实现权限风格模式与文本值的高效匹配。

## 功能特性

- **类 IAM 模式匹配**：支持类似 AWS IAM 策略模式的通配符匹配
- **位图索引**：使用位图索引加速模式匹配查询
- **权限评估**：评估给定操作是否与权限模式集合匹配

## 快速开始

```sql
CREATE EXTENSION biscuit CASCADE;  -- 需要 plpgsql

-- 创建含权限模式的表
CREATE TABLE permissions (
  id serial PRIMARY KEY,
  pattern text NOT NULL
);

-- 插入类 IAM 模式
INSERT INTO permissions (pattern) VALUES
  ('s3:GetObject'),
  ('s3:*'),
  ('ec2:Describe*'),
  ('iam:Create*');
```

## 模式语法

Biscuit 支持 IAM 风格的通配符模式：

- `*` 匹配任意字符序列
- `?` 匹配任意单个字符
- 精确字符串按字面匹配

## 说明

- 需要 `plpgsql` 扩展（使用 `CASCADE` 自动安装）
- 可用于 PostgreSQL 16、17 和 18
- MIT 许可证
