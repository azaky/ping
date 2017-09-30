# ping

Periodically hit urls, preventing apps to sleep in providers like [Heroku](https://blog.heroku.com/app_sleeping_on_heroku).
Use at your own risk.

## Running

```bash
X=https://example.com,https://example.org/home \
I=15000 \
go run main.go
```
