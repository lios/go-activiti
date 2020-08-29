/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50018
Source Host           : localhost:3306
Source Database       : go-workflow

Target Server Type    : MYSQL
Target Server Version : 50018
File Encoding         : 65001

Date: 2020-07-07 23:57:52
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for bytearry
-- ----------------------------
DROP TABLE IF EXISTS `bytearry`;
CREATE TABLE `bytearry` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `name` varchar(64) collate utf8_bin default NULL COMMENT '名称',
  `key` varchar(64) collate utf8_bin default NULL COMMENT '流程定义key',
  `version` bigint(2) default NULL COMMENT '版本',
  `deployment_id` bigint(64) default NULL COMMENT '流程部署id',
  `bytes` longblob COMMENT '资源文件',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for deployment
-- ----------------------------
DROP TABLE IF EXISTS `deployment`;
CREATE TABLE `deployment` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `name` varchar(64) collate utf8_bin default NULL COMMENT '名称',
  `key` varchar(64) collate utf8_bin default NULL COMMENT '流程定义key',
  `version` bigint(2) default NULL COMMENT '版本',
  `tenantI_id` bigint(64) default NULL COMMENT '租户id',
  `deploy_time` timestamp NOT NULL default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '部署时间',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for identity_link
-- ----------------------------
DROP TABLE IF EXISTS `identity_link`;
CREATE TABLE `identity_link` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `type` varchar(64) collate utf8_bin default NULL COMMENT '类型',
  `task_id` varchar(64) collate utf8_bin default NULL COMMENT '任务ID',
  `proc_inst_id` bigint(2) default NULL COMMENT '流程实例id',
  `group_id` bigint(64) default NULL COMMENT '组id',
  `user_id` timestamp NOT NULL default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '用户id',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for process_instance
-- ----------------------------
DROP TABLE IF EXISTS `process_instance`;
CREATE TABLE `process_instance` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `key` varchar(64) collate utf8_bin default NULL COMMENT '流程定义key',
  `name` varchar(64) collate utf8_bin default NULL COMMENT '名称',
  `version` bigint(2) default NULL COMMENT '版本',
  `business_key` varchar(64) collate utf8_bin default NULL,
  `tenant_id` varchar(64) collate utf8_bin default NULL COMMENT '租户id',
  `deployment_id` bigint(64) default NULL COMMENT '流程部署id',
  `start_time` timestamp NOT NULL default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '启动时间',
  `start_user_id` varchar(64) collate utf8_bin default NULL COMMENT '启动用户',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for task
-- ----------------------------
DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `task_define_key` varchar(64) collate utf8_bin default NULL COMMENT '任务key',
  `task_define_name` varchar(64) collate utf8_bin default NULL COMMENT '任务名称',
  `version` bigint(2) default NULL COMMENT '版本',
  `tenant_id` varchar(64) collate utf8_bin default NULL COMMENT '租户id',
  `deployment_id` bigint(64) default NULL COMMENT '流程部署id',
  `start_time` timestamp NOT NULL default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '启动时间',
  `assignee` varchar(64) collate utf8_bin default NULL COMMENT '审批人',
  `proc_inst_id` bigint(64) default NULL COMMENT '实例iD',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for task
-- ----------------------------
DROP TABLE IF EXISTS `variable`;
CREATE TABLE `variable` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `version` bigint(2) default NULL COMMENT '版本',
  `task_id` varchar(64) collate utf8_bin default NULL COMMENT '任务ID',
  `proc_inst_id` bigint(2) default NULL COMMENT '流程实例id',
  `name` varchar(64) collate utf8_bin default NULL COMMENT '变量名',
  `type` varchar(10) default NULL COMMENT '类型',
  `number` int default NULL COMMENT '类型',
  `date` datetime default NULL COMMENT '时间类型',
  `float` double collate utf8_bin default NULL COMMENT 'double',
  `text` varchar(4000) default NULL COMMENT '字符',
  `blob` longblob default NULL COMMENT '字符串',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

