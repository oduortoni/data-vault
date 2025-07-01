# data-vault
A privacy-first, extensible, pluggable data store system for structured personal data entry and access control, with minimal infrastructure overhead.

## usage

```bash
git clone https://github.com/oduortoni/data-vault.git
cd data-vault/dv
cargo run
```

install sqlx-cli if you have not:

```bash
cargo install sqlx-cli --no-default-features --features sqlite
```

ensure your .env file is present with:

```plaintext
DATABASE_URL=sqlite://./data.sqlite
```

then initialize the migration directory if you haven't:

```bash
sqlx migrate add create_users_table
```

run the database migrations

```bash
sqlx migrate run
```

the open your browser and navigate to the url

```plaintext
http://localhost:10000
```

## port

If you are experience port issues, you can defin ean environment variable called PORT and then restart the project

```bash
export PORT=10000
cargo run
```

## preview

You can check the live preview of this site at:

[data-vault](https://data-vault.onrender.com)

## contributions
