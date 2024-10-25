CREATE TABLE IF NOT EXISTS applications
(
    id            TEXT         NOT NULL PRIMARY KEY,
    title         VARCHAR(16)  NOT NULL,
    icon          TEXT         NOT NULL,
    page          TEXT         NOT NULL,
    width         INTEGER      NOT NULL DEFAULT 800,
    Height        INTEGER      NOT NULL DEFAULT 400,
    x             INTEGER,
    y             INTEGER,
    theme         VARCHAR(8)   NOT NULL DEFAULT 'light',
    background    VARCHAR(256) NOT NULL DEFAULT 'white',
    singleton     BOOLEAN      NOT NULL DEFAULT false,
    ban_uninstall BOOLEAN      NOT NULL DEFAULT false,
    fix_on_dock   BOOLEAN      NOT NULL DEFAULT false,
    fix_on_desk   BOOLEAN      NOT NULL DEFAULT false,
    created_at    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP
);

INSERT INTO applications (id, title, icon, page, width, height, x, y, theme, background, singleton, ban_uninstall,
                          fix_on_dock,
                          fix_on_desk)
VALUES ('ab95a563-f9c8-1bd2-5cf6-9e93c80d2c3b',
        '文件管理器',
        'internal://icon-oss',
        'internal://finder',
        800, 400, 0, 0,
        'light', 'white',
        false, true, true, true),
       ('d79d66f2-2932-0753-81a1-3b66ae6da94a',
        '启动台',
        'internal://icon-app',
        'system://launchpad',
        800, 400, 0, 0,
        'light', 'linear-gradient(to right, #4e54c8, #8f94fb)',
        true, true, true, false),
       ('40e3595a-7c10-cc09-08fb-54683ff39d74',
        '终端',
        'internal://icon-terminal',
        'internal://terminal',
        800, 400, 0, 0,
        'dark', 'white',
        false, true, true, true),
       ('ef708add-8c00-f5de-4279-40115b52321d',
        '应用商店',
        'internal://icon-app-store',
        'internal://app-store',
        800, 400, 0, 0,
        'light', 'linear-gradient(to right, #06beb6, #48b1bf)',
        true, true, true, true),
       ('039b3754-2eed-5c77-67e1-00549cfc44ce',
        '设置',
        'internal://icon-setting',
        'internal://setting',
        800, 400, 0, 0,
        'light', 'linear-gradient(to right, #536976, #292e49)',
        true, true, true, true);