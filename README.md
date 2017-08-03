# DEA - Disposable Email Address Detector

DEA provides api to check email addresses against a list of known "trash domains"

## Usage

To check email run application and make a call to: http://localhost:8000/v1/check/[domain]

For example check email `exmaple@getnada.com` make call to http://localhost:8000/v1/check/getnada.com

You will get a response showing if domain is trash domain.

```json
{
    "domain": "getnada.com",
    "blocked": true
}
```
