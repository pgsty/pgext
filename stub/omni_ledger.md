


## Usage

> [omni_ledger: Financial ledger](https://docs.omnigres.org/omni_ledger/basics/)

The `omni_ledger` extension provides primitives for building financial or finance-like systems where transfers of value are recorded and balanced.

### Core Concepts

- **Ledger** -- A set of accounts representing a specific financial entity (currency, unit of value, or company accounting system). Each ledger maintains independent balance tracking.

- **Account** -- A distinct entity within a ledger to which value can be credited or debited. Common uses: customer wallets, company bank accounts, project funds, liability accounts.

- **Transfer** -- The atomic movement of value between two accounts, ensuring balanced debits and credits. The term "transfer" is used instead of "transaction" to avoid confusion with database transactions.

- **Account Category** -- Classifies accounts by financial nature (assets, liabilities, equity). Once assigned, a category is permanent.

### Dependencies

Requires `omni_id` and `omni_polyfill`.
