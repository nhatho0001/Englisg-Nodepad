import { Grid } from "@mui/material";
export default function ChapterLayout({ children }: { children: React.ReactNode }) {
    return <Grid
        container
        justifyContent="center"
        alignItems="center"
        minHeight="100vh"
        minWidth="300px"
    >
        {children}
    </Grid>
}