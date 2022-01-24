# FTX Account API

## Functions:

* Enable Order service
* Disable Order service

ORDER_KEY = The plaintext key of the generated ORDER access pair.

The key is different from the auth key used for account requests.

### Enable Order service:

* Order service is enabled by default.
* Not required unless manually disabled before.

Request

```bash

```

Response

* Ok when enabled
* http error 505 internal error otherwise

### Disable order service:

* Disables order service
* Helps with scheduled maintenance i.e. disable order service before restart
* Order service will be enabled again by default after restart

Request

```bash

```

Response

* Ok when disabled
* http error 505 internal error otherwise