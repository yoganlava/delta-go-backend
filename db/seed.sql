INSERT INTO creator_rank (fee, name, importance)
  VALUES (0.05, 'normal', 1), (0.02, 'sponsor', 2), (0, 'owner', 2);

INSERT INTO users (username, email, PASSWORD, verified, social_id, avatar)
  VALUES ('admin', 'arif.iwamoto@gmail.com', '$2y$10$1M/LhZ.nIWp1xGuiwnYcfut9pMON2lIfI3W7yGAtosw9oCqJPsfpG', TRUE, NULL, NULL);

INSERT INTO file (LOCATION, size, file_name, mime_type, created_at, updated_at, user_id)
  VALUES ('/file/default_profile.png', 1000, 'default_profile.png', 'image/png', now(), now(), 1),
  ('/file/default_banner.png', 1000, 'default_banner.png', 'image/png', now(), now(), 1);

UPDATE
  users
SET
  avatar = '/file/default_profile.png'
WHERE
  id = 1;

INSERT INTO creator (name, avatar, user_id, creator_rank_id, is_company, created_at, updated_at)
  VALUES ('オンジン', '/file/default_profile.png', 1, 3, TRUE, now(), now());

-- INSERT into creator (name,avatar,user_id,creator_rank_id,is_company,created_at,updated_at)
-- values ('オンジン','/file/default_profile.png',1,3,true,now(),now());
INSERT INTO category (banner, name, description, created_at, updated_at)
  VALUES ('', '音楽', '音楽・歌関連のものを作ってる', now(), now()), ('', '絵', '音楽・歌関連のものを作ってる', now(), now());

INSERT into project (name,page_url,description, banner, creator_id,category_id,setting,created_at,updated_at) 
values('オンジン','onijn','オンジン公式プロジェクトページ、オンジン株式会社ではonjin.jpにおける管理！','/file/default_banner.png',1,1,{"show_total_earning": true},now(),now());
