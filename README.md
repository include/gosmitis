# README

This tool serves a RestAPI endpoint to send emails. Yes s simple as that.

### Config
Please edit and change ``config-apisendemail.tom``to match your email settings.

### Run

``go run apisendemail.go``

### Test
```
curl -X POST -d "{\"body\": \"that\", \"from\": \"foo\"}" http://127.0.0.1:8080/sendemail
```
