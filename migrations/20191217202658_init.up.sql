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
  avatar varchar,
  gender gender_type,
  refresh_token varchar,
  created_at timestamptz,
  updated_at timestamptz
);


CREATE TABLE file (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  location varchar UNIQUE,
  size int8,
  file_name varchar,
  mime_type varchar,
  created_at timestamptz,
  updated_at timestamptz,
  user_id int8,
  meta jsonb
);






CREATE TABLE user_profile (
  first_name varchar,
  last_name varchar,
  phone_number varchar,
  bio text,
  updated_at timestamptz,
  user_id int8 PRIMARY KEY 
);

CREATE TABLE user_info (
  user_id int8 PRIMARY KEY,
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
  avatar varchar REFERENCES file (LOCATION),
  user_id int8 REFERENCES users,
  creator_rank_id int4 ,
  is_company bool,
  stripe_account_id varchar,
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
  banner varchar REFERENCES file (LOCATION),
  avatar varchar REFERENCES file (LOCATION),
  creator_id int8 REFERENCES creator,
  category_id int4 REFERENCES category,
  setting jsonb,
  created_at timestamptz,
  updated_at timestamptz
);

CREATE TABLE project_info (
  project_id int8 PRIMARY KEY REFERENCES project
);


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
  PRIMARY KEY (creator_id, user_id, access_id)
);

CREATE TABLE tier (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  title varchar,
  description text,
  cover varchar REFERENCES file (LOCATION),
  price float8,
  project_id int8 REFERENCES project ON DELETE CASCADE,
  created_at timestamptz,
  updated_at timestamptz
);

CREATE TABLE subscription (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  user_id int8,
  expiry_date timestamptz,
  tier_id int8,
  created_at timestamptz,
  will_renew bool,
  meta jsonb,
  UNIQUE (user_id, tier_id)
);

CREATE TABLE post (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  title varchar,
  content text,
  mature bool,
  project_id int8,
  submit_user_id int8,
  created_at timestamptz,
  updated_at timestamptz,
  min_price float8
);

CREATE TABLE post_vote (
  user_id int8,
  vote int4,
  post_id int8,
  created_at timestamptz,
  PRIMARY KEY (user_id, post_id)
);

CREATE TABLE comment (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  user_id int8 ,
  content text,
  parent_id int8,
  post_id int8,
  created_at timestamptz,
  updated_at timestamptz
);

CREATE TABLE comment_vote (
  user_id int8,
  vote int4,
  comment_id int8 ,
  created_at timestamptz,
  PRIMARY KEY (user_id, comment_id)
);



CREATE TABLE tag (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  name varchar,
  created_at timestamptz
);

CREATE TABLE post_tag (
  post_id int8 ,
  tag_id int8 ,
  PRIMARY KEY (tag_id, post_id)
);

-- TODO BELOW
CREATE TABLE post_poll_option (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  name varchar,
  post_id int8
);
-- CREATE TABLE post_tier (
--   post_id int8 ,
--   tier_id int8,
--   PRIMARY KEY (post_id, tier_id)
-- );

CREATE TABLE post_text (
  post_id int8 PRIMARY KEY,
  word_count int8,
  char_count int8
);
CREATE TABLE post_poll_answer (
  user_id int8,
  post_poll_option_id int8 ,
  PRIMARY KEY (user_id, post_poll_option_id)
);

CREATE TABLE post_file (
  post_id int8,
  file_id int8 ,
  attachment bool,
  PRIMARY KEY (post_id, file_id)
);

CREATE TABLE post_link (
  post_id int8 primary key,
  link varchar
);

CREATE TABLE goal (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  description text,
  goal float8,
  TYPE goal_type,
  project_id int8,
  created_at timestamptz,
  updated_at timestamptz
);

CREATE TABLE benefit (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  description text,
  type benefit_period,
  created_at timestamptz,
  updated_at timestamptz
);

CREATE TABLE tier_benefit (
  tier_id int8 ,
  benefit_id int8,
  PRIMARY KEY (tier_id, benefit_id)
);

CREATE TABLE payment_method (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  full_name varchar,
  created_at timestamptz,
  updated_at timestamptz,
  user_id int8,
  default_method bool
);

CREATE TABLE stripe_card_payment_method (
  payment_method_id int8 PRIMARY KEY,
  last_four int4,
  brand card_brand_type,
  expire_date timestamptz,
  stripe_payment_method_id varchar
);

-- CREATE TABLE stripe_bank_payment_method (
--   payment_method_id int8 PRIMARY KEY REFERENCES payment_method ON DELETE CASCADE,
--   last_four int4,
--   stripe_payment_method_id varchar
-- );

CREATE TABLE payout_method (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  full_name varchar,
  created_at timestamptz,
  updated_at timestamptz,
  creator_id int8,
  default_method bool
);

CREATE TABLE paypal_payout_method (
    payout_method_id int8 PRIMARY KEY,
    paypal_email varchar
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
  sender_id int8,
  receiver_project_id int8,
  created_at timestamptz
);

CREATE TABLE payout_transaction (
  payout_method_id int8,
  transaction_id int8,
  PRIMARY KEY (payout_method_id, transaction_id)
);

CREATE TABLE subscription_transaction (
  subscription_id int8,
  transaction_id int8,
  PRIMARY KEY (subscription_id, transaction_id)
);

CREATE TABLE donation_transaction (
  transaction_id int8 PRIMARY KEY,
  message text,
  name varchar,
  private bool
);

CREATE TABLE message (
  id int8 GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  sender_id int8,
  receiver_id int8,
  subject varchar,
  body text,
  created_at timestamptz
);

Alter table post add constraint post_user_fk foreign key (submit_user_id) references users (id) on delete set null;
Alter table post add constraint post_project_fk foreign key (project_id) references project (id) on delete cascade;

Alter table comment add constraint comment_user_fk foreign key (user_id) references users (id) on delete cascade;
Alter table comment add constraint comment_parent_fk foreign key (parent_id) references comment (id) on delete cascade;
Alter table comment add constraint comment_post_fk foreign key (post_id) references post (id) on delete cascade;

Alter table comment_vote add constraint comment_vote_comment_fk foreign key (comment_id) references comment (id) on delete cascade;
Alter table comment_vote add constraint comment_vote_user_fk foreign key (user_id) references users (id) on delete cascade;

Alter table post_link add constraint post_link_post_fk foreign key (post_id) references post (id) on delete cascade;

Alter table post_text add constraint post_text_post_fk foreign key (post_id) references post (id) on delete cascade;

-- Alter table post_tier add constraint post_tier_tier_fk foreign key (tier_id) references tier (id) on delete set null;
-- Alter table post_tier add constraint post_tier_post_fk foreign key (post_id) references post (id) on delete cascade;

Alter table post_vote add constraint post_vote_post_fk foreign key (post_id) references post (id) on delete cascade;
Alter table post_vote add constraint post_vote_user_fk foreign key (user_id) references users (id) on delete set null;

Alter table subscription add constraint subscription_tier_fk foreign key (tier_id) references tier (id) on delete null;
ALTER TABLE subscription ADD CONSTRAINT subscription_user_fk foreign key (user_id) references users (id) on delete null;
ALTER TABLE post_tag ADD CONSTRAINT post_tag_post_fk FOREIGN KEY (post_id) REFERENCES post (id) on delete cascade;
ALTER TABLE post_tag ADD CONSTRAINT post_tag_tag_fk FOREIGN KEY (tag_id) REFERENCES tag (id) on delete cascade;

ALTER TABLE post_file ADD CONSTRAINT post_file_post_fk FOREIGN KEY (post_id) REFERENCES post (id) on delete cascade;
ALTER TABLE post_file ADD CONSTRAINT post_file_file_fk FOREIGN KEY (file_id) REFERENCES file (id) on delete cascade;

ALTER TABLE post_poll_option ADD CONSTRAINT post_poll_option_post_fk FOREIGN KEY (post_id) REFERENCES post (id) on delete cascade;

ALTER TABLE post_poll_answer ADD CONSTRAINT post_poll_answer_user_fk FOREIGN KEY (user_id) REFERENCES users (id) on delete cascade;
ALTER TABLE post_poll_answer ADD CONSTRAINT post_poll_answer_option_fk FOREIGN KEY (post_poll_option_id) REFERENCES post_poll_option (id) on delete cascade;

ALTER TABLE tier_benefit ADD CONSTRAINT tier_benefit_benefit_fk FOREIGN KEY (benefit_id) REFERENCES benefit (id) on delete cascade;
ALTER TABLE tier_benefit ADD CONSTRAINT tier_benefit_tier_fk FOREIGN KEY (tier_id) REFERENCES tier (id) on delete cascade;

ALTER TABLE users ADD CONSTRAINT user_avatar_fk FOREIGN KEY (avatar) REFERENCES file (location);
ALTER TABLE user_profile ADD CONSTRAINT user_profile_fk FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE user_info ADD CONSTRAINT user_info_fk FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE transaction ADD CONSTRAINT transaction_sender_fk FOREIGN KEY (sender_id) REFERENCES users (id) on delete set  NULL;
ALTER TABLE transaction ADD CONSTRAINT transaction_receiver_fk FOREIGN KEY (receiver_project_id) REFERENCES project (id) on delete set  NULL;

ALTER TABLE payout_transaction ADD CONSTRAINT payout_transaction_fk FOREIGN KEY (transaction_id) REFERENCES transaction (id) on delete cascade;
ALTER TABLE payout_transaction ADD CONSTRAINT payout_transaction_method_fk FOREIGN KEY (payout_method_id) REFERENCES payout_method (id) on delete set  NULL;

ALTER TABLE subscription_transaction ADD CONSTRAINT subscription_transaction_subscription_fk FOREIGN KEY (subscription_id) REFERENCES subscription (id) on delete set  NULL;
ALTER TABLE subscription_transaction ADD CONSTRAINT subscription_transaction_fk FOREIGN KEY (transaction_id) REFERENCES transaction (id) on delete cascade;
ALTER TABLE donation_transaction ADD CONSTRAINT donation_transaction_fk FOREIGN KEY (transaction_id) REFERENCES transaction (id) on delete cascade;

ALTER TABLE payout_method ADD CONSTRAINT payout_creator_fk FOREIGN KEY (creator_id) REFERENCES creator (id) on delete cascade;

ALTER TABLE file ADD CONSTRAINT file_user_fk FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE message ADD CONSTRAINT message_sender_fk FOREIGN KEY (sender_id) REFERENCES users (id) ON DELETE CASCADE;
ALTER TABLE message ADD CONSTRAINT message_receiver_fk FOREIGN KEY (receiver_id) REFERENCES users (id) ON DELETE CASCADE;

ALTER TABLE payment_method ADD CONSTRAINT payment_method_user_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;
ALTER TABLE stripe_card_payment_method ADD CONSTRAINT stripe_card_payment_method_fk FOREIGN KEY (payment_method_id) REFERENCES payment_method (id) ON DELETE CASCADE;

ALTER TABLE paypal_payout_method ADD CONSTRAINT paypal_payout_method_fk FOREIGN KEY (payout_method_id) REFERENCES payout_method (id) ON DELETE CASCADE;


ALTER TABLE goal ADD CONSTRAINT goal_project_fk FOREIGN KEY (project_id) REFERENCES project (id) ON DELETE CASCADE;

ALTER TABLE creator ADD CONSTRAINT creator_creator_rank_fk FOREIGN KEY (creator_rank_id) REFERENCES creator_rank (id) ON DELETE set  NULL;


CREATE UNIQUE INDEX email_unique_idx ON users (LOWER(email));

CREATE UNIQUE INDEX username_unique_idx ON users (LOWER(username));
