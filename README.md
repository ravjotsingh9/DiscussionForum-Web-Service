## Discussion Forum Web Service
It is a simple web service to support posting a comment in a forum and getting comments back.

### Running
Use `dep ensure` to download all the dependecies. Note: you need to set up go env.
Then use `docker-compose up -d --build` to run the web-service and the postgres databse.

### Endpoints

#### /topic/ [POST]
For new thread, expects request as:
```
    {
        "id": "",
        "content": "This is the content of the comment.",
        "pid": "",
        "tid": ""
    }
```

For comment on a thread, request as:
```
    {
        "id": "",
        "content": "This is the content of the comment.",
        "pid": "<id of the thread topic>",
        "tid": "<id of the thread>"
    }
```
For reply to a comment on a thread, request as:
```
    {
        "id": "",
        "content": "This is the content of the comment.",
        "pid": "<id of the comment>",
        "tid": "<id of the thread>"
    }
```

#### /getTopic/{commentID} [GET]
Should reveice a list of related topic, comments, reply.
