'use client';

import { useState, useEffect } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { apiFetch } from '@/lib/api';
import SendIcon from '@mui/icons-material/Send';
import { TextField, Button, Box, Typography, Container, Paper } from '@mui/material';
import styles from '../user.module.css';

export default function SignUpPage() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);
  const [mounted, setMounted] = useState(false);
  const router = useRouter();

  useEffect(() => {
    setMounted(true);
  }, []);

  if (!mounted) {
    return null;
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      await apiFetch('/user/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });

      // Redirect to login after successful registration
      router.push('/user/login');
    } catch (err: any) {
      setError(err.message || 'Sign up failed. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.card}>
        <h2 className={styles.title}>Sign Up</h2>

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
          Already have an account? <Link href="/user/login">Login</Link>
        </div>
      </div>
    </div>
  );
}