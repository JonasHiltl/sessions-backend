CREATE TABLE IF NOT EXISTS comments_by_party (
  id timeuuid,
  party_id text,
  author_id text,
  body text,
  PRIMARY KEY (party_id, comment_id)
) WITH CLUSTERING ORDER BY (comment_id DESC);

CREATE MATERIALIZED VIEW comments_by_user AS 
  SELECT * FROM comments_by_party
  WHERE author_id IS NOT NULL AND comment_id IS NOT NULL AND party_id IS NOT NULL
  PRIMARY KEY (author_id, comment_id, party_id)
WITH CLUSTERING ORDER BY (comment_id DESC);

CREATE TABLE IF NOT EXISTS replies (
  id timeuuid,
  comment_id timeuuid,
  author_id text,
  body text,
  PRIMARY KEY (comment_id, id)
) WITH CLUSTERING ORDER BY (id ASC);
