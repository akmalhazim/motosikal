# Motosikal.app

Motosikal.app is a survey tool used for a university research project. This repository contains data collection features.

Android (Device) -> Motosikal.app API -> Dashboard

The Android device and dashboard connects to the API via websocket for real-time activities.

This repository is available live at [https://api.motosikal.app](https://api.motosikal.app).

GitHub Actions is in-charge to ensure the Docker image is always updated with latest code changes. However, deployment is not yet automated.

## Installation

Use [Docker](https://www.docker.com) to run this API service.

```bash
docker run -p 8000 ghcr.io/akmalhazim/motosikal
```

## Usage

There's plenty room of changes. Here's just some draft on what is ready to consume.

```bash
GET /ws
GET /devices
POST /devices
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
