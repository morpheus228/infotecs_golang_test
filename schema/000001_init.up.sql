CREATE TABLE wallets
(
    id varchar(32) not null unique,
    balance numeric not null
);

CREATE TABLE transactions
(
    id serial not null unique,
    created_at TIMESTAMP WITH TIME ZONE not null,
    amount numeric not null,
    from_wallet_id varchar(32) references wallets (id) on delete cascade not null,
    to_wallet_id varchar(32) references wallets (id) on delete cascade not null
);
