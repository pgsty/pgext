

## Usage

> [GitHub: CrystallineCore/pg_biscuit](https://github.com/CrystallineCore/pg_biscuit)

`biscuit` (pg_biscuit) is a PostgreSQL extension that provides IAM-like pattern matching with bitmap indexing. It enables efficient matching of permission-style patterns against text values using a specialized bitmap index.

## Features

- **IAM-like pattern matching**: Supports wildcard pattern matching similar to AWS IAM policy patterns
- **Bitmap indexing**: Uses bitmap indexes for fast pattern matching queries
- **Permission evaluation**: Evaluate whether a given action matches a set of permission patterns

## Quick Start

```sql
CREATE EXTENSION biscuit CASCADE;  -- requires plpgsql

-- Create a table with permission patterns
CREATE TABLE permissions (
  id serial PRIMARY KEY,
  pattern text NOT NULL
);

-- Insert IAM-like patterns
INSERT INTO permissions (pattern) VALUES
  ('s3:GetObject'),
  ('s3:*'),
  ('ec2:Describe*'),
  ('iam:Create*');
```

## Pattern Syntax

Biscuit supports IAM-style wildcard patterns:

- `*` matches any sequence of characters
- `?` matches any single character
- Exact strings match literally

## Notes

- Requires the `plpgsql` extension (installed automatically with `CASCADE`)
- Available for PostgreSQL 16, 17, and 18
- Licensed under MIT
