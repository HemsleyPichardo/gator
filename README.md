# gator
An RSS feed aggregator in Go. We'll call it "Gator", you know, because aggreGATOR 🐊. 

Postgres and Go are needed to run this program. A json file in the user's home folder is used to track who is logged in and their connection to the PostgreSQL database. 

{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}

Postgres is an open-source database and was installed through the package manager. It needs to be started as a service on your machine. It runs on localhost:5432. 

Gator is a CLI application and was installed with Go from github.com/pressly/goose/v3/cmd/goose@latest

Valid commands: 
- `gator login <username>` - log in as an existing user
- `gator register <username>` - register a new user
- `gator reset` - wipes all users from the database
- `gator users` - list all registered users
- `gator agg <time>` - starts the aggregator. It runs continuously for a given duration
- `gator addfeed <url>` - add a new rss feed (requires login)
- `gator feeds` - list all feeds
- `gator follow <url>` - lets user follow an rss feed (requires login)
- `gator following` - list all feeds user is following (requires login)
- `gator unfollow <url>` - removes a feed-follow relation for the current user (requires login)
- `gator browse` - displays posts from feeds the current user follows. (requires login)
