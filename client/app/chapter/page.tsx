import ListChapter from "@/components/listChapter"
import { Grid } from "@mui/material";
import Box from '@mui/material/Box';
import { apiFetch } from "@/lib/api";
import { use } from "react";
import { cookies } from "next/headers";
export default async function PageListChapter() {
    const cookieStore = await cookies();
    const accessToken = cookieStore.get("refreshToken")?.value;
    const response = await apiFetch<{ data: any[] }>("/chapter/list-chapter", {
        method: 'GET',
        headers: {
            'Authorization': `${accessToken}`,
        },
    })

    return (
        <Box
            sx={{
                minWidth: 500,
                minHeighth: "100vh",
                borderRadius: 1,
                bgcolor: 'primary.main',
                p: 2,
            }}>
            <h1>
                Your Chapter
            </h1>
            <ListChapter list_data={response.data}></ListChapter>
        </Box>
    )
}