const BASE_URL = "http://localhost:8080"; // TODO: to get from env

// TODO: implement a more robust solution
export class HttpClient {
  public static async request(
    method: "GET" | "POST" | "PUT" | "DELETE",
    endpoint: string,
    body?: object,
    customConfig?: RequestInit,
  ) {
    return fetch(`${BASE_URL}/${endpoint}`, {
      method,
      ...customConfig,
      headers: { ...customConfig?.headers },
      body: body ? JSON.stringify(body) : undefined,
    }).then(async (response) =>
      response.ok
        ? await response.json()
        : Promise.reject(new Error(await response.text())),
    );
  }

  public static async get(endpoint: string) {
    return await this.request("GET", endpoint);
  }

  public static async post(endpoint: string, body: object) {
    return await this.request("POST", endpoint, body, {
      headers: {
        "Content-Type": "application/json",
      },
    });
  }

  public static async put(endpoint: string, body: object) {
    return await this.request("PUT", endpoint, body, {
      headers: {
        "Content-Type": "application/json",
      },
    });
  }

  public static async delete(endpoint: string) {
    return await this.request("DELETE", endpoint);
  }
}
