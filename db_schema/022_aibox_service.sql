-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

-- 软件更新表
CREATE TABLE o_aibox_update (
    id VARCHAR(36) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version VARCHAR(64) NOT NULL,
    type INTEGER NOT NULL,
    file_path VARCHAR(255) NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_key VARCHAR(32) NOT NULL,
    description TEXT,
    status INTEGER NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_aibox_update IS 'AI盒子软件更新表';
COMMENT ON COLUMN o_aibox_update.id IS '更新ID(version+type的md5)';
COMMENT ON COLUMN o_aibox_update.version IS '版本号';
COMMENT ON COLUMN o_aibox_update.type IS '更新类型(1:应用, 2:模型, 3:配置, 4:其他)';
COMMENT ON COLUMN o_aibox_update.file_path IS '文件存放路径';
COMMENT ON COLUMN o_aibox_update.file_name IS '文件名';
COMMENT ON COLUMN o_aibox_update.file_key IS '文件key';
COMMENT ON COLUMN o_aibox_update.description IS '更新描述';
COMMENT ON COLUMN o_aibox_update.status IS '状态(0:禁用, 1:启用)';

-- 创建索引

CREATE INDEX IF NOT EXISTS idx_aibox_update_version ON o_aibox_update(version);
CREATE INDEX IF NOT EXISTS idx_aibox_update_type ON o_aibox_update(type);
CREATE INDEX IF NOT EXISTS idx_aibox_update_status ON o_aibox_update(status);

-- 创建视图

-- 软件更新视图
CREATE VIEW v_aibox_update_info AS
SELECT
    u.id AS id,
    u.version AS version,
    u.type AS type,
    CASE u.type
        WHEN 1 THEN '应用'
        WHEN 2 THEN '模型'
        WHEN 3 THEN '配置'
        WHEN 4 THEN '其他'
        ELSE '未知'
    END AS type_name,
    u.file_path AS file_path,
    u.file_name AS file_name,
    u.file_key AS file_key,
    u.description AS description,
    u.status AS status,
    CASE u.status
        WHEN 0 THEN '禁用'
        WHEN 1 THEN '启用'
        ELSE '未知'
    END AS status_name,
    u.created_time AS created_time,
    u.updated_time AS updated_time
FROM
    o_aibox_update u;

COMMENT ON VIEW v_aibox_update_info IS 'AI盒子软件更新信息视图';
COMMENT ON COLUMN v_aibox_update_info.id IS '更新ID';
COMMENT ON COLUMN v_aibox_update_info.version IS '版本号';
COMMENT ON COLUMN v_aibox_update_info.type IS '更新类型';
COMMENT ON COLUMN v_aibox_update_info.type_name IS '更新类型名称';
COMMENT ON COLUMN v_aibox_update_info.file_path IS '文件存放路径';
COMMENT ON COLUMN v_aibox_update_info.file_name IS '文件名';
COMMENT ON COLUMN v_aibox_update_info.file_key IS '文件key';
COMMENT ON COLUMN v_aibox_update_info.description IS '更新描述';
COMMENT ON COLUMN v_aibox_update_info.status IS '状态';
COMMENT ON COLUMN v_aibox_update_info.status_name IS '状态名称';
COMMENT ON COLUMN v_aibox_update_info.created_time IS '创建时间';
COMMENT ON COLUMN v_aibox_update_info.updated_time IS '更新时间';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP VIEW IF EXISTS v_aibox_update_info;
DROP TABLE IF EXISTS o_aibox_update;


-- +goose StatementEnd
