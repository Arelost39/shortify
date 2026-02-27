CREATE SCHEMA shortify

CREATE TABLE shortify.encoded_links (
    id              BIGSERIAL   NOT NULL,
    decoded_link    TEXT        NOT NULL,
    encoded_link    TEXT        NOT NULL,
    created_at      TIMESTAMPTZ DEFAULT NOW()
)