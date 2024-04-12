# Vugopo
My strongly opinionated, fullstack template for web applications

## Features
-  Make with Go, Typescript and Vue
-  Simple authentication based in `jsonwebtokens`
-  Postgres database integration
-  Deploy the application in any environment that supports docker containers

## Why

I needed an easy-to-use template to easily develop and deploy single-page applications with high performance and use it to create back office applications.

## Try it now!

### Clone to local

If you prefer to do it manually with the cleaner git history

```bash
git clone git@github.com:bio426/vugopo.git <your-directory-name>
cd <your-directory-name>
go get
cd <your-directory-name>/ui
npm i
```

### Set up your environment variables for development

Create a file called `env.sh` with the same content as `env.example.sh` and replace the required environment variables

### Run database migrations

You can use the [tern](https://github.com/jackc/tern) tool to run the migrations inside the `migrations` directory

## Checklist

When you use this template, try follow the checklist to update your info properly

- [ ] Create the bash file and set your variables in `env.sh`
- [ ] Change the module name in `go.mod`
- [ ] Change the favicon in `ui/public`
- [ ] Remove the `.git` folder
- [ ] Initialize your own git repository

And, enjoy :)

## Usage

### Development

1. Start backend server

```bash
# in project root
go run main.go
```
2. Start frontend server

```bash
# in ui/ directory
npm run dev
```

### Docker Production Build

First, build the vitesse image by opening the terminal in the project's root directory.

```bash
docker build . -t vugopo:latest
```

Run the image and specify port mapping with the `-p` flag.

```bash
docker run --rm -it -p 8080:80 \
	-e AUTH_JWT_SECRET="your-secret"
	-e AUTH_COOKIE_NAME="your-token-name"
	-e PG_HOST="your-db-host"
	-e PG_PORT="your-db-port"
	-e PG_USER="your-db-user"
	-e PG_PASSWORD="your-db-password"
	-e PG_DATABASE="your-db-database"
	vugopo:latest
```
