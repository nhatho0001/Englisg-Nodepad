'use client';

import { useState, useEffect } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { apiFetch } from '@/lib/api';
import SendIcon from '@mui/icons-material/Send';
import { TextField, Button, Box, Typography, Container, Paper } from '@mui/material';
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
      // const data = await apiFetch<{ AccesToken: string; RefreshToken: string }>('/user/login', {
      //   method: 'POST',
      //   headers: {
      //     'Content-Type': 'application/json',
      //   },
      //   body: JSON.stringify({ email, password }),
      // });
      const res = await fetch('/api/user/login', {
        method: "POST",
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });
      res.ok ? router.push('/chapter') : setError('Login failed. Please try again.');;
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

          <TextField
            fullWidth
            id="standard-search"
            label="Email"
            type="email"
            variant="standard"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            disabled={loading}
            required
          />
          <TextField
            fullWidth
            id="standard-search"
            label="Password"
            type="password"
            variant="standard"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            disabled={loading}
            required
            sx={{ mt: 2, mb: 2 }}
          />
          <Box sx={{ display: 'flex', justifyContent: 'center' }}>
            {
              loading ? <Button loading variant="outlined" loadingPosition="start">
                Submit
              </Button> :
              (<Button type="submit" variant="outlined" endIcon={<SendIcon />}>
                Login
              </Button>) 
            }
          </Box>
        </form>

        <div className={styles.link}>
          Don’t have an account? <Link href="/user/sign-up">Sign Up</Link>
        </div>
      </div>
    </div>
  );
}