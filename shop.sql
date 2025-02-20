SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order_address
-- 订单地址表
-- ----------------------------
DROP TABLE IF EXISTS `order_address`;
CREATE TABLE `order_address`  (
                               `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '地址id',
                               `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户id',
                               `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '收货人姓名',
                               `phone` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '收货人电话',
                               `zip_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮编',
                               `state` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '省份',
                               `city` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '市',
                               `district` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '区/县',
                               `address` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '详细地址',
                               `default_address` tinyint(1) NULL DEFAULT NULL COMMENT '1：默认地址  0：非默认地址',
                               `label` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地址标签',
                               'is_delete'     TINYINT   DEFAULT 0                 NULL COMMENT '是否删除',
                               PRIMARY KEY (`id`) USING BTREE,
                               INDEX `userId`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for category
-- 商品类目表
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
                                `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '类目id',
                                `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '类目名称',
                                `parent_id` bigint(20) NOT NULL COMMENT '父类目id,顶级类目填0',
                                `is_parent` tinyint(1) NOT NULL COMMENT '是否为父节点，0为否，1为是',
                                'is_delete'     TINYINT   DEFAULT 0                 NULL COMMENT '是否删除',
                                PRIMARY KEY (`id`) USING BTREE,
                                INDEX `key_parent_id`(`parent_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '商品类目表，类目和商品(spu)是一对多关系，类目与品牌是多对多关系' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for category_brand
-- 商品分类和品牌的中间表
-- ----------------------------
DROP TABLE IF EXISTS `category_brand`;
CREATE TABLE `category_brand`  (
                                      `category_id` bigint(20) NOT NULL COMMENT '商品类目id',
                                      `brand_id` bigint(20) NOT NULL COMMENT '品牌id',
                                      'is_delete'     TINYINT   DEFAULT 0                 NULL COMMENT '是否删除',
                                      PRIMARY KEY (`category_id`, `brand_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '商品分类和品牌的中间表，两者是多对多关系' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for order
-- 商品订单表
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
                             `order_id` bigint(20) NOT NULL COMMENT '订单id',
                             `total_pay` bigint(20) NOT NULL COMMENT '总金额，单位为分',
                             `actual_pay` bigint(20) NOT NULL COMMENT '实付金额。单位:分。如:20007，表示:200元7分',
                             `promotion_ids` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT '',
                             `payment_type` tinyint(1) UNSIGNED ZEROFILL NOT NULL COMMENT '支付类型，1、在线支付，2、先用后付',
                             `post_fee` bigint(20) NOT NULL COMMENT '邮费。单位:分。如:20007，表示:200元7分',
                             `create_time` datetime(0) NULL DEFAULT NULL COMMENT '订单创建时间',
                             `shipping_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '物流名称',
                             `shipping_code` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '物流单号',
                             `user_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '用户id',
                             `buyer_message` varchar(100) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '买家留言',
                             `buyer_nick` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '买家昵称',
                             `buyer_rate` tinyint(1) NULL DEFAULT NULL COMMENT '买家是否已经评价,0未评价，1已评价',
                             `receiver_state` varchar(100) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT '' COMMENT '收获地址（省）',
                             `receiver_city` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT '' COMMENT '收获地址（市）',
                             `receiver_district` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT '' COMMENT '收获地址（区/县）',
                             `receiver_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT '' COMMENT '收获地址（街道、住址等详细地址）',
                             `receiver_mobile` varchar(12) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '收货人手机',
                             `receiver_zip` varchar(15) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '收货人邮编',
                             `receiver` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '收货人',
                             `invoice_type` int(1) NULL DEFAULT 0 COMMENT '发票类型(0无发票1普通发票，2电子发票，3增值税发票)',
                             `source_type` int(1) NULL DEFAULT 2 COMMENT '订单来源：1:app端，2：pc端，3：M端，4：微信端，5：手机qq端',
                             'is_delete'     TINYINT   DEFAULT 0                 NULL COMMENT '是否删除',
                             PRIMARY KEY (`order_id`) USING BTREE,
                             INDEX `create_time`(`create_time`) USING BTREE,
                             INDEX `buyer_nick`(`buyer_nick`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for order_detail
-- 订单详情表
-- ----------------------------
DROP TABLE IF EXISTS `order_detail`;
CREATE TABLE `order_detail`  (
                                    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '订单详情id ',
                                    `order_id` bigint(20) NOT NULL COMMENT '订单id',
                                    `sku_id` bigint(20) NOT NULL COMMENT 'sku商品id',
                                    `num` int(11) NOT NULL COMMENT '购买数量',
                                    `title` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品标题',
                                    `own_spec` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '商品动态属性键值集',
                                    `price` bigint(20) NOT NULL COMMENT '价格,单位：分',
                                    `image` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '商品图片',
                                    'is_delete'     TINYINT   DEFAULT 0                 NULL COMMENT '是否删除',
                                    PRIMARY KEY (`id`) USING BTREE,
                                    INDEX `key_order_id`(`order_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '订单详情表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for order_status
-- 订单状态表
-- ----------------------------
DROP TABLE IF EXISTS `order_status`;
CREATE TABLE `order_status`  (
                                    `order_id` bigint(20) NOT NULL COMMENT '订单id',
                                    `status` int(1) NULL DEFAULT NULL COMMENT '状态：1、未付款 2、已付款,未发货 3、已发货,未确认 4、交易成功 5、交易关闭 6、已评价',
                                    `create_time` datetime(0) NULL DEFAULT NULL COMMENT '订单创建时间',
                                    `payment_time` datetime(0) NULL DEFAULT NULL COMMENT '付款时间',
                                    `consign_time` datetime(0) NULL DEFAULT NULL COMMENT '发货时间',
                                    `end_time` datetime(0) NULL DEFAULT NULL COMMENT '交易完成时间',
                                    `close_time` datetime(0) NULL DEFAULT NULL COMMENT '交易关闭时间',
                                    `comment_time` datetime(0) NULL DEFAULT NULL COMMENT '评论时间',
                                    'is_delete'     TINYINT   DEFAULT 0                 NULL COMMENT '是否删除',
                                    PRIMARY KEY (`order_id`) USING BTREE,
                                    INDEX `status`(`status`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '订单状态表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for seckill_order
-- 秒杀订单表
-- ----------------------------
DROP TABLE IF EXISTS `seckill_order`;
CREATE TABLE `seckill_order`  (
                                     `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                     `user_id` bigint(20) NOT NULL COMMENT '用户id',
                                     `order_id` bigint(20) NOT NULL COMMENT '订单id',
                                     `sku_id` bigint(20) NOT NULL COMMENT '商品id',
                                     'is_delete'     TINYINT   DEFAULT 0                 NULL COMMENT '是否删除',
                                     PRIMARY KEY (`id`) USING BTREE,
                                     UNIQUE INDEX `u_id_sku_id`(`user_id`, `sku_id`) USING BTREE COMMENT '用户id和商品id唯一',
                                     INDEX `key_order_id`(`order_id`) USING BTREE COMMENT '订单id索引'
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for seckill_sku
-- 秒杀商品表
-- ----------------------------
DROP TABLE IF EXISTS `seckill_sku`;
CREATE TABLE `seckill_sku`  (
                                   `id` int(11) NOT NULL AUTO_INCREMENT,
                                   `sku_id` bigint(20) NULL DEFAULT NULL COMMENT '秒杀商品id',
                                   `start_time` datetime(0) NOT NULL COMMENT '秒杀开始时间',
                                   `end_time` datetime(0) NULL DEFAULT NULL COMMENT '秒杀结束时间',
                                   `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标题',
                                   `seckill_price` bigint(15) NULL DEFAULT NULL COMMENT '秒杀价格，单位为分',
                                   'is_delete'     TINYINT   DEFAULT 0                 NULL COMMENT '是否删除',
                                   `image` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '商品图片',
                                   `enable` tinyint(1) NOT NULL COMMENT '是否可以秒杀',
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sku
-- 商品sku表，通过spu和sku组合形成商品
-- ----------------------------
DROP TABLE IF EXISTS `sku`;
CREATE TABLE `sku`  (
                           `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'sku id',
                           `spu_id` bigint(20) NOT NULL COMMENT 'spu id',
                           `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品标题',
                           `images_url` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '商品的图片，多个图片以‘,’分割',
                           `price` bigint(15) NOT NULL DEFAULT 0 COMMENT '销售价格',
                           `indexes` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '特有规格属性在spu属性模板中的对应下标组合',
                           `own_spec` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT 'sku的特有规格参数键值对',
                           'is_delete'     TINYINT   DEFAULT 0                 NULL COMMENT '是否删除',
                           'create_time'   DATETIME  DEFAULT CURRENT_TIMESTAMP NULL COMMENT '创建时间',
                           'update_time'   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
                           PRIMARY KEY (`id`) USING BTREE,
                           INDEX `key_spu_id`(`spu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = 'sku表,该表表示具体的商品实体,如黑色的 64g的iphone 8' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for spu
-- 商品spu表，通过不同类目和品牌组合在一起形成spu
-- ----------------------------
DROP TABLE IF EXISTS `spu`;
CREATE TABLE `spu`  (
                           `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'spu id',
                           `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '标题，即商品名称',
                           `sub_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '子标题，即商品的简介',
                           `cid1` bigint(20) NOT NULL COMMENT '1级类目id',
                           `cid2` bigint(20) NOT NULL COMMENT '2级类目id',
                           `cid3` bigint(20) NOT NULL COMMENT '3级类目id',
                           `brand_id` bigint(20) NOT NULL COMMENT '商品所属品牌id',
                           `saleable` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否上架，0下架，1上架',
                           'is_delete'     TINYINT   DEFAULT 0                 NULL COMMENT '是否删除',
                           'create_time'   DATETIME  DEFAULT CURRENT_TIMESTAMP NULL COMMENT '创建时间',
                           'update_time'   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
                           PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = 'spu表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for spu_detail
-- 描述手机等数码产品的详细信息
-- ----------------------------
DROP TABLE IF EXISTS `spu_detail`;
CREATE TABLE `spu_detail`  (
                                  `spu_id` bigint(20) NOT NULL,
                                  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '商品描述信息',
                                  `specifications` varchar(10000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '全部规格参数数据',
                                  `spec_template` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '特有规格参数及可选值信息',
                                  `packing_list` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '包装清单',
                                  'is_delete'     TINYINT   DEFAULT 0                 NULL COMMENT '是否删除',
                                  `after_service` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '售后服务',
                                  PRIMARY KEY (`spu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock
-- 商品库存表
-- ----------------------------
DROP TABLE IF EXISTS `stock`;
CREATE TABLE `stock`  (
                             `sku_id` bigint(20) NOT NULL COMMENT '库存对应的商品sku id',
                             `seckill_stock` int(9) NULL DEFAULT 0 COMMENT '可秒杀库存',
                             `seckill_total` int(9) NULL DEFAULT 0 COMMENT '秒杀总数量',
                             `stock` int(9) NOT NULL COMMENT '库存数量',
                             'is_delete'     TINYINT   DEFAULT 0                 NULL COMMENT '是否删除',
                             PRIMARY KEY (`sku_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '库存表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- 用户表
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT,
                            `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
                            'avatar_url' varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像',
                            `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码，加密存储',
                            `phone` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '注册手机号',
                            `created` datetime(0) NOT NULL COMMENT '创建时间',
                            'is_delete'     TINYINT   DEFAULT 0                 NULL COMMENT '是否删除',
                            PRIMARY KEY (`id`) USING BTREE,
                            UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;
