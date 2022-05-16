CREATE TABLE IF NOT EXISTS parties (
    id text,
    user_id text,
    title text,
    is_public boolean,
    lat float,
    long float,
    street_address text,
    postal_code text,
    state text,
    country text,
    start_date timestamp,
    end_date timestamp,
    PRIMARY KEY (user_id, is_public, start_date)
) WITH CLUSTERING ORDER BY (is_public DESC, start_date DESC) 
AND transactions = { 'enabled' : true };

CREATE UNIQUE INDEX IF NOT EXISTS parties_by_id ON parties(id);

CREATE TABLE IF NOT EXISTS favorite_parties (
    user_id text,
    party_id text,
    favorited_at timestamp,
    PRIMARY KEY (user_id, party_id, favorited_at)
) WITH CLUSTERING ORDER BY (party_id DESC, favorited_at DESC) 
AND transactions = { 'enabled' : true };

CREATE UNIQUE INDEX favorite_parties_by_party ON favorite_parties(party_id);
