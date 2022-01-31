-- users
INSERT INTO user_profile (id, first_name, last_name, email, gender, created_at)
VALUES (1, 'Mariam', 'Ali', 'm.ali@gmail.com', 'FEMALE', '2020-11-24 23:42:47.000000');

INSERT INTO user_profile (id, first_name, last_name, email, gender, created_at)
VALUES (2, 'Joe', 'James', 'j.james@gmail.com', 'MALE', '2020-11-24 23:42:47.000000');

INSERT INTO user_profile (id, first_name, last_name, email, gender, created_at)
VALUES (3, 'Jamila', 'Ahmed', 'jamila@gmail.com', 'FEMALE', '2020-11-24 23:42:47.000000');

INSERT INTO user_profile (id, first_name, last_name, email, gender, created_at)
VALUES (4, 'Alex', 'Smith', 'alex2000@gmail.com', 'MALE', '2020-11-24 23:42:47.000000');

-- accounts
INSERT INTO youtube_account (id, created_at) VALUES (1, '2020-11-24 23:44:36.000000');
INSERT INTO youtube_account (id, created_at) VALUES (2, '2020-11-24 23:00:36.000000');
INSERT INTO youtube_account (id, created_at) VALUES (4, '2020-11-24 10:44:36.000000');

-- youtube channels
INSERT INTO youtube_channel (id, youtube_account_id, channel_name, created_at)
VALUES (1, 1, 'MariamBeauty', '2020-11-24 23:47:05.385073');

INSERT INTO youtube_channel (id, youtube_account_id, channel_name, created_at)
VALUES (2, 2, 'JoeTech', '2020-11-24 23:47:50.904706');

INSERT INTO youtube_channel (id, youtube_account_id, channel_name, created_at)
VALUES (3, 4, 'AlexTutorials', '2020-11-24 23:47:50.904706');

-- subscribers
INSERT INTO channel_subscriber (youtube_account_id, youtube_channel_id, created_at)
VALUES (1, 2, '2020-11-25 22:19:41.000000');

INSERT INTO channel_subscriber (youtube_account_id, youtube_channel_id, created_at)
VALUES (1, 3, '2020-11-25 22:19:58.000000');

INSERT INTO channel_subscriber (youtube_account_id, youtube_channel_id, created_at)
VALUES (4, 1, '2020-11-25 22:19:58.000000');

INSERT INTO channel_subscriber (youtube_account_id, youtube_channel_id, created_at)
VALUES (2, 1, '2020-11-25 22:19:58.000000');
