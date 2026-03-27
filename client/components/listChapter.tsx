import { apiFetch } from "@/lib/api";
import { cookies } from "next/headers";
import ComponetItemChapter from "./itemChapter";
import List from '@mui/material/List';
import Divider from '@mui/material/Divider';
import { Key, Fragment } from "react";

type ItemChapter = {
  ID: Key
  Title?: string,
  Status: string,
  Body: string,
}

export { type ItemChapter }


export default function ListChapter({ list_data }: { list_data: ItemChapter[] }) {
  return (
    <div className="ListItemChapter">
      <List sx={{ width: '100%', bgcolor: 'background.paper' }}>
        {list_data && list_data.length > 0 ? (
          list_data.map((element, index) => (
            <Fragment key={element.ID}>
              <ComponetItemChapter item={element}></ComponetItemChapter>
              {index < list_data.length - 1 && <Divider component="li" sx={{ my: 1 }} />}
            </Fragment>
          ))
        ) : (
          <p style={{ textAlign: "center", fontStyle: "italic", opacity: 0.6 }}>No chapters found.</p>
        )}
      </List>
    </div>
  )
}