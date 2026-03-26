import { apiFetch } from "@/lib/api";
import { cookies } from "next/headers";
import ComponetItemChapter from "./itemChapter";
import List from '@mui/material/List';
import { Key } from "react";

type ItemChapter = {
  ID : Key
  Title ?:  string , 
  Status :  string , 
  Body : string,
}

export {type ItemChapter}


export default function ListChapter({list_data} : {list_data : ItemChapter[]}) {
  return (
    <div className="ListItemChapter">
      <List sx={{ width: '100%', maxWidth: 720, bgcolor: 'background.paper' }}>
        {list_data.map((element) => {
          return <ComponetItemChapter key={element.ID} item={element}></ComponetItemChapter>
        })}
      </List>
    </div>
  )
}