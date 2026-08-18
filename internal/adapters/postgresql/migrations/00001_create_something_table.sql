-- +goose Up
CREATE TABLE IF NOT EXISTS something (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS something;
