import { ReverseGeoClient } from '../../sdk/js/src/index.ts';

async function main() {
  const apiKey = process.env.JELEO_API_KEY;
  if (!apiKey) throw new Error('Set JELEO_API_KEY env var');

  const client = new ReverseGeoClient({ apiKey });
  const result = await client.reverseGeo(48.239879, 9.216766);
  console.log(JSON.stringify(result, null, 2));
}

main().catch((err) => {
  console.error(err);
  process.exitCode = 1;
});
