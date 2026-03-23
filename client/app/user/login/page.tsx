'use client';

import { useState, useEffect } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { apiFetch } from '@/lib/api';
import styles from '../user.module.css';

export default function LoginPage() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);
  const [mounted, setMounted] = useState(false);
  const router = useRouter();

  console.log('LoginPage Render');

  useEffect(() => {
    setMounted(true);
  }, []);

  if (!mounted) {
    return null; // Or a simple loading skeleton to avoid hydration mismatch
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setLoading(true);
    try {
      const data = await apiFetch<{ AccesToken: string; RefreshToken: string }>('/user/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });

      localStorage.setItem('accessToken', data.AccesToken);
      localStorage.setItem('refreshToken', data.RefreshToken);

      router.push('/');
    } catch (err: any) {
      console.error('API Error:', err);
      setError(err.message || 'Login failed. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.card}>
        <h2 className={styles.title}>Login</h2>

        {error && <p className={styles.error}>{error}</p>}

        <form onSubmit={handleSubmit}>
          <input
            className={styles.input}
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            disabled={loading}
            required
          />

          <input
            className={styles.input}
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            disabled={loading}
            required
          />

          <button
            className={styles.button}
            type="submit"
            disabled={loading}
          >
            {loading ? 'Logging in...' : 'Login'}
          </button>
        </form>

        <div className={styles.link}>
          Don’t have an account? <Link href="/user/sign-up">Sign Up</Link>
        </div>
      </div>
    </div>
  );
}