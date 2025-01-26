grant all privileges on database debts_db to bot_service;

CREATE TABLE IF NOT EXISTS public.debts (
    debt_id TEXT NOT NULL PRIMARY KEY,
    amount INTEGER,
    currency TEXT,
    debtor_id TEXT NOT NULL,
    collector_id BIGINT NOT NULL,
    status TEXT NOT NULL,
    created_ts TIMESTAMPTZ NOT NULL,
    updated_ts TIMESTAMPTZ NOT NULL,
    version INTEGER NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS debtor_collector_status__idx on public.debts (debt_id, collector_id, status);