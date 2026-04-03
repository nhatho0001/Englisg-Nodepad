'use client';

import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import Link from 'next/link';

export default function Navbar() {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            <Link href="/" style={{ color: 'inherit', textDecoration: 'none' }}>
              Learning English
            </Link>
          </Typography>
          <Button color="inherit" component={Link} href="/">
            Home
          </Button>
          <Button color="inherit" component={Link} href="/chapter">
            List Chapter
          </Button>
          <Button color="inherit" component={Link} href="/chapter/create">
            Create Chapter
          </Button>
          <Button color="inherit" component={Link} href="/user/login">
            Login
          </Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
