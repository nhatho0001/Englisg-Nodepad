import { apiFetch } from "@/lib/api";
import { NextResponse } from "next/server";
import { cookies } from "next/headers";
import { VocabularyInput } from "@/app/chapter/create/page";
import { ChapterInput } from "@/app/chapter/create/page";
export async function  DELETE(req: Request) {
  const userId = req.url
  const url_page = new URL(req.url);
  const id: string | null = url_page.searchParams.get("id");
  if(!id) {
    return NextResponse.json({ success: false, message: 'Parse param is faild' }, { status: 500 });
  }
  const params_object = {
    'id': String(id),
  };
  const queryString = new URLSearchParams(params_object).toString();
  var url = `/chapter/delete?${queryString}`;
  const cookieStore = await cookies()
  const accessToken = cookieStore.get("refreshToken")?.value;
  console.log(url)
  try {
    const list_data = await apiFetch<{sucess : boolean}>(url , {
      method: 'DELETE',
      headers: {
      'Content-Type': 'application/json',
      'Authorization': `${accessToken}`,
      },
  })
    if(list_data) {
      const response = NextResponse.json({ success: true, data :  list_data });
      return response;
    }else {
      return NextResponse.json({ success: false,   message:  'Delete is faild' });
    }
  }catch (error : any) {
    return NextResponse.json({ success: false, message: error.message || 'Create Chapter Error' }, { status: 500 });
  }
  
}