CREATE TABLE `currency_conversion_rate` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'auto increment id',
  `from_currency_id` int(11) NOT NULL COMMENT 'from currency identifier',
  `to_currency_id` int(11) NOT NULL COMMENT 'to currency identifier',
  `rate` DECIMAL(19,4) UNSIGNED NOT NULL COMMENT 'conversion rate',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created date',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT 'user who create this entity',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update date',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT 'user who update this entity',
  `is_deleted` tinyint(1) DEFAULT '0' COMMENT 'deleted flag',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'delete date',
  `deleted_by` varchar(255) DEFAULT NULL COMMENT 'user who delete this entity',
  
  PRIMARY KEY (`id`),
  UNIQUE KEY `UC_from_to` (`from_currency_id`,`to_currency_id`,`is_deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
