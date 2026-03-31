'use client'
import { Key, useState } from "react"
import CreateListVocabulary from "@/components/addChapter";
import { Container, Paper, Typography, Box, Button, Divider } from "@mui/material";
import { apiFetch } from "@/lib/api";
import { FormControl, InputLabel, OutlinedInput, FormHelperText } from '@mui/material';
import { TextField, Stack } from '@mui/material';
import { VocabularyInput, ChapterInput } from '@/app/chapter/create/page'
export default function ListVocabulary({list_data , chapter} : {list_data : VocabularyInput[] , chapter :  ChapterInput}) {
  console.log(list_data)
  const [newChapter , setNewChapter] =  useState<ChapterInput>(chapter) ; 
  const [listVocabulary, setListVocabulary] = useState<VocabularyInput[]>(list_data)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const response = await fetch("/api/chapter/update-chapter" ,  {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({Chapter : newChapter , List_Vocabulary : listVocabulary })
    })
  }

  const addNewVocabulary = () => {
    setListVocabulary([
      ...listVocabulary,
      { ID: Date.now(), OriginContent: "", Description: "" }
    ]);
  }

  return (
    <Container maxWidth="md" sx={{ py: { xs: 4, md: 8 }, mb: 4 }}>
      <Paper elevation={4} sx={{ borderRadius: 3, overflow: 'hidden' }}>
        <Box sx={{ p: { xs: 3, md: 4 }, background: 'linear-gradient(135deg, #1976d2 0%, #1565c0 100%)', color: 'primary.contrastText' }}>
          <Typography variant="h4" component="h1" fontWeight="bold">
            Create Chapter
          </Typography>
          <Typography variant="subtitle1" sx={{ opacity: 0.9, mt: 1 }}>
            Add new vocabulary items to your lesson.
          </Typography>
        </Box>

        <Box component="form" onSubmit={handleSubmit} sx={{ p: { xs: 2, md: 4 } }}>
          <Stack spacing={2} sx={{ width: '100%', py: { xs: 2, md: 2 }}}>
            {/* Title Field */}
            <TextField 
              label="Title" 
              variant="outlined" 
              defaultValue={newChapter?.Title}
              onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                setNewChapter(prep => {
                  return {...prep ,  Title :  event.target.value}
                })
              }}
              fullWidth 
            />

            {/* Description Field */}
            <TextField 
              label="Description" 
              variant="outlined" 
              multiline 
              rows={4} 
              onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                setNewChapter(prep => {
                  return {...prep ,  Body :  event.target.value}
                })
              }}
              defaultValue={newChapter?.Body}
              fullWidth 
            />
          </Stack>
          <CreateListVocabulary listVocabulary={listVocabulary} setListVocabulary={setListVocabulary} />

          <Divider sx={{ my: 3 }} />

          <Box display="flex" justifyContent="space-between" alignItems="center" flexWrap="wrap" gap={2}>
            <Button
              variant="outlined"
              color="secondary"
              onClick={addNewVocabulary}
              sx={{ borderRadius: 2, fontWeight: 'bold' }}
            >
              + Add Word
            </Button>
            <Button
              type="submit"
              variant="contained"
              color="primary"
              size="large"
              sx={{ px: 4, py: 1.5, borderRadius: 2, fontWeight: 'bold' }}
            >
              Save Chapter
            </Button>
          </Box>
        </Box>
      </Paper>
    </Container>
  )
}