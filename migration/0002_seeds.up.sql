-- DELETE CURRENT DATA
DELETE FROM producao_pedidos;
TRUNCATE producao_pedidos RESTART IDENTITY;

-- INSERT ITEMS
INSERT INTO producao_pedidos (pedido_id, status, created_at, updated_at, deleted_at)
VALUES ('1', 'CRIADO', timestamp '2023-07-20 12:00', null, null),
       ('2', 'CRIADO', timestamp '2023-07-21 13:00', null, null),
       ('3', 'AGUARDANDO_PAGAMENTO', timestamp '2023-07-22 14:50', timestamp '2023-07-22 14:54', null),
       ('4', 'RECEBIDO', timestamp '2023-07-21 11:54:13', timestamp '2023-07-21 11:55:56', null),
       ('5', 'PRONTO', timestamp '2023-07-20 02:00', timestamp '2023-07-20 02:31', null),
       ('6', 'PRONTO', timestamp '2023-07-20 01:00', timestamp '2023-07-20 01:29', null),
       ('7', 'EM_PREPARACAO', timestamp '2023-07-23 14:26', timestamp '2023-07-23 14:36', null),
       ('8', 'FINALIZADO', timestamp '2023-07-22 12:26', timestamp '2023-07-22 12:56', null);

