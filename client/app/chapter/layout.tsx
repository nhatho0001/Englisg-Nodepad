import { Box } from "@mui/material";

export default function ChapterLayout({ children }: { children: React.ReactNode }) {
    return (
        <Box
            sx={{
                minHeight: "100vh",
                bgcolor: 'background.default',
                py: { xs: 4, md: 8 },
                px: { xs: 2, md: 4 },
                display: 'flex',
                justifyContent: 'center'
            }}
        >
            {children}
        </Box>
    );
}