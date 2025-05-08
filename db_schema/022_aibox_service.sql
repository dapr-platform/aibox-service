-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

-- 设备表
CREATE TABLE o_aibox_device (
    id VARCHAR(36) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    ip VARCHAR(255) NOT NULL,
    build_time_str VARCHAR(255),
    latest_heart_beat_time TIMESTAMP,
    status INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_aibox_device IS 'AI盒子设备表';
COMMENT ON COLUMN o_aibox_device.id IS '设备ID';
COMMENT ON COLUMN o_aibox_device.name IS '设备名称';
COMMENT ON COLUMN o_aibox_device.ip IS '设备IP地址';
COMMENT ON COLUMN o_aibox_device.build_time_str IS '设备构建时间';
COMMENT ON COLUMN o_aibox_device.latest_heart_beat_time IS '最近心跳时间';
COMMENT ON COLUMN o_aibox_device.status IS '设备状态(0:离线，1:在线)';

-- 事件表
CREATE TABLE o_aibox_event (
    id VARCHAR(36) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    dn VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    device_id VARCHAR(36) NOT NULL,
    content TEXT,
    picstr TEXT,
    level INTEGER NOT NULL DEFAULT 4,
    status INTEGER NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_aibox_event IS 'AI盒子事件表';
COMMENT ON COLUMN o_aibox_event.id IS '事件ID';
COMMENT ON COLUMN o_aibox_event.dn IS '设备编号';
COMMENT ON COLUMN o_aibox_event.title IS '事件标题';
COMMENT ON COLUMN o_aibox_event.device_id IS '关联设备ID';
COMMENT ON COLUMN o_aibox_event.content IS '事件内容';
COMMENT ON COLUMN o_aibox_event.picstr IS '图片信息';
COMMENT ON COLUMN o_aibox_event.level IS '事件级别(1:紧急, 2:严重, 3:轻微, 4:警告)';
COMMENT ON COLUMN o_aibox_event.status IS '事件状态(0:清除, 1:活动)';

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
COMMENT ON COLUMN o_aibox_update.description IS '更新描述';
COMMENT ON COLUMN o_aibox_update.status IS '状态(0:禁用, 1:启用)';

-- 创建索引
CREATE INDEX idx_aibox_device_status ON o_aibox_device(status);
CREATE INDEX idx_aibox_event_device_id ON o_aibox_event(device_id);
CREATE INDEX idx_aibox_event_level ON o_aibox_event(level);
CREATE INDEX idx_aibox_event_status ON o_aibox_event(status);
CREATE INDEX idx_aibox_update_version ON o_aibox_update(version);
CREATE INDEX idx_aibox_update_type ON o_aibox_update(type);
CREATE INDEX idx_aibox_update_status ON o_aibox_update(status);

-- 创建视图
CREATE VIEW v_aibox_device_info AS
SELECT
    d.id AS id,
    d.name AS name,
    d.ip AS ip,
    d.build_time_str AS build_time_str,
    d.latest_heart_beat_time AS latest_heart_beat_time,
    d.status AS status,
    CASE d.status 
        WHEN 0 THEN '离线'
        WHEN 1 THEN '在线'
        ELSE '未知'
    END AS status_name,
    (SELECT COUNT(*) FROM o_aibox_event WHERE device_id = d.id AND status = 1) AS active_event_count,
    (SELECT COUNT(*) FROM o_aibox_event WHERE device_id = d.id AND level = 1 AND status = 1) AS critical_event_count,
    (SELECT COUNT(*) FROM o_aibox_event WHERE device_id = d.id AND level = 2 AND status = 1) AS major_event_count,
    (SELECT COUNT(*) FROM o_aibox_event WHERE device_id = d.id AND level = 3 AND status = 1) AS minor_event_count,
    (SELECT COUNT(*) FROM o_aibox_event WHERE device_id = d.id AND level = 4 AND status = 1) AS warning_event_count
FROM
    o_aibox_device d;

COMMENT ON VIEW v_aibox_device_info IS 'AI盒子设备信息视图';
COMMENT ON COLUMN v_aibox_device_info.id IS '设备ID';
COMMENT ON COLUMN v_aibox_device_info.name IS '设备名称';
COMMENT ON COLUMN v_aibox_device_info.ip IS '设备IP地址';
COMMENT ON COLUMN v_aibox_device_info.build_time_str IS '设备构建时间';
COMMENT ON COLUMN v_aibox_device_info.latest_heart_beat_time IS '最近心跳时间';
COMMENT ON COLUMN v_aibox_device_info.status IS '设备状态(0:离线，1:在线)';
COMMENT ON COLUMN v_aibox_device_info.status_name IS '设备状态名称';
COMMENT ON COLUMN v_aibox_device_info.active_event_count IS '活动事件总数';
COMMENT ON COLUMN v_aibox_device_info.critical_event_count IS '紧急事件数';
COMMENT ON COLUMN v_aibox_device_info.major_event_count IS '严重事件数';
COMMENT ON COLUMN v_aibox_device_info.minor_event_count IS '轻微事件数';
COMMENT ON COLUMN v_aibox_device_info.warning_event_count IS '警告事件数';

-- 事件详情视图
CREATE VIEW v_aibox_event_info AS
SELECT
    e.id AS id,
    e.dn AS dn,
    e.title AS title,
    e.device_id AS device_id,
    e.content AS content,
    e.picstr AS picstr,
    e.level AS level,
    CASE e.level
        WHEN 1 THEN '紧急'
        WHEN 2 THEN '严重'
        WHEN 3 THEN '轻微'
        WHEN 4 THEN '警告'
        ELSE '未知'
    END AS level_name,
    e.status AS status,
    CASE e.status
        WHEN 0 THEN '清除'
        WHEN 1 THEN '活动'
        ELSE '未知'
    END AS status_name,
    e.created_time AS created_time,
    e.updated_time AS updated_time,
    d.name AS device_name,
    d.ip AS device_ip,
    d.status AS device_status,
    CASE d.status
        WHEN 0 THEN '离线'
        WHEN 1 THEN '在线'
        ELSE '未知'
    END AS device_status_name
FROM
    o_aibox_event e
LEFT JOIN
    o_aibox_device d ON e.device_id = d.id;

COMMENT ON VIEW v_aibox_event_info IS 'AI盒子事件详情视图';
COMMENT ON COLUMN v_aibox_event_info.id IS '事件ID';
COMMENT ON COLUMN v_aibox_event_info.dn IS '设备编号';
COMMENT ON COLUMN v_aibox_event_info.title IS '事件标题';
COMMENT ON COLUMN v_aibox_event_info.device_id IS '关联设备ID';
COMMENT ON COLUMN v_aibox_event_info.content IS '事件内容';
COMMENT ON COLUMN v_aibox_event_info.picstr IS '图片信息';
COMMENT ON COLUMN v_aibox_event_info.level IS '事件级别(1:紧急, 2:严重, 3:轻微, 4:警告)';
COMMENT ON COLUMN v_aibox_event_info.level_name IS '事件级别名称';
COMMENT ON COLUMN v_aibox_event_info.status IS '事件状态(0:清除, 1:活动)';
COMMENT ON COLUMN v_aibox_event_info.status_name IS '事件状态名称';
COMMENT ON COLUMN v_aibox_event_info.created_time IS '创建时间';
COMMENT ON COLUMN v_aibox_event_info.updated_time IS '更新时间';
COMMENT ON COLUMN v_aibox_event_info.device_name IS '设备名称';
COMMENT ON COLUMN v_aibox_event_info.device_ip IS '设备IP地址';
COMMENT ON COLUMN v_aibox_event_info.device_status IS '设备状态';
COMMENT ON COLUMN v_aibox_event_info.device_status_name IS '设备状态名称';

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
COMMENT ON COLUMN v_aibox_update_info.description IS '更新描述';
COMMENT ON COLUMN v_aibox_update_info.status IS '状态';
COMMENT ON COLUMN v_aibox_update_info.status_name IS '状态名称';
COMMENT ON COLUMN v_aibox_update_info.created_time IS '创建时间';
COMMENT ON COLUMN v_aibox_update_info.updated_time IS '更新时间';

-- 活动事件统计视图
CREATE VIEW v_aibox_active_event_stats AS
SELECT
    level,
    CASE level
        WHEN 1 THEN '紧急'
        WHEN 2 THEN '严重'
        WHEN 3 THEN '轻微'
        WHEN 4 THEN '警告'
        ELSE '未知'
    END AS level_name,
    COUNT(*) AS event_count
FROM
    o_aibox_event
WHERE
    status = 1
GROUP BY
    level
ORDER BY
    level;

COMMENT ON VIEW v_aibox_active_event_stats IS 'AI盒子活动事件统计视图';
COMMENT ON COLUMN v_aibox_active_event_stats.level IS '事件级别';
COMMENT ON COLUMN v_aibox_active_event_stats.level_name IS '事件级别名称';
COMMENT ON COLUMN v_aibox_active_event_stats.event_count IS '事件数量';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP VIEW IF EXISTS v_aibox_active_event_stats;
DROP VIEW IF EXISTS v_aibox_update_info;
DROP VIEW IF EXISTS v_aibox_event_info;
DROP VIEW IF EXISTS v_aibox_device_info;
DROP TABLE IF EXISTS o_aibox_update;
DROP TABLE IF EXISTS o_aibox_event;
DROP TABLE IF EXISTS o_aibox_device;

-- +goose StatementEnd
