CREATE TABLE IF NOT EXISTS blog_tags (
    tag_id CHAR(36) NOT NULL,
    blog_id CHAR(36),

    PRIMARY KEY (tag_id,blog_id), -- Composite Key

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Foreign Key (tag_id) REFERENCES tags(id) ON UPDATE CASCADE ON DELETE CASCADE,
    Foreign Key (blog_id) REFERENCES blogs(id) ON UPDATE CASCADE ON DELETE CASCADE
);