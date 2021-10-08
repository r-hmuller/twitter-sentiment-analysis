CREATE TABLE tweets (
    ID SERIAL CONSTRAINT pk_id_tweet PRIMARY KEY,
    IdTweet BIGINT NOT NULL,
    Tweet varchar(300) NOT NULL,
    Author varchar(200) NOT NULL,
    SentimentType varchar(25) NOT NULL,
    SentimentScore NUMERIC(7,6),
    CreatedAt timestamp
);