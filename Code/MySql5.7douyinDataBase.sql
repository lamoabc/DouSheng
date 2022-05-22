use douyin;
create table User_table  -- �û���
(
user_id bigint not null auto_increment primary key, -- �û�i
account_password varchar(32) not null,  -- �˺�����
user_name varchar(32) not null unique, -- �û��ǳ�
follow_count int default 0, -- ��ע����
follower_count int default 0 -- ��˿����
);
create table Video
(
id bigint not null auto_increment primary key,-- ��ƵΨһ��ʶ
author_id bigint not null, -- �����û�id
play_url varchar(1000) not null,-- ��Ƶ���ŵ�ַ
cover_url varchar(1000) not null,-- ��Ƶ�����ַ
video_title varchar(50) not null,  -- ��Ƶ����
favourite_count bigint default 0,  -- ��Ƶ�ĵ�������
comment_count bigint default 0,  -- ��Ƶ����������
upload_date timestamp default now(),  -- ��Ƶ��������
FOREIGN KEY(author_id) REFERENCES User_table(user_id)
);
create table user_follow  -- �û���ע��
(
follow_id bigint not null,  -- �û���id������ǰ�û�
follower_id bigint not null,  -- ����û��ķ�˿id
primary key(follow_id,follower_id),
Foreign key (follow_id) references User_table(user_id),
Foreign key (follower_id) references User_table(user_id)
);
create table user_favourite   -- �û�������Ƶ
(
fav_user_id bigint not null, -- ������Ϊ���û�id
fav_video_id bigint not null, -- �û����޵���Ƶid
primary key(fav_user_id,fav_video_id),
Foreign key (fav_user_id) references User_table(user_id),
Foreign key (fav_video_id) references Video(id)
);
create table comment  -- ���۱�
( 
comment_id bigint not null auto_increment primary key,-- ��Ƶ������id
com_video_id bigint not null, -- ���۵���Ƶid
com_user_id bigint not null,  -- �����û�id
content text not null, -- ��������
create_date timestamp default now(),  -- ���۷�������
Foreign key (com_video_id) references Video(id),
Foreign key(com_user_id) references User_table(user_id)
);



-- �û�����ͬ��������Ƶ������������
create trigger fix_favourite1
after insert on user_favourite
for each row
begin 
	update video set favourite_count=favourite_count+1 
	where id=new.fav_video_id;
end



-- �û�ȡ������ͬ��������Ƶ������������
create trigger fix_favourite2
after delete on user_favourite
for each row
begin 
	update video set favourite_count=favourite_count-1 
	where id = old.fav_video_id;
end



-- �û���עͬ�����´�����
create trigger fix_follow1
after insert on user_follow
for each row
begin 
	update user_table set follow_count=follow_count+1 
	where user_id=new.follower_id;
	update user_table set follower_count=follower_count+1 
	where user_id=new.follow_id;
end



-- �û�ȡ����עͬ�����´�����
create trigger fix_follow2
after delete on user_follow
for each row
begin 
	update user_table set follow_count=follow_count-1 
	where user_id=old.follower_id;
	update user_table set follower_count=follower_count-1 
	where user_id=old.follow_id;
end



-- �û�������Ƶͬ�����´�����
create trigger fix_comment1
after insert on comment
for each row
begin 
	update video set comment_count=comment_count+1
	where id=new.com_video_id;
end



-- �û�������Ƶͬ�����´�����
create trigger fix_comment2
after delete on comment
for each row
begin 
	update video set comment_count=comment_count-1
	where id=old.com_video_id;
end
