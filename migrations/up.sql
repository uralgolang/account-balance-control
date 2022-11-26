-- DROP SCHEMA accounts;

CREATE SCHEMA accounts AUTHORIZATION postgres;
-- accounts.account_balances definition

-- Drop table

-- DROP TABLE accounts.account_balances;

CREATE TABLE accounts.account_balances (
                                           accounts_id uuid NULL,
                                           balance int8 NULL
);

-- Permissions

ALTER TABLE accounts.account_balances OWNER TO postgres;
GRANT ALL ON TABLE accounts.account_balances TO postgres;


-- accounts.account_operations definition

-- Drop table

-- DROP TABLE accounts.account_operations;

CREATE TABLE accounts.account_operations (
                                             account_operations_id uuid NULL,
                                             operation_type int2 NULL,
                                             accounts_id uuid NULL,
                                             amount int8 NULL,
                                             operation_time timestamp NULL
);

-- Permissions

ALTER TABLE accounts.account_operations OWNER TO postgres;
GRANT ALL ON TABLE accounts.account_operations TO postgres;


-- accounts.account_transfers definition

-- Drop table

-- DROP TABLE accounts.account_transfers;

CREATE TABLE accounts.account_transfers (
                                            account_transfers_id uuid NULL,
                                            transfer_time timestamp NULL,
                                            payer_id uuid NULL,
                                            recipient_id uuid NULL,
                                            amount int8 NULL
);

-- Permissions

ALTER TABLE accounts.account_transfers OWNER TO postgres;
GRANT ALL ON TABLE accounts.account_transfers TO postgres;




-- Permissions

GRANT ALL ON SCHEMA accounts TO postgres;
