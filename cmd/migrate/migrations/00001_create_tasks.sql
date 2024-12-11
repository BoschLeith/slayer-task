-- +goose Up
CREATE TABLE IF NOT EXISTS tasks(
  id bigserial PRIMARY KEY,
  equipment TEXT,
  inventry TEXT,
  monster varchar(255),
  notes TEXT,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS tasks;
