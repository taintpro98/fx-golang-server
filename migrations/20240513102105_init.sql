-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- dong nay de tu dong sinh uuid
DROP TABLE IF EXISTS public.constants;

create table public.constants (
    "code" varchar primary key,
    "value" text not null
);

DROP TABLE IF EXISTS public.users;

create table public.users (
    id uuid DEFAULT uuid_generate_v4() primary key,
    "loyalty_id" int NULL,
    "email" varchar NULL,
    "phone" varchar NOT NULL,
    "cur_original_id" varchar NULL,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS public.feeds;

CREATE TABLE public.feeds (
    id bigserial primary key,
    user_id uuid not null,
    "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp NULL DEFAULT
);

CREATE UNIQUE INDEX slot_seats_slot_id_seat_id_u_idx ON public.slot_seats(slot_id, seat_id);

CREATE INDEX orders_user_id_idx ON public.orders (user_id);

CREATE INDEX orders_slot_id_idx ON public.orders (slot_id);

CREATE UNIQUE INDEX seats_seat_code_room_id_u_idx ON public.seats (seat_code, room_id);

CREATE UNIQUE INDEX users_loyalty_id_u_idx ON public.users(loyalty_id);

CREATE UNIQUE INDEX users_phone_u_idx ON public.users(phone);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd