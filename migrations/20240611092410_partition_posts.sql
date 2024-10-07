-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS public.posts;

CREATE TABLE public.posts (
  id uuid DEFAULT uuid_generate_v4(),
  title text not null,
  content text not null,
  mediaURL text null,
  user_id uuid not null,
  "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL DEFAULT,
  primary key (id, created_at)
) partition by range(created_at);

create index if not exists tx_post_created_at_idx on public.posts (created_at);

CREATE
OR REPLACE FUNCTION create_post_partitions(start_date DATE, end_date DATE) RETURN VOID AS $ $ DECLARE curr_start DATE := start_date;

curr_end DATE;

BEGIN WHILE curr_start < end_date LOOP curr_end := curr_start + INTERVAL '7 days';

EXECUTE format(
  'CREATE TABLE IF NOT EXISTS post_%s PARTITION OF public."posts" FOR VALUES FROM (%L) TO (%L)',
  to_char(curr_start, 'YYYYMMDD'),
  curr_start,
  curr_end
);

curr_start := curr_end;

END LOOP;

END;

$ $ LANGUAGE plpgsql;

SELECT
  create_post_partitions('2024-10-08' :: DATE, '2025-01-01' :: DATE);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd