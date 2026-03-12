


## Usage

> [dbt2: OSDL-DBT-2 test kit](https://github.com/osdldbt/dbt2)

`dbt2` is a TPC-C benchmark implementation for PostgreSQL. The extension provides stored procedures that implement the five standard TPC-C transaction types.

```sql
CREATE EXTENSION dbt2;
```

### TPC-C Transaction Types

The extension provides stored procedures for the five standard TPC-C transactions:

- **New Order**: Creates a new order with multiple line items, updating stock levels
- **Payment**: Processes a customer payment, updating warehouse and district balances
- **Order Status**: Retrieves the status of a customer's most recent order
- **Delivery**: Processes pending orders for delivery across all districts
- **Stock Level**: Checks the count of recently sold items with low stock

### Benchmark Workflow

The `dbt2` system consists of:

1. **Database extension** (`dbt2`): Stored procedures for TPC-C transactions
2. **Data loader**: Populates the benchmark tables with TPC-C data
3. **Driver**: Generates transaction workloads simulating terminal users
4. **Client**: Manages connections between the driver and database

### Running Benchmarks

The benchmark is typically run using the `dbt2` command-line tools (separate from the extension):

```bash
# Build the benchmark database
dbt2 build --dbms pgsql --warehouses 10

# Run the benchmark
dbt2 run --dbms pgsql --warehouses 10 --duration 300 --connections 10

# Generate report
dbt2 report --dbms pgsql
```

### TPC-C Schema

The benchmark uses these standard tables: `warehouse`, `district`, `customer`, `history`, `new_order`, `orders`, `order_line`, `item`, and `stock`.

Refer to the `doc/` directory in the repository for detailed configuration and tuning options.
