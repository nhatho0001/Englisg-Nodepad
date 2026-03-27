import { VocabularyInput } from '@/app/chapter/create/page'
import { Box, Stack } from '@mui/material'
import { Dispatch, SetStateAction } from 'react'
import { Key } from "react"
import FieldAddContent from './fieldAddContent';

export default function CreateListVocabulary({ listVocabulary, setListVocabulary }: { listVocabulary: VocabularyInput[], setListVocabulary: Dispatch<SetStateAction<VocabularyInput[]>> }) {
    const editOriginContent = (id: Key, content: string) => {
        setListVocabulary(prev => prev.map(item =>
            item.ID === id ? { ...item, OriginContent: content } : item
        ));
    }

    const editDescriptionContent = (id: Key, content: string) => {
        setListVocabulary(prev => prev.map(item =>
            item.ID === id ? { ...item, Description: content } : item
        ));
    }

    const deleteVocabulary = (id: Key) => {
        setListVocabulary(prev => prev.filter(item => item.ID !== id));
    }

    return (
        <Stack spacing={3}>
            {listVocabulary.map((element, index) => (
                <FieldAddContent
                    key={element.ID}
                    index={index}
                    element={element}
                    editOriginContent={editOriginContent}
                    editDescriptionContent={editDescriptionContent}
                    deleteVocabulary={deleteVocabulary}
                />
            ))}
        </Stack>
    )
}