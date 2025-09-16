-- +goose Up
-- +goose StatementBegin
create table if not exists orders
(
    id               serial primary key,
    order_uuid       varchar(36)      not null,
    user_uuid        varchar(36)      not null,
    transaction_uuid varchar(36),
    payment_method   varchar(20)      not null,
    status           varchar(20)      not null,
    total_price      double precision not null,
    unique (order_uuid),
    unique (transaction_uuid)
);

create table if not exists orders_parts
(
    id
              serial
        primary
            key,
    order_id
              integer
                          not
                              null
        references
            orders
                (
                 id
                    ) on delete cascade,
    part_uuid varchar(36) not null
);
CREATE INDEX idx_order_id_fk ON orders_parts (order_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
