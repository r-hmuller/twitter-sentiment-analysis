#TWITTER SEARCH AND SENTIMENT ANALYSIS 

This project is designed to get tweets, using Twitter's API
and use AWS Comprehend to identify if the tweet is positive/negative/neutral
and store the tweets on a database (postgres only).

To use it, you must inform your Twitter's credentials using env var.
You must also use AWS credentials (by default, it's located on ~/.aws/config)

## TO-DO

- [ ] Allow use others cloud providers to analyse the sentiment from tweets
- [ ] Allow use others databases (Mongodb)
- [ ] Upgrade API to get all stored tweets for a term, clustering by sentiment (POSITIVE/NEGATIVE) 