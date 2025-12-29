import os
import sys
from pathlib import Path

# Ensure the SDK is on sys.path without requiring installation
repo_root = Path(__file__).resolve().parents[2]
sdk_python_path = repo_root / "sdk" / "python"
if str(sdk_python_path) not in sys.path:
    sys.path.insert(0, str(sdk_python_path))

from jeleo_reverse_geo.client import ReverseGeoClient

if __name__ == "__main__":
    api_key = os.environ.get("JELEO_API_KEY")
    if not api_key:
        raise RuntimeError("Set JELEO_API_KEY environment variable")

    client = ReverseGeoClient(api_key=api_key)
    print(client.reverse_geo(48.239879, 9.216766))
