CREATE TABLE `ct_data_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `uuid` char(36) NOT NULL COMMENT '采集标识',
  `trace_id` char(32) NOT NULL COMMENT '链路标识',
  `dispatch` char(32) NOT NULL COMMENT '调度标识',
  `step_id` varchar(255) NOT NULL COMMENT '步骤标识',
  `info_data` text NOT NULL COMMENT '采集的数据(json)',
  `duration_ms` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '采集耗时(ms)',
  `create_time` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`),
  KEY `trace_id` (`trace_id`),
  KEY `dispatch` (`dispatch`)
) ENGINE=InnoDB COMMENT '数据采集日志';

CREATE TABLE `ct_steps` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `uuid` char(36) NOT NULL COMMENT '标识',
  `name` varchar(255) NOT NULL COMMENT '名称',
  `type` char(36) NOT NULL COMMENT '类型',
  `vars` text NOT NULL COMMENT '变量模板(json)',
  `data` varchar(1000) NOT NULL COMMENT '数据模板(json)',
  `request` text NOT NULL COMMENT '请求参数(json)',
  `response` text NOT NULL COMMENT '响应参数配置(json)',
  `create_time` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`)
) ENGINE=InnoDB COMMENT '采集步骤';

CREATE TABLE `ct_dispatch` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `uuid` char(36) NOT NULL COMMENT '标识',
  `name` varchar(255) NOT NULL COMMENT '名称',
  `type` tinyint NOT NULL COMMENT '类型',
  `way` tinyint NOT NULL COMMENT '方式',
  `rules` text NOT NULL COMMENT '规则(json)',
  `create_time` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_uuid` (`uuid`),
  KEY `idx_name` (`name`),
  KEY `idx_create_time` (`create_time`),
) ENGINE=InnoDB COMMENT '调度器';
