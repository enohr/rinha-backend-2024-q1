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
		JOIN LATERAL (
			SELECT *
			FROM transacoes t 
			WHERE user_id = u.id
			ORDER BY t.realizada_em desc 
			LIMIT 10
			) t on TRUE
		WHERE
			u.id = $1
	`
