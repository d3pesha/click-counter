-- +migrate Up
CREATE TABLE IF NOT EXISTS banners (
                                       id   SERIAL PRIMARY KEY,
                                       name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS banner_clicks (
                                             id          SERIAL PRIMARY KEY,
                                             banner_id   INT NOT NULL,
                                             timestamp   TIMESTAMP NOT NULL,
                                             click_count INT NOT NULL DEFAULT 0,
                                             UNIQUE (banner_id, timestamp),
    FOREIGN KEY (banner_id) REFERENCES banners(id) ON DELETE CASCADE
);
