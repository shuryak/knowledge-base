export default function sendRequest(method, route, body = null) {
  const url = `${process.env.REACT_APP_API_HOST}:${process.env.REACT_APP_API_PORT}/api/${route}`;

  const request = {
    method: method,
    headers: {
      'Content-Type': 'application/json',
    }
  }

  if (body != null) request.body = JSON.stringify(body);

  return fetch(url, request)
    .then(async (response) => ({
      status: response.status,
      data: await response.json(),
    }))
    .catch(() => {});
}
