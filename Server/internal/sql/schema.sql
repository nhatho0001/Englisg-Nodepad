-- ========================
-- USERS
-- ========================
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  email VARCHAR NOT NULL,
  hashed_password TEXT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS unique_email_active
ON users(email)
WHERE deleted_at IS NULL;


-- ========================
-- CHAPTERS
-- ========================
CREATE TABLE IF NOT EXISTS chapters (
  id SERIAL PRIMARY KEY,
  title VARCHAR,
  body TEXT,
  user_id INTEGER NOT NULL,
  status VARCHAR,
  created_at TIMESTAMP DEFAULT NOW(),

  CONSTRAINT fk_chapters_user
  FOREIGN KEY (user_id)
  REFERENCES users(id)
  ON DELETE CASCADE
);

COMMENT ON COLUMN chapters.body IS 'Description for chapter';


-- ========================
-- VOCABULARY
-- ========================
CREATE TABLE IF NOT EXISTS vocabulary (
  id SERIAL PRIMARY KEY,
  chapter_id INTEGER,
  origin_content TEXT,
  description TEXT,
  practice_time TIMESTAMP,
  created_at TIMESTAMP DEFAULT NOW(),

  CONSTRAINT fk_vocab_chapter
  FOREIGN KEY (chapter_id)
  REFERENCES chapters(id)
  ON DELETE CASCADE
);


-- ========================
-- REFRESH TOKENS
-- ========================
CREATE TABLE IF NOT EXISTS refresh_tokens (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  hashed_token TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  expires_at TIMESTAMP,

  CONSTRAINT fk_refresh_user
  FOREIGN KEY (user_id)
  REFERENCES users(id)
  ON DELETE CASCADE
);


-- ========================
-- MANAGE FILE (polymorphic)
-- ========================
CREATE TABLE IF NOT EXISTS manage_file (
  id SERIAL PRIMARY KEY,
  alt VARCHAR,
  link TEXT,
  type_file VARCHAR,

  owner_id INTEGER NOT NULL,
  owner_type VARCHAR NOT NULL,

  created_at TIMESTAMP DEFAULT NOW(),

  CONSTRAINT chk_owner_type
  CHECK (owner_type IN ('users', 'chapters', 'vocabulary'))
);

CREATE INDEX IF NOT EXISTS idx_manage_file_owner
ON manage_file(owner_id, owner_type);