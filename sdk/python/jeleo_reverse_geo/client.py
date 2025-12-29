from __future__ import annotations

from dataclasses import dataclass
from typing import Any, Dict, Optional
import requests

DEFAULT_BASE_URL = "https://slimigeo.jeleo.zone.id"


@dataclass
class ReverseGeoClient:
    api_key: str
    base_url: str = DEFAULT_BASE_URL
    session: Optional[requests.Session] = None

    def __post_init__(self) -> None:
        if not self.api_key:
            raise ValueError("api_key is required")
        if self.session is None:
            self.session = requests.Session()

    def reverse_geo(self, lat: float, lng: float, timeout: Optional[float] = 30) -> Dict[str, Any]:
        if not isinstance(lat, (int, float)) or not isinstance(lng, (int, float)):
            raise ValueError("lat and lng must be numbers")
        url = self.base_url.rstrip('/') + '/reverse-geo'
        params = {
            'lat': lat,
            'lng': lng,
            'api_key': self.api_key,
        }
        r = self.session.get(url, params=params, timeout=timeout)
        try:
            r.raise_for_status()
        except requests.HTTPError as e:
            # Attach response text for easier debugging
            raise requests.HTTPError(f"{e} | body: {r.text}") from e
        return r.json()
