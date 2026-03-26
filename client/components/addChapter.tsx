import {VocabularyInput} from '@/app/chapter/create/page'
import { Box } from '@mui/material'
import { Dispatch , SetStateAction ,  } from 'react'
import TextField from '@mui/material/TextField';
import { Key, useState } from "react"
export default function CreateListVocabulary({listVocabulary , setListVocabulary} : {listVocabulary : VocabularyInput[] ,setListVocabulary : Dispatch<SetStateAction<VocabularyInput[]>> }) {
    const editOriginContent = (id : Key , content : string ) => {
        var index = getVocabularyId(id)
        listVocabulary[index] = {
            ...listVocabulary[index] , 
            OriginContent : content
        }
        setListVocabulary(listVocabulary)
    }
    
    const editDescriptionContent = (id : Key , content : string ) => {
        var index = getVocabularyId(id)
        listVocabulary[index] = {
            ...listVocabulary[index] , 
            Description : content
        }
        setListVocabulary(listVocabulary)
    }
    const getVocabularyId = (id : Key)  => {
        return listVocabulary.findIndex((element) => {
            return element.ID == id
        })
    }
    return <Box>
        {listVocabulary.map((element) =>  {
            return <div key={element.ID} >
                <TextField
                    id="outlined-controlled"
                    label="Controlled"
                    defaultValue={element.OriginContent}
                    onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                        editOriginContent(element.ID , event.target.value);
                    }}
                />
                <TextField
                    id="outlined-controlled"
                    label="Controlled"
                    defaultValue={element.Description}
                    onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                        editDescriptionContent(element.ID , event.target.value);
                    }}
                />
            </div>
        })}
    </Box>
}