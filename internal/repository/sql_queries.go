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
	FORM public.debts
	WHERE collector_id = $1::TEXT AND
	status = 'active';
	`
	kInsertDebt = `
	INSERT INTO public.debts (
	debt_id,
    amount,
    currency, 
    debtor_id, 
    collector_id, 
    status, 
    created_ts, 
    updated_ts, 
    version
	) 
	VALUES (
	 SELECT * FROM jsonb_populate_record(NULL::debts, $1::JSONB)
	 );
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
    updated_ts, 
    version + 1
	 FORM jsonb_populate_record(NULL::debts, $1::JSONB)
	)
	WHERE debt_id = ($1::JSONB->>'debt_id')::TEXT AND version = ($1::JSONB->>'version')::INTEGER;
	`
)
