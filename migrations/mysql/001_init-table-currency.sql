CREATE TABLE `currency` (
  `id` int(11) NOT NULL COMMENT 'currency identifier',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT 'currency name',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created date',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT 'user who create this entity',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update date',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT 'user who update this entity',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
