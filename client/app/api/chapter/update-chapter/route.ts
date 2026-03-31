import { apiFetch } from "@/lib/api";
import { NextResponse } from "next/server";
import { cookies } from "next/headers";
import { VocabularyInput } from "@/app/chapter/create/page";
import { ChapterInput } from "@/app/chapter/create/page";
export async function PUT(req: Request) {
  const endpoint = "/chapter/update-chapter"
  const body = await req.json();
  const cookieStore = await cookies()
  const accessToken = cookieStore.get("refreshToken")?.value;

  try {
    const data = apiFetch<{Chapter : ChapterInput  , List_Vocabulary : VocabularyInput[]}>(endpoint , {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `${accessToken}`,
        },
        body: JSON.stringify(body),
    } )
    const response = NextResponse.json({ success: true, data :  data });
    return response;

  } catch (error : any) {
    return NextResponse.json({ success: false, message: error.message || 'Create Chapter Error' }, { status: 500 });
  }
}