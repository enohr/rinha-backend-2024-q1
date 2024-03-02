package database

const extratoQuery = `
		SELECT
			u.saldo,
			u.limite,
			t.valor,
			t.tipo,
			t.descricao,
			t.realizada_em
		FROM
			users u
		LEFT JOIN LATERAL (
			SELECT *
			FROM transacoes t 
			WHERE user_id = u.id
			ORDER BY t.realizada_em desc 
			LIMIT 10
			) t on TRUE
		WHERE
			u.id = $1
	`

const transacaoQuery = `
	WITH new_transaction AS (
		INSERT INTO transacoes (user_id, valor, tipo, descricao)
		VALUES ($1, $2, $3, $4)
	),
	updated_user AS (
		UPDATE users
		SET saldo = saldo + $2
		WHERE id = $1
		RETURNING saldo, limite
	)
	SELECT saldo, limite FROM updated_user;
	`
