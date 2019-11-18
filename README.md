# strange-server

## Usage
`docker-compose up --build`

1. To restart application without running migrations set environment variable "WITH_MIGRATIONS=false"

## Routes
1. [POST] /post (see ./testdata/post.http)
2. [GET] /list - list of stored requests
3. [GET] /source_list - list of current sources (C.O.)
