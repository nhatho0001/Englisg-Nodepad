import { Key, useState } from "react"

`use client`
type VocabularyInput = {
  ID : Key
  OriginContent : string
  Description : string
  
  
}
export default function CrateNewChapter() {

  const [listVocabulary , setListVocabulary] = useState<[]>()

  const handleSubmit = () => {
    
  }
    
  return (
    <>
      <form onSubmit={handleSubmit}>

      </form>
    </>
  )
}