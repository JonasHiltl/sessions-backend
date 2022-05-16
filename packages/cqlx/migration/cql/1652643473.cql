CREATE TABLE IF NOT EXISTS friend_relations (
    user_id text,
    friend_id text,
    accepted boolean,
    requested_at timestamp,
    accepted_at timestamp,
    PRIMARY KEY (user_id, accepted)
) WITH transactions = { 'enabled' : true };

CREATE UNIQUE INDEX IF NOT EXISTS friend_relation_by_friend_id ON parties(id);