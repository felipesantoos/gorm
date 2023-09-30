CREATE DATABASE gorm;
\C gorm;

CREATE TABLE users (
    id INT,
    login VARCHAR(256),
    email VARCHAR(256),
    avatar TEXT,
    active BOOLEAN,
    last_login TIMESTAMP,
    created TIMESTAMP,
    updated TIMESTAMP,
    CONSTRAINT users_pk PRIMARY KEY (id)
);

INSERT INTO users (id, login, email, avatar, active, last_login, created, updated)
VALUES (1, 'user1', 'user1@example.com', 'user1_avatar.jpg', true, NOW(), NOW(), NOW());

INSERT INTO users (id, login, email, avatar, active, last_login, created, updated)
VALUES (2, 'user2', 'user2@example.com', 'user2_avatar.jpg', true, NOW(), NOW(), NOW());

INSERT INTO users (id, login, email, avatar, active, last_login, created, updated)
VALUES (3, 'user3', 'user3@example.com', 'user3_avatar.jpg', true, NOW(), NOW(), NOW());

INSERT INTO users (id, login, email, avatar, active, last_login, created, updated)
VALUES (4, 'user4', 'user4@example.com', 'user4_avatar.jpg', true, NOW(), NOW(), NOW());

INSERT INTO users (id, login, email, avatar, active, last_login, created, updated)
VALUES (5, 'user5', 'user5@example.com', 'user5_avatar.jpg', true, NOW(), NOW(), NOW());

INSERT INTO users (id, login, email, avatar, active, last_login, created, updated)
VALUES (6, 'user6', 'user6@example.com', 'user6_avatar.jpg', true, NOW(), NOW(), NOW());

INSERT INTO users (id, login, email, avatar, active, last_login, created, updated)
VALUES (7, 'user7', 'user7@example.com', 'user7_avatar.jpg', true, NOW(), NOW(), NOW());

INSERT INTO users (id, login, email, avatar, active, last_login, created, updated)
VALUES (8, 'user8', 'user8@example.com', 'user8_avatar.jpg', true, NOW(), NOW(), NOW());

INSERT INTO users (id, login, email, avatar, active, last_login, created, updated)
VALUES (9, 'user9', 'user9@example.com', 'user9_avatar.jpg', true, NOW(), NOW(), NOW());

INSERT INTO users (id, login, email, avatar, active, last_login, created, updated)
VALUES (10, 'user10', 'user10@example.com', 'user10_avatar.jpg', true, NOW(), NOW(), NOW());
