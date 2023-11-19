# Wind 💨: A REST API to get your discord status

Wind is a simple REST API that allows users to get their discord
status through API instead of Websockets. It's a great service
for showing your status on a portfolio or something.

## Endpoints 🛣️
### `/presence`

This endpoint requires a `Authentication` header with the API key
that can be generated through the bot using `/get-api`. The API key
can get the status of the user that generated the API key for security purposes.

It also requires and `id` parameter with the id of the discord user
that wants to retrive the status information.

**Example Response:**
```
{
    user: {
        id: "the id of the user",
        //other fields would probably be empty
    },
    status: "online" | "idle" | "dnd" | "invisible" | "offline",
    activities: [
        name: "name of the activity",
        type: 4, //see bellow
        url: "",
        created_at: Date,
        application_id: "",
        state: ""
        timestamps: {
            "start": Date,
            "end": Date
        },
        emoji: {},
        party: {},
        assets: {},
        secrets: {},
    ],
    client_status: {
        "desktop": "online",
        "mobile": "",
        "web": ""
    }
}
```

**Activity Types:**
```
Game      = 0
Streaming = 1
Listening = 2
Watching  = 3
Custom    = 4
Competing = 5
```


## Installation 🚀

1. Pull the repository to your local machine using `git pull github.com/x1xo/Wind`

2. You need to change the `.env.example` file to `.env` and configure it
to your own liking. See bellow for more information on `.env` file.

*Note: If you change the port in `.env`, you also need to change it in Dockerfile*

3. You need to build the Docker image using `docker build . -t wind`

4. Run the image usin `docker run -itd --env-file .env wind`

## Environment Variables 🔐

`SERVER_IP` is the address that the server is bound to:
ex. SERVER_IP=127.0.0.1
`SERVER_PORT` is the port that the server is listening to:
ex. SERVER_PORT=3000

`LIMITER_MAX_REQUESTS` is the maximum requests that a client can send
in the duration that is configured in `LIMITER_DURATION`
ex. LIMITER_MAX_REQUESTS=100
ex. LIMITER_DURATION=3000000 #in miliseconds (5minutes)
*In this example the user can send maximum 100 requests in 5 minutes*

`DISCORD_TOKEN` is the discord bot's token from `discord.com`
ex. DISCORD_TOKEN=your-discord-token-here
`DISCORD_GUILD_ID` is the server where the bot is invited and 
all the users that want to use the API.
ex. DISCORD_GUILD_ID=735949272620466246

`DATABASE_TYPE` is the database of your choice:
ex. DATABASE_TYPE="sqlite" #or redis (not yet implemented)
`DATABASE_URL` is the path to the sqlite db if you chose sqlite
or the redis uri if you chose redis.
DATABASE_URL="sqlite.db"