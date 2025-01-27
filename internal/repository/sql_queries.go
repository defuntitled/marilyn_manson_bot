package repository

const (
	kGetDebtsByCollectorId = `
	SELECT 
	debt_id,
    amount,
    currency, 
    debtor_id, 
    collector_id, 
    status, 
    created_ts, 
    updated_ts, 
    version
	FROM public.debts
	WHERE collector_id = $1::BIGINT AND
	status = 1;
	`
	kInsertDebt = `
	INSERT INTO public.debts 
	SELECT * FROM jsonb_populate_record(NULL::debts, $1::JSONB)
	ON CONFLICT DO NOTHING;
	`
	kUpdateDebt = `
	UPDATE public.debts SET (
	debt_id,
    amount,
    currency, 
    debtor_id, 
    collector_id, 
    status, 
    updated_ts, 
    version
	) = (
	SELECT debt_id,
    amount,
    currency, 
    debtor_id, 
    collector_id, 
    status, 
    NOW(), 
    version + 1
	 FROM jsonb_populate_record(NULL::debts, $1::JSONB)
	)
	WHERE debt_id = ($1::JSONB->>'debt_id')::TEXT AND version = ($1::JSONB->>'version')::INTEGER;
	`
)
