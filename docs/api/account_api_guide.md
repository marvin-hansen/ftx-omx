# FTX Account API

## Functions:

* Enable API service
* Disable API service
* Register API
* Remove API
* Set API Leverage
* Reset API Leverage

API_KEY = The plaintext key of the generated API access pair.

The key is different from the auth key used for order requests.

### Enable API service:

* API service is disabled by default.
* Enable it to add a FTX (sub) account API key

Request:

```bash
curl -X POST http://localhost/api/on -H "auth: API_KEY"
```

Response

* Ok when enabled
* http error 505 internal error otherwise

### Disable API service:

* API service should be disabled after adding FTX (sub) accounts to prevent any external access

* Request:

```bash
curl -X POST http://localhost/api/off -H "auth: API_KEY"
```

Response

* Ok when disabled
* http error 505 internal error otherwise

### Register API

* Adds a new FTX (sub) account API key
* Returns a neutral & unique ID
* All orders require this ID to send orders to this account
* API KEY & SECRET stored only fully encrypted and are only decrypted when sending a request to FTX

Parameter (all string):

* id - leave always empty;
* accountName - i.e. mTXSubAccount
* key - FTX API key
* secret FTX API secret

Payload

```json 
{
   "id":"",
   "accountName": "API_NAME",
   "key":    "FTX_API_KEY",
   "secret": "FTX_API_SECRET"
}
```

Request

```bash
curl -X POST http://localhost/api/create -H "auth: a9f4fd720fb842dc66a3adc9f44d362b" -H 'Content-Type: application/json' -d '{"id":"", "accountName": "API_NAME", "key":    "FTX_API_KEY", "secret": "FTX_API_SECRET"}'
```

Response:

* id (string) the unique API identifier
* http error 505 internal error otherwise

```bash
api id: fdgd253 
```

### Remove API:

* Deletes the FTX account API key & secret matching to the ID

Parameter:

* API ID (string)

Request

```bash
curl http://localhost/api/delete?id=fdgd253 -H "auth: API_KEY" 
```

Response

* Ok when disabled
* http error 505 internal error otherwise

### Set Account Leverage:

Parameter:

* API ID (string)
* Market / Token (string)
* Leverage (int) between 1 and 20

Payload

```json 
{
   "api_id":"API_ID",
   "market":"ETHUSD",
   "leverage":3
}
```

Request

```bash
curl -X POST http://localhost/api/setleverage -H 'Content-Type: application/json' -d '{"api_id":"API_ID","market":"ETHUSD","leverage":3}' -H "auth: API_KEY" 
```

Response

* Ok when set
* http error 505 internal error otherwise

### Reset Account Leverage:

Parameter:

* API ID (string)
* Market / Token (string)
* Leverage (int) 0 - will be ignored by OMX

Payload

```json 
{
   "api_id":"API_ID",
   "market":"ETHUSD",
   "leverage":0
}
```

```bash
curl -X POST http://localhost/api/resetleverage -H 'Content-Type: application/json' -d '{"api_id":"API_ID","market":"ETHUSD","leverage":0}' -H "auth: API_KEY" 
```

Response

* Ok when reset
* http error 505 internal error otherwise
