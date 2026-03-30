import ListChapter from "@/components/listChapter"
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import Paper from '@mui/material/Paper';
import Typography from '@mui/material/Typography';
import { apiFetch } from "@/lib/api";
import { cookies } from "next/headers";

export default async function PageListChapter() {
    const cookieStore = await cookies();
    const accessToken = cookieStore.get("refreshToken")?.value;
    const response = await apiFetch<{ data: any[] }>("/chapter/list-chapter", {
        method: 'GET',
        headers: {
            'Authorization': `${accessToken}`,
        },
    });

    return (
        <Container maxWidth="md">
            <Paper
                elevation={6}
                sx={{
                    borderRadius: 3,
                    overflow: 'hidden',
                    bgcolor: 'background.paper',
                    mt: { xs: 2, md: 4 },
                    mb: { xs: 4, md: 8 }
                }}
            >
                <Box
                    sx={{
                        p: { xs: 3, md: 5 },
                        background: 'linear-gradient(135deg, #1976d2 0%, #1565c0 100%)',
                        color: 'primary.contrastText',
                        boxShadow: 'inset 0 -2px 10px rgba(0,0,0,0.1)'
                    }}
                >
                    <Typography
                        variant="h3"
                        component="h1"
                        fontWeight="800"
                        sx={{ fontSize: { xs: '2rem', md: '3rem' } }}
                        gutterBottom
                    >
                        Your Chapters
                    </Typography>
                    <Typography variant="subtitle1" sx={{ opacity: 0.9, fontSize: { xs: '1rem', md: '1.25rem' } }}>
                        Continue your English learning journey and track your progression.
                    </Typography>
                </Box>

                <Box sx={{ p: { xs: 2, md: 4 } }}>
                    <ListChapter list_data={response.data || []} />
                </Box>
            </Paper>
        </Container>
    );
}