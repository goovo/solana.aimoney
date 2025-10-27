-- 创建加密货币数据表
CREATE TABLE IF NOT EXISTS `sys_cryptos` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `rank` int(11) NOT NULL DEFAULT '0' COMMENT '排名',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '名称',
  `symbol` varchar(20) NOT NULL DEFAULT '' COMMENT '符号',
  `icon` varchar(500) NOT NULL DEFAULT '' COMMENT '图标URL',
  `price` decimal(20,8) NOT NULL DEFAULT '0.00000000' COMMENT '价格',
  `change_24h` decimal(10,4) NOT NULL DEFAULT '0.0000' COMMENT '24小时涨跌幅',
  `market_cap` decimal(30,2) NOT NULL DEFAULT '0.00' COMMENT '市值',
  `volume_24h` decimal(30,2) NOT NULL DEFAULT '0.00' COMMENT '24小时交易量',
  `circulating_supply` decimal(30,2) NOT NULL DEFAULT '0.00' COMMENT '流通量',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_rank` (`rank`),
  KEY `idx_symbol` (`symbol`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='加密货币数据表';