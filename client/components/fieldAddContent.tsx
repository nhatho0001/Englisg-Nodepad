import TextField from '@mui/material/TextField';
import {VocabularyInput} from '@/app/chapter/create/page'
import { Key } from "react"
import { Box, Paper, Typography, Button } from '@mui/material';

export default function FieldAddContent({element, index, editOriginContent, editDescriptionContent, deleteVocabulary } :  {
    element: VocabularyInput, 
    index: number,
    editOriginContent: (id: Key, content: string) => void,  
    editDescriptionContent: (id: Key, content: string) => void,
    deleteVocabulary?: (id: Key) => void
}) {
  return (
    <Paper 
      variant="outlined" 
      sx={{ 
        p: 3, 
        borderRadius: 2, 
        position: 'relative',
        display: 'flex',
        flexDirection: 'column',
        gap: 2,
        bgcolor: 'background.default',
        transition: 'all 0.2s',
        '&:hover': { borderColor: 'primary.main', boxShadow: '0 4px 12px rgba(0,0,0,0.05)' }
      }}
    >
      <Box display="flex" justifyContent="space-between" alignItems="center" mb={1}>
        <Typography variant="subtitle2" color="text.secondary" fontWeight="bold">
          Vocabulary #{index + 1}
        </Typography>
        {deleteVocabulary && (
            <Button size="small" color="error" onClick={() => deleteVocabulary(element.ID)}>
                Delete
            </Button>
        )}
      </Box>
      <Box display="flex" flexDirection={{ xs: 'column', sm: 'row' }} gap={2}>
        <TextField
          fullWidth
          label="English Word / Phrase"
          variant="outlined"
          value={element.OriginContent || ''}
          onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
            editOriginContent(element.ID, event.target.value);
          }}
        />
        <TextField
          fullWidth
          label="Translation / Description"
          variant="outlined"
          value={element.Description || ''}
          onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
            editDescriptionContent(element.ID, event.target.value);
          }}
        />
      </Box>
    </Paper>
  )
}