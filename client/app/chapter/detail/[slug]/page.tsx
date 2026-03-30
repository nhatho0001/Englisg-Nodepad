// import { Key, useState } from "react"
import { Container, Paper, Typography, Box, Button, Divider } from "@mui/material";
import { apiFetch } from "@/lib/api";
import { FormControl, InputLabel, OutlinedInput, FormHelperText } from '@mui/material';
import { TextField, Stack } from '@mui/material';
import { VocabularyInput, ChapterInput } from '@/app/chapter/create/page'
import { cookies } from "next/headers";
import ListVocabulary from "@/components/vocabularyOfChapter"
 

export default async function PageEditChapter({
  params,
}: {
  params: Promise<{ slug: string }>
}) {
  const { slug } = await params

  const params_object = {
      'id': String(slug),
  };
  const queryString = new URLSearchParams(params_object).toString();
  var url = `/chapter/list-vocabulary?${queryString}`;
  const cookieStore = await cookies()
  const accessToken = cookieStore.get("refreshToken")?.value;

  

  const list_data = await apiFetch<{Chapter : ChapterInput  , List_Vocabulary : VocabularyInput[]}>(url , {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `${accessToken}`,
    },
  })

  return <ListVocabulary list_data={list_data.List_Vocabulary} chapter={list_data.Chapter}></ListVocabulary>
}