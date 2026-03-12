


## Usage

> [omni_datasets: Dataset provisioning](https://docs.omnigres.org/omni_datasets/northwind/)

The `omni_datasets` extension provides access to sample datasets, currently including the Northwind database.

### Northwind Dataset

```sql
CREATE EXTENSION omni_datasets;
CREATE SCHEMA northwind;
SELECT omni_datasets.instantiate_northwind(schema => 'northwind');
```

The Northwind dataset is a classic sample database originally from Microsoft Access/SQL Server, using a PostgreSQL-compatible version maintained by Yugabyte. It provides a small but realistic business dataset with tables for customers, orders, products, employees, and more.
