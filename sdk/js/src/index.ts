export interface ReverseGeoClientOptions {
  apiKey: string;
  baseUrl?: string; // e.g., https://slimigeo.jeleo.zone.id
  fetchImpl?: typeof fetch; // for testing / environments without global fetch
}

export interface ReverseGeoResponse {
  // The exact shape depends on server. We keep it open but typed as a record.
  [key: string]: unknown;
}

export class ReverseGeoClient {
  private apiKey: string;
  private baseUrl: string;
  private fetchFn: typeof fetch;

  constructor(opts: ReverseGeoClientOptions) {
    if (!opts?.apiKey) throw new Error("apiKey is required");
    this.apiKey = opts.apiKey;
    this.baseUrl = (opts.baseUrl ?? 'https://slimigeo.jeleo.zone.id').replace(/\/$/, '');
    this.fetchFn = opts.fetchImpl ?? (globalThis.fetch as typeof fetch);
    if (!this.fetchFn) {
      throw new Error('No fetch implementation available. Provide fetchImpl in options or use Node 18+/browser.');
    }
  }

  async reverseGeo(lat: number, lng: number, signal?: AbortSignal): Promise<ReverseGeoResponse> {
    if (!Number.isFinite(lat) || !Number.isFinite(lng)) {
      throw new Error('lat and lng must be finite numbers');
    }
    const url = new URL(this.baseUrl + '/reverse-geo');
    url.searchParams.set('lat', String(lat));
    url.searchParams.set('lng', String(lng));
    url.searchParams.set('api_key', this.apiKey);

    const res = await this.fetchFn(url.toString(), { method: 'GET', signal });
    if (!res.ok) {
      const text = await res.text().catch(() => '');
      throw new Error(`ReverseGeo request failed: ${res.status} ${res.statusText} ${text}`.trim());
    }
    const data = (await res.json()) as ReverseGeoResponse;
    return data;
  }
}
