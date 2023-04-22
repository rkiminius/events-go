# events-go

`events-go` represents creation flow where users can create events with specified data.

## Prerequisites

In order to run this project you should have `mongodb` and `go` installed on your local setup. Or just use `docker` instead.

## Installation

### Local setup
1. Update `.env` file with valid `mongodb` host and port.
2. Build go application `go build -o envets-go`.
3. Run app `./events-go`.

### Docker setup
Run following docker command to build and launch docker containers.
```
    docker-compose up -d
```
App with run on `8000` port by default.

In order to stop all running containers.
```
    docker-compose down
```

## Usage 
### POST http://localhost:8000/events
Required data
```json
{
  "name": "My Event",
  "date": "2023-04-30T14:00:00Z",
  "languages": ["English", "French"],
  "videoQuality": ["720p", "1080p"],
  "audioQuality": ["High", "Low"],
  "invitees": ["example1@gmail.com", "example1@ss.com"]
}
```
Optional data
```json
{
  ...,
  "description": "A short description of the event",
  "options": {
    "default_video_quality": "1080p",
    "default_audio_quality": "High"
  }
}
```