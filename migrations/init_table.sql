CREATE TABLE IF NOT EXISTS public.debts (
    debt_id TEXT NOT NULL PRIMARY KEY,
    amount INTEGER,
    currency TEXT,
    debtor_id TEXT NOT NULL,
    collector_id BIGINT NOT NULL,
    status INTEGER NOT NULL,
    created_ts TIMESTAMPTZ NOT NULL,
    updated_ts TIMESTAMPTZ NOT NULL,
    version INTEGER NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS debtor_collector_status__idx on public.debts (debtor_id, collector_id, status) WHERE status = 1;