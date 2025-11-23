import { auth } from './firebase';

const BASE = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

async function withAuthHeaders() {
  const currentUser = auth.currentUser;
  const headers = { 'Content-Type': 'application/json' };
  if (currentUser) {
    const token = await currentUser.getIdToken();
    headers['Authorization'] = `Bearer ${token}`;
  }
  return headers;
}

export async function fetchForYou() {
  const headers = await withAuthHeaders();
  const res = await fetch(`${BASE}/foryou`, { headers });
  if (!res.ok) throw new Error(`Failed to fetch: ${res.status}`);
  return res.json();
}

export async function postComment(happeningId, text) {
  const headers = await withAuthHeaders();
  const res = await fetch(`${BASE}/happenings/${happeningId}/comments`, {
    method: 'POST',
    headers,
    body: JSON.stringify({ text })
  });
  if (!res.ok) throw new Error(`Failed to post comment: ${res.status}`);
  return res.json();
}