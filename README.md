# Jeleo Reverse Geocoding API SDKs and Examples

Turn GPS coordinate lat,lng into clean, human‑readable locations. Production‑ready Reverse Geocoding API with SDKs, examples, and a 60‑second quickstart for JavaScript/TypeScript and Python.

---

## Quick links
- Sign up (primary): https://www.jeleo.zone.id

Example cURL:

```
curl --location 'https://slimigeo.jeleo.zone.id/reverse-geo?lat=48.239879&lng=9.216766&api_key=YOUR_API_KEY'
```

---

## Why Jeleo Reverse Geocoding
- Accurate reverse geocoding: lat,lng → country, region, city, and more.
- Fast response times and simple pricing.
- Simple REST interface and lightweight SDKs.
- Works anywhere your app runs: browser, Node.js, serverless, Python backends, ETL.

---

## Getting started in 60 seconds

1) Get your API Key and base URL
- Create an account: https://www.jeleo.zone.id
- Copy your API key and base URL from the dashboard.

2) Pick your language

### JavaScript/TypeScript

Install (Node 18+ recommended):

```
cd sdk/js
npm ci
```

Use (TypeScript example):

```ts
import { ReverseGeoClient } from './src';

async function main() {
  const client = new ReverseGeoClient({ apiKey: process.env.JELEO_API_KEY! });
  const result = await client.reverseGeo(48.239879, 9.216766);
  console.log(result);
}

main().catch(console.error);
```

Run example:

```
JELEO_API_KEY=YOUR_API_KEY node --loader ts-node/esm examples/js/quickstart.ts
```

Or use plain cURL (no SDK):

```
curl --location 'https://<BASE_URL>/reverse-geo?lat=48.239879&lng=9.216766&api_key=<YOUR_API_KEY>'
```

### Python

```
cd sdk/python
pip install -r requirements.txt
```

```py
from jeleo_reverse_geo.client import ReverseGeoClient
import os

client = ReverseGeoClient(api_key=os.environ['JELEO_API_KEY'])
print(client.reverse_geo(48.239879, 9.216766))
```

---

## API Reference (OpenAPI)
See `openapi.yaml` for a minimal OpenAPI definition and `postman/Jeleo-Reverse-Geo.postman_collection.json` for a ready‑to‑use Postman collection.

Endpoint
- GET https://<BASE_URL>/reverse-geo
- Query parameters: `lat` (number), `lng` (number), `api_key` (string)

---

## Pricing and plans
- Compare plans and start free: https://www.jeleo.zone.id/?utm_source=github&utm_medium=repo&utm_campaign=reverse-geo#pricing

---

## Examples
- JavaScript/TypeScript: `examples/js/quickstart.ts`
- Python: `examples/python/quickstart.py`

---

## Benchmarks
Automated latency checks via GitHub Actions (scheduled) with historical trends. Badge will update here once enabled.

Planned metrics:
- P50, P95 latency (ms)
- Availability (%)
- Error rate

---

## FAQ
- Q: Is there a rate limit?
  A: See the pricing page for current tier limits: https://www.jeleo.zone.id/?utm_source=github&utm_medium=repo&utm_campaign=reverse-geo#pricing
- Q: Where do I get support?
  A: Open a GitHub Issue or contact us via the website.
- Q: Is there a free plan?
  A: Yes — start free and scale as needed.

---

## Contributing
- Open Issues for bugs or feature requests (SDKs, examples, docs).
- PRs welcome — please include tests where possible.

---

## Roadmap
- Additional SDKs (Go)
- CLI tool for batch reverse‑geo
- Advanced Benchmarks dashboard

---

## Star this repo
If this saved you time, please star the repo — it helps other developers find Reverse Geocoding tools.
