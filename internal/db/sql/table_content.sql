CREATE TABLE Users
(
    user_id        CHAR(36) PRIMARY KEY,
    last_posted_at TIMESTAMP NULL,
    dark_mode      bool           NOT NULL,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE PostSchedules
(
    post_schedule_id char(36) PRIMARY KEY,
    day_of_week      ENUM ('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday') NOT NULL,
    user_id          char(36)                                                                            NOT NULL,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users (user_id) ON DELETE CASCADE
);

CREATE TABLE Posts
(
    post_id    char(36) PRIMARY KEY,
    title      VARCHAR(20)  NOT NULL,
    content    VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE Tags
(
    tag_id char(36) PRIMARY KEY,
    name   VARCHAR(10) UNIQUE NOT NULL
);

CREATE TABLE PostTags
(
    post_id char(36),
    tag_id  char(36),
    PRIMARY KEY (post_id, tag_id),
    FOREIGN KEY (post_id) REFERENCES Posts (post_id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES Tags (tag_id) ON DELETE CASCADE
);