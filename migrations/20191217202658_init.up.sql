CREATE TYPE gender_type AS ENUM (
  'male',
  'female',
  'other'
);

CREATE TYPE login_strategy AS ENUM (
  'local',
  'facebook',
  'twitter',
  'google'
);

CREATE TYPE post_type AS ENUM (
  'text',
  'video',
  'image',
  'audio',
  'link',
  'poll'
);

CREATE TYPE goal_type AS ENUM (
  'subscription_count',
  'money'
);

CREATE TYPE benefit_period AS ENUM (
  'one_time',
  'monthly'
);

CREATE TYPE card_brand_type AS ENUM (
  'jcb',
  'visa',
  'mastercard',
  'american_express',
  'diners_club'
);

CREATE TABLE users (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  username varchar UNIQUE,
  email varchar,
  password varchar,
  verified bool,
  social_id varchar,
  strategy login_strategy,
  avatar_image_id int8 REFERENCES file,
  gender gender_type,
  refresh_token varchar,
  created_at timestamptz,
  updated_at timestamptz
);

CREATE UNIQUE INDEX email_unique_idx ON users (LOWER(email));

CREATE UNIQUE INDEX username_unique_idx ON users (LOWER(username));

CREATE TABLE user_profile (
  first_name varchar,
  last_name varchar,
  phone_number varchar,
  bio text,
  updated_at timestamptz,
  user_id int8 PRIMARY KEY REFERENCES users
);

CREATE TABLE file (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  location varchar,
  size int8,
  file_name varchar,
  mime_type varchar,
  created_at timestamptz,
  updated_at timestamptz,
  user_id int8 REFERENCES users,
  meta jsonb
);

CREATE TABLE user_info (
  user_id int8 PRIMARY KEY REFERENCES users,
  last_logged_in_date timestamptz,
  last_logged_in_ip varchar,
  created_ip varchar
);

--TODO commissions,aka shop;
CREATE TABLE category (
  id int8 PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
  banner varchar,
  name varchar,
  description varchar,
  created_at timestamptz,
  updated_at timestamptz
);

CREATE TABLE creator_rank (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  fee float8,
  name varchar,
  importance int4
);

CREATE TABLE creator (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  name varchar,
  avatar int8 REFERENCES file,
  user_id int8 REFERENCES users,
  creator_rank_id int4 REFERENCES creator_rank ON DELETE SET NULL,
  is_company bool,
  updated_at timestamptz,
  created_at timestamptz
);

CREATE TABLE creator_profile (
  legal_first_name varchar,
  legal_last_name varchar,
  legal_company_name varchar,
  bio text,
  -- earning_visibility          bool,
  -- subscriber_count_visibility bool,
  creator_id int8 PRIMARY KEY REFERENCES creator ON DELETE CASCADE
);

CREATE TABLE project (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  name varchar,
  page_url varchar,
  description text,
  creating text,
  banner_image_id int8 REFERENCES file,
  avatar_image_id int8 REFERENCES file,
  creator_id int8 REFERENCES creator,
  category_id int4 REFERENCES category,
  setting jsonb,
  created_at timestamptz,
  updated_at timestamptz
);

CREATE TABLE project_info (
  project_id int8 PRIMARY KEY REFERENCES project
);

-- create or replace function create_project(_user_id int8, _name varchar, _description text,
--                                _cover int8, _category_id int4, _banner int8) returns RECORD as
-- $$
-- declare
--     _project_id int8;
--     _access_id  int8;
-- BEGIN
--     insert into project(name, description, banner_image_id, cover_image_id, category_id, user_id, updated_at,
--                         created_at)
--     VALUES (_name, _description, _banner, _cover, _category_id, _user_id, now(), now())
--     returning id into _project_id;
--     insert into tier(title, description, cover_image_id, price, project_id, created_at, updated_at)
--     values ('公共', '投稿を表示する相手がアカウント保持者または登録者以外にもされる', 1, 0, _project_id, now(), now());
--     select id into _access_id from access where site = true and name = 'owner';
--     INSERT INTO project_access(project_id, user_id, access_id) values (_project_id, _user_id, _access_id);
-- END
-- $$ language plpgsql;
CREATE TABLE access (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  name varchar,
  setting jsonb
);

-- CREATE TABLE project_access (
--   project_id int8 REFERENCES project ON DELETE CASCADE,
--   user_id int8 REFERENCES users ON DELETE CASCADE,
--   access_id int8 REFERENCES access,
--   PRIMARY KEY (project_id, user_id, access_id)
-- );
-- CREATE TABLE site_access (
--   user_id int8 REFERENCES users ON DELETE CASCADE,
--   access_id int8 REFERENCES access,
--   PRIMARY KEY (project_id, user_id, access_id)
-- );
CREATE TABLE creator_access (
  creator_id int8 REFERENCES creator ON DELETE CASCADE,
  user_id int8 REFERENCES users ON DELETE CASCADE,
  access_id int8 REFERENCES access,
  PRIMARY KEY (project_id, user_id, access_id)
);

CREATE TABLE tier (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  title varchar,
  description text,
  cover_image_id int8 REFERENCES file,
  price float8,
  project_id int8 REFERENCES project ON DELETE CASCADE,
  created_at timestamptz,
  updated_at timestamptz
);

CREATE TABLE subscription (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  user_id int8 REFERENCES users ON DELETE CASCADE,
  expiry_date timestamptz,
  tier_id int8 REFERENCES tier ON DELETE CASCADE,
  created_at timestamptz,
  stripe_subscription_id varchar,
  UNIQUE (user_id, tier_id)
);

CREATE TABLE post (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  title varchar,
  content text,
  mature bool,
  project_id int8 REFERENCES project ON DELETE CASCADE,
  submit_user_id int8 REFERENCES users ON DELETE SET NULL,
  created_at timestamptz,
  updated_at timestamptz
);

CREATE TABLE post_vote (
  user_id int8 REFERENCES users ON DELETE CASCADE,
  vote int4,
  post_id int8 REFERENCES post ON DELETE CASCADE,
  created_at timestamptz,
  PRIMARY KEY (user_id, post_id)
);

CREATE TABLE comment (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  user_id int8 REFERENCES users ON DELETE SET NULL,
  content text,
  parent_id int8 REFERENCES comment,
  parent_user_id int8 REFERENCES users,
  post_id int8 REFERENCES post ON DELETE CASCADE,
  created_at timestamptz,
  updated_at timestamptz
);

CREATE TABLE comment_vote (
  user_id int8 REFERENCES users ON DELETE CASCADE,
  vote int4,
  comment_id int8 REFERENCES COMMENT ON DELETE CASCADE,
  created_at timestamptz,
  PRIMARY KEY (user_id, comment_id)
);

CREATE TABLE post_tier (
  post_id int8 REFERENCES post ON DELETE CASCADE,
  tier_id int8 REFERENCES tier,
  PRIMARY KEY (post_id, tier_id)
);

CREATE TABLE post_text (
  post_id int8 PRIMARY KEY REFERENCES post ON DELETE CASCADE,
  word_count int8,
  char_count int8
);

CREATE TABLE tag (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  name varchar,
  created_at timestamptz
);

CREATE TABLE post_tag (
  post_id int8 REFERENCES post ON DELETE CASCADE,
  tag_id int8 REFERENCES tag ON DELETE CASCADE,
  PRIMARY KEY (tag_id, post_id)
);

-- TODO BELOW
CREATE TABLE post_poll_option (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  name varchar,
  post_id int8 REFERENCES post ON DELETE CASCADE
);

CREATE TABLE post_poll_answer (
  user_id int8 REFERENCES users,
  post_poll_option_id int8 REFERENCES post_poll_option ON DELETE CASCADE,
  PRIMARY KEY (user_id, post_poll_option_id)
);

CREATE TABLE post_file (
  post_id int8 REFERENCES post ON DELETE CASCADE,
  file_id int8 REFERENCES file,
  attachment bool,
  PRIMARY KEY (post_id, file_id)
);

CREATE TABLE post_link (
  post_id int8 REFERENCES post,
  link varchar
);

CREATE TABLE goal (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  description text,
  goal float8,
  TYPE goal_type,
  project_id int8 REFERENCES project ON DELETE CASCADE,
  created_at timestamptz,
  updated_at timestamptz
);

CREATE TABLE benefit (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  description text,
  TYPE benefit_period,
  created_at timestamptz,
  updated_at timestamptz
);

CREATE TABLE tier_benefit (
  tier_id int8 REFERENCES tier ON DELETE CASCADE,
  benefit_id int8 REFERENCES benefit ON DELETE CASCADE,
  PRIMARY KEY (tier_id, benefit_id)
);

CREATE TABLE payment_method (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  full_name varchar,
  created_at timestamptz,
  updated_at timestamptz,
  user_id int8 REFERENCES users ON DELETE CASCADE,
  default_method bool
);

CREATE TABLE stripe_card_payment_method (
  payment_method_id int8 PRIMARY KEY REFERENCES payment_method ON DELETE CASCADE,
  last_four int4,
  brand card_brand_type,
  expire_date timestamptz,
  stripe_payment_method_id varchar
);

CREATE TABLE stripe_bank_payment_method (
  payment_method_id int8 PRIMARY KEY REFERENCES payment_method ON DELETE CASCADE,
  last_four int4,
  stripe_payment_method_id varchar
);

CREATE TABLE payout_method (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  full_name varchar,
  created_at timestamptz,
  updated_at timestamptz,
  creator_id int8 REFERENCES creator ON DELETE CASCADE,
  default_method bool
);

-- CREATE TABLE stripe_card_payout_method (
--   payout_method_id int8 PRIMARY KEY REFERENCES payout_method ON DELETE CASCADE,
--   last_four int4,
--   brand card_brand_type,
--   expire_date timestamptz,
--   stripe_payout_method_id varchar
-- );

-- CREATE TABLE stripe_bank_payout_method (
--   payout_method_id int8 PRIMARY KEY REFERENCES payout_method ON DELETE CASCADE,
--   last_four int4,
--   stripe_payout_method_id varchar
-- );

CREATE TABLE TRANSACTION (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  amount float8,
  sender_id int8 REFERENCES users,
  receiver_project_id int8 REFERENCES project NOT NULL,
  created_at timestamptz
);

CREATE TABLE payout_transaction (
  payout_method_id int8 REFERENCES payout_method ON DELETE CASCADE,
  transaction_id int8 REFERENCES TRANSACTION,
  PRIMARY KEY (payout_method_id, transaction_id)
);

CREATE TABLE subscription_transaction (
  subscription_id int8 REFERENCES subscription ON DELETE CASCADE,
  transaction_id int8 REFERENCES TRANSACTION ON DELETE CASCADE,
  PRIMARY KEY (subscription_id, transaction_id)
);

CREATE TABLE donation_transaction (
  transaction_id int8 REFERENCES TRANSACTION PRIMARY KEY,
  message text,
  name varchar,
  private bool
);

CREATE TABLE message (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  sender_id REFERENCES users ON DELETE CASCADE,
  receiver_id REFERENCES users ON DELETE CASCADE,
  subject varchar,
  body text,
  created_at timestamptz
);

-- create table commission(
--
-- );
-- create or replace function exist update_post(_post_id int8, _user_id int8, _title varchar, _content text, _tiers int8[], _mature bool,
--                             _type post_type, _post jsonb, _tags int8[],_attachments int8[]) as
-- $$
-- DECLARE
--     i           jsonb;
--     _tag_id     int8;
--     _tier_id    int8;
--     _file_id    int8;
--     _exist      bool;
--     _project_id int8;
--     _file       int8;
--     _attachment int8;
-- BEGIN
--     select project_id into _project_id from tier where id = _tiers[1];
--     select exists(select ca.*
--                   from project_access ca
--                            inner join access a on a.id = ca.access_id and a.update_post = true
--                   where ca.user_id = _user_id
--                     and ca.project_id = _project_id)
--     into _exist;
--     IF NOT _exist THEN
--          raise exception '401';
--     END IF;
--     update post
--     set title      = _title,
--         content    = _content,
--         mature     = _mature,
--         updated_at = now()
--     where id = _post_id;
--     case
--         when _type = 'text'::post_type then
--             update post_text
--             set word_count = _post ->> 'word_count'::int8,
--                 char_count = _post ->> 'char_count'::int8
--             where post_id = _post_id;
--         when _type = 'audio'::post_type or _type = 'video'::post_type or _type = 'image'::post_type then
--             delete from post_file where _post_id = _post_id;
--             for _file in SELECT * FROM jsonb_array_elements(_post)
--                 LOOP
--                     if _file ->> 'id' is not null then
--                         Insert into post_file(post_id, file_id, attachment)
--                         VALUES (_post_id, _file ->> 'id'::int8, false);
--                     else
--                         INSERT INTO file(location, size, file_name, mime_type, created_at, updated_at, user_id, meta)
--                         VALUES (_file ->> 'location', _file ->> 'size'::int8, _file ->> 'file_name',
--                                 _file ->> 'mime_type', now(),
--                                 now(), _user_id, _file -> 'meta')
--                         RETURNING id into _file_id;
--                         INSERT INTO post_file(post_id, file_id, attachment)
--                         VALUES (_post_id, _file_id, false);
--                     end if;
--                 end loop;
--         when _type = 'link'::post_type then
--             update post_link set link = _post ->> 'link' where post_id = _post_id;
--         when _type = 'poll'::post_type then
--             for i IN SELECT * FROM jsonb_array_elements(_post)
--                 LOOP
--                     INSERT INTO post_poll_option(name, post_id) VALUES (i ->> 'option', _post_id);
--                 end loop;
--         end case;
--     delete from post_tag where post_id = _post_id;
--     for _tag_id IN _tags
--         LOOP
--             INSERT INTO post_tag(tag_id, post_id) VALUES (_tag_id, _post_id);
--         end loop;
--     delete from post_tier where post_id = _post_id;
--     for _tier_id IN _tiers
--         LOOP
--             INSERT INTO post_tier(tier_id, post_id) VALUES (_tier_id, _post_id);
--         end loop;
--     delete from post_file where post_id = _post_id and attachment = true;
--          for _attachment in _attachments
--             LOOP
--                INSERT INTO post_file(post_id, file_id, attachment)
--                 VALUES (_post_id, _attachment, true);
--             end loop;
-- END;
-- $$ language plpgsql;
-- create or replace function create_post(_user_id int8, _title varchar, _content text, _tiers int8[], _mature bool,
--                             _type post_type, _post jsonb, _tags int8[], _attachments int8[]) as
-- $$
-- declare
--     _post_id    int8;
--     i           jsonb;
--     _tag_id     int8;
--     _tier_id    int8;
--     _file_id    int8;
--     _exist      bool;
--     _creator_id int8;
--     _file       jsonb;
--     _attachment int8;
-- BEGIN
--     select creator_id into _creator_id from tier where id = _tiers[1];
--     select exists(select ca.*
--                   from creator_access ca
--                            inner join access a on a.id = ca.access_id and a.create_post = true
--                   where ca.user_id = _user_id
--                     and ca.creator_id = _creator_id)
--     into _exist;
--     IF NOT _exist THEN
--         raise exception '401';
--     END IF;
--     insert into post(title, content, mature, created_at, updated_at)
--     VALUES (_title, _content, _mature, now(), now())
--     returning id into _post_id;
--     case
--         when _type = 'text'::post_type then
--             INSERT INTO post_text(post_id, word_count, char_count)
--             VALUES (_post_id, _post ->> 'word_count'::int8, _post ->> 'char_count'::int8);
--         when _type = 'audio'::post_type or _type = 'video'::post_type or _type = 'image'::post_type then
--             for _file in SELECT * FROM jsonb_array_elements(_post)
--                 LOOP
--                     INSERT INTO file(location, size, file_name, mime_type, created_at, updated_at, user_id, meta)
--                     VALUES (_file ->> 'location', _file ->> 'size'::int8, _file ->> 'file_name', _file ->> 'mime_type',
--                             now(),
--                             now(), _user_id, _file -> 'meta')
--                     RETURNING id into _file_id;
--                     INSERT INTO post_file(post_id, file_id, attachment)
--                     VALUES (_post_id, _file_id, false);
--                 end loop;
--         when _type = 'link'::post_type then
--             Insert into post_link(post_id, link) VALUES (_post_id, _post ->> 'link');
--         when _type = 'poll'::post_type then
--             for i IN SELECT * FROM jsonb_array_elements(_post)
--                 LOOP
--                     INSERT INTO post_poll_option(name, post_id) VALUES (i ->> 'option', _post_id);
--                 end loop;
--         end case;
--         for _tag_id IN _tags
--         LOOP
--             INSERT INTO post_tag(tag_id, post_id) VALUES (_tag_id, _post_id);
--         end loop;
--         for _tier_id IN _tiers
--         LOOP
--             INSERT INTO post_tier(tier_id, post_id) VALUES (_tier_id, _post_id);
--         end loop;
--     for _attachment in _attachments
--             LOOP
--                INSERT INTO post_file(post_id, file_id, attachment)
--                 VALUES (_post_id, _attachment, true);
--         end loop;
-- END
-- $$ language plpgsql;
