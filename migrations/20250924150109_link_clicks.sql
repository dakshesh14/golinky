-- +goose Up
-- +goose StatementBegin
CREATE TABLE link_clicks (
    id BIGSERIAL PRIMARY KEY,
    link_id UUID NOT NULL REFERENCES links (id) ON DELETE CASCADE,
    clicked_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ip_address INET,
    user_agent TEXT
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_link_clicks_link_id ON link_clicks (link_id);
CREATE INDEX idx_link_clicks_clicked_at ON link_clicks (clicked_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS link_clicks;
-- +goose StatementEnd
