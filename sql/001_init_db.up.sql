use registry;

create table access (
 access_id int NOT NULL AUTO_INCREMENT,
 access_code char(1),
 comment varchar (30),
 primary key (access_id)
);

insert into access values 
( 1, 'A', 'All access for the system'),
( 2, 'M', 'Management access for project'),
( 3, 'R', 'Read access for project'),
( 4, 'W', 'Write access for project'),
( 5, 'D', 'Delete access for project'),
( 6, 'S', 'Search access for project');


create table role (
 role_id int NOT NULL AUTO_INCREMENT,
 role_code varchar(20),
 name varchar (20),
 primary key (role_id)
);

insert into role values 
( 1, 'AMDRWS', 'sysAdmin'),
( 2, 'MDRWS', 'projectAdmin'),
( 3, 'RWS', 'developer'),
( 4, 'RS', 'guest');


create table user (
 user_id int NOT NULL AUTO_INCREMENT,
 username varchar(15),
 email varchar(30),
 password varchar(40) NOT NULL,
 realname varchar (20) NOT NULL,
 comment varchar (30),
 deleted tinyint (1) DEFAULT 0 NOT NULL,
 reset_uuid varchar(40) DEFAULT NULL,
 salt varchar(40) DEFAULT NULL,
 primary key (user_id),
 UNIQUE (username),
 UNIQUE (email)
);

insert into user values 
(1, 'admin', 'admin@example.com', '', 'system admin', 'admin user',0, null, ''),
(2, 'anonymous', 'anonymous@example.com', '', 'anonymous user', 'anonymous user', 1, null, '');
                                                                          
create table project (
 project_id int NOT NULL AUTO_INCREMENT,
 owner_id int NOT NULL,
 name varchar (30) NOT NULL,
 creation_time timestamp,
 deleted tinyint (1) DEFAULT 0 NOT NULL,
 public tinyint (1) DEFAULT 0 NOT NULL,
 primary key (project_id),
 FOREIGN KEY (owner_id) REFERENCES user(user_id),
 UNIQUE (name)
);

insert into project values 
(1, 1, 'library', NOW(), 0, 1);

create table project_role (
 pr_id int NOT NULL AUTO_INCREMENT,
 project_id int NOT NULL,
 role_id int NOT NULL,
 primary key (pr_id),
 FOREIGN KEY (role_id) REFERENCES role(role_id),
 FOREIGN KEY (project_id) REFERENCES project (project_id)
);

insert into project_role values
( 1,1,1 );

create table user_project_role (
 upr_id int NOT NULL AUTO_INCREMENT,
 user_id int NOT NULL,
 pr_id int NOT NULL,
 primary key (upr_id),
 FOREIGN KEY (user_id) REFERENCES user(user_id),
 FOREIGN KEY (pr_id) REFERENCES project_role (pr_id)
);

insert into user_project_role values
( 1,1,1 );

create table access_log (
 log_id int NOT NULL AUTO_INCREMENT,
 user_id int NOT NULL,
 project_id int NOT NULL,
 repo_name varchar (40), 
 GUID varchar(64), 
 operation varchar(20) NOT NULL,
 op_time timestamp,
 primary key (log_id),
 FOREIGN KEY (user_id) REFERENCES user(user_id),
 FOREIGN KEY (project_id) REFERENCES project (project_id)
);

CREATE TABLE repository (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL DEFAULT '',
  project_name varchar(255) NOT NULL DEFAULT '',
  project_id bigint(11) NOT NULL,
  created_at datetime DEFAULT NULL,
  updated_at datetime DEFAULT NULL,
  user_name varchar(155) NOT NULL,
  category varchar(255),
  is_public tinyint(2) NOT NULL DEFAULT 1,
  latest_tag varchar(255) NOT NULL DEFAULT 'latest',
  description varchar(512),

  PRIMARY KEY(`id`),
  KEY `index_repository_project_id` (`project_id`)
);

CREATE TABLE tag (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  project_id bigint(11) NOT NULL,
  repository_id bigint(11) NOT NULL,
  version varchar(255) NOT NULL DEFAULT '',
  created_at datetime DEFAULT NULL,
  updated_at datetime DEFAULT NULL,
  PRIMARY KEY(`id`),
  KEY `index_tag_reposotory_id` (`repository_id`)
);
