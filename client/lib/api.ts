const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export async function apiFetch<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
  console.log(`${API_URL}${endpoint}`)
  const response = await fetch(`${API_URL}${endpoint}`, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options.headers,
    },
  });

  if (!response.ok) {
    const statusText = response.statusText;
    const statusCode = response.status;
    const bodyText = await response.text().catch(() => '');

    let errorMessage = 'Something went wrong';
    try {
      const errorData = JSON.parse(bodyText);
      errorMessage = errorData.error || errorData.message || errorMessage;
    } catch (e) {
      errorMessage = bodyText || statusText || errorMessage;
    }

    throw new Error(`API Error ${statusCode}: ${errorMessage}`);
  }

  return response.json();
}
