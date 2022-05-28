create table user_table  -- 用户表
(
user_id bigint not null auto_increment primary key, -- 用户id
user_name varchar(32) not null unique, -- 用户昵称
account_password varchar(32) not null,  -- 账号密码
follow_count int default 0, -- 关注总数
follower_count int default 0, -- 粉丝总数
signature text, -- 个性签名
avatar varchar(200), -- 头像链接
background_image varchar(200)  -- 背景图链接
);
create table video
(
video_id bigint not null auto_increment primary key,-- 视频唯一标识
author_id bigint not null, -- 作者用户id
play_url varchar(200) not null,-- 视频播放地址
cover_url varchar(200) not null,-- 视频封面地址
video_title varchar(50) not null,  -- 视频标题
favourite_count bigint default 0,  -- 视频的点赞总数
comment_count bigint default 0,  -- 视频的评论总数
upload_date bigint default 0,  -- 视频发布时间戳
FOREIGN KEY(author_id) REFERENCES User_table(user_id)
);
create table user_follow  -- 用户关注表
(
follow_id bigint not null,  -- 用户的id，即当前用户
follower_id bigint not null,  -- 这个用户的粉丝id
primary key(follow_id,follower_id),
Foreign key (follow_id) references User_table(user_id),
Foreign key (follower_id) references User_table(user_id)
);
create table user_favourite   -- 用户点赞视频
(
fav_user_id bigint not null, -- 点赞行为的用户id
fav_video_id bigint not null, -- 用户点赞的视频id
primary key(fav_user_id,fav_video_id),
Foreign key (fav_user_id) references User_table(user_id),
Foreign key (fav_video_id) references Video(video_id)
);
create table comment  -- 评论表
( 
comment_id bigint not null auto_increment primary key,-- 视频的评论id
com_video_id bigint not null, -- 评论的视频id
com_user_id bigint not null,  -- 评论用户id
content text not null, -- 评论内容
create_date timestamp default now(),  -- 评论发布日期
Foreign key (com_video_id) references Video(video_id),
Foreign key(com_user_id) references User_table(user_id)
);



-- 用户点赞同步更新视频点赞数触发器
create trigger fix_favourite1
after insert on user_favourite
for each row
begin 
	update video set favourite_count=favourite_count+1 
	where video_id=new.fav_video_id;
end;



-- 用户取消点赞同步更新视频点赞数触发器
create trigger fix_favourite2
after delete on user_favourite
for each row
begin 
	update video set favourite_count=favourite_count-1 
	where video_id=old.fav_video_id;
end;



-- 用户关注同步更新触发器
create trigger fix_follow1
after insert on user_follow
for each row
begin 
	update user_table set follow_count=follow_count+1 
	where user_id=new.follower_id;
	update user_table set follower_count=follower_count+1 
	where user_id=new.follow_id;
end;



-- 用户取消关注同步更新触发器
create trigger fix_follow2
after delete on user_follow
for each row
begin 
	update user_table set follow_count=follow_count-1 
	where user_id=old.follower_id;
	update user_table set follower_count=follower_count-1 
	where user_id=old.follow_id;
end;



-- 用户评论视频同步更新触发器
create trigger fix_comment1
after insert on comment
for each row
begin 
	update video set comment_count=comment_count+1
	where video_id=new.com_video_id;
end;



-- 用户评论视频同步更新触发器
create trigger fix_comment2
after delete on comment
for each row
begin 
	update video set comment_count=comment_count-1
	where video_id=old.com_video_id;
end;

insert into user_table values(1,'Test001','test001',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/1.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(2,'Test002','test002',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/2.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(3,'Test003','test003',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/3.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(4,'Test004','test004',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/4.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(5,'Test005','test005',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/5.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(6,'Test006','test006',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/6.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(7,'Test007','test007',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/7.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(8,'Test008','test008',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/8.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(9,'Test009','test009',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/9.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(10,'Test010','test010',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/10.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(11,'Test011','test011',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/11.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(12,'Test012','test012',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/12.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(13,'Test013','test013',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/13.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(14,'Test014','test014',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/14.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(15,'Test015','test015',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/15.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(16,'Test016','test016',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/16.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(17,'Test017','test017',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/17.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(18,'Test018','test018',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/18.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(19,'Test019','test019',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/19.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(20,'Test020','test020',0,0,'欢迎使用抖声APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/20.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into video values
(0,
1,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/1.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/1.jpg',
'这里是Test001投稿的第一个视频,全部视频里第1个投稿的视频',
0,
0,
1653407743);
insert into video values
(0,
1,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/2.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/2.jpg',
'这里是用户Test001投稿的第二个视频,全部视频里第2个投稿的视频',
0,
0,
1653404743);
insert into video values
(0,
1,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/3.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/3.jpg',
'这里是用户Test001投稿的第三个视频,全部视频里第3个投稿的视频',
0,
0,
1653409773);
insert into video values
(0,
1,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/4.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/4.jpg',
'这里是用户Test001投稿的第四个视频,全部视频里第4个投稿的视频',
0,
0,
1653417743);
insert into video values
(0,
1,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/5.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/5.jpg',
'这里是用户Test001投稿的第五个视频,全部视频里第5个投稿的视频',
0,
0,
1653415957);
insert into video values
(0,
1,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/6.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/6.jpg',
'这里是用户Test001投稿的第六个视频,全部视频里第6个投稿的视频',
0,
0,
1653450262);
insert into video values
(0,
2,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/7.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/7.jpg',
'这里是用户Test002投稿的第一个视频,全部视频里第7个投稿的视频',
0,
0,
1653465173);
insert into video values
(0,
2,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/8.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/8.jpg',
'这里是用户Test002投稿的第二个视频,全部视频里第8个投稿的视频',
0,
0,
1653459399);
insert into video values
(0,
2,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/9.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/9.jpg',
'这里是用户Test002投稿的第三个视频,全部视频里第9个投稿的视频',
0,
0,
1653437923);
insert into video values
(0,
2,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/10.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/10.jpg',
'这里是用户Test002投稿的第四个视频,全部视频里第10个投稿的视频',
0,
0,
1653471845);
insert into video values
(0,
2,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/11.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/11.jpg',
'这里是用户Test002投稿的第五个视频,全部视频里第11个投稿的视频',
0,
0,
1653470548);
insert into video values
(0,
3,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/12.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/12.jpg',
'这里是用户Test003投稿的第一个视频,全部视频里第12个投稿的视频',
0,
0,
1653448790);
insert into video values
(0,
3,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/13.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/13.jpg',
'这里是用户Test003投稿的第二个视频,全部视频里第13个投稿的视频',
0,
0,
1653459468);
insert into video values
(0,
3,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/14.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/14.jpg',
'这里是用户Test003投稿的第三个视频,全部视频里第14个投稿的视频',
0,
0,
1653422382);
insert into video values
(0,
3,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/15.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/15.jpg',
'这里是用户Test003投稿的第四个视频,全部视频里第15个投稿的视频',
0,
0,
1653498486);
insert into video values
(0,
4,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/16.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/16.jpg',
'这里是用户Test004投稿的第一个视频,全部视频里第16个投稿的视频',
0,
0,
1653450279);
insert into video values
(0,
4,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/17.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/17.jpg',
'这里是用户Test004投稿的第二个视频,全部视频里第17个投稿的视频',
0,
0,
1653427155);
insert into video values
(0,
4,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/18.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/18.jpg',
'这里是用户Test001投稿的第三个视频,全部视频里第18个投稿的视频',
0,
0,
1653437833);
insert into video values
(0,
5,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/19.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/19.jpg',
'这里是用户Test005投稿的第一个视频,全部视频里第19个投稿的视频',
0,
0,
165349191);
insert into video values
(0,
5,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/20.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/20.jpg',
'这里是用户Test005投稿的第二个视频,全部视频里第20个投稿的视频',
0,
0,
1653459116);
insert into video values
(0,
5,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/21.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/21.jpg',
'这里是用户Test005投稿的第三个视频,全部视频里第21个投稿的视频',
0,
0,
1653434412);
insert into video values
(0,
6,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/22.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/22.jpg',
'这里是用户Test006投稿的第一个视频,全部视频里第22个投稿的视频',
0,
0,
1653489690);
insert into video values
(0,
6,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/23.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/23.jpg',
'这里是用户Test006投稿的第二个视频,全部视频里第23个投稿的视频',
0,
0,
1653411630);
insert into video values
(0,
6,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/24.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/24.jpg',
'这里是用户Test006投稿的第三个视频,全部视频里第24个投稿的视频',
0,
0,
1653494067);
insert into video values
(0,
7,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/25.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/25.jpg',
'这里是用户Test007投稿的第一个视频,全部视频里第25个投稿的视频',
0,
0,
1653443783);
insert into video values
(0,
7,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/26.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/26.jpg',
'这里是用户Test007投稿的第二个视频,全部视频里第26个投稿的视频',
0,
0,
1653429709);
insert into video values
(0,
7,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/27.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/27.jpg',
'这里是用户Test007投稿的第三个视频,全部视频里第27个投稿的视频',
0,
0,
1653490552);
insert into video values
(0,
8,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/28.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/28.jpg',
'这里是用户Test008投稿的第一个视频,全部视频里第28个投稿的视频',
0,
0,
1653431547);
insert into video values
(0,
8,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/29.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/29.jpg',
'这里是用户Test008投稿的第二个视频,全部视频里第29个投稿的视频',
0,
0,
1653455093);
insert into video values
(0,
9,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/30.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/30.jpg',
'这里是用户Test009投稿的第一个视频,全部视频里第30个投稿的视频',
0,
0,
1653478638);
insert into video values
(0,
9,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/31.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/31.jpg',
'这里是用户Test009投稿的第二个视频,全部视频里第31个投稿的视频',
0,
0,
1653410583);
insert into video values
(0,
10,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/32.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/32.jpg',
'这里是用户Test010投稿的第一个视频,全部视频里第32个投稿的视频',
0,
0,
1653451907);
insert into video values
(0,
10,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/33.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/33.jpg',
'这里是用户Test010投稿的第二个视频,全部视频里第33个投稿的视频',
0,
0,
1653464655);
insert into video values
(0,
11,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/34.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/34.jpg',
'这里是用户Test011投稿的第一个视频,全部视频里第34个投稿的视频',
0,
0,
1653428986);
insert into video values
(0,
11,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/35.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/35.jpg',
'这里是用户Test011投稿的第二个视频,全部视频里第35个投稿的视频',
0,
0,
1653465076);
insert into video values
(0,
12,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/36.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/36.jpg',
'这里是用户Test012投稿的第一个视频,全部视频里第36个投稿的视频',
0,
0,
16534838);
insert into video values
(0,
13,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/37.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/37.jpg',
'这里是用户Test013投稿的第一个视频,全部视频里第37个投稿的视频',
0,
0,
1653452622);
insert into video values
(0,
14,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/38.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/38.jpg',
'这里是用户Test014投稿的第一个视频,全部视频里第38个投稿的视频',
0,
0,
1653433644);
insert into video values
(0,
15,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/39.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/39.jpg',
'这里是用户Test015投稿的第一个视频,全部视频里第39个投稿的视频',
0,
0,
1653441388);
insert into video values
(0,
16,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/40.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/40.jpg',
'这里是用户Test016投稿的第一个视频,全部视频里第40个投稿的视频',
0,
0,
1653496915);
create view video_with_author
as
select *
from video v left join user_table ut 
on v.author_id =ut.user_id;
create view user_like_videoList
as
select *
from user_favourite uf left join video_with_author vwa
on uf.fav_video_id =vwa.video_id ;