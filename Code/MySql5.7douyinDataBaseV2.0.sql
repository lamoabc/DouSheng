create table user_table  -- �û���
(
user_id bigint not null auto_increment primary key, -- �û�id
user_name varchar(32) not null unique, -- �û��ǳ�
account_password varchar(32) not null,  -- �˺�����
follow_count int default 0, -- ��ע����
follower_count int default 0, -- ��˿����
signature text, -- ����ǩ��
avatar varchar(200), -- ͷ������
background_image varchar(200)  -- ����ͼ����
);
create table video
(
video_id bigint not null auto_increment primary key,-- ��ƵΨһ��ʶ
author_id bigint not null, -- �����û�id
play_url varchar(200) not null,-- ��Ƶ���ŵ�ַ
cover_url varchar(200) not null,-- ��Ƶ�����ַ
video_title varchar(50) not null,  -- ��Ƶ����
favourite_count bigint default 0,  -- ��Ƶ�ĵ�������
comment_count bigint default 0,  -- ��Ƶ����������
upload_date bigint default 0,  -- ��Ƶ����ʱ���
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
Foreign key (fav_video_id) references Video(video_id)
);
create table comment  -- ���۱�
( 
comment_id bigint not null auto_increment primary key,-- ��Ƶ������id
com_video_id bigint not null, -- ���۵���Ƶid
com_user_id bigint not null,  -- �����û�id
content text not null, -- ��������
create_date timestamp default now(),  -- ���۷�������
Foreign key (com_video_id) references Video(video_id),
Foreign key(com_user_id) references User_table(user_id)
);



-- �û�����ͬ��������Ƶ������������
create trigger fix_favourite1
after insert on user_favourite
for each row
begin 
	update video set favourite_count=favourite_count+1 
	where id=new.fav_video_id;
end;



-- �û�ȡ������ͬ��������Ƶ������������
create trigger fix_favourite2
after delete on user_favourite
for each row
begin 
	update video set favourite_count=favourite_count-1 
	where id=old.fav_video_id;
end;



-- �û���עͬ�����´�����
create trigger fix_follow1
after insert on user_follow
for each row
begin 
	update user_table set follow_count=follow_count+1 
	where user_id=new.follower_id;
	update user_table set follower_count=follower_count+1 
	where user_id=new.follow_id;
end;



-- �û�ȡ����עͬ�����´�����
create trigger fix_follow2
after delete on user_follow
for each row
begin 
	update user_table set follow_count=follow_count-1 
	where user_id=old.follower_id;
	update user_table set follower_count=follower_count-1 
	where user_id=old.follow_id;
end;



-- �û�������Ƶͬ�����´�����
create trigger fix_comment1
after insert on comment
for each row
begin 
	update video set comment_count=comment_count+1
	where id=new.com_video_id;
end;



-- �û�������Ƶͬ�����´�����
create trigger fix_comment2
after delete on comment
for each row
begin 
	update video set comment_count=comment_count-1
	where id=old.com_video_id;
end;

insert into user_table values(1,'Test001','test001',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/1.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(2,'Test002','test002',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/2.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(3,'Test003','test003',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/3.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(4,'Test004','test004',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/4.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(5,'Test005','test005',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/5.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(6,'Test006','test006',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/6.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(7,'Test007','test007',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/7.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(8,'Test008','test008',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/8.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(9,'Test009','test009',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/9.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(10,'Test010','test010',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/10.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(11,'Test011','test011',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/11.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(12,'Test012','test012',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/12.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(13,'Test013','test013',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/13.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(14,'Test014','test014',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/14.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(15,'Test015','test015',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/15.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(16,'Test016','test016',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/16.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(17,'Test017','test017',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/17.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(18,'Test018','test018',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/18.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(19,'Test019','test019',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/19.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into user_table values(20,'Test020','test020',0,0,'��ӭʹ�ö���APP','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/20.jpeg','https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png');
insert into video values
(0,
1,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/1.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/1.jpg',
'������Test001Ͷ��ĵ�һ����Ƶ,ȫ����Ƶ���1��Ͷ�����Ƶ',
0,
0,
1653407743);
insert into video values
(0,
1,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/2.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/2.jpg',
'�������û�Test001Ͷ��ĵڶ�����Ƶ,ȫ����Ƶ���2��Ͷ�����Ƶ',
0,
0,
1653404743);
insert into video values
(0,
1,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/3.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/3.jpg',
'�������û�Test001Ͷ��ĵ�������Ƶ,ȫ����Ƶ���3��Ͷ�����Ƶ',
0,
0,
1653409773);
insert into video values
(0,
1,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/4.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/4.jpg',
'�������û�Test001Ͷ��ĵ��ĸ���Ƶ,ȫ����Ƶ���4��Ͷ�����Ƶ',
0,
0,
1653417743);
insert into video values
(0,
1,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/5.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/5.jpg',
'�������û�Test001Ͷ��ĵ������Ƶ,ȫ����Ƶ���5��Ͷ�����Ƶ',
0,
0,
1653415957);
insert into video values
(0,
1,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/6.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/6.jpg',
'�������û�Test001Ͷ��ĵ�������Ƶ,ȫ����Ƶ���6��Ͷ�����Ƶ',
0,
0,
1653450262);
insert into video values
(0,
2,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/7.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/7.jpg',
'�������û�Test002Ͷ��ĵ�һ����Ƶ,ȫ����Ƶ���7��Ͷ�����Ƶ',
0,
0,
1653465173);
insert into video values
(0,
2,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/8.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/8.jpg',
'�������û�Test002Ͷ��ĵڶ�����Ƶ,ȫ����Ƶ���8��Ͷ�����Ƶ',
0,
0,
1653459399);
insert into video values
(0,
2,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/9.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/9.jpg',
'�������û�Test002Ͷ��ĵ�������Ƶ,ȫ����Ƶ���9��Ͷ�����Ƶ',
0,
0,
1653437923);
insert into video values
(0,
2,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/10.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/10.jpg',
'�������û�Test002Ͷ��ĵ��ĸ���Ƶ,ȫ����Ƶ���10��Ͷ�����Ƶ',
0,
0,
1653471845);
insert into video values
(0,
2,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/11.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/11.jpg',
'�������û�Test002Ͷ��ĵ������Ƶ,ȫ����Ƶ���11��Ͷ�����Ƶ',
0,
0,
1653470548);
insert into video values
(0,
3,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/12.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/12.jpg',
'�������û�Test003Ͷ��ĵ�һ����Ƶ,ȫ����Ƶ���12��Ͷ�����Ƶ',
0,
0,
1653448790);
insert into video values
(0,
3,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/13.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/13.jpg',
'�������û�Test003Ͷ��ĵڶ�����Ƶ,ȫ����Ƶ���13��Ͷ�����Ƶ',
0,
0,
1653459468);
insert into video values
(0,
3,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/14.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/14.jpg',
'�������û�Test003Ͷ��ĵ�������Ƶ,ȫ����Ƶ���14��Ͷ�����Ƶ',
0,
0,
1653422382);
insert into video values
(0,
3,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/15.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/15.jpg',
'�������û�Test003Ͷ��ĵ��ĸ���Ƶ,ȫ����Ƶ���15��Ͷ�����Ƶ',
0,
0,
1653498486);
insert into video values
(0,
4,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/16.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/16.jpg',
'�������û�Test004Ͷ��ĵ�һ����Ƶ,ȫ����Ƶ���16��Ͷ�����Ƶ',
0,
0,
1653450279);
insert into video values
(0,
4,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/17.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/17.jpg',
'�������û�Test004Ͷ��ĵڶ�����Ƶ,ȫ����Ƶ���17��Ͷ�����Ƶ',
0,
0,
1653427155);
insert into video values
(0,
4,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/18.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/18.jpg',
'�������û�Test001Ͷ��ĵ�������Ƶ,ȫ����Ƶ���18��Ͷ�����Ƶ',
0,
0,
1653437833);
insert into video values
(0,
5,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/19.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/19.jpg',
'�������û�Test005Ͷ��ĵ�һ����Ƶ,ȫ����Ƶ���19��Ͷ�����Ƶ',
0,
0,
165349191);
insert into video values
(0,
5,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/20.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/20.jpg',
'�������û�Test005Ͷ��ĵڶ�����Ƶ,ȫ����Ƶ���20��Ͷ�����Ƶ',
0,
0,
1653459116);
insert into video values
(0,
5,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/21.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/21.jpg',
'�������û�Test005Ͷ��ĵ�������Ƶ,ȫ����Ƶ���21��Ͷ�����Ƶ',
0,
0,
1653434412);
insert into video values
(0,
6,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/22.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/22.jpg',
'�������û�Test006Ͷ��ĵ�һ����Ƶ,ȫ����Ƶ���22��Ͷ�����Ƶ',
0,
0,
1653489690);
insert into video values
(0,
6,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/23.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/23.jpg',
'�������û�Test006Ͷ��ĵڶ�����Ƶ,ȫ����Ƶ���23��Ͷ�����Ƶ',
0,
0,
1653411630);
insert into video values
(0,
6,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/24.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/24.jpg',
'�������û�Test006Ͷ��ĵ�������Ƶ,ȫ����Ƶ���24��Ͷ�����Ƶ',
0,
0,
1653494067);
insert into video values
(0,
7,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/25.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/25.jpg',
'�������û�Test007Ͷ��ĵ�һ����Ƶ,ȫ����Ƶ���25��Ͷ�����Ƶ',
0,
0,
1653443783);
insert into video values
(0,
7,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/26.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/26.jpg',
'�������û�Test007Ͷ��ĵڶ�����Ƶ,ȫ����Ƶ���26��Ͷ�����Ƶ',
0,
0,
1653429709);
insert into video values
(0,
7,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/27.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/27.jpg',
'�������û�Test007Ͷ��ĵ�������Ƶ,ȫ����Ƶ���27��Ͷ�����Ƶ',
0,
0,
1653490552);
insert into video values
(0,
8,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/28.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/28.jpg',
'�������û�Test008Ͷ��ĵ�һ����Ƶ,ȫ����Ƶ���28��Ͷ�����Ƶ',
0,
0,
1653431547);
insert into video values
(0,
8,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/29.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/29.jpg',
'�������û�Test008Ͷ��ĵڶ�����Ƶ,ȫ����Ƶ���29��Ͷ�����Ƶ',
0,
0,
1653455093);
insert into video values
(0,
9,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/30.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/30.jpg',
'�������û�Test009Ͷ��ĵ�һ����Ƶ,ȫ����Ƶ���30��Ͷ�����Ƶ',
0,
0,
1653478638);
insert into video values
(0,
9,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/31.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/31.jpg',
'�������û�Test009Ͷ��ĵڶ�����Ƶ,ȫ����Ƶ���31��Ͷ�����Ƶ',
0,
0,
1653410583);
insert into video values
(0,
10,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/32.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/32.jpg',
'�������û�Test010Ͷ��ĵ�һ����Ƶ,ȫ����Ƶ���32��Ͷ�����Ƶ',
0,
0,
1653451907);
insert into video values
(0,
10,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/33.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/33.jpg',
'�������û�Test010Ͷ��ĵڶ�����Ƶ,ȫ����Ƶ���33��Ͷ�����Ƶ',
0,
0,
1653464655);
insert into video values
(0,
11,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/34.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/34.jpg',
'�������û�Test011Ͷ��ĵ�һ����Ƶ,ȫ����Ƶ���34��Ͷ�����Ƶ',
0,
0,
1653428986);
insert into video values
(0,
11,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/35.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/35.jpg',
'�������û�Test011Ͷ��ĵڶ�����Ƶ,ȫ����Ƶ���35��Ͷ�����Ƶ',
0,
0,
1653465076);
insert into video values
(0,
12,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/36.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/36.jpg',
'�������û�Test012Ͷ��ĵ�һ����Ƶ,ȫ����Ƶ���36��Ͷ�����Ƶ',
0,
0,
16534838);
insert into video values
(0,
13,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/37.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/37.jpg',
'�������û�Test013Ͷ��ĵ�һ����Ƶ,ȫ����Ƶ���37��Ͷ�����Ƶ',
0,
0,
1653452622);
insert into video values
(0,
14,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/38.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/38.jpg',
'�������û�Test014Ͷ��ĵ�һ����Ƶ,ȫ����Ƶ���38��Ͷ�����Ƶ',
0,
0,
1653433644);
insert into video values
(0,
15,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/39.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/39.jpg',
'�������û�Test015Ͷ��ĵ�һ����Ƶ,ȫ����Ƶ���39��Ͷ�����Ƶ',
0,
0,
1653441388);
insert into video values
(0,
16,
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20address/40.mp4',
'https://yygh-lamo.oss-cn-beijing.aliyuncs.com/Video%20cover/40.jpg',
'�������û�Test016Ͷ��ĵ�һ����Ƶ,ȫ����Ƶ���40��Ͷ�����Ƶ',
0,
0,
1653496915);