use registry;

CREATE TABLE categories (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY(`id`)
);
