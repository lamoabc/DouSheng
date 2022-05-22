use douyin;
create table User_table  -- 用户表
(
user_id bigint not null auto_increment primary key, -- 用户i
account_password varchar(32) not null,  -- 账号密码
user_name varchar(32) not null unique, -- 用户昵称
follow_count int default 0, -- 关注总数
follower_count int default 0 -- 粉丝总数
);
create table Video
(
id bigint not null auto_increment primary key,-- 视频唯一标识
author_id bigint not null, -- 作者用户id
play_url varchar(1000) not null,-- 视频播放地址
cover_url varchar(1000) not null,-- 视频封面地址
video_title varchar(50) not null,  -- 视频标题
favourite_count bigint default 0,  -- 视频的点赞总数
comment_count bigint default 0,  -- 视频的评论总数
upload_date timestamp default now(),  -- 视频发布日期
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
Foreign key (fav_video_id) references Video(id)
);
create table comment  -- 评论表
( 
comment_id bigint not null auto_increment primary key,-- 视频的评论id
com_video_id bigint not null, -- 评论的视频id
com_user_id bigint not null,  -- 评论用户id
content text not null, -- 评论内容
create_date timestamp default now(),  -- 评论发布日期
Foreign key (com_video_id) references Video(id),
Foreign key(com_user_id) references User_table(user_id)
);



-- 用户点赞同步更新视频点赞数触发器
create trigger fix_favourite1
after insert on user_favourite
for each row
begin 
	update video set favourite_count=favourite_count+1 
	where id=new.fav_video_id;
end



-- 用户取消点赞同步更新视频点赞数触发器
create trigger fix_favourite2
after delete on user_favourite
for each row
begin 
	update video set favourite_count=favourite_count-1 
	where id = old.fav_video_id;
end



-- 用户关注同步更新触发器
create trigger fix_follow1
after insert on user_follow
for each row
begin 
	update user_table set follow_count=follow_count+1 
	where user_id=new.follower_id;
	update user_table set follower_count=follower_count+1 
	where user_id=new.follow_id;
end



-- 用户取消关注同步更新触发器
create trigger fix_follow2
after delete on user_follow
for each row
begin 
	update user_table set follow_count=follow_count-1 
	where user_id=old.follower_id;
	update user_table set follower_count=follower_count-1 
	where user_id=old.follow_id;
end



-- 用户评论视频同步更新触发器
create trigger fix_comment1
after insert on comment
for each row
begin 
	update video set comment_count=comment_count+1
	where id=new.com_video_id;
end



-- 用户评论视频同步更新触发器
create trigger fix_comment2
after delete on comment
for each row
begin 
	update video set comment_count=comment_count-1
	where id=old.com_video_id;
end
