CREATE TABLE IF NOT EXISTS producao_pedidos (
	pedido_id bigserial NOT NULL,
	status varchar(255) NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT producao_pedidos_pkey PRIMARY KEY (pedido_id)
);
CREATE INDEX idx_items_deleted_at ON items USING btree (deleted_at);
