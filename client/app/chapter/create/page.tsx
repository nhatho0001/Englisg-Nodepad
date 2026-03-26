'use client'
import { Key, useState } from "react"
import CreateListVocabulary from "@/components/addChapter";


type VocabularyInput = {
  ID : Key
  OriginContent ?: string
  Description ?: string
}

export {type VocabularyInput}
export default function CrateNewChapter() {

  const [listVocabulary , setListVocabulary] = useState<VocabularyInput[]>([
    {
      ID : 1 ,
      OriginContent : "Home" , 
      Description : "Nhà"
    } , 
    {
      ID : 2 ,
      OriginContent : "Home" , 
      Description : "Nhà"
    }
  ])

  const editOriginContent = (id : Key , content : string ) => {
    var index = getVocabularyId(id)
    listVocabulary[index] = {
      ...listVocabulary[index] , 
      OriginContent : content
    }
    setListVocabulary(listVocabulary)
    console.log(listVocabulary)
  }

  const editDescriptionContent = (id : Key , content : string ) => {
    var index = getVocabularyId(id)
    listVocabulary[index] = {
      ...listVocabulary[index] , 
      Description : content
    }
    setListVocabulary(listVocabulary)
    console.log(listVocabulary)
  }
  const getVocabularyId = (id : Key)  => {
    return listVocabulary.findIndex((element) => {
      return element.ID == id
    })
  }
  const handleSubmit = () => {
    
  }
    
  return (
    <>
      <form onSubmit={handleSubmit}>
        <CreateListVocabulary listVocabulary = {listVocabulary} setListVocabulary = {setListVocabulary}></CreateListVocabulary>
      </form>
    </>
  )
}