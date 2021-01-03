INSERT INTO creator_rank (fee, name, importance)
  VALUES (0.05, 'normal', 1), (0.02, 'sponsor', 2),(0, 'owner', 2);

INSERT INTO users (username, email, PASSWORD, verified, social_id, avatar)
  VALUES ('admin', 'arif.iwamoto@gmail.com', '$2y$10$1M/LhZ.nIWp1xGuiwnYcfut9pMON2lIfI3W7yGAtosw9oCqJPsfpG', TRUE, NULL,NULL);

insert into file (location, size, file_name, mime_type, created_at,updated_at,user_id)
values ('/file/default_profile.png',1000,'default_profile.png','image/png',now(),now(),1);

UPDATE users set avatar = '/file/default_profile.png' where id = 1;

INSERT into creator (name,avatar,user_id,creator_rank_id,is_company,created_at,updated_at)
values ('オンジン','/file/default_profile.png',1,3,true,now(),now());

INSERT into creator (name,avatar,user_id,creator_rank_id,is_company,created_at,updated_at)
values ('オンジン','/file/default_profile.png',1,3,true,now(),now());
